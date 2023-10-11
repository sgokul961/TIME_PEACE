package repository

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{DB: DB}
}
func (p *categoryRepository) AddCategory(c domain.Category) (domain.Category, error) {

	var b string
	err := p.DB.Raw("INSERT INTO categories (category) VALUES (?) RETURNING category", c.Category).Scan(&b).Error
	if err != nil {
		return domain.Category{}, err
	}
	var categoryResponse domain.Category

	err = p.DB.Raw(`SELECT 
	                   P.id,
					   p.category
					   FROM 
					       categories p
						WHERE 
						     p.category = ?   `, b).Scan(&categoryResponse).Error

	if err != nil {
		return domain.Category{}, err
	}
	return categoryResponse, nil
}
func (p *categoryRepository) CheckCatogory(id uint) (bool, error) {
	var count int

	err := p.DB.Raw(`SELECT COUNT(*) FROM categories WHERE id=?`, id).Scan(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
func (p *categoryRepository) UpdateCategory(id uint, category string) (domain.Category, error) {

	//check database connection

	// if p.DB == nil {
	// 	return domain.Category{}, errors.New("databse connection is nil")
	// }

	//update the category

	if err := p.DB.Exec("UPDATE categories SET category = $1 WHERE id = $2", category, id).Error; err != nil {
		return domain.Category{}, err
	}
	//Retrive the updated category

	var newcat domain.Category

	if err := p.DB.Raw(`SELECT * FROM categories WHERE category=?`, category).Scan(&newcat).Error; err != nil {
		return domain.Category{}, err
	}
	fmt.Println(newcat)

	return newcat, nil
}
func (c *categoryRepository) DeleteCategory(categoryID string) error {

	id, err := strconv.Atoi(categoryID)

	if err != nil {
		return errors.New("convering into integer and not happend")

	}
	fmt.Println("this is the ID:", id)

	result := c.DB.Exec("DELETE FROM categories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")

	}
	return nil
}
