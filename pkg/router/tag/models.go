package tag

type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name" validate:"required"`
    Type int    `json:"type" validate:"required"`  // 0: product, 1: vendor
}