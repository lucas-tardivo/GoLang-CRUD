package router

import (
	"simple-crud/controller"

	"github.com/gin-gonic/gin"
)

type PersonRouteController struct {
	personController controller.PersonController
}

func NewRoutePersonController(personController controller.PersonController) PersonRouteController {
	return PersonRouteController{personController}
}

func (pc *PersonRouteController) StartRouter(rg *gin.RouterGroup) {
	router := rg.Group("persons")
	router.POST("/", pc.personController.PostPerson)
	router.GET("/", pc.personController.FindPersons)
	router.PUT("/:personId", pc.personController.UpdatePerson)
	router.GET("/:personId", pc.personController.FindPersonById)
	router.DELETE("/:personId", pc.personController.DeletePerson)
}