// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetUserUserIDHandlerFunc turns a function with the right signature into a get user user ID handler
type GetUserUserIDHandlerFunc func(GetUserUserIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserUserIDHandlerFunc) Handle(params GetUserUserIDParams) middleware.Responder {
	return fn(params)
}

// GetUserUserIDHandler interface for that can handle valid get user user ID params
type GetUserUserIDHandler interface {
	Handle(GetUserUserIDParams) middleware.Responder
}

// NewGetUserUserID creates a new http.Handler for the get user user ID operation
func NewGetUserUserID(ctx *middleware.Context, handler GetUserUserIDHandler) *GetUserUserID {
	return &GetUserUserID{Context: ctx, Handler: handler}
}

/*GetUserUserID swagger:route GET /user/{userId} getUserUserId

get a user

return the specified

*/
type GetUserUserID struct {
	Context *middleware.Context
	Handler GetUserUserIDHandler
}

func (o *GetUserUserID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUserUserIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
