package entity

import (
	"github.com/Vitokz/MoySkladSDK/models/general"
	"time"
)

//Country Страна
type Country struct {
	Meta         *general.Meta `json:"meta"`         // Метаданные
	Id           string        `json:"id"`           // Id Страны
	AccountId    string        `json:"account_id"`   // Id аккаунта
	Updated      time.Time     `json:"updated"`      // Момент последнего обновления сущности
	Shared       bool          `json:"shared"`       // Флаг общий доступ
	Owner        *Employee     `json:"owner"`        // Сотрудник владелец
	Group        *Group        `json:"group"`        // Отдел владелец
	Name         string        `json:"name"`         // Наимнование страны
	Description  string        `json:"description"`  // Описание страны
	Code         string        `json:"code"`         // Код страны
	ExternalCode string        `json:"externalCode"` // Внешний код страны
}

//Region Регион
type Region struct {
	Meta         *general.Meta `json:"meta"`         // Метаданные
	Id           string        `json:"id"`           // Id Региона
	AccountId    string        `json:"account_id"`   // Id аккаунта
	Version      int           `json:"version"`      // Версия сущности
	Updated      time.Time     `json:"updated"`      // Наименование региона
	Code         string        `json:"code"`         // Код региона
	ExternalCode string        `json:"externalCode"` // Внешний код региона
}
