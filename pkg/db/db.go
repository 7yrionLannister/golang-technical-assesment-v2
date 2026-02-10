package db

import (
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/env"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/log"
	"github.com/7yrionLannister/golang-technical-assesment-v2/pkg/util"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var gormConfig = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
	Logger: logger.Default.LogMode(logger.Silent),
}

// Database abstraction
//
//go:generate go tool counterfeiter . Database
type Database interface {
	Select(query string, args ...any) Database      // Add a select clause to the query
	Model(value any) Database                       // Specify a model to the query
	Scan(dest any) Database                         // Scan the result into the destination DTO
	Group(query string) Database                    // Add a group by clause to the query
	Where(query string, args ...any) Database       // Add a where clause to the query
	Error() error                                   // Returns the last error that occurred
	InitDatabaseConnection() error                  // Setup global database connection. Call once at the beggining of the application to initialize the connection.
	Find(out any, args ...any) Database             // Find records that match the conditions
	CreateInBatches(value any, batchSize int) error // Create records in batches
}

// Specific implementation of the Database interface using gorm
type GormDatabase struct {
	gormDb *gorm.DB
}

func (g *GormDatabase) Error() error {
	return g.gormDb.Error
}

func (g *GormDatabase) Model(value any) Database {
	return &GormDatabase{g.gormDb.Model(value)}
}

func (g *GormDatabase) Select(query string, args ...any) Database {
	return &GormDatabase{g.gormDb.Select(query, args...)}
}

func (g *GormDatabase) Where(query string, args ...any) Database {
	return &GormDatabase{g.gormDb.Where(query, args...)}
}

func (g *GormDatabase) Find(out any, conds ...any) Database {
	return &GormDatabase{g.gormDb.Find(out, conds...)}
}

func (g *GormDatabase) Scan(dest any) Database {
	return &GormDatabase{g.gormDb.Scan(dest)}
}

func (g *GormDatabase) Group(query string) Database {
	return &GormDatabase{g.gormDb.Group(query)}
}

func (g *GormDatabase) CreateInBatches(value any, batchSize int) error {
	return g.gormDb.CreateInBatches(value, batchSize).Error
}

func (g *GormDatabase) InitDatabaseConnection() error {
	// Connect gorm to database
	log.L.Debug("Connecting to database at " + env.Env.DataBaseUrl)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        env.Env.DataBaseUrl,
	}), gormConfig)
	if err != nil {
		return util.HandleError(err, "Failed to connect to database")
	}
	g.gormDb = db
	log.L.Debug("Connected to database")
	return nil
}
