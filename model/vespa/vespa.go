package vespa

type Error struct {
	Description string `json:"description"`
	Id          int    `json:"id"`
}

type ModResponse struct {
	Errors []Error `json:"errors"`
	Id     string  `json:"id"`
	Path   string  `json:"path"`
}

type OpinionAdd struct {
	Fields Opinion `json:"fields"`
}

type StringFieldUpdate struct {
	Assign string `json:"assign"`
}

type OpinionUpdateFields struct {
	Text StringFieldUpdate `json:"text"`
}

type OpinionUpdate struct {
	Fields OpinionUpdateFields `json:"fields"`
}

type PollAdd struct {
	Fields Poll `json:"fields"`
}

type Opinion struct {
	AgeSuitability int64 `json:"ageSuitability"`
	// Create datetime decimal bitmap
	CreateDtb     int64  `json:createDtb`
	LocationId    int32  `json:"locationId"`
	PollId        int64  `json:"pollId"`
	RootOpinionId int64  `json:rootOpinionId,omitempty`
	Text          string `json:"text"`
	ThemeId       int64  `json:"themeId"`
	UserId        int64  `json:"userId"`
}

type Poll struct {
	AgeSuitability int64  `json:"ageSuitability"`
	Contents       string `json:"contents"` // `json:"contents,omitempty"`
	// Create datetime decimal bitmap
	CreateDtb  int64  `json:"createDtb"`
	LocationId int32  `json:"locationId"`
	ThemeId    int64  `json:"themeId"`
	Title      string `json:"title"`
	UserId     int64  `json:"userId"`
}
