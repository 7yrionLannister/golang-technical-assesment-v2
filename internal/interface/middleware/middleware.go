package middleware

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	httpmiddleware "github.com/oapi-codegen/nethttp-middleware"
)

// TODO
// add logging middleware that logs requests and responses

func ValidatorMiddleware(swagger *openapi3.T) func(next http.Handler) http.Handler {
	return httpmiddleware.OapiRequestValidatorWithOptions(swagger, &httpmiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		},
	})
}
