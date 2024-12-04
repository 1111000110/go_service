package memberapi

type DeletememberRequest struct {
	Mid      int64  `json:"mid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
type DeletememberResponse struct {
	IsDeleted bool `json:"is_deleted"`
}

func (api *MemberApi) DeleteMember(req *DeletememberRequest) (*DeletememberResponse, error) {
	resp := &DeletememberResponse{}
	err := api.Post("/member/deletemember", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
