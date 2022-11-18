package gopay

type PaymentInstrument string

const PAYMENT_CARD = PaymentInstrument("PAYMENT_CARD") // Payment card
const BANK_ACCOUNT = PaymentInstrument("BANK_ACCOUNT") // Bank account
const PREMIUM_SMS = PaymentInstrument("PRSMS")         // Premium SMS
const MPAYMENT = PaymentInstrument("MPAYMENT")         // mPlatba (mobile payment)
const PAYSAFECARD = PaymentInstrument("PAYSAFECARD")   // PaySafeCard coupon
const GOPAY = PaymentInstrument("GOPAY")               // GoPay wallet
const PAYPAL = PaymentInstrument("PAYPAL")             // PayPal wallet
const BITCOIN = PaymentInstrument("BITCOIN")           // Bitcoin wallet
const MASTERPASS = PaymentInstrument("MASTERPASS")     // Masterpass
const GOOGLE_PAY = PaymentInstrument("GPAY")           // Google Pay

type BankSwift string

const CESKA_SPORITELNA = BankSwift("GIBACZPX")
const KOMERCNI_BANKA = BankSwift("KOMBCZPP")
const RAIFFEISENBANK = BankSwift("RZBCCZPP")
const MBANK = BankSwift("BREXCZPP")
const FIO_BANKA = BankSwift("FIOBCZPP")
const CSOB = BankSwift("CEKOCZPP")
const ERA = BankSwift("CEKOCZPP-ERA")
const UNICREDIT_BANK_CZ = BankSwift("BACXCZPP")
const VSEOBECNA_VEROVA_BANKA_BANKA = BankSwift("SUBASKBX")
const TATRA_BANKA = BankSwift("TATRSKBX")
const UNICREDIT_BANK_SK = BankSwift("UNCRSKBX")
const SLOVENSKA_SPORITELNA = BankSwift("GIBASKBX")
const POSTOVA_BANKA = BankSwift("POBNSKBA")
const CSOB_SK = BankSwift("CEKOSKBX")
const SBERBANK_SLOVENSKO = BankSwift("LUBASKBX")
const SPECIAL = BankSwift("OTHERS")
const MBANK1 = BankSwift("BREXPLPW")
const CITI_HANDLOWY = BankSwift("CITIPLPX")
const IKO = BankSwift("BPKOPLPW-IKO")
const INTELIGO = BankSwift("BPKOPLPW-INTELIGO")
const PLUS_BANK = BankSwift("IVSEPLPP")
const BANK_BPH_SA = BankSwift("BPHKPLPK")
const TOYOTA_BANK = BankSwift("TOBAPLPW")
const VOLKSWAGEN_BANK = BankSwift("VOWAPLP1")
const SGB = BankSwift("GBWCPLPP")
const POCZTOWY_BANK = BankSwift("POCZPLP4")
const BGZ_BANK = BankSwift("GOPZPLPW")
const IDEA = BankSwift("IEEAPLPA")
const BPS = BankSwift("POLUPLPR")
const GETIN_ONLINE = BankSwift("GBGCPLPK-GIO")
const BLIK = BankSwift("GBGCPLPK-BLIK")
const NOBLE_BANK = BankSwift("GBGCPLPK-NOB")
const ORANGE = BankSwift("BREXPLPW-OMB")
const BZ_WBK = BankSwift("WBKPPLPP")
const RAIFFEISEN_BANK_POLSKA_SA = BankSwift("RCBWPLPW")
const POWSZECHNA_KASA_OSZCZEDNOSCI_BANK_POLSKI_SA = BankSwift("BPKOPLPW")
const ALIOR_BANK = BankSwift("ALBPPLPW")
const ING_BANK_SLASKI = BankSwift("INGBPLPW")
const PEKAO_SA = BankSwift("PKOPPLPW")
const GETIN_ONLINE1 = BankSwift("GBGCPLPK")
const BANK_MILLENNIUM = BankSwift("BIGBPLPW")
const BANK_OCHRONY_SRODOWISKA = BankSwift("EBOSPLPW")
const BNP_PARIBAS_POLSKA = BankSwift("PPABPLPK")
const CREDIT_AGRICOLE = BankSwift("AGRIPLPR")
const DEUTSCHE_BANK_POLSKA_SA = BankSwift("DEUTPLPX")
const DNB_NORD = BankSwift("DNBANOKK")
const E_SKOK = BankSwift("NBPLPLPW")
const EUROBANK = BankSwift("SOGEPLPW")
const POLSKI_BANK_PRZEDSIEBIORCZOSCI_SPOLKA_AKCYJNA = BankSwift("PBPBPLPW")

