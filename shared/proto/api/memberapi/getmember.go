package memberapi

type GetMemberByMidRequest struct {
	Mid int64 `json:"mid"`
}
type GetMemberByMidResponse struct {
	Username string `json:"username"`
	Mid      int64  `json:"mid"`
	Name     string `json:"name"`
}

func (api *MemberApi) GetMember(req *GetMemberByMidRequest) (*GetMemberByMidResponse, error) {
	resp := &GetMemberByMidResponse{}
	err := api.Post("/member/getmemberbymid", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
