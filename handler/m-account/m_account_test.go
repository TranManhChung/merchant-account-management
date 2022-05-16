package m_account

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/common/err"
	"main.go/common/status"
	mAccountDB "main.go/db/m-account"
	"main.go/db/m-account/mocks"
	"testing"
)

func TestCreate(t *testing.T) {

	mockMAccountDB := &mocks.Service{}

	mockMAccountDBAddErr := &mocks.Service{}
	mockMAccountDBAddErr.On("Add", context.Background(), mock.Anything).Return(errors.New(""))

	mockMAccountDBAddSuccess := &mocks.Service{}
	mockMAccountDBAddSuccess.On("Add", context.Background(), mock.Anything).Return(nil)

	type fields struct {
		mAccountRepo mAccountDB.Service
	}
	type params struct {
		ctx context.Context
		req *CreateRequest
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   CreateResponse
	}{
		{
			name: "test nil request",
			fields: fields{
				mAccountRepo: mockMAccountDB,
			},
			params: params{
				ctx: context.Background(),
				req: nil,
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.NilRequest.Code(),
					Message: err.NilRequest.Error(),
				},
			},
		},
		{
			name: "test code is too long",
			fields: fields{
				mAccountRepo: mockMAccountDB,
			},
			params: params{
				ctx: context.Background(),
				req: &CreateRequest{Code: "1234567890123"},
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.TooLongCode.Code(),
					Message: err.TooLongCode.Error(),
				},
			},
		},
		{
			name: "test password is empty",
			fields: fields{
				mAccountRepo: mockMAccountDB,
			},
			params: params{
				ctx: context.Background(),
				req: &CreateRequest{
					Code:     "1234567",
					Password: "",
				},
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.HashPasswordFailed.Code(),
					Message: err.HashPasswordFailed.Error(),
				},
			},
		},
		{
			name: "test add merchant account db fail",
			fields: fields{
				mAccountRepo: mockMAccountDBAddErr,
			},
			params: params{
				ctx: context.Background(),
				req: &CreateRequest{
					Code:     "code",
					Name:     "name",
					UserName: "username",
					Password: "chungtm",
				},
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.AddMAccountFailed.Code(),
					Message: err.AddMAccountFailed.Error(),
				},
			},
		},
		{
			name: "test create success",
			fields: fields{
				mAccountRepo: mockMAccountDBAddSuccess,
			},
			params: params{
				ctx: context.Background(),
				req: &CreateRequest{
					Code:     "code",
					Name:     "name",
					UserName: "username",
					Password: "chungtm",
				},
			},
			want: CreateResponse{
				Status: status.Success,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mAccountRepo)
			got := h.Create(tt.params.ctx, tt.params.req)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestHashPassword(t *testing.T) {
	password := "chungtm"
	hash, err := HashPassword(password)
	assert.Nil(t, err)
	fmt.Println(hash)
	ok := CheckPasswordHash(password, hash)
	assert.Equal(t, ok, true)
}