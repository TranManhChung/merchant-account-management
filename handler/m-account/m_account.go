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
	Read(ctx context.Context, id string) ReadResponse
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
	if len(req.Code) > maxCodLen {
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.TooLongCode.Code(),
				Message: err.TooLongCode.Error(),
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
		return CreateResponse{
			Status: status.Failed,
			Error: &err.Error{
				Domain:  status.Domain,
				Code:    err.AddMAccountFailed.Code(),
				Message: err.AddMAccountFailed.Error(),
			},
		}
	}
	return CreateResponse{
		Status: status.Success,
	}
}

func (h Handler) Read(ctx context.Context, id string) ReadResponse {
	return ReadResponse{}
}

func (h Handler) Update(ctx context.Context, req *UpdateRequest) UpdateResponse {
	return UpdateResponse{}
}

func (h Handler) Delete(ctx context.Context, req *DeleteRequest) DeleteResponse {
	return DeleteResponse{}
}

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", err.NilPassword
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
