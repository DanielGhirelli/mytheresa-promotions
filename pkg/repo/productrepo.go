package repo

import (
	"encoding/json"
	"os"
	"sync"

	"mytheresa-promotions/pkg/domain/product"
)

type fileRepository struct {
	filePath string
	mu       sync.RWMutex
}

func NewFileRepository(filePath string) (product.Repository, error) {
	repo := &fileRepository{
		filePath: filePath,
	}

	// Verify the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, err
	}

	return repo, nil
}

func (r *fileRepository) GetAll() ([]product.RawProduct, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	file, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var data struct {
		Products []product.RawProduct `json:"products"`
	}
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data.Products, nil
}
