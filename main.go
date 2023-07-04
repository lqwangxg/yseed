package main

import (
	// "fmt"
	// "os"
	"github.com/lqwangxg/yseed/common"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//yseed()
	openServer()
}

// func yseed(){
// 	lenArgs :=len(os.Args)
// 	fmt.Printf("os.Args Length = %d, 3 integer numbers are required.\n", lenArgs) 
// 	if(lenArgs<3){
// 		return;
// 	}
// 	// to the console.
// 	msg := common.Seed(os.Args[1],os.Args[2],os.Args[3])
// 	fmt.Println(msg)
// }
func openServer(){
	e := echo.New()
    //===============================================
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      AllowOrigins: []string{"*"},
      AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
    }))
    //===============================================

    e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"Hello, get the prediction by /predict?n3=(n1,n2,n3)")
    })
    e.GET("/predict", common.Show)

    e.Logger.Fatal(e.Start(":8000"))
}
