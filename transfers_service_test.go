package flutterwave

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NdoleStudio/flutterwave-go/internal/helpers"
	"github.com/NdoleStudio/flutterwave-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestTransfersService_Rate(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, string(stubs.TransferRateResponse())) // Mock API response
	client := New(WithBaseURL(server.URL))

	// Act
	rate, response, err := client.Transfers.Rate(context.Background(), 1000, "USD", "NGN")

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	// Unmarshal expected response for assertion
	var expectedData TransferRateResponse
	err = json.Unmarshal([]byte(stubs.TransferRateResponse()), &expectedData)
	assert.Nil(t, err)

	assert.Equal(t, expectedData, *rate) // Ensure response struct matches expected output

	// Teardown
	server.Close()
}

func TestTransfersService_Rate_Failure(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	errorResponse := `{"error": "internal server error"}`
	server := helpers.MakeTestServer(http.StatusInternalServerError, errorResponse)
	client := New(WithBaseURL(server.URL))

	// Act
	rate, response, err := client.Transfers.Rate(context.Background(), 1000, "USD", "NGN")

	// Assert
	assert.NotNil(t, err) // Expect an error
	assert.Nil(t, rate)   // The rate should be nil due to failure
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusInternalServerError, response.HTTPResponse.StatusCode)

	assert.Contains(t, err.Error(), "500") // Ensure error message contains 500
	assert.Contains(t, err.Error(), "internal server error") // Ensure error message contains actual response error

	// Teardown
	server.Close()
}

func TestTransfersService_Rate_ErrCouldNotConstructNewRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange - Create a client with an invalid base URL
	client := New(WithBaseURL("://invalid-url"))

	// Act
	rate, response, err := client.Transfers.Rate(context.Background(), 1000, "USD", "NGN")

	// Assert
	assert.Nil(t, rate)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrCouldNotConstructNewRequest), "Expected ErrCouldNotConstructNewRequest")

	// Ensure error message contains relevant details
	assert.Contains(t, err.Error(), "could not construct new request")
}

func TestTransfersService_Rate_ErrRequestFailure(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange - Create a client pointing to a non-existent server
	client := New(WithBaseURL("http://127.0.0.1:54321")) // Non-listening port

	// Act
	rate, response, err := client.Transfers.Rate(context.Background(), 1000, "USD", "NGN")

	// Assert
	assert.Nil(t, rate)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrRequestFailure), "Expected ErrRequestFailure")
	assert.Contains(t, err.Error(), "request failed")
}

func TestTransfersService_Rate_ErrUnmarshalFailure(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange - Return malformed JSON from mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{invalid-json-response}`) // Malformed JSON
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	// Act
	rate, response, err := client.Transfers.Rate(context.Background(), 1000, "USD", "NGN")

	// Assert
	assert.Nil(t, rate)
	assert.NotNil(t, response)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrUnmarshalFailure), "Expected ErrUnmarshalFailure")
	assert.Contains(t, err.Error(), "failed to unmarshal response")
}

func TestTransfersService_GetTransfer(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := helpers.MakeTestServer(http.StatusOK, string(stubs.GetTransferResponse()))
	client := New(WithBaseURL(server.URL))

	// Act
	transfer, response, err := client.Transfers.GetTransfer(context.Background(), "12345")

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	var expectedData GetTransferResponse
	err = json.Unmarshal(stubs.GetTransferResponse(), &expectedData)
	assert.Nil(t, err)

	assert.Equal(t, expectedData, *transfer)

	// Teardown
	server.Close()
}

func TestTransfersService_GetTransfer_Failure(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	errorResponse := `{"status": "error", "message": "Transfer not found"}`
	server := helpers.MakeTestServer(http.StatusNotFound, errorResponse)
	client := New(WithBaseURL(server.URL))

	// Act
	transfer, response, err := client.Transfers.GetTransfer(context.Background(), "99999")

	// Assert
	assert.NotNil(t, err)
	assert.Nil(t, transfer)
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusNotFound, response.HTTPResponse.StatusCode)

	assert.Contains(t, err.Error(), "404")

	// Teardown
	server.Close()
}

func TestTransfersService_GetTransfer_ErrCouldNotConstructNewRequest(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	client := New(WithBaseURL("://invalid-url"))

	// Act
	transfer, response, err := client.Transfers.GetTransfer(context.Background(), "12345")

	// Assert
	assert.Nil(t, transfer)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrCouldNotConstructNewRequest), "Expected ErrCouldNotConstructNewRequest")
	assert.Contains(t, err.Error(), "could not construct new request")
}

func TestTransfersService_GetTransfer_ErrRequestFailure(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	client := New(WithBaseURL("http://127.0.0.1:54321"))

	// Act
	transfer, response, err := client.Transfers.GetTransfer(context.Background(), "12345")

	// Assert
	assert.Nil(t, transfer)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrRequestFailure), "Expected ErrRequestFailure")
	assert.Contains(t, err.Error(), "request failed")
}

func TestTransfersService_GetTransfer_ErrUnmarshalFailure(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{invalid-json-response}`)
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	// Act
	transfer, response, err := client.Transfers.GetTransfer(context.Background(), "12345")

	// Assert
	assert.Nil(t, transfer)
	assert.NotNil(t, response)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrUnmarshalFailure), "Expected ErrUnmarshalFailure")
	assert.Contains(t, err.Error(), "failed to unmarshal response")
}
