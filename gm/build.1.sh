GOARCH=amd64 GOOS=linux go build ./gamegm/gm/main.go
cp ./main ../部署/万世/gm/gm
cd ./gmweb && npm run build:prod
cp -r ./dist/* ../../部署/万世/gm/public/
