package m_account

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
	Add(ctx context.Context, account MerchantAccount) error
	Update(ctx context.Context, account MerchantAccount) error
	Get(ctx context.Context, id string) (*MerchantAccountEntity, error)
	Delete(ctx context.Context, id string) error
	IsExisted(ctx context.Context, id string) (bool, error)
}

type MAccountRepo struct {
	db        *bun.DB
	validator validator.Validator
}

func NewMAccountRepoRepo(db *bun.DB, lens map[string]int) (*MAccountRepo, error) {
	repo := &MAccountRepo{
		db:        db,
		validator: validator.New(lens),
	}
	if _, er := repo.db.NewCreateTable().IfNotExists().Model((*MerchantAccount)(nil)).Exec(context.Background()); er != nil {
		return nil, er
	}
	return repo, nil
}

func (o *MAccountRepo) Add(ctx context.Context, account MerchantAccount) error {
	if er := o.validator.CheckEmpty(map[string]interface{}{MerchantCode: account.Code, MerchantName: account.Name,
		Username: account.UserName, Password: account.Password}); er != nil {
		return er
	}
	if er := o.validator.CheckLen(map[string]int{MerchantName: len(account.Name), MerchantCode: len(account.Code), Username: len(account.UserName)}); er != nil {
		return er
	}
	account.IsActive = true
	account.CreatedAt = time.Now()
	account.ID = fmt.Sprintf("%v", time.Now().UnixNano())
	if _, er := o.db.NewInsert().Model(&account).Exec(ctx); er != nil {
		return err.AddMAccountFailed
	}
	return nil
}

func (o *MAccountRepo) Update(ctx context.Context, account MerchantAccount) error {
	if er := o.validator.CheckEmpty(map[string]interface{}{MerchantID: account.ID}); er != nil {
		return er
	}
	if er := o.validator.CheckLen(map[string]int{MerchantName: len(account.Name)}); er != nil {
		return er
	}
	query := o.db.NewUpdate().
		Model(&account)
	if account.Name != "" {
		query.Set("name = ?", account.Name)
	}
	if account.Password != "" {
		query.Set("password = ?", account.Password)
	}
	resp, er := query.Set("updated_at = ?", time.Now()).
		Where("id = ?", account.ID).
		Where("is_active = true").
		Exec(ctx)
	if er != nil {
		return err.UpdateMAccountFailed
	}
	r, er := resp.RowsAffected()
	if er != nil {
		return err.UpdateMAccountFailed
	}
	if r == 0 {
		return err.NotFound
	}
	return nil
}

func (o *MAccountRepo) Get(ctx context.Context, id string) (*MerchantAccountEntity, error) {
	if er := o.validator.CheckEmpty(map[string]interface{}{MerchantID: id}); er != nil {
		return nil, er
	}
	account := new(MerchantAccount)
	if er := o.db.NewSelect().Model(account).Where("id = ?", id).Where("is_active = true").Scan(ctx); er != nil {
		return nil, err.GetMAccountFailed
	}

	entity := account.ToEntity()

	return &entity, nil
}

func (o *MAccountRepo) Delete(ctx context.Context, id string) error {
	if er := o.validator.CheckEmpty(map[string]interface{}{MerchantID: id}); er != nil {
		return er
	}

	resp, er := o.db.NewUpdate().
		Model((*MerchantAccount)(nil)).
		Set("updated_at = ?", time.Now()).
		Set("is_active = false").
		Where("id = ?", id).
		Where("is_active = true").
		Exec(ctx)
	if er != nil {
		return err.DeleteMAccountFailed
	}
	r, er := resp.RowsAffected()
	if er != nil {
		return err.DeleteMAccountFailed
	}
	if r == 0 {
		return err.NotFound
	}

	return nil
}

func (o *MAccountRepo) IsExisted(ctx context.Context, id string) (bool, error) {
	if er := o.validator.CheckEmpty(map[string]interface{}{MerchantID: id}); er != nil {
		return false, er
	}
	exists, er := o.db.NewSelect().Model((*MerchantAccount)(nil)).Where("id = ?", id).Exists(ctx)
	if er != nil {
		return false, err.CheckExistenceFailed
	}
	if exists {
		return true, nil
	}
	return false, nil
}

var _ bun.AfterCreateTableHook = (*MerchantAccount)(nil)

func (*MerchantAccount) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().
		IfNotExists().
		Model((*MerchantAccount)(nil)).
		Index("code_idx").
		Column("code").
		Exec(ctx)
	return err
}
