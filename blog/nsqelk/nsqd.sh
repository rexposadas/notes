docker stop nsqd
docker rm nsqd
docker run -d --name nsqd -p 4150:4150 -p 4151:4151 -e BROADCAST_ADDRESS=172.17.42.1 -e LOOKUPD_ADDRESS=172.17.42.1:4160 rexposadas/nsqd
