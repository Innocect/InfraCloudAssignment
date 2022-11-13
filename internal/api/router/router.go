package api

import (
	"net/http"
	api "shortlink/internal/api/handler"
	"shortlink/internal/config"
	"shortlink/internal/services"

	"github.com/gorilla/mux"
)

const (
	protobufContentType = "application/x-protobuf"
)

// Router ...
type Router struct {
	*mux.Router
}

// NewRouter ...
func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

// InitializeRouter ...
func (r *Router) InitializeRouter(routerConfig *config.RouterConfig) {

	r.initializeRoutes(routerConfig)
}

func (r *Router) initializeRoutes(routerConfig *config.RouterConfig) {

	r.HandleFunc("/shorten-url", api.ShortenURLHandler(services.GetShortenURLService())).
		Methods(http.MethodPost).
		Name("ShortenURL")
}
