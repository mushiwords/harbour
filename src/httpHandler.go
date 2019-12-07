package main
import (
	"net/http"
	"encoding/json"
)

type HarbourResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}
/**
 * 404 返回函数
 **/
func MyNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["requestId"] = "requestId"
	data["result"] = &HarbourResponse{
		Code: 0,
		Message: "some message.",
	}

	w.WriteHeader(http.StatusNotFound)
	if data == nil {
		return 
	}

	body, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Write(body)
}

func ReturnResult(w http.ResponseWriter, r *http.Request, code int) {
	data := make(map[string]interface{})
	data["requestId"] = "requestId"
	data["result"] = &HarbourResponse{
		Code: code,
		Message: "some message.",
	}

	w.WriteHeader(http.StatusOK)
	
	if data == nil {
		return 
	}

	body, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Write(body)
}
