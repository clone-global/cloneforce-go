// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce

import (
	"github.com/stainless-sdks/cloneforce-go/option"
)

// V1IntegrationService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1IntegrationService] method instead.
type V1IntegrationService struct {
	options []option.RequestOption
	// Clone integration management (Slack, Email, MS Teams, Phone)
	Phone V1IntegrationPhoneService
}

// NewV1IntegrationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1IntegrationService(opts ...option.RequestOption) (r V1IntegrationService) {
	r = V1IntegrationService{}
	r.options = opts
	r.Phone = NewV1IntegrationPhoneService(opts...)
	return
}
