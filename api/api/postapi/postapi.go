package postapi

import "github.com/1111000110/go_service/api/apidata/postapidata"

type GetPostParam struct {
	Pid int64 `json:"pid"`
}
type GetPostResp struct {
	Data *postapidata.Post `json:"data"`
}
type AddPostParam struct {
	Mid  int64  `json:"mid"`
	Text string `json:"text"`
	Tid  int64  `json:"tid"`
}
type AddPostResp struct {
	Data *postapidata.Post `json:"data"`
}
type DeletePostParam struct {
	Mid int64 `json:"mid"`
	Pid int64 `json:"pid"`
}
type DeletePostResp struct {
	Data bool `json:"data"`
}
