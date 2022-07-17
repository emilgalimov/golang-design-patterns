package main

type responseLogger interface {
	print(string)
}

type logResponseDecorator struct {
	baseService appService
	log         responseLogger
}

func NewLogResponseDecorator(baseService appService, log requestLogger) *logResponseDecorator {
	return &logResponseDecorator{
		baseService: baseService,
		log:         log,
	}
}

func (d *logResponseDecorator) processRequest(request string) string {
	response := d.baseService.processRequest(request)
	d.log.print(response)
	return response
}
