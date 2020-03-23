package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Azure/go-autorest/autorest"
)

const (
	// AzureDatabricksApplicationID is the Application (client) ID of the "AzureDatabricks"
	// Azure AD Application
	AzureDatabricksApplicationID = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"

	headerAuthorization = "Authorization"

	headerXDatabricksAzureSPManagementToken   = "X-Databricks-Azure-SP-Management-Token"
	headerXDatabricksAzureWorkspaceResourceID = "X-Databricks-Azure-Workspace-Resource-Id"
)

// AzureDatabricksAuthorizer implements an authorization for Databricks Azure AD Authentication
// this can be used for interaction with Azure Databricks API
type AzureDatabricksAuthorizer struct {
	token           string
	managementToken string
	workspaceID     string
}

// NewAzureDatabricksAuthorizer creates a AzureTokenAuthorizer using the given credentials
func NewAzureDatabricksAuthorizer(token, managementToken, workspaceID string) (*AzureDatabricksAuthorizer, error) {
	if strings.TrimSpace(token) == "" {
		return nil, fmt.Errorf("token cannot be empty")
	}

	return &AzureDatabricksAuthorizer{
		token:           token,
		managementToken: managementToken,
		workspaceID:     workspaceID,
	}, nil
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header using the
// primary token along with the Azure Databricks authorization headers using the management token
// and workspace ID
// This can be used for Databricks Azure AD Authentication
func (ada *AzureDatabricksAuthorizer) WithAuthorization() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, err := p.Prepare(r)
			if err != nil {
				return r, err
			}

			return autorest.Prepare(r,
				autorest.WithHeader(headerAuthorization, ada.token),
				autorest.WithHeader(headerXDatabricksAzureSPManagementToken, ada.managementToken),
				autorest.WithHeader(headerXDatabricksAzureWorkspaceResourceID, ada.workspaceID),
			)
		})
	}
}
