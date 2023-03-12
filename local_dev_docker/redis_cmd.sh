#!/bin/bash

service="red-packet-backend-redis"

while getopts ":c:r:" opt; do
  case $opt in
    c) command="$OPTARG"
    ;;
    r) redis_cli="$OPTARG"
    ;;
    \?) echo "Invalid option -$OPTARG" >&2
    exit 1
    ;;
  esac

  case $OPTARG in
    -*) echo "Option $opt needs a valid argument"
    exit 1
    ;;
  esac
done

# echo "[DEBUG] command = $command"
# echo "[DEBUG] redis_cli = $redis_cli"

if [ $command = "start" ]
then
  docker start $service
  if [ $redis_cli ]
  then
    docker exec -it $service redis-cli
  fi
elif [ $command = "stop" ]
then
  docker stop $service
fi