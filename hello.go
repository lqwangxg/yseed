package main

import (
	"fmt"
	"os"
	"hello/common"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	//fmt.Printf("os.Args=%v, idx=%s \n", os.Args, idx)   
	// if g, err := common.GetGuaById(idx); err == nil {                                    
    //     fmt.Printf("id=%d, name=%s\n", g.Id, g.Name)                              
    // } else{
	// 	fmt.Printf("os.Args=%v, err: %v\n", os.Args, err)   
	// }
	//yseed()
	openServer()
}
func yseed(){
	lenArgs :=len(os.Args)
	fmt.Printf("os.Args Length = %d, 3 integer numbers are required.\n", lenArgs) 
	if(lenArgs<3){
		return;
	}
	// to the console.
	msg := common.Seed(os.Args[1],os.Args[2],os.Args[3])
	fmt.Println(msg)
}
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
    // e.GET("/users/:id", getUser)
    // e.PUT("/users/:id", updateUser)
    // e.DELETE("/users/:id", deleteUser)

    e.Logger.Fatal(e.Start(":8000"))
}
