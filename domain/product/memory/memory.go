package memory

import (
	"github.com/google/uuid"
	"github.com/shari4ov/ddd-go.git/aggregate"
	"github.com/shari4ov/ddd-go.git/domain/product"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}
func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product
	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
}
func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}
func (mpr *MemoryProductRepository) Update(pr aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[pr.Item().ID]; !ok {
		return product.ErrProductNotFound
	}
	mpr.products[pr.Item().ID] = pr
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
func (mpr *MemoryProductRepository) Add(pr aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[pr.Item().ID]; ok {
		return product.ErrProductExists
	}
	mpr.products[pr.Item().ID] = pr
	return nil
}
