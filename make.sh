#!/bin/bash

#rm -rf pkg
cd src/util/ && go build && go install && cd ../..
cd src/tool/ && go build && go install && cd ../..
cd src/timer/ && go build && go install && cd ../..
cd src/config/ && go build && go install && cd ../..
cd src/glog/ && go build && go install && cd ../..
cd src/golog/ && go build && go install && cd ../..
cd src/console/ && go build && go install && cd ../..
cd src/protocol/ && go build && go install && cd ../..
cd src/driver/ && go build && go install && cd ../..
cd src/ && go build -o ../bin/mornsail_exec .
#go install
