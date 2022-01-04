package synapse

import (
	"github.com/aquasecurity/defsec/rules/azure/synapse"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/scanner"
	"github.com/aquasecurity/tfsec/pkg/rule"
)

func init() {
	scanner.RegisterCheckRule(rule.Rule{
		LegacyID: "AZU027",
		BadExample: []string{`
 resource "azurerm_synapse_workspace" "bad_example" {
   name                                 = "example"
   resource_group_name                  = azurerm_resource_group.example.name
   location                             = azurerm_resource_group.example.location
   storage_data_lake_gen2_filesystem_id = azurerm_storage_data_lake_gen2_filesystem.example.id
   sql_administrator_login              = "sqladminuser"
   sql_administrator_login_password     = "H@Sh1CoR3!"
 
   aad_admin {
     login     = "AzureAD Admin"
     object_id = "00000000-0000-0000-0000-000000000000"
     tenant_id = "00000000-0000-0000-0000-000000000000"
   }
 
   tags = {
     Env = "production"
   }
 }
 `},
		GoodExample: []string{`
 resource "azurerm_synapse_workspace" "good_example" {
   name                                 = "example"
   resource_group_name                  = azurerm_resource_group.example.name
   location                             = azurerm_resource_group.example.location
   storage_data_lake_gen2_filesystem_id = azurerm_storage_data_lake_gen2_filesystem.example.id
   sql_administrator_login              = "sqladminuser"
   sql_administrator_login_password     = "H@Sh1CoR3!"
   managed_virtual_network_enabled	   = true
   aad_admin {
     login     = "AzureAD Admin"
     object_id = "00000000-0000-0000-0000-000000000000"
     tenant_id = "00000000-0000-0000-0000-000000000000"
   }
 
   tags = {
     Env = "production"
   }
 }
 `},
		Links: []string{
			"https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/synapse_workspace#managed_virtual_network_enabled",
		},
		RequiredTypes:  []string{"resource"},
		RequiredLabels: []string{"azurerm_synapse_workspace"},
		Base:           synapse.CheckVirtualNetworkEnabled,
	})
}
