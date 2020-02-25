package scylladb

type Row interface {
	TableName() string
}

type Opinion struct {
	PartitionPeriod int32
	RootOpinionId   int64
	OpinionId       int64
	//AgeSuitability  int64
	//PollId          int64
	//ThemeId         int64
	//LocationId      int32
	Version int16
	//ParentOpinionId int64
	//CreateDbt        int64
	//UserId          int64
	Data            []byte
	InsertProcessed bool
}

func (r Opinion) TableName() string {
	return "opinions"
}

type OpinionUpdate struct {
	PartitionPeriod int32
	RootOpinionId   int64
	OpinionId       int64
	Version         int16
	UpdateProcessed bool
}

func (r OpinionUpdate) TableName() string {
	return "opinion_updates"
}

type PeriodAddedToRootOpinionIds struct {
	PartitionPeriod  int64
	RootOpinionId    int64
	RootOpinionIdMod int32
}

func (r PeriodAddedToRootOpinionIds) TableName() string {
	return "period_added_to_root_opinion_ids"
}

type PeriodUpdatedRootOpinionIds struct {
	PartitionPeriod  int64
	RootOpinionId    int64
	RootOpinionIdMod int32
}

func (r PeriodUpdatedRootOpinionIds) TableName() string {
	return "period_updated_root_opinion_ids"
}

type Poll struct {
	PollId          int64
	PollIdMod       int32
	PartitionPeriod int32
	//ThemeId         int64
	//LocationId      int32
	//CreateDtb        int64
	//UserId          int64
	//AgeSuitability  int64
	Data            []byte
	InsertProcessed bool
}

func (r Poll) TableName() string {
	return "polls"
}

type RootOpinion struct {
	OpinionId int64
	PollId    int64
	CreateEs  int64
	Version   int32
	Data      []byte
}

func (r RootOpinion) TableName() string {
	return "root_opinions"
}

type UserSession struct {
	PartitionPeriod int32
	SessionId       string
	LastActionEs    int64
	KeepSignedIn    int8
	UserId          int64
	Data            []byte
}

func (r UserSession) TableName() string {
	return "user_sessions"
}
