package porcelain

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AzureAuth provides Azure Databricks auth info writer
func AzureAuth(token, managementToken, workspaceID string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		r.SetHeaderParam("Authorization", "Bearer "+token)
		r.SetHeaderParam("X-Databricks-Azure-SP-Management-Token", managementToken)
		r.SetHeaderParam("X-Databricks-Azure-Workspace-Resource-Id", workspaceID)
		return nil
	})
}
