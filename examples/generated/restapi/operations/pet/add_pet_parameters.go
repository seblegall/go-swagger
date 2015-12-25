package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/go-swagger/go-swagger/examples/generated/models"
)

// NewAddPetParams creates a new AddPetParams object
// with the default values initialized.
func NewAddPetParams() AddPetParams {
	return AddPetParams{}
}

// AddPetParams contains all the bound params for the add pet operation
// typically these are obtained from a http.Request
//
// swagger:parameters addPet
type AddPetParams struct {
	/*Pet object that needs to be added to the store
	  Required: true
	  In: body
	*/
	Body *models.Pet
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddPetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	var body models.Pet
	if err := route.Consumer.Consume(r.Body, &body); err != nil {
		res = append(res, errors.NewParseError("body", "body", "", err))
	} else {
		if err := body.Validate(route.Formats); err != nil {
			res = append(res, err)
		}

		if len(res) == 0 {
			o.Body = &body
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
