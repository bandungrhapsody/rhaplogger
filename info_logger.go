package rhaplogger

func (rl *RhapLogger) NewLogInfo() LogModel {
	return rl.getDefaultLogModel("INFO")
}
