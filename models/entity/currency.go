package entity

import "github.com/Vitokz/MoySkladSDK/models/general"

//Currency Валюта
type Currency struct {
	Meta           *general.Meta `json:"meta,omitempty"`           // Метаданные
	Id             string        `json:"id,omitempty"`             // ID валюты
	Name           string        `json:"name,omitempty"`           // Наименование валюты
	FullName       string        `json:"fullName,omitempty"`       // Полное наименование валюты
	Code           string        `json:"code,omitempty"`           // Цифровой код валюты
	IsoCode        string        `json:"isoCode,omitempty"`        // Буквенный код валюты
	Rate           float64       `json:"rate,omitempty"`           // Курс валюты
	Multiplicity   int           `json:"multiplicity,omitempty"`   // Кратность курса валюты
	Indirect       bool          `json:"indirect,omitempty"`       // Признак обратного курса валюты
	RateUpdateType bool          `json:"rateUpdateType,omitempty"` // Способ обновления курса валюты
	MajorUnit      *CurrencyUnit `json:"majorUnit"`                // Форма единиц целой части валюты
	MinorUnit      *CurrencyUnit `json:"minorUnit,omitempty"`      // Форма единиц дробной части Валюты.
	Archived       bool          `json:"archived,omitempty"`       // Добалена ли валюта в архив
	System         bool          `json:"system,omitempty"`         // Основана ли на валюте из системного справочника
	Default        bool          `json:"default,omitempty"`        // Является ли валюта дефолтной
}

// CurrencyUnit Форма единиц
type CurrencyUnit struct {
	Gender string `json:"gender,omitempty"` // Грамматический род единицы валюты (допустимые значения masculine - мужской, feminine - женский)
	S1     string `json:"s1,omitempty"`     // Форма единицы, используемая при числительном 1
	S2     string `json:"s2,omitempty"`     // Форма единицы, используемая при числительном 2
	S5     string `json:"s5,omitempty"`     // Форма единицы, используемая при числительном 5
}
