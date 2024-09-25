#!/bin/bash

docker run --name testdb -e POSTGRES_PASSWORD=mysecretpassword -d postgres

# docker run -d --name=roach1 --hostname=roach1 --net=roachnet -p 26257:26257 -p 8080:8080 -v "roach1:/cockroach/cockroach-data" cockroachdb/cockroach:v24.2.2 start --advertise-addr=roach1:26357 --http-addr=roach1:8080 --listen-addr=roach1:26357 --sql-addr=roach1:26257 --insecure --join=roach1:26357,roach2:26357,roach3:26357
