package handlers

import (
	"encoding/json"
	"ledger-service/model"
	"ledger-service/store"
	"net/http"
)

type TransactionHandler struct {
	Store *store.TransactionStore
}

func NewTransactionHandler(store *store.TransactionStore) *TransactionHandler {
	return &TransactionHandler{Store: store}
}

func (h *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var t model.Transaction
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if t.AccountID == "" {
		http.Error(w, "account_id is required", http.StatusBadRequest)
		return
	}

	h.Store.Add(t)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
