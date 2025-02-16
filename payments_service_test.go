package flutterwave

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/flutterwave-go/internal/helpers"
	"github.com/NdoleStudio/flutterwave-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func Test_paymentsService_InitiateMobileMoneyFranco_success(t *testing.T) {

	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, stubs.InitiateMobileMoneyFrancoResponse())
	defer server.Close()
	client := New(WithBaseURL(server.URL))
	ctx := context.Background()
	transRef := "MC-158523s09v5050e8"

	// Act
	data, response, err := client.Payments.InitiateMobileMoneyFranco(ctx, &InitiateMobileMoneyFrancoRequest{
		Amount:         10_000,
		Currency:       "XAF",
		PhoneNumber:    "+237678787878",
		Email:          "alice@example.com",
		TransactionRef: transRef,
		Country:        "CM",
		FullName:       "Alice Bob",
		Meta: map[string]interface{}{
			"flightID": "1234",
		},
	})

	// Assert
	assert.NoErrorf(t, err, "An error was not expected when initiating a mobile money franco payment")

	expectedResp := &InitiateMobileMoneyFrancoResponse{
		Status:  "success",
		Message: "Charge initiated",
		Meta:    map[string]interface{}{},
	}
	expectedResp.Data.ID = 1191267
	expectedResp.Data.TxRef = transRef
	expectedResp.Data.FlwRef = "FLW714021585320891879"
	expectedResp.Data.OrderID = "USS_URG_893982923s2323"
	expectedResp.Data.DeviceFingerprint = "62wd23423rq324323qew1"
	expectedResp.Data.Amount = 10_000
	expectedResp.Data.ChargedAmount = 10_000

	assert.Equal(t, expectedResp.Status, data.Status)
	assert.Equal(t, expectedResp.Message, data.Message)
	assert.Equal(t, expectedResp.Data.ID, data.Data.ID)
	assert.Equal(t, expectedResp.Data.TxRef, data.Data.TxRef)
	assert.Equal(t, expectedResp.Data.FlwRef, data.Data.FlwRef)
	assert.Equal(t, expectedResp.Data.OrderID, data.Data.OrderID)
	assert.Equal(t, expectedResp.Data.DeviceFingerprint, data.Data.DeviceFingerprint)
	assert.Equal(t, expectedResp.Data.Amount, data.Data.Amount)

	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

}

func Test_paymentsService_InitiateMobileMoneyFranco_BadRequest(t *testing.T) {

	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusBadRequest, "{}")
	defer server.Close()
	client := New(WithBaseURL(server.URL))
	ctx := context.Background()
	transRef := "MC-158523s09v5050e8"

	// Act
	_, response, err := client.Payments.InitiateMobileMoneyFranco(ctx, &InitiateMobileMoneyFrancoRequest{
		Amount:         10_000,
		Currency:       "XAF",
		PhoneNumber:    "+237678787878",
		Email:          "alice@example.com",
		TransactionRef: transRef,
		Country:        "CM",
		FullName:       "Alice Bob",
		Meta: map[string]interface{}{
			"flightID": "1234",
		},
	})

	// Assert
	assert.Errorf(t, err, "An error was expected when initiating a mobile money franco payment")

	assert.Equal(t, http.StatusBadRequest, response.HTTPResponse.StatusCode)

}