// https://doc.gopay.com/#bank-account
type BankAccount struct {
	Iban          string     `json:"iban,omitempty"`           // International bank account number
	Bic           *BankSwift `json:"bic,omitempty"`            // Business identification code (SWIFT)
	Prefix        string     `json:"prefix,omitempty"`         // Bank account prefix
	AccountNumber string     `json:"account_number,omitempty"` // Bank account number
	BankCode      string     `json:"bank_code,omitempty"`      // Bank account code
	AccountName   string     `json:"account_name,omitempty"`   // Bank account name
}

// https://doc.gopay.com/#payment-card
type PaymentCard struct {
	CardNumber        string `json:"card_number,omitempty"`         // Masked payment card´s number
	CardExpiration    string `json:"card_expiration,omitempty"`     // Expiration date
	CardBrand         string `json:"card_brand,omitempty"`          // Payment card´s type
	CardIssuerCountry string `json:"card_issuer_country,omitempty"` // Country code of issuing bank
	CardIssuerBank    string `json:"card_issuer_bank,omitempty"`    // Issuing bank
	CardToken         string `json:"card_token,omitempty"`          // Token for identification payment purposes
	Card3dsResult     string `json:"3ds_result,omitempty"`          // 3D Secure authorization’s result for identification payment purposes
}

// https://doc.gopay.com/#contact
type Contact struct {
	FirstName   string `json:"first_name,omitempty"`   // First name
	LastName    string `json:"last_name,omitempty"`    // Last name
	Email       string `json:"email,omitempty"`        // Valid e-mail
	PhoneNumber string `json:"phone_number,omitempty"` // Phone number with country code
	City        string `json:"city,omitempty"`         // City
	Street      string `json:"street,omitempty"`       // Street
	PostalCode  string `json:"postal_code,omitempty"`  // Postal code
	CountryCode string `json:"country_code,omitempty"` // Country code ISO 3166-1 alpha-3
}

// https://doc.gopay.com/#payer
type Payer struct {
	AllowedPaymentInstruments []PaymentInstrument `json:"allowed_payment_instruments,omitempty"` // Array of allowed payment methods
	DefaultPaymentInstrument  *PaymentInstrument  `json:"default_payment_instrument,omitempty"`  // Preferred payment method
	DefaultSwift              *BankSwift          `json:"default_swift,omitempty"`               // Preferred bank if default_payment_instrument is set to BANK_ACCOUNT, set by SWIFT code
	AllowedSwifts             []BankSwift         `json:"allowed_swifts,omitempty"`              // Array of allowed bank codes
	BankAccount               *BankAccount        `json:"bank_account,omitempty"`                // Bank account´s information
	PaymentCard               *PaymentCard        `json:"payment_card,omitempty"`                // Payment card´s information
	Contact                   *Contact            `json:"contact,omitempty"`                     // Customer´s data
	VerifyPin                 string              `json:"verify_pin,omitempty"`                  // PIN for identification payment purposes
	AllowedCardToken          string              `json:"allowed_card_token,omitempty"`          // Token for identification payment purposes
}

type Address struct {
	Street     string `json:"street,omitempty"`      // Street
	PostalCode string `json:"postal_code,omitempty"` // PostalCode
	City       string `json:"city,omitempty"`        // City
	Country    string `json:"country,omitempty"`     // Country code ISO 3166-1 alpha-3
}

type Accounts struct {
	Id       int    `json:"id,omitempty"`       // Account ID
	Balance  int    `json:"balance,omitempty"`  // Balance
	Currency string `json:"currency,omitempty"` // Currency [ISO 4217](http://www.iso.org/iso/home/standards/currency_codes.htm)
}

type Target struct {
	Type string `json:"type,omitempty"` // Always ACCOUNT
	Goid int64  `json:"goid,omitempty"` // Unique identifier of an e-shop in the payment gateway system
}

type Callback struct {
	ReturnURL       string `json:"return_url,omitempty"`       // URL address for return to e-shop (with protocol)
	NotificationUrl string `json:"notification_url,omitempty"` // URL address for sending asynchronous notification in the case of changes in the payment status (with protocol)
}

type AdditionalParam struct {
	Name  string `json:"name,omitempty"`  // Parameter name
	Value string `json:"value,omitempty"` // Value of optional parameter
}

type RecurrenceCycle string

const DAY = RecurrenceCycle("DAY")
const WEEK = RecurrenceCycle("WEEK")
const MONTH = RecurrenceCycle("MONTH")
const ON_DEMAND = RecurrenceCycle("ON_DEMAND")

