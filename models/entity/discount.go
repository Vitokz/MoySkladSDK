package entity

type Discount struct {
	Meta        *Meta    `json:"meta"`        // Метаданные
	Id          string   `json:"id"`          // Id контактного
	AccountId   string   `json:"accountId"`   // Id учетной записи
	Name        string   `json:"name"`        // Наименование скидки
	Active      bool     `json:"active"`      // Индикатор является ли скидка активной
	AllProducts bool     `json:"allProducts"` // Индикатор действует ли скидка на все товары
	AgentTags   []string `json:"agentTags"`   // Тэги контрагентов
	Assortment  []Meta   `json:"assortment"`  // Массив метаданных товаров и услуг
}
