package rest

import (
	"fmt"
	"github.com/TomKKlalala/superchainer/util"
	"github.com/TomKKlalala/superchainer/x/lotteryservice/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	"net/http"
)

var logger = util.GetLogger("lotteryservice")

type createLotteryReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Rounds      string       `json:"rounds"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Hashed      bool         `json:"hashed"`
}

func createLotteryHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createLotteryReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rounds := util.StringToArray(req.Rounds, ",")

		// create the message
		msg := types.NewMsgCreateLottery(req.Title, req.Description, addr, rounds, req.Hashed)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

func getLotteryHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		lotteryID := vars[lotteryID]
		logger.Info("receive lotteryID from rest request: " + lotteryID)
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/lottery/%s", storeName, lotteryID), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getLotteriesHandler(cliCtx context.CLIContext, storeName string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/lotteries", storeName), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getCandidatesHandler(cliCtx context.CLIContext, storeName string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		lotteryID := vars[lotteryID]
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/lottery/%s/candidates", storeName, lotteryID), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getWinnersHandler(cliCtx context.CLIContext, storeName string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		lotteryID := vars[lotteryID]
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/lottery/%s/winners", storeName, lotteryID), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
