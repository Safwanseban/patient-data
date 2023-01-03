package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Patient_name string `json:"patient_name"`
	Age          uint   `json:"age"`
	Disease      string `json:"disease"`
}
