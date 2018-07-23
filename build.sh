rm -rf ./deploy
GOOS=linux go build -o ./deploy/hello
zip hello.zip deploy/hello