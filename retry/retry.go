package retry

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/radhianamri/go-libs/errors"
)

func Repeat[T any](retryCount int, retryErrors []error, fn func() (*T, error)) (resp *T, err error) {
	for i := 0; i < retryCount; i++ {
		resp, err = fn()
		if err == nil || !errors.Contains(retryErrors, err) {
			return resp, nil
		}
	}

	return nil, fmt.Errorf("failed to execute function %s after %d retries: %w", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), retryCount, err)
}
