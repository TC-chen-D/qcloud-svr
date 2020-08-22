package main

import (
	"qcloud-svr/services"
	"qcloud-svr/utils"

	log "github.com/sirupsen/logrus"
)


func main() {
	var conf utils.Conf
	config,err := conf.GetConf()
	if err != nil{
		log.Error(err)
	}
	config.InitLogs()
	requestInfo := &services.RequestInfo{}
	if err := requestInfo.BuildRequestInfo(config); err != nil {
		log.Error(err)
		return
	}
	requestInfo.QcApiSvr()
}
