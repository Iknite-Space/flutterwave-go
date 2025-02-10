package stubs

// BillsCreateDStvPaymentResponse is a dummy JSOn response for checking a dstv user
func BillsCreateDStvPaymentResponse() string {
	return `
	{
		"status": "success",
		"message": "Bill payment successful",
		"data": {
			"phone_number": "+23490803840303",
			"amount": 500,
			"network": "9MOBILE",
			"flw_ref": "CF-FLYAPI-20200311081921359990",
			"tx_ref": "BPUSSD1583957963415840"
		}
	}
`
}

// BillsValidateDstvResponse is a dummy response for validating a DStv payment
func BillsValidateDstvResponse() string {
	return `
	{
		"status": "success",
		"message": "Item validated successfully",
		"data": {
			"response_code": "00",
			"address": null,
			"response_message": "Successful",
			"name": "MTN",
			"biller_code": "BIL099",
			"customer": "08038291822",
			"product_code": "AT099",
			"email": null,
			"fee": 100,
			"maximum": 0,
			"minimum": 0
		}
	}
`
}

// BillsGetStatusVerboseResponse is a dummy response of a verbose bill status
func BillsGetStatusVerboseResponse() string {
	return `
	{
		"status": "success",
		"message": "Bill status fetch successful",
		"data": {
			"currency": "NGN",
			"customer_id": "+23490803840303",
			"frequency": "One Time",
			"amount": "500.0000",
			"product": "AIRTIME",
			"product_name": "9MOBILE",
			"commission": 10,
			"transaction_date": "2020-03-11T20:19:21.27Z",
			"country": "NG",
			"tx_ref": "CF-FLYAPI-20200311081921359990",
			"extra": null,
			"product_details": "FLY-API-NG-AIRTIME-9MOBILE",
			"status": "successful"
		}
	}
`
}

// TransactionRefundResponse is a dummy response for refunding a transaction
func TransactionRefundResponse() []byte {
	return []byte(`
{
  "status": "success",
  "message": "Transaction refund initiated",
  "data": {
    "id": 75923,
    "account_id": 73362,
    "tx_id": 908790,
    "flw_ref": "URF_1577867664541_3572735",
    "wallet_id": 74639,
    "amount_refunded": 6900,
    "status": "completed",
    "destination": "payment_source",
    "meta": {
      "source": "availablebalance"
    },
    "created_at": "2021-01-24T09:18:37.366Z"
  }
}
`)
}

// TransferRatesResponse is a dummy response for fetching the tranfer rate of a given transaction.
func TransferRateResponse() []byte {
	return []byte(`
	{
  "status": "success",
  "message": "Transfer amount fetched",
  "data": {
    "rate": 624.24,
    "source": {
      "currency": "NGN",
      "amount": 624240
    },
    "destination": {
      "currency": "USD",
      "amount": 1000
    }
  }
}
	`)
}

// InitiateMobileMoneyFrancoResponse is a dummy response for initiating a mobile money franco payment
func InitiateMobileMoneyFrancoResponse() string {
	return `
	{
  "status": "success",
  "message": "Charge initiated",
  "data": {
    "id": 1191267,
    "tx_ref": "MC-158523s09v5050e8",
    "flw_ref": "FLW714021585320891879",
    "order_id": "USS_URG_893982923s2323",
    "device_fingerprint": "62wd23423rq324323qew1",
    "amount": 10000,
    "charged_amount": 1500,
    "app_fee": 1703,
    "merchant_fee": 0,
    "auth_model": "AUTH",
    "currency": "XAF",
    "ip": "154.123.220.1",
    "narration": "John Madakin",
    "status": "pending",
    "payment_type": "mobilemoneysn",
    "fraud_status": "ok",
    "charge_type": "normal",
    "created_at": "2020-03-27T14:54:50.000Z",
    "account_id": 73362,
    "customer": {
      "id": 349061,
      "phone_number": "237******20",
      "name": "John Madakin",
      "email": "user@flw.com",
      "created_at": "2020-03-27T09:12:54.000Z"
    },
    "processor_response": "Transaction in progress"
  },
  "meta": {
    "authorization": {
      "mode": "redirecturl",
      "redirect_url": "http://test.com/example"
    }
  }
}
  `
}
