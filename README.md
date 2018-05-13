## Preface
this project is forked from [golang/protobuf](https://github.com/golang/protobuf)

all my codes are in "my" branche.

## Features

* contain all features from the official project
* autogenerate the struct's MOCK value.

For instance, there is proto definition like:

```proto
// message definition in .proto
message KK {
  message Love {
    int64 duration=1;
  }

  Love status=1;
  int32 age=2;
  double height=3;
  string name=4;
  bytes talk=5;
  bool isloving=6;
}
```

And the corresponding golang sdk would be like:
```go
type KK struct {
	Status               *KK_Love `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Height               float64  `protobuf:"fixed64,3,opt,name=height" json:"height,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Talk                 []byte   `protobuf:"bytes,5,opt,name=talk,proto3" json:"talk,omitempty"`
	Isloving             bool     `protobuf:"varint,6,opt,name=isloving" json:"isloving,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// mock parts
var KKMock = KK{
	Status:   &KK_LoveMock,
	Age:      2018,
	Height:   2018.0513,
	Name:     "abcdefg",
	Talk:     nil,
	Isloving: true,
}

type KK_Love struct {
	Duration             int64    `protobuf:"varint,1,opt,name=duration" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// mock parts
var KK_LoveMock = KK_Love{
	Duration: 2018,
}
```

## install

premonition: the installition would overwrite the official protoc-gen-go!!!

```bash
go get -v github.com/chalvern/protobuf

mv $GOPATH/src/github.com/chalvern $GOPATH/github.com/golang

cd $GOPATH/src/github.com/golang/protobuf/protoc-gen-go

go install -v
```