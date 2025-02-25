//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azblob

type AccessTier string

const (
	AccessTierArchive AccessTier = "Archive"
	AccessTierCool    AccessTier = "Cool"
	AccessTierHot     AccessTier = "Hot"
	AccessTierP10     AccessTier = "P10"
	AccessTierP15     AccessTier = "P15"
	AccessTierP20     AccessTier = "P20"
	AccessTierP30     AccessTier = "P30"
	AccessTierP4      AccessTier = "P4"
	AccessTierP40     AccessTier = "P40"
	AccessTierP50     AccessTier = "P50"
	AccessTierP6      AccessTier = "P6"
	AccessTierP60     AccessTier = "P60"
	AccessTierP70     AccessTier = "P70"
	AccessTierP80     AccessTier = "P80"
)

// PossibleAccessTierValues returns the possible values for the AccessTier const type.
func PossibleAccessTierValues() []AccessTier {
	return []AccessTier{
		AccessTierArchive,
		AccessTierCool,
		AccessTierHot,
		AccessTierP10,
		AccessTierP15,
		AccessTierP20,
		AccessTierP30,
		AccessTierP4,
		AccessTierP40,
		AccessTierP50,
		AccessTierP6,
		AccessTierP60,
		AccessTierP70,
		AccessTierP80,
	}
}

type AccountKind string

const (
	AccountKindStorage          AccountKind = "Storage"
	AccountKindBlobStorage      AccountKind = "BlobStorage"
	AccountKindStorageV2        AccountKind = "StorageV2"
	AccountKindFileStorage      AccountKind = "FileStorage"
	AccountKindBlockBlobStorage AccountKind = "BlockBlobStorage"
)

// PossibleAccountKindValues returns the possible values for the AccountKind const type.
func PossibleAccountKindValues() []AccountKind {
	return []AccountKind{
		AccountKindStorage,
		AccountKindBlobStorage,
		AccountKindStorageV2,
		AccountKindFileStorage,
		AccountKindBlockBlobStorage,
	}
}

type ArchiveStatus string

const (
	ArchiveStatusRehydratePendingToCool ArchiveStatus = "rehydrate-pending-to-cool"
	ArchiveStatusRehydratePendingToHot  ArchiveStatus = "rehydrate-pending-to-hot"
)

// PossibleArchiveStatusValues returns the possible values for the ArchiveStatus const type.
func PossibleArchiveStatusValues() []ArchiveStatus {
	return []ArchiveStatus{
		ArchiveStatusRehydratePendingToCool,
		ArchiveStatusRehydratePendingToHot,
	}
}

type BlobExpiryOptions string

const (
	BlobExpiryOptionsAbsolute           BlobExpiryOptions = "Absolute"
	BlobExpiryOptionsNeverExpire        BlobExpiryOptions = "NeverExpire"
	BlobExpiryOptionsRelativeToCreation BlobExpiryOptions = "RelativeToCreation"
	BlobExpiryOptionsRelativeToNow      BlobExpiryOptions = "RelativeToNow"
)

// PossibleBlobExpiryOptionsValues returns the possible values for the BlobExpiryOptions const type.
func PossibleBlobExpiryOptionsValues() []BlobExpiryOptions {
	return []BlobExpiryOptions{
		BlobExpiryOptionsAbsolute,
		BlobExpiryOptionsNeverExpire,
		BlobExpiryOptionsRelativeToCreation,
		BlobExpiryOptionsRelativeToNow,
	}
}

type BlobImmutabilityPolicyMode string

const (
	BlobImmutabilityPolicyModeUnlocked BlobImmutabilityPolicyMode = "Unlocked"
	BlobImmutabilityPolicyModeLocked   BlobImmutabilityPolicyMode = "Locked"
	BlobImmutabilityPolicyModeMutable  BlobImmutabilityPolicyMode = "Mutable"
)

// PossibleBlobImmutabilityPolicyModeValues returns the possible values for the BlobImmutabilityPolicyMode const type.
func PossibleBlobImmutabilityPolicyModeValues() []BlobImmutabilityPolicyMode {
	return []BlobImmutabilityPolicyMode{
		BlobImmutabilityPolicyModeUnlocked,
		BlobImmutabilityPolicyModeLocked,
		BlobImmutabilityPolicyModeMutable,
	}
}

type BlobType string

const (
	BlobTypeBlockBlob  BlobType = "BlockBlob"
	BlobTypePageBlob   BlobType = "PageBlob"
	BlobTypeAppendBlob BlobType = "AppendBlob"
)

// PossibleBlobTypeValues returns the possible values for the BlobType const type.
func PossibleBlobTypeValues() []BlobType {
	return []BlobType{
		BlobTypeBlockBlob,
		BlobTypePageBlob,
		BlobTypeAppendBlob,
	}
}

type BlockListType string

const (
	BlockListTypeCommitted   BlockListType = "committed"
	BlockListTypeUncommitted BlockListType = "uncommitted"
	BlockListTypeAll         BlockListType = "all"
)

// PossibleBlockListTypeValues returns the possible values for the BlockListType const type.
func PossibleBlockListTypeValues() []BlockListType {
	return []BlockListType{
		BlockListTypeCommitted,
		BlockListTypeUncommitted,
		BlockListTypeAll,
	}
}

type CopyStatusType string

const (
	CopyStatusTypePending CopyStatusType = "pending"
	CopyStatusTypeSuccess CopyStatusType = "success"
	CopyStatusTypeAborted CopyStatusType = "aborted"
	CopyStatusTypeFailed  CopyStatusType = "failed"
)

// PossibleCopyStatusTypeValues returns the possible values for the CopyStatusType const type.
func PossibleCopyStatusTypeValues() []CopyStatusType {
	return []CopyStatusType{
		CopyStatusTypePending,
		CopyStatusTypeSuccess,
		CopyStatusTypeAborted,
		CopyStatusTypeFailed,
	}
}

