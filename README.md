# comment
Mango API: Comment

## Run with Docker
*$ go build
*$ docker build -t avosa/comment:dev .
*$ docker rm commentDEV
*$ docker run -d -p 8084:8084 --network mango_net --name commentDEV avosa/comment:dev
*$ docker logs commentDEV