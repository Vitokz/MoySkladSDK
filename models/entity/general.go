package entity

// MinPrice Цена
type MinPrice struct {
	Value    float64   `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате Метаданных
}

// SalePrice Цена продажи
type SalePrice struct {
	Value     float64    `json:"value,omitempty"`     // Значение цены
	Currency  *Currency  `json:"currency,omitempty"`  // Ссылка на валюту в формате
	PriceType *PriceType `json:"priceType,omitempty"` // Тип цены
}

// BuyPrice Закупочная цена
type BuyPrice struct {
	Value    float64   `json:"value,omitempty"`    // Значение цены
	Currency *Currency `json:"currency,omitempty"` // Ссылка на валюту в формате
}

// Barcode Штрихкоды
type Barcode struct {
	EAN13   string `json:"ean13,omitempty"`   // штрихкод в формате EAN13, если требуется создать штрихкод в формате EAN13
	EAN8    string `json:"ean8,omitempty"`    // штрихкод в формате EAN8, если требуется создать штрихкод в формате EAN8
	Code128 string `json:"code128,omitempty"` // штрихкод в формате Code128, если требуется создать штрихкод в формате Code128
	GTIN    string `json:"gtin,omitempty"`    // штрихкод в формате GTIN, если требуется создать штрихкод в формате GTIN. Валидируется на соответствие формату GS1
	UPC     string `json:"upc,omitempty"`     // штрихкод в формате UPC, если требуется создать штрихкод в формате UPC.
}

//Uom Единица измерения
type Uom struct {
	Meta         *Meta     `json:"meta,omitempty"`         // Метаданные
	ID           string    `json:"id,omitempty"`           // ID Единиицы измерения
	AccountID    string    `json:"accountId,omitempty"`    // ID учетной записи
	Owner        *Employee `json:"owner,omitempty"`        // Владелец
	Shared       bool      `json:"shared,omitempty"`       // Общий доступ
	Group        *Group    `json:"group,omitempty"`        // Отдел сотрудника
	Updated      string    `json:"updated,omitempty"`      // Момент последнего обновления
	Name         string    `json:"name,omitempty"`         // Наимнование единицы измерения
	Description  string    `json:"description,omitempty"`  // Описание единицы измерения
	Code         string    `json:"code,omitempty"`         // Код единицы измерения
	ExternalCode string    `json:"externalCode,omitempty"` // Внешний код Единицы измерения
}

//PriceType Тип цены
type PriceType struct {
	Meta         *Meta  `json:"meta,omitempty"`         // Метаданные
	Id           string `json:"id,omitempty"`           // Id типа цены
	Name         string `json:"name,omitempty"`         // Наимнование типа цены
	ExternalCode string `json:"externalCode,omitempty"` // Внешний код типа цены
}

// Attribute дополнительные поля
type Attribute struct {
	Meta        *Meta       `json:"meta,omitempty"`     // Метаданные
	ID          string      `json:"id"`                 // Id до поля
	Name        string      `json:"name"`               // Наименование доп. поля
	Type        string      `json:"type"`               // Тип доп. поля
	Value       interface{} `json:"value,omitempty"`    // Значение доп. поля
	Required    bool        `json:"required,omitempty"` // Является ли доп. поле обязательным
	Description string      `json:"description"`        // Описание доп поля
}

//Meta метаданных какой-либо сущности.
//Не все поля структуры могут присутствовать в ответе сервера
type Meta struct {
	Href         string `json:"href,omitempty"`         // Ссылка на объект
	MetadataHref string `json:"metadataHref,omitempty"` // Ссылка на метаданные сущности
	Type         string `json:"type,omitempty"`         // Тип объекта
	MediaType    string `json:"mediaType,omitempty"`    // Тип данных которые приходят от сервера
	UuidHref     string `json:"uuidHref"`               // Ссылка на объект UI
	DownloadHref string `json:"downloadHref,omitempty"` // Ссылка на скачивание
	Size         int    `json:"size,omitempty"`         // Размер возвращенного списка
	Limit        int    `json:"limit,omitempty"`        // Максимальное количество элементов(по дефолту 1000)
	Offset       int    `json:"offset,omitempty"`       // Порядковый номер элемента на котором закончился список
	NextHref     string `json:"nextHref"`               // Ссылка на следующую страницу сущностей
	PreviousHref string `json:"previousHref,omitempty"` // Ссылка на предыдущую страницу сущностей
}

// Cashier Кассир
type Cashier struct {
	Meta        *Meta        `json:"meta"`        // Метаданные Кассира (Только для чтения)
	ID          string       `json:"id"`          // ID Сотрудника (Только для чтения)
	AccountID   string       `json:"accountId"`   // ID учетной записи (Только для чтения)
	Employee    *Employee    `json:"employee"`    // Метаданные сотрудника, которого представляет собой кассир
	RetailStore *RetailStore `json:"retailStore"` // Метаданные точки продаж, к которой прикреплен кассир
}

// EmbeddedTemplate Шаблон печатной формы
type EmbeddedTemplate struct {
	Meta    *Meta  `json:"meta"`    //Метаданные
	Id      string `json:"id"`      // Идентификатор шаблона
	Name    string `json:"name"`    // Наименование шаблона
	Type    string `json:"type"`    // Тип шаблона (entity - документ)
	Content string `json:"content"` // Ссылка на скачивание
}

// ListResponse Структура для ответа в виде списка сущностей
type ListResponse struct {
	Meta    *Meta                    `json:"meta"`    // Метаданные
	Context Meta                     `json:"context"` // Информация о сотруднике сделавшем запрос
	Rows    []map[string]interface{} `json:"rows"`    // Массив сущностей, интерфейс тут для того чтобы не захламлять код подобными структурами для каждой сущности
}
