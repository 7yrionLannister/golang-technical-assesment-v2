package main

import (
	builtinlog "log"
	"net"
	"net/http"
	"os"

	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/application/repository"
	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/api"
	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/middleware"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/db"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/env"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
)

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

	// Start DB
	db.DB = new(db.GormDatabase)
	err = db.DB.InitDatabaseConnection()
	if err != nil {
		log.L.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	err = db.ImportTestData()
	if err != nil {
		log.L.Warn("Error importing test data", "error", err)
	}

	swagger, err := api.GetSwagger()
	if err != nil {
		log.L.Error("Failed to load swagger spec")
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil
	consumptionRepo := &repository.EnergyConsumptionRepository{} // TODO: use constructor
	consumptionApi := api.NewEnergyConsumptionService(consumptionRepo)
	strictServer := api.NewStrictHandler(consumptionApi, nil)

	r := http.NewServeMux()

	validator := middleware.ValidatorMiddleware(swagger)

	// register our strictServer above as the handler for the interface
	api.HandlerFromMux(strictServer, r)

	s := &http.Server{
		Handler: validator(r),
		Addr:    net.JoinHostPort("0.0.0.0", env.Env.ServerPort),
	}

	builtinlog.Fatal(s.ListenAndServe())
}
