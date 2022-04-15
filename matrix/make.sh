#!/bin/sh

gox -os=linux -arch=amd64
mv matrix_linux_amd64 matrix
docker build -t 172.16.84.121/dtstack-dev/matrix:feat_4.3.3_fenle .
docker push 172.16.84.121/dtstack-dev/matrix:feat_4.3.3_fenle

