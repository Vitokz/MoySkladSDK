package entity

import (
	"github.com/Vitokz/MoySkladSDK/models/general"
	"time"
)

type Consignment struct {
	Meta         *general.Meta     `json:"meta"`                  // Метаданные
	Id           string            `json:"id"`                    // ID Серии (Только для чтения)
	AccountId    string            `json:"accountId"`             // ID учетной записи (Только для чтения)
	Updated      time.Time         `json:"updated"`               // Момент последнего обновления сущности (Только для чтения)
	Name         string            `json:"name"`                  // Наименование Серии
	Description  string            `json:"description,omitempty"` // Описание Серии
	Code         string            `json:"code"`                  // Код Серии
	ExternalCode string            `json:"externalCode"`          // Внешний код Серии
	Label        string            `json:"label"`                 // Метка Серии
	Barcodes     []general.Barcode `json:"barcodes,omitempty"`    // Штрихкоды Комплекта
	Image        *general.Image    `json:"image"`                 // Изображение товара к которому относится данная серия
	Attributes   *general.Meta     `json:"attributes"`            // Метаданные ссылки или модификации
}
