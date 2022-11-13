package api

import (
	"net/http"
	"shortlink/internal/models"
	"shortlink/internal/services"
	"shortlink/internal/utils"
)

// ShortenURLHandler  Shortens the given long form of the URL
// service services.ShortenURLService
func ShortenURLHandler(service services.ShortenURLService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req := &models.ShortLinkRedisData{}

		if marshErr := utils.UnmarshalRequest(ctx, r, req); marshErr != nil {
			w.WriteHeader(500)
			w.Write([]byte(marshErr.Error()))
			return
		}

		err := service.ValidateRequest(ctx, req)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		shortURL, err := service.ProcessRequest(ctx, req)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(err.Error()))
			return
		}

		response, err := utils.MarshalRequest(ctx, shortURL)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(200)
		w.Write(response)
	}
}
