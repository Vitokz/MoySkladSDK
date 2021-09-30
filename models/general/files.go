package general

import (
	"github.com/Vitokz/MoySkladSDK/models/entity"
	"time"
)

type File struct {
	Meta      *Meta            `json:"meta"`      // Метаданные
	Title     string           `json:"title"`     // Название Файла
	Filename  string           `json:"filename"`  // Имя файла
	Size      int              `json:"size"`      // Размер файла в байтах
	Created   time.Time        `json:"created"`   // Время загрузки файла на сервер
	CreatedBy *entity.Employee `json:"createdBy"` // Метаданные сотрудника
	Miniature *Meta            `json:"miniature"` // Метаданные миниатюры изображения
	Tiny      *Meta            `json:"tiny"`      // Метаданные уменьшенного изображения
}
