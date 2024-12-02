package models

type Post struct {
	Tid   int64  `json:"tid"`
	Mid   int64  `json:"mid"`
	Pid   int64  `json:"pid"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Ct    int64  `json:"ct"`
	Ut    int64  `json:"ut"`
}
