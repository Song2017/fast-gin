package order_domain

type ShipFromAddress struct {
	Address1    string `json:"address1"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	County      string `json:"county"`
	Zip         string `json:"zip"`
}
