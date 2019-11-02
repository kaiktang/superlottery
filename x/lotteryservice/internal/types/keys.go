package types

const (
	ModuleName = "lotteryservice"

	StoreKey = ModuleName

	// can't share the same prefix
	LotteryPrefix           = "lottery_"
	LotteryIDKey            = "id"
	LotteryCandidatesPrefix = "candidates_"
)
