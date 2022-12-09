package main

import (
	"errors"
	"fmt"
	"os"
)

var _ error = (*SQLError)(nil)

type SQLError struct {
	message string
}

func (de SQLError) Error() string {
	return de.message
}

func newSQLError(message string) *SQLError {
	return &SQLError{message: message}
}

var (
	ErrSQLNoRows = newSQLError("no rows error")
)

// Let's force an error from the nested function call chain that simulate a database call.
// main() -> getData() -> decodeData() -> fetchDataFromMockDB()

func fetchDataFromMockDB() (string, error) {
	return "", ErrSQLNoRows
}

func decodeData() (string, error) {
	data, err := fetchDataFromMockDB()
	if err != nil {
		return "", fmt.Errorf("decode data error: %w", err)
	}

	return data, nil
}

func getData() (string, error) {
	data, err := decodeData()
	if err != nil {
		return "", fmt.Errorf("get data error: %w", err)
	}

	return data, nil
}

func main() {
	data, err := getData()
	if err != nil {
		// err is ErrSQLNoRows but wrapped. So the following won't work.
		// err == ErrSQLNoRows yields false.
		// Use errors.Is() to compare ErrSQLNoRows against the wrapped errors.
		if errors.Is(err, ErrSQLNoRows) {
			fmt.Println("err is ErrSQLNoRows")
		}

		// For errors.As(), the second parameter is a pointer to pointer type.
		var targetErr *SQLError
		if errors.As(err, &targetErr) {
			fmt.Println("err is a SQLError")
		}

		// Output format: topmost error message -> ... -> bottommost error message
		// Output: get data error: decode data error: no rows error
		fmt.Println(err)

		// Unwrap the previous error.
		// Output: decode data error: no rows error
		unwrapErr := errors.Unwrap(err)
		fmt.Println(unwrapErr)

		os.Exit(1)
	}

	fmt.Println(data)
}
