package server

import (
	"fmt"
	"net/http"
	"os"
	api "shortlink/internal/api/router"
	"shortlink/internal/config"
	"shortlink/internal/dao"
	"shortlink/internal/services"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// Server ...
type Server struct {
	Router *api.Router
}

func RunServer() (err error) {

	err = godotenv.Load("../../.env")
	if err != nil {
		return err
	}

	routerConfig := &config.RouterConfig{
		RoutePrefix: "127.0.0.1:8000",
		Port:        os.Getenv("PORT"),
	}

	//InitRedisDao
	redisConfig := &config.RedisConfig{
		Redis: dao.InitShortlinkRedisDAO(),
	}

	// InitServices
	services.InitShortenURLService(redisConfig)

	// InitRoutes
	server := &Server{
		Router: api.NewRouter(),
	}
	server.Router.InitializeRouter(routerConfig)

	return startServer(server, routerConfig)
}

func startServer(server *Server, config *config.RouterConfig) error {
	fmt.Print("Starting Https Server on Port :50051")

	corsOpt := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodPost, http.MethodGet},
	})

	return http.ListenAndServe(config.RoutePrefix, corsOpt.Handler(server.Router))
}
