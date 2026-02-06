package repository

import (
	"time"

	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/domain/model"
	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/domain/view"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/db"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/util"
)

//go:generate go tool counterfeiter . EnergyConsumptionRepositoryInterface
type EnergyConsumptionRepositoryInterface interface {
	GetEnergyConsumptionsByMeterIdBetweenDates(metersIds []uint8, startDate time.Time, endDate time.Time) ([]view.EnergyConsumption, error)
}

type EnergyConsumptionRepository struct {
	// TODO: put db here instead of accessing it via package
}

func (*EnergyConsumptionRepository) GetEnergyConsumptionsByMeterIdBetweenDates(metersIds []uint8, startDate time.Time, endDate time.Time) ([]view.EnergyConsumption, error) {
	var result []view.EnergyConsumption
	err := db.DB.
		Model(&model.EnergyConsumption{}).
		Select("device_id as meter_id, sum(consumption) as total_consumption").
		Where("device_id IN ? AND created_at BETWEEN ? AND ?", metersIds, startDate, endDate).
		Group("device_id").
		Scan(&result).
		Error()
	if err != nil {
		return nil, util.HandleError(err, "Failed to query energy consumptions")
	}
	log.L.Debug("query result", "dto", result)
	return result, nil
}
