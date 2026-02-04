package repository

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/domain/view"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/db"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/db/dbfakes"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/env"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
	"github.com/stretchr/testify/assert"
)

var mockDB *dbfakes.FakeDatabase

// Initialize the mock database and the log.L.
func TestMain(m *testing.M) {
	log.L.InitLogger(env.Env.LogLevel)
	db.DB = new(dbfakes.FakeDatabase)
	db.DB.InitDatabaseConnection()
	mockDB = db.DB.(*dbfakes.FakeDatabase)

	mockDB.ModelReturns(mockDB)
	mockDB.SelectReturns(mockDB)
	mockDB.WhereReturns(mockDB)
	mockDB.GroupReturns(mockDB)

	code := m.Run()
	os.Exit(code)
}

func TestGetEnergyConsumptionsByMeterIdBetweenDates_Success(t *testing.T) {
	// For
	meterId := uint8(123)
	startDate := time.Now().Add(-24 * time.Hour)
	endDate := time.Now()

	// Expect
	expectedData := []view.EnergyConsumption{
		{MeterId: meterId, Address: "address", TotalConsumption: 1000.35},
	}

	// When
	mockDB.ScanStub = func(a any) db.Database {
		aPointer := a.(*[]view.EnergyConsumption)
		*aPointer = expectedData
		return mockDB
	}
	mockDB.ErrorReturns(nil)

	// Test
	repo := EnergyConsumptionRepository{}
	result, err := repo.GetEnergyConsumptionsByMeterIdBetweenDates([]uint8{meterId}, startDate, endDate)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedData, result)
}

func TestGetEnergyConsumptionsByMeterIdBetweenDates_Error(t *testing.T) {
	// For
	meterId := uint8(123)
	startDate := time.Now().Add(-24 * time.Hour)
	endDate := time.Now()

	// Expect
	expectedErr := errors.New("database error")

	// When
	mockDB.ErrorReturns(expectedErr)
	mockDB.ScanReturns(mockDB)

	// Test
	repo := EnergyConsumptionRepository{}
	result, err := repo.GetEnergyConsumptionsByMeterIdBetweenDates([]uint8{meterId}, startDate, endDate)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to query energy consumptions")
}
