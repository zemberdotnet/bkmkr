package main

import (
	"context"
)


func main() {
	ctx := context.Background()
	client, err := connectInit(ctx)
	if err != nil {
		panic(err)
	}
	Read(client,ctx)	

	defer connectClose(client)
}

