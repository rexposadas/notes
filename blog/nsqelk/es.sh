docker stop es
docker rm es
docker run -d --name es -p 9200:9200 -p 9300:9300 dockerfile/elasticsearch
