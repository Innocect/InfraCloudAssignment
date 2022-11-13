package server

import (
	"fmt"
	"net/http"
	"net/url"
	api "shortlink/internal/api/router"
	"shortlink/internal/config"

	"github.com/golang/glog"
	"github.com/rs/cors"
)

const (
	dependencyURL = "https://github.com/innocect/software/shortlink/"
)

// Server ...
type Server struct {
	Router *api.Router
}

func RunServer() (err error) {

	baseURL, err := url.Parse(dependencyURL)
	if err != nil {
		glog.Error("ServerError" + err.Error())
		return err
	}

	routerConfig := config.RouterConfig{
		DependencyURL: *baseURL,
	}

	// InitServices

	// InitRoutes
	server := &Server{
		Router: api.NewRouter(),
	}
	server.Router.InitializeRouter(&routerConfig)

	return startServer(server)
}

func startServer(server *Server) error {
	fmt.Print("Starting Https Server on Port :50051")

	corsOpt := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodPost, http.MethodGet},
	})

	return http.ListenAndServe("127.0.0.1:8000", corsOpt.Handler(server.Router))
}
