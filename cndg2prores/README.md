# cdng2prores

Quick and dirty dcraw/ffmpeg standalone CinemaDNG to prores proxy conversion.

## Usage

```
Usage of ./cdng2prores:
  -dcraw="./dcraw": Path to DCRAW binary
  -extension="mov": File extension
  -ffmpeg="./ffmpeg": Path to FFMPEG binary
  -format-extension="mov": File extension used for determining format
  -pnmtotiff="./pnmtotiff": Path to PNMTOTIFF binary
  -proxy="proxy": Proxy files subdirectory name
  -scaleh=0: Scale height
  -scalew=0: Scale width
```

If no additional arguments are given, the current working directory will be
processed, otherwise all listed directories will be scanned and all .mov
files will be processed.

If scalew and scaleh are specified, the video will be rescaled to that size.

If extension is specified, that extension will be used for output files. This
can be useful when dealing with annoying hacks like Premiere not dealing with
Quicktime MOV files properly when using the Quicktime decoder -- the
workaround for which is to use ```.mpg``` as the file extension for the
proxy files.

The threading option allows parallelism (more than one file to be encoded
at once). It is disabled by default.

## Dependencies

 * [Go](http://golang.org) - for compilation
 * ffmpeg - Binary required to execute. This should have prores support compiled in.
   - [Windows ffmpeg binaries](http://ffmpeg.zeranoe.com/builds/)
   - [ffmpeg source](https://github.com/FFmpeg/FFmpeg)
 * dcraw - Requred for RAW frame to PPM conversion
 * ppmtotiff - Required for PPM to TIFF conversion
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
