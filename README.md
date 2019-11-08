# clipload
> clipboard image upload made simple (macOS)

Requirements:
- `pngpaste` from Homebrew
- endpoint for uploading images

Building:
```
go build -ldflags "-X main.ApiUrl=https://image-api.tld/endpoint" clipload.go
```