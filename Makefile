rabbit-plugin:
	go build -o bin/rabbit-plugin .

release:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o bin/rabbit-plugin-linux-amd64 .
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -ldflags '-extldflags "-static"' -o bin/rabbit-plugin-linux-386 .
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -ldflags '-extldflags "-static"' -o bin/rabbit-plugin-linux-arm .
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -ldflags '-extldflags "-static"' -o bin/rabbit-plugin-linux-arm64 .
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -ldflags '-extldflags "-static"' -o bin/rabbit-plugin-windows-386.exe .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o bin/rabbit-plugin-windows-amd64.exe .
	CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -a -ldflags '-extldflags "-static"' -o bin/rabbit-plugin-darwin-386 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o bin/rabbit-plugin-darwin-amd64 .
