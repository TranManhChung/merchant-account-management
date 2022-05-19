package m_member

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"main.go/common/err"
	"main.go/common/validator"
	"time"
)

type Service interface {
	Add(ctx context.Context, member MerchantMember) error
	Update(ctx context.Context, member MerchantMember) error
	Get(ctx context.Context, email string) (*MerchantMemberEntity, error)
	Delete(ctx context.Context, email string) error
	IsExisted(ctx context.Context, email string) (bool, error)
	GetByMerchantID(ctx context.Context, mID string, offset, limit int) ([]MerchantMemberEntity, error)
	DeleteByMerchantID(ctx context.Context, mID string) error
}

type MMemberRepo struct {
	db        *bun.DB
	validator validator.Validator
}

func NewMMemberRepoRepo(db *bun.DB, lens map[string]int) (*MMemberRepo, error) {
	repo := &MMemberRepo{
		db:        db,
		validator: validator.New(lens),
	}
	if _, er := repo.db.NewCreateTable().IfNotExists().Model((*MerchantMember)(nil)).Exec(context.Background()); er != nil {
		return nil, er
	}
	return repo, nil
}

func (o *MMemberRepo) Add(ctx context.Context, member MerchantMember) error {
	if er := o.validator.CheckEmpty(map[string]interface{}{Email: member.Email, Name: member.Name,
		Address: member.Address, Phone: member.Phone, MerchantID: member.MerchantID}); er != nil {
		return er
	}
	if er := o.validator.CheckLen(map[string]int{Name: len(member.Name), Email: len(member.Email), Address: len(member.Address), Phone: len(member.Phone)}); er != nil {
		return er
	}

	member.IsActive = true
	member.CreatedAt = time.Now()
	member.ID = fmt.Sprintf("%v", time.Now().UnixNano())
	if _, er := o.db.NewInsert().Model(&member).Exec(ctx); er != nil {
		return err.AddMMemberFailed
	}
	return nil
}

func (o *MMemberRepo) Update(ctx context.Context, member MerchantMember) error {
	if er := o.validator.CheckEmpty(map[string]interface{}{Email: member.Email}); er != nil {
		return er
	}
	if er := o.validator.CheckLen(map[string]int{Name: len(member.Name), Phone: len(member.Phone), Address: len(member.Address)}); er != nil {
		return er
	}
	query := o.db.NewUpdate().
		Model(&member)
	if member.Name != "" {
		query.Set("name = ?", member.Name)
	}
	if member.Phone != "" {
		query.Set("phone = ?", member.Phone)
	}
	if member.Address != "" {
		query.Set("address = ?", member.Address)
	}
	resp, er := query.Set("updated_at = ?", time.Now()).
		Where("email = ?", member.Email).
		Where("is_active = true").
		Exec(ctx)
	if er != nil {
		return err.UpdateMMemberFailed
	}
	r, er := resp.RowsAffected()
	if er != nil {
		return err.UpdateMMemberFailed
	}
	if r == 0 {
		return err.NotFound
	}
	return nil
}

func (o *MMemberRepo) Get(ctx context.Context, email string) (*MerchantMemberEntity, error) {
	if er := o.validator.CheckEmpty(map[string]interface{}{Email: email}); er != nil {
		return nil, er
	}
	member := new(MerchantMember)
	if er := o.db.NewSelect().Model(member).Where("email = ?", email).Where("is_active = true").Scan(ctx); er != nil {
		return nil, err.GetMMemberFailed
	}

	entity := member.toEntity()

	return &entity, nil
}

func (o *MMemberRepo) Delete(ctx context.Context, email string) error {
	if er := o.validator.CheckEmpty(map[string]interface{}{Email: email}); er != nil {
		return er
	}

	resp, er := o.db.NewUpdate().
		Model((*MerchantMember)(nil)).
		Set("updated_at = ?", time.Now()).
		Set("is_active = false").
		Where("email = ?", email).
		Where("is_active = true").
		Exec(ctx)
	if er != nil {
		return err.DeleteMMemberFailed
	}
	r, er := resp.RowsAffected()
	if er != nil {
		return err.DeleteMMemberFailed
	}
	if r == 0 {
		return err.NotFound
	}

	return nil
}

func (o *MMemberRepo) DeleteByMerchantID(ctx context.Context, mID string) error {
	if er := o.validator.CheckEmpty(map[string]interface{}{MerchantID: mID}); er != nil {
		return er
	}

	resp, er := o.db.NewUpdate().
		Model((*MerchantMember)(nil)).
		Set("updated_at = ?", time.Now()).
		Set("is_active = false").
		Where("merchant_id = ?", mID).
		Where("is_active = true").
		Exec(ctx)
	if er != nil {
		return err.DeleteMMemberFailed
	}
	r, er := resp.RowsAffected()
	if er != nil {
		return err.DeleteMMemberFailed
	}
	if r == 0 {
		return err.NotFound
	}

	return nil
}

func (o *MMemberRepo) IsExisted(ctx context.Context, email string) (bool, error) {
	if er := o.validator.CheckEmpty(map[string]interface{}{Email: email}); er != nil {
		return false, er
	}
	exists, er := o.db.NewSelect().Model((*MerchantMember)(nil)).Where("email = ?", email).Exists(ctx)
	if er != nil {
		return false, err.CheckExistenceFailed
	}
	if exists {
		return true, nil
	}
	return false, nil
}

func (o *MMemberRepo) GetByMerchantID(ctx context.Context, mID string, offset, limit int) ([]MerchantMemberEntity, error) {
	var entities []MerchantMemberEntity
	var members []MerchantMember
	if er := o.validator.CheckEmpty(map[string]interface{}{MerchantID: mID}); er != nil {
		return entities, er
	}
	if er := o.db.NewSelect().Model(&members).Where("merchant_id = ?", mID).Where("is_active = true").Limit(limit).Offset(offset).Scan(ctx); er != nil {
		return entities, err.GetMMemberFailed
	}
	for _, m := range members {
		entities = append(entities, m.toEntity())
	}
	return entities, nil
}

var _ bun.AfterCreateTableHook = (*MerchantMember)(nil)

func (*MerchantMember) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		IfNotExists().
		Model((*MerchantMember)(nil)).
		Index("email_idx").
		Column("email").
		Exec(ctx)
	return err
}
