package memberapi

type CreateMemberRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
}
type CreateMemberResponse struct {
	Mid int64 `json:"mid"`
}

func (api *MemberApi) CreateMember(req *CreateMemberRequest) (*CreateMemberResponse, error) {
	resp := &CreateMemberResponse{}
	err := api.Post("/member/createmember", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
