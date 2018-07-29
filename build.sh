echo Deleting deploy folder
rm -rf ./deploy

echo building binary
GOOS=linux go build -o ./deploy/hello

echo zipping
cd deploy
zip hello.zip hello