package publishing

type ArticleRepository interface {
	GetAll() ([]*Article, error)
	GetByID(id string) (*Article, error)
	Create(article *Article) error
	Update(article *Article) (string, error)
	Delete(id string) (string, error)
}
