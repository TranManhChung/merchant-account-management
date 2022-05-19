package m_member

import (
	"time"
)

const (
	Email      = "member email"
	Name       = "member name"
	Address    = "member address"
	Phone      = "member phone"
	MerchantID = "merchant id"
)

type MerchantMember struct {
	ID         string `bun:",pk,"`
	Email      string `bun:",unique"`
	MerchantID string
	Name       string
	Address    string
	Phone      string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type MerchantMemberEntity struct {
	MerchantID string `json:"merchant_id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
}

func (m MerchantMember) toEntity() MerchantMemberEntity {
	return MerchantMemberEntity{
		MerchantID: m.MerchantID,
		Email:      m.Email,
		Name:       m.Name,
		Address:    m.Address,
		Phone:      m.Phone,
	}
}
