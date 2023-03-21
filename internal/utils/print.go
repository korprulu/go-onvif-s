package utils

import (
	"encoding/json"
	"fmt"
)

// PrintJSON ...
func PrintJSON(obj any) {
	out, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
}
