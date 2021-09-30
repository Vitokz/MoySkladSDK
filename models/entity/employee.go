package entity

import (
	"github.com/Vitokz/MoySkladSDK/models/general"
	"time"
)

type Employee struct {
	Meta         *general.Meta       `json:"meta"`                   // Метаданные Сотрудника
	ID           string              `json:"id,omitempty"`           // ID Сотрудника (Только для чтения)
	AccountID    string              `json:"accountId,omitempty"`    // ID учетной записи (Только для чтения)
	Owner        *Employee           `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       bool                `json:"shared,omitempty"`       // Общий доступ
	Group        *Group              `json:"group,omitempty"`        // Отдел сотрудника
	Updated      time.Time           `json:"updated,omitempty"`      // Момент последнего обновления Сотрудника (Только для чтения)
	Name         string              `json:"name,omitempty"`         // Наименование Сотрудника (Только для чтения)
	Description  string              `json:"description,omitempty"`  // Комментарий к Сотруднику
	ExternalCode string              `json:"externalCode,omitempty"` // Внешний код Сотрудника (Только для чтения)
	Archived     bool                `json:"archived,omitempty"`     // Добавлен ли Сотрудник в архив
	Created      time.Time           `json:"created,omitempty"`      // Момент создания Сотрудника (Только для чтения)
	UID          string              `json:"uid,omitempty"`          // Логин Сотрудника (Только для чтения)
	Email        string              `json:"email,omitempty"`        // Электронная почта сотрудника
	Phone        string              `json:"phone,omitempty"`        // Телефон сотрудника
	FirstName    string              `json:"firstName,omitempty"`    // Имя
	MiddleName   string              `json:"middleName,omitempty"`   // Отчество
	LastName     string              `json:"lastName,omitempty"`     // Фамилия
	FullName     string              `json:"fullName,omitempty"`     // Имя Отчество Фамилия (Только для чтения)
	ShortFio     string              `json:"shortFio,omitempty"`     // Краткое ФИО (Только для чтения)
	Cashiers     []general.Cashier   `json:"cashiers,omitempty"`     // Массив кассиров (Только для чтения)
	Attributes   []general.Attribute `json:"attributes,omitempty"`   // Дополнительные поля Сотрудника
	Image        *general.Image      `json:"image,omitempty"`        // Фотография сотрудника
	INN          string              `json:"inn,omitempty"`          // ИНН сотрудника (в формате ИНН физического лица)
	Position     string              `json:"position,omitempty"`     // Должность сотрудника
}

// Organization структура сущности Юрлицо
type Organization struct {
	Meta              *general.Meta `json:"meta,omitempty"`              // Метаданные Юрлица
	Id                string        `json:"id,omitempty"`                // ID Юрлица (Только для чтения)
	AccountId         string        `json:"accountId,omitempty"`         // ID учетной записи (Только для чтения)
	Owner             *Employee     `json:"owner,omitempty"`             // Владелец (Сотрудник)
	Shared            bool          `json:"shared,omitempty"`            // Общий доступ
	Group             *Group        `json:"group,omitempty"`             // Отдел сотрудника
	SyncID            string        `json:"syncId,omitempty"`            // ID синхронизации
	Updated           time.Time     `json:"updated,omitempty"`           // Момент последнего обновления Юрлица (Только для чтения)
	Name              string        `json:"name,omitempty"`              // Наименование Юрлица
	Description       string        `json:"description,omitempty"`       // Комментарий к Юрлицу
	Code              string        `json:"code,omitempty"`              // Код Юрлица
	ExternalCode      string        `json:"externalCode,omitempty"`      // Внешний код Юрлица (Только для чтения)
	Achived           bool          `json:"archived,omitempty"`          // Добавлено ли Юрлицо в архив
	Created           time.Time     `json:"created,omitempty"`           // Дата создания
	ActualAddress     string        `json:"actualAddress,omitempty"`     // Фактический адрес Юрлица
	ActualAddressFull *AddressFull  `json:"actualAddressFull,omitempty"` // Фактический адрес Юрлица с детализацией по отдельным полям
	CompanyType       string        `json:"companyType,omitempty"`       // Тип Юрлица. В зависимости от значения данного поля набор выводимых реквизитов контрагента может меняться
	// --
	// companyType
	// legal	Юридическое лицо
	// entrepreneur	Индивидуальный предприниматель
	// individual	Физическое лицо
	// --
	TrackingContractNumber string              `json:"trackingContractNumber,omitempty"` // Номер договора с ЦРПТ
	TrackingContractDate   time.Time           `json:"trackingContractDate,omitempty"`   // Дата договора с ЦРПТ
	BonusProgram           *BonusProgram       `json:"bonusProgram,omitempty"`           // Метаданные активной бонусной программы
	BonusPoints            int                 `json:"bonusPoints,omitempty"`            // Бонусные баллы по активной бонусной программе
	LegalTitle             string              `json:"legalTitle,omitempty"`             // Полное наименование. Игнорируется, если передано одно из значений для ФИО. Формируется автоматически на основе получаемых ФИО Юрлица
	LegalLastName          string              `json:"legalLastName,omitempty"`          // Фамилия для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalFirstName         string              `json:"legalFirstName,omitempty"`         // Имя для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalMiddleName        string              `json:"legalMiddleName,omitempty"`        // Отчество для Юрлица типа [Индивидуальный предприниматель, Физическое лицо]. Игнорируется для Юрлиц типа [Юридическое лицо]
	LegalAddress           string              `json:"legalAddress,omitempty"`           // Юридический адреса Юрлица
	LegalAddressFull       *AddressFull        `json:"legalAddressFull,omitempty"`       // Юридический адрес Юрлица с детализацией по отдельным полям
	INN                    string              `json:"inn,omitempty"`                    // ИНН
	KPP                    string              `json:"kpp,omitempty"`                    // КПП
	OGRN                   string              `json:"ogrn,omitempty"`                   // ОГРН
	OGRNIP                 string              `json:"ogrnip,omitempty"`                 // ОГРНИП
	OKPO                   string              `json:"okpo,omitempty"`                   // ОКПО
	CertificateNumber      general.Meta        `json:"certificate_number"`               // Номер свидетельства
	CertificateDate        time.Time           `json:"certificateDate,omitempty"`        // Дата свидетельства
	Email                  string              `json:"email,omitempty"`                  // Адрес электронной почты
	Phone                  string              `json:"phone,omitempty"`                  // Номер городского телефона
	Fax                    string              `json:"fax,omitempty"`                    // Номер факса
	Accounts               *Account            `json:"accounts,omitempty"`               // Метаданные счетов юрлица
	Attributes             []general.Attribute `json:"attributes,omitempty"`             // Массив метаданных дополнительных полей юрлица
	IsEgaisEnable          bool                `json:"isEgaisEnable,omitempty"`          // Включен ли ЕГАИС для данного юрлица
	FSRARID                string              `json:"fsrarId,omitempty"`                // Идентификатор в ФСРАР
	PayerVat               bool                `json:"payerVat,omitempty"`               // Является ли данное юрлицо плательщиком НДС
	UTMURL                 string              `json:"utmUrl,omitempty"`                 // IP-адрес УТМ
	Director               string              `json:"director,omitempty"`               // Руководитель
	DirectorPosition       string              `json:"directorPosition"`                 // Позиция руководителя
	DirectorSign           *Sign               `json:"directorSign"`                     // Пдпись руководителя
	ChiefAccountant        string              `json:"chiefAccountant,omitempty"`        // Главный бухгалтер
	ChiefAccountSign       *Sign               `json:"chiefAccountSign"`                 // Подпись главного бухгалтера
	Stamp                  *Sign               `json:"stamp"`                            // Печать
}

// Accounts ...
type Accounts struct {
	Meta *general.Meta `json:"meta,omitempty"`
	Rows []Account     `json:"rows,omitempty"`
}

//Sign Подписка и печать
type Sign struct {
	Meta      *general.Meta `json:"meta"`      // Метаданные
	Title     string        `json:"title"`     // Название изображения
	Filename  string        `json:"filename"`  // Имя файла
	Size      int           `json:"size"`      // Размер файла в байтах
	Updated   time.Time     `json:"updated"`   // Время загрузки файла на сервер
	Miniature general.Meta  `json:"miniature"` // Метданые миниатюры изображения
}
