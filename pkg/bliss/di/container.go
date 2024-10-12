package di

import (
	"fmt"
	"reflect"
)

type Provider func(args []any) any

// Container interface defines the methods for dependency injection container.
type Container interface {
	SetValue(interfaceType reflect.Type, value any)
	SetProvider(interfaceType reflect.Type, provider any)

	Resolve(interfaceType reflect.Type) (reflect.Value, bool)
	MustResolve(interfaceType reflect.Type) reflect.Value

	CreateInvoker(fn any) func()
}

type DependencyType int

const (
	DependencyTypeAssignment DependencyType = iota
	DependencyTypeInvocation
)

type DependencyTarget struct {
	Type            DependencyType
	ReflectionType  reflect.Type
	ReflectionValue reflect.Value
}

// ContainerImpl is the concrete implementation of the Container interface.
type ContainerImpl struct {
	dependencies map[reflect.Type]DependencyTarget
	// mutex    sync.RWMutex
}

var _ Container = (*ContainerImpl)(nil)

var (
	reflectTypeError     = reflect.TypeOf((*error)(nil)).Elem()     //nolint:gochecknoglobals
	reflectTypeContainer = reflect.TypeOf((*Container)(nil)).Elem() //nolint:gochecknoglobals
)

// NewContainer creates a new dependency injection container.
func NewContainer() *ContainerImpl {
	return &ContainerImpl{
		dependencies: make(map[reflect.Type]DependencyTarget),
		// mutex:    sync.RWMutex{},
	}
}

func (c *ContainerImpl) SetValue(interfaceType reflect.Type, value any) {
	reflectionValue := reflect.ValueOf(value)
	reflectionType := reflectionValue.Type()

	if !reflectionType.AssignableTo(interfaceType) {
		panic(fmt.Sprintf("Implementation type %s is not assignable to %s", reflectionType, interfaceType))
	}

	// c.mutex.Lock()
	// defer c.mutex.Unlock()
	c.dependencies[interfaceType] = DependencyTarget{
		Type:            DependencyTypeAssignment,
		ReflectionType:  reflectionType,
		ReflectionValue: reflectionValue,
	}
}

func (c *ContainerImpl) SetProvider(interfaceType reflect.Type, provider any) {
	fnValue := reflect.ValueOf(provider)

	fnType := fnValue.Type()
	if fnType.Kind() != reflect.Func {
		panic("Provider must be a function")
	}

	outNum := fnType.NumOut()
	if outNum == 0 || (outNum > 1 && !fnType.Out(1).AssignableTo(reflectTypeError)) {
		panic(
			fmt.Sprintf("Provider must return a single value or a (value, error) pair that is assignable to %s", interfaceType),
		)
	}

	if !fnType.Out(0).AssignableTo(interfaceType) {
		panic(fmt.Sprintf("Provider %s does not return a value that is assignable to %s", fnType, interfaceType))
	}

	c.dependencies[interfaceType] = DependencyTarget{
		Type:            DependencyTypeInvocation,
		ReflectionType:  fnType,
		ReflectionValue: fnValue,
	}
}

func (c *ContainerImpl) Resolve(t reflect.Type) (reflect.Value, bool) {
	if t.Implements(reflectTypeContainer) {
		return reflect.ValueOf(c), true
	}

	// c.mutex.RLock()
	target, ok := c.dependencies[t]
	// c.mutex.RUnlock()

	if !ok {
		return reflect.Value{}, false
	}

	switch target.Type {
	case DependencyTypeAssignment:
		return target.ReflectionValue, true

	case DependencyTypeInvocation:
		args := c.resolveArgs(target.ReflectionType)

		results := target.ReflectionValue.Call(args)
		if len(results) == 2 && !results[1].IsNil() {
			panic(results[1].Interface().(error).Error()) //nolint:forcetypeassert
		}

		return results[0], true
	}

	return reflect.Value{}, false
}

func (c *ContainerImpl) MustResolve(t reflect.Type) reflect.Value {
	value, ok := c.Resolve(t)

	if !ok {
		panic("No implementation registered for type " + t.String())
	}

	return value
}

func (c *ContainerImpl) CreateInvoker(fn any) func() {
	fnValue := reflect.ValueOf(fn)

	fnType := fnValue.Type()
	if fnType.Kind() != reflect.Func {
		panic("Invoke parameter must be a function")
	}

	args := c.resolveArgs(fnType)

	return func() {
		fnValue.Call(args)
	}
}

func (c *ContainerImpl) resolveArgs(fnType reflect.Type) []reflect.Value {
	numIn := fnType.NumIn()
	args := make([]reflect.Value, numIn)

	for i := range args {
		paramType := fnType.In(i)
		argValue := c.MustResolve(paramType)
		args[i] = argValue
	}

	return args
}
