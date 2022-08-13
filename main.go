package main

import (
	"net/http"

	_ "github.com/blennster/swag-repro/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Struct1 struct {
	Field1 string `json:"field1"`
}

type Struct2 struct {
	Struct1
	Field2 string `json:"field2"`
}

type Struct3 struct {
	Field3 []Struct2 `json:"field3"`
}

// @success 200 {object} Struct3{field3=[]Struct1}
// @router /ex1 [get]
func Response(c *gin.Context) {
	s2 := []Struct2{
		{
			Struct1: Struct1{
				Field1: "I am first",
			},
		},
		{
			Struct1: Struct1{
				Field1: "I am second",
			},
		},
	}

	res := Struct3{
		Field3: s2,
	}

	c.JSON(http.StatusOK, res)
}

// @success 200 {object} Struct3{field3=[]string}
// @router /ex2 [get]
func OtherReponse(c *gin.Context) {
	Response(c)
}

func main() {
	r := gin.Default()
	r.GET("/ex1", Response)
	r.GET("/ex2", OtherReponse)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
