package validator

import (
	"github.com/stretchr/testify/assert"
	"main.go/common/err"
	"testing"
)

func TestCheckLen(t *testing.T) {

	tests := []struct {
		name   string
		field  map[string]int
		params map[string]int
		want   error
	}{
		{
			name:   "1 length is ok",
			field:  map[string]int{"name": 10},
			params: map[string]int{"name": 1},
			want:   nil,
		},
		{
			name:   "2 length is not ok",
			field:  map[string]int{"name": 10},
			params: map[string]int{"name": 111},
			want:   err.InternalError{Code: err.InvalidParameter.Code, Mess: "name is too long"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := New(tt.field)
			got := validator.CheckLen(tt.params)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCheckEmpty(t *testing.T) {

	tests := []struct {
		name   string
		field  map[string]int
		params map[string]interface{}
		want   error
	}{
		{
			name:   "1 empty string",
			params: map[string]interface{}{"name": ""},
			want:   err.InternalError{Code: err.InvalidParameter.Code, Mess: "name is empty"},
		},
		{
			name:   "2 invalid number",
			params: map[string]interface{}{"name": int64(0)},
			want:   err.InternalError{Code: err.InvalidParameter.Code, Mess: "name is empty"},
		},
		{
			name:   "3 invalid data type",
			params: map[string]interface{}{"name": 0},
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := New(tt.field)
			got := validator.CheckEmpty(tt.params)
			assert.Equal(t, tt.want, got)
		})
	}
}
