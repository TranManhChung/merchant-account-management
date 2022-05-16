package m_account

import (
	"context"
	mMemberDB "main.go/db/m-member"
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
	return CreateResponse{}
}

func (h Handler) Read(ctx context.Context, email string) ReadResponse {
	return ReadResponse{}
}

func (h Handler) Update(ctx context.Context, req *UpdateRequest) UpdateResponse {
	return UpdateResponse{}
}

func (h Handler) Delete(ctx context.Context, req *DeleteRequest) DeleteResponse {
	return DeleteResponse{}
}
