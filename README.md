# QCloudSvr 

QingCloud API接口调用SDK封装

### 说明
目前通过配置文件获取请求参数,请将项目下配置文件拷贝至需要引用此包的项目路径中，并根据需要修改相关配置文件中的赋值
`实际应用可以通过其他方式传入参数`

目前只针对Get请求API进行了封装，可拓展其他如POST、PUT、DELETE等其他方法

输入参数验证根据实际各个字段要求增加相应相应验证方法，注册到validator验证器里面

### 使用实例
```
	c := conf.InitInstances()
    
    	requestInfo := &services.RequestInfo{}
    	if err := requestInfo.BuildRequestInfo(c); err != nil {
    		log.Error(err)
    		return
    	}
    	result := requestInfo.ProcessRequest()
    
    	// Print the response.
    	fmt.Println(result)
    
    	// Print the return code.
    	fmt.Println(result.RetCode)
```