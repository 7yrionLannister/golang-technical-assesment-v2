package api

import (
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/application/repository/repositoryfakes"
	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/domain/dto"
	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/domain/view"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/env"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
	"github.com/go-faker/faker/v4"
	"github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/assert"
)

var (
	mockRepo   *repositoryfakes.FakeEnergyConsumptionRepositoryInterface
	controller EnergyConsumptionServiceInterface
)

var expectedResult = []dto.EnergyConsumption{
	{MeterId: 0, Active: []float32{1, 0}},
	{MeterId: 1, Active: []float32{0, 11}},
}

func TestMain(m *testing.M) {
	log.L.InitLogger(env.Env.LogLevel)
	mockRepo = new(repositoryfakes.FakeEnergyConsumptionRepositoryInterface)
	controller = NewEnergyConsumptionService(mockRepo)

	code := m.Run()
	os.Exit(code)
}

func mockForPeriods(metersIds []uint8) {
	// Expect
	data := make([]view.EnergyConsumption, 2)
	for i := range 2 {
		data[i] = view.EnergyConsumption{
			MeterId:          metersIds[i],
			Address:          faker.GetRealAddress().Address,
			TotalConsumption: float32(i*10 + 1),
		}
	}

	// When
	i := mockRepo.GetEnergyConsumptionsByMeterIdBetweenDatesCallCount()
	mockRepo.GetEnergyConsumptionsByMeterIdBetweenDatesReturnsOnCall(i, []view.EnergyConsumption{data[0], {}}, nil)
	mockRepo.GetEnergyConsumptionsByMeterIdBetweenDatesReturnsOnCall(i+1, []view.EnergyConsumption{{}, data[1]}, nil)
}

func TestGetEnergyConsumptionsMonthly_Success(t *testing.T) {
	// For
	// Input
	metersIds := []uint8{0, 1}
	startDate, _ := time.Parse("2006-01-02", "2025-01-01")
	endDate, _ := time.Parse("2006-01-02", "2025-02-28")

	mockForPeriods(metersIds)

	// Test
	res, err := controller.GetConsumption(t.Context(), GetConsumptionRequestObject{
		Params: GetConsumptionParams{
			MeterId:   metersIds,
			StartDate: types.Date{Time: startDate},
			EndDate:   types.Date{Time: endDate},
			Period:    Monthly,
		},
	})

	// Assert
	assert.NoError(t, err)
	result, ok := res.(GetConsumption200JSONResponse)
	assert.True(t, ok, "Conversion to 200 response failed")
	assert.Len(t, result.DataGraph, 2)
	assert.Condition(t, func() bool {
		for index, expected := range expectedResult {
			actual := result.DataGraph[index]
			// Ignore address in comparison
			equal := expected.MeterId == actual.MeterId && reflect.DeepEqual(expected.Active, actual.Active)
			if !equal {
				return false
			}
		}
		return true
	}, "result.DataGraph is not equal to expectedResult")
	assert.Equal(t, []string{"January 2025", "February 2025"}, result.Period)
}

func TestGetEnergyConsumptionsWeekly_Success(t *testing.T) {
	// For
	// Input
	metersIds := []uint8{0, 1}
	startDate, _ := time.Parse("2006-01-02", "2025-06-01")
	endDate, _ := time.Parse("2006-01-02", "2025-06-15")

	// Expect
	mockForPeriods(metersIds)

	// Test
	res, err := controller.GetConsumption(t.Context(), GetConsumptionRequestObject{
		Params: GetConsumptionParams{
			MeterId:   metersIds,
			StartDate: types.Date{Time: startDate},
			EndDate:   types.Date{Time: endDate},
			Period:    Weekly,
		},
	})

	// Assert
	assert.NoError(t, err)
	result, ok := res.(GetConsumption200JSONResponse)
	assert.True(t, ok, "Conversion to 200 response failed")
	assert.Condition(t, func() bool {
		for index, expected := range expectedResult {
			actual := result.DataGraph[index]
			// Ignore address in comparison
			equal := expected.MeterId == actual.MeterId && reflect.DeepEqual(expected.Active, actual.Active)
			if !equal {
				return false
			}
		}
		return true
	}, "result.DataGraph is not equal to expectedResult")
	assert.Equal(t, []string{"June 1 - June 7", "June 8 - June 14"}, result.Period)
}

