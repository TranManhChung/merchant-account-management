package m_account

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"main.go/common/err"
	"main.go/model"
	"time"
)

type Service interface {
	Add(ctx context.Context, account model.MerchantAccount) error
	Update(ctx context.Context, account model.MerchantAccount) error
	Get(ctx context.Context, code string) (*model.MerchantAccountEntity, error)
	Delete(ctx context.Context, code string) error
	IsExisted(ctx context.Context, code string) (bool, error)
}

type MAccountRepo struct {
	db *bun.DB
}

func NewBunDB(format, username, password, address, database, driverName string) (*bun.DB, error) {
	dataSource := fmt.Sprintf(format, username, password, address, database)
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	bunDB := bun.NewDB(db, sqlitedialect.New())

	//bunDB.AddQueryHook(bundebug.NewQueryHook(
	//	bundebug.WithVerbose(true),
	//	bundebug.FromEnv("BUNDEBUG"),
	//))

	return bunDB, nil
}

func NewMAccountRepoRepo(db *bun.DB) *MAccountRepo {
	return &MAccountRepo{
		db: db,
	}
}

func (o *MAccountRepo) Add(ctx context.Context, account model.MerchantAccount) error {
	if account.Code == "" {
		return err.EmptyMerchantCode
	}
	if account.Name == "" {
		return err.EmptyMerchantName
	}
	if account.UserName == "" {
		return err.EmptyUserName
	}
	if account.Password == "" {
		return err.EmptyPassword
	}
	if _, er := o.db.NewCreateTable().IfNotExists().Model((*model.MerchantAccount)(nil)).Exec(ctx); er != nil {
		return err.AddMAccountFailed
	}
	account.IsActive = true
	account.CreatedAt = time.Now()
	if _, er := o.db.NewInsert().Model(&account).Exec(ctx); er != nil {
		return err.AddMAccountFailed
	}
	return nil
}

func (o *MAccountRepo) Update(ctx context.Context, account model.MerchantAccount) error {
	if account.Code == "" {
		return err.EmptyMerchantCode
	}
	query := o.db.NewUpdate().
		Model(&account)
	if account.Name != "" {
		query.Set("name = ?", account.Name)
	}
	if account.Password != "" {
		query.Set("password = ?", account.Password)
	}
	if _, er := query.Set("updated_at = ?", time.Now()).
		Where("code = ?", account.Code).
		Where("is_active != true").
		Exec(ctx); er != nil {
		return err.UpdateMAccountFailed
	}
	return nil
}

func (o *MAccountRepo) Get(ctx context.Context, code string) (*model.MerchantAccountEntity, error) {
	if code == "" {
		return nil, err.EmptyMerchantCode
	}

	account := new(model.MerchantAccount)
	if er := o.db.NewSelect().Model(account).Where("code = ?", code).Where("is_active != true").Scan(ctx); er != nil {
		return nil, err.GetMAccountFailed
	}
	entity := account.ToEntity()

	return &entity, nil
}

func (o *MAccountRepo) Delete(ctx context.Context, code string) error {
	if code == "" {
		return err.EmptyMerchantCode
	}

	if _, er := o.db.NewUpdate().
		Model((*model.MerchantAccount)(nil)).
		Set("updated_at = ?", time.Now()).
		Set("is_active = true").
		Where("code = ?", code).Exec(ctx); er != nil {
		return err.DeleteMAccountFailed
	}

	return nil
}

func (o *MAccountRepo) IsExisted(ctx context.Context, code string) (bool, error) {
	if code == "" {
		return false, err.EmptyMerchantCode
	}
	exists, er := o.db.NewSelect().Model((*model.MerchantAccount)(nil)).Where("code = ?", code).Exists(ctx)
	if er != nil {
		return false, err.CheckExistenceFailed
	}
	if exists {
		return true, nil
	}
	return false, nil
}
