package m_member

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
	GlobalMemberNameForTesting = "nike"
	GlobalLensForTesting       = map[string]int{Email: 20, Name: 20, Address: 20, Phone: 20, MerchantID: 20}
)

func TestMain(m *testing.M) {
	bunDB, err := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	if err != nil {
		panic(err)
	}
	defer bunDB.Close()
	if _, err = bunDB.NewDropTable().Model((*MerchantMember)(nil)).IfExists().Exec(context.Background()); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestNewMMemberRepoRepo(t *testing.T) {
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
		repo    *MMemberRepo
	}
	tests := []struct {
		name   string
		params params
		want   want
	}{
		{
			name: "1 new member repo success",
			params: params{
				db: bunDB,
			},
			want: want{
				wantErr: false,
				repo: &MMemberRepo{
					db:        bunDB,
					validator: validator.New(GlobalLensForTesting),
				},
			},
		},
		{
			name: "2 new member repo fail",
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
			gotRepo, er := NewMMemberRepoRepo(tt.params.db, GlobalLensForTesting)
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
		ctx    context.Context
		member MerchantMember
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
			name: "1 add new member success",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Name:       "name",
					Address:    "address",
					Phone:      "phone",
					MerchantID: "1234",
				},
			},
			want: want{
				er: nil,
			},
		},
		{
			name: "2 add new member fail because empty merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:   "email",
					Name:    "name",
					Address: "address",
					Phone:   "phone",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "merchant id is empty"},
			},
		},
		{
			name: "3 add new member fail because empty email",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Name:       "name",
					Address:    "address",
					Phone:      "phone",
					MerchantID: "1234",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member email is empty"},
			},
		},
		{
			name: "4 add new member fail because empty name",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Address:    "address",
					Phone:      "phone",
					MerchantID: "1234",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member name is empty"},
			},
		},
		{
			name: "5 add new member fail because empty address",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Name:       "name",
					Phone:      "phone",
					MerchantID: "1234",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member address is empty"},
			},
		},
		{
			name: "6 add new member fail because empty phone",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Address:    "address",
					Name:       "name",
					MerchantID: "1234",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member phone is empty"},
			},
		},
		{
			name: "7 add new member fail because too long phone",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Address:    "address",
					Name:       "name",
					MerchantID: "1234",
					Phone:      "phone21271862176218621217",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member phone is too long"},
			},
		},
		{
			name: "8 add new member fail because too long address",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Address:    "addresshfahkdfhkahfkadshkf",
					Name:       "name",
					MerchantID: "1234",
					Phone:      "phone",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member address is too long"},
			},
		},
		{
			name: "9 add new member fail because too long email",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "emailfalskfjalsfjfdlajdflajafds",
					Address:    "address",
					Name:       "name",
					MerchantID: "1234",
					Phone:      "phone",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member email is too long"},
			},
		},
		{
			name: "10 add new member fail because too long name",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Address:    "address",
					Name:       "namefajjflajfkajflajsfklalkfdjla",
					MerchantID: "1234",
					Phone:      "phone",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member name is too long"},
			},
		},
		{
			name: "11 add new member fail because member already exists",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Name:       "name",
					Address:    "address",
					Phone:      "phone",
					MerchantID: "1234",
				},
			},
			want: want{
				er: err.AddMMemberFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMMemberRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got := repo.Add(tt.params.ctx, tt.params.member)
			assert.Equal(t, got, tt.want.er)
		})
	}
}

func TestUpdate(t *testing.T) {

	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx    context.Context
		member MerchantMember
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
			name: "1 update member success",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Name:       "name-update",
					Address:    "address-update",
					Phone:      "phone-update",
					MerchantID: "12345",
				},
			},
			want: want{
				er: nil,
			},
		},
		{
			name: "2 add new member fail because empty email",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Name:       "name",
					Address:    "address",
					Phone:      "phone",
					MerchantID: "1234",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member email is empty"},
			},
		},
		{
			name: "3 add new member fail because too long phone",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Address:    "address",
					Name:       "name",
					MerchantID: "1234",
					Phone:      "phone21271862176218621217",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member phone is too long"},
			},
		},
		{
			name: "4 add new member fail because too long address",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Address:    "addresshfahkdfhkahfkadshkf",
					Name:       "name",
					MerchantID: "1234",
					Phone:      "phone",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member address is too long"},
			},
		},
		{
			name: "5 add new member fail because too long name",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email",
					Address:    "address",
					Name:       "namefajjflajfkajflajsfklalkfdjla",
					MerchantID: "1234",
					Phone:      "phone",
				},
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member name is too long"},
			},
		},
		{
			name: "6 update member fail because member doesn't exist",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				member: MerchantMember{
					Email:      "email11",
					Address:    "address",
					Name:       "name",
					MerchantID: "1234",
					Phone:      "phone",
				},
			},
			want: want{
				er: err.NotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMMemberRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got := repo.Update(tt.params.ctx, tt.params.member)
			assert.Equal(t, got, tt.want.er)
		})
	}
}

func TestGet(t *testing.T) {
	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx   context.Context
		email string
	}
	type want struct {
		er     error
		entity *MerchantMemberEntity
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   want
	}{
		{
			name: "1 get member",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:   context.Background(),
				email: "email",
			},
			want: want{
				er: nil,
				entity: &MerchantMemberEntity{
					Email:      "email",
					Address:    "address-update",
					Name:       "name-update",
					MerchantID: "1234",
					Phone:      "phone-update",
				},
			},
		},
		{
			name: "2 get member fail because empty email",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member email is empty"},
			},
		},
		{
			name: "3 get member fail because member does not exist",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:   context.Background(),
				email: fmt.Sprintf("code_1_%v", time.Now().Unix()),
			},
			want: want{
				er: err.GetMMemberFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMMemberRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got, er := repo.Get(tt.params.ctx, tt.params.email)
			assert.Equal(t, er, tt.want.er)
			assert.Equal(t, got, tt.want.entity)
		})
	}
}

