package zoom

import (
	"context"
	"time"
)

type Token struct {
	AccessToken string
	ExpiresIn   time.Duration
}

type RoleType int

const (
	RoleTypeParticipant RoleType = 0
	RoleTypeHost        RoleType = 1
)

type Authentication interface {
	GrantVideoToken(ctx context.Context, req GrantVideoTokenRequest) (Token, error)
}