type RecurrenceState string

const REQUESTED = RecurrenceState("REQUESTED")
const STARTED = RecurrenceState("STARTED")
const STOPPED = RecurrenceState("STOPPED")

type Recurrence struct {
	RecurrenceCycle  RecurrenceCycle `json:"recurrence_cycle,omitempty"`   // Time period of recurring
	RecurrencePeriod int64           `json:"recurrence_period,omitempty"`  // Recurring period of recurring payment
	RecurrenceDateTo string          `json:"recurrence_date_to,omitempty"` // The period of validity recurring payment (yyyy-mm-dd)
	RecurrenceState  RecurrenceState `json:"recurrence_state,omitempty"`   // Describes state of recurring payment
}

type ItemType string

const ITEM = ItemType("ITEM")
const DELIVERY = ItemType("DELIVERY")
const DISCOUNT = ItemType("DISCOUNT")

type ItemVatRate int

const ZERO = ItemVatRate(0)
const TEN = ItemVatRate(10)
const FIFTEEN = ItemVatRate(15)
const TWENTYONE = ItemVatRate(21)

type Items struct {
	Type       ItemType    `json:"type,omitempty"`        // Type of row, for registration of sales
	ProductUrl string      `json:"product_url,omitempty"` // URL address of the product
	Ean        string      `json:"ean,omitempty"`         // [EAN code of the product](https://en.wikipedia.org/wiki/International_Article_Number)
	Count      int64       `json:"count,omitempty"`       // Number of items
	Name       string      `json:"name,omitempty"`        // Product name
	Amount     int64       `json:"amount,omitempty"`      // Total price of items in cents
	VatRate    ItemVatRate `json:"vat_rate,omitempty"`    // VAT rate
}

// TODO: možná anglicky? (GoPay to má taky v čj)
type EET struct {
	DicPoverujiciho string `json:"dic_poverujiciho,omitempty"` //	DIČ pověřujícího poplatníka	varchar
	CelkTrzba       int64  `json:"celk_trzba,omitempty"`       //	Celková částka tržby	long v centech
	ZaklNepodlDph   int64  `json:"zakl_nepodl_dph,omitempty"`  //	Celková částka plnění osvobozených od DPH	long v centech
	ZaklDan1        int64  `json:"zakl_dan1,omitempty"`        //	Celkový základ daně se základní sazbou DPH	long v centech
	Dan1            int64  `json:"dan1,omitempty"`             //	Celková DPH se základní sazbou	long v centech
	ZaklDan2        int64  `json:"zakl_dan2,omitempty"`        //	Celkový základ daně s první sníženou sazbou DPH	long v centech
	Dan2            int64  `json:"dan2,omitempty"`             //	Celková DPH s první sníženou sazbou	long v centech
	ZaklDan3        int64  `json:"zakl_dan3,omitempty"`        //	Celkový základ daně s druhou sníženou sazbou DPH	long v centech
	Dan3            int64  `json:"dan3,omitempty"`             //	Celková DPH s druhou sníženou sazbou	long v centech
	CestSluz        int64  `json:"cest_sluz,omitempty"`        //	Celková částka v režimu DPH pro cestovní službu	long v centech
	PouzitZboz1     int64  `json:"pouzit_zboz1,omitempty"`     //	Celková částka v režimu DPH pro prodej použitého zboží se základní sazbou	long v centech
	PouzitZboz2     int64  `json:"pouzit_zboz2,omitempty"`     //	Celková částka v režimu DPH pro prodej použitého zboží s první sníženou sazbou	long v centech
	PouzitZboz3     int64  `json:"pouzit_zboz3,omitempty"`     //	Celková částka v režimu DPH pro prodej použitého zboží s druhou sníženou sazbou	long v centech
	UrcenoCerpZuct  int64  `json:"urceno_cerp_zuct,omitempty"` //	Celková částka plateb určená k následnému čerpání nebo zúčtování	long v centech
	CerpZuct        int64  `json:"cerp_zuct,omitempty"`        //	Celková částka plateb, které jsou následným čerpáním nebo zúčtováním platby	long v centech
	Mena            int64  `json:"mena,omitempty"`             //	Měna, ve které jsou údaje předávány
}

type EETCode struct {
	FIK string `json:"fik,omitempty"` // Fiskální identifikační kód (FIK)
	BKP string `json:"bkp,omitempty"` // Bezpečnostní kód poplatníka (BKP)
	PKP string `json:"pkp,omitempty"` // Podpisový kód poplatníka (PKP)
}

type PaymentState string

