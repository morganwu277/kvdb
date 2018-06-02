LevelDB based Net-KV-Store
---
# Motivation
I think most people can ignore this project since it's really meaningless. So timeline is: 
1. I do an app using leveldb as its store, since it will be quite large and need high fequrency of access, that's why I didn't choose RDBMS such as MySQL. 
2. One day, we need to scale the app. Boom... I have to make this part leveldb files able to be accessed from network! Here we began the story

# Tips
1. Trade off between response time and throughput
 - if you need throughput: try batch put, which means put lot of keys in one client, it will make latter keys response time increase
 - if you need response time, try not to put lot of keys at one time

# Credits

1. grpc
   https://github.com/grpc/grpc-go

2. leveldb
   - https://github.com/google/leveldb
   - https://github.com/google/leveldb/blob/master/doc/impl.md
