// +build all resource_resource_authorization
// +build !exclude_resource_authorization

package azuredevops

// The tests in this file use the mock clients in mock_client.go to mock out
// the Azure DevOps client operations.

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/microsoft/terraform-provider-azuredevops/azdosdkmocks"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/utils/config"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/utils/converter"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/utils/testhelper"
	"github.com/stretchr/testify/require"
)

var projectId = "projectid"
var endpointId = uuid.New()

var resourceReferenceAuthorized = build.DefinitionResourceReference{
	Authorized: converter.Bool(true),
	Id:         converter.String(endpointId.String()),
	Name:       nil,
	Type:       converter.String("endpoint"),
}

var resourceReferenceNotAuthorized = build.DefinitionResourceReference{
	Authorized: converter.Bool(false),
	Id:         converter.String(endpointId.String()),
	Name:       nil,
	Type:       converter.String("endpoint"),
}

func init() {
	InitProvider()
}

/**
 * Begin unit tests
 */

func TestAzureDevOpsResourceAuthorization_FlattenExpand_RoundTripTestAzureDevOpsResourceAuthorization_FlattenExpand_RoundTrip(t *testing.T) {
	resourceData := schema.TestResourceDataRaw(t, resourceResourceAuthorization().Schema, nil)
	flattenAuthorizedResource(resourceData, &resourceReferenceAuthorized, projectId)

	resourceReferenceAfterRoundtrip, projectIdAfterRoundtrip, err := expandAuthorizedResource(resourceData)
	require.Nil(t, err)
	require.Equal(t, resourceReferenceAuthorized, *resourceReferenceAfterRoundtrip)
	require.Equal(t, projectId, projectIdAfterRoundtrip)
}

func TestAzureDevOpsResourceAuthorization_Create_DoesNotSwallowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r, resourceData, clients := prepareForCreateOrUpdate(t, ctrl, "CreateResourceAuthorization() Failed")

	err := r.Create(resourceData, clients)
	require.Contains(t, err.Error(), "CreateResourceAuthorization() Failed")
}

func TestAzureDevOpsResourceAuthorization_Update_DoesNotSwallowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r, resourceData, clients := prepareForCreateOrUpdate(t, ctrl, "UpdateResourceAuthorization() Failed")

	err := r.Update(resourceData, clients)
	require.Contains(t, err.Error(), "UpdateResourceAuthorization() Failed")
}

func prepareForCreateOrUpdate(t *testing.T, ctrl *gomock.Controller, expectedMessage string) (*schema.Resource, *schema.ResourceData, *config.AggregatedClient) {
	r := resourceResourceAuthorization()
	resourceData := schema.TestResourceDataRaw(t, r.Schema, nil)
	flattenAuthorizedResource(resourceData, &resourceReferenceAuthorized, projectId)

	buildClient := azdosdkmocks.NewMockBuildClient(ctrl)
	clients := &config.AggregatedClient{BuildClient: buildClient, Ctx: context.Background()}

	expectedArgs := build.AuthorizeProjectResourcesArgs{
		Resources: &[]build.DefinitionResourceReference{resourceReferenceAuthorized},
		Project:   &projectId,
	}
	buildClient.
		EXPECT().
		AuthorizeProjectResources(clients.Ctx, expectedArgs).
		Return(nil, errors.New(expectedMessage)).
		Times(1)
	return r, resourceData, clients
}

func TestAzureDevOpsResourceAuthorization_Read_DoesNotSwallowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := resourceResourceAuthorization()
	resourceData := schema.TestResourceDataRaw(t, r.Schema, nil)
	flattenAuthorizedResource(resourceData, &resourceReferenceAuthorized, projectId)

	buildClient := azdosdkmocks.NewMockBuildClient(ctrl)
	clients := &config.AggregatedClient{BuildClient: buildClient, Ctx: context.Background()}

	expectedArgs := build.GetProjectResourcesArgs{
		Project: &projectId,
		Type:    resourceReferenceAuthorized.Type,
		Id:      resourceReferenceAuthorized.Id,
	}
	buildClient.
		EXPECT().
		GetProjectResources(clients.Ctx, expectedArgs).
		Return(nil, errors.New("ReadResourceAuthorization() Failed")).
		Times(1)

	err := r.Read(resourceData, clients)
	require.Contains(t, err.Error(), "ReadResourceAuthorization() Failed")
}

func TestAzureDevOpsResourceAuthorization_Delete_DoesNotSwallowError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := resourceResourceAuthorization()
	resourceData := schema.TestResourceDataRaw(t, r.Schema, nil)
	flattenAuthorizedResource(resourceData, &resourceReferenceNotAuthorized, projectId)

	buildClient := azdosdkmocks.NewMockBuildClient(ctrl)
	clients := &config.AggregatedClient{BuildClient: buildClient, Ctx: context.Background()}

	expectedArgs := build.AuthorizeProjectResourcesArgs{
		Resources: &[]build.DefinitionResourceReference{resourceReferenceNotAuthorized},
		Project:   &projectId,
	}
	buildClient.
		EXPECT().
		AuthorizeProjectResources(clients.Ctx, expectedArgs).
		Return(nil, errors.New("DeleteResourceAuthorization() Failed")).
		Times(1)

	err := r.Delete(resourceData, clients)
	require.Contains(t, err.Error(), "DeleteResourceAuthorization() Failed")
}

// Verifies that resource authorization can be created, updated and deleted
func TestAccResourceAuthorization_CRUD(t *testing.T) {
	projectName := testhelper.TestAccResourcePrefix + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
	serviceEndpointName := testhelper.TestAccResourcePrefix + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resourcesHCL := testhelper.TestAccServiceEndpointGitHubResource(projectName, serviceEndpointName)
	authedHCL := testhelper.TestAccResourceAuthorization("azuredevops_serviceendpoint_github.serviceendpoint.id", true)
	unAuthedHCL := testhelper.TestAccResourceAuthorization("azuredevops_serviceendpoint_github.serviceendpoint.id", false)

	tfAuthNode := "azuredevops_resource_authorization.auth"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testhelper.TestAccPreCheck(t, nil) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf("%s\n%s", resourcesHCL, authedHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tfAuthNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfAuthNode, "resource_id"),
					resource.TestCheckResourceAttr(tfAuthNode, "authorized", "true"),
				),
			}, {
				Config: fmt.Sprintf("%s\n%s", resourcesHCL, unAuthedHCL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tfAuthNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfAuthNode, "resource_id"),
					resource.TestCheckResourceAttr(tfAuthNode, "authorized", "false"),
				),
			},
		},
	})
}

func init() {
	InitProvider()
}
