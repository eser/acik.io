package services

import (
	"reflect"
	"sync"
)

// Container interface defines the methods for dependency injection container.
type Container interface {
	Set(interfaceType reflect.Type, implValue reflect.Value)
	Resolve(interfaceType reflect.Type) (reflect.Value, bool)
	MustResolve(interfaceType reflect.Type) reflect.Value
}

// ContainerImpl is the concrete implementation of the Container interface.
type ContainerImpl struct {
	registry map[reflect.Type]reflect.Value
	mutex    sync.RWMutex
}

var _ Container = (*ContainerImpl)(nil)

// NewContainer creates a new dependency injection container.
func NewContainer() *ContainerImpl {
	return &ContainerImpl{
		registry: make(map[reflect.Type]reflect.Value),
		mutex:    sync.RWMutex{},
	}
}

func (c *ContainerImpl) Set(interfaceType reflect.Type, implValue reflect.Value) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.registry[interfaceType] = implValue
}

func (c *ContainerImpl) Resolve(t reflect.Type) (reflect.Value, bool) {
	c.mutex.RLock()
	implValue, ok := c.registry[t]
	c.mutex.RUnlock()

	return implValue, ok
}

func (c *ContainerImpl) MustResolve(t reflect.Type) reflect.Value {
	implValue, ok := c.Resolve(t)

	if !ok {
		panic("No implementation registered for type " + t.String())
	}

	return implValue
}
