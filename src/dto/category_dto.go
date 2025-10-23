package dto

type CategoryAddRequest struct {
	Name string `json:"name" validate:"required,max=50"`
}
type CategoryAddSubRequest struct {
	Name       string `json:"name" validate:"required,max=50"`
	CategoryID int    `json:"category_id" validate:"required"`
}
type CategoryUpdateSubRequest struct {
	Name string `json:"name" validate:"required,max=50"`
}
type CategoryResponse struct {
	ID            int                `json:"id"`
	Name          string             `json:"name"`
	SubCategories []CategoryResponse `json:"sub_categories"`
}
