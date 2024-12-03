package postapi

import "github.com/1111000110/go_service/shared/models"

type GetPostsByPidsRequest struct {
	Pids []int64 `json:"pids"`
}
type GetPostsByPidsResponse struct {
	Posts *[]models.Post `json:"posts"`
}
type GetPostsByMidsRequest struct {
	Mids []int64 `json:"mids"`
}
type GetPostsByMidsResponse struct {
	Posts *[]models.Post `json:"posts"`
}
