package models

type Product struct {
	Title         string        `json:"title"`
	Specification string        `json:"specification"`
	Brand         string        `json:"brand"`
	Model         string        `json:"model"`
	Description   string        `json:"description"`
	Price         ProductPrice  `json:"price"`
	Source        ProductSource `json:"source"`
}

type ProductPrice struct {
	Value    int `json:"value"`
	Shipping int `json:"shipping"`
	Cashback int `json:"cashback"`
}

type ProductSource struct {
	Site     string `json:"site"`
	ImageURL string `json:"imageUrl"`
	PageURL  string `json:"pageUrl"`
}
