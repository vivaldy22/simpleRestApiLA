package account

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/simpleRestApiLA/models"
	"github.com/vivaldy22/simpleRestApiLA/tools/respJson"
	"github.com/vivaldy22/simpleRestApiLA/tools/varMux"
)

type accController struct {
	accUseCase models.AccountUseCase
}

func NewController(accUseCase models.AccountUseCase, r *mux.Router) {
	handler := &accController{accUseCase}
	pref := r.PathPrefix("/account").Subrouter()
	pref.HandleFunc("/{account_number}", handler.getAccByAccNum).Methods(http.MethodGet)
	pref.HandleFunc("/{from_account_number}/transfer", handler.transferBalance).Methods(http.MethodPost)
}

func (a *accController) getAccByAccNum(w http.ResponseWriter, r *http.Request) {
	accNum := varMux.GetVarsMux("account_number", r)

	data, err := a.accUseCase.GetByAccNum(accNum)
	if err != nil {
		respJson.WriteJSON(false, http.StatusNotFound, "Data not found", nil, err, w)
		return
	}
	respJson.WriteJSON(true, http.StatusOK, "Data found", data, nil, w)
}

func (a *accController) transferBalance(w http.ResponseWriter, r *http.Request) {
	var transferInfo = new(models.Transfer)
	from := varMux.GetVarsMux("from_account_number", r)

	err := json.NewDecoder(r.Body).Decode(transferInfo)
	if err != nil {
		respJson.WriteJSON(false, http.StatusBadRequest, "Decoding json failed", nil, err, w)
		return
	}

	msg, err := a.accUseCase.Transfer(from, transferInfo.ToAccountNumber, transferInfo.Amount)
	if err != nil {
		respJson.WriteJSON(false, http.StatusNotFound, msg, nil, err, w)
		return
	}

	respJson.WriteJSON(true, http.StatusCreated, msg, nil, nil, w)
}
