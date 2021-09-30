package entity

import (
	"github.com/Vitokz/MoySkladSDK/models/general"
	"time"
)

// Store Склад
type Store struct {
	Meta         *general.Meta       `json:"meta,omitempty"`         // Метаданные Склада
	Id           string              `json:"id,omitempty"`           // ID Склада (Только для чтения)
	AccountId    string              `json:"accountId,omitempty"`    // ID учетной записи (Только для чтения)
	Owner        *Employee           `json:"owner,omitempty"`        // Владелец (Сотрудник)
	Shared       bool                `json:"shared,omitempty"`       // Общий доступ
	Group        *Group              `json:"group,omitempty"`        // Отдел сотрудника
	Updated      time.Time           `json:"updated,omitempty"`      // Момент последнего обновления Склада (Только для чтения)
	Name         string              `json:"name,omitempty"`         // Наименование Склада
	Description  string              `json:"description,omitempty"`  // Комментарий к Складу
	Code         string              `json:"code,omitempty"`         // Код Склада
	ExternalCode string              `json:"externalCode,omitempty"` // Внешний код Склада (Только для чтения)
	Archived     bool                `json:"archived,omitempty"`     // Добавлен ли Склад в архив
	Address      string              `json:"address,omitempty"`      // Адрес склада
	AddressFull  *AddressFull        `json:"addressFull,omitempty"`  // Адрес с детализацией по отдельным полям
	Parent       *Store              `json:"parent,omitempty"`       // Метаданные родительского склада (Группы)
	PathName     string              `json:"pathName,omitempty"`     // Группа Склада
	Attributes   []general.Attribute `json:"attributes,omitempty"`   // Массив метаданных дополнительных полей склада
}

