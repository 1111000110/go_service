package postapi

type DeletePostByPidRequest struct {
	Mid int64 `json:"mid"`
	Pid int64 `json:"pid"`
}
type DeletePostByPidResponse struct {
	IsDeleted bool `json:"is_deleted"` // 是否删除
}
