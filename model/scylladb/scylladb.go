package scylladb

type Opinion struct {
	PartitionPeriod int32
	RootOpinionId   int64
	OpinionId       int64
	AgeSuitability  int64
	PollId          int64
	ThemeId         int64
	LocationId      int32
	Version         int16
	ParentOpinionId int64
	CreateEs        int64
	UserId          int64
	Data            []byte
	InsertProcessed bool
}

type OpinionUpdate struct {
	PartitionPeriod int32
	RootOpinionId   int64
	OpinionId       int64
	Version         int16
	UpdateProcessed bool
}

type PeriodAddedToRootOpinionIds struct {
	PartitionPeriod  int64
	RootOpinionId    int64
	RootOpinionIdMod int32
}

type PeriodUpdatedRootOpinionIds struct {
	PartitionPeriod  int64
	RootOpinionId    int64
	RootOpinionIdMod int32
}

type Poll struct {
	PollId          int64
	ThemeId         int64
	LocationId      int32
	PollIdMod       int32
	CreateEs        int64
	UserId          int64
	PartitionPeriod int32
	AgeSuitability  int64
	Data            []byte
	InsertProcessed bool
}

type RootOpinion struct {
	OpinionId int64
	PollId    int64
	Version   int32
	Data      []byte
}

type UserSession struct {
	PartitionPeriod int32
	SessionId       string
	LastActionEs    int64
	KeepSignedIn    int8
	UserId          int64
	Data            []byte
}
