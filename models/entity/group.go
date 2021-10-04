package entity

//Group Отдел
type Group struct {
	Meta      *Meta  `json:"meta"`      // Метаданные
	Id        string `json:"id"`        // Id отдела
	AccountId string `json:"accountId"` // Id учетной записи
	Name      string `json:"name"`      // Наимнование отдела
	Index     int    `json:"index"`     // Порядковый номер в списке отделов
}
