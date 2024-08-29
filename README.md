# cent-inout

read centrifugo inout msg and parse them with protobuf


# usage

- config centrifugo to use nats as broker
- read nats topic from nats

```bash
./cent-inout --nats-address=127.0.0.1:4222
```