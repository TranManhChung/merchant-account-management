package m_account

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"main.go/common/err"
	"main.go/common/status"
	mAccountDB "main.go/db/m-account"
	mMemberDB "main.go/db/m-member"
)

type Service interface {
	Create(ctx context.Context, req *CreateRequest) CreateResponse
	Read(ctx context.Context, id string) ReadResponse
	Update(ctx context.Context, req *UpdateRequest) UpdateResponse
	Delete(ctx context.Context, id string) DeleteResponse
}

type Handler struct {
	mAccountRepo mAccountDB.Service
	mMemberRepo  mMemberDB.Service
}

func New(mAccountRepo mAccountDB.Service, mMemberRepo mMemberDB.Service) Handler {
	return Handler{
		mAccountRepo: mAccountRepo,
		mMemberRepo:  mMemberRepo,
	}
}

func (h Handler) Create(ctx context.Context, req *CreateRequest) CreateResponse {
	if req == nil {
		return CreateResponse{
			Status: status.Failed,
			Error:  err.NilRequest.ToExternalError(nil),
		}
	}

	pwd, er := HashPassword(req.Password)
	if er != nil {
		return CreateResponse{
			Status: status.Failed,
			Error:  err.HashPasswordFailed.ToExternalError(nil),
		}
	}
	if er = h.mAccountRepo.Add(ctx, mAccountDB.MerchantAccount{
		Code:     req.Code,
		Name:     req.Name,
		UserName: req.UserName,
		Password: pwd,
	}); er != nil {
		return CreateResponse{
			Status: status.Failed,
			Error:  err.AddMAccountFailed.ToExternalError(er),
		}
	}
	return CreateResponse{
		Status: status.Success,
	}
}

func (h Handler) Read(ctx context.Context, id string) ReadResponse {
	if id == "" {
		return ReadResponse{
			Status: status.Failed,
			Error:  err.EmptyMerchantCode.ToExternalError(nil),
		}
	}
	entity, er := h.mAccountRepo.Get(ctx, id)
	if er != nil {
		return ReadResponse{
			Status: status.Failed,
			Error:  err.GetMAccountFailed.ToExternalError(er),
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
	var pwd string
	var er error
	pwd, er = HashPassword(req.Password)
	if er != nil && er.Error() != err.EmptyPassword.Error() {
		return UpdateResponse{
			Status: status.Failed,
			Error:  err.HashPasswordFailed.ToExternalError(nil),
		}
	}

	if er = h.mAccountRepo.Update(ctx, mAccountDB.MerchantAccount{
		ID: req.MerchantID,
		Name:     req.Name,
		Password: pwd,
	}); er != nil {
		return UpdateResponse{
			Status: status.Failed,
			Error:  err.UpdateMAccountFailed.ToExternalError(er),
		}
	}
	return UpdateResponse{
		Status: status.Success,
	}
}

func (h Handler) Delete(ctx context.Context, id string) DeleteResponse{

	if er := h.mAccountRepo.Delete(ctx, id); er != nil {
		return DeleteResponse{
			Status: status.Failed,
			Error:  err.DeleteMAccountFailed.ToExternalError(er),
		}
	}

	if er := h.mMemberRepo.DeleteByMerchantID(ctx, id); er != nil && err.DeleteMAccountFailed.ToExternalError(er).Code != err.NotFound.Code{
		return DeleteResponse{
			Status: status.Failed,
			Error:  err.DeleteMAccountFailed.ToExternalError(er),
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
