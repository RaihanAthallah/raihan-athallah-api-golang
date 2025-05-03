package repository

type ArticeRepository interface {
	// GetArticleByID(id string) (Article, error)
	// GetAllArticles() ([]Article, error)
	// CreateArticle(article Article) (Article, error)
	// UpdateArticle(id string, article Article) (Article, error)
	// DeleteArticle(id string) error
}

type articleRepository struct {
	cache ArticleCache
	db    ArticleDB
}

func NewArticleRepository() ArticeRepository {
	return &articleRepository{}
}
