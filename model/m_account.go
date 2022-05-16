package model

import "github.com/golang/protobuf/ptypes/timestamp"

type MerchantAccount struct {
	ID        int64  `bun:",pk,autoincrement"`
	Code      string `bun:",unique"`
	Name      string
	UserName  string
	Password  string
	IsActive  bool
	CreatedAt timestamp.Timestamp
	UpdatedAt timestamp.Timestamp
}

type MerchantAccountEntity struct {
	Code     string `json:"merchant_code"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (m MerchantAccount) toEntity() MerchantAccountEntity {
	return MerchantAccountEntity{
		Code:     m.Code,
		Name:     m.Name,
		IsActive: m.IsActive,
	}
}
