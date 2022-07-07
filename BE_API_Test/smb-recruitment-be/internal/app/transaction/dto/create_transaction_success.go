package dto

import "net/http"

type CreateTransactionSuccess struct {
	TransactionID string `json:"transaction_id"`
}

func (resp *CreateTransactionSuccess) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusCreated)
	return nil
}
