package types

import (
	"fmt"
	"github.com/TomKKlalala/superchainer/util"
	"strconv"
	"strings"
	"sync"
)
import sdk "github.com/cosmos/cosmos-sdk/types"

var lotteryID int64
var lock sync.Mutex

func LotteryID() string {
	lock.Lock()
	defer lock.Unlock()
	lotteryID++
	return strconv.FormatInt(lotteryID, 10)
}

type Lottery struct {
	Rounds       []int          `json:"rounds"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Owner        sdk.AccAddress `json:"owner"`
	Hashed       bool           `json:"hashed"`
	StopEnroll   bool           `json:"stopEnroll"`
	CurrentRound int            `json:"currentRound"`
}

func NewLottery() Lottery {
	return Lottery{Rounds: []int{}, CurrentRound: -1}
}

func (lottery Lottery) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Title: %s
Description: %s
Owner: %s
Rounds: %s
Hashed: %s
StopEnroll: %s
CurrentRound: %d`,
		lottery.Title,
		lottery.Description,
		lottery.Owner,
		util.ArrayToString(lottery.Rounds, ","),
		strconv.FormatBool(lottery.Hashed),
		strconv.FormatBool(lottery.StopEnroll),
		lottery.CurrentRound))
}
// TODO: 还需要表示每轮抽中人数
type Candidates []string

func (c Candidates) String() string {
	return strings.Join(c, ",")
}
