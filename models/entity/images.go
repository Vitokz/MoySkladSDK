package entity

// Image структура изображения (image)
type Image struct {
	Meta      *Meta  `json:"meta"`                // Метаданные объекта
	Title     string `json:"title,omitempty"`     // Название Изображения
	Filename  string `json:"filename,omitempty"`  // Имя файла
	Size      int    `json:"size,omitempty"`      // Размер файла в байтах
	Updated   string `json:"updated,omitempty"`   // Время последнего изменения
	Miniature *Meta  `json:"miniature,omitempty"` // Метаданные миниатюры изображения
	Tiny      *Meta  `json:"tiny,omitempty"`      // Метаданные уменьшенного изображения
}
