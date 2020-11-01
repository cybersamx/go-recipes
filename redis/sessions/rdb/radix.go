package rdb

import (
	"bytes"

	"github.com/mediocregopher/radix/v3"
	"github.com/mediocregopher/radix/v3/resp"
)

// See Radix documentation on NewLenReader and FlatCmd:
// https://godoc.org/github.com/mediocregopher/radix#FlatCmd

func RadixSetSession(client *radix.Pool, session *Session, seconds int) error {
	buf, err := GOBEncode(&session)
	if err != nil {
		return err
	}

	// Set the object.
	reader := resp.NewLenReader(buf, int64(buf.Len()))
	if err = client.Do(radix.FlatCmd(nil, "SET", session.SessionID, reader)); err != nil {
		return err
	}

	// Set expiration for the object.
	if err = client.Do(radix.FlatCmd(nil, "EXPIRE", session.SessionID, seconds)); err != nil {
		return err
	}

	return nil
}

func RadixGetSession(client *radix.Pool, session *Session, key string) error {
	var buf bytes.Buffer
	mn := radix.MaybeNil{
		Rcv: &buf,
	}

	// Get object matching key.
	if err := client.Do(radix.FlatCmd(&mn, "GET", key)); err != nil {
		return err
	}

	// Check for nil (not found).
	if mn.Nil {
		return NotFoundErr
	}

	// Found object, decode it.
	if err := GOBDecode(&buf, session); err != nil {
		return err
	}

	return nil
}
