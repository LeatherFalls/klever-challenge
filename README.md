# Golang gRPC upvote service
* A grpc cryptocurencies upvote service.

## Start:

* Install [Go](https://go.dev/dl/)

* Clone this repository
```
- http: git clone https://github.com/LeatherFalls/klever-challenge.git
- ssh: git clone git@github.com:LeatherFalls/klever-challenge.git
```

* Run docker-compose
```
  docker-compose up -d
```
* Build
```
  make build
```
* Start server
```
  make server
```
* Start client
```
  make client
```
* Remove generated files
```
  make prune
```

## gRPC Methods:

```
- rpc Create(Crypto) returns (CryptoId);
- rpc ListAll(google.protobuf.Empty) returns (stream Crypto);
- rpc ListById(CryptoId) returns (Crypto);
- rpc Update(Crypto) returns (google.protobuf.Empty);
- rpc Delete(CryptoId) returns (google.protobuf.Empty);
- rpc UpvoteCrypto(CryptoId) returns (google.protobuf.Empty);
```
## Built with
- [Go](https://go.dev/dl/)
- [gRPC](https://grpc.io/)
- [MongoDB](https://www.mongodb.com/)
- [Docker](https://www.docker.com/)
