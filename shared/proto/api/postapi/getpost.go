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

func (api *PostApi) GetPostsByPids(req *GetPostsByPidsRequest) (*GetPostsByPidsResponse, error) {
	resp := &GetPostsByPidsResponse{}
	err := api.Post("/post/getpostsbypids", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *PostApi) GetPostsByMids(req *GetPostsByMidsRequest) (*GetPostsByMidsResponse, error) {
	resp := &GetPostsByMidsResponse{}
	err := api.Post("/post/getpostsbymids", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
