package m_account

import (
	"time"
)

const (
	MerchantID   = "merchant id"
	MerchantCode = "merchant code"
	MerchantName = "merchant name"
	Username     = "username"
	Password     = "password"
)

type MerchantAccount struct {
	ID        string `bun:"id,pk,"`
	Code      string `bun:",unique"`
	Name      string
	UserName  string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MerchantAccountEntity struct {
	ID       string `json:"id"`
	Code     string `json:"merchant_code"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (m MerchantAccount) ToEntity() MerchantAccountEntity {
	return MerchantAccountEntity{
		ID:       m.ID,
		Code:     m.Code,
		Name:     m.Name,
		IsActive: m.IsActive,
	}
}
