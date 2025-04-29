package flutterwave

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// chargebacksService is the API client for the `/v3/chargeback` endpoint
type chargebacksService service

// ListChargebacks returns a list of chargebacks.
//
// API Docs: https://developer.flutterwave.com/reference/list-a-chargeback
func (service *chargebacksService) ListChargebacks(ctx context.Context, flwRef string) (*ListChargebacksResp, *Response, error) {
	uri := "/v3/chargebacks"

	request, err := service.client.newRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("%w: %v", ErrCouldNotConstructNewRequest, err)
	}

	// Adding the parameters
	params := map[string]string{}
	if flwRef != "" {
		params["flw_ref"] = flwRef
	}
	requestWithParams := service.client.addURLParams(request, params)

	response, err := service.client.do(requestWithParams)
	if err != nil {
		return nil, response, fmt.Errorf("%w: %v", ErrRequestFailure, err)
	}

	var data ListChargebacksResp
	if err = json.Unmarshal(*response.Body, &data); err != nil {
		return nil, response, fmt.Errorf("%w: %v", ErrUnmarshalFailure, err)
	}

	return &data, response, nil

}
