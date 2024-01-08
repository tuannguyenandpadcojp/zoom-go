package zoom

import "encoding/json"

type WebhookCallBackRequest struct {
	Payload json.RawMessage `json:"payload"`
	Event   string          `json:"event"`
	EventTs int64           `json:"event_ts"`
}
