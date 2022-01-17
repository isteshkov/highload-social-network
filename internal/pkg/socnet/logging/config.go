package logging

const (
	lvlDebug = iota
	lvlInfo
	lvlWarn
	lvlError
	lvlFatal
)

type Config struct {
	LogLvl        string
	Version       string
	Release       string
	CommitSha     string
	AppName       string

	logLvl int
}

func getLvlFromString(tag string) int {
	switch tag {
	case LevelDebug:
		return lvlDebug
	case LevelInfo:
		return lvlInfo
	case LevelWarning:
		return lvlWarn
	case LevelError:
		return lvlError
	case LevelFatal:
		return lvlFatal
	default:
		return lvlDebug
	}
}
