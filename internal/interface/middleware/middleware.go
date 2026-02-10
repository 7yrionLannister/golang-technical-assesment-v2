package middleware

import (
	"net/http"
	"time"

	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	httpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func ValidatorMiddleware(swagger *openapi3.T) func(next http.Handler) http.Handler {
	return httpmiddleware.OapiRequestValidatorWithOptions(swagger, &httpmiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		},
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Process the request
		next.ServeHTTP(w, r)

		// Log after the request has been processed
		log.L.Debug("Request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)
	})
}
