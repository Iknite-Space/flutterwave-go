package flutterwave

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/NdoleStudio/flutterwave-go/internal/helpers"
	"github.com/NdoleStudio/flutterwave-go/internal/stubs"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBillsService_CreatePayment(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillsCreateDStvPaymentResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	data, response, err := client.Bills.CreatePayment(context.Background(), &BillsCreatePaymentRequest{
		Country:    "NG",
		Customer:   "7034504232",
		Amount:     100,
		Recurrence: "ONCE",
		Type:       "DStv",
		Reference:  uuid.New().String(),
		BillerName: "DStv",
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "success", data.Status)
	assert.Equal(t, "Bill payment successful", data.Message)
	assert.Equal(t, "+23490803840303", data.Data.PhoneNumber)
	assert.Equal(t, "500", data.Data.Amount.String())
	assert.Equal(t, "9MOBILE", data.Data.Network)
	assert.Equal(t, "CF-FLYAPI-20200311081921359990", data.Data.FlwRef)
	assert.Equal(t, "BPUSSD1583957963415840", data.Data.TxRef)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.True(t, data.IsSuccessfull())

	// Teardown
	server.Close()
}

func TestBillsService_Validate(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillsValidateDstvResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	data, response, err := client.Bills.Validate(context.Background(), "CB177", "BIL099", "08038291822")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "success", data.Status)
	assert.Equal(t, "Item validated successfully", data.Message)
	assert.Equal(t, "00", data.Data.ResponseCode)
	assert.Equal(t, "Successful", data.Data.ResponseMessage)
	assert.Equal(t, "MTN", data.Data.Name)
	assert.Equal(t, "BIL099", data.Data.BillerCode)
	assert.Equal(t, "08038291822", data.Data.Customer)
	assert.Equal(t, "AT099", data.Data.ProductCode)
	assert.Equal(t, "100", data.Data.Fee.String())
	assert.Equal(t, "0", data.Data.Maximum.String())
	assert.Equal(t, "0", data.Data.Minimum.String())
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.True(t, data.IsSuccessfull())

	// Teardown
	server.Close()
}

func TestBillsService_GetStatusVerbose(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.BillsGetStatusVerboseResponse())
	client := New(WithBaseURL(server.URL))

	// Act
	data, response, err := client.Bills.GetStatusVerbose(context.Background(), "9300049404444")

	// Assert
	assert.Nil(t, err)

	transactionDate, err := dateparse.ParseAny("2020-03-11T20:19:21.27Z")
	assert.Nil(t, err)

	assert.Equal(t, &BillsStatusVerboseResponse{
		Status:  "success",
		Message: "Bill status fetch successful",
		Data: struct {
			Currency        string      `json:"currency"`
			CustomerID      string      `json:"customer_id"`
			Frequency       string      `json:"frequency"`
			Amount          string      `json:"amount"`
			Product         string      `json:"product"`
			ProductName     string      `json:"product_name"`
			Commission      int         `json:"commission"`
			TransactionDate time.Time   `json:"transaction_date"`
			Country         string      `json:"country"`
			TxRef           string      `json:"tx_ref"`
			Extra           interface{} `json:"extra"`
			ProductDetails  string      `json:"product_details"`
			Status          string      `json:"status"`
		}{
			"NGN",
			"+23490803840303",
			"One Time",
			"500.0000",
			"AIRTIME",
			"9MOBILE",
			10,
			transactionDate,
			"NG",
			"CF-FLYAPI-20200311081921359990",
			nil,
			"FLY-API-NG-AIRTIME-9MOBILE",
			"successful",
		},
	}, data)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)
	assert.True(t, data.IsSuccessfull())

	// Teardown
	server.Close()
}
