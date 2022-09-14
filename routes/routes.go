package routes

import (
	"getAdvice/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Getadvices(c *gin.Context) {
	advices, err := models.Getalladvices()
	checkErr(err)

	if advices == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No records found!"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": advices})
	}

}

func GetAdviceid(c *gin.Context) {
	id := c.Param("id")

	advice, err := models.GeteachAdvice(id)
	checkErr(err)

	// if the description is blank we can assume nothing is found
	if advice.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": advice})
	}
}

func Addadvice(c *gin.Context) {
	var json models.Advice

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := models.Addadvcie(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func Updateadvice(c *gin.Context) {
	var json models.Advice

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := models.Updateanadvice(json, personId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func GetsurpriseAdvice(c *gin.Context) {
	advice, err := models.GetrandomAdvice()
	checkErr(err)

	// if the description is blank we can assume nothing is found
	if advice.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": advice})
	}
}
