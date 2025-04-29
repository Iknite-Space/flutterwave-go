package flutterwave

type ListChargebacksResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Meta    struct {
		PageInfo struct {
			Total       int `json:"total"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
			PageSize    int `json:"page_size"`
		} `json:"page_info"`
	} `json:"meta"`
	Data []Chargeback `json:"data"`
}

type Chargeback struct {
	ID      int    `json:"id"`
	Amount  int    `json:"amount"`
	FlwRef  string `json:"flw_ref"`
	Status  string `json:"status"`
	Stage   string `json:"stage"`
	Comment string `json:"comment"`
	Meta    struct {
		UploadedProof interface{} `json:"uploaded_proof"`
		History       []struct {
			Action      string `json:"action"`
			Stage       string `json:"stage,omitempty"`
			Date        string `json:"date"`
			Description string `json:"description"`
			Source      string `json:"source,omitempty"`
		} `json:"history"`
	} `json:"meta"`
	DueDate       string `json:"due_date"`
	SettlementID  string `json:"settlement_id"`
	CreatedAt     string `json:"created_at"`
	TransactionID int    `json:"transaction_id"`
	TxRef         string `json:"tx_ref"`
}
