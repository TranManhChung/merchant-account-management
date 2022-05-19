package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBunDB(t *testing.T) {
	type params struct {
		format, username, password, address, database, driverName string
	}

	tests := []struct {
		name    string
		params  params
		wantErr bool
	}{
		{
			name: "test new bunDB success",
			params: params{
				"postgres://%s:%s@%s/%s?sslmode=disable",
				"postgres",
				"chungtm",
				"localhost:5432",
				"postgres",
				"postgres",
			},
			wantErr: false,
		},
		{
			name: "test new bunDB fail",
			params: params{
				"",
				"",
				"",
				"",
				"",
				"",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBunDB(tt.params.format, tt.params.username, tt.params.password, tt.params.address, tt.params.database, tt.params.driverName)
			got := err != nil
			assert.Equal(t, got, tt.wantErr)
		})
	}
}