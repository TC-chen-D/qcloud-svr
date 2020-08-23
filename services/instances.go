/*
 * Created by Chen on Fri.Aug.2020
 */

package services

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
	"github.com/tcchend/qcloud-svr/conf"
	"github.com/tcchend/qcloud-svr/models"
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

// Verification of input information
func Validate(input *models.InstancesRequest) (*models.InstancesRequest, error) {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println("err:", err)
		}
		return nil,err
	}
	return input, nil
}

// Get parameter validation from configuration information
func NewInstancesRequest(c *conf.Conf) (*models.InstancesRequest,error) {
	res,err := Validate(&c.InstancesRequest)
	if err != nil {
		return nil,err
	}
	return res,nil
}


// Build the request information
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

// Parse GET method
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

// parse post
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


// Encapsulation the http request information
func (r *RequestInfo) BuildRequestInfo(c *conf.Conf) error {
	r.RequestURL = c.ApiUrl.RequestURL
	r.RequestMethod = c.ApiUrl.RequestMethod
	instancesRequest,err := NewInstancesRequest(c)
	if err != nil {
		log.Error(err)
		return err
	}
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
		fmt.Println(instanceResponse)
		return &models.Rtn{http.StatusOK, instanceResponse, "sucessful response!"}
	} else {
		log.Error(err)
		return &models.Rtn{}
	}
}
