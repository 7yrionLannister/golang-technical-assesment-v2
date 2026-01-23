package main

import (
	builtinlog "log"
	"net"
	"net/http"
	"os"

	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/api"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/db"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/env"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"

	middleware "github.com/oapi-codegen/nethttp-middleware"
)

// "context"
// "os"
// "os/signal"
// "syscall"

// messagingRouter "github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/messaging/router"
// "github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/middleware"
// webRouter "github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/router"
// "github.com/7yrionLannister/golang-technical-assesment-v2/pkg/config"
// "github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
// "github.com/7yrionLannister/golang-technical-assesment-v2/pkg/util"

func main() {
	// Load environment variables
	err := env.LoadEnv()
	if err != nil {
		builtinlog.Fatalf("Error loading environment variables: %v", err) // use log as our logger is not initialized yet
	}
	// Set the structured logger
	log.L.InitLogger(env.Env.LogLevel)

	// Migrate the database
	err = db.MigrateUp()
	if err != nil {
		log.L.Error("Error migrating database", "error", err)
		os.Exit(1)
	}
	log.L.Debug("Initialized application")

	swagger, err := api.GetSwagger()
	if err != nil {
		log.L.Error("Failed to load swagger spec")
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	consumptionApi := api.NewEnergyConsumptionController()
	strictServer := api.NewStrictHandler(consumptionApi, nil)

	r := http.NewServeMux()

	// register our strictServer above as the handler for the interface
	api.HandlerFromMux(strictServer, r)

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	h := middleware.OapiRequestValidator(swagger)(r)

	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort("0.0.0.0", env.Env.ServerPort),
	}

	builtinlog.Fatal(s.ListenAndServe())
}
