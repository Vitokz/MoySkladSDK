package entity

//Product Товар
type Product struct {
	Meta                *Meta          `json:"meta,omitempty"`                // Метаданные Товара
	Id                  string         `json:"id,omitempty"`                  // ID Товара (Только для чтения)
	AccountId           string         `json:"accountId,omitempty"`           // ID учетной записи (Только для чтения)
	Owner               *Employee      `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Shared              bool           `json:"shared,omitempty"`              // Общий доступ
	Group               *Group         `json:"group,omitempty"`               // Метаданные отдела сотрудника
	SyncId              string         `json:"syncId,omitempty"`              // ID синхронизации
	Updated             string         `json:"updated,omitempty"`             // Момент последнего обновления сущности (Только для чтения)
	Name                string         `json:"name,omitempty"`                // Наименование Товара
	Description         string         `json:"description,omitempty"`         // Описание Товара
	Code                string         `json:"code,omitempty"`                // Код Товара
	ExternalCode        string         `json:"externalCode,omitempty"`        // Внешний код Товара
	Archived            bool           `json:"archived,omitempty"`            // Добавлен ли Товар в архив
	PathName            string         `json:"pathName,omitempty"`            // Наименование группы, в которую входит Товар (Только для чтения)
	Vat                 int            `json:"vat,omitempty"`                 // НДС %
	VatEnabled          bool           `json:"vatEnabled,omitempty"`          // Включен ли НДС для группы
	UseParentVat        bool           `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы
	EffectiveVat        int            `json:"effectiveVat,omitempty"`        // Реальный НДС % (Только для чтения)
	EffectiveVatEnabled int            `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`       // Метаданные группы Товара
	Uom                 *Uom           `json:"uom,omitempty"`                 // Единицы измерения
	Images              *Images        `json:"images,omitempty"`              // Изображения
	MinPrice            *MinPrice      `json:"minPrice,omitempty"`            // Минимальная цена
	SalePrices          []SalePrice    `json:"salePrices,omitempty"`          // Цены продажи
	BuyPrice            *BuyPrice      `json:"buyPrice,omitempty"`            // Закупочная цена
	Supplier            *Counterparty  `json:"supplier,omitempty"`            // Метаданные контрагента-поставщика
	Attributes          []Attribute    `json:"attributes,omitempty"`          // Коллекция доп. полей
	Country             *Country       `json:"country,omitempty"`             // Метаданные Страны
	Article             string         `json:"article,omitempty"`             // Артикул
	Weight              float64        `json:"weight,omitempty"`              // Вес
	Volume              float64        `json:"volume,omitempty"`              // Объем
	Packs               []Pack         `json:"packs,omitempty"`               // Упаковки Товара
	Alcoholic           *Alcoholic     `json:"alcoholic,omitempty"`           // Объект, содержащий поля алкогольной продукции
	VariantsCount       int            `json:"variantsCount,omitempty"`       // Количество модификаций у данного товара (Только для чтения)
	MinimumBalance      int            `json:"minimumBalance,omitempty"`      // Неснижаемый остаток
	IsSerialTrackable   bool           `json:"isSerialTrackable,omitempty"`   // Учет по серийным номерам. Не может быть указан вместе с alcoholic и weighed
	Things              []string       `json:"things,omitempty"`              // Серийные номера
	Barcodes            []Barcode      `json:"barcodes,omitempty"`            // Штрихкоды
	DiscountProhibited  bool           `json:"discountProhibited,omitempty"`  // Признак запрета скидок
	Tnved               string         `json:"tnved,omitempty"`               // Код ТН ВЭД
	PartialDisposal     bool           `json:"partialDisposal,omitempty"`     // Управление состоянием частичного выбытия маркированного товара. "true" возможность включена
	TrackingType        string         `json:"trackingType,omitempty"`        // Тип маркируемой продукции
	// --
	// Значения поля trackingType.
	// NOT_TRACKED	Без маркировки
	// TOBACCO	Тип маркировки "Табак"
	// SHOES	Тип маркировки "Обувь"
	// LP_CLOTHES	Тип маркировки "Одежда"
	// LP_LINENS	Тип маркировки "Постельное белье"
	// PERFUMERY	Духи и туалетная вода
	// ELECTRONICS	Фотокамеры и лампы-вспышки
	// --

	PaymentItemType string `json:"paymentItemType,omitempty"` // Признак предмета расчета

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

	PPEType string `json:"ppeType,omitempty"` // Код вида номенклатурной классификации медицинских средств индивидуальной защиты (EAN-13)
	// --
	// Значения поля ppeType.
	// 2400001323807	маска лицевая для защиты дыхательных путей, многоразового использования
	// 2400003675805	маска лицевая для защиты дыхательных путей, одноразового использования
	// 2400001807703	респиратор общего применения
	// 2400001818303	респиратор хирургический
	// 2400002186203	респиратор хирургический антибактериальный
	// 2400001368105	средство назальное для защиты от загрязненного воздуха, местного действия
	// 2400001225408	перчатки смотровые (процедурные) из латекса гевеи, неопудренные, нестерильные
	// 2400001225606	перчатки смотровые (процедурные) из латекса гевеи, опудренные
	// 2400001226108	перчатки смотровые (процедурные) из латекса гевеи, неопудренные, стерильные
	// 2400001393503	перчатки смотровые (процедурные) из полихлоропрена, неопудренные
	// 2400001858309	перчатки смотровые (процедурные) нитриловые, неопудренные, нестерильные
	// 2400001858507	перчатки смотровые (процедурные) нитриловые, опудренные
	// 2400002052805	перчатки смотровые (процедурные) виниловые, неопудренные
	// 2400002052904	перчатки смотровые (процедурные) виниловые, опудренные
	// 2400002984502	перчатки смотровые (процедурные) из гваюлового латекса, неопудренные
	// 2400003117107	перчатки смотровые (процедурные) из этиленвинилацетата, неопудренные, стерильные
	// 2400003117206	перчатки смотровые (процедурные) из этиленвинилацетата, неопудренные, нестерильные
	// 2400003207907	перчатки смотровые (процедурные) нитриловые, неопудренные, антибактериальные
	// 2400003215308	перчатки смотровые (процедурные) полиизопреновые, неопудренные
	// 2400003297700	перчатки смотровые (процедурные) нитриловые, неопудренные, стерильные
	// 2400003356704	перчатки смотровые (процедурные) виниловые, неопудренные, стерильные
	// 2400003356803	перчатки смотровые (процедурные) виниловые, опудренные, стерильные
	// 2400003433108	перчатки смотровые (процедурные) из латекса гевеи, опудренные, стерильные
	// 2400003492303	перчатки смотровые (процедурные) полиизопреновые, опудренные
	// 2400003495700	перчатки смотровые (процедурные) из полихлоропрена, неопудренные, стерильные
	// 2400003495809	перчатки смотровые (процедурные) из полихлоропрена, неопудренные, стерильные
	// 2400003495908	перчатки смотровые (процедурные) нитриловые, опудренные, стерильные
	// 2400003496004	перчатки смотровые (процедурные) полиизопреновые, неопудренные, стерильные
	// 2400003496103	перчатки смотровые (процедурные) полиизопреновые, опудренные, стерильные
	// 2400001226306	перчатки хирургические из латекса гевеи, неопудренные
	// 2400001226405	перчатки хирургические из латекса гевеи, опудренные
	// 2400001393107	перчатки хирургические из полихлоропрена, неопудренные
	// 2400001393602	перчатки смотровые (процедурные) из полихлоропрена, опудренные
	// 2400001565306	перчатки хирургические из блоксополимера стирола, неопудренные, антибактериальные
	// 2400001857203	перчатки хирургические нитриловые, опудренные
	// 2400001857005	перчатки хирургические нитриловые, неопудренные
	// 2400002015909	перчатки хирургические полиизопреновые, неопудренные
	// 2400002016005	перчатки хирургические полиизопреновые, неопудренные, антибактериальные
	// 2400002016104	перчатки хирургические полиизопреновые, опудренные
	// 2400003161209	перчатки хирургические из блоксополимера стирола, неопудренные
	// 2400003227806	перчатки хирургические полимерно-композитные, неопудренные
	// 2400003237409	перчатки хирургические полимерно-композитные, неопудренные
	// 2400003263408	перчатки хирургические из латекса гевеи, неопудренные, антибактериальные
	// 2400003356902	перчатки хирургические из гваюлового латекса, неопудренные
	// 2400003356902	перчатки хирургические из полихлоропрена, опудренные
	// 2400002886806	набор гигиенической одежды для посетителей
	// 2400002886707	комбинезон гигиенический для посетителей
	// --
	Files File `json:"files,omitempty"` // Массив метаданных Файлов (Максимальное количество файлов - 100)
}

