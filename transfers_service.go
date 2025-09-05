package flutterwave

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// transfersService is the API client for the `/v3/transfers` endpoint
type transfersService service

// Estimate the Transfer Rate of a transaction
//
// API Docs: https://developer.flutterwave.com/reference/transfer_rates_post
func (service *transfersService) Rate(ctx context.Context, amount int, destination_currency, source_currency string) (*TransferRateResponse, *Response, error) {
	uri := "/v3/transfers/rates"

	const (
		// Precision is used to signify how many decimal places 
		// you want amount returned. if no precision is set, 6 will be used
		precision = 2
	)

	body := map[string]interface{}{
		"source": map[string]string{
			"currency": source_currency,
		},
		"destination": map[string]string{
			"currency": destination_currency,
			"amount": strconv.Itoa(amount),
		},
		"precision": precision,
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, uri, body)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrCouldNotConstructNewRequest, err)
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, fmt.Errorf("%w: %v", ErrRequestFailure, err)
	}

	var data TransferRateResponse
	if err = json.Unmarshal(*response.Body, &data); err != nil {
		return nil, response, fmt.Errorf("%w: %v", ErrUnmarshalFailure, err)
	}

	return &data, response, nil

}

func (service *transfersService) CreateTransfer(ctx context.Context, req CreateTransferRequest) (*CreateTransferResponse, *Response, error) {
	uri := "/v3/transfers"

	request, err := service.client.newRequest(ctx, http.MethodPost, uri, req)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrCouldNotConstructNewRequest, err)
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, fmt.Errorf("%w: %v", ErrRequestFailure, err)
	}

	var data CreateTransferResponse
	if err = json.Unmarshal(*response.Body, &data); err != nil {
		return nil, response, fmt.Errorf("%w: %v", ErrUnmarshalFailure, err)
	}

	return &data, response, nil
}
