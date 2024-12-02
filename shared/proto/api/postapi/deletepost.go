package postapi

type DeletePostByPidRequest struct {
	Pid int64 `json:"pid"`
}
type DeletePostByPidResponse struct {
	IsDeleted bool `json:"is_deleted"` // 是否删除
}
