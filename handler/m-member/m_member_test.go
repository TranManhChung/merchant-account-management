package m_account

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/common/err"
	"main.go/common/status"
	mMemberDB "main.go/db/m-member"
	"main.go/db/m-member/mocks"
	"main.go/model"
	"testing"
)

func TestCreate(t *testing.T) {
	mockMMemberDB := &mocks.Service{}
	mockMMemberDB.On("IsExisted", context.Background(), mock.Anything).Return(false, nil)

	reqErr := CreateRequest{
		Email: "email@gmail.com",
	}
	mockMMemberDBIsExistedErr := &mocks.Service{}
	mockMMemberDBIsExistedErr.On("IsExisted", context.Background(), reqErr.Email).Return(false, errors.New(""))

	reqErr = CreateRequest{
		Email: "email@gmail.com",
	}
	mockMMemberDBAccountExisted := &mocks.Service{}
	mockMMemberDBAccountExisted.On("IsExisted", context.Background(), reqErr.Email).Return(true, nil)

	reqAddErr := CreateRequest{
		Email: "email@gmail.com",
		Name:  "name",
	}
	mockMMemberDBAddErr := &mocks.Service{}
	mockMMemberDBAddErr.On("IsExisted", context.Background(), mock.Anything).Return(false, nil)
	mockMMemberDBAddErr.On("Add", context.Background(), model.MerchantMember{
		Email: reqAddErr.Email,
		Name:  reqAddErr.Name,
	}).Return(errors.New(""))

	mockMMemberDBAddSuccess := &mocks.Service{}
	mockMMemberDBAddSuccess.On("IsExisted", context.Background(), mock.Anything).Return(false, nil)
	mockMMemberDBAddSuccess.On("Add", context.Background(), mock.Anything).Return(nil)

	type fields struct {
		mMemberRepo mMemberDB.Service
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
				mMemberRepo: mockMMemberDB,
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
			name: "test check existence fail",
			fields: fields{
				mMemberRepo: mockMMemberDBIsExistedErr,
			},
			params: params{
				ctx: context.Background(),
				req: &reqErr,
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.CheckExistenceFailed.Code(),
					Message: err.CheckExistenceFailed.Error(),
				},
			},
		},
		{
			name: "test member existed",
			fields: fields{
				mMemberRepo: mockMMemberDBAccountExisted,
			},
			params: params{
				ctx: context.Background(),
				req: &reqErr,
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.EmailExisted.Code(),
					Message: err.EmailExisted.Error(),
				},
			},
		},
		{
			name: "test add merchant account db fail",
			fields: fields{
				mMemberRepo: mockMMemberDBAddErr,
			},
			params: params{
				ctx: context.Background(),
				req: &reqAddErr,
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.AddMMemberFailed.Code(),
					Message: err.AddMMemberFailed.Error(),
				},
			},
		},
		{
			name: "test create success",
			fields: fields{
				mMemberRepo: mockMMemberDBAddSuccess,
			},
			params: params{
				ctx: context.Background(),
				req: &CreateRequest{
				},
			},
			want: CreateResponse{
				Status: status.Success,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mMemberRepo)
			got := h.Create(tt.params.ctx, tt.params.req)
			assert.Equal(t, got, tt.want)
		})
	}
}
