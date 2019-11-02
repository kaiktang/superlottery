package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

const (
	restName = "lottery"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/lottery", storeName), createLotteryHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/lottery/{lotteryID}"), getLotteryHandler(cliCtx, "lotteryID")).Methods("GET")
	//r.HandleFunc(fmt.Sprintf("/%s/lottery/{%s}/whois", storeName, restName), whoIsHandler(cliCtx, storeName)).Methods("GET")
	//r.HandleFunc(fmt.Sprintf("/%s/lottery", storeName), deleteNameHandler(cliCtx)).Methods("DELETE")
}
