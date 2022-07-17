package main

type requestLogger interface {
	print(string)
}

type logRequestDecorator struct {
	baseService appService
	log         requestLogger
}

func NewLogRequestDecorator(baseService appService, log requestLogger) *logRequestDecorator {
	return &logRequestDecorator{
		baseService: baseService,
		log:         log,
	}
}

func (d *logRequestDecorator) processRequest(request string) string {
	d.log.print(request)
	return d.baseService.processRequest(request)
}
