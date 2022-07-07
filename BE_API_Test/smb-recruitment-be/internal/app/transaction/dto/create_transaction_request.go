package dto

import (
	"encoding/json"
	"math/big"
	"net/http"
)

type CreateTransactionRequest struct {
	TransactionCode     string     `json:"transaction_code"`
	Amount              *big.Float `json:"amount"`
	DestinationAccount  string     `json:"destination_account"`
	AuthorizationMethod string     `json:"auth_method"`
}

func (payload *CreateTransactionRequest) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
		return err
	}
	return nil
}
