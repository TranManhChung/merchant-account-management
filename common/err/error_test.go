package err

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"main.go/common/status"
	"testing"
)

func TestToExternalError(t *testing.T) {

	tests := []struct {
		name   string
		fields InternalError
		params error
		want   *Error
	}{
		{
			name:   "1 test nil input error",
			fields: NilRequest,
			params: errors.New("test"),
			want: &Error{
				Domain:  status.Domain,
				Code:    NilRequest.Code,
				Message: NilRequest.Error(),
			},
		},
		{
			name:   "2 test parse error success",
			fields: NilRequest,
			params: HashPasswordFailed,
			want: &Error{
				Domain:  status.Domain,
				Code:    HashPasswordFailed.Code,
				Message: HashPasswordFailed.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.ToExternalError(tt.params)
			assert.Equal(t, tt.want, got)
		})
	}
}
