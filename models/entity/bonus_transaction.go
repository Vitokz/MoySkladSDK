package entity

import (
	"github.com/Vitokz/MoySkladSDK/models/general"
	"time"
)

// BonusProgram Бонусная программа
type BonusProgram struct {
	Meta                      *general.Meta `json:"meta"`                              // Метаданные Бонусной программы
	ID                        string        `json:"id"`                                // ID Бонусной программы (Только для чтения)
	AccountID                 string        `json:"accountId"`                         // ID учетной записи (Только для чтения)
	Name                      string        `json:"name,omitempty"`                    // Наименование Бонусной программы
	Active                    bool          `json:"active"`                            // Индикатор, является ли бонусная программа активной на данный момент
	AllProducts               bool          `json:"allProducts"`                       // Индикатор, действует ли бонусная программа на все товары (всегда true)
	AllAgents                 bool          `json:"allAgents"`                         // Индикатор, действует ли скидка на всех контрагентов
	AgentTags                 []string      `json:"agentTags"`                         // Тэги контрагентов, к которым применяется бонусная программа. В случае пустого значения контрагентов в результате выводится пустой массив
	EarnRateRoublesToPoint    int           `json:"earnRateRoublesToPoint,omitempty"`  // Курс начисления
	SpendRatePointsToRouble   int           `json:"spendRatePointsToRouble,omitempty"` // Курс списания
	MaxPaidRatePercents       int           `json:"maxPaidRatePercents,omitempty"`     // Максимальный процент оплаты баллами
	PostponedBonusesDelayDays int           `json:"postponedBonusesDelayDays"`         // Баллы начисляются через [N] дней
	EarnWhileRedeeming        bool          `json:"earnWhileRedeeming"`                // Разрешить одновременное начисление и списание бонусов. Если true, то бонусы начислены
	WelcomeBonusesEnabled     bool          `json:"welcomeBonusesEnabled"`             // Возможность начисления приветственных баллов
	WelcomeBonusesValue       int           `json:"welcomeBonusesValue"`               // Колчество приветственных бонусов начисляемых участникам бонусной программы
	WelcomeBonusesEnabledEnum string        `json:"welcomeBonusesEnabled"` //Не трогать
}

type BonusTransaction struct {
	Meta              *general.Meta `json:"meta"`              // Метаданные
	Id                string        `json:"id"`                // ID Бонусной операции (Только для чтения)
	AccountId         string        `json:"accountId"`         // ID учетной записи (Только для чтения)
	Owner             *Employee     `json:"owner"`             // Метаданные владельца (Сотрудника)
	Shared            bool          `json:"shared"`            // Общий доступ
	Group             *Group        `json:"group"`             // Метаданные отдела сотрудника
	Updated           time.Time     `json:"updated"`           // Момент последнего обновления сущности (Только для чтения)
	Created           time.Time     `json:"created"`           // Время создания бонусной операции
	Code              string        `json:"code"`              // Код бонусной операции
	ExternalCode      string        `json:"externalCode"`      // Внешний код бонусной операции
	Name              string        `json:"name"`              // Наименование Бонусной операции
	Applicable        bool          `json:"applicable"`        // Отметка о поведении
	Moment            time.Time     `json:"moment"`            // Время проведения бонусной операции
	Agent             *Counterparty `json:"agent,omitempty"`   // Метаданные Контрагента	(Необходимое при создании)
	ParentDocument    general.Meta  `json:"parentDocument"`    // Метаданные связанного документа
	BonusProgram      *BonusProgram `json:"bonusProgram"`      // Метаданные бонусной программы
	BonusValue        int           `json:"bonusValue"`        // Количество бонусных баллов
	Organization      *Organization `json:"organization"`      // Метаданные юрлица
	TransactionType   string        `json:"transactionType"`   // Тип бонусной транзации
	TransactionStatus string        `json:"transactionStatus"` // Статус бонусной операции
	ExecutionDate     time.Time     `json:"executionDate"`     //Дата начисления бонусной операции
	CategoryType      string        `json:"categoryType"`      // Категория бонусной операции
}
