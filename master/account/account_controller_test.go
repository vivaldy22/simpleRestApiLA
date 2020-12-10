package account

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vivaldy22/simpleRestApiLA/models"

	"github.com/vivaldy22/simpleRestApiLA/tools/dbTest"

	"github.com/magiconair/properties/assert"

	"github.com/gorilla/mux"
)

type BadTransfer struct {
	Amountt int `json:"amount"`
}

var mockTrans = &models.Transfer{
	ToAccountNumber: "555112",
	Amount:          "100",
}

var badMockTrans = &BadTransfer{
	Amountt: 100,
}

func mockRouter() *mux.Router {
	db := dbTest.NewDbTest()
	r := mux.NewRouter()
	repo := NewRepo(db)
	uc := NewUseCase(repo)
	ctrl := &accController{c: uc}
	accPref := r.PathPrefix("/account").Subrouter()
	accPref.HandleFunc("/{account_number}", ctrl.GetAccByAccNum).Methods(http.MethodGet)
	accPref.HandleFunc("/{from_account_number}/transfer", ctrl.TransferBalance).Methods(http.MethodPost)

	return r
}

func TestAccController_GetAccByAccNum(t *testing.T) {
	router := mockRouter()
	req, err := http.NewRequest("GET", "/account/555111", nil)
	if err != nil {
		t.Fatalf("error occurred %v", err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, 200, response.Code, "Response 200 is expected")
}

func TestAccController_GetAccByAccNumError(t *testing.T) {
	router := mockRouter()
	req, err := http.NewRequest("GET", "/account/5551113", nil)
	if err != nil {
		t.Fatalf("error occurred %v", err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, 404, response.Code, "Response 404 is expected")
}

func TestAccController_TransferBalance(t *testing.T) {
	router := mockRouter()
	data, _ := json.Marshal(mockTrans)
	req, err := http.NewRequest("POST", "/account/555111/transfer", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("error occurred %v", err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, 201, response.Code, "Response 201 is expected")
}

func TestAccController_TransferBalanceDecodeFailed(t *testing.T) {
	router := mockRouter()
	data, _ := json.Marshal(badMockTrans)
	req, err := http.NewRequest("POST", "/account/555111/transfer", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("error occurred %v", err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, 400, response.Code, "Response 400 is expected")
}

func TestAccController_TransferBalanceNotFound(t *testing.T) {
	router := mockRouter()
	mockTrans.Amount = "999999"
	data, _ := json.Marshal(mockTrans)
	req, err := http.NewRequest("POST", "/account/555111/transfer", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("error occurred %v", err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, 404, response.Code, "Response 404 is expected")
}
