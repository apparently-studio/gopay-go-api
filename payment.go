package gopay

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Payment body used when creating Payment..
type PaymentBody struct {
	Payer            *Payer            `json:"payer,omitempty"`             // Payment method settings and payer information
	Target           *Target           `json:"target,omitempty"`            // Payee information
	Amount           int64             `json:"amount,omitempty"`            // Payment amount in cents
	Currency         string            `json:"currency,omitempty"`          // Payment currency [ISO 4217](http://www.iso.org/iso/home/standards/currency_codes.htm)
	OrderNumber      string            `json:"order_number,omitempty"`      // Merchant’s order id, alphanumeric
	OrderDescription string            `json:"order_description,omitempty"` // Order description, alphanumeric
	Items            []Items           `json:"items,omitempty"`             // Details of the payment items
	EET              *EET              `json:"eet,omitempty"`               // EET údaje (povinné pro funkci EET)
	Callback         *Callback         `json:"callback,omitempty"`          // Callback URL for processing of the payment result / Notification URL for processing of change of payment status
	AdditionalParams []AdditionalParam `json:"additional_params,omitempty"` // Additional payment parameters
	Lang             string            `json:"lang,omitempty"`              // Payment gateway language
	Preauthorization bool              `json:"preauthorization,omitempty"`  // true if the payment should be preauthorized
	Recurrence       *Recurrence       `json:"recurrence,omitempty"`        // Contains object describing recurrence, if the payment should be recurrent
}

// Payment struct returned from REST API.
type Payment struct {
	ID               int64             `json:"id,omitempty"`                // Payment ID
	OrderNumber      string            `json:"order_number,omitempty"`      // Order ID
	State            *PaymentState     `json:"state,omitempty"`             // Payment status
	Amount           int64             `json:"amount,omitempty"`            // Amount in cents
	Currency         string            `json:"currency,omitempty"`          // Payment currency [ISO 4217](http://www.iso.org/iso/home/standards/currency_codes.htm)
	Payer            *Payer            `json:"payer,omitempty"`             // Information about the payer and payment methods
	Target           *Target           `json:"target,omitempty"`            // Payee information
	AdditionalParams []AdditionalParam `json:"additional_params,omitempty"` // Additional parameters
	Lang             string            `json:"string,omitempty"`            // Payment gateway language
	GWURL            string            `json:"gw_url,omitempty"`            // URL for initiation of the payment gate
	Preauthorization bool              `json:"preauthorization,omitempty"`  // true if the payment should be preauthorized
	Recurrence       *Recurrence       `json:"recurrence,omitempty"`        // Contains object describing recurrence, if the payment should be recurrent
}

// Create Payment
// https://doc.gopay.com/#payment-creation
func CreatePayment(client *Client, body PaymentBody) (*Payment, error) {
	if body.Target.Goid == 0 || body.Target.Type != "ACCOUNT" {
		body.Target = &Target{
			Type: "ACCOUNT",
			Goid: int64(client.goId),
		}
	}

	res, err := client.doRequest("/payments/payment", "POST", requestBodyJSON{body: body})
	if err != nil {
		return nil, err
	}

	payment := Payment{}

	err = json.NewDecoder(res.Body).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

// Get Payment
// https://doc.gopay.com/#payment-status
func GetPayment(client *Client, ID int64) (*Payment, error) {
	res, err := client.doRequest(fmt.Sprintf("/payments/payment/%d", ID), "GET", nil)
	if err != nil {
		return nil, err
	}

	payment := Payment{}

	err = json.NewDecoder(res.Body).Decode(&payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

// Refund Payment
// https://doc.gopay.com/#payment-refund
func RefundPayment(client *Client, ID int64, amount int) error {
	body := url.Values{}
	body.Add("amount", strconv.Itoa(amount))

	_, err := client.doRequest(fmt.Sprintf("/payments/payment/%d/refund", ID), "POST", requestBodyForm{body: body})
	if err != nil {
		return err
	}

	return nil
}
