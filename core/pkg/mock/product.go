package mock

import "github.com/monitprod/core/pkg/models"

var NintendoSwitch = models.Product{
	Title:         "Nintendo Switch",
	Specification: "with Neon Blue and Neon Red Joy‑Con - HAC-001(-01)",
	Brand:         "Nintendo",
	Model:         "Switch",
	Description:   "",
	Price: models.ProductPrice{
		Value:    161666,
		Shipping: 50,
		Cashback: 5,
	},
	Source: models.ProductSource{
		Site:     "Amazon",
		ImageURL: "https://m.media-amazon.com/images/I/61-PblYntsL._AC_SX466_.jpg",
		PageURL:  "https://www.amazon.com.br/Nintendo-Switch-Azul-Vermelho-Neon/dp/B07VGRJDFY",
	},
}

var TectonAllChair = models.Product{
	Title:         "Cadeira Tecton All Black Unique",
	Specification: "Tecton All Black Unique",
	Brand:         "Flexform",
	Model:         "Tecton All Black Unique",
	Description:   "",
	Price: models.ProductPrice{
		Value:    166050,
		Shipping: 40,
		Cashback: 6,
	},
	Source: models.ProductSource{
		Site:     "Flexform E-commerce",
		ImageURL: "https://assets.betalabs.net/production/flexform/item-images/23dc7b0869e882ad51f7a8e2e4675201.png",
		PageURL:  "https://www.flexform.com.br/loja/office-chairs/cadeira-de-escritorio-flexform-tecton-all-black-unique",
	},
}

var Products = []models.Product{
	NintendoSwitch,
	TectonAllChair,
}
