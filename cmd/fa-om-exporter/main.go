package main

import (
	"fmt"
	"purestorage/fa-openmetrics-exporter/internal/rest-client"
)


func main() {
	c := client.NewRestClient("10.225.112.90", "b5cb29e7-a93c-b40a-b02f-da2b90c8c65e", "latest")
        defer c.Close()
        var cli client.Client
        cli = c
        al := cli.GetArrays()
        for _, a := range al.Items {
		fmt.Println(a)
        }
}