type DeleteSnapshotsOptionType string

const (
	DeleteSnapshotsOptionTypeInclude DeleteSnapshotsOptionType = "include"
	DeleteSnapshotsOptionTypeOnly    DeleteSnapshotsOptionType = "only"
)

// PossibleDeleteSnapshotsOptionTypeValues returns the possible values for the DeleteSnapshotsOptionType const type.
func PossibleDeleteSnapshotsOptionTypeValues() []DeleteSnapshotsOptionType {
	return []DeleteSnapshotsOptionType{
		DeleteSnapshotsOptionTypeInclude,
		DeleteSnapshotsOptionTypeOnly,
	}
}

type Enum0 string

const (
	Enum0Service Enum0 = "service"
)

// PossibleEnum0Values returns the possible values for the Enum0 const type.
func PossibleEnum0Values() []Enum0 {
	return []Enum0{
		Enum0Service,
	}
}

type Enum1 string

const (
	Enum1Properties Enum1 = "properties"
)

// PossibleEnum1Values returns the possible values for the Enum1 const type.
func PossibleEnum1Values() []Enum1 {
	return []Enum1{
		Enum1Properties,
	}
}

type Enum10 string

const (
	Enum10Blobs Enum10 = "blobs"
)

// PossibleEnum10Values returns the possible values for the Enum10 const type.
func PossibleEnum10Values() []Enum10 {
	return []Enum10{
		Enum10Blobs,
	}
}

type Enum11 string

const (
	Enum11Container Enum11 = "container"
)

// PossibleEnum11Values returns the possible values for the Enum11 const type.
func PossibleEnum11Values() []Enum11 {
	return []Enum11{
		Enum11Container,
	}
}

type Enum12 string

const (
	Enum12Metadata Enum12 = "metadata"
)

// PossibleEnum12Values returns the possible values for the Enum12 const type.
func PossibleEnum12Values() []Enum12 {
	return []Enum12{
		Enum12Metadata,
	}
}

type Enum13 string

const (
	Enum13ACL Enum13 = "acl"
)

// PossibleEnum13Values returns the possible values for the Enum13 const type.
func PossibleEnum13Values() []Enum13 {
	return []Enum13{
		Enum13ACL,
	}
}

type Enum14 string

const (
	Enum14Undelete Enum14 = "undelete"
)

// PossibleEnum14Values returns the possible values for the Enum14 const type.
func PossibleEnum14Values() []Enum14 {
	return []Enum14{
		Enum14Undelete,
	}
}

type Enum15 string

const (
	Enum15Rename Enum15 = "rename"
)

// PossibleEnum15Values returns the possible values for the Enum15 const type.
func PossibleEnum15Values() []Enum15 {
	return []Enum15{
		Enum15Rename,
	}
}

type Enum16 string

const (
	Enum16Lease Enum16 = "lease"
)

// PossibleEnum16Values returns the possible values for the Enum16 const type.
func PossibleEnum16Values() []Enum16 {
	return []Enum16{
		Enum16Lease,
	}
}

type Enum2 string

const (
	Enum2TwoThousandTwenty0612 Enum2 = "2020-06-12"
)

// PossibleEnum2Values returns the possible values for the Enum2 const type.
func PossibleEnum2Values() []Enum2 {
	return []Enum2{
		Enum2TwoThousandTwenty0612,
	}
}

type Enum20 string

const (
	Enum20Directory Enum20 = "directory"
)

// PossibleEnum20Values returns the possible values for the Enum20 const type.
func PossibleEnum20Values() []Enum20 {
	return []Enum20{
		Enum20Directory,
	}
}

type Enum21 string

const (
	Enum21SetAccessControl Enum21 = "setAccessControl"
)

// PossibleEnum21Values returns the possible values for the Enum21 const type.
func PossibleEnum21Values() []Enum21 {
	return []Enum21{
		Enum21SetAccessControl,
	}
}

type Enum22 string

const (
	Enum22GetAccessControl Enum22 = "getAccessControl"
)

// PossibleEnum22Values returns the possible values for the Enum22 const type.
func PossibleEnum22Values() []Enum22 {
	return []Enum22{
		Enum22GetAccessControl,
	}
}

type Enum24 string

const (
	Enum24Expiry Enum24 = "expiry"
)

// PossibleEnum24Values returns the possible values for the Enum24 const type.
func PossibleEnum24Values() []Enum24 {
	return []Enum24{
		Enum24Expiry,
	}
}

type Enum26 string

const (
	Enum26ImmutabilityPolicies Enum26 = "immutabilityPolicies"
)

// PossibleEnum26Values returns the possible values for the Enum26 const type.
func PossibleEnum26Values() []Enum26 {
	return []Enum26{
		Enum26ImmutabilityPolicies,
	}
}

type Enum27 string

const (
	Enum27Legalhold Enum27 = "legalhold"
)

// PossibleEnum27Values returns the possible values for the Enum27 const type.
func PossibleEnum27Values() []Enum27 {
	return []Enum27{
		Enum27Legalhold,
	}
}

type Enum28 string

const (
	Enum28Snapshot Enum28 = "snapshot"
)

// PossibleEnum28Values returns the possible values for the Enum28 const type.
func PossibleEnum28Values() []Enum28 {
	return []Enum28{
		Enum28Snapshot,
	}
}

type Enum29 string

const (
	Enum29True Enum29 = "true"
)

// PossibleEnum29Values returns the possible values for the Enum29 const type.
func PossibleEnum29Values() []Enum29 {
	return []Enum29{
		Enum29True,
	}
}

type Enum3 string

const (
	Enum3Stats Enum3 = "stats"
)

// PossibleEnum3Values returns the possible values for the Enum3 const type.
func PossibleEnum3Values() []Enum3 {
	return []Enum3{
		Enum3Stats,
	}
}

