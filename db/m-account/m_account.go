package m_account

import (
	"context"
	_ "github.com/lib/pq"
	"main.go/model"
)

type Service interface {
	Add(ctx context.Context, account model.MerchantAccount) error
	Update(ctx context.Context, account model.MerchantAccount) error
	Get(ctx context.Context, code string) (model.MerchantAccountEntity, error)
	Delete(ctx context.Context, code string) error
}