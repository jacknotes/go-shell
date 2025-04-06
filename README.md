# build 
```bash
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/go-shell-linux-amd64 main.go
```


# running
```bash
chmod a+x dist/go-shell-linux-amd64
dist/go-shell-linux-amd64 start -f code.toml
```