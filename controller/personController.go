package controller

import (
	"simple-crud/model/entity"
	"simple-crud/model/request"
	"strings"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonController struct {
	DB *gorm.DB
}

func NewPersonController(DB *gorm.DB) PersonController {
	return PersonController{DB}
}

func (pc *PersonController) PostPerson(c *gin.Context) {
	var payload *request.CreatePersonRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newPerson := entity.Person{
		Name:     payload.Name,
		DateBirth:   payload.DateBirth,
		Contact:     payload.Contact,
	}

	result := pc.DB.Create(&newPerson)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPerson})
}

func (pc *PersonController) UpdatePerson(ctx *gin.Context) {
	personId := ctx.Param("personId")

	var payload *request.UpdatePersonRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var updatedPerson entity.Person
	result := pc.DB.First(&updatedPerson, "id = ?", personId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No person with that ID exists"})
		return
	}

	personToUpdate := entity.Person{
		Name:     payload.Name,
		DateBirth:   payload.DateBirth,
		Contact:     payload.Contact,
	}

	pc.DB.Model(&updatedPerson).Updates(personToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPerson})
}

func (pc *PersonController) FindPersonById(ctx *gin.Context) {
	personId := ctx.Param("personId")

	var person entity.Person
	result := pc.DB.First(&person, "id = ?", personId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No person with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": person})
}

func (pc *PersonController) FindPersons(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var persons []entity.Person
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&persons)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(persons), "data": persons})
}

func (pc *PersonController) DeletePerson(ctx *gin.Context) {
	personId := ctx.Param("personId")

	result := pc.DB.Delete(&entity.Person{}, "id = ?", personId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No person with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}