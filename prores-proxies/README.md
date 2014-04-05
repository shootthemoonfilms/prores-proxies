# prores-proxies

Quick and dirty ffmpeg standalone proxy conversion.

## Usage

```
Usage of ./prores-proxies:
  -ffmpeg="./ffmpeg": Path to FFMPEG binary
  -proxy="proxy": Proxy files subdirectory name
```

If no additional arguments are given, the current working directory will be
processed, otherwise all listed directories will be scanned and all .mov
files will be processed.

