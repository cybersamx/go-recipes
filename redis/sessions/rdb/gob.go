package rdb

import (
	"bytes"
	"encoding/gob"
)

func GOBEncode(val interface{}) (*bytes.Buffer, error) {
	// bytes.Buffer implements io.Reader and io.Writer.
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(val); err != nil {
		return nil, err
	}

	return &buf, nil
}

func GOBDecode(buf *bytes.Buffer, val interface{}) error {
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(val); err != nil {
		return err
	}

	return nil
}
