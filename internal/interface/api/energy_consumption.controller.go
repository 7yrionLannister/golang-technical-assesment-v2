package api

import (
	"context"
	"encoding/json"

	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
)

// TODO: mock for unit tests
type EnergyConsumptionControllerInterface interface {
	StrictServerInterface
}

type EnergyConsumptionController struct{}

func NewEnergyConsumptionController() *EnergyConsumptionController {
	return &EnergyConsumptionController{}
}

func (EnergyConsumptionController) GetConsumption(ctx context.Context, request GetConsumptionRequestObject) (GetConsumptionResponseObject, error) {
	// TODO: implement
	return nil, nil
}

func (EnergyConsumptionController) GetOpenapi(ctx context.Context, request GetOpenapiRequestObject) (GetOpenapiResponseObject, error) {
	swagger, err := GetSwagger()
	if err != nil {
		log.L.Error("Failed to load swagger spec")
	}
	swaggerJSON, err := json.Marshal(swagger)
	if err != nil {
		log.L.Error("Failed to turn spec into json")
	}
	v := &map[string]any{}
	err = json.Unmarshal(swaggerJSON, v)
	if err != nil {
		log.L.Error("Failed to unmarshall swagger")
	}
	return GetOpenapi200JSONResponse(*v), nil
}
