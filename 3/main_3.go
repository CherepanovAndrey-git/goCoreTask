package main

import (
	"3/pkg"
	"fmt"
)

func main() {

	testMap := pkg.NewMap()
	fmt.Println("OG map:", testMap)
	testMap.Add("key1", 12)
	testMap.Add("key2", 13)
	fmt.Println("OG map after adding elements:", testMap)

	value, exists := testMap.Get("key1")
	fmt.Printf("Get 'key1': value - %d, exists - %v\n", value, exists)
	fmt.Println("Exists 'key1:", testMap.Exist("key1"))

	copyMap := testMap.MapCopy()
	fmt.Println("Copied map:", copyMap)

	copyMap["key3"] = 34
	fmt.Println("Copied map after modification:", copyMap)
	fmt.Println("OG map:", testMap)

	_, existsInOg := testMap.Get("key3")
	fmt.Println("Exists 'key3' in original map:", existsInOg)
}
