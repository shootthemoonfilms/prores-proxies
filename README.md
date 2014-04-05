# ProRes Proxies

## Singleton Processing

See ``prores-proxies`` directory for README file.

## Distributed Processing

Uses gearman to create ProRes proxies using a series of distributed machines.

### Requirements

 * ffmpeg - Compiled with libx264 and prores support
 * golang - For compilation -- and for fun!

### Building

```
go get -d
go build
```

