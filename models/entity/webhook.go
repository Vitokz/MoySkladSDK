package entity

const (
	Delete    WebHookActionType = "DELETE"
	Create    WebHookActionType = "CREATE"
	Update    WebHookActionType = "UPDATE"
	Processed WebHookActionType = "PROCESSED"

	CustomerOrder WebhookType = "CustomerOrder" // Заказ покупателя
	InvoiceOut    WebhookType = "InvoiceOut"    // Счёт покупателя
	InvoiceIn     WebhookType = "InvoiceIn"     // Счёт поставщика
	PurchaseOrder WebhookType = "PurchaseOrder" // Заказ поставщику
	Demand        WebhookType = "Demand"        // Отгрузка
	Supply        WebhookType = "Supply"        // Приемка
)

type (
	WebhookType       string
	WebHookActionType string

	WebhookRequest struct {
		Id        string      `json:"id" form:"id" query:"id"`
		Type      WebhookType `json:"type" form:"type" query:"type"`
		RequestId string      `json:"requestId" form:"requestId" query:"requestId"`
	}

	Webhook struct {
		Meta       *Meta             `json:"meta,omitempty"`
		ID         string            `json:"id,omitempty"`
		AccountID  string            `json:"accountId,omitempty"`
		EntityType string            `json:"entityType,omitempty"`
		URL        string            `json:"url,omitempty"`
		Method     string            `json:"method,omitempty"`
		Enable     bool              `json:"enable,omitempty"`
		Action     WebHookActionType `json:"action,omitempty"`
	}
)
