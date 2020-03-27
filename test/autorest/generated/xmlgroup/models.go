// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package xmlgroup

import (
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// An Access policy
type AccessPolicy struct {
	// the date-time the policy expires
	Expiry *time.Time `xml:"Expiry"`

	// the permissions for the acl policy
	Permission *string `xml:"Permission"`

	// the date-time the policy is active
	Start *time.Time `xml:"Start"`
}

func (a AccessPolicy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias:  (*alias)(&a),
		Expiry: (*timeRFC3339)(a.Expiry),
		Start:  (*timeRFC3339)(a.Start),
	}
	return e.EncodeElement(aux, start)
}

func (a *AccessPolicy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(a),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	a.Expiry = (*time.Time)(aux.Expiry)
	a.Start = (*time.Time)(aux.Start)
	return nil
}

// A barrel of apples.
type AppleBarrel struct {
	BadApples  *[]string `xml:"BadApples>Apple"`
	GoodApples *[]string `xml:"GoodApples>Apple"`
}

// AppleBarrelResponse is the response envelope for operations that return a AppleBarrel type.
type AppleBarrelResponse struct {
	// A barrel of apples.
	AppleBarrel *AppleBarrel `xml:"AppleBarrel"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// A banana.
type Banana struct {
	// The time at which you should reconsider eating this banana
	Expiration *time.Time `xml:"expiration"`
	Flavor     *string    `xml:"flavor"`
	Name       *string    `xml:"name"`
}

func (b Banana) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "banana"
	type alias Banana
	aux := &struct {
		*alias
		Expiration *timeRFC3339 `xml:"expiration"`
	}{
		alias:      (*alias)(&b),
		Expiration: (*timeRFC3339)(b.Expiration),
	}
	return e.EncodeElement(aux, start)
}

func (b *Banana) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias Banana
	aux := &struct {
		*alias
		Expiration *timeRFC3339 `xml:"expiration"`
	}{
		alias: (*alias)(b),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	b.Expiration = (*time.Time)(aux.Expiration)
	return nil
}

// BananaArrayResponse is the response envelope for operations that return a []Banana type.
type BananaArrayResponse struct {
	// Array of Banana
	Bananas *[]Banana `xml:"banana"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BananaResponse is the response envelope for operations that return a Banana type.
type BananaResponse struct {
	// A banana.
	Banana *Banana `xml:"banana"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// An Azure Storage blob
type Blob struct {
	Deleted *bool `xml:"Deleted"`

	// Dictionary of <string>
	Metadata *map[string]*string `xml:"String"`
	Name     *string             `xml:"Name"`

	// Properties of a blob
	Properties *BlobProperties `xml:"Properties"`
	Snapshot   *string         `xml:"Snapshot"`
}

type BlobPrefix struct {
	Name *string `xml:"Name"`
}

// Properties of a blob
type BlobProperties struct {
	AccessTier         *AccessTier    `xml:"AccessTier"`
	AccessTierInferred *bool          `xml:"AccessTierInferred"`
	ArchiveStatus      *ArchiveStatus `xml:"ArchiveStatus"`
	BlobSequenceNumber *int32         `xml:"x-ms-blob-sequence-number"`
	BlobType           *BlobType      `xml:"BlobType"`
	CacheControl       *string        `xml:"Cache-Control"`
	ContentDisposition *string        `xml:"Content-Disposition"`
	ContentEncoding    *string        `xml:"Content-Encoding"`
	ContentLanguage    *string        `xml:"Content-Language"`

	// Size in bytes
	ContentLength          *int64             `xml:"Content-Length"`
	ContentMd5             *string            `xml:"Content-MD5"`
	ContentType            *string            `xml:"Content-Type"`
	CopyCompletionTime     *time.Time         `xml:"CopyCompletionTime"`
	CopyID                 *string            `xml:"CopyId"`
	CopyProgress           *string            `xml:"CopyProgress"`
	CopySource             *string            `xml:"CopySource"`
	CopyStatus             *CopyStatusType    `xml:"CopyStatus"`
	CopyStatusDescription  *string            `xml:"CopyStatusDescription"`
	DeletedTime            *time.Time         `xml:"DeletedTime"`
	DestinationSnapshot    *string            `xml:"DestinationSnapshot"`
	Etag                   *string            `xml:"Etag"`
	IncrementalCopy        *bool              `xml:"IncrementalCopy"`
	LastModified           *time.Time         `xml:"Last-Modified"`
	LeaseDuration          *LeaseDurationType `xml:"LeaseDuration"`
	LeaseState             *LeaseStateType    `xml:"LeaseState"`
	LeaseStatus            *LeaseStatusType   `xml:"LeaseStatus"`
	RemainingRetentionDays *int32             `xml:"RemainingRetentionDays"`
	ServerEncrypted        *bool              `xml:"ServerEncrypted"`
}

func (b BlobProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias BlobProperties
	aux := &struct {
		*alias
		CopyCompletionTime *timeRFC1123 `xml:"CopyCompletionTime"`
		DeletedTime        *timeRFC1123 `xml:"DeletedTime"`
		LastModified       *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias:              (*alias)(&b),
		CopyCompletionTime: (*timeRFC1123)(b.CopyCompletionTime),
		DeletedTime:        (*timeRFC1123)(b.DeletedTime),
		LastModified:       (*timeRFC1123)(b.LastModified),
	}
	return e.EncodeElement(aux, start)
}

func (b *BlobProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias BlobProperties
	aux := &struct {
		*alias
		CopyCompletionTime *timeRFC1123 `xml:"CopyCompletionTime"`
		DeletedTime        *timeRFC1123 `xml:"DeletedTime"`
		LastModified       *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias: (*alias)(b),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	b.CopyCompletionTime = (*time.Time)(aux.CopyCompletionTime)
	b.DeletedTime = (*time.Time)(aux.DeletedTime)
	b.LastModified = (*time.Time)(aux.LastModified)
	return nil
}

type Blobs struct {
	Blob       *[]Blob       `xml:"Blob"`
	BlobPrefix *[]BlobPrefix `xml:"BlobPrefix"`
}

// I am a complex type with no XML node
type ComplexTypeNoMeta struct {
	// The id of the res
	ID *string `xml:"ID"`
}

// I am a complex type with XML node
type ComplexTypeWithMeta struct {
	// The id of the res
	ID *string `xml:"ID"`
}

// An Azure Storage container
type Container struct {
	// Dictionary of <string>
	Metadata *map[string]*string `xml:"String"`
	Name     *string             `xml:"Name"`

	// Properties of a container
	Properties *ContainerProperties `xml:"Properties"`
}

// Properties of a container
type ContainerProperties struct {
	Etag          *string            `xml:"Etag"`
	LastModified  *time.Time         `xml:"Last-Modified"`
	LeaseDuration *LeaseDurationType `xml:"LeaseDuration"`
	LeaseState    *LeaseStateType    `xml:"LeaseState"`
	LeaseStatus   *LeaseStatusType   `xml:"LeaseStatus"`
	PublicAccess  *PublicAccessType  `xml:"PublicAccess"`
}

func (c ContainerProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias ContainerProperties
	aux := &struct {
		*alias
		LastModified *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias:        (*alias)(&c),
		LastModified: (*timeRFC1123)(c.LastModified),
	}
	return e.EncodeElement(aux, start)
}

func (c *ContainerProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias ContainerProperties
	aux := &struct {
		*alias
		LastModified *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias: (*alias)(c),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	c.LastModified = (*time.Time)(aux.LastModified)
	return nil
}

// CORS is an HTTP feature that enables a web application running under one domain to access resources in another domain.
// Web browsers implement a security restriction known as same-origin policy that prevents a web page from calling APIs in
// a different domain; CORS provides a secure way to allow one domain (the origin domain) to call APIs in another domain
type CorsRule struct {
	// the request headers that the origin domain may specify on the CORS request.
	AllowedHeaders *string `xml:"AllowedHeaders"`

	// The methods (HTTP request verbs) that the origin domain may use for a CORS request. (comma separated)
	AllowedMethods *string `xml:"AllowedMethods"`

	// The origin domains that are permitted to make a request against the storage service via CORS. The origin domain is the
	// domain from which the request originates. Note that the origin must be an exact case-sensitive match with the origin that
	// the user age sends to the service. You can also use the wildcard character '*' to allow all origin domains to make requests
	// via CORS.
	AllowedOrigins *string `xml:"AllowedOrigins"`

	// The response headers that may be sent in the response to the CORS request and exposed by the browser to the request issuer
	ExposedHeaders *string `xml:"ExposedHeaders"`

	// The maximum amount time that a browser should cache the preflight OPTIONS request.
	MaxAgeInSeconds *int32 `xml:"MaxAgeInSeconds"`
}

type Error struct {
	Message *string `xml:"message"`
	Status  *int32  `xml:"status"`
}

func newError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

type JSONInput struct {
	ID *int32 `json:"id,omitempty"`
}

type JSONOutput struct {
	ID *int32 `json:"id,omitempty"`
}

// JSONOutputResponse is the response envelope for operations that return a JSONOutput type.
type JSONOutputResponse struct {
	JSONOutput *JSONOutput

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// An enumeration of blobs
type ListBlobsResponse struct {
	Blobs           *Blobs  `xml:"Blobs"`
	ContainerName   *string `xml:"ContainerName,attr"`
	Delimiter       *string `xml:"Delimiter"`
	Marker          *string `xml:"Marker"`
	MaxResults      *int32  `xml:"MaxResults"`
	NextMarker      *string `xml:"NextMarker"`
	Prefix          *string `xml:"Prefix"`
	ServiceEndpoint *string `xml:"ServiceEndpoint,attr"`
}

// ListBlobsResponseResponse is the response envelope for operations that return a ListBlobsResponse type.
type ListBlobsResponseResponse struct {
	// An enumeration of blobs
	EnumerationResults *ListBlobsResponse `xml:"EnumerationResults"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// An enumeration of containers
type ListContainersResponse struct {
	Containers      *[]Container `xml:"Containers>Container"`
	Marker          *string      `xml:"Marker"`
	MaxResults      *int32       `xml:"MaxResults"`
	NextMarker      *string      `xml:"NextMarker"`
	Prefix          *string      `xml:"Prefix"`
	ServiceEndpoint *string      `xml:"ServiceEndpoint,attr"`
}

// ListContainersResponseResponse is the response envelope for operations that return a ListContainersResponse type.
type ListContainersResponseResponse struct {
	// An enumeration of containers
	EnumerationResults *ListContainersResponse `xml:"EnumerationResults"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Azure Analytics Logging settings.
type Logging struct {
	// Indicates whether all delete requests should be logged.
	Delete *bool `xml:"Delete"`

	// Indicates whether all read requests should be logged.
	Read *bool `xml:"Read"`

	// the retention policy
	RetentionPolicy *RetentionPolicy `xml:"RetentionPolicy"`

	// The version of Storage Analytics to configure.
	Version *string `xml:"Version"`

	// Indicates whether all write requests should be logged.
	Write *bool `xml:"Write"`
}

type Metrics struct {
	// Indicates whether metrics are enabled for the Blob service.
	Enabled *bool `xml:"Enabled"`

	// Indicates whether metrics should generate summary statistics for called API operations.
	IncludeApIs *bool `xml:"IncludeAPIs"`

	// the retention policy
	RetentionPolicy *RetentionPolicy `xml:"RetentionPolicy"`

	// The version of Storage Analytics to configure.
	Version *string `xml:"Version"`
}

// the retention policy
type RetentionPolicy struct {
	// Indicates the number of days that metrics or logging or soft-deleted data should be retained. All data older than this
	// value will be deleted
	Days *int32 `xml:"Days"`

	// Indicates whether a retention policy is enabled for the storage service
	Enabled *bool `xml:"Enabled"`
}

// I am root, and I ref a model WITH meta
type RootWithRefAndMeta struct {
	// XML will use XMLComplexTypeWithMeta
	RefToModel *ComplexTypeWithMeta `xml:"XMLComplexTypeWithMeta"`

	// Something else (just to avoid flattening)
	Something *string `xml:"Something"`
}

// RootWithRefAndMetaResponse is the response envelope for operations that return a RootWithRefAndMeta type.
type RootWithRefAndMetaResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// I am root, and I ref a model WITH meta
	RootWithRefAndMeta *RootWithRefAndMeta `xml:"RootWithRefAndMeta"`
}

// I am root, and I ref a model with no meta
type RootWithRefAndNoMeta struct {
	// XML will use RefToModel
	RefToModel *ComplexTypeNoMeta `xml:"RefToModel"`

	// Something else (just to avoid flattening)
	Something *string `xml:"Something"`
}

// RootWithRefAndNoMetaResponse is the response envelope for operations that return a RootWithRefAndNoMeta type.
type RootWithRefAndNoMetaResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// I am root, and I ref a model with no meta
	RootWithRefAndNoMeta *RootWithRefAndNoMeta `xml:"RootWithRefAndNoMeta"`
}

// signed identifier
type SignedIDentifier struct {
	// The access policy
	AccessPolicy *AccessPolicy `xml:"AccessPolicy"`

	// a unique id
	ID *string `xml:"Id"`
}

// SignedIDentifierArrayResponse is the response envelope for operations that return a []SignedIDentifier type.
type SignedIDentifierArrayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// a collection of signed identifiers
	SignedIdentifiers *[]SignedIDentifier `xml:"SignedIdentifier"`
}

// A slide in a slideshow
type Slide struct {
	Items *[]string `xml:"item"`
	Title *string   `xml:"title"`
	Type  *string   `xml:"type,attr"`
}

// Data about a slideshow
type Slideshow struct {
	Author *string  `xml:"author,attr"`
	Date   *string  `xml:"date,attr"`
	Slides *[]Slide `xml:"slide"`
	Title  *string  `xml:"title,attr"`
}

func (s Slideshow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "slideshow"
	type alias Slideshow
	aux := &struct {
		*alias
	}{
		alias: (*alias)(&s),
	}
	return e.EncodeElement(aux, start)
}

// SlideshowResponse is the response envelope for operations that return a Slideshow type.
type SlideshowResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Data about a slideshow
	Slideshow *Slideshow `xml:"slideshow"`
}

// Storage Service Properties.
type StorageServiceProperties struct {
	// The set of CORS rules.
	Cors *[]CorsRule `xml:"Cors>CorsRule"`

	// The default version to use for requests to the Blob service if an incoming request's version is not specified. Possible
	// values include version 2008-10-27 and all more recent versions
	DefaultServiceVersion *string `xml:"DefaultServiceVersion"`

	// The Delete Retention Policy for the service
	DeleteRetentionPolicy *RetentionPolicy `xml:"DeleteRetentionPolicy"`

	// A summary of request statistics grouped by API in hourly aggregates for blobs
	HourMetrics *Metrics `xml:"HourMetrics"`

	// Azure Analytics Logging settings
	Logging *Logging `xml:"Logging"`

	// a summary of request statistics grouped by API in minute aggregates for blobs
	MinuteMetrics *Metrics `xml:"MinuteMetrics"`
}

// StorageServicePropertiesResponse is the response envelope for operations that return a StorageServiceProperties type.
type StorageServicePropertiesResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Storage Service Properties.
	StorageServiceProperties *StorageServiceProperties `xml:"StorageServiceProperties"`
}

// XMLGetHeadersResponse contains the response from method XML.GetHeaders.
type XMLGetHeadersResponse struct {
	// CustomHeader contains the information returned from the Custom-Header header response.
	CustomHeader *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
