package zoom

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GrantVideoTokenRequest struct {
	RoleType               RoleType
	TPC                    string
	ExpiresIn              time.Duration
	UserIdentity           *string
	SessionKey             *string
	GeoRegions             *string
	CloudRecordingOption   *int
	CloudRecordingElection *int
	TelemetryTrackingID    *string
}

type VideoVersion int

const (
	VideoVersion1 VideoVersion = 1
)

// VideoClaims references: https://developers.zoom.us/docs/video-sdk/auth/#generate-a-video-sdk-jwt
type VideoClaims struct {
	AppKey                 string       `json:"app_key"`
	RoleType               RoleType     `json:"role_type"`
	TPC                    string       `json:"tpc"`
	Version                VideoVersion `json:"version"`
	UserIdentity           *string      `json:"user_identity,omitempty"`
	SessionKey             *string      `json:"session_key,omitempty"`
	GeoRegions             *string      `json:"geo_regions,omitempty"`
	CloudRecordingOption   *int         `json:"cloud_recording_option,omitempty"`
	CloudRecordingElection *int         `json:"cloud_recording_election,omitempty"`
	TelemetryTrackingID    *string      `json:"telemetry_tracking_id,omitempty"`

	jwt.RegisteredClaims
}

func (c *Client) GrantVideoToken(ctx context.Context, req GrantVideoTokenRequest) (Token, error) {
	iat := time.Now()
	exp := iat.Add(req.ExpiresIn)
	claims := VideoClaims{
		AppKey:                 c.key,
		RoleType:               req.RoleType,
		TPC:                    req.TPC,
		Version:                VideoVersion1,
		UserIdentity:           req.UserIdentity,
		SessionKey:             req.SessionKey,
		GeoRegions:             req.GeoRegions,
		CloudRecordingOption:   req.CloudRecordingOption,
		CloudRecordingElection: req.CloudRecordingElection,
		TelemetryTrackingID:    req.TelemetryTrackingID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(iat),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(c.secret))
	if err != nil {
		return Token{}, err
	}
	return Token{
		AccessToken: ss,
		ExpiresIn:   req.ExpiresIn,
	}, nil
}
