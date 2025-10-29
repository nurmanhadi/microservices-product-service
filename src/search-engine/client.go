package searchengine

import (
	"context"
	"log"
	"product-service/src/search-engine/collection"

	"github.com/typesense/typesense-go/v3/typesense"
	"github.com/typesense/typesense-go/v3/typesense/api"
)

type ClientSearchEngine interface {
	UpsertProduct(product collection.ProductCollection) error
}

type clientSearchEngine struct {
	ctx          context.Context
	searchEngine *typesense.Client
}

func NewClientSearchEngine(searchEngine *typesense.Client) ClientSearchEngine {
	newCtx := context.Background()
	return &clientSearchEngine{
		searchEngine: searchEngine,
		ctx:          newCtx,
	}
}
func (c *clientSearchEngine) UpsertProduct(product collection.ProductCollection) error {
	newProduct := collection.ProductCollection{
		ID:         product.ID,
		Name:       product.Name,
		Quantity:   product.Quantity,
		Price:      product.Price,
		CreatedAt:  product.CreatedAt,
		UpdatedAt:  product.UpdatedAt,
		Categories: product.Categories,
	}
	resp, err := c.searchEngine.Collection("products").Documents().Upsert(c.ctx, newProduct, &api.DocumentIndexParameters{})
	if err != nil {
		return err
	}
	log.Println(resp)
	return nil
}
