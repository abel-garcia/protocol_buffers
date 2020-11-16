# Protocol Buffers

### What is Protocol Buffers?.

Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data think XML, but smaller, faster, and simpler.

#### Content.
```bash
.
├── README.md
├── book
│   ├── person.pb.go
│   └── person_grpc.pb.go
├── go.mod
├── go.sum
├── main.go
└── person.proto
```
#### Messages definition

```protobuf
message Person {
  int32 id = 1;
  string name = 2;
  string last_name = 3;
  int32 age = 4;
  string email = 5;

  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }

  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }

  repeated PhoneNumber phones = 6;
  google.protobuf.Timestamp last_updated = 7;
}
```

#### Services definition
```protobuf
service AddressBook {
    rpc Get(Person) returns(Person) {}
    rpc Put(Person) returns(Person) {}
    rpc Del(Person) returns(Person) {}
}
```

#### Use the protocol buffer compiler.

```bash
> protoc --go_out=book/. --go_opt=paths=source_relative \
> --go-grpc_out=book/. --go-grpc_opt=paths=source_relative \
> *.proto
```

#### Run Go protocol buffer API to write and read messages

```bash
> go run main.go
```

### Output in bash

```bash
Bytes data:  [8 1 18 4 65 98 101 108 26 6 71 97 114 99 105 97 32 28 42 25 97 98 101 108 103 97 114 99 105 97 51 56 51 52 56 64 103 109 97 105 108 46 99 111 109 50 12 10 8 53 53 53 45 52 51 50 49 16 1 58 12 8 196 250 199 253 5 16 224 143 211 228 2]
=============================
ID:  1
Name:  Abel
Last Name:  Garcia
Age:  28
Email:  abelgarcia38348@gmail.com
Phones:  [number:"555-4321"  type:HOME]
```
