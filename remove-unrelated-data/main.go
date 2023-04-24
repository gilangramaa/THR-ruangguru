package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
}

func main() {
	dataMap1 := map[string]interface{}{
		"name":    "Edo",
		"age":     20,
		"address": "Jakarta",
	}
	fmt.Println(removeUnrelated(dataMap1, "address"))

	dataMap2 := map[string]interface{}{
		"run":  "lari",
		"jump": "loncat",
		"swim": "berenang",
	}
	fmt.Println(removeUnrelated(dataMap2, "jump"))

	dataMap3 := map[string]interface{}{
		"satu":  "ichi",
		"dua":   "ni",
		"tiga":  "san",
		"empat": "yon",
		"lima":  "go",
	}
	fmt.Println(removeUnrelated(dataMap3, "tiga"))
}
