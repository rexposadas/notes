#!/bin/bash
bash dockerclean.sh
bash es.sh 
bash logstash.sh
bash kibana.sh
bash consumer.sh
bash lookupd.sh
bash nsqd.sh
