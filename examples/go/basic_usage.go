package main

import (
	"fmt"

	waf "github.com/hitechcloud-vietnam/hitechcloud-waf-sdk/go"
)

func main() {
	client, err := waf.NewClient(waf.WithCredentials("admin", "********"))
	if err != nil {
		panic(err)
	}

	sites, err := client.Sites.List(map[string]string{"page": "1", "page_size": "20"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Sites:", sites)

	_, err = client.Acl.Create("123", map[string]interface{}{
		"type": "ip", "value": "203.0.113.10", "action": "deny",
	})
	if err != nil {
		panic(err)
	}

	qps, _ := client.Statistics.Qps(nil)
	fmt.Println("QPS:", qps)

	status, _ := client.License.Status(nil)
	fmt.Println("License:", status)
}
