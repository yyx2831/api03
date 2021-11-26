package main

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	//config.OutputPaths = []string{"stdout"}
	//config.ErrorOutputPaths = []string{"stderr"}
	config.OutputPaths = []string{"zap_test/zap_log_file/log.log", "stderr"}
	config.ErrorOutputPaths = []string{"zap_test/zap_log_file/log.log", "stderr"}
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("info")
	logger.Error("error")
}
