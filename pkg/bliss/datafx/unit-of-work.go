package datafx

import "database/sql"

type UnitOfWork interface {
	Commit() error
	Rollback() error
	GetTx() *sql.Tx
}

type UnitOfWorkImpl struct {
	tx *sql.Tx
}

var _ UnitOfWork = (*UnitOfWorkImpl)(nil)

func NewUnitOfWork(tx *sql.Tx) *UnitOfWorkImpl {
	return &UnitOfWorkImpl{tx: tx}
}

func (uow *UnitOfWorkImpl) Commit() error {
	return uow.tx.Commit() //nolint:wrapcheck
}

func (uow *UnitOfWorkImpl) Rollback() error {
	return uow.tx.Rollback() //nolint:wrapcheck
}

func (uow *UnitOfWorkImpl) GetTx() *sql.Tx {
	return uow.tx
}
