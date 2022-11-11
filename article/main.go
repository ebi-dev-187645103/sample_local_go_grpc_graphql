package main

import (
	"fmt"

	"github.com/ebi-dev-187645103/sample_local_go_grpc_graphql/article/client"
)


func main() {
	fmt.Println("start gRPC Client.")

	port := "8080"
	// c,err := client.NewClient(port)
	err := client.NewClient(port)
	if err != nil{
		fmt.Println(err)
	}else{
		// 4. 実行
		// c.Hello()
		fmt.Println("end gRPC Client.")
	}
}
