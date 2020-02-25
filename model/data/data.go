package data

type Opinion struct {
	AgeSuitability int64 `json:"ageSuitability"`
	// Create datetime decimal bitmap
	CreateDtb       int64  `json:"createDtb"`
	Id              int64  `json:"id"`
	LocationId      int32  `json:"locationId"`
	ParentOpinionId int64  `json:parentOpinionId,omitempty`
	PollId          int64  `json:"pollId"`
	RootOpinionId   int64  `json:rootOpinionId,omitempty`
	Text            string `json:"text"`
	ThemeId         int64  `json:"themeId"`
	UserId          int64  `json:"userId"`
	Version         int16  `json:"version"`
}

type Poll struct {
	AgeSuitability int64  `json:"ageSuitability"`
	Contents       string `json:"contents"` // `json:"contents,omitempty"`
	// Create datetime decimal bitmap
	CreateDtb  int64  `json:"createDtb"`
	Id         int64  `json:"id"`
	LocationId int32  `json:"locationId"`
	ThemeId    int64  `json:"themeId"`
	Title      string `json:"title"`
	UserId     int64  `json:"userId"`
}
