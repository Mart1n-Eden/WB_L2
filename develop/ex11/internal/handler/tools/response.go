package tools

import (
	"encoding/json"
	"net/http"
)

type error struct {
	Err string `json:"error,omitempty"`
}

//type sucsess struct {
//	Res []model.Event `json:"result,omitempty"`
//}

type sucsess struct {
	Res interface{} `json:"result,omitempty"`
}

func SendError(w http.ResponseWriter, code int, err string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&error{Err: err})
}

//func SendSucsess(w http.ResponseWriter, code int, events []model.Event) {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(code)
//	json.NewEncoder(w).Encode(&sucsess{Res: events})
//}

func SendSucsess(w http.ResponseWriter, code int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&sucsess{Res: res})
}
