package iothub

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-03-31/devices"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/locks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iothub/parse"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/iothub/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceIotHubEndpointServiceBusQueue() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceIotHubEndpointServiceBusQueueCreateUpdate,
		Read:   resourceIotHubEndpointServiceBusQueueRead,
		Update: resourceIotHubEndpointServiceBusQueueCreateUpdate,
		Delete: resourceIotHubEndpointServiceBusQueueDelete,

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.EndpointServiceBusQueueID(id)
			return err
		}),

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.IoTHubEndpointName,
			},

			"resource_group_name": azure.SchemaResourceGroupName(),

			// TODO remove in 3.0
			"iothub_name": {
				Type:          pluginsdk.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ValidateFunc:  validate.IoTHubName,
				Deprecated:    "Deprecated in favour of `iothub_id`",
				ConflictsWith: []string{"iothub_id"},
			},

			"iothub_id": {
				Type: pluginsdk.TypeString,
				// TODO add Required: true in 3.0
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ValidateFunc:  validate.IotHubID,
				ConflictsWith: []string{"iothub_name"},
			},

			"connection_string": {
				Type:     pluginsdk.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *pluginsdk.ResourceData) bool {
					sharedAccessKeyRegex := regexp.MustCompile("SharedAccessKey=[^;]+")
					sbProtocolRegex := regexp.MustCompile("sb://([^:]+)(:5671)?/;")

					maskedNew := sbProtocolRegex.ReplaceAllString(new, "sb://$1:5671/;")
					maskedNew = sharedAccessKeyRegex.ReplaceAllString(maskedNew, "SharedAccessKey=****")
					return (new == d.Get(k).(string)) && (maskedNew == old)
				},
				Sensitive: true,
			},
		},
	}
}

func resourceIotHubEndpointServiceBusQueueCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).IoTHub.ResourceClient
	subscriptionId := meta.(*clients.Client).IoTHub.ResourceClient.SubscriptionID
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()
	subscriptionID := meta.(*clients.Client).Account.SubscriptionId

	endpointRG := d.Get("resource_group_name").(string)
	iotHubName := d.Get("iothub_name").(string)
	iotHubRG := endpointRG

	if iotHubName == "" {
		id, err := parse.IotHubID(d.Get("iothub_id").(string))
		if err != nil {
			return err
		}
		iotHubName = id.Name
		iotHubRG = id.ResourceGroup
	}

	id := parse.NewEndpointServiceBusQueueID(subscriptionId, iotHubRG, iotHubName, d.Get("name").(string))

	locks.ByName(iotHubName, IothubResourceName)
	defer locks.UnlockByName(iotHubName, IothubResourceName)

	iothub, err := client.Get(ctx, iotHubRG, iotHubName)
	if err != nil {
		if utils.ResponseWasNotFound(iothub.Response) {
			return fmt.Errorf("IotHub %q (Resource Group %q) was not found", iotHubName, iotHubRG)
		}

		return fmt.Errorf("loading IotHub %q (Resource Group %q): %+v", iotHubName, iotHubRG, err)
	}

	queueEndpoint := devices.RoutingServiceBusQueueEndpointProperties{
		ConnectionString: utils.String(d.Get("connection_string").(string)),
		Name:             utils.String(id.EndpointName),
		SubscriptionID:   utils.String(subscriptionID),
		ResourceGroup:    utils.String(endpointRG),
	}

	routing := iothub.Properties.Routing
	if routing == nil {
		routing = &devices.RoutingProperties{}
	}

	if routing.Endpoints == nil {
		routing.Endpoints = &devices.RoutingEndpoints{}
	}

	if routing.Endpoints.EventHubs == nil {
		queues := make([]devices.RoutingServiceBusQueueEndpointProperties, 0)
		routing.Endpoints.ServiceBusQueues = &queues
	}
	endpoints := make([]devices.RoutingServiceBusQueueEndpointProperties, 0)

	alreadyExists := false
	for _, existingEndpoint := range *routing.Endpoints.ServiceBusQueues {
		if existingEndpointName := existingEndpoint.Name; existingEndpointName != nil {
			if strings.EqualFold(*existingEndpointName, id.EndpointName) {
				if d.IsNewResource() {
					return tf.ImportAsExistsError("azurerm_iothub_endpoint_servicebus_queue", id.ID())
				}
				endpoints = append(endpoints, queueEndpoint)
				alreadyExists = true
			} else {
				endpoints = append(endpoints, existingEndpoint)
			}
		}
	}

	if d.IsNewResource() {
		endpoints = append(endpoints, queueEndpoint)
	} else if !alreadyExists {
		return fmt.Errorf("unable to find %s", id)
	}
	routing.Endpoints.ServiceBusQueues = &endpoints

	future, err := client.CreateOrUpdate(ctx, iotHubRG, iotHubName, iothub, "")
	if err != nil {
		return fmt.Errorf("creating/updating %s: %+v", id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for the completion of the creating/updating of %s: %+v", id, err)
	}

	d.SetId(id.ID())

	return resourceIotHubEndpointServiceBusQueueRead(d, meta)
}

func resourceIotHubEndpointServiceBusQueueRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).IoTHub.ResourceClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.EndpointServiceBusQueueID(d.Id())
	if err != nil {
		return err
	}

	iothub, err := client.Get(ctx, id.ResourceGroup, id.IotHubName)
	if err != nil {
		return fmt.Errorf("loading IotHub %q (Resource Group %q): %+v", id.IotHubName, id.ResourceGroup, err)
	}

	d.Set("name", id.EndpointName)
	d.Set("iothub_name", id.IotHubName)

	iotHubId := parse.NewIotHubID(id.SubscriptionId, id.ResourceGroup, id.IotHubName)
	d.Set("iothub_id", iotHubId.ID())

	if iothub.Properties == nil || iothub.Properties.Routing == nil || iothub.Properties.Routing.Endpoints == nil {
		return nil
	}

	if endpoints := iothub.Properties.Routing.Endpoints.ServiceBusQueues; endpoints != nil {
		for _, endpoint := range *endpoints {
			if existingEndpointName := endpoint.Name; existingEndpointName != nil {
				if strings.EqualFold(*existingEndpointName, id.EndpointName) {
					d.Set("connection_string", endpoint.ConnectionString)
					d.Set("resource_group_name", endpoint.ResourceGroup)
				}
			}
		}
	}

	return nil
}

func resourceIotHubEndpointServiceBusQueueDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).IoTHub.ResourceClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.EndpointServiceBusQueueID(d.Id())
	if err != nil {
		return err
	}

	locks.ByName(id.IotHubName, IothubResourceName)
	defer locks.UnlockByName(id.IotHubName, IothubResourceName)

	iothub, err := client.Get(ctx, id.ResourceGroup, id.IotHubName)
	if err != nil {
		if utils.ResponseWasNotFound(iothub.Response) {
			return fmt.Errorf("IotHub %q (Resource Group %q) was not found", id.IotHubName, id.ResourceGroup)
		}

		return fmt.Errorf("loading IotHub %q (Resource Group %q): %+v", id.IotHubName, id.ResourceGroup, err)
	}

	if iothub.Properties == nil || iothub.Properties.Routing == nil || iothub.Properties.Routing.Endpoints == nil {
		return nil
	}
	endpoints := iothub.Properties.Routing.Endpoints.ServiceBusQueues

	if endpoints == nil {
		return nil
	}

	updatedEndpoints := make([]devices.RoutingServiceBusQueueEndpointProperties, 0)
	for _, endpoint := range *endpoints {
		if existingEndpointName := endpoint.Name; existingEndpointName != nil {
			if !strings.EqualFold(*existingEndpointName, id.EndpointName) {
				updatedEndpoints = append(updatedEndpoints, endpoint)
			}
		}
	}

	iothub.Properties.Routing.Endpoints.ServiceBusQueues = &updatedEndpoints

	future, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.IotHubName, iothub, "")
	if err != nil {
		return fmt.Errorf("updating %s: %+v", id, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for %s to finish updating: %+v", id, err)
	}

	return nil
}
