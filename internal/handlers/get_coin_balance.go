package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/diyabhandari/Go-API/api"
	"github.com/diyabhandari/Go-API/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

// in this func we assume that weve alr passed the auth middleware and we just need to grab the username from the db
// easiest way is to decode our params into a struct
func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query()) //grabs params in url and sets them to the values in the struct
	//here itll grab username from url and put it into username field in struct
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
	var tokenDetails = *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
