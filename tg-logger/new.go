package tg_logger

type ServiceLogger struct {
	API string
}

func NewServiceLogger(api string) *ServiceLogger {
	return &ServiceLogger{API: api}
}
