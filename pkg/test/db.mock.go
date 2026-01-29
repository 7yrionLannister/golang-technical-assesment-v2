// Package test contains mocks for testing purposes.
// Do not use these mocks in production code.
package test

import (
	"github.com/7yrionLannister/golang-technical-assesment-v2/internal/domain/view"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/db"
	"github.com/stretchr/testify/mock"
)

// [MockDatabase] is a mock implementation of the [Database] interface
type MockDatabase struct {
	mock.Mock
	MockDB *MockDatabase
}

// naive mock [db.Database.InitDatabaseConnection] implementation.
// It does not perform any database connection, but simply returns nil error.
func (m *MockDatabase) InitDatabaseConnection() error {
	return nil
}

// naive mock [db.Database.Find] implementation that solely retrieves the expected output slice
// or returns the expected error.
func (m *MockDatabase) Find(out any, conds ...any) db.Database {
	m.Called(out, conds)
	return m
}

// naive mock [db.Database.CreateInBatches] implementation that solely returns the expected error
func (m *MockDatabase) CreateInBatches(value any, batchSize int) error {
	argsCall := m.Called(value, batchSize)
	return argsCall.Error(0)
}

// naive mock [db.Database.Model] implementation that solely returns the expected output
func (m *MockDatabase) Model(value any) db.Database {
	m.Called(value)
	return m
}

// naive mock [db.Database.Select] implementation that solely returns the expected output
func (m *MockDatabase) Select(query string, args ...any) db.Database {
	m.Called(query, args)
	return m
}

// naive mock [db.Database.Where] implementation that solely returns the expected output
func (m *MockDatabase) Where(query string, args ...any) db.Database {
	m.Called(query, args)
	return m
}

// naive mock [db.Database.Group] implementation that solely returns the expected output
func (m *MockDatabase) Group(query string) db.Database {
	m.Called(query)
	return m
}

// naive mock [db.Database.Scan] implementation that solely returns the expected output
func (m *MockDatabase) Scan(dest any) db.Database {
	argsCall := m.Called(dest)
	if argsCall.Get(1) == nil {
		switch result := dest.(type) {
		case *[]view.EnergyConsumption:
			*result = argsCall.Get(0).([]view.EnergyConsumption)
		default:
			panic("unsupported type for Scan output")
		}
	}
	return m
}

// naive mock [db.Database.Error] implementation that solely returns the expected error
func (m *MockDatabase) Error() error {
	argsCall := m.Called()
	if argsCall.Get(0) == nil {
		return nil
	}
	return argsCall.Error(0)
}

var MockDB *MockDatabase
