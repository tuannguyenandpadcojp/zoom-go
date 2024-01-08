package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/smithy-go/ptr"
	"github.com/google/uuid"

	zoom "github.com/tuannguyenandpadcojp/zoom-go"
)

func main() {
	var zoomAPI zoom.API
	zoomAPI, err := zoom.NewClient(
		zoom.WithHTTPClient(
			http.DefaultClient,
		),
	)
	if err != nil {
		log.Fatalf("initializing zoom SDK: %v", err)
	}
	http.HandleFunc("/zoom/callbacks", func(w http.ResponseWriter, r *http.Request) {
		callbackReq, err := zoom.CallBackRequest(r)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		switch callbackReq.Event {
		case zoom.CallBackURLValidationEvent:
			URLCallbackChallenger(w, callbackReq)
			return
		case zoom.SessionStartedEvent:
			SessionStarted(w, callbackReq)
			return
		case zoom.SessionEndedEvent:
			SessionEnded(w, callbackReq)
			return
		default:
			fmt.Fprintf(w, "unsupported the event:%s", callbackReq.Event)
		}
	})
	http.HandleFunc("/sessions", func(w http.ResponseWriter, r *http.Request) {
		sessionID := uuid.NewString()
		token, err := zoomAPI.GrantVideoToken(context.Background(), zoom.GrantVideoTokenRequest{
			RoleType:     zoom.RoleTypeHost, // grant token for host to start session
			TPC:          sessionID,         // session name
			ExpiresIn:    60 * time.Second,
			UserIdentity: ptr.String("issuer"),
			SessionKey:   ptr.String(sessionID),
			GeoRegions:   ptr.String("JP"),
		})
		if err != nil {
			fmt.Fprintf(w, "grant video token: %v", err)
			return
		}

		fmt.Printf("access_token: %s\n", token.AccessToken)
		w.Write([]byte(fmt.Sprintf(`{"access_token": "%s","session_id": "%s"}`, token.AccessToken, sessionID)))
	})
	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		log.Fatalf("start server: %v\n", err)
	}
}

func URLCallbackChallenger(w http.ResponseWriter, r zoom.WebhookCallBackRequest) {
	resp, err := zoom.Challenger(r.Payload, os.Getenv("ZOOM_WEBHOOK_SECRET_TOKEN"))
	if err != nil {
		fmt.Fprintf(w, "url callback challenge: %v", err)
		return
	}
	b, err := json.Marshal(&resp)
	if err != nil {
		fmt.Fprintf(w, "marshal url callback challenge response: %v", err)
		return
	}
	w.Write(b)
}

func SessionStarted(w http.ResponseWriter, r zoom.WebhookCallBackRequest) {
	var pl zoom.SessionStartedPayload
	if err := json.Unmarshal(r.Payload, &pl); err != nil {
		fmt.Fprintf(w, "marshal url session started payload event: %v", err)
		return
	}
	fmt.Printf(">>>> session started: %+v\n", pl)
}

func SessionEnded(w http.ResponseWriter, r zoom.WebhookCallBackRequest) {
	var pl zoom.SessionEndedPayload
	if err := json.Unmarshal(r.Payload, &pl); err != nil {
		fmt.Fprintf(w, "marshal url session started payload event: %v", err)
		return
	}
	fmt.Printf(">>>> session ended: %+v\n", pl)
}
