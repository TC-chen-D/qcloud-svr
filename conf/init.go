package conf

import (
	log "github.com/sirupsen/logrus"
	)

func InitInstances() *Conf{
	var conf Conf
	config, err := conf.GetConf()
	if err != nil {
		log.Error(err)
	}
	config.InitLogs()
	return config
}

