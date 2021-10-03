package gopay

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type PaymentBody struct {
	Payer            Payer             `json:"payer,omitempty"`             // Objekt popisující plátce platby
	Target           Target            `json:"target,omitempty"`            // Objekt popisující příjemce platby
	Amount           int64             `json:"amount,omitempty"`            // Částka v haléřích
	Currency         string            `json:"currency,omitempty"`          // Určuje měnu platby, formát měny odpovídá [ISO 4217](http://www.iso.org/iso/home/standards/currency_codes.htm)
	OrderNumber      string            `json:"order_number,omitempty"`      // Identifikace objednávky v rámci prodejního místa (ID objednávky)
	OrderDescription string            `json:"order_description,omitempty"` // Popis objednávky
	Items            []Items           `json:"items,omitempty"`             // Detailně rozepsané jednotlivé položky objednávky
	EET              *EET              `json:"eet,omitempty"`               // EET údaje (povinné pro funkci EET)
	Callback         Callback          `json:"callback,omitempty"`          // Návratové URL a notifikační URL pro oznámení změny stavu platby
	AdditionalParams []AdditionalParam `json:"additional_params,omitempty"` // Doplňkové parametry platby
	Lang             string            `json:"lang,omitempty"`              // Parametr definuje jazyk platby
	Preauthorization bool              `json:"preauthorization,omitempty"`  // Pokud se jedná o předautorizovanou platbu, je předáván s hodnotou true
	Recurrence       Recurrence        `json:"recurrence,omitempty"`        // Pokud se jedná o opakovanou platbu, je předáván objekt, popisující opakování platby
}

type Payment struct {
	ID               int64             `json:"id,omitempty"`                // ID platby
	OrderNumber      string            `json:"order_number,omitempty"`      // ID objednávky
	State            PaymentState      `json:"state,omitempty"`             // Stav platby
	Amount           int64             `json:"amount,omitempty"`            // Částka v haléřích
	Currency         string            `json:"currency,omitempty"`          // Určuje měnu platby, formát měny odpovídá [ISO 4217](http://www.iso.org/iso/home/standards/currency_codes.htm)
	Payer            Payer             `json:"payer,omitempty"`             // Objekt popisující plátce platby
	Target           Target            `json:"target,omitempty"`            // Objekt popisující příjemce platby
	AdditionalParams []AdditionalParam `json:"additional_params,omitempty"` // Doplňkové parametry platby
	Lang             string            `json:"string,omitempty"`            // Parametr definuje jazyk platby
	GWURL            string            `json:"gw_url,omitempty"`            // URL pro inicializaci platební brány
}

func CreatePayment(client *Client, body PaymentBody) (*Payment, error) {
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

func RefundPayment(client *Client, ID int64, amount int) error {
	body := url.Values{}
	body.Add("amount", strconv.Itoa(amount))

	_, err := client.doRequest(fmt.Sprintf("/payments/payment/%d/refund", ID), "POST", requestBodyForm{body: body})
	if err != nil {
		return err
	}

	return nil
}
