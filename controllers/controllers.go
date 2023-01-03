package controllers

import (
	"fmt"
	"strconv"

	"github.com/Safwanseban/interview-patient/configs"
	"github.com/Safwanseban/interview-patient/models"
	"github.com/gin-gonic/gin"
)

func InsertPatient(c *gin.Context) {

	var patient models.Patient
	fmt.Println(patient.Patient_name, patient.Age)
	if err := c.ShouldBindJSON(&patient); err != nil {

		c.JSON(404, gin.H{

			"err": err,
		})
		c.Abort()
		return
	}

	record := configs.Db.Create(&patient)
	if record.Error != nil {
		c.JSON(500, gin.H{

			"err": record.Error.Error(),
		})
		c.Abort()
		return

	}
	c.JSON(200, gin.H{
		"status": true,
		"msg":    "patient data inserted",
	})

}
func GEtPatientnData(c *gin.Context) {

	var patient []models.Patient
	record := configs.Db.Find(&patient)
	if record.Error != nil {

		c.JSON(500, gin.H{

			"err": record.Error.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200,gin.H{

		"status":true,
		"msg":patient,
	})

}
var pupdate struct{
	Patient_name string `json:"patient_name"`
	Age          uint   `json:"age"`
	Disease      string `json:"disease"`

}
func UpdatePatient(c *gin.Context) {
	var patient models.Patient
	id := c.Query("id")
	ID, _ := strconv.Atoi(id)

	if err := c.ShouldBindJSON(&pupdate); err != nil {

		c.JSON(500, gin.H{"err": err.Error()})
		c.Abort()
		return
	}

	record := configs.Db.Raw("update patients set patient_name=?,age=?,disease=? where id=?", pupdate.Patient_name, pupdate.Age, pupdate.Disease, ID).Scan(&patient)

	if record.Error != nil {
		c.JSON(500, gin.H{
			"err": record.Error.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200,gin.H{

		"status":true,
		"msg":"patient details updated",
	})


}
func DeleteFromPatient(c *gin.Context) {
	var patient models.Patient
	id := c.Query("id")
	ID, _ := strconv.Atoi(id)


	record := configs.Db.Raw("delete from patients where id=?", ID).Scan(&patient)
	if record.Error != nil {
		c.JSON(404, gin.H{

			"err": record.Error.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200,gin.H{

		"status":true,
		"msg":"patient deleted",
	})

}
