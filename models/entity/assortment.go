package entity

//Assortment Ассортимент товара
type Assortment struct {
	Meta                *Meta          `json:"meta,omitempty"`
	ID                  string         `json:"id,omitempty"`
	AccountID           string         `json:"accountId,omitempty"`
	Owner               *Employee      `json:"owner,omitempty"`
	Shared              bool           `json:"shared,omitempty"`
	Group               *Group         `json:"group,omitempty"`
	Updated             string         `json:"updated,omitempty"`
	Name                string         `json:"name,omitempty"`
	Code                string         `json:"code,omitempty"`
	ExternalCode        string         `json:"externalCode,omitempty"`
	Archived            bool           `json:"archived,omitempty"`
	PathName            string         `json:"pathName,omitempty"`
	Article             string         `json:"article,omitempty"`
	UseParentVat        bool           `json:"useParentVat,omitempty"`
	Vat                 int            `json:"vat,omitempty"`
	VatEnabled          bool           `json:"vatEnabled,omitempty"`
	EffectiveVat        int            `json:"effectiveVat,omitempty"`
	EffectiveVatEnabled int            `json:"effectiveVatEnabled,omitempty"`
	Label               string         `json:"label"`
	Images              *Images        `json:"images,omitempty"`
	MinPrice            *MinPrice      `json:"minPrice,omitempty"`
	SalePrices          []SalePrice    `json:"salePrices,omitempty"`
	BuyPrice            *BuyPrice      `json:"buyPrice,omitempty"`
	Barcodes            []Barcode      `json:"barcodes,omitempty"`
	Supplier            *Counterparty  `json:"supplier,omitempty"`
	PaymentItemType     string         `json:"paymentItemType,omitempty"`
	Description         string         `json:"description,omitempty"`
	DiscountProhibited  bool           `json:"discountProhibited,omitempty"`
	Weight              float64        `json:"weight,omitempty"`
	Volume              float64        `json:"volume,omitempty"`
	VariantsCount       int            `json:"variantsCount,omitempty"`
	IsSerialTrackable   bool           `json:"isSerialTrackable,omitempty"`
	TrackingType        string         `json:"trackingType,omitempty"`
	Files               File           `json:"files,omitempty"`
	Stock               float64        `json:"stock"`
	Reserve             float64        `json:"reserve"`
	InTransit           float64        `json:"inTransit"`
	Quantity            float64        `json:"quantity"`
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`
	Uom                 *Uom           `json:"uom,omitempty"`
	Weighed             bool           `json:"weighed,omitempty"`
	Attributes          []Attribute    `json:"attributes,omitempty"`
	Components          *Components    `json:"components"`
}

type Images struct {
	Meta    *Meta   `json:"meta"`           // Метаданные
	Context *Meta   `json:"context"`        // Метаданные о сотруднике выполнившем запрос
	Rows    []Image `json:"rows,omitempty"` // Массив структур изображений
}
