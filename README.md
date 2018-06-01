LevelDB based Net-KV-Store
---

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
