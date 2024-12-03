package memberapi

type DeletememberRequest struct {
	Mid      int64  `json:"mid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}
type DeletememberResponse struct {
}