type Enum30 string

const (
	Enum30Copy Enum30 = "copy"
)

// PossibleEnum30Values returns the possible values for the Enum30 const type.
func PossibleEnum30Values() []Enum30 {
	return []Enum30{
		Enum30Copy,
	}
}

type Enum31 string

const (
	Enum31Abort Enum31 = "abort"
)

// PossibleEnum31Values returns the possible values for the Enum31 const type.
func PossibleEnum31Values() []Enum31 {
	return []Enum31{
		Enum31Abort,
	}
}

type Enum32 string

const (
	Enum32Tier Enum32 = "tier"
)

// PossibleEnum32Values returns the possible values for the Enum32 const type.
func PossibleEnum32Values() []Enum32 {
	return []Enum32{
		Enum32Tier,
	}
}

type Enum33 string

const (
	Enum33Block Enum33 = "block"
)

// PossibleEnum33Values returns the possible values for the Enum33 const type.
func PossibleEnum33Values() []Enum33 {
	return []Enum33{
		Enum33Block,
	}
}

type Enum34 string

const (
	Enum34Blocklist Enum34 = "blocklist"
)

// PossibleEnum34Values returns the possible values for the Enum34 const type.
func PossibleEnum34Values() []Enum34 {
	return []Enum34{
		Enum34Blocklist,
	}
}

type Enum35 string

const (
	Enum35Page Enum35 = "page"
)

// PossibleEnum35Values returns the possible values for the Enum35 const type.
func PossibleEnum35Values() []Enum35 {
	return []Enum35{
		Enum35Page,
	}
}

type Enum36 string

const (
	Enum36Pagelist Enum36 = "pagelist"
)

// PossibleEnum36Values returns the possible values for the Enum36 const type.
func PossibleEnum36Values() []Enum36 {
	return []Enum36{
		Enum36Pagelist,
	}
}

type Enum37 string

const (
	Enum37Incrementalcopy Enum37 = "incrementalcopy"
)

// PossibleEnum37Values returns the possible values for the Enum37 const type.
func PossibleEnum37Values() []Enum37 {
	return []Enum37{
		Enum37Incrementalcopy,
	}
}

type Enum38 string

const (
	Enum38Appendblock Enum38 = "appendblock"
)

// PossibleEnum38Values returns the possible values for the Enum38 const type.
func PossibleEnum38Values() []Enum38 {
	return []Enum38{
		Enum38Appendblock,
	}
}

type Enum39 string

const (
	Enum39Seal Enum39 = "seal"
)

// PossibleEnum39Values returns the possible values for the Enum39 const type.
func PossibleEnum39Values() []Enum39 {
	return []Enum39{
		Enum39Seal,
	}
}

type Enum40 string

const (
	Enum40Query Enum40 = "query"
)

// PossibleEnum40Values returns the possible values for the Enum40 const type.
func PossibleEnum40Values() []Enum40 {
	return []Enum40{
		Enum40Query,
	}
}

type Enum42 string

const (
	Enum42Tags Enum42 = "tags"
)

// PossibleEnum42Values returns the possible values for the Enum42 const type.
func PossibleEnum42Values() []Enum42 {
	return []Enum42{
		Enum42Tags,
	}
}

type Enum5 string

const (
	Enum5List Enum5 = "list"
)

// PossibleEnum5Values returns the possible values for the Enum5 const type.
func PossibleEnum5Values() []Enum5 {
	return []Enum5{
		Enum5List,
	}
}

type Enum7 string

const (
	Enum7Userdelegationkey Enum7 = "userdelegationkey"
)

// PossibleEnum7Values returns the possible values for the Enum7 const type.
func PossibleEnum7Values() []Enum7 {
	return []Enum7{
		Enum7Userdelegationkey,
	}
}

type Enum8 string

const (
	Enum8Account Enum8 = "account"
)

// PossibleEnum8Values returns the possible values for the Enum8 const type.
func PossibleEnum8Values() []Enum8 {
	return []Enum8{
		Enum8Account,
	}
}

type Enum9 string

const (
	Enum9Batch Enum9 = "batch"
)

// PossibleEnum9Values returns the possible values for the Enum9 const type.
func PossibleEnum9Values() []Enum9 {
	return []Enum9{
		Enum9Batch,
	}
}

// GeoReplicationStatusType - The status of the secondary location
type GeoReplicationStatusType string

const (
	GeoReplicationStatusTypeBootstrap   GeoReplicationStatusType = "bootstrap"
	GeoReplicationStatusTypeLive        GeoReplicationStatusType = "live"
	GeoReplicationStatusTypeUnavailable GeoReplicationStatusType = "unavailable"
)

// PossibleGeoReplicationStatusTypeValues returns the possible values for the GeoReplicationStatusType const type.
func PossibleGeoReplicationStatusTypeValues() []GeoReplicationStatusType {
	return []GeoReplicationStatusType{
		GeoReplicationStatusTypeBootstrap,
		GeoReplicationStatusTypeLive,
		GeoReplicationStatusTypeUnavailable,
	}
}

type LeaseDurationType string

const (
	LeaseDurationTypeInfinite LeaseDurationType = "infinite"
	LeaseDurationTypeFixed    LeaseDurationType = "fixed"
)

// PossibleLeaseDurationTypeValues returns the possible values for the LeaseDurationType const type.
func PossibleLeaseDurationTypeValues() []LeaseDurationType {
	return []LeaseDurationType{
		LeaseDurationTypeInfinite,
		LeaseDurationTypeFixed,
	}
}

type LeaseStateType string

const (
	LeaseStateTypeAvailable LeaseStateType = "available"
	LeaseStateTypeLeased    LeaseStateType = "leased"
	LeaseStateTypeExpired   LeaseStateType = "expired"
	LeaseStateTypeBreaking  LeaseStateType = "breaking"
	LeaseStateTypeBroken    LeaseStateType = "broken"
)

