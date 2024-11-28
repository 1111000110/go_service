package postapidata

type Post struct {
	Pid  int64  `json:"pid"`
	Mid  int64  `json:"mid"`
	Tid  int64  `json:"tid"`
	Ct   int64  `json:"ct"`
	Ut   int64  `json:"ut"`
	Text string `json:"text"`
	Img  string `json:"img"`
}

func (p *Post) GetPid() int64 {
	if p != nil {
		return p.Pid
	}
	return 0
}
func (p *Post) GetMid() int64 {
	if p != nil {
		return p.Mid
	}
	return 0
}
func (p *Post) GetTid() int64 {
	if p != nil {
		return p.Tid
	}
	return 0
}
func (p *Post) GetCt() int64 {
	if p != nil {
		return p.Ct
	}
	return 0
}
func (p *Post) GetUt() int64 {
	if p != nil {
		return p.Ut
	}
	return 0
}
func (p *Post) GetText() string {
	if p != nil {
		return p.Text
	}
	return ""
}
func (p *Post) GetImg() string {
	if p != nil {
		return p.Img
	}
	return ""
}
