package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fathimtiaz/galaxy-merchant/pkg/service/galaxymerchant"
)

type Request struct {
	Input string
}

type Response struct {
	Result []string
	Error  string
}

func Result(w http.ResponseWriter, r *http.Request) {
	var req Request
	var resp Response
	var err error

	defer func() {
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(jsonResp)
	}()

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp.Error = err.Error()
		return
	}

	gm := galaxymerchant.NewGalaxyMerchant()

	if err = gm.ParseInput(req.Input); err != nil {
		w.WriteHeader(http.StatusOK)
		resp.Error = err.Error()
		return
	}

	if err = gm.SetResults(); err != nil {
		w.WriteHeader(http.StatusOK)
		resp.Error = err.Error()
		return
	}

	resp.Result = gm.Results
}