// PossibleLeaseStateTypeValues returns the possible values for the LeaseStateType const type.
func PossibleLeaseStateTypeValues() []LeaseStateType {
	return []LeaseStateType{
		LeaseStateTypeAvailable,
		LeaseStateTypeLeased,
		LeaseStateTypeExpired,
		LeaseStateTypeBreaking,
		LeaseStateTypeBroken,
	}
}

type LeaseStatusType string

const (
	LeaseStatusTypeLocked   LeaseStatusType = "locked"
	LeaseStatusTypeUnlocked LeaseStatusType = "unlocked"
)

// PossibleLeaseStatusTypeValues returns the possible values for the LeaseStatusType const type.
func PossibleLeaseStatusTypeValues() []LeaseStatusType {
	return []LeaseStatusType{
		LeaseStatusTypeLocked,
		LeaseStatusTypeUnlocked,
	}
}

type ListBlobsIncludeItem string

const (
	ListBlobsIncludeItemCopy               ListBlobsIncludeItem = "copy"
	ListBlobsIncludeItemDeleted            ListBlobsIncludeItem = "deleted"
	ListBlobsIncludeItemMetadata           ListBlobsIncludeItem = "metadata"
	ListBlobsIncludeItemSnapshots          ListBlobsIncludeItem = "snapshots"
	ListBlobsIncludeItemUncommittedblobs   ListBlobsIncludeItem = "uncommittedblobs"
	ListBlobsIncludeItemVersions           ListBlobsIncludeItem = "versions"
	ListBlobsIncludeItemTags               ListBlobsIncludeItem = "tags"
	ListBlobsIncludeItemImmutabilitypolicy ListBlobsIncludeItem = "immutabilitypolicy"
	ListBlobsIncludeItemLegalhold          ListBlobsIncludeItem = "legalhold"
)

// PossibleListBlobsIncludeItemValues returns the possible values for the ListBlobsIncludeItem const type.
func PossibleListBlobsIncludeItemValues() []ListBlobsIncludeItem {
	return []ListBlobsIncludeItem{
		ListBlobsIncludeItemCopy,
		ListBlobsIncludeItemDeleted,
		ListBlobsIncludeItemMetadata,
		ListBlobsIncludeItemSnapshots,
		ListBlobsIncludeItemUncommittedblobs,
		ListBlobsIncludeItemVersions,
		ListBlobsIncludeItemTags,
		ListBlobsIncludeItemImmutabilitypolicy,
		ListBlobsIncludeItemLegalhold,
	}
}

type ListContainersIncludeType string

const (
	ListContainersIncludeTypeMetadata ListContainersIncludeType = "metadata"
	ListContainersIncludeTypeDeleted  ListContainersIncludeType = "deleted"
)

// PossibleListContainersIncludeTypeValues returns the possible values for the ListContainersIncludeType const type.
func PossibleListContainersIncludeTypeValues() []ListContainersIncludeType {
	return []ListContainersIncludeType{
		ListContainersIncludeTypeMetadata,
		ListContainersIncludeTypeDeleted,
	}
}

type PathRenameMode string

const (
	PathRenameModeLegacy PathRenameMode = "legacy"
	PathRenameModePosix  PathRenameMode = "posix"
)

// PossiblePathRenameModeValues returns the possible values for the PathRenameMode const type.
func PossiblePathRenameModeValues() []PathRenameMode {
	return []PathRenameMode{
		PathRenameModeLegacy,
		PathRenameModePosix,
	}
}

type PremiumPageBlobAccessTier string

const (
	PremiumPageBlobAccessTierP10 PremiumPageBlobAccessTier = "P10"
	PremiumPageBlobAccessTierP15 PremiumPageBlobAccessTier = "P15"
	PremiumPageBlobAccessTierP20 PremiumPageBlobAccessTier = "P20"
	PremiumPageBlobAccessTierP30 PremiumPageBlobAccessTier = "P30"
	PremiumPageBlobAccessTierP4  PremiumPageBlobAccessTier = "P4"
	PremiumPageBlobAccessTierP40 PremiumPageBlobAccessTier = "P40"
	PremiumPageBlobAccessTierP50 PremiumPageBlobAccessTier = "P50"
	PremiumPageBlobAccessTierP6  PremiumPageBlobAccessTier = "P6"
	PremiumPageBlobAccessTierP60 PremiumPageBlobAccessTier = "P60"
	PremiumPageBlobAccessTierP70 PremiumPageBlobAccessTier = "P70"
	PremiumPageBlobAccessTierP80 PremiumPageBlobAccessTier = "P80"
)

// PossiblePremiumPageBlobAccessTierValues returns the possible values for the PremiumPageBlobAccessTier const type.
func PossiblePremiumPageBlobAccessTierValues() []PremiumPageBlobAccessTier {
	return []PremiumPageBlobAccessTier{
		PremiumPageBlobAccessTierP10,
		PremiumPageBlobAccessTierP15,
		PremiumPageBlobAccessTierP20,
		PremiumPageBlobAccessTierP30,
		PremiumPageBlobAccessTierP4,
		PremiumPageBlobAccessTierP40,
		PremiumPageBlobAccessTierP50,
		PremiumPageBlobAccessTierP6,
		PremiumPageBlobAccessTierP60,
		PremiumPageBlobAccessTierP70,
		PremiumPageBlobAccessTierP80,
	}
}

type PublicAccessType string

const (
	PublicAccessTypeBlob      PublicAccessType = "blob"
	PublicAccessTypeContainer PublicAccessType = "container"
)

// PossiblePublicAccessTypeValues returns the possible values for the PublicAccessType const type.
func PossiblePublicAccessTypeValues() []PublicAccessType {
	return []PublicAccessType{
		PublicAccessTypeBlob,
		PublicAccessTypeContainer,
	}
}

