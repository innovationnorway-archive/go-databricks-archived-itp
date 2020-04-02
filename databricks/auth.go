package databricks

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest"
)

const (
	// AzureDatabricksApplicationID is the Application (client) ID of the "AzureDatabricks"
	// Azure AD Application
	AzureDatabricksApplicationID = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"

	headerAuthorization = "Authorization"

	headerXDatabricksOrgID = "X-Databricks-Org-Id"

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

// TokenAuthorizerWithOrgID implements the bearer authorization
type TokenAuthorizerWithOrgID struct {
	token string
	orgID string
}

// NewTokenAuthorizerWithOrgID crates a TokenAuthorizer using the given token provider
func NewTokenAuthorizerWithOrgID(token, orgID string) *TokenAuthorizerWithOrgID {
	return &TokenAuthorizerWithOrgID{
		token: token,
		orgID: orgID,
	}
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header using the
// primary token along with the Databricks authorization headers using the organization ID
func (t *TokenAuthorizerWithOrgID) WithAuthorization() autorest.PrepareDecorator {
	headers := make(map[string]interface{})
	headers[headerAuthorization] = fmt.Sprintf("Bearer %s", t.token)
	headers[headerXDatabricksOrgID] = t.orgID

	return autorest.NewAPIKeyAuthorizerWithHeaders(headers).WithAuthorization()
}

// TokenAuthorizerWithWorkspaceID implements the bearer authorization
type TokenAuthorizerWithWorkspaceID struct {
	token       string
	workspaceID string
}

// NewTokenAuthorizerWithWorkspaceID crates a TokenAuthorizer using the given token provider
func NewTokenAuthorizerWithWorkspaceID(token, workspaceID string) *TokenAuthorizerWithWorkspaceID {
	return &TokenAuthorizerWithWorkspaceID{
		token:       token,
		workspaceID: workspaceID,
	}
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header using the
// primary token along with the Azure Databricks authorization headers using the workspace ID
func (t *TokenAuthorizerWithWorkspaceID) WithAuthorization() autorest.PrepareDecorator {
	headers := make(map[string]interface{})
	headers[headerAuthorization] = fmt.Sprintf("Bearer %s", t.token)
	headers[headerXDatabricksAzureWorkspaceResourceID] = t.workspaceID

	return autorest.NewAPIKeyAuthorizerWithHeaders(headers).WithAuthorization()
}

// TokenAuthorizerWithServicePrincipal implements authorization for Azure Databricks.
type TokenAuthorizerWithServicePrincipal struct {
	token           string
	managementToken string
	workspaceID     string
}

// NewTokenAuthorizerWithServicePrincipal creates a AzureTokenAuthorizer using the given credentials
func NewTokenAuthorizerWithServicePrincipal(token, managementToken, workspaceID string) *TokenAuthorizerWithServicePrincipal {
	return &TokenAuthorizerWithServicePrincipal{
		token:           token,
		managementToken: managementToken,
		workspaceID:     workspaceID,
	}
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header using the
// primary token along with the Azure Databricks authorization headers using the management token
// and workspace ID
func (t *TokenAuthorizerWithServicePrincipal) WithAuthorization() autorest.PrepareDecorator {
	headers := make(map[string]interface{})
	headers[headerAuthorization] = fmt.Sprintf("Bearer %s", t.token)
	headers[headerXDatabricksAzureSPManagementToken] = t.managementToken
	headers[headerXDatabricksAzureWorkspaceResourceID] = t.workspaceID

	return autorest.NewAPIKeyAuthorizerWithHeaders(headers).WithAuthorization()
}
