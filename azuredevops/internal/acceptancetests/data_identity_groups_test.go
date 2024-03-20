//go:build (all || core || data_sources || data_groups) && (!exclude_data_sources || !exclude_data_groups)
// +build all core data_sources data_groups
// +build !exclude_data_sources !exclude_data_groups

package acceptancetests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
)

func generateIdentityGroupsDataSourceConfig(projectName string) string {
	return fmt.Sprintf(`
resource "azuredevops_project" "test" {
	name               = "%[1]s"
	work_item_template = "Agile"
	version_control    = "Git"
	visibility         = "private"
	description        = "Managed by Terraform"
}

data "azuredevops_identity_groups" "test" {
	project_id = azuredevops_project.test.id
}
`, projectName)
}

func testIdentityGroupsDataSource(t *testing.T, projectName string) {
	tfNode := "data.azuredevops_identity_groups.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		Steps: []resource.TestStep{
			{
				Config: generateIdentityGroupsDataSourceConfig(projectName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "groups.#"),
				),
			},
		},
	})
}

func TestAccIdentityGroupsDataSource(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	testIdentityGroupsDataSource(t, projectName)
}
