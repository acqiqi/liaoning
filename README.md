# vge

git push -u origin master
export GO111MODULE=on 
export GOPROXY=https://goproxy.io 
go mod vendor

GOOS=linux GOARCH=mipsle GOMIPS=softfloat  go build main.go

scp -r main root@192.168.8.1:/tmp
scp -r cfg.json root@192.168.8.1:/tmp