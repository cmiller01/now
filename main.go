package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.UTC().Format(time.RFC3339))
	fmt.Println(t.Unix())
}
