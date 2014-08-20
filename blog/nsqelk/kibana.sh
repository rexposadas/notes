docker stop kibana
docker rm kibana
docker run -d --name kibana -p 80:80 -e ES_HOST=172.17.42.1 -e ES_PORT=9200 arcus/kibana