// QueryFormatType - The quick query format type.
type QueryFormatType string

const (
	QueryFormatTypeDelimited QueryFormatType = "delimited"
	QueryFormatTypeJSON      QueryFormatType = "json"
	QueryFormatTypeArrow     QueryFormatType = "arrow"
)

// PossibleQueryFormatTypeValues returns the possible values for the QueryFormatType const type.
func PossibleQueryFormatTypeValues() []QueryFormatType {
	return []QueryFormatType{
		QueryFormatTypeDelimited,
		QueryFormatTypeJSON,
		QueryFormatTypeArrow,
	}
}

// QueryRequestQueryType - the query type
type QueryRequestQueryType string

const (
	QueryRequestQueryTypeSQL QueryRequestQueryType = "SQL"
)

// PossibleQueryRequestQueryTypeValues returns the possible values for the QueryRequestQueryType const type.
func PossibleQueryRequestQueryTypeValues() []QueryRequestQueryType {
	return []QueryRequestQueryType{
		QueryRequestQueryTypeSQL,
	}
}

// RehydratePriority - If an object is in rehydrate pending state then this header is returned with priority of rehydrate.
// Valid values are High and Standard.
type RehydratePriority string

const (
	RehydratePriorityHigh     RehydratePriority = "High"
	RehydratePriorityStandard RehydratePriority = "Standard"
)

// PossibleRehydratePriorityValues returns the possible values for the RehydratePriority const type.
func PossibleRehydratePriorityValues() []RehydratePriority {
	return []RehydratePriority{
		RehydratePriorityHigh,
		RehydratePriorityStandard,
	}
}

type SKUName string

const (
	SKUNameStandardLRS   SKUName = "Standard_LRS"
	SKUNameStandardGRS   SKUName = "Standard_GRS"
	SKUNameStandardRAGRS SKUName = "Standard_RAGRS"
	SKUNameStandardZRS   SKUName = "Standard_ZRS"
	SKUNamePremiumLRS    SKUName = "Premium_LRS"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameStandardLRS,
		SKUNameStandardGRS,
		SKUNameStandardRAGRS,
		SKUNameStandardZRS,
		SKUNamePremiumLRS,
	}
}

type SequenceNumberActionType string

const (
	SequenceNumberActionTypeMax       SequenceNumberActionType = "max"
	SequenceNumberActionTypeUpdate    SequenceNumberActionType = "update"
	SequenceNumberActionTypeIncrement SequenceNumberActionType = "increment"
)

// PossibleSequenceNumberActionTypeValues returns the possible values for the SequenceNumberActionType const type.
func PossibleSequenceNumberActionTypeValues() []SequenceNumberActionType {
	return []SequenceNumberActionType{
		SequenceNumberActionTypeMax,
		SequenceNumberActionTypeUpdate,
		SequenceNumberActionTypeIncrement,
	}
}

// StorageErrorCode - Error codes returned by the service
type StorageErrorCode string

