#download dependencies
go mod vendor

docker build --platform linux/amd64 -t vm-backend:0.0.0 .


# test
docker run -it vm-backend:0.0.0