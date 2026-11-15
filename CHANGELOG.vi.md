# Nhật ký thay đổi

> [English](CHANGELOG.md) | **Tiếng Việt**

---

Mọi thay đổi đáng chú ý của HiTechCloud WAF SDK đều được ghi lại trong tệp này.
Định dạng dựa trên [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
và dự án tuân theo [Semantic Versioning](https://semver.org/).

## [1.0.1] - 2026-11-16

### Đã sửa
- SDK Java: `WafClient` giờ tham chiếu tường minh
  `vn.hitechcloud.waf.resources.System` - tên `System` bị mơ hồ với
  `java.lang.System` nên build Maven thất bại.
- SDK Node.js: thêm script `npm test` còn thiếu và smoke test để bước test
  chạy thay vì fail với "Missing script: test".
- SDK PHP: thêm `phpunit.xml.dist` và smoke test để `vendor/bin/phpunit`
  chạy bộ test thay vì in help và fail job test.

- SDK Node.js: chạy test runner tích hợp bằng `node --test` không kèm đối số
  đường dẫn để cùng một lệnh chạy được trên Node.js 18, 20 và 22.
- SDK Python: thêm smoke test pytest để `pytest` có bộ test thật (bộ test rỗng
  thoát mã 5 và fail job test).

## [1.0.0] - 2026-11-01

### Thêm
- Phát hành đầu tiên với 5 SDK ngôn ngữ (Node.js/TypeScript, Python, PHP, Go, Java).
- 227 endpoint thuộc 25 nhóm tài nguyên.
- Luồng xác thực tự động CSRF -> login -> JWT với TFA tuỳ chọn.
- Đặc tả OpenAPI 3.0.
- Tài liệu tiếng Việt.
- Máy chủ MCP cung cấp các thao tác WAF cốt lõi.
- Ví dụ cho mọi ngôn ngữ được hỗ trợ.
