package v1

import (
	fmt "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"
)

func TestCommonErrorsWrapping(t *testing.T) {
	errors := []struct {
		name           string
		errorFunc      func(string) error
		errorFuncf     func(string, ...interface{}) error
		validationFunc func(error) bool
	}{
		{"Canceled", Canceled, Canceledf, IsCanceled},
		{"DeadlineExceeded", DeadlineExceeded, DeadlineExceededf, IsDeadlineExceeded},
		{"InvalidArgument", InvalidArgument, InvalidArgumentf, IsInvalidArgument},
		{"NotFound", NotFound, NotFoundf, IsNotFound},
		{"AlreadyExists", AlreadyExists, AlreadyExistsf, IsAlreadyExists},
		{"PermissionDenied", PermissionDenied, PermissionDeniedf, IsPermissionDenied},
		{"PreconditionFailed", PreconditionFailed, PreconditionFailedf, IsPreconditionFailed},
		{"Unauthenticated", Unauthenticated, Unauthenticatedf, IsUnauthenticated},
		{"ResourceExhausted", ResourceExhausted, ResourceExhaustedf, IsResourceExhausted},
		{"Unknown", Unknown, Unknownf, IsUnknown},
		{"Unavailable", Unavailable, Unavailablef, IsUnavailable},
		{"Aborted", Aborted, Abortedf, IsAborted},
	}
	for idx, testCase := range errors {
		t.Run(testCase.name, func(t *testing.T) {
			plain := testCase.errorFunc(testCase.name + " error")
			formatted := testCase.errorFuncf("%s error", testCase.name)
			require.Equal(t, testCase.name+" error", status.Convert(plain).Message())
			require.Equal(t, testCase.name+" error", status.Convert(formatted).Message())

			for _, e := range []error{plain, formatted} {
				wrapped := fmt.Errorf("Wraps: %w", e)
				wrapped2 := fmt.Errorf("Wraps another one: %w", wrapped)
				for idx2 := range errors {
					assert.Equal(t, idx == idx2, errors[idx2].validationFunc(e))
					assert.Equal(t, idx == idx2, errors[idx2].validationFunc(wrapped), "wrapped error %s is not detected as an error", testCase.name)
					assert.Equal(t, idx == idx2, errors[idx2].validationFunc(wrapped2), "wrapped error %s is not detected as an error", testCase.name)
				}
			}
		})
	}
}
