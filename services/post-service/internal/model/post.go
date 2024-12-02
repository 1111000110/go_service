package model

type Post struct {
	Tid   int64  `json:"tid" ,bson:"tid"`
	Mid   int64  `json:"mid" ,bson:"mid"`
	Pid   int64  `json:"pid" ,bson:"pid"`
	Title string `json:"title" ,bson:"title"`
	Desc  string `json:"desc" ,bson:"desc"`
	Ct    int64  `json:"ct" ,bson:"ct"`
	Ut    int64  `json:"ut" ,bson:"ut"`
}
