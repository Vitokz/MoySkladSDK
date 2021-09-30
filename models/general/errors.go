package general

//Error ошибка
type Error struct {
	Error        string `json:"error"`        // Заголовок ошибки
	Parameter    string `json:"parameter"`    // Параметр, на котором произошла ошибка
	Code         int    `json:"code"`         // Код ошибки
	ErrorMessage string `json:"errorMessage"` // Сообщение, прилагаемое к ошибке
}
