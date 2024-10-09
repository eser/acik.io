package services_test

import (
	"context"
	"testing"

	"github.com/eser/acik.io/pkg/bliss/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Define test interfaces and implementations.
type Adder interface {
	Add(ctx context.Context, x int, y int) (int, error)
}

type AdderImpl struct {
	services.Implements[Adder]
}

func (AdderImpl) Add(_ context.Context, x int, y int) (int, error) {
	return x + y, nil
}

type Multiplier interface {
	Multiply(ctx context.Context, x int, y int) (int, error)
}

type MultiplierImpl struct {
	services.Implements[Multiplier]
}

func (MultiplierImpl) Multiply(_ context.Context, x int, y int) (int, error) {
	return x * y, nil
}

func TestRegisterAndGet(t *testing.T) {
	t.Parallel()

	c := services.NewContainer()

	// Register the implementation
	services.Register[Adder](c, AdderImpl{}) //nolint:exhaustruct

	// Get the implementation
	adder, ok := services.Get[Adder](c)
	assert.True(t, ok, "Expected to retrieve Adder implementation")
	assert.NotNil(t, adder, "Adder implementation should not be nil")

	// Use the implementation
	result, err := adder.Add(context.Background(), 2, 3)

	require.NoError(t, err, "Adder.Add should not return an error")
	assert.Equal(t, 5, result, "Adder.Add should return the correct sum")
}

func TestMustGet(t *testing.T) {
	t.Parallel()

	c := services.NewContainer()

	// Register the implementation
	services.Register[Multiplier](c, MultiplierImpl{}) //nolint:exhaustruct

	// MustGet should return the implementation
	multiplier := services.MustGet[Multiplier](c)
	assert.NotNil(t, multiplier, "Multiplier implementation should not be nil")

	// Use the implementation
	result, err := multiplier.Multiply(context.Background(), 2, 3)

	require.NoError(t, err, "Multiplier.Multiply should not return an error")
	assert.Equal(t, 6, result, "Multiplier.Multiply should return the correct product")
}

func TestMustGet_Panic(t *testing.T) {
	t.Parallel()

	c := services.NewContainer()

	// Ensure that MustGet panics if not registered
	assert.PanicsWithValue(t,
		"No implementation registered for type services_test.Adder",
		func() {
			services.MustGet[Adder](c)
		},
		"Expected MustGet to panic when the implementation is not registered",
	)
}

func TestInvoke(t *testing.T) {
	t.Parallel()

	c := services.NewContainer()

	// Register implementations
	services.Register[Adder](c, AdderImpl{})           //nolint:exhaustruct
	services.Register[Multiplier](c, MultiplierImpl{}) //nolint:exhaustruct

	// Invoke a function that takes dependencies
	services.Invoke(c, func(adder Adder, multiplier Multiplier) {
		sum, err := adder.Add(context.Background(), 2, 3)

		require.NoError(t, err, "Adder.Add should not return an error")
		assert.Equal(t, 5, sum, "Adder.Add should return the correct sum")

		product, err := multiplier.Multiply(context.Background(), 2, 3)

		require.NoError(t, err, "Multiplier.Multiply should not return an error")
		assert.Equal(t, 6, product, "Multiplier.Multiply should return the correct product")
	})
}

func TestInvoke_MissingDependency(t *testing.T) {
	t.Parallel()

	c := services.NewContainer()

	// Register only the Adder implementation
	services.Register[Adder](c, AdderImpl{}) //nolint:exhaustruct

	// Ensure that Invoke panics if a dependency is missing
	assert.PanicsWithValue(t,
		"No implementation registered for type services_test.Multiplier",
		func() {
			services.Invoke(c, func(multiplier Multiplier) {
				// This function should not be called
			})
		},
		"Expected Invoke to panic when a dependency is not registered",
	)
}

func TestInvoke_NonFunction(t *testing.T) {
	t.Parallel()

	c := services.NewContainer()

	// Ensure that Invoke panics if passed a non-function
	assert.PanicsWithValue(t,
		"Invoke parameter must be a function",
		func() {
			services.Invoke(c, 42) // Passing a non-function
		},
		"Expected Invoke to panic when a non-function is passed",
	)
}

func TestGet_NotRegistered(t *testing.T) {
	t.Parallel()

	c := services.NewContainer()

	// Ensure that Get returns false if not registered
	adder, ok := services.Get[Adder](c)

	assert.False(t, ok, "Expected Get to return false when not registered")
	assert.Equal(t, zeroValue[Adder](), adder, "Adder should be zero value when not registered")
}

// Helper function to get zero value of a type.
func zeroValue[T any]() T { //nolint:ireturn
	var zero T

	return zero
}
