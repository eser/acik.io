package di_test

import (
	"context"
	"testing"

	"github.com/eser/acik.io/pkg/bliss/di"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Define test interfaces and implementations.
type Adder interface {
	Add(ctx context.Context, x int, y int) (int, error)
}

type AdderImpl struct {
	di.Implements[Adder]
}

func (AdderImpl) Add(_ context.Context, x int, y int) (int, error) {
	return x + y, nil
}

type Multiplier interface {
	Multiply(ctx context.Context, x int, y int) (int, error)
}

type MultiplierImpl struct {
	di.Implements[Multiplier]
}

func (MultiplierImpl) Multiply(_ context.Context, x int, y int) (int, error) {
	return x * y, nil
}

func TestRegisterAndGetValue(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Register the implementation
	di.Register[Adder](c, AdderImpl{}) //nolint:exhaustruct

	// Get the implementation
	adder, ok := di.Get[Adder](c)
	assert.True(t, ok, "Expected to retrieve Adder implementation")
	assert.NotNil(t, adder, "Adder implementation should not be nil")

	// Use the implementation
	result, err := adder.Add(context.Background(), 2, 3)

	require.NoError(t, err, "Adder.Add should not return an error")
	assert.Equal(t, 5, result, "Adder.Add should return the correct sum")
}

func TestRegisterProviderAndGetValue(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Register the implementation using provider
	di.RegisterProvider[Adder](c, func() Adder {
		return AdderImpl{} //nolint:exhaustruct
	})

	// Get the implementation
	adder, ok := di.Get[Adder](c)
	assert.True(t, ok, "Expected to retrieve Adder implementation")
	assert.NotNil(t, adder, "Adder implementation should not be nil")

	// Use the implementation
	result, err := adder.Add(context.Background(), 2, 3)

	require.NoError(t, err, "Adder.Add should not return an error")
	assert.Equal(t, 5, result, "Adder.Add should return the correct sum")
}

func TestMustGetValue(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Register the implementation
	di.Register[Multiplier](c, MultiplierImpl{}) //nolint:exhaustruct

	// MustGet should return the implementation
	multiplier := di.MustGet[Multiplier](c)
	assert.NotNil(t, multiplier, "Multiplier implementation should not be nil")

	// Use the implementation
	result, err := multiplier.Multiply(context.Background(), 2, 3)

	require.NoError(t, err, "Multiplier.Multiply should not return an error")
	assert.Equal(t, 6, result, "Multiplier.Multiply should return the correct product")
}

func TestMustGetValueWithProvider(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Register the implementation using provider
	di.RegisterProvider[Multiplier](c, func() Multiplier {
		return MultiplierImpl{} //nolint:exhaustruct
	})

	// MustGet should return the implementation
	multiplier := di.MustGet[Multiplier](c)
	assert.NotNil(t, multiplier, "Multiplier implementation should not be nil")

	// Use the implementation
	result, err := multiplier.Multiply(context.Background(), 2, 3)

	require.NoError(t, err, "Multiplier.Multiply should not return an error")
	assert.Equal(t, 6, result, "Multiplier.Multiply should return the correct product")
}

func TestMustGetValue_Panic(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Ensure that MustGet panics if not registered
	assert.PanicsWithValue(t,
		"No implementation registered for type di_test.Adder",
		func() {
			di.MustGet[Adder](c)
		},
		"Expected MustGet to panic when the implementation is not registered",
	)
}

func TestMustGetValue_PanicWithProvider(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Ensure that MustGet panics if not registered
	assert.PanicsWithValue(t,
		"No implementation registered for type di_test.Adder",
		func() {
			di.MustGet[Adder](c)
		},
		"Expected MustGet to panic when the implementation is not registered",
	)
}

func TestGetValue_NotRegistered(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Ensure that Get returns false if not registered
	adder, ok := di.Get[Adder](c)

	assert.False(t, ok, "Expected Get to return false when not registered")
	assert.Zero(t, adder, "Adder should be zero value when not registered")
}

func TestGetValue_NotRegisteredWithProvider(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Ensure that Get returns false if not registered
	adder, ok := di.Get[Adder](c)

	assert.False(t, ok, "Expected Get to return false when not registered")
	assert.Zero(t, adder, "Adder should be zero value when not registered")
}

func TestInvoke(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Register implementations
	di.Register[Adder](c, AdderImpl{})           //nolint:exhaustruct
	di.Register[Multiplier](c, MultiplierImpl{}) //nolint:exhaustruct

	// Invoke a function that takes dependencies
	di.Invoke(c, func(adder Adder, multiplier Multiplier) {
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

	c := di.NewContainer()

	// Register only the Adder implementation
	di.Register[Adder](c, AdderImpl{}) //nolint:exhaustruct

	// Ensure that Invoke panics if a dependency is missing
	assert.PanicsWithValue(t,
		"No implementation registered for type di_test.Multiplier",
		func() {
			di.Invoke(c, func(multiplier Multiplier) {
				// This function should not be called
			})
		},
		"Expected Invoke to panic when a dependency is not registered",
	)
}

func TestInvoke_NonFunction(t *testing.T) {
	t.Parallel()

	c := di.NewContainer()

	// Ensure that Invoke panics if passed a non-function
	assert.PanicsWithValue(t,
		"Invoke parameter must be a function",
		func() {
			di.Invoke(c, 42) // Passing a non-function
		},
		"Expected Invoke to panic when a non-function is passed",
	)
}
