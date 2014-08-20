docker stop logstash-stdout
docker rm logstash-stdout
docker run --name logstash-stdout -p 514:514 -p 9292:9292 -p 7000:7000 -e LOGSTASH_CONFIG_STRING='input {tcp {port => 7000}} output {stdout { }}' rexposadas/docker-logstash
