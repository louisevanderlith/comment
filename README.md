# comment
Mango API: Comment
Any comments made on the system, should all be controlled by the Comment API.

## Run with Docker
* $ docker build -t avosa/comment:dev .
* $ docker rm commentDEV
* $ docker run -d -e RUNMODE=DEV -p 8084:8084 --network mango_net --name CommentDEV avosa/comment:dev
* $ docker logs commentDEV