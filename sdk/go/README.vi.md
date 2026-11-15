# HiTechCloud WAF SDK (Go)

> [English](README.md) | **Tiếng Việt**

---

```bash
go get github.com/hitechcloudvn/hitechcloud-waf-sdk/go
```

```go
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
	sites, _ := client.Sites.List(nil)
	fmt.Println(sites)

	site, _ := client.Sites.Get("123", nil)
	fmt.Println(site)
}
```

Dùng `waf.WithSandbox()` để trỏ tới `https://vcwaf-sandbox.hitechcloud.vn`.
