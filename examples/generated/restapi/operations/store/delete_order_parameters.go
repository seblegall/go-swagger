package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteOrderParams creates a new DeleteOrderParams object
// with the default values initialized.
func NewDeleteOrderParams() DeleteOrderParams {
	return DeleteOrderParams{}
}

// DeleteOrderParams contains all the bound params for the delete order operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteOrder
type DeleteOrderParams struct {
	/*ID of the order that needs to be deleted
	  Required: true
	  Min Length: 1
	  In: path
	*/
	OrderID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DeleteOrderParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	rOrderID, rhkOrderID, _ := route.Params.GetOK("orderId")
	if err := o.bindOrderID(rOrderID, rhkOrderID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteOrderParams) bindOrderID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.OrderID = raw

	if err := o.validateOrderID(formats); err != nil {
		return err
	}

	return nil
}

func (o *DeleteOrderParams) validateOrderID(formats strfmt.Registry) error {

	if err := validate.MinLength("orderId", "path", string(o.OrderID), 1); err != nil {
		return err
	}

	return nil
}
