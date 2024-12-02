package postapi

import "github.com/1111000110/go_service/shared/models"

type CreatePostRequest struct {
	Mid   int64  `json:"mid"`
	Tid   int64  `json:"tid"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type CreatePostResponse struct {
	Post models.Post `json:"post"`
}
