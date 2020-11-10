<!--

    SOURCE: https://itnext.io/learning-go-mongodb-crud-with-grpc-98e425aeaae6

    #1. PROTO FILE
    ---------------------------
        to compile proto file
        > protoc proto/blog.proto --go_out=plugins=grpc:.



    #2. SERVER IMPLEMENTATION
    ---------------------------
        First things first, we’ll need some boilerplate to create a gRPC server and mongoDB connection, give some feedback to the console and handle server shutdowns by the user through a shutdown hook.
 -->
