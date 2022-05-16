package model

import "github.com/golang/protobuf/ptypes/timestamp"

type MerchantAccount struct {
	ID        int64 `bun:",pk,autoincrement"`
	Code      string
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

type MerchantMember struct {
	ID           int64 `bun:",pk,autoincrement"`
	MerchantCode string
	Email        string
	Name         string
	Address      string
	DoB          string
	Phone        string
	Gender       string
	CreatedAt    timestamp.Timestamp
	UpdatedAt    timestamp.Timestamp
}

type MerchantMemberEntity struct {
	MerchantCode string `json:"merchant_code"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	DoB          string `json:"do_b"`
	Phone        string `json:"phone"`
	Gender       string `json:"gender"`
}

func (m MerchantMember) toEntity() MerchantMemberEntity {
	return MerchantMemberEntity{
		MerchantCode: m.MerchantCode,
		Email:        m.Email,
		Name:         m.Name,
		Address:      m.Address,
		DoB:          m.DoB,
		Phone:        m.Phone,
		Gender:       m.Gender,
	}
}
