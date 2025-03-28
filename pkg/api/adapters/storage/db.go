package storage

import (
	"errors"
	"fmt"

	"github.com/eser/ajan/datafx"
)

var ErrDatasourceNotFound = errors.New("datasource not found")

func NewFromDefault(dataRegistry *datafx.Registry) (*Queries, error) {
	datasource := dataRegistry.GetDefault()

	if datasource == nil {
		return nil, fmt.Errorf("%w - default", ErrDatasourceNotFound)
	}

	db := datasource.GetConnection()

	return &Queries{db: db}, nil
}

func NewFromNamed(dataRegistry *datafx.Registry, name string) (*Queries, error) {
	datasource := dataRegistry.GetNamed(name)

	if datasource == nil {
		return nil, fmt.Errorf("%w - %s", ErrDatasourceNotFound, name)
	}

	db := datasource.GetConnection()

	return &Queries{db: db}, nil
}
