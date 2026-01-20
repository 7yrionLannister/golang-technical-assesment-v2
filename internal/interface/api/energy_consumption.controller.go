package api

import "net/http"

// TODO: mock for unit tests
type EnergyConsumptionControllerInterface interface {
	ServerInterface
}

type EnergyConsumptionController struct{}

func (EnergyConsumptionController) GetConsumption(w http.ResponseWriter, r *http.Request, params GetConsumptionParams)
