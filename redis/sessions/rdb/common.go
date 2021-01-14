package rdb

import (
	"errors"
)

// Session represents a session of an application.
type Session struct {
	SessionID string
	UserID    uint
	Username  string
}

var (
	NotFoundErr = errors.New("not found")
)
