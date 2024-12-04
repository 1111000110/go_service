package postapi

type DeletePostByPidRequest struct {
	Mid int64 `json:"mid"`
	Pid int64 `json:"pid"`
}
type DeletePostByPidResponse struct {
	IsDeleted bool `json:"is_deleted"` // 是否删除
}
type DeletePostByMidRequest struct {
	Mid int64 `json:"mid"`
}
type DeletePostByMidResponse struct {
	IsDeleted bool `json:"is_deleted"`
}

func (api *PostApi) DeletePostByPid(req *DeletePostByPidRequest) (*DeletePostByPidResponse, error) {
	resp := &DeletePostByPidResponse{}
	err := api.Post("/post/deletepostbypid", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (api *PostApi) DeletePostByMid(req *DeletePostByMidRequest) (*DeletePostByMidResponse, error) {
	resp := &DeletePostByMidResponse{}
	err := api.Post("/post/deletepostbymid", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
