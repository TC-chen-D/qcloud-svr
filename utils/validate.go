package utils

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"github.com/tcchend/qcloud-svr/models"
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

// Self-define validation function
//func checkName(fl validator.FieldLevel) bool {
//	count := utf8.RuneCountInString(fl.Field().String())
//	fmt.Printf("length: %v \n", count)
//	if  count > 5 {
//		return false
//	}
//	return true
//}