const (
	StorageErrorCodeAccountAlreadyExists                              StorageErrorCode = "AccountAlreadyExists"
	StorageErrorCodeAccountBeingCreated                               StorageErrorCode = "AccountBeingCreated"
	StorageErrorCodeAccountIsDisabled                                 StorageErrorCode = "AccountIsDisabled"
	StorageErrorCodeAppendPositionConditionNotMet                     StorageErrorCode = "AppendPositionConditionNotMet"
	StorageErrorCodeAuthenticationFailed                              StorageErrorCode = "AuthenticationFailed"
	StorageErrorCodeAuthorizationFailure                              StorageErrorCode = "AuthorizationFailure"
	StorageErrorCodeAuthorizationPermissionMismatch                   StorageErrorCode = "AuthorizationPermissionMismatch"
	StorageErrorCodeAuthorizationProtocolMismatch                     StorageErrorCode = "AuthorizationProtocolMismatch"
	StorageErrorCodeAuthorizationResourceTypeMismatch                 StorageErrorCode = "AuthorizationResourceTypeMismatch"
	StorageErrorCodeAuthorizationServiceMismatch                      StorageErrorCode = "AuthorizationServiceMismatch"
	StorageErrorCodeAuthorizationSourceIPMismatch                     StorageErrorCode = "AuthorizationSourceIPMismatch"
	StorageErrorCodeBlobAlreadyExists                                 StorageErrorCode = "BlobAlreadyExists"
	StorageErrorCodeBlobArchived                                      StorageErrorCode = "BlobArchived"
	StorageErrorCodeBlobBeingRehydrated                               StorageErrorCode = "BlobBeingRehydrated"
	StorageErrorCodeBlobImmutableDueToPolicy                          StorageErrorCode = "BlobImmutableDueToPolicy"
	StorageErrorCodeBlobNotArchived                                   StorageErrorCode = "BlobNotArchived"
	StorageErrorCodeBlobNotFound                                      StorageErrorCode = "BlobNotFound"
	StorageErrorCodeBlobOverwritten                                   StorageErrorCode = "BlobOverwritten"
	StorageErrorCodeBlobTierInadequateForContentLength                StorageErrorCode = "BlobTierInadequateForContentLength"
	StorageErrorCodeBlobUsesCustomerSpecifiedEncryption               StorageErrorCode = "BlobUsesCustomerSpecifiedEncryption"
	StorageErrorCodeBlockCountExceedsLimit                            StorageErrorCode = "BlockCountExceedsLimit"
	StorageErrorCodeBlockListTooLong                                  StorageErrorCode = "BlockListTooLong"
	StorageErrorCodeCannotChangeToLowerTier                           StorageErrorCode = "CannotChangeToLowerTier"
	StorageErrorCodeCannotVerifyCopySource                            StorageErrorCode = "CannotVerifyCopySource"
	StorageErrorCodeConditionHeadersNotSupported                      StorageErrorCode = "ConditionHeadersNotSupported"
	StorageErrorCodeConditionNotMet                                   StorageErrorCode = "ConditionNotMet"
	StorageErrorCodeContainerAlreadyExists                            StorageErrorCode = "ContainerAlreadyExists"
	StorageErrorCodeContainerBeingDeleted                             StorageErrorCode = "ContainerBeingDeleted"
	StorageErrorCodeContainerDisabled                                 StorageErrorCode = "ContainerDisabled"
	StorageErrorCodeContainerNotFound                                 StorageErrorCode = "ContainerNotFound"
	StorageErrorCodeContentLengthLargerThanTierLimit                  StorageErrorCode = "ContentLengthLargerThanTierLimit"
	StorageErrorCodeCopyAcrossAccountsNotSupported                    StorageErrorCode = "CopyAcrossAccountsNotSupported"
	StorageErrorCodeCopyIDMismatch                                    StorageErrorCode = "CopyIdMismatch"
	StorageErrorCodeEmptyMetadataKey                                  StorageErrorCode = "EmptyMetadataKey"
	StorageErrorCodeFeatureVersionMismatch                            StorageErrorCode = "FeatureVersionMismatch"
	StorageErrorCodeIncrementalCopyBlobMismatch                       StorageErrorCode = "IncrementalCopyBlobMismatch"
	StorageErrorCodeIncrementalCopyOfEralierVersionSnapshotNotAllowed StorageErrorCode = "IncrementalCopyOfEralierVersionSnapshotNotAllowed"
	StorageErrorCodeIncrementalCopySourceMustBeSnapshot               StorageErrorCode = "IncrementalCopySourceMustBeSnapshot"
	StorageErrorCodeInfiniteLeaseDurationRequired                     StorageErrorCode = "InfiniteLeaseDurationRequired"
	StorageErrorCodeInsufficientAccountPermissions                    StorageErrorCode = "InsufficientAccountPermissions"
	StorageErrorCodeInternalError                                     StorageErrorCode = "InternalError"
	StorageErrorCodeInvalidAuthenticationInfo                         StorageErrorCode = "InvalidAuthenticationInfo"
	StorageErrorCodeInvalidBlobOrBlock                                StorageErrorCode = "InvalidBlobOrBlock"
	StorageErrorCodeInvalidBlobTier                                   StorageErrorCode = "InvalidBlobTier"
	StorageErrorCodeInvalidBlobType                                   StorageErrorCode = "InvalidBlobType"
	StorageErrorCodeInvalidBlockID                                    StorageErrorCode = "InvalidBlockId"
	StorageErrorCodeInvalidBlockList                                  StorageErrorCode = "InvalidBlockList"
	StorageErrorCodeInvalidHTTPVerb                                   StorageErrorCode = "InvalidHttpVerb"
	StorageErrorCodeInvalidHeaderValue                                StorageErrorCode = "InvalidHeaderValue"
	StorageErrorCodeInvalidInput                                      StorageErrorCode = "InvalidInput"
	StorageErrorCodeInvalidMD5                                        StorageErrorCode = "InvalidMd5"
	StorageErrorCodeInvalidMetadata                                   StorageErrorCode = "InvalidMetadata"
	StorageErrorCodeInvalidOperation                                  StorageErrorCode = "InvalidOperation"
	StorageErrorCodeInvalidPageRange                                  StorageErrorCode = "InvalidPageRange"
	StorageErrorCodeInvalidQueryParameterValue                        StorageErrorCode = "InvalidQueryParameterValue"
	StorageErrorCodeInvalidRange                                      StorageErrorCode = "InvalidRange"
	StorageErrorCodeInvalidResourceName                               StorageErrorCode = "InvalidResourceName"
	StorageErrorCodeInvalidSourceBlobType                             StorageErrorCode = "InvalidSourceBlobType"
	StorageErrorCodeInvalidSourceBlobURL                              StorageErrorCode = "InvalidSourceBlobUrl"
	StorageErrorCodeInvalidURI                                        StorageErrorCode = "InvalidUri"
	StorageErrorCodeInvalidVersionForPageBlobOperation                StorageErrorCode = "InvalidVersionForPageBlobOperation"
	StorageErrorCodeInvalidXMLDocument                                StorageErrorCode = "InvalidXmlDocument"
	StorageErrorCodeInvalidXMLNodeValue                               StorageErrorCode = "InvalidXmlNodeValue"
	StorageErrorCodeLeaseAlreadyBroken                                StorageErrorCode = "LeaseAlreadyBroken"
	StorageErrorCodeLeaseAlreadyPresent                               StorageErrorCode = "LeaseAlreadyPresent"
	StorageErrorCodeLeaseIDMismatchWithBlobOperation                  StorageErrorCode = "LeaseIdMismatchWithBlobOperation"
	StorageErrorCodeLeaseIDMismatchWithContainerOperation             StorageErrorCode = "LeaseIdMismatchWithContainerOperation"
	StorageErrorCodeLeaseIDMismatchWithLeaseOperation                 StorageErrorCode = "LeaseIdMismatchWithLeaseOperation"
	StorageErrorCodeLeaseIDMissing                                    StorageErrorCode = "LeaseIdMissing"
	StorageErrorCodeLeaseIsBreakingAndCannotBeAcquired                StorageErrorCode = "LeaseIsBreakingAndCannotBeAcquired"
	StorageErrorCodeLeaseIsBreakingAndCannotBeChanged                 StorageErrorCode = "LeaseIsBreakingAndCannotBeChanged"
	StorageErrorCodeLeaseIsBrokenAndCannotBeRenewed                   StorageErrorCode = "LeaseIsBrokenAndCannotBeRenewed"
	StorageErrorCodeLeaseLost                                         StorageErrorCode = "LeaseLost"
	StorageErrorCodeLeaseNotPresentWithBlobOperation                  StorageErrorCode = "LeaseNotPresentWithBlobOperation"
	StorageErrorCodeLeaseNotPresentWithContainerOperation             StorageErrorCode = "LeaseNotPresentWithContainerOperation"
	StorageErrorCodeLeaseNotPresentWithLeaseOperation                 StorageErrorCode = "LeaseNotPresentWithLeaseOperation"
	StorageErrorCodeMD5Mismatch                                       StorageErrorCode = "Md5Mismatch"
	StorageErrorCodeMaxBlobSizeConditionNotMet                        StorageErrorCode = "MaxBlobSizeConditionNotMet"
	StorageErrorCodeMetadataTooLarge                                  StorageErrorCode = "MetadataTooLarge"
	StorageErrorCodeMissingContentLengthHeader                        StorageErrorCode = "MissingContentLengthHeader"
	StorageErrorCodeMissingRequiredHeader                             StorageErrorCode = "MissingRequiredHeader"
	StorageErrorCodeMissingRequiredQueryParameter                     StorageErrorCode = "MissingRequiredQueryParameter"
	StorageErrorCodeMissingRequiredXMLNode                            StorageErrorCode = "MissingRequiredXmlNode"
	StorageErrorCodeMultipleConditionHeadersNotSupported              StorageErrorCode = "MultipleConditionHeadersNotSupported"
	StorageErrorCodeNoAuthenticationInformation                       StorageErrorCode = "NoAuthenticationInformation"
	StorageErrorCodeNoPendingCopyOperation                            StorageErrorCode = "NoPendingCopyOperation"
	StorageErrorCodeOperationNotAllowedOnIncrementalCopyBlob          StorageErrorCode = "OperationNotAllowedOnIncrementalCopyBlob"
	StorageErrorCodeOperationTimedOut                                 StorageErrorCode = "OperationTimedOut"
	StorageErrorCodeOutOfRangeInput                                   StorageErrorCode = "OutOfRangeInput"
	StorageErrorCodeOutOfRangeQueryParameterValue                     StorageErrorCode = "OutOfRangeQueryParameterValue"
	StorageErrorCodePendingCopyOperation                              StorageErrorCode = "PendingCopyOperation"
	StorageErrorCodePreviousSnapshotCannotBeNewer                     StorageErrorCode = "PreviousSnapshotCannotBeNewer"
	StorageErrorCodePreviousSnapshotNotFound                          StorageErrorCode = "PreviousSnapshotNotFound"
	StorageErrorCodePreviousSnapshotOperationNotSupported             StorageErrorCode = "PreviousSnapshotOperationNotSupported"
	StorageErrorCodeRequestBodyTooLarge                               StorageErrorCode = "RequestBodyTooLarge"
	StorageErrorCodeRequestURLFailedToParse                           StorageErrorCode = "RequestUrlFailedToParse"
	StorageErrorCodeResourceAlreadyExists                             StorageErrorCode = "ResourceAlreadyExists"
	StorageErrorCodeResourceNotFound                                  StorageErrorCode = "ResourceNotFound"
	StorageErrorCodeResourceTypeMismatch                              StorageErrorCode = "ResourceTypeMismatch"
	StorageErrorCodeSequenceNumberConditionNotMet                     StorageErrorCode = "SequenceNumberConditionNotMet"
	StorageErrorCodeSequenceNumberIncrementTooLarge                   StorageErrorCode = "SequenceNumberIncrementTooLarge"
	StorageErrorCodeServerBusy                                        StorageErrorCode = "ServerBusy"
	StorageErrorCodeSnaphotOperationRateExceeded                      StorageErrorCode = "SnaphotOperationRateExceeded"
	StorageErrorCodeSnapshotCountExceeded                             StorageErrorCode = "SnapshotCountExceeded"
	StorageErrorCodeSnapshotsPresent                                  StorageErrorCode = "SnapshotsPresent"
	StorageErrorCodeSourceConditionNotMet                             StorageErrorCode = "SourceConditionNotMet"
	StorageErrorCodeSystemInUse                                       StorageErrorCode = "SystemInUse"
	StorageErrorCodeTargetConditionNotMet                             StorageErrorCode = "TargetConditionNotMet"
	StorageErrorCodeUnauthorizedBlobOverwrite                         StorageErrorCode = "UnauthorizedBlobOverwrite"
	StorageErrorCodeUnsupportedHTTPVerb                               StorageErrorCode = "UnsupportedHttpVerb"
	StorageErrorCodeUnsupportedHeader                                 StorageErrorCode = "UnsupportedHeader"
	StorageErrorCodeUnsupportedQueryParameter                         StorageErrorCode = "UnsupportedQueryParameter"
	StorageErrorCodeUnsupportedXMLNode                                StorageErrorCode = "UnsupportedXmlNode"
)

