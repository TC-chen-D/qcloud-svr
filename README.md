# QCloudSvr 

QingCloud API接口调用SDK封装

### 说明
目前通过配置文件获取请求参数，实际应用可以通过其他方式

目前只针对Get请求API进行了封装，可拓展其他如POST、PUT、DELETE等其他方法

输入参数验证根据实际各个字段要求增加相应相应验证方法，注册到validator验证器里面

### 使用实例
```cassandraql
	var conf conf.Conf
	config, err := conf.GetConf()
	if err != nil {
		log.Error(err)
	}
	config.InitLogs()
	requestInfo := &services.RequestInfo{}
	if err := requestInfo.BuildRequestInfo(config); err != nil {
		log.Error(err)
		return
	}
	requestInfo.QcApiSvr()
```