package m_account

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"main.go/common/err"
	"main.go/common/status"
	mAccountDB "main.go/db/m-account"
	"main.go/db/m-account/mocks"
	mMemberDB "main.go/db/m-member"
	mockMember "main.go/db/m-member/mocks"
	"testing"
)

func TestCreate(t *testing.T) {

	mockMAccountDB := &mocks.Service{}
	mockMAccountDB.On("IsExisted", context.Background(), mock.Anything).Return(false, nil)

	reqLibErr := CreateRequest{
		Code: "nike",
	}
	mockMAccountDBIsExistedLibErr := &mocks.Service{}
	mockMAccountDBIsExistedLibErr.On("IsExisted", context.Background(), reqLibErr.Code).Return(false, errors.New(""))

	reqCustomErr := CreateRequest{
		Code: "nike",
	}
	mockMAccountDBIsExistedCustomErr := &mocks.Service{}
	mockMAccountDBIsExistedCustomErr.On("IsExisted", context.Background(), reqCustomErr.Code).Return(false, err.InvalidParameter)

	mockMAccountDBAddErr := &mocks.Service{}
	mockMAccountDBAddErr.On("IsExisted", context.Background(), mock.Anything).Return(false, nil)
	mockMAccountDBAddErr.On("Add", context.Background(), mock.Anything).Return(errors.New(""))

	mockMAccountDBAddCustomErr := &mocks.Service{}
	mockMAccountDBAddCustomErr.On("IsExisted", context.Background(), mock.Anything).Return(false, nil)
	mockMAccountDBAddCustomErr.On("Add", context.Background(), mock.Anything).Return(err.InvalidParameter)

	mockMAccountDBAddSuccess := &mocks.Service{}
	mockMAccountDBAddSuccess.On("IsExisted", context.Background(), mock.Anything).Return(false, nil)
	mockMAccountDBAddSuccess.On("Add", context.Background(), mock.Anything).Return(nil)

	type fields struct {
		mAccountRepo mAccountDB.Service
		mMemberRepo  mMemberDB.Service
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
					Code:    err.NilRequest.Code,
					Message: err.NilRequest.Error(),
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
					Code:    err.HashPasswordFailed.Code,
					Message: err.HashPasswordFailed.Error(),
				},
			},
		},
		{
			name: "test add merchant account db fail because lib",
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
					Code:    err.AddMAccountFailed.Code,
					Message: err.AddMAccountFailed.Error(),
				},
			},
		},
		{
			name: "test add merchant account db fail because system",
			fields: fields{
				mAccountRepo: mockMAccountDBAddCustomErr,
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
					Code:    err.InvalidParameter.Code,
					Message: err.InvalidParameter.Error(),
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
			h := New(tt.fields.mAccountRepo, tt.fields.mMemberRepo)
			got := h.Create(tt.params.ctx, tt.params.req)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestRead(t *testing.T) {

	mockMAccountDBGetLibErr := &mocks.Service{}
	code := "nike"
	mockMAccountDBGetLibErr.On("Get", context.Background(), code).Return(nil, errors.New(""))

	mockMAccountDBGetSysErr := &mocks.Service{}
	code = "nike"
	mockMAccountDBGetSysErr.On("Get", context.Background(), code).Return(nil, err.InvalidParameter)

	mockMAccountDBGetSuccess := &mocks.Service{}
	code = "nike"
	entity := &mAccountDB.MerchantAccountEntity{
		Code: code,
	}
	mockMAccountDBGetSuccess.On("Get", context.Background(), code).Return(entity, nil)

	type fields struct {
		mAccountRepo mAccountDB.Service
		mMemberRepo  mMemberDB.Service
	}
	type params struct {
		ctx  context.Context
		code string
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   ReadResponse
	}{
		{
			name: "test merchant code is nil",
			params: params{
				code: "",
			},
			want: ReadResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.EmptyMerchantCode.Code,
					Message: err.EmptyMerchantCode.Error(),
				},
			},
		},
		{
			name: "test get merchant account db fail because lib",
			params: params{
				ctx:  context.Background(),
				code: code,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBGetLibErr,
			},
			want: ReadResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.GetMAccountFailed.Code,
					Message: err.GetMAccountFailed.Error(),
				},
			},
		},
		{
			name: "test get merchant account db fail because system",
			params: params{
				ctx:  context.Background(),
				code: code,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBGetSysErr,
			},
			want: ReadResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.InvalidParameter.Code,
					Message: err.InvalidParameter.Error(),
				},
			},
		},
		{
			name: "test get success",
			params: params{
				ctx:  context.Background(),
				code: code,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBGetSuccess,
			},
			want: ReadResponse{
				Status: status.Success,
				Data:   entity,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mAccountRepo, tt.fields.mMemberRepo)
			got := h.Read(tt.params.ctx, tt.params.code)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestUpdate(t *testing.T) {

	mockMAccountDBUpdateLibErr := &mocks.Service{}
	reqLibErr := &UpdateRequest{
		MerchantID: "nike",
		Name:       "nike",
	}
	mAccountErr := mAccountDB.MerchantAccount{
		ID:   reqLibErr.MerchantID,
		Name: reqLibErr.Name,
	}
	mockMAccountDBUpdateLibErr.On("Update", context.Background(), mAccountErr).Return(errors.New(""))

	mockMAccountDBUpdateSysErr := &mocks.Service{}
	reqSysErr := &UpdateRequest{
		MerchantID: "nike",
		Name:       "nike",
	}
	mAccountSysErr := mAccountDB.MerchantAccount{
		ID:   reqSysErr.MerchantID,
		Name: reqSysErr.Name,
	}
	mockMAccountDBUpdateSysErr.On("Update", context.Background(), mAccountSysErr).Return(err.InvalidParameter)

	mockMAccountDBUpdateSuccess := &mocks.Service{}
	reqOk := &UpdateRequest{
		MerchantID: "adidas",
		Name:       "adidas",
	}
	mAccountOk := mAccountDB.MerchantAccount{
		ID:   reqOk.MerchantID,
		Name: reqOk.Name,
	}
	mockMAccountDBUpdateSuccess.On("Update", context.Background(), mAccountOk).Return(nil)

	type fields struct {
		mAccountRepo mAccountDB.Service
		mMemberRepo  mMemberDB.Service
	}
	type params struct {
		ctx context.Context
		req *UpdateRequest
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   UpdateResponse
	}{
		{
			name: "test req is nil",
			params: params{
				ctx: context.Background(),
				req: nil,
			},
			want: UpdateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.NilRequest.Code,
					Message: err.NilRequest.Error(),
				},
			},
		},
		{
			name: "test update merchant account fail because lib",
			params: params{
				ctx: context.Background(),
				req: reqLibErr,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBUpdateLibErr,
			},
			want: UpdateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.UpdateMAccountFailed.Code,
					Message: err.UpdateMAccountFailed.Error(),
				},
			},
		},
		{
			name: "test update merchant account fail because sys",
			params: params{
				ctx: context.Background(),
				req: reqSysErr,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBUpdateSysErr,
			},
			want: UpdateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.InvalidParameter.Code,
					Message: err.InvalidParameter.Error(),
				},
			},
		},
		{
			name: "test update merchant account success",
			params: params{
				ctx: context.Background(),
				req: reqOk,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBUpdateSuccess,
			},
			want: UpdateResponse{
				Status: status.Success,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mAccountRepo, tt.fields.mMemberRepo)
			got := h.Update(tt.params.ctx, tt.params.req)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestDelete(t *testing.T) {

	mockMAccountDBDeleteLibErr := &mocks.Service{}
	MerchantID := "id"
	mockMAccountDBDeleteLibErr.On("Delete", context.Background(), MerchantID).Return(errors.New(""))

	mockMAccountDBDeleteSysErr := &mocks.Service{}
	mockMAccountDBDeleteSysErr.On("Delete", context.Background(), MerchantID).Return(err.InvalidParameter)

	mockMAccountDBDeleteSuccess := &mocks.Service{}
	mockMAccountDBDeleteSuccess.On("Delete", context.Background(), MerchantID).Return(nil)
	mockMMemberDBDeleteSuccess := &mockMember.Service{}
	mockMMemberDBDeleteSuccess.On("DeleteByMerchantID", context.Background(), MerchantID).Return(nil)

	mockMMemberDBDeleteErr := &mockMember.Service{}
	mockMMemberDBDeleteErr.On("DeleteByMerchantID", context.Background(), MerchantID).Return(errors.New(""))

	type fields struct {
		mAccountRepo mAccountDB.Service
		mMemberRepo  mMemberDB.Service
	}
	type params struct {
		ctx context.Context
		mID string
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   DeleteResponse
	}{
		{
			name: "test delete merchant account fail because lib",
			params: params{
				ctx: context.Background(),
				mID: MerchantID,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBDeleteLibErr,
			},
			want: DeleteResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.DeleteMAccountFailed.Code,
					Message: err.DeleteMAccountFailed.Error(),
				},
			},
		},
		{
			name: "test delete merchant account fail because system",
			params: params{
				ctx: context.Background(),
				mID: MerchantID,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBDeleteSysErr,
			},
			want: DeleteResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.InvalidParameter.Code,
					Message: err.InvalidParameter.Error(),
				},
			},
		},
		{
			name: "test delete merchant account success",
			params: params{
				ctx: context.Background(),
				mID: MerchantID,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBDeleteSuccess,
				mMemberRepo:  mockMMemberDBDeleteSuccess,
			},
			want: DeleteResponse{
				Status: status.Success,
			},
		},
		{
			name: "test delete merchant account success but delete member of merchant fail",
			params: params{
				ctx: context.Background(),
				mID: MerchantID,
			},
			fields: fields{
				mAccountRepo: mockMAccountDBDeleteSuccess,
				mMemberRepo:  mockMMemberDBDeleteErr,
			},
			want: DeleteResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.DeleteMAccountFailed.Code,
					Message: err.DeleteMAccountFailed.Error(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mAccountRepo, tt.fields.mMemberRepo)
			got := h.Delete(tt.params.ctx, tt.params.mID)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestHashPassword(t *testing.T) {
	password := "chungtm"
	hash, err := HashPassword(password)
	assert.Nil(t, err)

	ok := CheckPasswordHash(password, hash)
	assert.Equal(t, ok, true)
}
