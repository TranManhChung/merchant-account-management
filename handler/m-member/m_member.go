package m_account

import (
	"context"
	"main.go/common/err"
	"main.go/common/status"
	mMemberDB "main.go/db/m-member"
	"main.go/model"
)

type Service interface {
	Create(ctx context.Context, req *CreateRequest) CreateResponse
	Read(ctx context.Context, email string) ReadResponse
	Update(ctx context.Context, req *UpdateRequest) UpdateResponse
	Delete(ctx context.Context, req *DeleteRequest) DeleteResponse
}

type Handler struct {
	mMemberRepo mMemberDB.Service
}

func New(mMemberRepo mMemberDB.Service) Handler {
	return Handler{
		mMemberRepo: mMemberRepo,
	}
}

func (h Handler) Create(ctx context.Context, req *CreateRequest) CreateResponse {
	if req == nil {
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.NilRequest.Code(),
				Message: err.NilRequest.Error(),
			},
		}
	}
	if isExisted, er := h.mMemberRepo.IsExisted(ctx, req.Email); er != nil {
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.CheckExistenceFailed.Code(),
				Message: err.CheckExistenceFailed.Error(),
			},
		}
	} else if isExisted {
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.EmailExisted.Code(),
				Message: err.EmailExisted.Error(),
			},
		}
	}
	if er := h.mMemberRepo.Add(ctx, model.MerchantMember{
		Email:        req.Email,
		MerchantCode: req.MerchantCode,
		Name:         req.Name,
		Address:      req.Address,
		DoB:          req.DoB,
		Phone:        req.Phone,
		Gender:       req.Gender,
	}); er != nil {
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.AddMMemberFailed.Code(),
				Message: err.AddMMemberFailed.Error(),
			},
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
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.EmptyEmail.Code(),
				Message: err.EmptyEmail.Error(),
			},
		}
	}
	entity, er := h.mMemberRepo.Get(ctx, email)
	if er != nil {
		return ReadResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.GetMMemberFailed.Code(),
				Message: err.GetMMemberFailed.Error(),
			},
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
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.NilRequest.Code(),
				Message: err.NilRequest.Error(),
			},
		}
	}

	if er := h.mMemberRepo.Update(ctx, model.MerchantMember{
		Email:   req.Email,
		Name:    req.Name,
		Address: req.Address,
		DoB:     req.DoB,
		Phone:   req.Phone,
		Gender:  req.Gender,
	}); er != nil {
		return UpdateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.UpdateMMemberFailed.Code(),
				Message: err.UpdateMMemberFailed.Error(),
			},
		}
	}
	return UpdateResponse{
		Status: status.Success,
	}
}

func (h Handler) Delete(ctx context.Context, req *DeleteRequest) DeleteResponse {
	if req == nil {
		return DeleteResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.NilRequest.Code(),
				Message: err.NilRequest.Error(),
			},
		}
	}

	if er := h.mMemberRepo.Delete(ctx, req.Email); er != nil {
		return DeleteResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.DeleteMMemberFailed.Code(),
				Message: err.DeleteMMemberFailed.Error(),
			},
		}
	}
	return DeleteResponse{
		Status: status.Success,
	}
}
