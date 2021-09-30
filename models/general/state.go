package general

type StateType string

const (
	Regular      StateType = "Regular"      // Обычный (значение по умолчанию)
	Successful   StateType = "Successful"   // Финальный положительный
	Unsuccessful StateType = "Unsuccessful" // Финальный отрицательный
)

// State Статус договра
type State struct {
	Meta       *Meta `json:"meta,omitempty"`       // Метаданные Статуса (Только для чтения)
	Id         string        `json:"id,omitempty"`         // ID Статуса (Только для чтения)
	AccountId  string        `json:"accountId,omitempty"`  // ID учетной записи (Только для чтения)
	Name       string        `json:"name,omitempty"`       // Наименование Статуса (Необходимое при создании)
	Color      int           `json:"color,omitempty"`      // Цвет Статуса (Необходимое при создании)
	StateType  StateType     `json:"stateType,omitempty"`  // Тип Статуса (Необходимое при создании) !! Дефолтное значение - Regular
	EntityType string        `json:"entityType,omitempty"` // Тип сущности, к которой относится Статус (ключевое слово в рамках JSON API) (Только для чтения)
}
