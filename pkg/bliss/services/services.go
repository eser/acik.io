package services

import (
	"fmt"
	"reflect"
)

// Implements is a marker type to associate implementations with interfaces.
type Implements[I any] struct{}

var Default = NewContainer() //nolint:gochecknoglobals

// Register registers an implementation instance for a given interface.
func Register[I any](c Container, impl I) {
	targetType := reflect.TypeOf((*I)(nil)).Elem()
	implValue := reflect.ValueOf(impl)

	if targetType.Kind() == reflect.Interface && !implValue.Type().Implements(targetType) {
		panic(fmt.Sprintf("Implementation type %s does not implement interface %s", implValue.Type(), targetType))
	}

	c.Set(targetType, implValue)
}

// Get retrieves a registered implementation for the given interface.
func Get[I any](c Container) (I, bool) { //nolint:ireturn
	interfaceType := reflect.TypeOf((*I)(nil)).Elem()

	impl, ok := c.Resolve(interfaceType)
	if !ok {
		var zero I

		return zero, false
	}

	return impl.Interface().(I), true //nolint:forcetypeassert
}

// MustGet retrieves a registered implementation or panics if not found.
func MustGet[I any](c Container) I { //nolint:ireturn
	interfaceType := reflect.TypeOf((*I)(nil)).Elem()

	impl := c.MustResolve(interfaceType)

	return impl.Interface().(I) //nolint:forcetypeassert
}

// Invoke calls a function, injecting dependencies based on its parameters.
func Invoke(c Container, fn any) {
	fnValue := reflect.ValueOf(fn)

	fnType := fnValue.Type()
	if fnType.Kind() != reflect.Func {
		panic("Invoke parameter must be a function")
	}

	numIn := fnType.NumIn()
	args := make([]reflect.Value, numIn)

	for i := range numIn {
		paramType := fnType.In(i)
		argValue := c.MustResolve(paramType)
		args[i] = argValue
	}

	fnValue.Call(args)
}
