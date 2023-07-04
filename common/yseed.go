package common

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"gopkg.in/yaml.v3"
)
type GuaYaml struct {
	Guas []Gua `yaml:"gua"`
}
type Gua struct {
	Id 	  int	 `yaml:"id"`
	Name  string `yaml:"name"`
	Mx	  int    `yaml:"mx"`
	Alias string `yaml:"alias"`
}

func Seed(x string, y string, z string) string {
	fmt.Printf("%s,%s,%s\n", x, y, z)
	
	top, err := mod(x, 8);
	if  (err != nil) {
		return fmt.Sprintf("x:%v",err)
	}
	bom, err := mod(y, 8);
	if  (err!=nil) {
		return fmt.Sprintf("y:%v",err)
	}
	yao, err := mod(z, 6);
	if  (err!=nil) {
		return fmt.Sprintf("z:%v",err)
	}
	
	index := top*10+bom
	g, err := GetGuaById(index)
	if  (err!=nil) {
		return fmt.Sprintf("z:%v",err)
	}
	return fmt.Sprintf("index=%d (%d,%d): %s %s卦 第%d爻.", index, top,bom, g.Alias, g.Name, yao)
}

func mod(p string, m int) (int, error) {
	x, err := strconv.Atoi(p); 
	if err != nil {
		return -1, ParamErr
	}
	
	fmt.Printf("before x=%d, m=%d ", x, m)
	if (x > m){
		x %= m
		if (x==0){
			x = m
		}
	}
	fmt.Printf("after mod x=%d, m=%d \n", x, m)
	return x, nil
}

func GetGuaById(mx int) (Gua, error) {
	var g Gua
	
	guas := readYml();	
	for _, g := range guas {
		if (g.Mx == mx){
			return g, nil
		}
	}
	return g, OutRangeErr
}

func readYml() []Gua {
	s, err := ioutil.ReadFile("./config.yml")
	if err != nil {                                    
        fmt.Println(err)                               
        return make([]Gua,0)
    }
	
	// define a variable as GuaYaml struct  
	var guaList GuaYaml
	// get data from byte(s) read from config.yml.
	yaml.Unmarshal([]byte(s), &guaList)
	// for loop the guas by range
	// for i, g := range guaList.Guas {
	// 	fmt.Printf("%d: id=%d, name=%s\n",i, g.Id, g.Name)
	// }
	return guaList.Guas
}
