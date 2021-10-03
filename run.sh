#go mod init app
#!/bin/bash
go mod download
go build -o app .
sleep 30
./app

echo "App is running, try localhost:7070 or <public-ip>:7070"
