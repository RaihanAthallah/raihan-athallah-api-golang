package repository

import (
	"raihan-athallah-api-golang/app/model"

	"gorm.io/gorm"
)

type ArticleDB interface {
	GetArticleByID(id string) (model.Article, error)
	GetAllArticles() ([]model.Article, error)
	CreateArticle(article model.Article) (model.Article, error)
	UpdateArticle(id string, article model.Article) (model.Article, error)
	DeleteArticle(id string) error
}

type articleDB struct {
	// db *gorm.DB // Assuming you're using a SQL database, you can use sql.DB here
}

func NewArticleDB(db *gorm.DB) ArticleDB {
	return &articleDB{}
}

func (a *articleDB) GetArticleByID(id string) (model.Article, error) {
	// Implement the logic to get an article by ID from the database
	// Example:
	// var article Article
	// err := a.db.First(&article, id).Error
	// return article, err
	return model.Article{}, nil
}

func (a *articleDB) GetAllArticles() ([]model.Article, error) {
	// Implement the logic to get all articles from the database
	// Example:
	// var articles []Article
	// err := a.db.Find(&articles).Error
	// return articles, err
	return nil, nil
}

func (a *articleDB) CreateArticle(article model.Article) (model.Article, error) {
	// Implement the logic to create an article in the database
	// Example:
	// err := a.db.Create(&article).Error
	// return article, err
	return model.Article{}, nil
}

func (a *articleDB) UpdateArticle(id string, article model.Article) (model.Article, error) {
	// Implement the logic to update an article in the database
	// Example:
	// err := a.db.Model(&Article{}).Where("id = ?", id).Updates(article).Error
	// return article, err
	return model.Article{}, nil
}

func (a *articleDB) DeleteArticle(id string) error {
	// Implement the logic to delete an article from the database
	// Example:
	// err := a.db.Delete(&Article{}, id).Error
	// return err
	return nil
}
