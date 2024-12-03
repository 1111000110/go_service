package memberapi

type GetMemberRequest struct {
	Mid int64 `json:"mid"`
}
type GetMemberResponse struct {
	Username string `json:"username"`
	Mid      int64  `json:"mid"`
	Name     string `json:"name"`
}
