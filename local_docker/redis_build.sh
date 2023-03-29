# source: hub.docker.com/_/redis
docker pull redis:3.2.0
docker run -p 6379:6379 --name red-packet-backend-redis -d redis:3.2.0