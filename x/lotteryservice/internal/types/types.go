package types

import (
	"fmt"
	"github.com/TomKKlalala/superchainer/util"
	"strconv"
	"strings"
)
import sdk "github.com/cosmos/cosmos-sdk/types"

type Lottery struct {
	ID           string         `json:"id"`
	Rounds       []int          `json:"rounds"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Owner        sdk.AccAddress `json:"owner"`
	Hashed       bool           `json:"hashed"`
	StopEnroll   bool           `json:"stopEnroll"`
	CurrentRound int            `json:"currentRound"`
	CandidateNum int            `json:"candidateNum"`
	// for seeds
	CreateTime int64 `json:"createTime"`
}

func NewLottery() Lottery {
	return Lottery{Rounds: []int{}, CurrentRound: -1}
}

func (lottery Lottery) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
ID: %s
Title: %s
Description: %s
Owner: %s
Rounds: %s
Hashed: %s
StopEnroll: %s
CurrentRound: %d
CandidateNum: %d
CreateTime: %d`,
		lottery.ID,
		lottery.Title,
		lottery.Description,
		lottery.Owner,
		util.ArrayToString(lottery.Rounds, ","),
		strconv.FormatBool(lottery.Hashed),
		strconv.FormatBool(lottery.StopEnroll),
		lottery.CurrentRound,
		lottery.CandidateNum,
		lottery.CreateTime),
	)
}

type Candidates []string

func (c Candidates) String() string {
	return strings.Join(c, ",")
}
