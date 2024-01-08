package zoom

import "time"

type SessionStartedPayload struct {
	AccountID string                      `json:"account_id"`
	Object    SessionStartedPayloadObject `json:"object"`
}

type SessionStartedPayloadObject struct {
	ID          string    `json:"id"`
	SessionID   string    `json:"session_id"`
	SessionName string    `json:"session_name"`
	SessionKey  string    `json:"session_key"`
	StartTime   time.Time `json:"start_time"`
}

type SessionEndedPayload struct {
	AccountID string                    `json:"account_id"`
	Object    SessionEndedPayloadObject `json:"object"`
}

type SessionEndedPayloadObject struct {
	ID          string    `json:"id"`
	SessionID   string    `json:"session_id"`
	SessionName string    `json:"session_name"`
	SessionKey  string    `json:"session_key"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
