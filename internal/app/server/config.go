package application

type Config struct {
	ProfilingApiPort   string
	PublicApiPort      string
	MetricApiPort      string
	TimeOutSecond      int
	MaxMultipartMemory int64
}
