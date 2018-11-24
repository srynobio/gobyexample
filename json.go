package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type responce1 struct {
	Page   int
	Fruits []string
}

type responce2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var pf = fmt.Println

func main() {

	bolB, _ := json.Marshal(true)
	pf(string(bolB))

	intB, _ := json.Marshal(1)
	pf(string(intB))

	fltB, _ := json.Marshal(2.34)
	pf(string(fltB))

	strB, _ := json.Marshal("gopher")
	pf(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	pf(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	pf(string(mapB))

	res1D := &responce1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	pf(string(res1B))

	res2D := &responce2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	pf(string(res2B))

	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)

	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	pf(dat)

	num := dat["num"].(float64)
	pf(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	pf(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	pf("str: ", str)
	res := responce2{}
	json.Unmarshal([]byte(str), &res)
	pf(res)
	pf(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
