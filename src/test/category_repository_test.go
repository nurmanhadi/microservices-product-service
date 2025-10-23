package test

import (
	"encoding/json"
	"fmt"
	"product-service/pkg/env"
	"product-service/src/config"
	"product-service/src/dto"
	"product-service/src/internal/repository"
	"testing"
)

func TestFindAll(t *testing.T) {
	env.NewEnv()
	db := config.NewSql()
	defer db.Close()
	catRepo := repository.NewCategoryRepository(db)
	categories, err := catRepo.FindAll()
	if err != nil {
		t.Error(err)
	}
	var cat []dto.CategoryResponse
	for _, x := range categories {
		if x.ParentID == nil {
			var subCat []dto.CategoryResponse
			for _, y := range categories {
				if y.ParentID != nil && *y.ParentID == x.ID {
					subCat = append(subCat, dto.CategoryResponse{
						ID:   y.ID,
						Name: y.Name,
					})
				}
			}
			cat = append(cat, dto.CategoryResponse{
				ID:            x.ID,
				Name:          x.Name,
				SubCategories: subCat,
			})
		}
	}
	json, _ := json.Marshal(cat)
	fmt.Println(string(json))
}
