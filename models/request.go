/*
 * Created by Chen on Fri.Aug.2020
 */
package models

// The Instances API request.
type InstancesRequest struct {
	AccessKeyId      string `yaml:"access_key_id"json:"access_key_id"validate:"required"`
	Action           string `yaml:"action"json:"action"validate:"required"`
	Expires          string `yaml:"expires"json:"expires"validate:"required"`
	Limit            string `yaml:"limit"json:"limit"validate:"required"`
	SignatureMethod  string `yaml:"signature_method"json:"signature_method"validate:"required"`
	SignatureVersion string `yaml:"signature_version"json:"signature_version"validate:"required"`
	StatusOne        string `yaml:"status.1"json:"status_one"validate:"required"`
	TimeStamp        string `yaml:"time_stamp"json:"time_stamp"validate:"required"`
	Version          string `yaml:"version"json:"version"validate:"required"`
	Zone             string `yaml:"zone"json:"zone"validate:"required"`
	Signature        string `yaml:"signature"json:"signature"validate:"required"`
}
