package respJson

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type failJSONResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
}

func WriteJSON(success bool, code int, msg string, data interface{}, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if success {
		if data != nil {
			jsonResp, _ := json.Marshal(data)
			w.Write(jsonResp)
		}
	} else {
		resp := &failJSONResponse{
			Message:   fmt.Sprintf("%v", msg),
			ErrorCode: code,
		}
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
		log.Printf("[%v]%v, %v\n", code, msg, err)
	}
}
