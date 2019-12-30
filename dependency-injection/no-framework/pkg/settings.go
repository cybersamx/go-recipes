package pkg

type Settings struct {
	HTTPPort int
	DSN      string
}

func NewSettings(port int, dsn string) *Settings {
	return &Settings{
		HTTPPort: port,
		DSN:      dsn,
	}
}
