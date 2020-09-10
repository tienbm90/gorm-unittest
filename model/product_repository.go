package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProductRepository struct {
	DB *gorm.DB
}

func ProvideProductRepostiory(DB *gorm.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) FindAll() ([]Product, error) {
	var products []Product
	result := p.DB.Model(&Product{}).Select("id, code, price").Find(&products)
	//result := p.DB.Raw("SELECT id, code, price FROM products WHERE products.deleted_at IS NULL").Scan(&products)
	return products, result.Error
}

func (p *ProductRepository) FindByID(id uint) ([]Product, error) {
	var products []Product
	//result := p.DB.Model(&Product{}).Select("id, code, price").Where("id = ?", id).Scan(&products)
	result := p.DB.Raw(
		`SELECT id, code, price FROM "products" WHERE "products"."deleted_at" IS NULL  AND id = ?`, id).Scan(&products)
	return products, result.Error
}

func (p *ProductRepository) Create(product Product) (Product, error) {
	//err := p.DB.Debug().Model(&Product{}).Create(&product).Error
	err := p.DB.Raw(`INSERT INTO "products" ("id","created_at","updated_at","deleted_at","code","price") VALUES (?,?,?,?,?,?)`,
		product.ID, product.CreatedAt, product.UpdatedAt, product.DeletedAt, product.Code, product.Price).Error
	return product, err
}

func (p *ProductRepository) Update(product Product) (Product, error) {
	result := p.DB.Update(&product)
	return product, result.Error
}

func (p *ProductRepository) Delete(product Product) {
	p.DB.Delete(&product)
}
