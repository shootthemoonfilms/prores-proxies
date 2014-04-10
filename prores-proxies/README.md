# prores-proxies

Quick and dirty ffmpeg standalone proxy conversion.

## Usage

```
Usage of ./prores-proxies:
  -ffmpeg="./ffmpeg": Path to FFMPEG binary
  -proxy="proxy": Proxy files subdirectory name
  -scaleh=0: Scale height
  -scalew=0: Scale width

```

If no additional arguments are given, the current working directory will be
processed, otherwise all listed directories will be scanned and all .mov
files will be processed.

If scalew and scaleh are specified, the video will be rescaled to that size.

## Dependencies

 * [Go](http://golang.org) - for compilation
 * ffmpeg - Binary required to execute. This should have libx264 support compiled in.
   - [Windows ffmpeg binaries](http://ffmpeg.zeranoe.com/builds/)
   - [ffmpeg source](https://github.com/FFmpeg/FFmpeg)
 * [golang-crosscompile](https://github.com/davecheney/golang-crosscompile) - for cross-compiling for other platforms

## Building

```
go build
```

To cross-compile, make sure that you have installed and setup
``golang-crosscompile`` first. Then run

```
make
```

