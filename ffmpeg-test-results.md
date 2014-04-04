# FFMPEG PRORES -> H.264 TEST RESULTS

```
Test Clip: 7.76 sec, 178 MB (sourced from BMPCC)
Original at https://www.dropbox.com/sh/aitqd5tddqu3zi4/GCJRv7pSie
```

## MPEG4 VCODEC TEST

Using:
``time ffmpeg -y -i G008_1_2013-07-27_1514_C0001.mov -vcodec mpeg4 -acodec aac -strict -2 G008_1_2013-07-27_1514_C0001.mp4``

```
video:1876kB audio:122kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 0.281198%

real	0m6.530s
user	0m16.760s
sys	0m0.180s
```

## LIBX264 DEFAULTS TEST

Using:
``time ffmpeg -y -i G008_1_2013-07-27_1514_C0001.mov -vcodec libx264 -acodec aac -strict -2 G008_1_2013-07-27_1514_C0001.mp4``

```
video:3155kB audio:122kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 0.214723%

real	0m28.846s
user	1m41.855s
sys	0m0.658s
```

Quality:
Approaches original ProRes quality

## LIBX264 CRF51 EXTREME LOW QUALITY TEST

Using:
``time ffmpeg -y -i G008_1_2013-07-27_1514_C0001.mov -vcodec libx264 -acodec aac -strict -2 -crf 51 G008_1_2013-07-27_1514_C0001.mp4``

```
video:147kB audio:122kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 2.618429%

real	0m13.215s
user	0m44.814s
sys	0m0.502s
```

Quality:
Crap. Really terrible.

## LIBX264 ULTRAFAST SETTING TEST

Using:
``time ffmpeg -y -i G008_1_2013-07-27_1514_C0001.mov -vcodec libx264 -acodec aac -strict -2 -preset ultrafast G008_1_2013-07-27_1514_C0001.mp4``

```
video:7353kB audio:122kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 0.073811%

real	0m5.880s
user	0m18.954s
sys	0m0.257s
```

Quality:
Very good -- almost as nice as the original HQ test, but much faster.

