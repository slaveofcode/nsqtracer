<img src="https://raw.github.com/slaveofcode/nsqtracer/main/img/logo.png" align="right" width="200" />

# NSQTracer
Simple NSQ tracer tool to see any messages that being published to specific NSQ `topic`. This library will not consume the message, only listening for new messages and display, that's it. However it also can consume the message by providing `--auto-finish` option, so the message will be assumed as finished.

<img src="https://raw.github.com/slaveofcode/nsqtracer/main/img/preview.png" align="center" />

### Installation
Simply go to the [latest binary release](https://github.com/slaveofcode/nsqtracer/releases) page to download the binary, and then run via CLI command. or if you has Go installed locally, NSQTracer can be installed via `go install`

**Tips:**
Place the binary on `/usr/local/bin` (for *\*nix* users) so it accessible anywhere in your system.

```
$ go install github.com/slaveofcode/nsqtracer
```

#### Start from default address
```
$ ./nsqtracer --topic SOMETOPIC
```

The command above will start a tracer listening from default `nsqd` host address, which is located on `localhost:4151`. 


#### Start from specific address
If you want to use specific address of NSQd, you can attach more options like below

```
$ ./nsqtracer --topic SOMETOPIC --nsqd-tcp localhost:4180
```

#### Start from multiple address
NSQTracer is able to listen from multiple `nsqd` addresses, so you can watch through all `nsqd` instances. Simply add more on the `--nsqd-tcp` options.

```
$ ./nsqtracer --topic SOMETOPIC --nsqd-tcp localhost:4150 --nsqd-tcp localhost:4140 --nsqd-tcp localhost:4180
```

### Using NSQLookupd address
By using `nsqlookupd` host address you'll discover all `nsqd`'s that connected to the `nsqlookupd` instance, it's really helpful when you want to listen from all `nsqd`'s by supplying only one `nsqlookupd` http address.

```
$ ./nsqtracer --topic SOMETOPIC --nsqlookup-http localhost:4161
```

The same rule applies for the `--nsqlookup-http` option, you can also provide more than one `nsqlookupd` addresses.

#### Read Available Options
Just run `./nsqtracer --help` to see all available commands


### LICENSE
MIT License

Copyright (c) 2020 Aditya Kresna