func TestGetEnergyConsumptionsDaily_Success(t *testing.T) {
	// For
	// Input
	metersIds := []uint8{0, 1}
	startDate, _ := time.Parse("2006-01-02", "2025-07-01")
	endDate, _ := time.Parse("2006-01-02", "2025-07-03")

	// Expect
	mockForPeriods(metersIds)

	// Test
	res, err := controller.GetConsumption(t.Context(), GetConsumptionRequestObject{
		Params: GetConsumptionParams{
			MeterId:   metersIds,
			StartDate: types.Date{Time: startDate},
			EndDate:   types.Date{Time: endDate},
			Period:    Daily,
		},
	})

	// Assert
	assert.NoError(t, err)
	result, ok := res.(GetConsumption200JSONResponse)
	assert.True(t, ok, "Conversion to 200 response failed")
	assert.Len(t, result.DataGraph, 2)
	assert.Condition(t, func() bool {
		for index, expected := range expectedResult {
			actual := result.DataGraph[index]
			// Ignore address in comparison
			equal := expected.MeterId == actual.MeterId && reflect.DeepEqual(expected.Active, actual.Active)
			if !equal {
				return false
			}
		}
		return true
	}, "result.DataGraph is not equal to expectedResult")
	assert.Equal(t, []string{"July 1", "July 2"}, result.Period)
}

func TestGetEnergyConsumptionsMonthly_Error(t *testing.T) {
	// For
	// Input
	metersIds := []uint8{0, 1}
	startDate, _ := time.Parse("2006-01-02", "2024-01-01")
	endDate, _ := time.Parse("2006-01-02", "2024-02-27")

	// Expect
	expectedErr := errors.New("database error")

	// When
	mockRepo.GetEnergyConsumptionsByMeterIdBetweenDatesReturns(nil, expectedErr)

	// Test
	res, err := controller.GetConsumption(t.Context(), GetConsumptionRequestObject{
		Params: GetConsumptionParams{
			MeterId:   metersIds,
			StartDate: types.Date{Time: startDate},
			EndDate:   types.Date{Time: endDate},
			Period:    Monthly,
		},
	})

	// Assert
	result, ok := res.(GetConsumption500JSONResponse)
	assert.True(t, ok, "Conversion to 500 response failed")
	assert.Equal(t, "A500", result.Code)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")
}

func TestGetEnergyConsumptionsWeekly_Error(t *testing.T) {
	// For
	// Input
	metersIds := []uint8{0, 1}
	startDate, _ := time.Parse("2006-01-02", "2024-02-01")
	endDate, _ := time.Parse("2006-01-02", "2024-03-01")

	// Expect
	expectedErr := errors.New("database error")

	// When
	mockRepo.GetEnergyConsumptionsByMeterIdBetweenDatesReturns(nil, expectedErr)

	// Test
	res, err := controller.GetConsumption(t.Context(), GetConsumptionRequestObject{
		Params: GetConsumptionParams{
			MeterId:   metersIds,
			StartDate: types.Date{Time: startDate},
			EndDate:   types.Date{Time: endDate},
			Period:    Weekly,
		},
	})

	// Assert
	result, ok := res.(GetConsumption500JSONResponse)
	assert.True(t, ok, "Conversion to 500 response failed")
	assert.Equal(t, "A500", result.Code)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")
}

func TestGetEnergyConsumptionsDaily_Error(t *testing.T) {
	// For
	// Input
	metersIds := []uint8{0, 1}
	startDate, _ := time.Parse("2006-01-02", "2024-02-01")
	endDate, _ := time.Parse("2006-01-02", "2024-03-01")

	// Expect
	expectedErr := errors.New("database error")

	// When
	mockRepo.GetEnergyConsumptionsByMeterIdBetweenDatesReturns(nil, expectedErr)

	// Test
	res, err := controller.GetConsumption(t.Context(), GetConsumptionRequestObject{
		Params: GetConsumptionParams{
			MeterId:   metersIds,
			StartDate: types.Date{Time: startDate},
			EndDate:   types.Date{Time: endDate},
			Period:    Daily,
		},
	})

	// Assert
	result, ok := res.(GetConsumption500JSONResponse)
	assert.True(t, ok, "Conversion to 500 response failed")
	assert.Equal(t, "A500", result.Code)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database error")
}
