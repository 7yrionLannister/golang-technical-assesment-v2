package api

import (
	"context"
	"encoding/json"

	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/application/service"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/util"
)

// TODO: mock for unit tests
type EnergyConsumptionControllerInterface interface {
	StrictServerInterface
}

type EnergyConsumptionController struct {
	svc service.EnergyConsumptionServiceInterface
}

func NewEnergyConsumptionController(svc service.EnergyConsumptionServiceInterface) *EnergyConsumptionController {
	return &EnergyConsumptionController{
		svc: svc,
	}
}

func (ctr *EnergyConsumptionController) GetConsumption(ctx context.Context, request GetConsumptionRequestObject) (GetConsumptionResponseObject, error) {
	result, err := ctr.svc.GetEnergyConsumptions(request.Params.MeterId, request.Params.StartDate.Time, request.Params.EndDate.Time, string(request.Params.Period))
	if err != nil {
		return nil, util.HandleError(err, "Failed to query energy consumption")
	}
	return GetConsumption200JSONResponse(*result), nil
}

func (*EnergyConsumptionController) GetOpenapi(ctx context.Context, request GetOpenapiRequestObject) (GetOpenapiResponseObject, error) {
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
