package secrets

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "github.com/innovationnorway/go-databricks/secrets"

        // BackendType enumerates the values for backend type.
    type BackendType string

    const (
                // DATABRICKS ...
        DATABRICKS BackendType = "DATABRICKS"
            )
    // PossibleBackendTypeValues returns an array of possible values for the BackendType const type.
    func PossibleBackendTypeValues() []BackendType {
        return []BackendType{DATABRICKS}
    }

        // Permission enumerates the values for permission.
    type Permission string

    const (
                // MANAGE ...
        MANAGE Permission = "MANAGE"
                // READ ...
        READ Permission = "READ"
                // WRITE ...
        WRITE Permission = "WRITE"
            )
    // PossiblePermissionValues returns an array of possible values for the Permission const type.
    func PossiblePermissionValues() []Permission {
        return []Permission{MANAGE,READ,WRITE}
    }

        // Permission1 enumerates the values for permission 1.
    type Permission1 string

    const (
                // Permission1MANAGE ...
        Permission1MANAGE Permission1 = "MANAGE"
                // Permission1READ ...
        Permission1READ Permission1 = "READ"
                // Permission1WRITE ...
        Permission1WRITE Permission1 = "WRITE"
            )
    // PossiblePermission1Values returns an array of possible values for the Permission1 const type.
    func PossiblePermission1Values() []Permission1 {
        return []Permission1{Permission1MANAGE,Permission1READ,Permission1WRITE}
    }

        // Permission2 enumerates the values for permission 2.
    type Permission2 string

    const (
                // Permission2MANAGE ...
        Permission2MANAGE Permission2 = "MANAGE"
                // Permission2READ ...
        Permission2READ Permission2 = "READ"
                // Permission2WRITE ...
        Permission2WRITE Permission2 = "WRITE"
            )
    // PossiblePermission2Values returns an array of possible values for the Permission2 const type.
    func PossiblePermission2Values() []Permission2 {
        return []Permission2{Permission2MANAGE,Permission2READ,Permission2WRITE}
    }

            // ACLItemAttributes ...
            type ACLItemAttributes struct {
            Principal *string `json:"principal,omitempty"`
            // Permission - Possible values include: 'Permission2READ', 'Permission2WRITE', 'Permission2MANAGE'
            Permission Permission2 `json:"permission,omitempty"`
            }

            // AclsAttributes ...
            type AclsAttributes struct {
            Scope *string `json:"scope,omitempty"`
            Principal *string `json:"principal,omitempty"`
            }

            // AclsResult ...
            type AclsResult struct {
            autorest.Response `json:"-"`
            Principal *string `json:"principal,omitempty"`
            // Permission - Possible values include: 'Permission1READ', 'Permission1WRITE', 'Permission1MANAGE'
            Permission Permission1 `json:"permission,omitempty"`
            }

            // Attributes ...
            type Attributes struct {
            StringValue *string `json:"string_value,omitempty"`
            BytesValue *[]byte `json:"bytes_value,omitempty"`
            Scope *string `json:"scope,omitempty"`
            Key *string `json:"key,omitempty"`
            }

            // CreateScopeAttributes ...
            type CreateScopeAttributes struct {
            Scope *string `json:"scope,omitempty"`
            InitialManagePrincipal *string `json:"initial_manage_principal,omitempty"`
            }

            // DeleteScopeAttributes ...
            type DeleteScopeAttributes struct {
            Scope *string `json:"scope,omitempty"`
            }

            // ListResult ...
            type ListResult struct {
            autorest.Response `json:"-"`
            SecretsProperty *[]MetadataAttributes `json:"secrets,omitempty"`
            }

            // ListScopesResult ...
            type ListScopesResult struct {
            autorest.Response `json:"-"`
            Scopes *[]ScopeAttributes `json:"scopes,omitempty"`
            }

            // ListSecretAclsAttributes ...
            type ListSecretAclsAttributes struct {
            Scope *string `json:"scope,omitempty"`
            }

            // ListSecretAclsResult ...
            type ListSecretAclsResult struct {
            autorest.Response `json:"-"`
            Scope *[]ACLItemAttributes `json:"scope,omitempty"`
            }

            // ListSecretsAttributes ...
            type ListSecretsAttributes struct {
            Scope *string `json:"scope,omitempty"`
            }

            // MetadataAttributes ...
            type MetadataAttributes struct {
            Key *string `json:"key,omitempty"`
            LastUpdatedTimestamp *int64 `json:"last_updated_timestamp,omitempty"`
            }

            // PutSecretAclsAttributes ...
            type PutSecretAclsAttributes struct {
            Scope *string `json:"scope,omitempty"`
            Principal *string `json:"principal,omitempty"`
            // Permission - Possible values include: 'READ', 'WRITE', 'MANAGE'
            Permission Permission `json:"permission,omitempty"`
            }

            // ScopeAttributes ...
            type ScopeAttributes struct {
            Name *string `json:"name,omitempty"`
            // BackendType - Possible values include: 'DATABRICKS'
            BackendType BackendType `json:"backend_type,omitempty"`
            }

