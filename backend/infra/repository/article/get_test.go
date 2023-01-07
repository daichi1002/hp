package article

import (
	"backend/domain/repository"
	"backend/infra"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbMock() (repository.ArticleRepository, sqlmock.Sqlmock, error) {
	sqlDB, mock, _ := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	// repository 初期化
	gormHandler := new(infra.GormHandler)
	gormHandler.DB = mockDB
	repository := NewArticleRepository(gormHandler)
	return repository, mock, err
}

func TestListArticles(t *testing.T) {
	repository, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	// mock設定
	tagRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(1, "テストタグ1", time.Now(), time.Now())
	articleTagRows := sqlmock.NewRows([]string{"article_id", "tag_id"}).
		AddRow("01GKTYV0YZ3MK4ZK34HVWFHAZH", 1)
	articleRows := sqlmock.NewRows([]string{"id", "title", "content", "created_at", "updated_at"}).
		AddRow("01GKTYV0YZ3MK4ZK34HVWFHAZH", "テストタイトル1", "テストコンテンツ1", time.Now(), time.Now())

	// 期待されるSQL
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles"`)).
		WillReturnRows(articleRows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "article_tags" WHERE "article_tags"."article_id" = $1`)).
		WithArgs("01GKTYV0YZ3MK4ZK34HVWFHAZH").
		WillReturnRows(articleTagRows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tags" WHERE "tags"."id" = $1`)).WithArgs(1).
		WillReturnRows(tagRows)

	// テスト対象の関数実行
	_, err = repository.ListArticles()

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find Articles: %v", err)
	}
}

func TestGetArticle(t *testing.T) {
	repository, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	// mock設定
	tagRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(1, "テストタグ1", time.Now(), time.Now()).AddRow(2, "テストタグ2", time.Now(), time.Now())
	articleTagRows := sqlmock.NewRows([]string{"article_id", "tag_id"}).
		AddRow("01GKTYV0YZ3MK4ZK34HVWFHAZH", 1).
		AddRow("01GKTYV0YZ3MK4ZK34HVWFHAZH", 2)
	articleRows := sqlmock.NewRows([]string{"id", "title", "content", "created_at", "updated_at"}).
		AddRow("01GKTYV0YZ3MK4ZK34HVWFHAZH", "テストタイトル1", "テストコンテンツ1", time.Now(), time.Now())

	// 期待されるSQL
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "articles"`)).
		WillReturnRows(articleRows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "article_tags" WHERE "article_tags"."article_id" = $1`)).
		WithArgs("01GKTYV0YZ3MK4ZK34HVWFHAZH").
		WillReturnRows(articleTagRows)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tags" WHERE "tags"."id" IN ($1,$2)`)).WithArgs(1, 2).
		WillReturnRows(tagRows)

	// テスト対象の関数実行
	_, err = repository.GetArticle("01GKTYV0YZ3MK4ZK34HVWFHAZH")

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find Article: %v", err)
	}
}
