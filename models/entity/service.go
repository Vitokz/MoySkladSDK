package entity

import (
	"github.com/Vitokz/MoySkladSDK/models/general"
	"time"
)

//Service Услуга
type Service struct {
	Meta                *general.Meta       `json:"meta"`                          // Метаданные
	Id                  string              `json:"id"`                            // ID Услуги (Только для чтения)
	AccountId           string              `json:"accountId"`                     // ID учетной записи (Только для чтения)
	Owner               *Employee           `json:"owner"`                         // Метаданные владельца (Сотрудника)
	Shared              bool                `json:"shared"`                        // Общий доступ
	Group               *Group              `json:"group"`                         // Метаданные отдела сотрудника
	SyncId              string              `json:"syncId,omitempty"`              // ID синхронизации
	Updated             time.Time           `json:"updated"`                       // Момент последнего обновления сущности (Только для чтения)
	Name                string              `json:"name"`                          // Наименование Услуги
	Description         string              `json:"description,omitempty"`         // Описание Услуги
	Code                string              `json:"code"`                          // Код Услуги
	ExternalCode        string              `json:"externalCode"`                  // Внешний код Услуги
	Archived            bool                `json:"archived"`                      // Добавлена ли Услуга в архив
	PathName            string              `json:"pathName"`                      // Наименование группы, в которую входит Услуга (Только для чтения)
	Vat                 int                 `json:"vat,omitempty"`                 // НДС %
	VatEnabled          bool                `json:"vatEnabled,omitempty"`          // Включен ли НДС для группы
	UseParentVat        bool                `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы
	EffectiveVat        int                 `json:"effectiveVat,omitempty"`        // Реальный НДС % (Только для чтения)
	EffectiveVatEnabled int                 `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС
	ProductFolder       *ProductFolder      `json:"productFolder,omitempty"`       // Метаданные группы Услуги
	Uom                 *general.Uom        `json:"uom,omitempty"`                 // Единицы измерения
	MinPrice            *general.MinPrice   `json:"minPrice,omitempty"`            // Минимальная цена
	SalePrices          []general.SalePrice `json:"salePrices,omitempty"`          // Цены продажи
	BuyPrice            *general.BuyPrice   `json:"buyPrice,omitempty"`            // Закупочная цена
	Barcodes            []general.Barcode   `json:"barcodes,omitempty"`            // Штрихкоды
	Attributes          []general.Attribute `json:"attributes,omitempty"`          // Массив метаданных доп. полей
	DiscountProhibited  bool                `json:"discountProhibited,omitempty"`  // Признак запрета скидок
	PaymentItemType     string              `json:"paymentItemType,omitempty"`     // Признак предмета расчета

	// --
	// Значения поля paymentItemType.
	// GOOD	Товар
	// EXCISABLE_GOOD	Подакцизный товар
	// COMPOUND_PAYMENT_ITEM	Составной предмет расчета
	// ANOTHER_PAYMENT_ITEM	Иной предмет расчета
	// --
	TaxSystem string `json:"taxSystem,omitempty"` // Код системы налогообложения

	// --
	// Значения поля taxSystem.
	// TAX_SYSTEM_SAME_AS_GROUP	Совпадает с группой
	// GENERAL_TAX_SYSTEM	ОСН
	// SIMPLIFIED_TAX_SYSTEM_INCOME	УСН. Доход
	// SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME	УСН. Доход-Расход
	// UNIFIED_AGRICULTURAL_TAX	ЕСХН
	// PRESUMPTIVE_TAX_SYSTEM	ЕНВД
	// PATENT_BASED	Патент
	// --
	Files []general.File `json:"files,omitempty"` // Метаданные массива файлов
}
