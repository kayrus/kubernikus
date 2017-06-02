package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetAPIV1ClustersParams creates a new GetAPIV1ClustersParams object
// with the default values initialized.
func NewGetAPIV1ClustersParams() GetAPIV1ClustersParams {
	var ()
	return GetAPIV1ClustersParams{}
}

// GetAPIV1ClustersParams contains all the bound params for the get API v1 clusters operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAPIV1Clusters
type GetAPIV1ClustersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetAPIV1ClustersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}