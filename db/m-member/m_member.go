package m_member

import (
	"context"
	_ "github.com/lib/pq"
	"main.go/model"
)

type Service interface {
	Add(ctx context.Context, account model.MerchantMember) error
	Update(ctx context.Context, account model.MerchantMember) error
	Get(ctx context.Context, email string) (model.MerchantMemberEntity, error)
	Delete(ctx context.Context, email string) error
	IsExisted(ctx context.Context, email string) (bool, error)
}
