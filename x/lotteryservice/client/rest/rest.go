package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

const (
	lotteryID = "id"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/lottery", storeName), createLotteryHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/lotteries", storeName), getLotteriesHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/lottery/{%s}", storeName, lotteryID), getLotteryHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/lottery/{%s}/candidates", storeName, lotteryID), getCandidatesHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/lottery/{%s}/winners", storeName, lotteryID), getWinnersHandler(cliCtx, storeName)).Methods("GET")
}
