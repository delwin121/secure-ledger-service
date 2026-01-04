package store

import (
	"ledger-service/model"
	"sync"
)

type TransactionStore struct {
	mu           sync.Mutex
	transactions []model.Transaction
}

func NewTransactionStore() *TransactionStore {
	return &TransactionStore{
		transactions: make([]model.Transaction, 0),
	}
}

func (s *TransactionStore) Add(t model.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.transactions = append(s.transactions, t)
}

func (s *TransactionStore) GetAll() []model.Transaction {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Return a copy to avoid race conditions if caller modifies it
	result := make([]model.Transaction, len(s.transactions))
	copy(result, s.transactions)
	return result
}