func TestDelete(t *testing.T) {
	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx   context.Context
		email string
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
			name: "1 delete member",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:   context.Background(),
				email: "email",
			},
			want: want{
				er: nil,
			},
		},
		{
			name: "2 delete member fail because empty email",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
			},
			want: want{
				er: err.InternalError{Code: -22, Mess: "member email is empty"},
			},
		},
		{
			name: "3 delete member fail because member has been deleted",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:   context.Background(),
				email: "email",
			},
			want: want{
				er: err.NotFound,
			},
		},
		{
			name: "4 delete member fail because member does not exist",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:   context.Background(),
				email: fmt.Sprintf("code_1_%v", time.Now().Unix()),
			},
			want: want{
				er: err.NotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMMemberRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got := repo.Delete(tt.params.ctx, tt.params.email)
			assert.Equal(t, got, tt.want.er)
		})
	}
}

func TestIsExisted(t *testing.T) {
	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx   context.Context
		email string
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
			name: "1 check existed member",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:   context.Background(),
				email: "email",
			},
			want: want{
				er:        nil,
				isExisted: true,
			},
		},
		{
			name: "2 check existed member fail because empty email",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
			},
			want: want{
				er:        err.InternalError{Code: -22, Mess: "member email is empty"},
				isExisted: false,
			},
		},
		{
			name: "3 check existed member fail because member does not exist",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:   context.Background(),
				email: fmt.Sprintf("code_1_%v", time.Now().Unix()),
			},
			want: want{
				er:        nil,
				isExisted: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMMemberRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got, er := repo.IsExisted(tt.params.ctx, tt.params.email)
			assert.Equal(t, er, tt.want.er)
			assert.Equal(t, got, tt.want.isExisted)
		})
	}
}

func TestGetByMerchantID(t *testing.T) {
	// prepare data

	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	var ets []MerchantMemberEntity
	members := []MerchantMember{
		{
			Email:      "email-1",
			Name:       "name-1",
			Address:    "address-1",
			Phone:      "phone-1",
			MerchantID: "111",
		},
		{
			Email:      "email-2",
			Name:       "name-2",
			Address:    "address-2",
			Phone:      "phone-2",
			MerchantID: "111",
		},
		{
			Email:      "email-3",
			Name:       "name-3",
			Address:    "address-3",
			Phone:      "phone-3",
			MerchantID: "111",
		},
	}
	repo, er := NewMMemberRepoRepo(bunDB, GlobalLensForTesting)
	assert.Nil(t, er)
	for _, v := range members {
		er := repo.Add(context.Background(), v)
		assert.Nil(t, er)
		ets = append(ets, v.toEntity())
	}

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx    context.Context
		mID    string
		offset int
		limit  int
	}
	type want struct {
		er       error
		entities []MerchantMemberEntity
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   want
	}{
		{
			name: "1 get members by merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:    context.Background(),
				mID:    "111",
				offset: 0,
				limit:  10,
			},
			want: want{
				er:       nil,
				entities: ets,
			},
		},
		{
			name: "2 get members by merchant id fail because empty merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:    context.Background(),
				offset: 0,
				limit:  10,
			},
			want: want{
				er:       err.InternalError(err.InternalError{Code: -22, Mess: "merchant id is empty"}),
				entities: nil,
			},
		},
		{
			name: "3 get members by merchant id 2",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx:    context.Background(),
				mID:    "111",
				offset: 0,
				limit:  2,
			},
			want: want{
				er:       nil,
				entities: ets[:2],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMMemberRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got, er := repo.GetByMerchantID(tt.params.ctx, tt.params.mID, tt.params.offset, tt.params.limit)
			assert.Equal(t, er, tt.want.er)
			assert.Equal(t, got, tt.want.entities)
		})
	}
}

func TestDeleteByMerchantID(t *testing.T) {
	// prepare data

	bunDB, er := db.NewBunDB("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "chungtm", "localhost:5432", "postgres", "postgres")
	assert.Nil(t, er)
	defer bunDB.Close()

	var ets []MerchantMemberEntity
	members := []MerchantMember{
		{
			Email:      "email-4",
			Name:       "name-4",
			Address:    "address-4",
			Phone:      "phone-4",
			MerchantID: "111",
		},
	}
	repo, er := NewMMemberRepoRepo(bunDB, GlobalLensForTesting)
	assert.Nil(t, er)
	for _, v := range members {
		er := repo.Add(context.Background(), v)
		assert.Nil(t, er)
		ets = append(ets, v.toEntity())
	}

	type fields struct {
		db *bun.DB
	}
	type params struct {
		ctx context.Context
		mID string
	}
	tests := []struct {
		name   string
		fields fields
		params params
		want   error
	}{
		{
			name: "1 delete members by merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				mID: "111",
			},
			want: nil,
		},
		{
			name: "2 delete members by merchant id fail because empty merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
			},
			want: err.InternalError{Data: interface{}(nil), Code: -22, Mess: "merchant id is empty"},
		},
		{
			name: "3 delete members by merchant id fail because invalid merchant id",
			fields: fields{
				db: bunDB,
			},
			params: params{
				ctx: context.Background(),
				mID: "faldkfjaldf",
			},
			want: err.NotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, er := NewMMemberRepoRepo(tt.fields.db, GlobalLensForTesting)
			assert.Nil(t, er)
			got := repo.DeleteByMerchantID(tt.params.ctx, tt.params.mID)
			assert.Equal(t, got, tt.want)
		})
	}
}
