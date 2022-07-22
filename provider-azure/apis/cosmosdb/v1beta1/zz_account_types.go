/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccountObservation struct {
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	GeoLocation []GeoLocationObservation `json:"geoLocation,omitempty" tf:"geo_location,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	ReadEndpoints []*string `json:"readEndpoints,omitempty" tf:"read_endpoints,omitempty"`

	WriteEndpoints []*string `json:"writeEndpoints,omitempty" tf:"write_endpoints,omitempty"`
}

type AccountParameters struct {

	// +kubebuilder:validation:Optional
	AccessKeyMetadataWritesEnabled *bool `json:"accessKeyMetadataWritesEnabled,omitempty" tf:"access_key_metadata_writes_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	AnalyticalStorage []AnalyticalStorageParameters `json:"analyticalStorage,omitempty" tf:"analytical_storage,omitempty"`

	// +kubebuilder:validation:Optional
	AnalyticalStorageEnabled *bool `json:"analyticalStorageEnabled,omitempty" tf:"analytical_storage_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Backup []BackupParameters `json:"backup,omitempty" tf:"backup,omitempty"`

	// +kubebuilder:validation:Optional
	Capabilities []CapabilitiesParameters `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	// +kubebuilder:validation:Optional
	Capacity []CapacityParameters `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Required
	ConsistencyPolicy []ConsistencyPolicyParameters `json:"consistencyPolicy" tf:"consistency_policy,omitempty"`

	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// +kubebuilder:validation:Optional
	CreateMode *string `json:"createMode,omitempty" tf:"create_mode,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultIdentityType *string `json:"defaultIdentityType,omitempty" tf:"default_identity_type,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAutomaticFailover *bool `json:"enableAutomaticFailover,omitempty" tf:"enable_automatic_failover,omitempty"`

	// +kubebuilder:validation:Optional
	EnableFreeTier *bool `json:"enableFreeTier,omitempty" tf:"enable_free_tier,omitempty"`

	// +kubebuilder:validation:Optional
	EnableMultipleWriteLocations *bool `json:"enableMultipleWriteLocations,omitempty" tf:"enable_multiple_write_locations,omitempty"`

	// +kubebuilder:validation:Required
	GeoLocation []GeoLocationParameters `json:"geoLocation" tf:"geo_location,omitempty"`

	// +kubebuilder:validation:Optional
	IPRangeFilter *string `json:"ipRangeFilter,omitempty" tf:"ip_range_filter,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	IsVirtualNetworkFilterEnabled *bool `json:"isVirtualNetworkFilterEnabled,omitempty" tf:"is_virtual_network_filter_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// +kubebuilder:validation:Optional
	LocalAuthenticationDisabled *bool `json:"localAuthenticationDisabled,omitempty" tf:"local_authentication_disabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MongoServerVersion *string `json:"mongoServerVersion,omitempty" tf:"mongo_server_version,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkACLBypassForAzureServices *bool `json:"networkAclBypassForAzureServices,omitempty" tf:"network_acl_bypass_for_azure_services,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkACLBypassIds []*string `json:"networkAclBypassIds,omitempty" tf:"network_acl_bypass_ids,omitempty"`

	// +kubebuilder:validation:Required
	OfferType *string `json:"offerType" tf:"offer_type,omitempty"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Restore []RestoreParameters `json:"restore,omitempty" tf:"restore,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkRule []VirtualNetworkRuleParameters `json:"virtualNetworkRule,omitempty" tf:"virtual_network_rule,omitempty"`
}

type AnalyticalStorageObservation struct {
}

type AnalyticalStorageParameters struct {

	// +kubebuilder:validation:Required
	SchemaType *string `json:"schemaType" tf:"schema_type,omitempty"`
}

type BackupObservation struct {
}

type BackupParameters struct {

	// +kubebuilder:validation:Optional
	IntervalInMinutes *float64 `json:"intervalInMinutes,omitempty" tf:"interval_in_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionInHours *float64 `json:"retentionInHours,omitempty" tf:"retention_in_hours,omitempty"`

	// +kubebuilder:validation:Optional
	StorageRedundancy *string `json:"storageRedundancy,omitempty" tf:"storage_redundancy,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type CapabilitiesObservation struct {
}

type CapabilitiesParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type CapacityObservation struct {
}

type CapacityParameters struct {

	// +kubebuilder:validation:Required
	TotalThroughputLimit *float64 `json:"totalThroughputLimit" tf:"total_throughput_limit,omitempty"`
}

type ConsistencyPolicyObservation struct {
}

type ConsistencyPolicyParameters struct {

	// +kubebuilder:validation:Required
	ConsistencyLevel *string `json:"consistencyLevel" tf:"consistency_level,omitempty"`

	// +kubebuilder:validation:Optional
	MaxIntervalInSeconds *float64 `json:"maxIntervalInSeconds,omitempty" tf:"max_interval_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	MaxStalenessPrefix *float64 `json:"maxStalenessPrefix,omitempty" tf:"max_staleness_prefix,omitempty"`
}

type CorsRuleObservation struct {
}

type CorsRuleParameters struct {

	// +kubebuilder:validation:Required
	AllowedHeaders []*string `json:"allowedHeaders" tf:"allowed_headers,omitempty"`

	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// +kubebuilder:validation:Required
	ExposedHeaders []*string `json:"exposedHeaders" tf:"exposed_headers,omitempty"`

	// +kubebuilder:validation:Required
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds" tf:"max_age_in_seconds,omitempty"`
}

type DatabaseObservation struct {
}

type DatabaseParameters struct {

	// +kubebuilder:validation:Optional
	CollectionNames []*string `json:"collectionNames,omitempty" tf:"collection_names,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type GeoLocationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GeoLocationParameters struct {

	// +kubebuilder:validation:Required
	FailoverPriority *float64 `json:"failoverPriority" tf:"failover_priority,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RestoreObservation struct {
}

type RestoreParameters struct {

	// +kubebuilder:validation:Optional
	Database []DatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// +kubebuilder:validation:Required
	RestoreTimestampInUtc *string `json:"restoreTimestampInUtc" tf:"restore_timestamp_in_utc,omitempty"`

	// +crossplane:generate:reference:type=Account
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SourceCosmosDBAccountID *string `json:"sourceCosmosdbAccountId,omitempty" tf:"source_cosmosdb_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceCosmosDBAccountIDRef *v1.Reference `json:"sourceCosmosDbAccountIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceCosmosDBAccountIDSelector *v1.Selector `json:"sourceCosmosDbAccountIdSelector,omitempty" tf:"-"`
}

type VirtualNetworkRuleObservation struct {
}

type VirtualNetworkRuleParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Account is the Schema for the Accounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountSpec   `json:"spec"`
	Status            AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
