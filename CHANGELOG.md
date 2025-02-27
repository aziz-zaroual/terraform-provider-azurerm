## 2.92.0 (Unreleased)

FEATURES:

* `azurerm_api_management_custom_domain` - the `proxy` property has been deprecated in favour of the `gateway` for the 3.0 release [GH-14628]

ENHANCEMENTS:

* `azurerm_monitor_action_group` - support for the `event_hub_receiver` block [GH-14771]

BUG FIXES:

* `azurerm_consumption_budget_subscription` - Fix issue in migration logic [GH-14898]
* `azurerm_cosmosdb_mongo_collection` - now validates that "_id" is included as an index key [GH-14857]
* `azurem_hdinsight` - hdinsight resources using oozie metastore can now be created without error [GH-14880]


## 2.91.0 (January 07, 2022)

FEATURES:

* **New Data Source:** `azurerm_aadb2c_directory` ([#14671](https://github.com/hashicorp/terraform-provider-azurerm/issues/14671))
* **New Data Source:** `azurerm_sql_managed_instance` ([#14739](https://github.com/hashicorp/terraform-provider-azurerm/issues/14739))
* **New Resource:** `azurerm_aadb2c_directory` ([#14671](https://github.com/hashicorp/terraform-provider-azurerm/issues/14671))
* **New Resource:** `azurerm_app_service_slot_custom_hostname_binding` ([#13097](https://github.com/hashicorp/terraform-provider-azurerm/issues/13097))
* **New Resource:** `azurerm_data_factory_linked_service_odbc` ([#14787](https://github.com/hashicorp/terraform-provider-azurerm/issues/14787))
* **New Resource:** `azurerm_disk_pool` ([#14675](https://github.com/hashicorp/terraform-provider-azurerm/issues/14675))
* **New Resource:** `azurerm_load_test` ([#14724](https://github.com/hashicorp/terraform-provider-azurerm/issues/14724))
* **New Resource:** `azurerm_virtual_desktop_scaling_plan` ([#14188](https://github.com/hashicorp/terraform-provider-azurerm/issues/14188))

ENHANCEMENTS:

* dependencies: upgrading `appplatform` to API version `2021-09-01-preview` ([#14365](https://github.com/hashicorp/terraform-provider-azurerm/issues/14365))
* dependencies: upgrading `network` to API Version `2021-05-01` ([#14164](https://github.com/hashicorp/terraform-provider-azurerm/issues/14164))
* dependencies: upgrading to `v60.2.0` of `github.com/Azure/azure-sdk-for-go` ([#14688](https://github.com/hashicorp/terraform-provider-azurerm/issues/14688)] and [[#14667](https://github.com/hashicorp/terraform-provider-azurerm/issues/14667))
* dependencies: upgrading to `v2.10.1` of `github.com/hashicorp/terraform-plugin-sdk` ([#14666](https://github.com/hashicorp/terraform-provider-azurerm/issues/14666))
* `azurerm_application_gateway` - support for the `key_vault_secret_id` and `force_firewall_policy_association` properties ([#14413](https://github.com/hashicorp/terraform-provider-azurerm/issues/14413))
* `azurerm_application_gateway` - support the `fips_enagled` property ([#14797](https://github.com/hashicorp/terraform-provider-azurerm/issues/14797))
* `azurerm_cdn_endpoint_custom_domain` - support for HTTPS ([#13283](https://github.com/hashicorp/terraform-provider-azurerm/issues/13283))
* `azurerm_hdinsight_hbase_cluster` - support for the `network` property ([#14825](https://github.com/hashicorp/terraform-provider-azurerm/issues/14825))
* `azurerm_iothub` - support for the `identity` block ([#14354](https://github.com/hashicorp/terraform-provider-azurerm/issues/14354))
* `azurerm_iothub_endpoint_servicebus_queue_resource` - depracating the `iothub_name` propertyin favour of `iothub_id` property ([#14690](https://github.com/hashicorp/terraform-provider-azurerm/issues/14690))
* `azurerm_iothub_endpoint_storage_container_resource` - depracating the `iothub_name` property in favour of `iothub_id` property [[#14690](https://github.com/hashicorp/terraform-provider-azurerm/issues/14690)] 
* `azurerm_iot_fallback_route` - support for the `source` property ([#14836](https://github.com/hashicorp/terraform-provider-azurerm/issues/14836))
* `azurerm_kubernetes_cluster` - support for the `public_network_access_enabled`, `scale_down_mode`, and `workload_runtime` properties ([#14386](https://github.com/hashicorp/terraform-provider-azurerm/issues/14386))
* `azurerm_linux_function_app` - (Beta Resource) fix the filtering of `app_settings` for `WEBSITE_CONTENTSHARE` and `WEBSITE_CONTENTAZUREFILECONNECTIONSTRING` ([#14815](https://github.com/hashicorp/terraform-provider-azurerm/issues/14815))
* `azurerm_linux_virtual_machine` - support for the `user_data` property ([#13888](https://github.com/hashicorp/terraform-provider-azurerm/issues/13888))
* `azurerm_linux_virtual_machine_scale_set` - support for the `user_data` property ([#13888](https://github.com/hashicorp/terraform-provider-azurerm/issues/13888))
* `azurerm_managed_disk` - support for the `gallery_image_reference_id` property ([#14121](https://github.com/hashicorp/terraform-provider-azurerm/issues/14121))
* `azurerm_mysql_server` - support capacities up to `16TB` for the `storage_mb` property ([#14838](https://github.com/hashicorp/terraform-provider-azurerm/issues/14838))
* `azurerm_postgresql_flexible_server` - support for the `geo_redundant_backup_enabled` property ([#14661](https://github.com/hashicorp/terraform-provider-azurerm/issues/14661))
* `azurerm_recovery_services_vault` - support for the `storage_mode_type` property ([#14659](https://github.com/hashicorp/terraform-provider-azurerm/issues/14659))
* `azurerm_spring_cloud_certificate` - support for the `certificate_content` property ([#14689](https://github.com/hashicorp/terraform-provider-azurerm/issues/14689))
* `azurerm_servicebus_namespace_authorization_rule` - the `resource_group_name` and `namespace_name` properties have been deprecated in favour of the `namespace_id` property ([#14784](https://github.com/hashicorp/terraform-provider-azurerm/issues/14784))
* `azurerm_servicebus_namespace_network_rule_set` - the `resource_group_name` and `namespace_name` properties have been deprecated in favour of the `namespace_id` property ([#14784](https://github.com/hashicorp/terraform-provider-azurerm/issues/14784))
* `azurerm_servicebus_namespace_authorization_rule` - the `resource_group_name` and `namespace_name` properties have been deprecated in favour of the `namespace_id` property ([#14784](https://github.com/hashicorp/terraform-provider-azurerm/issues/14784))
* `azurerm_servicebus_queue` - the `resource_group_name` and `namespace_name` properties have been deprecated in favour of the `namespace_id` property ([#14784](https://github.com/hashicorp/terraform-provider-azurerm/issues/14784))
* `azurerm_servicebus_queue_authorization_rule` - the `resource_group_name`, `namespace_name`, and `queue_name` properties have been deprecated in favour of the `queue_id` property ([#14784](https://github.com/hashicorp/terraform-provider-azurerm/issues/14784))
* `azurerm_servicebus_subscription` - the `resource_group_name`, `namespace_name`, and `topic_name` properties have been deprecated in favour of the `topic_id` property ([#14784](https://github.com/hashicorp/terraform-provider-azurerm/issues/14784))
* `azurerm_servicebus_subscription_rule` - the `resource_group_name`, `namespace_name`, `topic_name`, and `subscription_name` properties have been deprecated in favour of the `subscription_id` property ([#14784](https://github.com/hashicorp/terraform-provider-azurerm/issues/14784))
* `azurerm_servicebus_topic` - the `resource_group_name` and `namespace_name` properties have been deprecated in favour of the `namespace_id` property ([#14784](https://github.com/hashicorp/terraform-provider-azurerm/issues/14784))
* `azurerm_servicebus_topic_authorization_rule` - the `resource_group_name`, `namespace_name`, and `topic_name` properties have been deprecated in favour of the `topic_id` property ([#14784](https://github.com/hashicorp/terraform-provider-azurerm/issues/14784))
* `azurerm_shared_image_version` - images can now be sorted by semver ([#14708](https://github.com/hashicorp/terraform-provider-azurerm/issues/14708))
* `azurerm_virtual_network_gateway_connection` - support for the `connection_mode` property ([#14738](https://github.com/hashicorp/terraform-provider-azurerm/issues/14738))
* `azurerm_web_application_firewall_policy` - the `file_upload_limit_in_mb` property within the `policy_settings` block can now be set to `4000` ([#14715](https://github.com/hashicorp/terraform-provider-azurerm/issues/14715))
* `azurerm_windows_virtual_machine` - support for the `user_data` property ([#13888](https://github.com/hashicorp/terraform-provider-azurerm/issues/13888))
* `azurerm_windows_virtual_machine_scale_set` - support for the `user_data` property ([#13888](https://github.com/hashicorp/terraform-provider-azurerm/issues/13888))

BUG FIXES:

* `azurerm_app_service_environment_v3` - fix the default value of the `allow_new_private_endpoint_connections` property ([#14805](https://github.com/hashicorp/terraform-provider-azurerm/issues/14805))
* `azurerm_consumption_budget_subscription` - added an additional state migration to fix the bug introduced by the first one and to parse the `subscription_id` from the resource's ID ([#14803](https://github.com/hashicorp/terraform-provider-azurerm/issues/14803))
* `azurerm_network_interface_security_group_association` - checking the ID matches the expected format during import ([#14753](https://github.com/hashicorp/terraform-provider-azurerm/issues/14753))
* `azurerm_storage_management_policy` - handle the unexpected deletion of the storage account ([#14799](https://github.com/hashicorp/terraform-provider-azurerm/issues/14799))

## 2.90.0 (December 17, 2021)

FEATURES:

* **New Data Source:** `azurerm_app_configuration_key` ([#14484](https://github.com/hashicorp/terraform-provider-azurerm/issues/14484))
* **New Resource:** `azurerm_container_registry_task` ([#14533](https://github.com/hashicorp/terraform-provider-azurerm/issues/14533))
* **New Resource:** `azurerm_maps_creator` ([#14566](https://github.com/hashicorp/terraform-provider-azurerm/issues/14566))
* **New Resource:** `azurerm_netapp_snapshot_policy` ([#14230](https://github.com/hashicorp/terraform-provider-azurerm/issues/14230))
* **New Resource:** `azurerm_synapse_sql_pool_workload_classifier` ([#14412](https://github.com/hashicorp/terraform-provider-azurerm/issues/14412))
* **New Resource:** `azurerm_synapse_workspace_sql_aad_admin` ([#14341](https://github.com/hashicorp/terraform-provider-azurerm/issues/14341))
* **New Resource:** `azurerm_vpn_gateway_nat_rule` ([#14527](https://github.com/hashicorp/terraform-provider-azurerm/issues/14527))

ENHANCEMENTS:

* dependencies: updating `apimanagement` to API Version `2021-08-01` ([#14312](https://github.com/hashicorp/terraform-provider-azurerm/issues/14312))
* dependencies: updating `managementgroups` to API Version `2020-05-01` ([#14635](https://github.com/hashicorp/terraform-provider-azurerm/issues/14635))
* dependencies: updating `redisenterprise` to use an Embedded SDK ([#14502](https://github.com/hashicorp/terraform-provider-azurerm/issues/14502))
* dependencies: updating to `v0.19.1` of `github.com/hashicorp/go-azure-helpers` ([#14627](https://github.com/hashicorp/terraform-provider-azurerm/issues/14627))
* dependencies: updating to `v2.10.0` of `github.com/hashicorp/terraform-plugin-sdk` ([#14596](https://github.com/hashicorp/terraform-provider-azurerm/issues/14596))
* Data Source: `azurerm_function_app_host_keys` - support for `signalr_extension_key` and `durabletask_extension_key` ([#13648](https://github.com/hashicorp/terraform-provider-azurerm/issues/13648))
* `azurerm_application_gateway ` - support for private link configurations ([#14583](https://github.com/hashicorp/terraform-provider-azurerm/issues/14583))
* `azurerm_blueprint_assignment` - support for the `lock_exclude_actions` property ([#14648](https://github.com/hashicorp/terraform-provider-azurerm/issues/14648))
* `azurerm_container_group` - support for `ip_address_type = None` ([#14460](https://github.com/hashicorp/terraform-provider-azurerm/issues/14460))
* `azurerm_cosmosdb_account` - support for the `create_mode` property and `restore` block ([#14362](https://github.com/hashicorp/terraform-provider-azurerm/issues/14362))
* `azurerm_data_factory_dataset_*` - deprecate `data_factory_name` in favour of `data_factory_id` for consistency across all data factory dataset resources ([#14610](https://github.com/hashicorp/terraform-provider-azurerm/issues/14610))
* `azurerm_data_factory_integration_runtime_*`- deprecate `data_factory_name` in favour of `data_factory_id` for consistency across all data factory integration runtime resources ([#14610](https://github.com/hashicorp/terraform-provider-azurerm/issues/14610))
* `azurerm_data_factory_trigger_*`- deprecate `data_factory_name` in favour of `data_factory_id` for consistency across all data factory trigger resources ([#14610](https://github.com/hashicorp/terraform-provider-azurerm/issues/14610))
* `azurerm_data_factory_pipeline`- deprecate `data_factory_name` in favour of `data_factory_id` for consistency across all data factory resources ([#14610](https://github.com/hashicorp/terraform-provider-azurerm/issues/14610))
* `azurerm_iothub` - support for the `cloud_to_device` block ([#14546](https://github.com/hashicorp/terraform-provider-azurerm/issues/14546))
* `azurerm_iothub_endpoint_eventhub` - the `iothub_name` property has been deprecated in favour of the `iothub_id` property ([#14632](https://github.com/hashicorp/terraform-provider-azurerm/issues/14632))
* `azurerm_logic_app_workflow` - support for the `open_authentication_policy` block ([#14007](https://github.com/hashicorp/terraform-provider-azurerm/issues/14007))
* `azurerm_signalr` - support for the `live_trace_enabled` property ([#14646](https://github.com/hashicorp/terraform-provider-azurerm/issues/14646))
* `azurerm_xyz_policy_assignment` add support for `non_compliance_message` ([#14518](https://github.com/hashicorp/terraform-provider-azurerm/issues/14518))

BUG FIXES:

* `azurerm_cosmosdb_account` - will now set a default value for `default_identity_type` when the API return a nil value ([#14643](https://github.com/hashicorp/terraform-provider-azurerm/issues/14643))
* `azurerm_function_app` - address `app_settings` during creation rather than just updates ([#14638](https://github.com/hashicorp/terraform-provider-azurerm/issues/14638))
* `azurerm_marketplace_agreement` - fix crash when the import check triggers ([#14614](https://github.com/hashicorp/terraform-provider-azurerm/issues/14614))
* `azurerm_postgresql_configuration` - now locks during write operations to prevent conflicts ([#14619](https://github.com/hashicorp/terraform-provider-azurerm/issues/14619))
* `azurerm_postgresql_flexible_server_configuration` - now locks during write operations to prevent conflicts ([#14607](https://github.com/hashicorp/terraform-provider-azurerm/issues/14607))

---

For information on changes between the v2.89.0 and v2.0.0 releases, please see [the previous v2.x changelog entries](https://github.com/hashicorp/terraform-provider-azurerm/blob/main/CHANGELOG-v2.md).

For information on changes between the v2.00.0 and v1.0.0 releases, please see [the previous v1.x changelog entries](https://github.com/hashicorp/terraform-provider-azurerm/blob/main/CHANGELOG-v1.md).

For information on changes prior to the v1.0.0 release, please see [the v0.x changelog](https://github.com/hashicorp/terraform-provider-azurerm/blob/main/CHANGELOG-v0.md).
