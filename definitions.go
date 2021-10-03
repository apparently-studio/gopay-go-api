package gopay

type PaymentInstrument string

const PAYMENT_CARD = PaymentInstrument("PAYMENT_CARD") // Platební karty
const BANK_ACCOUNT = PaymentInstrument("BANK_ACCOUNT") // Bankovní převody
const PREMIUM_SMS = PaymentInstrument("PRSMS")         // Premium SMS
const MPAYMENT = PaymentInstrument("MPAYMENT")         // Mplatba
const PAYSAFECARD = PaymentInstrument("PAYSAFECARD")   // paysafecard
const GOPAY = PaymentInstrument("GOPAY")               // GoPay účet
const PAYPAL = PaymentInstrument("PAYPAL")             // PayPal účet
const BITCOIN = PaymentInstrument("BITCOIN")           // Platba bitcoiny
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

type BankAccount struct {
	Iban          string    `json:"iban,omitempty"`           // Kód IBAN bankovního účtu zákazníka
	Bic           BankSwift `json:"bic,omitempty"`            // SWIFT kód banky zákazníka
	Prefix        string    `json:"prefix,omitempty"`         // Předčíslí bankovního účtu zákazníka
	AccountNumber string    `json:"account_number,omitempty"` // Číslo bankovního účtu zákazníka
	BankCode      string    `json:"bank_code,omitempty"`      // Kód banky zákazníka
	AccountName   string    `json:"account_name,omitempty"`   // Jméno majitele bankovního účtu
}

type PaymentCard struct {
	CardNumber        string `json:"card_number,omitempty"`         // Vymaskované číslo platební karty
	CardExpiration    string `json:"card_expiration,omitempty"`     // Datum expirace
	CardBrand         string `json:"card_brand,omitempty"`          // Typ platební karty
	CardIssuerCountry string `json:"card_issuer_country,omitempty"` // Kód země vydavatelské banky
	CardIssuerBank    string `json:"card_issuer_bank,omitempty"`    // Vydavatelská banka
	CardToken         string `json:"card_token,omitempty"`          // Token platební karty pro účely identifikační platby
	Card3dsResult     string `json:"3ds_result,omitempty"`          // 3 není validní začátek struct memberu proto ten prefix Card
}

type Contact struct {
	FirstName   string `json:"first_name,omitempty"`   // Jméno zákazníka
	LastName    string `json:"last_name,omitempty"`    // Příjmení zákazníka
	Email       string `json:"email,omitempty"`        // Validní e-mail zákazníka
	PhoneNumber string `json:"phone_number,omitempty"` // Telefonní číslo zákazníka s předvolbou
	City        string `json:"city,omitempty"`         // Město zákazníka
	Street      string `json:"street,omitempty"`       // Ulice zákazníka
	PostalCode  string `json:"postal_code,omitempty"`  // Poštovní směrovací číslo zákazníka
	CountryCode string `json:"country_code,omitempty"` // Kód státu zákazníka
}

type Payer struct {
	AllowedPaymentInstruments []PaymentInstrument `json:"allowed_payment_instruments,omitempty"` // Pole povolených platebních metod
	DefaultPaymentInstrument  PaymentInstrument   `json:"default_payment_instrument,omitempty"`  // Preferovaná platební metoda
	DefaultSwift              BankSwift           `json:"default_swift,omitempty"`               // Preferová banka pokud je default_payment_instrument nastaveno na BANK_ACCOUNT, nastaveno pomocí SWIFT kódu banky
	AllowedSwifts             []BankSwift         `json:"allowed_swifts,omitempty"`              // Pole povolených kódů bank
	BankAccount               BankAccount         `json:"bank_account,omitempty"`                // Údaje o bankovním účtu plátce
	PaymentCard               PaymentCard         `json:"payment_card,omitempty"`                // Údaje o použité platební kartě
	Contact                   Contact             `json:"contact,omitempty"`                     // Údaje o zákaníkovi
	VerifyPin                 string              `json:"verify_pin,omitempty"`                  // PIN pro účely [identifikační platby](https://doc.gopay.com/cs/?lang=php#identifikacni-platba)
	AllowedCardToken          string              `json:"allowed_card_token,omitempty"`          // Token pro účely [identifikační platby](https://doc.gopay.com/cs/?lang=php#identifikacni-platba)
}

type Address struct {
	Street     string `json:"street,omitempty"`      // Ulice
	PostalCode string `json:"postal_code,omitempty"` // Poštovní směrovací číslo
	City       string `json:"city,omitempty"`        // Město
	Country    string `json:"country,omitempty"`     // Kód státu
}

type Accounts struct {
	Id       int    `json:"id,omitempty"`       // ID účtu
	Balance  int    `json:"balance,omitempty"`  // Zůstatek účtu
	Currency string `json:"currency,omitempty"` // Měna, formát měny odpovídá [ISO 4217](http://www.iso.org/iso/home/standards/currency_codes.htm)
}

type Target struct {
	Type string `json:"type,omitempty"` // Popis příjemce platby
	Goid int64  `json:"goid,omitempty"` // Jedinečný identifikátor eshopu v systému platební brány
}

type Callback struct {
	ReturnURL       string `json:"return_url,omitempty"`       // URL adresa pro návrat na eshop (včetně protokolu)
	NotificationUrl string `json:"notification_url,omitempty"` // URL adresa pro odeslání asynchronní notifikace v případě změny stavu platby (včetně protokolu)
}

