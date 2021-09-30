package entity

import "github.com/Vitokz/MoySkladSDK/models/general"

//Group Отдел
type Group struct {
	Meta      *general.Meta `json:"meta"`      // Метаданные
	Id        string        `json:"id"`        // Id отдела
	AccountId string        `json:"accountId"` // Id учетной записи
	Name      string        `json:"name"`      // Наимнование отдела
	Index     int           `json:"index"`     // Порядковый номер в списке отделов
}