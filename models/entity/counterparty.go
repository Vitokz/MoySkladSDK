package entity

//Counterparty Контрагент
type Counterparty struct {
	Meta               *Meta           `json:"meta,omitempty"`               // Метаданные
	Id                 string          `json:"id,omitempty"`                 // Id контрагента
	AccountId          string          `json:"account_id,omitempty"`         // Id аккаунта
	Owner              *Employee       `json:"owner,omitempty"`              // Владелец
	Shared             bool            `json:"shared,omitempty"`             // Общий доступ
	Group              *Group          `json:"group,omitempty"`              // Отдел сотрудника
	SyncId             string          `json:"syncId,omitempty"`             // Синхронизация
	Updated            string          `json:"updated,omitempty"`            // Момент последнего обновления
	Name               string          `json:"name,omitempty"`               // Наимнование Контрагента
	Description        string          `json:"description,omitempty"`        // Комментарий к контрагенту
	Code               string          `json:"code,omitempty"`               // Код контрагента
	ExternalCode       string          `json:"externalCode,omitempty"`       // Внешний код контрагента
	Archived           bool            `json:"archived,omitempty"`           // Добавлен ли контрагент в архив
	Created            string          `json:"created,omitempty"`            // Момент создания
	Email              string          `json:"email,omitempty"`              // Адрес электронной почты
	Phone              string          `json:"phone,omitempty"`              // Номер телефона
	Fax                string          `json:"fax,omitempty"`                // Номер факса
	ActualAddress      string          `json:"actualAddress,omitempty"`      // Фактический адрес контрагента
	ActualAddressFull  *AddressFull    `json:"actualAddressFull,omitempty"`  // Фактический адрес контрагента с доп. полями
	Accounts           []Account       `json:"accounts,omitempty"`           // Массив счетов контрагента
	CompanyType        string          `json:"companyType,omitempty"`        // Тип контрагента.
	DiscountCardNumber string          `json:"discountCardNumber,omitempty"` // Номер дисконтной карты
	State              *State          `json:"state,omitempty"`              // Метаданные статуса контрагента
	SalesAmount        int             `json:"salesAmount,omitempty"`        // Сумма продаж
	BonusProgram       *BonusProgram   `json:"bonusProgram,omitempty"`       // Метаданные активной бонусной программы
	BonusPoints        int             `json:"bonus_points,omitempty"`       // Бонусные баллы по активной бонусной программе
	Files              File            `json:"files,omitempty"`              // Метаданные массива файлов
	LegalType          string          `json:"legalType,omitempty"`          // Полное наименование контрагента
	LegalLastName      string          `json:"legalLastName,omitempty"`      // Фамилия для контрагента
	LegalFirstName     string          `json:"legalFirstName,omitempty"`     // Имя для контрагента
	LegalMiddleName    string          `json:"legalMiddleName,omitempty"`    // Отчество для контрагента
	LegalAddress       string          `json:"legalAddress,omitempty"`       // Юридический адрес Контрагента
	LegalAddressFull   *AddressFull    `json:"legalAddressFull,omitempty"`   // Юридический адрес Контрагента с детализацией по отдельным полям
	Inn                string          `json:"inn,omitempty"`                // ИНН
	Kpp                string          `json:"kpp,omitempty"`                // КПП
	Ogrn               string          `json:"ogrn,omitempty"`               // ОГРН
	Ogrnip             string          `json:"ogrnip,omitempty"`             // ОГРНИП
	Okpo               string          `json:"okpo,omitempty"`               // ОКПО
	CertificateNumber  string          `json:"certificateNumber,omitempty"`  // Номер свидетельства
	CertificateDate    string          `json:"certificateDate,omitempty"`    // Дата свидетельства
	Tags               []string        `json:"tags,omitempty"`               // Группы контрагента
	ContactPersons     []ContactPerson `json:"contactpersons,omitempty"`     // Массив контактных лиц фирмы контрагента
	Attributes         []Attribute     `json:"attributes,omitempty"`         // Массив метаданных доп. полей
	Discounts          []Discount      `json:"discounts,omitempty"`          // Массив метаданных скидок
	Notes              []Note          `json:"notes,omitempty"`              // Массив событий контрагента
	PriceType          *PriceType      `json:"priceType,omitempty"`          // Тип цены Контрагента
}

//AddressFull Полный адрес
type AddressFull struct {
	PostalCode string   `json:"postalCode,omitempty"` // Почтовый индекс
	Country    *Country `json:"country,omitempty"`    // Метаданные страны
	Region     *Region  `json:"region,omitempty"`     // Метаданные региона
	City       string   `json:"city,omitempty"`       // Город
	Street     string   `json:"street,omitempty"`     // Улица
	House      string   `json:"house,omitempty"`      // Дом
	Apartment  string   `json:"apartment,omitempty"`  // Квартира
	AddInfo    string   `json:"addInfo,omitempty"`    // Другое
	Comment    string   `json:"comment,omitempty"`    // Комментарий
}

//Account Счета Контрагента
type Account struct {
	Meta                 *Meta  `json:"meta,omitempty"`
	ID                   string `json:"id,omitempty"`                   // ID Счета (Только для чтения)
	AccountID            string `json:"accountId,omitempty"`            // ID учетной записи (Только для чтения)
	Updated              string `json:"updated,omitempty"`              // Момент последнего обновления юрлица (Только для чтения)
	IsDefault            bool   `json:"isDefault,omitempty"`            // Является ли счет основным счетом юрлица
	AccountNumber        string `json:"accountNumber,omitempty"`        // Номер счета	Необходимое при создании
	BankName             string `json:"bankName,omitempty"`             // Наименование банка
	BankLocation         string `json:"bankLocation,omitempty"`         // Адрес банка
	CorrespondentAccount string `json:"correspondentAccount,omitempty"` // Корр счет
	BIC                  string `json:"bic,omitempty"`                  // БИК
}

//ContactPerson Контактный список контрагента
type ContactPerson struct {
	Meta         *Meta         `json:"meta,omitempty"`         // Метаданные
	Id           string        `json:"id,omitempty"`           // Id контактного
	AccountId    string        `json:"accountId,omitempty"`    // Id учетной записи
	Updated      string        `json:"updated,omitempty"`      // Момент последнего обновления
	Name         string        `json:"name,omitempty"`         // ФИО контактного лица
	Description  string        `json:"description,omitempty"`  // Описание контактного лица
	ExternalCode string        `json:"externalCode,omitempty"` // Внешний код контактного лица
	Email        string        `json:"email,omitempty"`        // Адрес электронной почты контактного лица
	Phone        string        `json:"phone,omitempty"`        // Номер телефона контакного лица
	Position     string        `json:"position,omitempty"`     // Должность контактного лица
	Agent        *Counterparty `json:"agent,omitempty"`        // Метаданные контрагента
}

//Note Событие
type Note struct {
	Meta              *Meta         `json:"meta,omitempty"`              // Метаданные
	Id                string        `json:"id,omitempty"`                // Id контактного
	AccountId         string        `json:"accountId,omitempty"`         // Id учетной записи
	Created           string        `json:"created,omitempty"`           // Время создания
	Description       string        `json:"description,omitempty"`       // Описание контактного лица
	Agent             *Counterparty `json:"agent,omitempty"`             // Метаданные контрагента
	Author            *Employee     `json:"author,omitempty"`            // Метаданные сотрудника-создателя события
	AuthorApplication *Meta         `json:"authorApplication,omitempty"` // Метаданные приложения
}
