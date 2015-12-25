package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/examples/generated/models"
)

// NewUpdateUserParams creates a new UpdateUserParams object
// with the default values initialized.
func NewUpdateUserParams() UpdateUserParams {
	return UpdateUserParams{}
}

// UpdateUserParams contains all the bound params for the update user operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateUser
type UpdateUserParams struct {
	/*Updated user object
	  Required: true
	  In: body
	*/
	Body *models.User
	/*name that need to be deleted
	  Required: true
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *UpdateUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	var body models.User
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

	rUsername, rhkUsername, _ := route.Params.GetOK("username")
	if err := o.bindUsername(rUsername, rhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateUserParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Username = raw

	return nil
}
