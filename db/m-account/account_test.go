package m_account

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
	"main.go/common/err"
	"main.go/common/validator"
	"main.go/db"
	"os"
	"testing"
	"time"
)

var (
	GlobalLensForTesting = map[string]int{MerchantCode: 20, MerchantName: 20, Username: 20}
	originalStatus       = MerchantAccount{
		Code:     "code",
		Name:     "nike",
		UserName: "nikevn",
		Password: "123",
	}
)

func TestMain(m *testing.M) {
	bunDB, err := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	if err != nil {
		panic(err)
	}
	defer bunDB.Close()
	if _, err = bunDB.NewDropTable().Model((*MerchantAccount)(nil)).IfExists().Exec(context.Background()); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestNewMAccountRepoRepo(t *testing.T) {
	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	defer bunDB.Close()
	assert.Nil(t, er)

	bunDBErr, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:54332", "postgres", "postgres")
	defer bunDBErr.Close()
	assert.Nil(t, er)

	type params struct {
		db *bun.DB
	}
	type want struct {
		wantErr bool
		repo    *MAccountRepo
	}
	tests := []struct {
		name   string
		params params
		want   want
	}{
		{
			name: "test new account repo success",
			params: params{
				db: bunDB,
			},
			want: want{
				wantErr: false,
				repo: &MAccountRepo{
					db:        bunDB,
					validator: validator.New(GlobalLensForTesting),
				},
			},
		},
		{
			name: "test new account repo fail",
			params: params{
				db: bunDBErr,
			},
			want: want{
				wantErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRepo, er := NewMAccountRepoRepo(tt.params.db, GlobalLensForTesting)
			assert.Equal(t, gotRepo, tt.want.repo)
			isErr := er != nil
			assert.Equal(t, isErr, tt.want.wantErr)
		})
	}
}

func TestAdd(t *testing.T) {
	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx     context.Context
		account MerchantAccount
	}
	type want struct {
		er error
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   want
	}{
		{
			name: "1 add new account success ",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:     context.Background(),
				account: originalStatus,
			},
			want: want{
				er: nil,
			},
		},
		{
			name: "2 add new account fail because empty merchant code",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					Name:     "name",
					UserName: "username",
					Password: "password",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "merchant code is empty"},
			},
		},
		{
			name: "3 add new account fail because empty merchant name",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					Code:     fmt.Sprintf("code_%v", time.Now().Unix()),
					UserName: "username",
					Password: "password",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "merchant name is empty"},
			},
		},
		{
			name: "4 add new account fail because empty username name",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					Code:     fmt.Sprintf("code_%v", time.Now().Unix()),
					Name:     "nike",
					Password: "password",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "username is empty"},
			},
		},
		{
			name: "5 add new account fail because empty password",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					Code:     fmt.Sprintf("code_%v", time.Now().Unix()),
					Name:     "nike",
					UserName: "username",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "password is empty"},
			},
		},
		{
			name: "6 add new account fail because account already exists",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					Code:     "code",
					Name:     "nike",
					UserName: "username",
					Password: "123",
				},
			},
			want: want{
				er: err.AddMAccountFailed,
			},
		},
		{
			name: "7 add new account fail because too long merchant code",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					Code:     "code212121212121212121212121212",
					Name:     "nike",
					UserName: "username",
					Password: "123",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "merchant code is too long"},
			},
		},
		{
			name: "8 add new account fail because too long name",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					Code:     "code",
					Name:     "nike212121212121212121212121212",
					UserName: "username",
					Password: "123",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "merchant name is too long"},
			},
		},
		{
			name: "9 add new account fail because too long username",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					Code:     "code",
					Name:     "qwq",
					UserName: "usernamenike212121212121212121212121212",
					Password: "123",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "username is too long"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMAccountRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got := repo.Add(tt.params.ctx, tt.params.account)
			assert.Equal(t, got, tt.want.er)
		})
	}
}

