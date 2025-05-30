// Code generated by goa v3.20.0, DO NOT EDIT.
//
// user client
//
// Command:
// $ goa gen be/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "user" service client.
type Client struct {
	CreateEndpoint goa.Endpoint
	GetEndpoint    goa.Endpoint
	ListEndpoint   goa.Endpoint
	UpdateEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
}

// NewClient initializes a "user" service client given the endpoints.
func NewClient(create, get, list, update, delete_ goa.Endpoint) *Client {
	return &Client{
		CreateEndpoint: create,
		GetEndpoint:    get,
		ListEndpoint:   list,
		UpdateEndpoint: update,
		DeleteEndpoint: delete_,
	}
}

// Create calls the "create" endpoint of the "user" service.
// Create may return the following errors:
//   - "unauthorized" (type *Unauthorized): Auth Failed
//   - "internalServerError" (type *InternalServerError): Internal Server Error
//   - "notFound" (type *NotFound): Not Found
//   - "badRequest" (type *BadRequest): Invalid Request
//   - "forbidden" (type *Forbidden): Accesso negato
//   - error: internal error
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res *User, err error) {
	var ires any
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*User), nil
}

// Get calls the "get" endpoint of the "user" service.
// Get may return the following errors:
//   - "unauthorized" (type *Unauthorized): Auth Failed
//   - "internalServerError" (type *InternalServerError): Internal Server Error
//   - "notFound" (type *NotFound): Not Found
//   - "badRequest" (type *BadRequest): Invalid Request
//   - "forbidden" (type *Forbidden): Accesso negato
//   - error: internal error
func (c *Client) Get(ctx context.Context, p *GetPayload) (res *User, err error) {
	var ires any
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*User), nil
}

// List calls the "list" endpoint of the "user" service.
// List may return the following errors:
//   - "unauthorized" (type *Unauthorized): Auth Failed
//   - "internalServerError" (type *InternalServerError): Internal Server Error
//   - "notFound" (type *NotFound): Not Found
//   - "badRequest" (type *BadRequest): Invalid Request
//   - "forbidden" (type *Forbidden): Accesso negato
//   - error: internal error
func (c *Client) List(ctx context.Context, p *ListPayload) (res []*User, err error) {
	var ires any
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*User), nil
}

// Update calls the "update" endpoint of the "user" service.
// Update may return the following errors:
//   - "unauthorized" (type *Unauthorized): Auth Failed
//   - "internalServerError" (type *InternalServerError): Internal Server Error
//   - "notFound" (type *NotFound): Not Found
//   - "badRequest" (type *BadRequest): Invalid Request
//   - "forbidden" (type *Forbidden): Accesso negato
//   - error: internal error
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (res *User, err error) {
	var ires any
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*User), nil
}

// Delete calls the "delete" endpoint of the "user" service.
// Delete may return the following errors:
//   - "unauthorized" (type *Unauthorized): Auth Failed
//   - "internalServerError" (type *InternalServerError): Internal Server Error
//   - "notFound" (type *NotFound): Not Found
//   - "badRequest" (type *BadRequest): Invalid Request
//   - "forbidden" (type *Forbidden): Accesso negato
//   - error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}