type AdditionalParam struct {
	Name  string `json:"name,omitempty"`  // Název parametru
	Value string `json:"value,omitempty"` // 	Hodnota volitelného parametru
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
	RecurrenceCycle  RecurrenceCycle `json:"recurrence_cycle,omitempty"`   // Časový úsek opakování
	RecurrencePeriod int64           `json:"recurrence_period,omitempty"`  // Perioda opakování opakované platby
	RecurrenceDateTo string          `json:"recurrence_date_to,omitempty"` // Doba platnosti opakované platby (yyyy-mm-dd)
	RecurrenceState  RecurrenceState `json:"recurrence_state,omitempty"`   // Popis stavu opakované platby
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
	Type       ItemType    `json:"type,omitempty"`        // Typ položky
	ProductUrl string      `json:"product_url,omitempty"` // URL adresa produktu
	Ean        string      `json:"ean,omitempty"`         // [EAN kód produktu](https://cs.wikipedia.org/wiki/European_Article_Number)
	Count      int64       `json:"count,omitempty"`       // Počet položek produktu
	Name       string      `json:"name,omitempty"`        // Název produktu
	Amount     int64       `json:"amount,omitempty"`      // Součet cen položek s DPH v haléřích
	VatRate    ItemVatRate `json:"vat_rate,omitempty"`    // Součet cen položek s DPH v haléřích
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

const CREATED = PaymentState("CREATED")                             //	Platba založena
const PAYMENT_METHOD_CHOSEN = PaymentState("PAYMENT_METHOD_CHOSEN") //	Platební metoda vybrána
const PAID = PaymentState("PAID")                                   //	Platba zaplacena
const AUTHORIZED = PaymentState("AUTHORIZED")                       //	Platba předautorizována
const CANCELED = PaymentState("CANCELED")                           //	Platba zrušena
const TIMEOUTED = PaymentState("TIMEOUTED")                         //	Vypršelá platnost platby
const REFUNDED = PaymentState("REFUNDED")                           //	Platba refundována
const PARTIALLY_REFUNDED = PaymentState("PARTIALLY_REFUNDED")       //	Platba částečně refundována

var PaymentSecondaryStates = map[int]string{
	101:  "Čekáme na provedení online platby.",
	102:  "Čekáme na provedení offline platby.",
	3001: "Bankovní platba potvrzena avízem.",
	3002: "Bankovní platba potvrzena výpisem.",
	3003: "Bankovní platba nebyla potvrzena.",
	5001: "Schváleno s nulovou částkou",
	5002: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu dosažení limitů na platební kartě.",
	5003: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu problémů na straně vydavatele platební karty.",
	5004: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu problému na straně vydavatele platební karty.",
	5005: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu zablokované platební karty.",
	5006: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu nedostatku peněžních prostředků na platební kartě.",
	5007: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu expirované platební karty.",
	5008: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu zamítnutí CVV/CVC kódu.",
	5009: "Zamítnutí platby v systému 3D Secure banky zákazníka.",
	5015: "Zamítnutí platby v systému 3D Secure banky zákazníka.",
	5017: "Zamítnutí platby v systému 3D Secure banky zákazníka.",
	5018: "Zamítnutí platby v systému 3D Secure banky zákazníka.",
	5019: "Zamítnutí platby v systému 3D Secure banky zákazníka.",
	6502: "Zamítnutí platby v systému 3D Secure banky zákazníka.",
	6504: "Zamítnutí platby v systému 3D Secure banky zákazníka.",
	5010: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu problémů na platební kartě.",
	5014: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu problémů na platební kartě.",
	5011: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu problémů na účtu platební karty.",
	5036: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu problémů na účtu platební karty.",
	5012: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu technických problémů v autorizačním centru banky zákazníka.",
	5013: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu chybného zadání čísla platební karty.",
	5016: "Zamítnutí platby v autorizačním centru banky zákazníka, platba nebyla povolena na platební kartě zákazníka.",
	5020: "Neznámá konfigurace",
	5021: "Zamítnutí platby v autorizačním centru banky zákazníka z důvodu dosažení nastavených limitů na platební kartě.",
	5022: "Nastal technický problém spojený s autorizačním centrem banky zákazníka.",
	5023: "Platba nebyla provedena.",
	5038: "Platba nebyla provedena.",
	5024: "Platba nebyla provedena. Platební údaje nebyly zadány v časovém limitu na platební bráně.",
	5025: "Platba nebyla provedena. Konkrétní důvod zamítnutí je sdělen přímo zákazníkovi.",
	5026: "Platba nebyla provedena. Součet kreditovaných částek překročil uhrazenou částku.",
	5027: "Platba nebyla provedena. Uživatel není oprávněn k provedení operace.",
	5028: "Platba nebyla provedena. Částka k úhradě překročila autorizovanou částku.",
	5029: "Platba zatím nebyla provedena.",
	5030: "Platba nebyla provedena z důvodu opakovaného zadání platby.",
	5031: "Při platbě nastal technický problém na straně banky.",
	5033: "SMS se nepodařilo doručit.",
	5035: "Platební karta je vydaná v regionu, ve kterém nejsou podporovány platby kartou.",
	5037: "Držitel platební karty zrušil platbu.",
	5039: "Platba byla zamítnuta v autorizačním centru banky zákazníka z důvodu zablokované platební karty.",
	5040: "Duplicitni reversal transakce",
	5041: "Duplicitní transakce",
	5042: "Bankovní platba byla zamítnuta.",
	5043: "Platba zrušena uživatelem.",
	5044: "SMS byla odeslána. Zatím se ji nepodařilo doručit.",
	5045: "Platba byla přijata. Platba bude připsána po zpracování v síti Bitcoin.",
	5046: "Platba nebyla uhrazena v plné výši.",
	5047: "Platba byla provedena po splatnosti.",
}

// TODO: groups
// TODO: group-codes
// TODO: enabledPaymentInstruments
// TODO: enabledSwifts
