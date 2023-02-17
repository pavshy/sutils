package sutils

type SimpleLogger interface {
	Printf(format string, a ...any)
	//Println(a ...any)
}

type MultiLog struct {
	loggers map[string]SimpleLogger // map[tag]SimpleLogger
}

func NewMultiLog() *MultiLog {
	return &MultiLog{}
}

func (ml *MultiLog) AddLogger(tag string, logger SimpleLogger) {
	ml.loggers[tag] = logger
}

func (ml *MultiLog) Printf(format string, a ...any) {
	for _, logger := range ml.loggers {
		logger.Printf(format, a...)
	}
}

//func (ml *MultiLog) Println(a ...any) {
//	for _, logger := range ml.loggers {
//		logger.Println(a...)
//	}
//}

func (ml *MultiLog) Tag(tags ...string) *MultiLog {
	newML := &MultiLog{
		loggers: make(map[string]SimpleLogger),
	}
	for _, tag := range tags {
		newML.loggers[tag] = ml.loggers[tag]
	}
	return newML
}
