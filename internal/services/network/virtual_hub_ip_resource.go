package network

import (
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-05-01/network"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/locks"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/network/parse"
	networkValidate "github.com/hashicorp/terraform-provider-azurerm/internal/services/network/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceVirtualHubIP() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceVirtualHubIPCreateUpdate,
		Read:   resourceVirtualHubIPRead,
		Update: resourceVirtualHubIPCreateUpdate,
		Delete: resourceVirtualHubIPDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(60 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(60 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(60 * time.Minute),
		},

		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.VirtualHubIpConfigurationID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"virtual_hub_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: networkValidate.VirtualHubID,
			},

			"subnet_id": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: networkValidate.SubnetID,
			},

			"private_ip_address": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPv4Address,
			},

			"private_ip_allocation_method": {
				Type:     pluginsdk.TypeString,
				Optional: true,
				Default:  network.IPAllocationMethodDynamic,
				ValidateFunc: validation.StringInSlice([]string{
					string(network.IPAllocationMethodDynamic),
					string(network.IPAllocationMethodStatic),
				}, false),
			},

			"public_ip_address_id": {
				Type:         pluginsdk.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: networkValidate.PublicIpAddressID,
			},
		},
	}
}

func resourceVirtualHubIPCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.VirtualHubIPClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.VirtualHubID(d.Get("virtual_hub_id").(string))
	if err != nil {
		return err
	}

	locks.ByName(id.Name, virtualHubResourceName)
	defer locks.UnlockByName(id.Name, virtualHubResourceName)

	name := d.Get("name").(string)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.Name, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for present of existing Virtual Hub IP %q (Resource Group %q / Virtual Hub %q): %+v", name, id.ResourceGroup, id.Name, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_virtual_hub_ip", *existing.ID)
		}

		if d.Get("public_ip_address_id").(string) == "" {
			return fmt.Errorf("`public_ip_address_id` is required for new resources, created after September 1st 2021")
		}
	}

	parameters := network.HubIPConfiguration{
		Name: utils.String(d.Get("name").(string)),
		HubIPConfigurationPropertiesFormat: &network.HubIPConfigurationPropertiesFormat{
			Subnet: &network.Subnet{
				ID: utils.String(d.Get("subnet_id").(string)),
			},
		},
	}

	if v, ok := d.GetOk("private_ip_address"); ok {
		parameters.HubIPConfigurationPropertiesFormat.PrivateIPAddress = utils.String(v.(string))
	}

	if v, ok := d.GetOk("private_ip_allocation_method"); ok {
		parameters.HubIPConfigurationPropertiesFormat.PrivateIPAllocationMethod = network.IPAllocationMethod(v.(string))
	}

	if v, ok := d.GetOk("public_ip_address_id"); ok {
		parameters.HubIPConfigurationPropertiesFormat.PublicIPAddress = &network.PublicIPAddress{
			ID: utils.String(v.(string)),
		}
	}

	future, err := client.CreateOrUpdate(ctx, id.ResourceGroup, id.Name, name, parameters)
	if err != nil {
		return fmt.Errorf("creating/updating Virtual Hub IP %q (Resource Group %q / Virtual Hub %q): %+v", name, id.ResourceGroup, id.Name, err)
	}

	if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting on creating/updating future for Virtual Hub IP %q (Resource Group %q / Virtual Hub %q): %+v", name, id.ResourceGroup, id.Name, err)
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.Name, name)
	if err != nil {
		return fmt.Errorf("retrieving Virtual Hub IP %q (Resource Group %q / Virtual Hub %q): %+v", name, id.ResourceGroup, id.Name, err)
	}

	if resp.ID == nil || *resp.ID == "" {
		return fmt.Errorf("empty or nil ID returned for Virtual Hub IP %q (Resource Group %q / Virtual Hub %q) ID", name, id.ResourceGroup, id.Name)
	}

	d.SetId(*resp.ID)

	return resourceVirtualHubIPRead(d, meta)
}

func resourceVirtualHubIPRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.VirtualHubIPClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.VirtualHubIpConfigurationID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.VirtualHubName, id.IpConfigurationName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] Virtual Hub IP %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("retrieving Virtual Hub IP %q (Resource Group %q / Virtual Hub %q): %+v", id.IpConfigurationName, id.ResourceGroup, id.VirtualHubName, err)
	}

	d.Set("name", id.IpConfigurationName)
	d.Set("virtual_hub_id", parse.NewVirtualHubID(id.SubscriptionId, id.ResourceGroup, id.VirtualHubName).ID())

	if props := resp.HubIPConfigurationPropertiesFormat; props != nil {
		d.Set("private_ip_address", props.PrivateIPAddress)
		d.Set("private_ip_allocation_method", props.PrivateIPAllocationMethod)

		if v := props.PublicIPAddress; v != nil {
			d.Set("public_ip_address_id", v.ID)
		}

		if v := props.Subnet; v != nil {
			d.Set("subnet_id", v.ID)
		}
	}

	return nil
}

func resourceVirtualHubIPDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).Network.VirtualHubIPClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.VirtualHubIpConfigurationID(d.Id())
	if err != nil {
		return err
	}

	locks.ByName(id.VirtualHubName, virtualHubResourceName)
	defer locks.UnlockByName(id.VirtualHubName, virtualHubResourceName)

	future, err := client.Delete(ctx, id.ResourceGroup, id.VirtualHubName, id.IpConfigurationName)
	if err != nil {
		return fmt.Errorf("deleting Virtual Hub IP %q (Resource Group %q / virtualHubName %q): %+v", id.IpConfigurationName, id.ResourceGroup, id.VirtualHubName, err)
	}

	if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting on deleting future for Virtual Hub IP %q (Resource Group %q / virtualHubName %q): %+v", id.IpConfigurationName, id.ResourceGroup, id.VirtualHubName, err)
	}

	return nil
}
