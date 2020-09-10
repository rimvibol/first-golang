package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEmployees ... get All Employee
func GetEmployees(c *gin.Context) {
	var employee []Models.Employee
	err := Models.GetEmployees(&employee)
	if err != nil {
		fmt.Print("print IF")
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Print("print ElSE")
		c.JSON(http.StatusOK, employee)
	}
}

//CreateEmployee ... CreateEmployee
func CreateEmployee(c *gin.Context) {
	var employee Models.Employee
_:
	c.BindJSON(&employee)
	err := Models.AddEmployee(&employee)
	if err != nil {
		fmt.Print("print IF")
		fmt.Println(err.Error())
		//c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(200, employee)
	}

}
