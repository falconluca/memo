package memo

import "github.com/shaohsiung/memo/internal/pkg/dbcontext"

type Repository interface {
	Create(item *Item) error
	Update(item *Item) error
	Delete(id int64) error
	Get(id int64) (*Item, error)
	List() ([]*Item, error)
}

type repository struct {
	db *dbcontext.DB
}

func NewRepository(db *dbcontext.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Create(item *Item) error {
	return r.db.Create(item).Error
}

func (r repository) Update(item *Item) error {
	return r.db.Save(item).Error
}

func (r repository) Delete(id int64) error {
	return r.db.Delete(&Item{ID: id}).Error
}

func (r repository) Get(id int64) (item *Item, err error) {
	item = &Item{ID: id}
	err = r.db.First(item).Error
	return
}

func (r repository) List() ([]*Item, error) {
	var items []*Item
	err := r.db.Find(&items).Error
	return items, err
}
