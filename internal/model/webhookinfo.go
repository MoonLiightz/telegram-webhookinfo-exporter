package model

// WebhookInfo ...
type WebhookInfo struct {
	Ok     bool `json:"ok"`
	Result struct {
		URL                  string `json:"url"`
		HasCustomCertificate bool   `json:"has_custom_certificate"`
		PendingUpdateCount   int64  `json:"pending_update_count"`
		LastErrorDate        int64  `json:"last_error_date"`
		LastErrorMessage     string `json:"last_error_message"`
		MaxConnections       int8   `json:"max_connections"`
	} `json:"result"`
}