func TestUpdate(t *testing.T) {

	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	// prepare data
	account := new(MerchantAccount)
	er = bunDB.NewSelect().Model(account).Where("is_active = true").Limit(1).Scan(context.Background())
	assert.Nil(t, er)

	defer func() {
		// reset to original status
		_, er = bunDB.NewUpdate().Model(account).Set("password = ?", originalStatus.Password).Set("name = ?", originalStatus.Name).Set("is_active = true").Where("id = ?", account.ID).Exec(context.Background())
		assert.Nil(t, er)
	}()

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx     context.Context
		account MerchantAccount
	}
	type want struct {
		er error
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   want
	}{
		{
			name: "1 update account success",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					ID:       account.ID,
					Name:     originalStatus.Name + "-update",
					UserName: originalStatus.UserName + "-update",
					Password: originalStatus.Password + "-password",
				},
			},
			want: want{
				er: nil,
			},
		},
		{
			name: "2 update account fail because merchant name is too long",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					ID:       account.ID,
					Name:     "nikefdhsfhakfhkashfkahfkjasdhfkdshafa",
					UserName: "nike-update-usr",
					Password: "nike-update-pwd",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "merchant name is too long"},
			},
		},
		{
			name: "3 update account fail because empty merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "merchant id is empty"},
			},
		},
		{
			name: "4 update account fail because account doesn't exist",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				account: MerchantAccount{
					ID:       "fjadlfjalf",
					Name:     "nike-update-name",
					UserName: "nike-update-usr",
					Password: "nike-update-pwd",
				},
			},
			want: want{
				er: err.NotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMAccountRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got := repo.Update(tt.params.ctx, tt.params.account)
			assert.Equal(t, got, tt.want.er)
		})
	}
}

func TestGet(t *testing.T) {
	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	// prepare data
	account := new(MerchantAccount)
	er = bunDB.NewSelect().Model(account).Where("is_active = true").Limit(1).Scan(context.Background())
	assert.Nil(t, er)

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx context.Context
		id  string
	}
	type want struct {
		er     error
		entity *MerchantAccountEntity
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   want
	}{
		{
			name: "1 get account",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				id:  account.ID,
			},
			want: want{
				er: nil,
				entity: &MerchantAccountEntity{
					ID:       account.ID,
					Code:     "code",
					Name:     "nike",
					IsActive: true,
				},
			},
		},
		{
			name: "2 get account fail because empty merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "merchant id is empty"},
			},
		},
		{
			name: "3 get account fail because account does not exist",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				id:  fmt.Sprintf("code_1_%v", time.Now().Unix()),
			},
			want: want{
				er: err.GetMAccountFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMAccountRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got, er := repo.Get(tt.params.ctx, tt.params.id)
			assert.Equal(t, er, tt.want.er)
			assert.Equal(t, got, tt.want.entity)
		})
	}
}

func TestDelete(t *testing.T) {
	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	// prepare data
	account := new(MerchantAccount)
	er = bunDB.NewSelect().Model(account).Where("is_active = true").Limit(1).Scan(context.Background())
	assert.Nil(t, er)

	defer func() {
		// reset to original status
		_, er = bunDB.NewUpdate().Model(account).Set("password = ?", originalStatus.Password).Set("name = ?", originalStatus.Name).Set("is_active = true").Where("id = ?", account.ID).Exec(context.Background())
		assert.Nil(t, er)
	}()

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx context.Context
		id  string
	}
	type want struct {
		er error
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   want
	}{
		{
			name: "1 case delete account",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				id:  account.ID,
			},
			want: want{
				er: nil,
			},
		},
		{
			name: "2 delete account fail because empty merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "merchant id is empty"},
			},
		},
		{
			name: "3 delete account fail because account has been deleted",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				id:  account.ID,
			},
			want: want{
				er: err.NotFound,
			},
		},
		{
			name: "4 delete account fail because account does not exist",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				id:  fmt.Sprintf("code_1_%v", time.Now().Unix()),
			},
			want: want{
				er: err.NotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMAccountRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got := repo.Delete(tt.params.ctx, tt.params.id)
			assert.Equal(t, got, tt.want.er)
		})
	}
}

func TestIsExisted(t *testing.T) {
	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	// prepare data
	account := new(MerchantAccount)
	er = bunDB.NewSelect().Model(account).Where("is_active = true").Limit(1).Scan(context.Background())
	assert.Nil(t, er)

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx context.Context
		id  string
	}
	type want struct {
		er        error
		isExisted bool
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   want
	}{
		{
			name: "1 check existed account",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				id:  account.ID,
			},
			want: want{
				er:        nil,
				isExisted: true,
			},
		},
		{
			name: "2 check existed account fail because empty merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
			},
			want: want{
				er:        err.InternalError{Code: -22, Mess: "merchant id is empty"},
				isExisted: false,
			},
		},
		{
			name: "3 check existed account fail because account does not exist",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				id:  "faldfjkdjfia",
			},
			want: want{
				er:        nil,
				isExisted: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMAccountRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got, er := repo.IsExisted(tt.params.ctx, tt.params.id)
			assert.Equal(t, er, tt.want.er)
			assert.Equal(t, got, tt.want.isExisted)
		})
	}
}
