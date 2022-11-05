# source: github.com/mysql/mysql-docker/blob/mysql-server/5.7/Dockerfile
# for setting can be found at dev.mysql.com/doc/refman/5.7/en/docker-mysql-more-topics.html#docker-persisting-data-configuration

# mount data into container, in order to update schema
docker pull mysql/mysql-server:5.7.24
docker run -p 3306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD='true' -e MYSQL_ROOT_PASSWORD='' -e MYSQL_ROOT_HOST=% \
    -v $(pwd)/scripts:/scripts \
    --name=red-packet-mysql -d mysql/mysql-server:5.7.24 \
    --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
