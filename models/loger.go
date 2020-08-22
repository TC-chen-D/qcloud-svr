package models

type LogConf struct {
	LogPath       string `yaml:"log_path"`
	LogFile       string `yaml:"log_file"`
	RotationTime  int64  `yaml:"rotation_time"`
	RotationCount uint   `yaml:"rotation_count"`
	LogLevel      string `yaml:"log_level"`
}
