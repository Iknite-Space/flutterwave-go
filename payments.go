package flutterwave

// GetPaymentLinkRequest is data needed to create a payment link
type GetPaymentLinkRequest struct {
	TransactionReference string                       `json:"tx_ref"`
	Amount               string                       `json:"amount"`
	Currency             string                       `json:"currency"`
	Meta                 map[string]interface{}       `json:"meta"`
	RedirectURL          string                       `json:"redirect_url"`
	Customer             GetPaymentLinkCustomer       `json:"customer"`
	Customizations       GetPaymentLinkCustomizations `json:"customizations"`
}

// GetPaymentLinkCustomer contains the customer details.
type GetPaymentLinkCustomer struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phonenumber"`
}

// GetPaymentLinkCustomizations contains options to customize the look of the payment modal.
type GetPaymentLinkCustomizations struct {
	Title string `json:"title"`
	Logo  string `json:"logo"`
}

// GetPaymentLinkResponse is the data returned after creating a payment link
type GetPaymentLinkResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Link string `json:"link"`
	} `json:"data"`
}

// InitiateMobileMoneyFrancoRequest is the data needed to create a payment request
type InitiateMobileMoneyFrancoRequest struct {
	// Amount is the amount to be charged it is passed as - ('amount':'100') (required).
	// // N.B. Amount should not be less than 100."
	Amount string `json:"amount"`

	// Currency This is the specified currency to charge in. (expected value: XAF or XOF) (required).
	Currency string `json:"currency"`

	// PhoneNumber is the phone number linked to the customer's mobile money account (required).
	PhoneNumber string `json:"phone_number"`

	// Email is the email address of the customer (required)
	Email string `json:"email"`

	// TransactionRef is a unique reference, unique to the particular transaction being carried (required).
	TransactionRef string `json:"tx_ref"`

	// Country is the country code of the francophone country making the mobile money payment. Possible values are CM
	// // (Cameroon), SN (Senegal), BF (Burkina Faso) and so on.
	Country string `json:"country"`

	// FullName is the customers full name. It should include first and last name of the customer
	FullName string `json:"fullname,omitempty"`

	// ClientIP represents the current IP address of the customer carrying out the transaction
	ClientIP string `json:"client_ip,omitempty"`

	// DeviceFingerprint is the fingerprint for the device being used. It can be generated using a library on whatever
	// //platform is being used.
	DeviceFingerprint string `json:"device_fingerprint,omitempty"`

	// Meta is used to include additional payment information.
	Meta map[string]interface{} `json:"meta,omitempty"`
}

// InitiateMobileMoneyFrancoResponse represents the response from a mobile money payment request.
type InitiateMobileMoneyFrancoResponse struct {
	Status  string `json:"status"`  // Status of the charge request
	Message string `json:"message"` // Message related to the charge request
	Data    struct {
		ID                int    `json:"id"`
		TxRef             string `json:"tx_ref"`
		FlwRef            string `json:"flw_ref"`
		OrderID           string `json:"order_id"`
		DeviceFingerprint string `json:"device_fingerprint"`
		Amount            uint32 `json:"amount"`
		ChargedAmount     uint32 `json:"charged_amount"`
		AppFee            uint32 `json:"app_fee"`
		MerchantFee       int    `json:"merchant_fee"`
		AuthModel         string `json:"auth_model"`
		Currency          string `json:"currency"`
		IP                string `json:"ip"`
		Narration         string `json:"narration"`
		Status            string `json:"status"`
		PaymentType       string `json:"payment_type"`
		FraudStatus       string `json:"fraud_status"`
		ChargeType        string `json:"charge_type"`
		CreatedAt         string `json:"created_at"`
		AccountID         int    `json:"account_id"`
		Customer          struct {
			ID          int    `json:"id"`
			PhoneNumber string `json:"phone_number"`
			Name        string `json:"name"`
			Email       string `json:"email"`
			CreatedAt   string `json:"created_at"`
		} `json:"customer"`
		ProcessorResponse string `json:"processor_response"`
	} `json:"data"`
	Meta map[string]interface{} `json:"meta"`
}
