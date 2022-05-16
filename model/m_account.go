package model

import (
	"time"
)

type MerchantAccount struct {
	ID        int64  `bun:",pk,autoincrement"`
	Code      string `bun:",unique"`
	Name      string
	UserName  string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MerchantAccountEntity struct {
	Code     string `json:"merchant_code"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (m MerchantAccount) ToEntity() MerchantAccountEntity {
	return MerchantAccountEntity{
		Code:     m.Code,
		Name:     m.Name,
		IsActive: m.IsActive,
	}
}
