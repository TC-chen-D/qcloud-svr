/*
 * Created by Chen on Fri.Aug.2020
 */

package services

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"qcloud-svr/models"
	"qcloud-svr/utils"
	"strings"

	log "github.com/sirupsen/logrus"
)

// RequestInfo is the request information for QingCloud API services.
type RequestInfo struct {
	RequestURL        string
	RequestMethod     string
	RequestParameters map[string]string
	RequestHeader     map[string]string
}

func (r *RequestInfo) NewRequestInfo() error {
	switch r.RequestMethod {
	case "GET":
		if err := r.ParseGet(); err != nil {
			return err
		}
	case "POST":
		if err := r.ParsePost(); err != nil {
			return err
		}
	case "HEAD":
		if err := r.ParseHead(); err != nil {
			return err
		}
	case "PUT":
		if err := r.ParsePut(); err != nil {
			return err
		}
	case "DELETE":
		if err := r.ParseDelete(); err != nil {
			return err
		}
	default:
		if err := r.ParseAny(); err != nil {
			return err
		}
	}
	return nil
}

func (r *RequestInfo) ParseGet() error {
	paramsParts := []string{}
	parameters, err := json.Marshal(r.RequestParameters)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(parameters, &r.RequestParameters); err != nil {
		return err
	}
	for key, value := range r.RequestParameters {
		paramsParts = append(paramsParts, fmt.Sprintf("%s=%s", key, value))
	}
	joined := strings.Join(paramsParts, "&")
	if joined != "" {
		r.RequestURL += "?" + joined
	}
	return nil
}

func (r *RequestInfo) ParsePost() error {
	// parse post
	return nil
}

func (r *RequestInfo) ParseHead() error {
	return nil
}

func (r *RequestInfo) ParsePut() error {
	return nil
}

func (r *RequestInfo) ParseDelete() error {
	return nil
}

func (r *RequestInfo) ParseAny() error {
	return nil
}

func NewInstancesRequest(c *utils.Conf) *models.InstancesRequest {
	return &c.InstancesRequest
}

func (r *RequestInfo) BuildRequestInfo(c *utils.Conf) error {
	r.RequestURL = c.InstancesRequest.RequestURL
	r.RequestMethod = c.InstancesRequest.RequestMethod
	instancesRequest := NewInstancesRequest(c)
	j, err := json.Marshal(instancesRequest)
	if err != nil {
		log.Error(err)
		return err
	}
	if err = json.Unmarshal(j, &r.RequestParameters); err != nil {
		log.Error(err)
		return err
	}
	if err = r.NewRequestInfo(); err != nil {
		return err
	}
	return nil
}

func (r *RequestInfo) QcApiSvr() *models.Rtn {
	req, err := http.NewRequest(r.RequestMethod, r.RequestURL, nil)
	if err != nil {
		log.Error(err)
		return &models.Rtn{http.StatusBadRequest,nil,"The bad request!"}
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true, Renegotiation: tls.RenegotiateOnceAsClient},
	}

	cli := http.Client{
		Transport:     tr,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	response, err := cli.Do(req)
	if err != nil {
		log.Error(err)
		return &models.Rtn{}
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return &models.Rtn{}
	}
	var instanceResponse models.InstancesResponse
	if response.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &instanceResponse)
		log.Info(http.StatusOK)
		return &models.Rtn{http.StatusOK, instanceResponse, "sucessful response!"}
	} else {
		log.Error(err)
		return &models.Rtn{}
	}
}
