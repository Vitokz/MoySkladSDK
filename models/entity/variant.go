package entity

// Variant Модификация
type Variant struct {
	Meta               *Meta            `json:"meta,omitempty"`
	Id                 string           `json:"id,omitempty"`
	AccountId          string           `json:"accountId,omitempty"`
	Updated            string           `json:"updated,omitempty"`
	Name               string           `json:"name,omitempty"`
	Code               string           `json:"code,omitempty"`
	ExternalCode       string           `json:"externalCode,omitempty"`
	Archived           bool             `json:"archived,omitempty"`
	Characteristics    *Characteristics `json:"characteristics,omitempty"`
	Images             *Images          `json:"images,omitempty"`
	MinPrice           *MinPrice        `json:"minPrice,omitempty"`   // Минимальная цена
	SalePrices         []SalePrice      `json:"salePrices,omitempty"` // Цены продажи
	BuyPrice           *BuyPrice        `json:"buyPrice,omitempty"`   // Закупочная цена
	Barcodes           []Barcode        `json:"barcodes,omitempty"`   // Штрихкоды
	Product            *Product         `json:"product,omitempty"`
	Things             []string         `json:"things,omitempty"`
	DiscountProhibited bool             `json:"discountProhibited,omitempty"` // Признак запрета скидок
	Packs              []Pack           `json:"packs,omitempty"`
}

// Characteristics ...
type Characteristics []struct {
	Meta  *Meta  `json:"meta,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
