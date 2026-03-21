// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce

import (
	"github.com/clone-global/cloneforce-go/option"
)

// V1Service contains methods and other services that help with interacting with
// the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	options []option.RequestOption
	// Clone profile management and asset generation
	Clones V1CloneService
	// Skill marketplace search and clone skill management
	Skills       V1SkillService
	Integrations V1IntegrationService
	// Organization-level connection credential management
	Connections V1ConnectionService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.options = opts
	r.Clones = NewV1CloneService(opts...)
	r.Skills = NewV1SkillService(opts...)
	r.Integrations = NewV1IntegrationService(opts...)
	r.Connections = NewV1ConnectionService(opts...)
	return
}
