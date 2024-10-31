package datafx

type ContextKey string

const (
	ContextKeyUnitOfWork ContextKey = "unit-of-work"
)

type UnitOfWork interface {
	Commit() error
	Rollback() error
	GetScope() DbTransactionManager
}

type UnitOfWorkImpl struct {
	scope DbTransactionManager
}

var _ UnitOfWork = (*UnitOfWorkImpl)(nil)

func NewUnitOfWork() *UnitOfWorkImpl {
	// tx, _ := scope.Begin()
	// if err != nil {
	// 	return nil, err //nolint:wrapcheck
	// }

	return &UnitOfWorkImpl{}
}

func (uow *UnitOfWorkImpl) Commit() error {
	return uow.scope.Commit() //nolint:wrapcheck
}

func (uow *UnitOfWorkImpl) Rollback() error {
	return uow.scope.Rollback() //nolint:wrapcheck
}

func (uow *UnitOfWorkImpl) GetScope() DbTransactionManager { //nolint:ireturn
	return uow.scope
}

func (uow *UnitOfWorkImpl) Use(fn func(DbExecutor) any) {
	fn(uow.scope)
}
