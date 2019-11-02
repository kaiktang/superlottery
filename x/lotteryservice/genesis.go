package lotteryservice

import (
	"fmt"
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	LotteryRecords []types.Lottery `json:"lottery_records"`
}

func NewGenesisState(lotteryRecords []types.Lottery) GenesisState {
	return GenesisState{LotteryRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.LotteryRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid LotteryRecords: Owner: %s. Error: Missing Owner", record.Owner)
		}
		if record.Title == "" {
			return fmt.Errorf("invalid LotteryRecords: Title: %s. Error: Missing Value", record.Owner)
		}
		if record.Description == "" {
			return fmt.Errorf("invalid LotteryRecords: Description: %s. Error: Missing Price", record.Description)
		}

		if len(record.Rounds) < 1 {
			return fmt.Errorf("invalid LotteryRecords: Rounds: %s. Error: Missing Rounds", record.Rounds)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		LotteryRecords: []types.Lottery{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.LotteryRecords {
		keeper.CreateLottery(ctx, &record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Lottery
	iterator := k.GetLotteryIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		lotteryID := string(iterator.Key())
		lottery := k.GetLottery(ctx, lotteryID[len(types.LotteryPrefix):])
		records = append(records, *lottery)

	}
	return GenesisState{LotteryRecords: records}
}
