package entity

type File struct {
	Meta      *Meta     `json:"meta"`      // Метаданные
	Title     string    `json:"title"`     // Название Файла
	Filename  string    `json:"filename"`  // Имя файла
	Size      int       `json:"size"`      // Размер файла в байтах
	Created   string    `json:"created"`   // Время загрузки файла на сервер
	CreatedBy *Employee `json:"createdBy"` // Метаданные сотрудника
	Miniature *Meta     `json:"miniature"` // Метаданные миниатюры изображения
	Tiny      *Meta     `json:"tiny"`      // Метаданные уменьшенного изображения
}
