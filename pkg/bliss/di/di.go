package di

import (
	"reflect"
)

// Implements is a marker type to associate implementations with interfaces.
type Implements[I any] struct{}

var Default = NewContainer() //nolint:gochecknoglobals

// Register registers an implementation instance for a given interface.
func Register[I any](c Container, impl I) {
	targetType := reflect.TypeOf((*I)(nil)).Elem()

	c.SetValue(targetType, impl)
}

func RegisterProvider[I any](c Container, provider any) {
	targetType := reflect.TypeOf((*I)(nil)).Elem()

	c.SetProvider(targetType, provider)
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
	invoker := c.CreateInvoker(fn)

	invoker()
}
