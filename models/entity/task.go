package entity

import (
	"github.com/Vitokz/MoySkladSDK/models/general"
	"time"
)

type Task struct {
	Meta *general.Meta `json:"meta"` // Метаданные
	Id string `json:"id"` // ID задачи
	AccountId string	`json:"accountId"` // ID учетной записи
	Author *Employee `json:"author"` // Метаданные Сотрудника, создавшего задачу (администратор аккаунта, если автор - Приложение)
	AuthorApplication general.Meta `json:"authorApplication"` //Метаданные Приложения, создавшего задачу
	Created time.Time `json:"created"` // Момент создания
	Updated time.Time `json:"updated"` // Момент последнего обновления
	Description string `json:"description"` // Текст задачи
	DueToDate time.Time `json:"dueToDate"` // Срок задачи
	Assignee general.Meta `json:"assignee"` // Метаданные ответственного за выполнение задачи
	Done bool `json:"done"` // Отметка о выполнении задачи
	Completed time.Time `json:"completed"` // Время выполнения задачи
	Implementer *Employee // Метаданные Сотрудника, выполнившего задачу
	Agent *Counterparty `json:"agent"` // Метаданные Контрагента или юрлица, связанного с задачей. Задача может быть привязана либо к конрагенту, либо к юрлицу, либо к документу
	Operation general.Meta `json:"operation"` //Метаданные Документа, связанного с задачей. Задача может быть привязана либо к конрагенту, либо к юрлицу, либо к документу //TODO Поменять вместо Meta структуру документа
	Notes *Note `json:"notes"` //Метаданные комментария к задаче
	Files []general.File `json:"files"` // Метаданные массива файлов
}