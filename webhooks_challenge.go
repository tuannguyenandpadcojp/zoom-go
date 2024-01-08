package zoom

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type URLCallbackChallengePayload struct {
	PlainToken string `json:"plainToken"`
}

type URLCallbackChallengeResponse struct {
	PlainToken     string `json:"plainToken"`
	EncryptedToken string `json:"encryptedToken"`
}

func CallBackRequest(r *http.Request) (WebhookCallBackRequest, error) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return WebhookCallBackRequest{}, fmt.Errorf("read request body: %w", err)
	}
	defer r.Body.Close()
	var callbackReq WebhookCallBackRequest
	if err := json.Unmarshal(b, &callbackReq); err != nil {
		return WebhookCallBackRequest{}, fmt.Errorf("unmarshal zoom callback request: %w", err)
	}
	return callbackReq, nil
}

func Challenger(payload json.RawMessage, webhookSecretToken string) (URLCallbackChallengeResponse, error) {
	var challengePayload URLCallbackChallengePayload
	if err := json.Unmarshal(payload, &challengePayload); err != nil {
		return URLCallbackChallengeResponse{}, fmt.Errorf("unmarshal url callback challenge payload: %w", err)
	}
	plainTokenEncrypted, err := HMAC256(challengePayload.PlainToken, webhookSecretToken)
	if err != nil {
		return URLCallbackChallengeResponse{}, fmt.Errorf("hashing plainToken: %w", err)
	}

	return URLCallbackChallengeResponse{
		PlainToken:     challengePayload.PlainToken,
		EncryptedToken: plainTokenEncrypted,
	}, nil
}
