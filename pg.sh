#!/bin/bash

docker run --name testdb -e POSTGRES_PASSWORD=mysecretpassword -d postgres
