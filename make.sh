#!/bin/bash

rm -rf pkg
cd src/timer/ && go build && go install && cd ../..
cd src/config/ && go build && go install && cd ../..
cd src/glog/ && go build && go install && cd ../..
cd src/golog/ && go build && go install && cd ../..
cd src/console/ && go build && go install && cd ../..
cd src/protocol/ && go build && go install && cd ../..
cd src/servlet/ && go build && go install && cd ../..
cd src/player/ && go build && go install && cd ../..
cd src/driver/ && go build && go install && cd ../..
cd src/ && go build -o ../bin/mornsail_exec mornsail
#go install