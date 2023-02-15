#!/bin/bash
set -e
PID=$(lsof -i:8848 | grep LISTEN | awk '{print $2}')

if [ -n "$PID" ]; then
  kill -9 "$PID"
fi

go build -o main .
nohup ./main > nohup.log 2>&1 &