package vespa

type FeedbackCommentAdd struct {
	Fields FeedbackComment `json:"fields"`
}

type FeedbackCommentUpdateFields struct {
	Text StringFieldUpdate `json:"text"`
}

type FeedbackCommentUpdate struct {
	Fields FeedbackCommentUpdateFields `json:"fields"`
}

type FeedbackAdd struct {
	Fields Feedback `json:"fields"`
}

type Feedback struct {
	AgeSuitability int64  `json:"ageSuitability"`
	Contents       string `json:"contents"` // `json:"contents,omitempty"`
	// Create datetime decimal bitmap
	CreateDtb      int64  `json:"createDtb"`
	FeedbackTypeId int32  `json:"themeId"`
	Title          string `json:"title"`
	UserId         int64  `json:"userId"`
}

type FeedbackComment struct {
	AgeSuitability int64 `json:"ageSuitability"`
	// Create datetime decimal bitmap
	CreateDtb      int64  `json:createDtb`
	FeedbackTypeId int32  `json:"themeId"`
	FeedbackId     int64  `json:"pollId"`
	Text           string `json:"text"`
	UserId         int64  `json:"userId"`
}

type GetFeedbackResponse struct {
	Errors []Error `json:"errors"`
	// TODO: test
	Record Feedback
}
