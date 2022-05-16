package m_account

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"main.go/common/err"
	"main.go/common/status"
	mAccountDB "main.go/db/m-account"
	"main.go/model"
)

type Service interface {
	Create(ctx context.Context, req *CreateRequest) CreateResponse
	Read(ctx context.Context, code string) ReadResponse
	Update(ctx context.Context, req *UpdateRequest) UpdateResponse
	Delete(ctx context.Context, req *DeleteRequest) DeleteResponse
}

type Handler struct {
	mAccountRepo mAccountDB.Service
}

func New(mAccountRepo mAccountDB.Service) Handler {
	return Handler{
		mAccountRepo: mAccountRepo,
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
	if isExisted, er := h.mAccountRepo.IsExisted(ctx, req.Code); er != nil {
		customErr := er.(err.InternalError)
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    customErr.Code(),
				Message: customErr.Error(),
			},
		}
	} else if isExisted {
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.MerchantCodeExisted.Code(),
				Message: err.MerchantCodeExisted.Error(),
			},
		}
	}
	pwd, er := HashPassword(req.Password)
	if er != nil {
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.HashPasswordFailed.Code(),
				Message: err.HashPasswordFailed.Error(),
			},
		}
	}
	if er = h.mAccountRepo.Add(ctx, model.MerchantAccount{
		Code:     req.Code,
		Name:     req.Name,
		UserName: req.UserName,
		Password: pwd,
	}); er != nil {
		customErr := er.(err.InternalError)
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    customErr.Code(),
				Message: customErr.Error(),
			},
		}
	}
	return CreateResponse{
		Status: status.Success,
	}
}

func (h Handler) Read(ctx context.Context, code string) ReadResponse {
	if code == "" {
		return ReadResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.EmptyMerchantCode.Code(),
				Message: err.EmptyMerchantCode.Error(),
			},
		}
	}
	entity, er := h.mAccountRepo.Get(ctx, code)
	if er != nil {
		customErr := er.(err.InternalError)
		return ReadResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    customErr.Code(),
				Message: customErr.Error(),
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
	var pwd string
	var er error
	pwd, er = HashPassword(req.Password)
	if er != nil && er.Error() != err.EmptyPassword.Error() {
		return UpdateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.HashPasswordFailed.Code(),
				Message: err.HashPasswordFailed.Error(),
			},
		}
	}

	if er = h.mAccountRepo.Update(ctx, model.MerchantAccount{
		Code:     req.Code,
		Name:     req.Name,
		Password: pwd,
	}); er != nil {
		customErr := er.(err.InternalError)
		return UpdateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    customErr.Code(),
				Message: customErr.Error(),
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

	if er := h.mAccountRepo.Delete(ctx, req.Code); er != nil {
		customErr := er.(err.InternalError)
		return DeleteResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    customErr.Code(),
				Message: customErr.Error(),
			},
		}
	}
	return DeleteResponse{
		Status: status.Success,
	}
}

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", err.EmptyPassword
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
