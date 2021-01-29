package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)
	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]int      //m3 == nil
	fmt.Println(m, m2, m3)
	fmt.Println("travering map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Getting values")

	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName, ok)
	}
	fmt.Println("delete")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
