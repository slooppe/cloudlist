package schema

import (
	"context"
	"fmt"
)

// Provider is an interface implemented by any cloud service provider.
//
// It provides the bare minimum of methods to allow complete overview of user
// data.
type Provider interface {
	// Name returns the name of the provider
	Name() string
	// ProfileName returns the name of the provider profile
	ProfileName() string
	// Resources returns the provider for an resource deployment source.
	Resources(ctx context.Context) (*Resources, error)
}

// Resources is a container of multiple resource returned from providers
type Resources struct {
	Items []*Resource
}

// Append appends a single resource to the resource list
func (r *Resources) Append(resource *Resource) {
	r.Items = append(r.Items, resource)
}

// Merge merges a list of resources into the main list
func (r *Resources) Merge(resources *Resources) {
	r.Items = append(r.Items, resources.Items...)
}

// Resource is a cloud resource belonging to the organization
type Resource struct {
	// Public specifies whether the asset is public facing or private
	Public bool `json:"public"`
	// Provider is the name of provider for instance
	Provider string `json:"provider"`
	// Profile is the profile name of the resource provider
	Profile string `json:"profile,omitempty"`
	// ProfileName is the name of the key profile
	ProfileName string `json:"profile_name,omitempty"`
	// PublicIPv4 is the public ipv4 address of the instance.
	PublicIPv4 string `json:"public_ipv4,omitempty"`
	// PrivateIpv4 is the private ipv4 address of the instance
	PrivateIpv4 string `json:"private_ipv4,omitempty"`
	// DNSName is the DNS name of the resource
	DNSName string `json:"dns_name,omitempty"`
}

// ErrNoSuchKey means no such key exists in metadata.
type ErrNoSuchKey struct {
	Name string
}

// Error returns the value of the metadata key
func (e *ErrNoSuchKey) Error() string {
	return fmt.Sprintf("no such key: %s", e.Name)
}

// Options contains configuration options for a provider
type Options []OptionBlock

// OptionBlock is a single option on which operation is possible
type OptionBlock map[string]string

// GetMetadata returns the value for a key if it exists.
func (o OptionBlock) GetMetadata(key string) (string, bool) {
	data, ok := o[key]
	if !ok || data == "" {
		return "", false
	}
	return data, true
}
