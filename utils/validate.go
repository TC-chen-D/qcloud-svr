package utils

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"qcloud-svr/models"
)

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
