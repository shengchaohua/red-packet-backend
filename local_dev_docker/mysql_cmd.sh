#!/bin/bash

service="red-packet-backend-mysql"

while getopts ":c:m:" opt; do
  case $opt in
    c) command="$OPTARG"
    ;;
    m) mysql_client="$OPTARG"
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
# echo "[DEBUG] mysql_client = $mysql_client"

if [ $command = "start" ]
then
  docker start $service
  if [ $mysql_client ]
  then
    sleep 1 # to avoid "ERROR 2002 (HY000): Can't connect to local MySQL server through socket '/var/lib/mysql/mysql.sock'"
    docker exec -it $service mysql -uroot
  fi
elif [ $command = "stop" ]
then
  docker stop $service
fi
