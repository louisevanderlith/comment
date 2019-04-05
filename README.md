# comment
Mango API: Comment

## Run with Docker
* $ docker build -t avosa/comment:latest .
* $ docker rm commentDEV
* $ docker run -d -e RUNMODE=DEV -p 8084:8084 --network mango_net --name CommentDEV avosa/comment:latest
* $ docker logs commentDEV