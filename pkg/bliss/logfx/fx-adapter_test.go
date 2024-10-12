package logfx_test

import (
	"bufio"
	"errors"
	"log/slog"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/eser/acik.io/pkg/bliss/logfx"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx/fxevent"
)

type MockWriter struct{}

func (m *MockWriter) Write(p []byte) (int, error) {
	return 0, nil
}

func generateFxLogger() (*logfx.FxLogger, *slog.Logger) {
	logger := slog.New(slog.NewJSONHandler(&MockWriter{}, &slog.HandlerOptions{ //nolint:exhaustruct
		Level: slog.LevelDebug,
	}))

	return &logfx.FxLogger{Logger: logger}, logger
}

func TestGetFxLogger(t *testing.T) {
	t.Parallel()

	fxLogger := logfx.GetFxLogger(nil)

	assert.NotNil(t, fxLogger, "GetFxLogger() = nil, want not nil")
}

func TestFxLogger_LogEvent(t *testing.T) { //nolint:paralleltest,funlen
	fxLogger, _ := generateFxLogger()

	tests := []struct {
		name  string
		event fxevent.Event
		want  string
	}{
		{
			name: "OnStartExecuting",
			event: &fxevent.OnStartExecuting{
				FunctionName: "startFunc",
				CallerName:   "callerFunc",
			},
			want: `{"level":"debug","message":"OnStart hook executing: ","callee":"startFunc","caller":"callerFunc"}`,
		},
		{
			name: "OnStartExecuted",
			event: &fxevent.OnStartExecuted{ //nolint:exhaustruct
				FunctionName: "startFunc",
				CallerName:   "callerFunc",
				Runtime:      time.Second,
			},
			want: `{"level":"debug","message":"OnStart hook executing: ","callee":"startFunc","caller":"callerFunc"}`,
		},
		{
			name: "OnStartExecuted with err",
			event: &fxevent.OnStartExecuted{ //nolint:exhaustruct
				FunctionName: "startFunc",
				CallerName:   "callerFunc",
				Runtime:      time.Second,
				Err:          errors.New("error"), //nolint:err113
			},
			want: `{"level":"debug","message":"OnStart hook failed: ","callee":"startFunc","caller":"callerFunc","error":"error"}`, //nolint:lll
		},
		{
			name: "OnStopExecuting",
			event: &fxevent.OnStopExecuting{
				FunctionName: "stopFunc",
				CallerName:   "callerFunc",
			},
			want: `{"level":"debug","message":"OnStart hook executing: ","callee":"stopFunc","caller":"callerFunc"}`,
		},
		{
			name: "OnStopExecuted",
			event: &fxevent.OnStopExecuted{ //nolint:exhaustruct
				FunctionName: "stopFunc",
				CallerName:   "callerFunc",
				Runtime:      time.Second,
			},
			want: `{"level":"debug","message":"OnStart hook executing: ","callee":"stopFunc","caller":"callerFunc"}`,
		},
		{
			name: "OnStopExecuted with err",
			event: &fxevent.OnStopExecuted{
				FunctionName: "stopFunc",
				CallerName:   "callerFunc",
				Runtime:      time.Second,
				Err:          errors.New("error"), //nolint:err113
			},
			want: `{"level":"debug","message":"OnStart hook failed: ","callee":"stopFunc","caller":"callerFunc","error":"error"}`, //nolint:lll
		},
		{
			name: "Supplied",
			event: &fxevent.Supplied{ //nolint:exhaustruct
				TypeName: "typeA",
				Err:      errors.New("error"), //nolint:err113
			},
			want: `{"level":"debug","message":"supplied: ","type":"typeA","error":"error"}`,
		},
		{
			name: "Provided",
			event: &fxevent.Provided{ //nolint:exhaustruct
				ConstructorName: "constructorA",
				OutputTypeNames: []string{"typeA", "typeB"},
			},
			want: `{"level":"debug","message":"provided: ","constructor":"constructorA","types":["typeA","typeB"]}`,
		},
		{
			name: "Decorated",
			event: &fxevent.Decorated{ //nolint:exhaustruct
				DecoratorName:   "decoratorA",
				OutputTypeNames: []string{"typeA", "typeB"},
			},
			want: `{"level":"debug","message":"decorated: ","decorator":"decoratorA","types":["typeA","typeB"]}`,
		},
		{
			name: "Invoking",
			event: &fxevent.Invoking{ //nolint:exhaustruct
				FunctionName: "invokeFunc",
			},
			want: `{"level":"debug","message":"invoking: ","callee":"invokeFunc"}`,
		},
		{
			name: "Started",
			event: &fxevent.Started{
				Err: nil,
			},
			want: `{"level":"debug","message":"Started"}`,
		},
		{
			name: "LoggerInitialized",
			event: &fxevent.LoggerInitialized{
				ConstructorName: "constructorA",
				Err:             nil,
			},
			want: `{"level":"debug","message":"Logger initialized: ","constructor":"constructorA"}`,
		},
	}

	var wg sync.WaitGroup

	wg.Add(len(tests))

	for _, tt := range tests { //nolint:paralleltest
		t.Run(tt.name, func(t *testing.T) {
			go func() {
				defer wg.Done()

				scanner := bufio.NewScanner(os.Stdin)

				for scanner.Scan() {
					assert.Equal(t, tt.want, scanner.Text())
				}

				assert.NoError(t, scanner.Err())
			}()

			fxLogger.LogEvent(tt.event)
		})
	}

	wg.Wait()
}
