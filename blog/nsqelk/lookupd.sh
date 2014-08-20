# run nsqlookupd
docker stop lookupd
docker rm lookupd
docker run -d --name lookupd -p 4160:4160 -p 4161:4161 mreiferson/nsqlookupd
