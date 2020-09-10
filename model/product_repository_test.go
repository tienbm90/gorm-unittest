package model

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	//"gorm-ut/pkg/model"
)
var products = []Product{
	Product{
		Model: gorm.Model{ID: 1},
		Code:  "code 1",
		Price: 1,
	}, {
		Model: gorm.Model{ID: 2},
		Code:  "code 2",
		Price: 2,
	},{
		Model: gorm.Model{ID: 3},
		Code:  "code 3",
		Price: 3,
	},
}
//func (s *Suite) TestProduct_FindAll() {
//	s.mock.ExpectQuery(
//		//regexp.QuoteMeta(
//		`SELECT id, code, price FROM "products" WHERE "products"."deleted_at" IS NULL`).
//		//WithArgs(1).
//		WillReturnRows(
//			sqlmock.NewRows([]string{"id", "code", "price"}).
//				AddRow(products[0].ID, products[0].Code, products[0].Price).
//				AddRow(products[1].ID, products[1].Code, products[1].Price))
//
//	res, err := s.productRepository.FindAll()
//
//	require.NoError(s.T(), err)
//	require.Equal(s.T(), res, products)
//	//require.Nil(s.T(), deep.Equal(products, res))
//
//}
//
//func (s *Suite) TestProduct_FindByID() {
//	s.mock.ExpectQuery(
//		//regexp.QuoteMeta(
//		`SELECT id, code, price FROM "products" WHERE "products"."deleted_at" IS NULL  AND id = ?`).
//		WithArgs(1).
//		WillReturnRows(
//			sqlmock.NewRows([]string{"id", "code", "price"}).
//				AddRow(products[0].ID, products[0].Code, products[0].Price))
//				//AddRow(products[0].ID, products[0].Code, products[0].Price).
//				//AddRow(products[1].ID, products[1].Code, products[1].Price))
//
//	res, err := s.productRepository.FindByID(1)
//
//	require.NoError(s.T(), err)
//	require.Equal(s.T(), res[0], products[0])
//
//}
//func (s *Suite) TestProduct_Create() {
//	s.mock.ExpectExec(
//		`INSERT INTO "products" ("id","created_at","updated_at","deleted_at","code","price") VALUES ('3','2020-09-04 10:19:47','2020-09-04 10:19:47',NULL,'code 3','3') RETURNING "products"."id"`).
//		//`INSERT INTO "products" ("id","created_at","updated_at","deleted_at","code","price") VALUES ('3','2020-09-04 10:18:45','2020-09-04 10:18:45',NULL,'code 3','3') RETURNING "products"."id"`).
//		//WithArgs(products[2].ID,products[2].CreatedAt, products[2].UpdatedAt, products[2].DeletedAt, products[2].Code, products[2].Price)
//		WillReturnResult(sqlmock.NewResult(1, 1))
//	_, err := s.productRepository.Create(products[2])
//
//	require.NoError(s.T(), err)
//
//}

func TestProductRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb , err := gorm.Open("postgres", db)

	productRepo := ProvideProductRepostiory(gdb)

	mock.ExpectExec(`INSERT INTO "products" ("id","created_at","updated_at","deleted_at","code","price") VALUES (?,?,?,?,?,?) RETURNING "products"."id"`).WithArgs(products[2].ID,products[2].CreatedAt, products[2].UpdatedAt, products[2].DeletedAt, products[2].Code, products[2].Price).
	//mock.ExpectExec(`INSERT INTO "products" ("id","created_at","updated_at","deleted_at","code","price") VALUES ('3','2020-09-04 15:29:40','2020-09-04 15:29:40',NULL,'code 3','3') RETURNING "products"."id" `).
		WillReturnResult(sqlmock.NewResult(1,1))
		//WithArgs("ST")
	_, err = productRepo.Create(products[2])
	assert.NoError(t, err)
}