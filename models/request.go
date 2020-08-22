/*
 * Created by Chen on Fri.Aug.2020
 */
package models

// The Instances API requset
type InstancesRequest struct {
	RequestURL       string `yaml:"request_url" json:"-"`
	RequestMethod    string `yaml:"request_method" json:"-"`
	AccessKeyId      string `yaml:"access_key_id"json:"access_key_id"`
	Action           string `yaml:"action"json:"action"`
	Expires          string `yaml:"expires"json:"expires"`
	Limit            string `yaml:"limit"json:"limit"`
	SignatureMethod  string `yaml:"signature_method"json:"signature_method"`
	SignatureVersion string `yaml:"signature_version"json:"signature_version"`
	StatusOne        string `yaml:"status.1"json:"status_one"`
	TimeStamp        string `yaml:"time_stamp"json:"time_stamp"`
	Version          string `yaml:"version"json:"version"`
	Zone             string `yaml:"zone"json:"zone"`
	Signature        string `yaml:"signature"json:"signature"`
}
