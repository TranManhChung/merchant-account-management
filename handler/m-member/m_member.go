package m_account

import (
	"context"
	"main.go/common/err"
	"main.go/common/status"
	mAccountDB "main.go/db/m-account"
	mMemberDB "main.go/db/m-member"
)

type Service interface {
	Create(ctx context.Context, req *CreateRequest) CreateResponse
	Read(ctx context.Context, email string) ReadResponse
	Update(ctx context.Context, req *UpdateRequest) UpdateResponse
	Delete(ctx context.Context, email string) DeleteResponse
	Reads(ctx context.Context, merchantID string, offset, limit int) ReadsResponse
}

type Handler struct {
	mMemberRepo  mMemberDB.Service
	mAccountRepo mAccountDB.Service
}

func New(mMemberRepo mMemberDB.Service, mAccountRepo mAccountDB.Service) Handler {
	return Handler{
		mMemberRepo:  mMemberRepo,
		mAccountRepo: mAccountRepo,
	}
}

func (h Handler) Create(ctx context.Context, req *CreateRequest) CreateResponse {
	if req == nil {
		return CreateResponse{
			Status: status.Failed,
			Error:  err.NilRequest.ToExternalError(nil),
		}
	}

	if isExisted, er := h.mAccountRepo.IsExisted(ctx, req.MerchantID); er != nil || !isExisted {
		return CreateResponse{
			Status: status.Failed,
			Error:  err.CheckExistenceFailed.ToExternalError(er),
		}
	}

	if isExisted, er := h.mMemberRepo.IsExisted(ctx, req.Email); er != nil {
		return CreateResponse{
			Status: status.Failed,
			Error:  err.CheckExistenceFailed.ToExternalError(er),
		}
	} else if isExisted {
		return CreateResponse{
			Status: status.Failed,
			Error:  err.EmailExisted.ToExternalError(nil),
		}
	}
	if er := h.mMemberRepo.Add(ctx, mMemberDB.MerchantMember{
		Email:      req.Email,
		MerchantID: req.MerchantID,
		Name:       req.Name,
		Address:    req.Address,
		Phone:      req.Phone,
	}); er != nil {
		return CreateResponse{
			Status: status.Failed,
			Error:  err.AddMMemberFailed.ToExternalError(er),
		}
	}
	return CreateResponse{
		Status: status.Success,
	}
}

func (h Handler) Read(ctx context.Context, email string) ReadResponse {
	if email == "" {
		return ReadResponse{
			Status: status.Failed,
			Error:  err.EmptyEmail.ToExternalError(nil),
		}
	}
	entity, er := h.mMemberRepo.Get(ctx, email)
	if er != nil {
		return ReadResponse{
			Status: status.Failed,
			Error:  err.GetMMemberFailed.ToExternalError(er),
		}
	}
	return ReadResponse{
		Status: status.Success,
		Data:   entity,
	}
}

func (h Handler) Update(ctx context.Context, req *UpdateRequest) UpdateResponse {
	if req == nil {
		return UpdateResponse{
			Status: status.Failed,
			Error:  err.NilRequest.ToExternalError(nil),
		}
	}

	if er := h.mMemberRepo.Update(ctx, mMemberDB.MerchantMember{
		Email:   req.Email,
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}); er != nil {
		return UpdateResponse{
			Status: status.Failed,
			Error:  err.UpdateMMemberFailed.ToExternalError(er),
		}
	}
	return UpdateResponse{
		Status: status.Success,
	}
}

func (h Handler) Delete(ctx context.Context, email string) DeleteResponse{
	if er := h.mMemberRepo.Delete(ctx, email); er != nil {
		return DeleteResponse{
			Status: status.Failed,
			Error:  err.DeleteMMemberFailed.ToExternalError(er),
		}
	}
	return DeleteResponse{
		Status: status.Success,
	}
}

func (h Handler) Reads(ctx context.Context, merchantID string, offset, limit int) ReadsResponse {
	entities, er := h.mMemberRepo.GetByMerchantID(ctx, merchantID, offset, limit)
	if er != nil {
		return ReadsResponse{
			Status: status.Failed,
			Error:  err.GetMMemberFailed.ToExternalError(er),
		}
	}
	return ReadsResponse{
		Status: status.Success,
		Data:   &entities,
	}
}
