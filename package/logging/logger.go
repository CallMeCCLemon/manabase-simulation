package logging

import (
	"github.com/onsi/ginkgo/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateLogger Creates a new logger with the default configuration.
func CreateLogger() *zap.Logger {
	// Create a custom logger configuration
	config := zap.NewProductionConfig()

	// Set the output to stdout
	config.OutputPaths = []string{"stdout"}

	// Set the error output to stderr
	config.ErrorOutputPaths = []string{"stderr"}

	// Configure the encoder to use a human-readable format (good for testing)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Optional: Colorize log levels

	// Build the logger
	logger, err := config.Build(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		// Redirect the core output to GinkgoWriter
		return zapcore.NewCore(
			zapcore.NewConsoleEncoder(config.EncoderConfig),
			zapcore.AddSync(ginkgo.GinkgoWriter), // Send logs to GinkgoWriter
			zapcore.InfoLevel,
		)
	}))
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any

	// Use the logger
	return logger
}
