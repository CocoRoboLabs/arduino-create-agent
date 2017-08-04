// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "arduino-create-agent": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/arduino/arduino-create-agent/design
// --out=$(GOPATH)/src/github.com/arduino/arduino-create-agent
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// CommandsV1Controller is the controller interface for the CommandsV1 actions.
type CommandsV1Controller interface {
	goa.Muxer
	Exec(*ExecCommandsV1Context) error
	List(*ListCommandsV1Context) error
}

// MountCommandsV1Controller "mounts" a CommandsV1 resource controller on the given service.
func MountCommandsV1Controller(service *goa.Service, ctrl CommandsV1Controller) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/:id", ctrl.MuxHandler("preflight", handleCommandsV1Origin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/", ctrl.MuxHandler("preflight", handleCommandsV1Origin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewExecCommandsV1Context(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(ExecCommandsV1Payload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Exec(rctx)
	}
	h = handleCommandsV1Origin(h)
	service.Mux.Handle("POST", "/:id", ctrl.MuxHandler("exec", h, unmarshalExecCommandsV1Payload))
	service.LogInfo("mount", "ctrl", "CommandsV1", "action", "Exec", "route", "POST /:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListCommandsV1Context(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleCommandsV1Origin(h)
	service.Mux.Handle("GET", "/", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "CommandsV1", "action", "List", "route", "GET /")
}

// handleCommandsV1Origin applies the CORS response headers corresponding to the origin.
func handleCommandsV1Origin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalExecCommandsV1Payload unmarshals the request body into the context request data Payload field.
func unmarshalExecCommandsV1Payload(ctx context.Context, service *goa.Service, req *http.Request) error {
	var payload ExecCommandsV1Payload
	if err := service.DecodeRequest(req, &payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload
	return nil
}

// ConnectV1Controller is the controller interface for the ConnectV1 actions.
type ConnectV1Controller interface {
	goa.Muxer
	Websocket(*WebsocketConnectV1Context) error
}

// MountConnectV1Controller "mounts" a ConnectV1 resource controller on the given service.
func MountConnectV1Controller(service *goa.Service, ctrl ConnectV1Controller) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/connect", ctrl.MuxHandler("preflight", handleConnectV1Origin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewWebsocketConnectV1Context(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Websocket(rctx)
	}
	h = handleConnectV1Origin(h)
	service.Mux.Handle("GET", "/v1/connect", ctrl.MuxHandler("websocket", h, nil))
	service.LogInfo("mount", "ctrl", "ConnectV1", "action", "Websocket", "route", "GET /v1/connect")
}

// handleConnectV1Origin applies the CORS response headers corresponding to the origin.
func handleConnectV1Origin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// DiscoverV1Controller is the controller interface for the DiscoverV1 actions.
type DiscoverV1Controller interface {
	goa.Muxer
	List(*ListDiscoverV1Context) error
}

// MountDiscoverV1Controller "mounts" a DiscoverV1 resource controller on the given service.
func MountDiscoverV1Controller(service *goa.Service, ctrl DiscoverV1Controller) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v1/discover", ctrl.MuxHandler("preflight", handleDiscoverV1Origin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListDiscoverV1Context(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleDiscoverV1Origin(h)
	service.Mux.Handle("GET", "/v1/discover", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "DiscoverV1", "action", "List", "route", "GET /v1/discover")
}

// handleDiscoverV1Origin applies the CORS response headers corresponding to the origin.
func handleDiscoverV1Origin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// PublicController is the controller interface for the Public actions.
type PublicController interface {
	goa.Muxer
	goa.FileServer
}

// MountPublicController "mounts" a Public resource controller on the given service.
func MountPublicController(service *goa.Service, ctrl PublicController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/debug", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/docs", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "swagger/swagger.json", "route", "GET /swagger.json")

	h = ctrl.FileHandler("/debug", "templates/debug.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/debug", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "templates/debug.html", "route", "GET /debug")

	h = ctrl.FileHandler("/docs", "templates/docs.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/docs", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "templates/docs.html", "route", "GET /docs")
}

// handlePublicOrigin applies the CORS response headers corresponding to the origin.
func handlePublicOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}