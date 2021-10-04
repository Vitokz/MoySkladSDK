package entity

type Task struct {
	Meta              *Meta         `json:"meta"`              // Метаданные
	Id                string        `json:"id"`                // ID задачи
	AccountId         string        `json:"accountId"`         // ID учетной записи
	Author            *Employee     `json:"author"`            // Метаданные Сотрудника, создавшего задачу (администратор аккаунта, если автор - Приложение)
	AuthorApplication Meta          `json:"authorApplication"` //Метаданные Приложения, создавшего задачу
	Created           string        `json:"created"`           // Момент создания
	Updated           string        `json:"updated"`           // Момент последнего обновления
	Description       string        `json:"description"`       // Текст задачи
	DueToDate         string        `json:"dueToDate"`         // Срок задачи
	Assignee          Meta          `json:"assignee"`          // Метаданные ответственного за выполнение задачи
	Done              bool          `json:"done"`              // Отметка о выполнении задачи
	Completed         string        `json:"completed"`         // Время выполнения задачи
	Implementer       *Employee     // Метаданные Сотрудника, выполнившего задачу
	Agent             *Counterparty `json:"agent"`     // Метаданные Контрагента или юрлица, связанного с задачей. Задача может быть привязана либо к конрагенту, либо к юрлицу, либо к документу
	Operation         Meta          `json:"operation"` //Метаданные Документа, связанного с задачей. Задача может быть привязана либо к конрагенту, либо к юрлицу, либо к документу //TODO Поменять вместо Meta структуру документа
	Notes             *Note         `json:"notes"`     //Метаданные комментария к задаче
	Files             File          `json:"files"`     // Метаданные массива файлов
}
