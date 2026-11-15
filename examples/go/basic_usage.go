// HiTechCloud WAF SDK — Go example / Ví dụ Go
//
// Run: / Cách chạy:
//   go mod init example && go get github.com/hitechcloudvn/hitechcloud-waf-sdk/go
//   go run basic_usage.go
package main

import (
	"fmt"

	waf "github.com/hitechcloudvn/hitechcloud-waf-sdk/go"
)

func main() {
	client, err := waf.NewClient(waf.WithCredentials("admin", "********"))
	if err != nil {
		panic(err)
	}

	// List sites / Liệt kê site
	sites, err := client.Sites.List(map[string]string{"page": "1", "page_size": "20"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Sites:", sites)

	// Create an ACL rule / Tạo quy tắc ACL
	_, err = client.Acl.Create("123", map[string]interface{}{
		"type": "ip", "value": "203.0.113.10", "action": "deny",
	})
	if err != nil {
		panic(err)
	}

	// Read QPS statistics / Đọc thống kê QPS
	qps, _ := client.Statistics.Qps(nil)
	fmt.Println("QPS:", qps)

	// Check license status / Kiểm tra trạng thái giấy phép
	status, _ := client.License.Status(nil)
	fmt.Println("License:", status)
}
