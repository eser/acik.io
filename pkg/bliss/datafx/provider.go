package datafx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
)

const DefaultDB = "DEFAULT"

type DataProvider interface {
	GetDefault() *sql.DB
	GetNamed(name string) *sql.DB
	CreateUnitOfWork(ctx context.Context) (context.Context, error)
}

type DataProviderImpl struct {
	dbs    map[string]*sql.DB
	logger *slog.Logger
}

var _ DataProvider = (*DataProviderImpl)(nil)

func NewDataProvider(logger *slog.Logger) *DataProviderImpl {
	dbs := make(map[string]*sql.DB)

	return &DataProviderImpl{
		dbs:    dbs,
		logger: logger,
	}
}

func (dataProvider *DataProviderImpl) GetDefault() *sql.DB {
	return dataProvider.dbs[DefaultDB]
}

func (dataProvider *DataProviderImpl) GetNamed(name string) *sql.DB {
	if db, exists := dataProvider.dbs[name]; exists {
		return db
	}

	return nil
}

func (dataProvider *DataProviderImpl) AddConnection(name string, dsn string) error {
	dataProvider.logger.Info(
		"adding database connection",
		slog.String("name", name),
		slog.String("dialect", string(DetermineDialect(dsn))),
	)

	dialect := DetermineDialect(dsn)

	database, err := sql.Open(string(dialect), dsn)
	if err != nil {
		dataProvider.logger.Error(
			"failed to open database connection",
			slog.String("error", err.Error()),
			slog.String("name", name),
		)

		return fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := database.PingContext(context.TODO()); err != nil {
		dataProvider.logger.Error("failed to ping database", slog.String("error", err.Error()), slog.String("name", name))

		return fmt.Errorf("failed to ping database: %w", err)
	}

	dataProvider.dbs[name] = database
	dataProvider.logger.Info("successfully added database connection", slog.String("name", name))

	return nil
}

func (dataProvider *DataProviderImpl) LoadFromConfig(config *Config) error {
	for name, source := range config.Sources {
		err := dataProvider.AddConnection(name, source.DSN)
		if err != nil {
			return fmt.Errorf("failed to add connection for %s: %w", name, err)
		}
	}

	return nil
}

func (dataProvider *DataProviderImpl) CreateUnitOfWork(ctx context.Context) (context.Context, error) {
	defaultDb := dataProvider.GetDefault()
	if defaultDb == nil {
		return nil, errors.New("default database connection not available") //nolint:err113
	}

	uow, err := NewUnitOfWork(defaultDb)
	if err != nil {
		return nil, fmt.Errorf("failed to create unit of work: %w", err)
	}

	newCtx := context.WithValue(ctx, ContextKeyUnitOfWork, uow)

	return newCtx, nil
}

func GetUnitOfWork(ctx context.Context) (UnitOfWork, bool) { //nolint:ireturn
	uow, ok := ctx.Value(ContextKeyUnitOfWork).(UnitOfWork)

	return uow, ok
}