// RetailStore Точка продаж
type RetailStore struct {
	Meta                 *general.Meta      `json:"meta,omitempty"`                 // Метаданные Точки продаж
	Id                   string             `json:"id,omitempty"`                   // ID Точки продаж (Только для чтения)
	AccountId            string             `json:"accountId,omitempty"`            // ID учетной записи (Только для чтения)
	Owner                *Employee          `json:"owner,omitempty"`                // Владелец (Сотрудник)
	Shared               bool               `json:"shared,omitempty"`               // Общий доступ
	Group                *Group             `json:"group,omitempty"`                // Отдел сотрудника
	Updated              time.Time          `json:"updated,omitempty"`              // Момент последнего обновления Точки продаж (Только для чтения)
	Name                 string             `json:"name,omitempty"`                 // Наименование Точки продаж
	Description          string             `json:"description,omitempty"`          // Комментарий к Точке продаж
	ExternalCode         string             `json:"externalCode,omitempty"`         // Внешний код Точки продаж (Только для чтения)
	Archived             bool               `json:"archived,omitempty"`             // Добавлена ли Точка продаж в архив
	Address              string             `json:"address,omitempty"`              // Адрес Точки продаж
	AddressFull          *AddressFull       `json:"addressFull,omitempty"`          // Адрес с детализацией по отдельным полям
	ControlShippingStock bool               `json:"controlShippingStock,omitempty"` // Контроль остатков. Не может быть true, если AllowCreateProducts имеет значение true
	OnlyInStock          bool               `json:"onlyInStock,omitempty"`          // Выгружать только товары в наличии. Доступно только при активном контроле остатков. Влияет только на выгрузку остатков в POS API
	Active               bool               `json:"active,omitempty"`               // Состояние точки продаж (Включена/Отключена)
	ControlCashierChoice bool               `json:"controlCashierChoice,omitempty"` // Выбор продавца
	DiscountEnable       bool               `json:"discountEnable,omitempty"`       // Разрешить скидки
	DiscountMaxPercent   int                `json:"discountMaxPercent,omitempty"`   // Максимальная скидка (в процентах)
	PriceType            *general.PriceType `json:"priceType,omitempty"`            // Тип цены, с которыми будут продаваться товары в рознице
	Cashiers             []general.Cashier  `json:"cashiers,omitempty"`             // Метаданные Кассиров
	Organization         *Organization      `json:"organization,omitempty"`         // Метаданные Юрлица
	Store                *Store             `json:"store,omitempty"`                // Метаданные Склада
	Acquire              *Counterparty      `json:"acquire,omitempty"`              // Метаданные Банка-эквайера
	BankPercent          int                `json:"bankPercent,omitempty"`          // Комиссия банка-эквайера (в процентах)
	IssueOrders          bool               `json:"issueOrders,omitempty"`          // Выдача заказов
	SellReserves         bool               `json:"sellReserves,omitempty"`         // Учет резервов
	LastOperationNames   []LastOperation    `json:"lastOperationNames,omitempty"`   // Последние операции (Только для чтения)
	OfdEnabled           bool               `json:"ofdEnabled,omitempty"`           // Отправлять электронный чек через ОФД (Только для чтения)
	PriorityOfdSend      string             `json:"priorityOfdSend,omitempty"`      // Приоритет отправки электронного чека. Активен только, когда отправка электронных чеков через ОФД включена.
	// --
	// priorityOfdSend
	// phone	Приоритет отправки на телефон
	// email	Приоритет отправки на e-mail
	// none	Отсутствие отправки чека
	// --
	AllowCustomPrice    bool              `json:"allowCustomPrice,omitempty"`    // Разрешить продажу по свободной цене (Только для чтения)
	AuthTokenAttached   bool              `json:"authTokenAttached,omitempty"`   // Создан ли токен для точки продаж (Только для чтения)
	OrderToState        *general.State    `json:"orderToState,omitempty"`        // Метаданные статуса, который проставится заказу после проведения продажи на его основании (если указано)
	CustomerOrderStates []general.State   `json:"customerOrderStates,omitempty"` // Метаданные статусов, в которых выгружаются заказы в точку продаж (если указано)
	Environment         *Environment      `json:"environment,omitempty"`         // Информация об окружении. Подробнее тут	Только для чтения	да
	State               *stateRetailStore `json:"state,omitempty"`               // Информация статусе точки продаж (Только для чтения)
	DefaultTaxSystem    string            `json:"defaultTaxSystem,omitempty"`    // Код системы налогообложения по умолчанию
	// --
	// Значения поля defaultTaxSystem.
	// GENERAL_TAX_SYSTEM	ОСН
	// SIMPLIFIED_TAX_SYSTEM_INCOME	УСН. Доход
	// SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME	УСН. Доход-Расход
	// UNIFIED_AGRICULTURAL_TAX	ЕСХН
	// PRESUMPTIVE_TAX_SYSTEM	ЕНВД
	// PATENT_BASED	Патент
	// --

	OrderTaxSystem string `json:"orderTaxSystem,omitempty"` // Код системы налогообложения для заказов
	// --
	// Значения поля orderTaxSystem.
	// GENERAL_TAX_SYSTEM	ОСН
	// SIMPLIFIED_TAX_SYSTEM_INCOME	УСН. Доход
	// SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME	УСН. Доход-Расход
	// UNIFIED_AGRICULTURAL_TAX	ЕСХН
	// PRESUMPTIVE_TAX_SYSTEM	ЕНВД
	// PATENT_BASED	Патент
	// --

	DemandPrefix               string `json:"demandPrefix,omitempty"`               // Префикс номера продаж
	AllowSellTobaccoWithoutMRC bool   `json:"allowSellTobaccoWithoutMRC,omitempty"` // Разрешить продавать табачную продукцию не по МРЦ
	TobaccoMrcControlType      string `json:"tobaccoMrcControlType"`                //Контроль МРЦ для табачной продукции
	// --
	// tobaccoMrcControlType
	// USER_PRICE	Не контролировать МРЦ
	// MRC_PRICE	Продавать по МРЦ указанной на пачке
	// SAME_PRICE	Запрещать продажу, если цена продажи не совпадает с МРЦ
	// --
	MarkingSellingMode string `json:"markingSellingMode"` // Режим продалжи маркированной продукции
	// --
	// markingSellingMode
	// CORRECT_MARKS_ONLY	Только с правильными кодами маркировки
	// WITHOUT_ERRORS	С правильными кодами и те, которые не удалось проверить
	// ALL	Все – независимо от результатов проверки кодов маркировки
	// --
	AllowCreateProducts bool            `json:"allowCreateProducts,omitempty"` // Контроль остатков. Не может быть true, если controlShippingStock имеет значение true
	ProductFolders      []ProductFolder `json:"productFolders,omitempty"`      // Коллекция Метаданных групп товаров, из которых можно выгружать товары
	// createAgentsTags	Array(Meta)	Коллекция групп покупателей, представленных в формате строк. Определяет группы, в которые добавляются новые покупатели. Значения null игнорируются	—	нет
	// filterAgentsTags	Array(Meta)	Коллекция групп покупателей, представленных в формате строк. Определяет группы, из которых выгружаются покупатели. Значения null игнорируются	—	нет
	PrintAlways                         bool                      `json:"printAlways,omitempty"`                         // Всегда печатать кассовые чеки
	ReceiptTemplate                     *general.EmbeddedTemplate `json:"receiptTemplate,omitempty"`                     // Метаданные шаблона печати кассовых чеков
	CreatePaymentInOnRetailShiftClosing bool                      `json:"createPaymentInOnRetailShiftClosing,omitempty"` // Создавать входящий платеж при закрытии смены
	CreateCashInOnRetailShiftClosing    bool                      `json:"createCashInOnRetailShiftClosing,omitempty"`    // Создавать ПКО при закрытии смены
	ReturnFromClosedShiftEnabled        bool                      `json:"returnFromClosedShiftEnabled,omitempty"`        // Разрешить возвраты в закрытых сменах
	EnableReturnsWithNoReason           bool                      `json:"enableReturnsWithNoReason,omitempty"`           // Разрешить возвраты без основания
	CreateOrderWithState                *general.State            `json:"createOrderWithState,omitempty"`                // Метаданные статуса, который будет указан при создании заказа
	ReservePrepaidGoods                 bool                      `json:"reservePrepaidGoods,omitempty"`                 // Резервировать товары, за которые внесена предоплата
	FiscalType                          string                    `json:"fiscalType,omitempty"`                          // Тип формирования чеков
	// --
	// fiscalType
	// STANDARD	Стандартное
	// MASTER	Стандартное с обработкой облачных операций
	// CLOUD	Облачное
	// --
	MinionToMasterType string `json:"minionToMasterType,omitempty"` // Стратегия выбора кассы для фискализации облачных чеков
	// --
	// ANY	Любая мастер касса
	// SAME_GROUP	Только кассы из того же отдела
	// CHOSEN	Выбранные кассы из списка в поле masterRetailStores
	// --
	MasterRetailStores []*RetailStore `json:"masterRetailStores,omitempty"` // Ссылка на точки продаж, которые могут фискализировать операции с текущей точки продаж, если minionToMaster = CHOSEN
	QrAcquire          *general.Meta  `json:"qrAcquire"`                    // Метаданные банка-эквайера по операциям по QR-коду
	QrBankPercent      float64        `json:"qrBankPercent"`                // Комиссия банка-эквайера по операциям по QR-коду
	QrPayEnabled       bool           `json:"qrPayEnabled"`                 // Возможность продаж по QR-коду на точке продаж
	IdQR               string         `json:"idQR"`                         // Идентификатор устройства QR для приложения оплаты по QR
	QrTerminalId       string         `json:"qrTerminalId"`                 // Идентификатор терминала для приложения по QR
}

