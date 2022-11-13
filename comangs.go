package main

import (
	"fmt"
)

func getCancelButtonData() string {
	return "c"
}

func getFileButtonData(filename string) string {
	return fmt.Sprintf("f:%s", filename)
}
