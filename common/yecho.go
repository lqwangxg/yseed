package common

import (
	"regexp"
	"net/http" 
	"github.com/labstack/echo"
)

func Show(c echo.Context) error {
	//n3 := strings.Split(c.QueryParam("n3"),",")
	n3 := c.QueryParam("n3")
	re := regexp.MustCompile(`,\s*`)
	ss := re.Split(n3, -1)  //=> ["AA", "BB", "CC"]

	x :=ss[0]
	y :=ss[1]
	z :=ss[2]
	result :=Seed(x,y,z)	
	return c.String(http.StatusOK, result)
}