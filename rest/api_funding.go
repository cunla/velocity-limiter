/*
 * Controller for /fund api
 *
 * API version: 1.0.0
 * Contact: daniel.maruani@gmail.com
 */
package rest

import (
	"encoding/json"
	_ "encoding/json"
	logic "github.com/cunla/velocity-limiter/logic"
	model "github.com/cunla/velocity-limiter/model"
	"net/http"
	"strconv"
	"time"
)

func fundRequestToFund(fundRequest FundRequest) model.Fund {
	var fund model.Fund
	fund.FundId, _ = strconv.Atoi(fundRequest.Id)
	fund.CustomerId, _ = strconv.Atoi(fundRequest.CustomerId)
	fund.LoadAmount, _ = strconv.ParseFloat(fundRequest.LoadAmount[1:], 64)
	fund.Time = fundRequest.Time.Truncate(24 * time.Hour)
	return fund
}

func FundAccountApi(w http.ResponseWriter, r *http.Request) {
	var fundRequest FundRequest
	err := json.NewDecoder(r.Body).Decode(&fundRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fund := fundRequestToFund(fundRequest)
	accepted, err := logic.FundAccount(fund)
	if err != nil {
		http.Error(w, "", http.StatusAlreadyReported)
		return
	}
	fundResponse := FundResponse{Id: fundRequest.Id,
		CustomerId: fundRequest.CustomerId,
		Accepted:   accepted}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	bytes, err := json.Marshal(fundResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}
