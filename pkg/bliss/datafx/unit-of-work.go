package datafx

import (
	"context"
	"fmt"
)

type ContextKey string

const (
	ContextKeyUnitOfWork ContextKey = "unit-of-work"
)

type UnitOfWork interface {
	Commit() error
	Rollback() error
	GetScope() DbExecutorTx
}

type UnitOfWorkImpl struct {
	scope DbExecutorTx
}

var _ UnitOfWork = (*UnitOfWorkImpl)(nil)

func NewUnitOfWork(db DbExecutorDb) (*UnitOfWorkImpl, error) {
	transaction, err := db.BeginTx(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	return &UnitOfWorkImpl{
		scope: transaction,
	}, nil
}

func (uow *UnitOfWorkImpl) Commit() error {
	return uow.scope.Commit() //nolint:wrapcheck
}

func (uow *UnitOfWorkImpl) Rollback() error {
	return uow.scope.Rollback() //nolint:wrapcheck
}

func (uow *UnitOfWorkImpl) GetScope() DbExecutorTx { //nolint:ireturn
	return uow.scope
}

func (uow *UnitOfWorkImpl) Use(fn func(DbExecutor) any) {
	fn(uow.scope)
}
