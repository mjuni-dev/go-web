package module

type Config struct {
	ConnectionString string
	// Logger
	// Email/Notification Service
}

type Module interface {
	Name() string
	Initialize() error
}
