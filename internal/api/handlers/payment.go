package handlers

import (
	"encoding/json"
	"net/http"

	"simple_bank/internal/usecase"
)

type PaymentHandler struct {
	paymentUseCase *usecase.PaymentUseCase
}

func NewPaymentHandler(paymentUseCase *usecase.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{paymentUseCase: paymentUseCase}
}

func (h *PaymentHandler) MakePayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		MerchantID string  `json:"merchant_id"`
		Amount     float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	customerID := r.Header.Get("CustomerID")
	if customerID == "" {
		http.Error(w, "Customer ID not found", http.StatusBadRequest)
		return
	}

	err := h.paymentUseCase.MakePayment(customerID, req.MerchantID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment successful"})
}
