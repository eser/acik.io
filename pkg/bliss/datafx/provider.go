package datafx

import "context"

const DefaultDB = "default"

type DataProvider interface {
	GetDefault() DbTransactionManager
	GetNamed(name string) DbTransactionManager

	CreateUnitOfWork(ctx context.Context) context.Context
}

type DataProviderImpl struct {
	dbs map[string]DbTransactionManager
}

var _ DataProvider = (*DataProviderImpl)(nil)

func NewDataProvider() *DataProviderImpl {
	return &DataProviderImpl{
		dbs: map[string]DbTransactionManager{
			DefaultDB: nil,
		},
	}
}

func (dp *DataProviderImpl) GetDefault() DbTransactionManager { //nolint:ireturn
	return dp.dbs[DefaultDB]
}

func (dp *DataProviderImpl) GetNamed(name string) DbTransactionManager { //nolint:ireturn
	return dp.dbs[name]
}

func (dp *DataProviderImpl) CreateUnitOfWork(ctx context.Context) context.Context {
	uow := NewUnitOfWork()
	newCtx := context.WithValue(ctx, ContextKeyUnitOfWork, uow)

	return newCtx
}

func GetUnitOfWork(ctx context.Context) (UnitOfWork, bool) { //nolint:ireturn
	uow, ok := ctx.Value(ContextKeyUnitOfWork).(UnitOfWork)

	return uow, ok
}
