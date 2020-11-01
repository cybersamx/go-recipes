package rdb

import (
	"bytes"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisTimeout = 10 * time.Second
)

// See Go-Redis documentation: https://redis.uptrace.dev/

func RedisSetSession(client *redis.Client, session *Session, seconds int) error {
	buf, err := GOBEncode(&session)
	if err != nil {
		return err
	}

	ctx, closeFn := context.WithTimeout(context.Background(), redisTimeout)
	defer closeFn()

	expires := time.Duration(seconds) * time.Second
	if err := client.Set(ctx, session.SessionID, buf.Bytes(), expires).Err(); err != nil {
		return err
	}

	return nil
}

func RedisGetSession(client *redis.Client, session *Session, key string) error {
	ctx, closeFn := context.WithTimeout(context.Background(), redisTimeout)
	defer closeFn()

	// Get object matching key.
	data, err := client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return NotFoundErr
	} else if err != nil {
		return err
	}

	// Found object, decode it.
	buf := bytes.NewBuffer(data)
	if err := GOBDecode(buf, session); err != nil {
		return err
	}

	return nil
}
