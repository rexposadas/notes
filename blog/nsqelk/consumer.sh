# consume messages from nsq and send it to LogStash
# One consumer per topic.  
docker stop consumer
docker rm consumer
docker run -d --name consumer -e TOPIC=test -e LOOKUPD_ADDR=172.17.42.1:4161 -e OUTPUT_TCP_ADDR=172.17.42.1:7000 rexposadas/docker-nsq-to-tcp
