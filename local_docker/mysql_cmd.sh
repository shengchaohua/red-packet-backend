#!/bin/bash

#mysql_master="red_packet_backend_mysql_master"
#mysql_slave="red_packet_backend_mysql_slave"

# mysql master shell
docker exec -it red_packet_backend_mysql_master mysql -uroot -p123456

# mysql slave shell
docker exec -it red_packet_backend_mysql_slave mysql -uroot -p123456
