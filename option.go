package log

const (
	ConsoleFormat = "console"
	JsonFormat    = "json"
)

// Option represents a log option.
type Option struct {
	Level            string
	Format           string
	EnableColor      bool
	EnableCaller     bool
	OutputPaths      []string
	ErrorOutputPaths []string
	CallerSkip       int
}

// DefaultOption returns Option with the default value.
func DefaultOption() *Option {
	return &Option{
		Level:            InfoLevel.String(),
		Format:           ConsoleFormat,
		EnableColor:      true,
		EnableCaller:     true,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		CallerSkip:       1,
	}
}
