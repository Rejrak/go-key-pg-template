// Code generated by goa v3.20.0, DO NOT EDIT.
//
// user endpoints
//
// Command:
// $ goa gen be/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "user" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
	Get    goa.Endpoint
	List   goa.Endpoint
	Update goa.Endpoint
	Delete goa.Endpoint
}

// NewEndpoints wraps the methods of the "user" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Create: NewCreateEndpoint(s, a.OAuth2Auth),
		Get:    NewGetEndpoint(s, a.OAuth2Auth),
		List:   NewListEndpoint(s, a.OAuth2Auth),
		Update: NewUpdateEndpoint(s, a.OAuth2Auth),
		Delete: NewDeleteEndpoint(s, a.OAuth2Auth),
	}
}

// Use applies the given middleware to all the "user" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.Get = m(e.Get)
	e.List = m(e.List)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "user".
func NewCreateEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "oauth2",
			Scopes:         []string{"openid"},
			RequiredScopes: []string{"openid"},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "password",
					TokenURL:   "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
					RefreshURL: "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
				},
			},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Create(ctx, p)
	}
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "user".
func NewGetEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "oauth2",
			Scopes:         []string{"openid"},
			RequiredScopes: []string{"openid"},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "password",
					TokenURL:   "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
					RefreshURL: "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
				},
			},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Get(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "user".
func NewListEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListPayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "oauth2",
			Scopes:         []string{"openid"},
			RequiredScopes: []string{"openid"},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "password",
					TokenURL:   "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
					RefreshURL: "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
				},
			},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.List(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "user".
func NewUpdateEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdatePayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "oauth2",
			Scopes:         []string{"openid"},
			RequiredScopes: []string{"openid"},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "password",
					TokenURL:   "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
					RefreshURL: "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
				},
			},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Update(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "user".
func NewDeleteEndpoint(s Service, authOAuth2Fn security.AuthOAuth2Func) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeletePayload)
		var err error
		sc := security.OAuth2Scheme{
			Name:           "oauth2",
			Scopes:         []string{"openid"},
			RequiredScopes: []string{"openid"},
			Flows: []*security.OAuthFlow{
				&security.OAuthFlow{
					Type:       "password",
					TokenURL:   "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
					RefreshURL: "http://localhost:8080/realms/LastingDynamics/protocol/openid-connect/token",
				},
			},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authOAuth2Fn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Delete(ctx, p)
	}
}
