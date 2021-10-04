package entity

type Project struct {
	Meta         *Meta       `json:"meta"`                  // Метаданные Проекта
	Id           string      `json:"id"`                    // ID Проекта (Только для чтения)
	AccountId    string      `json:"accountId"`             // ID учетной записи (Только для чтения)
	Owner        *Employee   `json:"owner"`                 // Метаданные владельца (Сотрудника)
	Shared       bool        `json:"shared"`                // Общий доступ
	Group        *Group      `json:"group"`                 // Метаданные отдела сотрудника
	Updated      string      `json:"updated"`               // Момент последнего обновления сущности (Только для чтения)
	Name         string      `json:"name"`                  // Наименование Проекта
	Description  string      `json:"description,omitempty"` // Описание Проекта
	Code         string      `json:"code"`                  // Код Проекта
	ExternalCode string      `json:"externalCode"`          // Внешний код Проекта
	Archived     bool        `json:"archived"`              // Заархивирован ли проект
	Attributes   []Attribute `json:"attributes,omitempty"`  // Коллекция доп. полей
}
