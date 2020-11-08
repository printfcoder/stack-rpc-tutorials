package main

import (
	"os"

	"github.com/micro/go-micro/v2/logger"
	lr "github.com/micro/go-plugins/logger/logrus/v2"
)

func main() {
	l := lr.NewLogger(
		logger.WithOutput(os.Stdout)).Fields(map[string]interface{}{
		"header1": "头1",
		"header2": 8080,
	})

	logger.DefaultLogger = l

	logger.Info("testing: Info")
	logger.Infof("testing: %s", "Infof")
}
