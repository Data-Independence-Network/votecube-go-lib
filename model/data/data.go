package data

type Opinion struct {
	CreateEs        int64  `json:"createEs"`
	Id              int64  `json:"id"`
	RootOpinionId   int64  `json:rootOpinionId,omitempty`
	ParentOpinionId int64  `json:parentOpinionId,omitempty`
	PollId          int64  `json:"pollId"`
	Text            string `json:"text"`
	UserId          int64  `json:"userId"`
}

type Poll struct {
	Contents string `json:"contents"` // `json:"contents,omitempty"`
	CreateEs int64  `json:"createEs"`
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	UserId   int64  `json:"userId"`
}
