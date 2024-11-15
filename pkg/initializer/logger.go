package initializer

import "github.com/sirupsen/logrus"

func InitializeLogger() error {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	return nil
}
