package main

import (
	"fmt"
	"os"
)

const clusterName = "CLUSTER_NAME"

func main(){
	fmt.Println("Hello from DOGE 🐩🐩 action")
	x := GetEnvAsStringOrFallback(clusterName, "ERROR")
	fmt.Println(x)
}



// GetEnvAsStringOrFallback returns the env variable for the given key
// and falls back to the given defaultValue if not set
func GetEnvAsStringOrFallback(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}