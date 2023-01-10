package main

type config struct {
	// These fields will bind to the global flags.
	addr      string
	keyID     string
	secretKey string
	useSSL    bool

	// These fields will bind to the flags in the list command.
	recursive  bool
	longFormat bool
}
