# Đóng góp cho HiTechCloud WAF SDK

> [English](CONTRIBUTING.md) | **Tiếng Việt**

---

Cảm ơn bạn đã quan tâm đến việc đóng góp!

## Phát triển

1. Clone kho mã.
2. Chọn một SDK ngôn ngữ dưới `sdk/`.
3. Cài đặt phụ thuộc và chạy các kiểm tra:
   - Python: `python -m py_compile hitechcloud_waf/**/*.py`
   - Node.js: `npm run lint` (tsc --noEmit)
   - PHP: `php -l src/**/*.php`
   - Go: `gofmt -l .` / `go build ./...`
   - Java: `mvn compile`

## Nguyên tắc

- Mọi phương thức phải ánh xạ tới một endpoint WAF thật.
- Giữ cách nhóm tài nguyên và đặt tên nhất quán giữa các ngôn ngữ.
- Cập nhật `openapi.json` và tài liệu khi endpoint thay đổi.
- Tài liệu hoá các phương thức công khai.

## Pull Request

1. Fork và tạo một nhánh tính năng.
2. Thêm hoặc cập nhật ví dụ/kiểm tra.
3. Đảm bảo mọi kiểm tra theo từng ngôn ngữ đều đạt.
4. Gửi PR.

Khi đóng góp, bạn đồng ý rằng công việc của bạn được cấp phép theo Giấy phép MIT.
