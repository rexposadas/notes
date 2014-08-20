docker stop logstash
docker rm logstash
docker run -d --name logstash -link es:es -p 514:514 -p 9292:9292 -p 7000:7000 -e LOGSTASH_CONFIG_STRING='input {tcp {port => 7000}} output {elasticsearch {host => "172.17.42.1"}}' rexposadas/docker-logstash
