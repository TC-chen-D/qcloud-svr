/*
 * Created by Chen on Fri.Aug.2020
 */
package models

// The detail information of vxnet
type VxnetInfo struct {
	VxnetName string `json:"vxnet_name"`
	VxnetType int    `json:"vxnet_type"`
	VxnetId   string `json:"vxnet_id"`
	NicId     string `json:"nic_id"`
	PrivateIp string `json:"private_ip"`
}

// The detail information of image
type ImageInfo struct {
	ProcessorType string `json:"processor_type"`
	Platform      string `json:"platform"`
	ImageSize     int    `json:"image_size"`
	ImageName     string `json:"image_name"`
	ImageId       string `json:"image_id"`
	OsFamily      string `json:"os_family"`
	Provider      string `json:"provider"`
}

// The detail information of instance
type InstanceInfo struct {
	VcpusCurrent     int         `json:"vcpus_current"`
	InstanceId       string      `json:"instance_id"`
	Vxnets           []VxnetInfo `json:"vxnets"`
	MemoryCurrent    int         `json:"memory_current"`
	SubCode          int         `json:"sub_code"`
	TransitionStatus string      `json:"transition_status"`
	InstanceName     string      `json:"instance_name"`
	InstanceType     string      `json:"instance_type"`
	CreateTime       string      `json:"create_time"`
	Status           string      `json:"status"`
	StatusTime       string      `json:"status_time"`
	Image            ImageInfo   `json:"image"`
	Description      string      `json:"description"`
}

// The Instances API response
type InstancesResponse struct {
	Action      string         `json:"action"`
	InstanceSet []InstanceInfo `json:"instance_set"`
	RetCode     int            `json:"ret_code"`
	TotalCount  int            `json:"total_count"`
}

// The general response
type Rtn struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