// ProductFolder Группа Товаров
type ProductFolder struct {
	Meta                *Meta          `json:"meta,omitempty"`                // Метаданные Группы Товара (Только для чтения)
	ID                  string         `json:"id,omitempty"`                  // ID Группы товаров (Только для чтения)
	AccountID           string         `json:"accountId,omitempty"`           // ID учетной записи (Только для чтения)
	Owner               *Employee      `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Shared              bool           `json:"shared,omitempty"`              // Общий доступ
	Group               *Group         `json:"group,omitempty"`               // Метаданные отдела сотрудника
	Updated             string         `json:"updated,omitempty"`             // Момент последнего обновления сущности (Только для чтения)
	Name                string         `json:"name,omitempty"`                // Наименование Группы товаров
	Description         string         `json:"description,omitempty"`         // Описание Группы товаров
	Code                string         `json:"Description,omitempty"`         // Код Группы товаров
	ExternalCode        string         `json:"externalCode,omitempty"`        // Внешний код Группы товаров
	Archived            bool           `json:"archived,omitempty"`            // Добавлена ли Группа товаров в архив (Только для чтения)
	PathName            string         `json:"pathName,omitempty"`            // Наименование Группы товаров, в которую входит данная Группа товаров (Только для чтения)
	Vat                 int            `json:"vat,omitempty"`                 // НДС %
	VatEnabled          bool           `json:"vatEnabled,omitempty"`          // Включен ли НДС для группы
	UseParentVat        bool           `json:"useParentVat,omitempty"`        // Используется ли ставка НДС родительской группы
	EffectiveVat        int            `json:"effectiveVat,omitempty"`        // Реальный НДС % (Только для чтения)
	EffectiveVatEnabled int            `json:"effectiveVatEnabled,omitempty"` // Дополнительный признак для определения разграничения реального НДС
	ProductFolder       *ProductFolder `json:"productFolder,omitempty"`       // Ссылка на Группу товаров, в которую входит данная Группа товаров, в формате Метаданных
	TaxSystem           string         `json:"taxSystem,omitempty"`           // Код системы налогообложения
	// --
	// taxSystem
	// TAX_SYSTEM_SAME_AS_GROUP	Совпадает с группой
	// GENERAL_TAX_SYSTEM	ОСН
	// SIMPLIFIED_TAX_SYSTEM_INCOME	УСН. Доход
	// SIMPLIFIED_TAX_SYSTEM_INCOME_OUTCOME	УСН. Доход-Расход
	// UNIFIED_AGRICULTURAL_TAX	ЕСХН
	// PRESUMPTIVE_TAX_SYSTEM	ЕНВД
	// PATENT_BASED	Патент
	// --
}

type Pack struct {
	ID       string    `json:"id"`                 // ID упаковки товара (Только для чтения)
	Uom      *Uom      `json:"uom"`                // Метаданные единиц измерения
	Quantity int       `json:"quantity"`           // Количество Товаров в упаковке данного вида
	Barcodes []Barcode `json:"barcodes,omitempty"` // Массив штрихкодов упаковок товаров
}

type Alcoholic struct {
	Excise   bool    `json:"excise,omitempty"`   // Содержит акцизную марку
	Type     int     `json:"type,omitempty"`     // Код вида продукции
	Strength float64 `json:"strength,omitempty"` // Крепость
	Volume   float64 `json:"volume,omitempty"`   // Объём тары
}
