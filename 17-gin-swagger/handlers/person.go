package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webapp/models"
)

// GetPerson godoc
// @Summary      get a person
// @Description  get string by Name
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        name   path      string  true  "Person Name"
// @Success      200  {object}  models.Person
// @Router       /persons/{name} [get]
func GetPerson(c *gin.Context) {
	if personName := c.Param("name"); personName != "" {
		// check if the person exists
		if person, err := models.GetPersonByName(personName); err == nil && person != nil {
			c.JSON(http.StatusOK, person)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// AddPerson godoc
// @Summary add person
// @Description add person by information
// @tags persons
// @ID add-person
// @Accept json
// @Produce json
// @Param person body models.Person true "query params"
// @Success 200 {string} string
// @Header 200 {string} string
// @Router /persons [post]
func AddPerson(c *gin.Context) {
	var person models.Person
	err := c.BindJSON(&person)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, err)
	}

	models.AddNewPersonObject(person)

	c.JSON(http.StatusOK, "person was added")
}
