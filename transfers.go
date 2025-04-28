package flutterwave

import "encoding/json"

// ExchangeRateResponse is data returned when querying a transaction.
type TransferRateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Rate        float64 `json:"rate"`
		Source      Payment `json:"source"`
		Destination Payment `json:"destination"`
	}
}

// Payment is data returned for either the source or the destination.
type Payment struct {
	Currency string      `json:"currency"`
	Amount   json.Number `json:"amount"`
}

// CreateTransferRequest represents the payload for creating a transfer.
type CreateTransferRequest struct {
	AccountBank           string `json:"account_bank"`
	AccountNumber         string `json:"account_number"`
	Amount                string `json:"amount"`
	Currency              string `json:"currency"`
	DebitSubaccount       string `json:"debit_subaccount"`
	Beneficiary           string `json:"beneficiary"`
	BeneficiaryName       string `json:"beneficiary_name"`
	Reference             string `json:"reference"`
	DebitCurrency         string `json:"debit_currency"`
	DestinationBranchCode string `json:"destination_branch_code"`
	CallbackURL           string `json:"callback_url"`
	Narration             string `json:"narration"`
}

// CreateTransferResponse represents the response for creating a transfer.
type CreateTransferResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID               json.Number `json:"id"`
		AccountNumber    string      `json:"account_number"`
		BankCode         string      `json:"bank_code"`
		FullName         string      `json:"full_name"`
		CreatedAt        string      `json:"created_at"`
		Currency         string      `json:"currency"`
		DebitCurrency    string      `json:"debit_currency"`
		Amount           json.Number `json:"amount"`
		Fee              json.Number `json:"fee"`
		Status           string      `json:"status"`
		Reference        string      `json:"reference"`
		Meta             interface{} `json:"meta"`
		Narration        string      `json:"narration"`
		CompleteMessage  string      `json:"complete_message"`
		RequiresApproval int         `json:"requires_approval"`
		IsApproved       int         `json:"is_approved"`
		BankName         string      `json:"bank_name"`
	} `json:"data"`
}