const CREATED = PaymentState("CREATED")                             //	Payment created
const PAYMENT_METHOD_CHOSEN = PaymentState("PAYMENT_METHOD_CHOSEN") //	Payment method confirmed
const PAID = PaymentState("PAID")                                   //	Payment has already been paid
const AUTHORIZED = PaymentState("AUTHORIZED")                       //	Payment pre-authorized
const CANCELED = PaymentState("CANCELED")                           //	Payment declined
const TIMEOUTED = PaymentState("TIMEOUTED")                         //	The payment has expired
const REFUNDED = PaymentState("REFUNDED")                           //	Payment refunded
const PARTIALLY_REFUNDED = PaymentState("PARTIALLY_REFUNDED")       //	Payment partially refunded

var PaymentSecondaryStates = map[int]string{
	101:  "We are waiting for an online payment.",
	102:  "We are waiting for an offline payment.",
	3001: "Bank payment confirmed by advice",
	3002: "Bank payment confirmed by statement.",
	3003: "Bank payment was not confirmed.",
	5001: "Approved with zero amount.",
	5002: "Rejection of the payment in the authorization center of the customer’s bank due to reaching the limits on the payment card.",
	5003: "Refusal of payment in the authorization center of the customer’s bank due to problems on the part of the payment card issuer.",
	5004: "Rejection of the payment in the authorization center of the customer’s bank due to a problem on the part of the payment card issuer.",
	5005: "Rejection of the payment in the authorization center of the customer’s bank due to a blocked payment card.",
	5006: "Refusal of payment in the authorization center of the customer’s bank due to lack of funds on the payment card.",
	5007: "Rejection of the payment in the authorization center of the customer’s bank due to the expired payment card.",
	5008: "Rejection of payment in the authorization center of the customer’s bank due to rejection of the CVV/CVC code.",
	5009: "Rejection of payment in the customer’s 3D Secure bank system.",
	5015: "Rejection of payment in the customer’s 3D Secure bank system.",
	5017: "Rejection of payment in the customer’s 3D Secure bank system.",
	5018: "Rejection of payment in the customer’s 3D Secure bank system.",
	5019: "Rejection of payment in the customer’s 3D Secure bank system.",
	6502: "Rejection of payment in the customer’s 3D Secure bank system.",
	6504: "Rejection of payment in the customer’s 3D Secure bank system.",
	5010: "Rejection of the payment in the authorization center of the customer’s bank due to problems on the payment card.",
	5014: "Rejection of the payment in the authorization center of the customer’s bank due to problems on the payment card.",
	5011: "Rejection of the payment in the authorization center of the customer’s bank due to problems on the payment card.",
	5036: "Rejection of the payment in the authorization center of the customer’s bank due to problems on the payment card.",
	5012: "Rejection of the payment in the authorization center of the customer’s bank due to problems on the payment card.",
	5013: "Rejection of the payment in the authorization center of the customer’s bank due to problems on the payment card.",
	5016: "Rejection of the payment in the authorization center of the customer’s bank due to problems on the payment card.",
	5020: "Unknown configuration",
	5021: "Rejection of the payment in the authorization center of the customer’s bank due to reaching the set limits on the payment card.",
	5022: "There was a technical problem associated with the customer’s bank authorization center.",
	5023: "Payment was not made.",
	5038: "Payment was not made.",
	5024: "Payment was not made. Payment details were not entered within the time limit on the payment gateway.",
	5025: "Payment was not made. The specific reason for rejection is communicated directly to the customer.",
	5026: "Payment was not made. The sum of the credited amounts exceeded the amount paid.",
	5027: "Payment was not made. The user is not authorized to perform the operation.",
	5028: "Payment was not made. The amount to be paid exceeded the authorized amount.",
	5029: "Payment has not been made yet.",
	5030: "The payment was not made due to a re-entry.",
	5031: "There was a technical problem with the payment on the part of the bank.",
	5033: "SMS could not be delivered.",
	5035: "The payment card is issued in a region where card payments are not supported.",
	5037: "The credit card holder canceled the payment.",
	5039: "The payment was declined at the customer’s bank authorization center due to a blocked credit card.",
	5040: "Duplicate reversal transactions",
	5041: "Duplicate transactions",
	5042: "Bank payment declined.",
	5043: "Payment canceled by user.",
	5044: "SMS has been sent. It has not yet been delivered.",
	5045: "Payment received. Payment will be credited after processing on Bitcoin.",
	5046: "The payment was not paid in full.",
	5047: "Payment was made overdue.",
}

// TODO: groups
// TODO: group-codes
// TODO: enabledPaymentInstruments
// TODO: enabledSwifts
