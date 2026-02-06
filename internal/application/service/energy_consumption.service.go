package service

import (
	"time"

	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/application/repository"
	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/domain/dto"
	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/interface/api"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/util"
	"github.com/go-faker/faker/v4"
)

// TODO: mock for unit tests
type EnergyConsumptionServiceInterface interface {
	GetEnergyConsumptions(metersIds []uint8, startDate time.Time, endDate time.Time, kindPeriod string) (*dto.PeriodicConsumption, error)
}

type EnergyConsumptionService struct {
	repo repository.EnergyConsumptionRepositoryInterface
}

func (svc *EnergyConsumptionService) GetEnergyConsumptions(metersIds []uint8, startDate time.Time, endDate time.Time, kindPeriod api.GetConsumptionParamsPeriod) (*dto.PeriodicConsumption, error) {
	var periodDto = &dto.PeriodicConsumption{
		Period:    make([]string, 0),
		DataGraph: make([]dto.EnergyConsumption, 0),
	}
	err := svc.stepThroughPeriod(periodDto, metersIds, startDate, endDate, kindPeriod)
	if err != nil {
		return nil, err
	}
	return periodDto, nil
}

// Iterates through the period between startDate and endDate, incrementing the date by the kindPeriod to find the sub-periods.
// For each sub-period, it computes the energy consumption for each meter in the metersIds slice.
func (svc *EnergyConsumptionService) stepThroughPeriod(periodDto *dto.PeriodicConsumption, metersIds []uint8, startDate time.Time, endDate time.Time, kindPeriod api.GetConsumptionParamsPeriod) error {
	// map to keep track of the *dto.EnergyConsumptionDTO that is part of the DataGraph
	energyConsumptionDTOForMeter := make(map[uint8]*dto.EnergyConsumption)
	for startDate.Before(endDate) {
		periodEndDate, periodString := stepDateAndGetPeriodString(kindPeriod, startDate)
		periodDto.Period = append(periodDto.Period, periodString)
		if periodEndDate.After(endDate) {
			periodEndDate = endDate
		}
		err := svc.populateDataGraphForPeriod(periodDto, metersIds, energyConsumptionDTOForMeter, startDate, periodEndDate)
		if err != nil {
			return err
		}
		startDate = periodEndDate
	}
	for _, v := range energyConsumptionDTOForMeter {
		periodDto.DataGraph = append(periodDto.DataGraph, *v)
	}
	return nil
}

// Increments the date by the kindPeriod.
// Gets the period string for the kindPeriod.
func stepDateAndGetPeriodString(kindPeriod api.GetConsumptionParamsPeriod, initialDate time.Time) (newDate time.Time, periodString string) {
	switch kindPeriod {
	case api.Daily:
		return initialDate.AddDate(0, 0, 1), initialDate.Format("January 2") // TODO format as "JAN 2"
	case api.Weekly:
		periodString = initialDate.Format("January 2") + " - " + initialDate.AddDate(0, 0, 6).Format("January 2") // TODO format as "JAN 2 - JAN 2"
		return initialDate.AddDate(0, 0, 7), periodString
	case api.Monthly:
		return initialDate.AddDate(0, 1, 0), initialDate.Format("January 2006") // TODO format as "JAN 2006"
	default:
		return initialDate, "TODO"
	}
}

// Queries the energy consumption for each meter in the metersIds slice for the period between periodStartDate and periodEndDate
func (svc *EnergyConsumptionService) populateDataGraphForPeriod(periodDto *dto.PeriodicConsumption, metersIds []uint8, energyConsumptionDTOForMeter map[uint8]*dto.EnergyConsumption, periodStartDate time.Time, periodEndDate time.Time) error {
	energyConsumptions, err := svc.repo.GetEnergyConsumptionsByMeterIdBetweenDates(metersIds, periodStartDate, periodEndDate)
	if err != nil {
		return util.HandleError(err, "Failed to fetch energy consumptions")
	}
	for index, meterId := range metersIds {
		energyConsumptionDTO, present := energyConsumptionDTOForMeter[meterId]
		if !present {
			energyConsumptionDTO = &dto.EnergyConsumption{
				MeterId: meterId,
				Active:  make([]float32, 0),
				Address: faker.GetRealAddress().Address, // Asume faker to be an http client that gets the address for the meter
			}
			energyConsumptionDTOForMeter[meterId] = energyConsumptionDTO
		}
		energyConsumptionDTO.Active = append(energyConsumptionDTO.Active, energyConsumptions[index].TotalConsumption)
	}
	return nil
}
