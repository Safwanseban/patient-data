package routes

import ("github.com/gin-gonic/gin"

"github.com/Safwanseban/interview-patient/controllers"
)

func Routes(c *gin.Engine) {


c.GET("/",controllers.GEtPatientnData)
c.POST("/insert",controllers.InsertPatient)
c.PUT("/update",controllers.UpdatePatient)
c.DELETE("/delete",controllers.DeleteFromPatient)

}
