package main

import (
	"github.com/Safwanseban/interview-patient/routes"
	"github.com/gin-gonic/gin"

	"github.com/Safwanseban/interview-patient/configs"
)

var R = gin.New()

func init() {

	configs.Connecttodb()
}

func main() {

	routes.Routes(R)
	R.Run(":7876")

}
