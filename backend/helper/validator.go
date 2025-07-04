package helper

import (
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ValidatorErrorMsg(err error) map[string]string {
	errsMap := make(map[string]string)

	if validatorErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validatorErrors {
			field := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				errsMap[field] = fmt.Sprintf("%s harus diisi", field)
			case "email":
				errsMap[field] = "format email tidak valid"
			case "unique":
				errsMap[field] = fmt.Sprintf("%s sudah digunakan", field)
			case "min":
				errsMap[field] = fmt.Sprintf("%s minimal harus sebanyak %s karakter", field, fieldError.Param())
			case "max":
				errsMap[field] = fmt.Sprintf("%s maksimal tidak lebih dari %s karakter", field, fieldError.Param())
			case "numeric":
				errsMap[field] = fmt.Sprintf("%s harus bernilai angka", field)
			default:
				errsMap[field] = "nilai tidak valid"
			}
		}
	}

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errsMap["username"] = "username sudah tersedia"
			}
			if strings.Contains(err.Error(), "email") {
				errsMap["email"] = "email sudah tersedia"
			}
		}
	} else if err == gorm.ErrRecordNotFound {
		errsMap["error"] = "data tidak ditemukan"
	}
	return errsMap
}

func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}
