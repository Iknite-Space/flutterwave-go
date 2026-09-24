package flutterwave

import (
	"context"
	"github.com/NdoleStudio/flutterwave-go/internal/helpers"
	"github.com/NdoleStudio/flutterwave-go/internal/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTransactionsService_VerifyAcceptsDecimalAmount(t *testing.T) {
	t.Parallel()

	const body = `{
  "status": "success",
  "message": "Transaction fetched successfully",
  "data": {
    "id": 987654,
    "tx_ref": "collection-ref",
    "flw_ref": "FLW-MOCK",
    "amount": 34.33,
    "currency": "GBP",
    "charged_amount": 34.33,
    "app_fee": 0.99,
    "merchant_fee": 0,
    "status": "successful",
    "payment_type": "card",
    "created_at": "2026-09-24T08:57:00.000Z",
    "amount_settled": 33.34,
    "customer": {
      "id": 1,
      "name": "Test User",
      "phone_number": "",
      "email": "test@example.com",
      "created_at": "2026-09-24T08:57:00.000Z"
    }
  }
}`

	server := helpers.MakeTestServer(http.StatusOK, body)
	defer server.Close()
	client := New(WithBaseURL(server.URL))

	txn, response, err := client.Transactions.Verify(context.Background(), 987654)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, "34.33", txn.Data.Amount.String())
	assert.Equal(t, "34.33", txn.Data.ChargedAmount.String())
	assert.Equal(t, "GBP", txn.Data.Currency)
	assert.Equal(t, "successful", txn.Data.Status)
}

func TestTransactionsService_VerifyAcceptsIntegerAmount(t *testing.T) {
	t.Parallel()

	const body = `{
  "status": "success",
  "message": "Transaction fetched successfully",
  "data": {
    "id": 123,
    "tx_ref": "momo-ref",
    "amount": 20000,
    "currency": "XAF",
    "charged_amount": 20000,
    "app_fee": 0,
    "merchant_fee": 0,
    "status": "successful",
    "payment_type": "mobilemoney",
    "created_at": "2026-09-24T08:57:00.000Z",
    "customer": {
      "id": 1,
      "name": "Test User",
      "phone_number": "",
      "email": "test@example.com",
      "created_at": "2026-09-24T08:57:00.000Z"
    }
  }
}`

	server := helpers.MakeTestServer(http.StatusOK, body)
	defer server.Close()
	client := New(WithBaseURL(server.URL))

	txn, response, err := client.Transactions.Verify(context.Background(), 123)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, "20000", txn.Data.Amount.String())
}

func TestTransactionsService_Refund(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, string(stubs.TransactionRefundResponse()))
	client := New(WithBaseURL(server.URL))

	// Act
	refund, response, err := client.Transactions.Refund(context.Background(), 123, 200)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.Equal(t, stubs.TransactionRefundResponse(), *response.Body)
	assert.Equal(t, 75923, refund.Data.ID)

	// Teardown
	server.Close()
}

func TestTransactionsService_RefundWithError(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusInternalServerError, "")
	client := New(WithBaseURL(server.URL))

	// Act
	_, response, err := client.Transactions.Refund(context.Background(), 123, 200)

	// Assert
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	// Teardown
	server.Close()
}
