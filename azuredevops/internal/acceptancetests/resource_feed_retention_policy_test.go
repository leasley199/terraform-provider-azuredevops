package acceptancetests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/feed"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/internal/client"
)

func TestAccFeedRetentionPolicy_basic(t *testing.T) {
	name := testutils.GenerateResourceName()

	tfNode := "azuredevops_feed_retention_policy.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      checkFeedRetentionPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: hclFeedRetentionPolicyBasic(name),
				Check: resource.ComposeTestCheckFunc(
					CheckFeedRetentionPolicyExist(20),
					resource.TestCheckResourceAttr(tfNode, "count_limit", "20"),
					resource.TestCheckResourceAttr(tfNode, "days_to_keep_recently_downloaded_packages", "30"),
				),
			},
		},
	})
}

func TestAccFeedRetentionPolicy_update(t *testing.T) {
	name := testutils.GenerateResourceName()

	tfNode := "azuredevops_feed_retention_policy.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      checkFeedRetentionPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: hclFeedRetentionPolicyBasic(name),
				Check: resource.ComposeTestCheckFunc(
					CheckFeedRetentionPolicyExist(20),
					resource.TestCheckResourceAttrSet(tfNode, "count_limit"),
				),
			},
			{
				Config: hclFeedRetentionPolicyUpdate(name),
				Check: resource.ComposeTestCheckFunc(
					CheckFeedRetentionPolicyExist(21),
					resource.TestCheckResourceAttr(tfNode, "count_limit", "21"),
					resource.TestCheckResourceAttr(tfNode, "days_to_keep_recently_downloaded_packages", "31"),
				),
			},
		},
	})
}

func TestAccFeedRetentionPolicy_requiresImportError(t *testing.T) {
	name := testutils.GenerateResourceName()

	tfNode := "azuredevops_feed_retention_policy.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		CheckDestroy:      checkFeedRetentionPolicyDestroyed,
		Steps: []resource.TestStep{
			{
				Config: hclFeedRetentionPolicyBasic(name),
				Check: resource.ComposeTestCheckFunc(
					CheckFeedRetentionPolicyExist(20),
					resource.TestCheckResourceAttrSet(tfNode, "count_limit"),
				),
			},
			{
				Config: hclFeedRetentionPolicyImport(name),
				Check: resource.ComposeTestCheckFunc(
					CheckFeedRetentionPolicyExist(20),
					resource.TestCheckResourceAttr(tfNode, "count_limit", "20"),
					resource.TestCheckResourceAttr(tfNode, "days_to_keep_recently_downloaded_packages", "30"),
				),
			},
		},
	})
}

func checkFeedRetentionPolicyDestroyed(s *terraform.State) error {
	clients := testutils.GetProvider().Meta().(*client.AggregatedClient)
	for _, res := range s.RootModule().Resources {
		if res.Type != "azuredevops_feed_retention_policy" {
			continue
		}
		id := res.Primary.ID
		projectID := res.Primary.Attributes["project_id"]

		_, err := clients.FeedClient.GetFeedRetentionPolicies(clients.Ctx, feed.GetFeedRetentionPoliciesArgs{
			FeedId:  &id,
			Project: &projectID,
		})
		if err == nil {
			return fmt.Errorf(" Feed Retention Policy (Feed ID: %s) should not exist", id)
		}
	}
	return nil
}

func CheckFeedRetentionPolicyExist(expectedCountLimit int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, ok := s.RootModule().Resources["azuredevops_feed_retention_policy.test"]
		if !ok {
			return fmt.Errorf(" Did not find a `azuredevops_feed_retention_policy` in the TF state")
		}

		clients := testutils.GetProvider().Meta().(*client.AggregatedClient)
		id := res.Primary.ID
		projectID := res.Primary.Attributes["project_id"]

		policy, err := clients.FeedClient.GetFeedRetentionPolicies(clients.Ctx, feed.GetFeedRetentionPoliciesArgs{
			FeedId:  &id,
			Project: &projectID,
		})

		if err != nil {
			return fmt.Errorf(" Feed Retention Policy with Feed ( Feed ID=%s ) cannot be found!. Error=%v", id, err)
		}

		if *policy.CountLimit != expectedCountLimit {
			return fmt.Errorf(" Feed Retention Policy with ( Feed ID=%s ) has CountLimit=%d, but expected CountLimit=%d", id, *policy.CountLimit, expectedCountLimit)
		}
		return nil
	}
}

func hclFeedRetentionPolicyTemplate(name string) string {
	return fmt.Sprintf(`
resource "azuredevops_project" "test" {
  name = "%[1]s"
}

resource "azuredevops_feed" "test" {
  project_id = azuredevops_project.test.id
  name       = "%[1]s"
}`, name)
}

func hclFeedRetentionPolicyBasic(name string) string {
	return fmt.Sprintf(`
%s

resource "azuredevops_feed_retention_policy" "test" {
  project_id                                = azuredevops_project.test.id
  feed_id                                   = azuredevops_feed.test.id
  count_limit                               = 20
  days_to_keep_recently_downloaded_packages = 30
}
`, hclFeedRetentionPolicyTemplate(name))
}

func hclFeedRetentionPolicyUpdate(name string) string {
	return fmt.Sprintf(`
%s

resource "azuredevops_feed_retention_policy" "test" {
  project_id                                = azuredevops_project.test.id
  feed_id                                   = azuredevops_feed.test.id
  count_limit                               = 21
  days_to_keep_recently_downloaded_packages = 31
}
`, hclFeedRetentionPolicyTemplate(name))
}

func hclFeedRetentionPolicyImport(name string) string {
	return fmt.Sprintf(`
%s

resource "azuredevops_feed_retention_policy" "import" {
  feed_id                                   = azuredevops_feed.test.id
  project_id                                = azuredevops_project.test.id
  count_limit                               = azuredevops_feed_retention_policy.test.count_limit
  days_to_keep_recently_downloaded_packages = azuredevops_feed_retention_policy.test.days_to_keep_recently_downloaded_packages
}
`, hclFeedRetentionPolicyBasic(name))
}
