package rhaplogger

func (rl *RhapLogger) NewLogError() LogModel {
	return rl.getDefaultLogModel("ERROR")
}
