package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"github.com/skadoodle1201/goAPI/api"
	"github.com/skadoodle1201/goAPI/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.GetCoinRequest{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var response = api.GetCoinResponse{
		Coins:    (*tokenDetails).Coins,
		Code:     http.StatusOK,
		Message:  "Success",
		Username: (*tokenDetails).Username,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