// LastOperation Последние операции
type LastOperation struct {
	Entity string `json:"entity,omitempty"` // Ключевое слово, обозначающее тип последней операции (Только для чтения)
	Name   string `json:"name,omitempty"`   // Наименование (номер) последней операции (Только для чтения)
}

type Environment struct {
	Device          string    `json:"device,omitempty"`          // Информация об устройстве
	OS              string    `json:"os,omitempty"`              // Информация об операционной системе
	Software        *software `json:"software,omitempty"`        // Информация о ПО
	ChequePrinter   *kkt      `json:"chequePrinter,omitempty"`   // Данные о ККТ
	PaymentTerminal string    `json:"paymentTerminal,omitempty"` // Информация о платежном терминале
}

type stateRetailStore struct {
	Sync            *sync            `json:"sync,omitempty"`            // Состояние синхронизации
	LastCheckMoment string           `json:"lastCheckMoment,omitempty"` // Дата и время последней синхронизации
	FiscalMemory    *fiscalMemory    `json:"fiscalMemory,omitempty"`    // Информация о фискальном накопителе
	PaymentTerminal *paymentTerminal `json:"paymentTerminal,omitempty"` // Информация о платежном терминале
}

type paymentTerminal struct {
	AcquiringType string `json:"acquiringType,omitempty"` // Информация о типе эквайера (например: inpas/payme)
}

type sync struct {
	Message          string `json:"message,omitempty"`          // Состояние синхронизации
	LastAttempMoment string `json:"lastAttempMoment,omitempty"` // Дата последней сихронизации (не обязательно успешной)
}

type fiscalMemory struct {
	Error           errorF `json:"error,omitempty"`           // Информация об ошибке ФН
	NotSendDocCount int    `json:"notSendDocCount,omitempty"` // Количество неотправленных документов в ОФД
}

type errorF struct {
	Code    int    `json:"сode,omitempty"`    // Код ошибки ФН
	Message string `json:"message,omitempty"` // Описание ошибки
}

type software struct {
	Name    string `json:"name,omitempty"`    // Наименование ПО
	Vendor  string `json:"vendor,omitempty"`  // Производитель
	Version string `json:"version,omitempty"` // Версия ПО
}

type kkt struct {
	Vendor            string        `json:"vendor,omitempty"`            // Производитель
	Name              string        `json:"name,omitempty"`              // Наименование ПО
	Serial            string        `json:"serial,omitempty"`            // Серийный номер
	FiscalDataVersion string        `json:"fiscalDataVersion,omitempty"` // Формат фискальных данных
	Driver            *driver       `json:"driver,omitempty"`            // Информация об используемом драйвере
	FiscalMemory      *fiscalDevice `json:"fiscalMemory,omitempty"`      // Информация о фискальном накопителе
	FirmwareVersion   string        `json:"firmwareVersion,omitempty"`   // Версия прошивки ККТ
}

type driver struct {
	Name    string `json:"name,omitempty"`    // Наименование драйвера
	Version string `json:"version,omitempty"` // Версия драйвера
}

type fiscalDevice struct {
	FiscalDataVersion  string `json:"fiscalDataVersion,omitempty"`  // Версия фискальной памяти
	FiscalValidityDate string `json:"fiscalValidityDate,omitempty"` // Версия фискальной памяти
}
