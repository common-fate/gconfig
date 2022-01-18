# gconfig

Config for Granted

## Serialisation

Config is serialized to be stored as binary blobs using Protobuf. To serialise a Config into binary, run:

```go
c := Config{}
pb := c.SerializeProtobuf()
// marshal to b ([]byte type)
b, err := proto.Marshal(pb)
```

To deserialize a Config from binary, run:

```go
// load `b` []byte from a database, or somewhere else
providers := &gconfigv1alpha1.Providers{} // providers must be provided separately to rehydrate the config

var pb gconfigv1alpha1.Config
err = proto.Unmarshal(b, &pb)
c := FromProtobuf(&pb, providers)
```
