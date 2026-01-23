package api

import "context"

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
