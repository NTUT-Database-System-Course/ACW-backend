package favorite

type AddRequest struct {
    ProductID int `json:"product_id" validate:"required"`
}

type DeleteRequest struct {
    ProductID int `json:"product_id" validate:"required"`
}

type FavoriteProduct struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       int    `json:"price"`
    Remain      int    `json:"remain"`
    Disability  bool   `json:"disability"`
    ImageURL    string `json:"image_url"`
}