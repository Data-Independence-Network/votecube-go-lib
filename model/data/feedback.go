package data

type Feedback struct {
	AgeSuitability int64  `json:"ageSuitability"`
	Contents       string `json:"contents"` // `json:"contents,omitempty"`
	// Create datetime decimal bitmap
	CreateDtb      int64  `json:"createDtb"`
	FeedbackTypeId int32  `json:"feedbackTypeId"`
	Id             int64  `json:"id"`
	Title          string `json:"title"`
	Text           string `json:"text"`
	UserId         int64  `json:"userId"`
}

type FeedbackComment struct {
	AgeSuitability int64 `json:"ageSuitability"`
	// Create datetime decimal bitmap
	CreateDtb      int64  `json:createDtb`
	FeedbackId     int64  `json:"feedbackId"`
	FeedbackTypeId int32  `json:"feedbackTypeId"`
	Id             int64  `json:"id"`
	Text           string `json:"text"`
	UserId         int64  `json:"userId"`
}