// PossibleStorageErrorCodeValues returns the possible values for the StorageErrorCode const type.
func PossibleStorageErrorCodeValues() []StorageErrorCode {
	return []StorageErrorCode{
		StorageErrorCodeAccountAlreadyExists,
		StorageErrorCodeAccountBeingCreated,
		StorageErrorCodeAccountIsDisabled,
		StorageErrorCodeAppendPositionConditionNotMet,
		StorageErrorCodeAuthenticationFailed,
		StorageErrorCodeAuthorizationFailure,
		StorageErrorCodeAuthorizationPermissionMismatch,
		StorageErrorCodeAuthorizationProtocolMismatch,
		StorageErrorCodeAuthorizationResourceTypeMismatch,
		StorageErrorCodeAuthorizationServiceMismatch,
		StorageErrorCodeAuthorizationSourceIPMismatch,
		StorageErrorCodeBlobAlreadyExists,
		StorageErrorCodeBlobArchived,
		StorageErrorCodeBlobBeingRehydrated,
		StorageErrorCodeBlobImmutableDueToPolicy,
		StorageErrorCodeBlobNotArchived,
		StorageErrorCodeBlobNotFound,
		StorageErrorCodeBlobOverwritten,
		StorageErrorCodeBlobTierInadequateForContentLength,
		StorageErrorCodeBlobUsesCustomerSpecifiedEncryption,
		StorageErrorCodeBlockCountExceedsLimit,
		StorageErrorCodeBlockListTooLong,
		StorageErrorCodeCannotChangeToLowerTier,
		StorageErrorCodeCannotVerifyCopySource,
		StorageErrorCodeConditionHeadersNotSupported,
		StorageErrorCodeConditionNotMet,
		StorageErrorCodeContainerAlreadyExists,
		StorageErrorCodeContainerBeingDeleted,
		StorageErrorCodeContainerDisabled,
		StorageErrorCodeContainerNotFound,
		StorageErrorCodeContentLengthLargerThanTierLimit,
		StorageErrorCodeCopyAcrossAccountsNotSupported,
		StorageErrorCodeCopyIDMismatch,
		StorageErrorCodeEmptyMetadataKey,
		StorageErrorCodeFeatureVersionMismatch,
		StorageErrorCodeIncrementalCopyBlobMismatch,
		StorageErrorCodeIncrementalCopyOfEralierVersionSnapshotNotAllowed,
		StorageErrorCodeIncrementalCopySourceMustBeSnapshot,
		StorageErrorCodeInfiniteLeaseDurationRequired,
		StorageErrorCodeInsufficientAccountPermissions,
		StorageErrorCodeInternalError,
		StorageErrorCodeInvalidAuthenticationInfo,
		StorageErrorCodeInvalidBlobOrBlock,
		StorageErrorCodeInvalidBlobTier,
		StorageErrorCodeInvalidBlobType,
		StorageErrorCodeInvalidBlockID,
		StorageErrorCodeInvalidBlockList,
		StorageErrorCodeInvalidHTTPVerb,
		StorageErrorCodeInvalidHeaderValue,
		StorageErrorCodeInvalidInput,
		StorageErrorCodeInvalidMD5,
		StorageErrorCodeInvalidMetadata,
		StorageErrorCodeInvalidOperation,
		StorageErrorCodeInvalidPageRange,
		StorageErrorCodeInvalidQueryParameterValue,
		StorageErrorCodeInvalidRange,
		StorageErrorCodeInvalidResourceName,
		StorageErrorCodeInvalidSourceBlobType,
		StorageErrorCodeInvalidSourceBlobURL,
		StorageErrorCodeInvalidURI,
		StorageErrorCodeInvalidVersionForPageBlobOperation,
		StorageErrorCodeInvalidXMLDocument,
		StorageErrorCodeInvalidXMLNodeValue,
		StorageErrorCodeLeaseAlreadyBroken,
		StorageErrorCodeLeaseAlreadyPresent,
		StorageErrorCodeLeaseIDMismatchWithBlobOperation,
		StorageErrorCodeLeaseIDMismatchWithContainerOperation,
		StorageErrorCodeLeaseIDMismatchWithLeaseOperation,
		StorageErrorCodeLeaseIDMissing,
		StorageErrorCodeLeaseIsBreakingAndCannotBeAcquired,
		StorageErrorCodeLeaseIsBreakingAndCannotBeChanged,
		StorageErrorCodeLeaseIsBrokenAndCannotBeRenewed,
		StorageErrorCodeLeaseLost,
		StorageErrorCodeLeaseNotPresentWithBlobOperation,
		StorageErrorCodeLeaseNotPresentWithContainerOperation,
		StorageErrorCodeLeaseNotPresentWithLeaseOperation,
		StorageErrorCodeMD5Mismatch,
		StorageErrorCodeMaxBlobSizeConditionNotMet,
		StorageErrorCodeMetadataTooLarge,
		StorageErrorCodeMissingContentLengthHeader,
		StorageErrorCodeMissingRequiredHeader,
		StorageErrorCodeMissingRequiredQueryParameter,
		StorageErrorCodeMissingRequiredXMLNode,
		StorageErrorCodeMultipleConditionHeadersNotSupported,
		StorageErrorCodeNoAuthenticationInformation,
		StorageErrorCodeNoPendingCopyOperation,
		StorageErrorCodeOperationNotAllowedOnIncrementalCopyBlob,
		StorageErrorCodeOperationTimedOut,
		StorageErrorCodeOutOfRangeInput,
		StorageErrorCodeOutOfRangeQueryParameterValue,
		StorageErrorCodePendingCopyOperation,
		StorageErrorCodePreviousSnapshotCannotBeNewer,
		StorageErrorCodePreviousSnapshotNotFound,
		StorageErrorCodePreviousSnapshotOperationNotSupported,
		StorageErrorCodeRequestBodyTooLarge,
		StorageErrorCodeRequestURLFailedToParse,
		StorageErrorCodeResourceAlreadyExists,
		StorageErrorCodeResourceNotFound,
		StorageErrorCodeResourceTypeMismatch,
		StorageErrorCodeSequenceNumberConditionNotMet,
		StorageErrorCodeSequenceNumberIncrementTooLarge,
		StorageErrorCodeServerBusy,
		StorageErrorCodeSnaphotOperationRateExceeded,
		StorageErrorCodeSnapshotCountExceeded,
		StorageErrorCodeSnapshotsPresent,
		StorageErrorCodeSourceConditionNotMet,
		StorageErrorCodeSystemInUse,
		StorageErrorCodeTargetConditionNotMet,
		StorageErrorCodeUnauthorizedBlobOverwrite,
		StorageErrorCodeUnsupportedHTTPVerb,
		StorageErrorCodeUnsupportedHeader,
		StorageErrorCodeUnsupportedQueryParameter,
		StorageErrorCodeUnsupportedXMLNode,
	}
}
