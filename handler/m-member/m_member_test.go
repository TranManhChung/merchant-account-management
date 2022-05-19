package m_account

import (
	"context"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"main.go/common/err"
	"main.go/common/status"
	mAccountDB "main.go/db/m-account"
	mockAccount "main.go/db/m-account/mocks"
	mMemberDB "main.go/db/m-member"
	"main.go/db/m-member/mocks"
	"testing"
)

func TestCreate(t *testing.T) {

	req2 := CreateRequest{
		Email:      "email@gmail.com",
		MerchantID: "123",
	}
	mockMA2 := &mockAccount.Service{}
	mockMA2.On("IsExisted", context.Background(), req2.MerchantID).Return(false, errors.New(""))

	req3 := CreateRequest{
		Email:      "email@gmail.com",
		MerchantID: "123",
	}
	mockMA3 := &mockAccount.Service{}
	mockMA3.On("IsExisted", context.Background(), req3.MerchantID).Return(false, nil)

	req4 := CreateRequest{
		Email:      "email@gmail.com",
		MerchantID: "123",
	}
	mockMA4 := &mockAccount.Service{}
	mockMA4.On("IsExisted", context.Background(), req4.MerchantID).Return(true, nil)
	mockMM4 := &mocks.Service{}
	mockMM4.On("IsExisted", context.Background(), req4.Email).Return(false, errors.New(""))

	req5 := CreateRequest{
		Email:      "email@gmail.com",
		MerchantID: "123",
	}
	mockMA5 := &mockAccount.Service{}
	mockMA5.On("IsExisted", context.Background(), req5.MerchantID).Return(true, nil)
	mockMM5 := &mocks.Service{}
	mockMM5.On("IsExisted", context.Background(), req5.Email).Return(true, nil)

	req6 := CreateRequest{
		Email:      "email@gmail.com",
		MerchantID: "123",
	}
	mockMA6 := &mockAccount.Service{}
	mockMA6.On("IsExisted", context.Background(), req6.MerchantID).Return(true, nil)
	mockMM6 := &mocks.Service{}
	mockMM6.On("IsExisted", context.Background(), req6.Email).Return(false, nil)
	mockMM6.On("Add", context.Background(), mMemberDB.MerchantMember{
		Email:      req6.Email,
		MerchantID: req6.MerchantID,
	}).Return(errors.New(""))

	req7 := CreateRequest{
		Email:      "email@gmail.com",
		MerchantID: "123",
	}
	mockMA7 := &mockAccount.Service{}
	mockMA7.On("IsExisted", context.Background(), req7.MerchantID).Return(true, nil)
	mockMM7 := &mocks.Service{}
	mockMM7.On("IsExisted", context.Background(), req7.Email).Return(false, nil)
	mockMM7.On("Add", context.Background(), mMemberDB.MerchantMember{
		Email:      req6.Email,
		MerchantID: req6.MerchantID,
	}).Return(nil)

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
			name: "1 test nil request",
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
			name: "2 test check account existence occur error",
			fields: fields{
				mAccountRepo: mockMA2,
			},
			params: params{
				ctx: context.Background(),
				req: &req2,
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.CheckExistenceFailed.Code,
					Message: err.CheckExistenceFailed.Error(),
				},
			},
		},
		{
			name: "3 test account does not exist",
			fields: fields{
				mAccountRepo: mockMA3,
			},
			params: params{
				ctx: context.Background(),
				req: &req3,
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.CheckExistenceFailed.Code,
					Message: err.CheckExistenceFailed.Error(),
				},
			},
		},
		{
			name: "4 test check member existence occur error",
			fields: fields{
				mAccountRepo: mockMA4,
				mMemberRepo:  mockMM4,
			},
			params: params{
				ctx: context.Background(),
				req: &req4,
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.CheckExistenceFailed.Code,
					Message: err.CheckExistenceFailed.Error(),
				},
			},
		},
		{
			name: "5 test member already exists",
			fields: fields{
				mAccountRepo: mockMA5,
				mMemberRepo:  mockMM5,
			},
			params: params{
				ctx: context.Background(),
				req: &req5,
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.EmailExisted.Code,
					Message: err.EmailExisted.Error(),
				},
			},
		},
		{
			name: "6 test add member occur error",
			fields: fields{
				mAccountRepo: mockMA6,
				mMemberRepo:  mockMM6,
			},
			params: params{
				ctx: context.Background(),
				req: &req6,
			},
			want: CreateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.AddMMemberFailed.Code,
					Message: err.AddMMemberFailed.Error(),
				},
			},
		},
		{
			name: "7 test add member success",
			fields: fields{
				mAccountRepo: mockMA7,
				mMemberRepo:  mockMM7,
			},
			params: params{
				ctx: context.Background(),
				req: &req7,
			},
			want: CreateResponse{
				Status: status.Success,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mMemberRepo, tt.fields.mAccountRepo)
			got := h.Create(tt.params.ctx, tt.params.req)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestRead(t *testing.T) {

	req2 := "email@gmail.com"
	mockMMDB2 := &mocks.Service{}
	mockMMDB2.On("Get", context.Background(), req2).Return(nil, errors.New(""))

	entity3 := &mMemberDB.MerchantMemberEntity{
		Email:   "email",
		Name:    "name",
		Address: "address",
	}
	req3 := "email@gmail.com"
	mockMMDB3 := &mocks.Service{}
	mockMMDB3.On("Get", context.Background(), req3).Return(entity3, nil)

	type fields struct {
		mAccountRepo mAccountDB.Service
		mMemberRepo  mMemberDB.Service
	}
	type params struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   ReadResponse
	}{
		{
			name: "1 test empty email",
			params: params{
				ctx: context.Background(),
			},
			want: ReadResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.EmptyEmail.Code,
					Message: err.EmptyEmail.Error(),
				},
			},
		},
		{
			name: "2 test get member fail",
			fields: fields{
				mMemberRepo: mockMMDB2,
			},
			params: params{
				ctx:   context.Background(),
				email: req2,
			},
			want: ReadResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.GetMMemberFailed.Code,
					Message: err.GetMMemberFailed.Error(),
				},
			},
		},
		{
			name: "3 test get member success",
			fields: fields{
				mMemberRepo: mockMMDB3,
			},
			params: params{
				ctx:   context.Background(),
				email: req3,
			},
			want: ReadResponse{
				Status: status.Success,
				Data:   entity3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mMemberRepo, tt.fields.mAccountRepo)
			got := h.Read(tt.params.ctx, tt.params.email)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestUpdate(t *testing.T) {
	req2 := &UpdateRequest{
		Email: "email",
	}
	mockMM2 := &mocks.Service{}
	mockMM2.On("Update", context.Background(), mMemberDB.MerchantMember{
		Email: req2.Email,
	}).Return(errors.New(""))

	req3 := &UpdateRequest{
		Email: "email",
	}
	mockMM3 := &mocks.Service{}
	mockMM3.On("Update", context.Background(), mMemberDB.MerchantMember{
		Email: req3.Email,
	}).Return(nil)

	type fields struct {
		mMemberRepo  mMemberDB.Service
		mAccountRepo mAccountDB.Service
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
			name: "1 test nil request",
			params: params{
				ctx: context.Background(),
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
			name: "2 test update member fail",
			params: params{
				ctx: context.Background(),
				req: req2,
			},
			fields: fields{
				mMemberRepo: mockMM2,
			},
			want: UpdateResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.UpdateMMemberFailed.Code,
					Message: err.UpdateMMemberFailed.Error(),
				},
			},
		},
		{
			name: "3 test update member success",
			params: params{
				ctx: context.Background(),
				req: req3,
			},
			fields: fields{
				mMemberRepo: mockMM3,
			},
			want: UpdateResponse{
				Status: status.Success,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mMemberRepo, tt.fields.mAccountRepo)
			got := h.Update(tt.params.ctx, tt.params.req)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestDelete(t *testing.T) {

	email2:= "email"
	mockMM2 := &mocks.Service{}
	mockMM2.On("Delete", context.Background(), email2).Return(errors.New(""))

	email3:= "email"
	mockMM3 := &mocks.Service{}
	mockMM3.On("Delete", context.Background(), email3).Return(nil)

	type fields struct {
		mMemberRepo  mMemberDB.Service
		mAccountRepo mAccountDB.Service
	}
	type params struct {
		ctx context.Context
		email string
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   DeleteResponse
	}{
		{
			name: "2 test delete member fail",
			params: params{
				ctx: context.Background(),
				email: email2,
			},
			fields: fields{
				mMemberRepo: mockMM2,
			},
			want: DeleteResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.DeleteMMemberFailed.Code,
					Message: err.DeleteMMemberFailed.Error(),
				},
			},
		},
		{
			name: "3 test delete member success",
			params: params{
				ctx: context.Background(),
				email: email3,
			},
			fields: fields{
				mMemberRepo: mockMM3,
			},
			want: DeleteResponse{
				Status: status.Success,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mMemberRepo, tt.fields.mAccountRepo)
			got := h.Delete(tt.params.ctx, tt.params.email)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestReads(t *testing.T) {

	MerchantID2 := "id"
	Offset2 := 1
	Limit2 := 1
	mockMM2 := &mocks.Service{}
	mockMM2.On("GetByMerchantID", context.Background(), MerchantID2, Offset2, Limit2).Return(nil, errors.New(""))

	entities3 := []mMemberDB.MerchantMemberEntity{
		{
			MerchantID: "1",
		},
		{
			MerchantID: "2",
		},
	}
	MerchantID3 := "id"
	Offset3 := 1
	Limit3 := 1
	mockMM3 := &mocks.Service{}
	mockMM3.On("GetByMerchantID", context.Background(), MerchantID3, Offset3, Limit3).Return(entities3, nil)

	type fields struct {
		mMemberRepo  mMemberDB.Service
		mAccountRepo mAccountDB.Service
	}
	type params struct {
		ctx        context.Context
		MerchantID string
		Offset     int
		Limit      int
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   ReadsResponse
	}{
		{
			name: "2 test get members fail",
			params: params{
				ctx:        context.Background(),
				MerchantID: MerchantID2,
				Offset:     Offset2,
				Limit:      Limit2,
			},
			fields: fields{
				mMemberRepo: mockMM2,
			},
			want: ReadsResponse{
				Status: status.Failed,
				Error: &err.Error{
					Domain:  status.Domain,
					Code:    err.GetMMemberFailed.Code,
					Message: err.GetMMemberFailed.Error(),
				},
			},
		},
		{
			name: "3 test get members success",
			params: params{
				ctx:        context.Background(),
				MerchantID: MerchantID3,
				Offset:     Offset3,
				Limit:      Limit3,
			},
			fields: fields{
				mMemberRepo: mockMM3,
			},
			want: ReadsResponse{
				Status: status.Success,
				Data:   &entities3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(tt.fields.mMemberRepo, tt.fields.mAccountRepo)
			got := h.Reads(tt.params.ctx, tt.params.MerchantID, tt.params.Offset, tt.params.Limit)
			assert.Equal(t, got, tt.want)
		})
	}
}
