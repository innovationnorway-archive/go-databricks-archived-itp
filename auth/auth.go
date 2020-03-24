package auth

import (
	"fmt"

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

// TokenAuthorizer implements the bearer authorization
type TokenAuthorizer struct {
	token string
}

// NewTokenAuthorizer crates a TokenAuthorizer using the given token provider
func NewTokenAuthorizer(token string) *TokenAuthorizer {
	return &TokenAuthorizer{token: token}
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header whose
// value is "Bearer " followed by the token.
func (t *TokenAuthorizer) WithAuthorization() autorest.PrepareDecorator {
	headers := make(map[string]interface{})
	headers[headerAuthorization] = fmt.Sprintf("Bearer %s", t.token)

	return autorest.NewAPIKeyAuthorizerWithHeaders(headers).WithAuthorization()
}

// AzureDatabricksAuthorizer implements authorization for Azure Databricks.
type AzureDatabricksAuthorizer struct {
	token           string
	managementToken string
	workspaceID     string
}

// NewAzureDatabricksAuthorizer creates a AzureTokenAuthorizer using the given credentials
func NewAzureDatabricksAuthorizer(token, managementToken, workspaceID string) *AzureDatabricksAuthorizer {
	return &AzureDatabricksAuthorizer{
		token:           token,
		managementToken: managementToken,
		workspaceID:     workspaceID,
	}
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header using the
// primary token along with the Azure Databricks authorization headers using the management token
// and workspace ID
func (ada *AzureDatabricksAuthorizer) WithAuthorization() autorest.PrepareDecorator {
	headers := make(map[string]interface{})
	headers[headerAuthorization] = fmt.Sprintf("Bearer %s", ada.token)
	headers[headerXDatabricksAzureSPManagementToken] = ada.managementToken
	headers[headerXDatabricksAzureWorkspaceResourceID] = ada.workspaceID

	return autorest.NewAPIKeyAuthorizerWithHeaders(headers).WithAuthorization()
}
