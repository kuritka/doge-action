package main

import (
	"fmt"
	"github.com/AbsaOSS/gopkg/env"
)

const clusterName = "CLUSTER_NAME"

func main(){
	fmt.Println("Hello from DOGE 🐩🐩 action")
	x := env.GetEnvAsStringOrFallback(clusterName, "ERROR")
	fmt.Println(x)
}
