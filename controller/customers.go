package customer

import (
	"fullstack/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customers struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Age  int    `json:"age" gorm:"column:age;"`
}

func GetCustomers(c *gin.Context) {
	var customerList []Customers
	if err := db.DB.Table("customers").Find(&customerList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"payload": customerList,
	})
}
