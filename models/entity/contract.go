package entity

import (
	"github.com/Vitokz/MoySkladSDK/models/general"
	"time"
)

type ContractType string
type RewardType string

const (
	Commission ContractType = "Commission" // Договор коммиссии
	Sales      ContractType = "Sales"      // Договор купли-продажи

	PercentOfSales RewardType = "PercentOfSales" // Процент от суммы продажи
	None           RewardType = "None"           // Не рассчитывать
)

// Contract Договор
type Contract struct {
	Meta                *general.Meta       `json:"meta,omitempty"`
	Id                  string              `json:"id,omitempty"`                  // ID Договора	(Только для чтения)
	AccountId           string              `json:"accountId,omitempty"`           // ID учетной записи	(Только для чтения)
	Owner               *Employee           `json:"owner,omitempty"`               // Метаданные владельца (Сотрудника)
	Shared              bool                `json:"shared,omitempty"`              // Общий доступ
	Group               *Group              `json:"group,omitempty"`               // Метаданные отдела сотрудника
	Updated             time.Time           `json:"updated,omitempty"`             // Момент последнего обновления сущности	(Только для чтения)
	Name                string              `json:"name,omitempty"`                // Номер договора	(Необходимое при создании)
	Description         string              `json:"description,omitempty"`         // Описание Договора
	Code                string              `json:"code,omitempty"`                // Код Договора
	ExternalCode        string              `json:"externalCode,omitempty"`        // Внешний код Договора
	Archived            bool                `json:"archived,omitempty"`            // Добавлен ли Договор в архив
	Moment              time.Time           `json:"moment,omitempty"`              // Дата Договора
	Sum                 int                 `json:"sum,omitempty"`                 // Сумма Договора
	ContractType        ContractType        `json:"contractType,omitempty"`        // Тип договора. Возможные значения договра
	RewardType          RewardType          `json:"rewardType,omitempty"`          // Тип Вознаграждения
	RewardPercent       int                 `json:"rewardPercent,omitempty"`       // Вознаграждение в процентах (от 0 до 100)
	OwnAgent            *Organization       `json:"ownAgent,omitempty"`            // Метаданные вашего юрлица - организации (Необходимое при создании)
	Agent               *Counterparty       `json:"agent,omitempty"`               // Метаданные Контрагента	(Необходимое при создании)
	State               *general.State      `json:"state,omitempty"`               // Метаданные статуса договора
	OrganizationAccount *Account            `json:"organizationAccount,omitempty"` // Метаданные счета вашего юрлица
	AgentAccount        *Account            `json:"agentAccount,omitempty"`        // Метаданные счета контрагента
	Rate                *Currency           `json:"rate,omitempty"`                // Метаданные валюты
	Attributes          []general.Attribute `json:"attributes,omitempty"`          // Коллекция доп. полей
}
