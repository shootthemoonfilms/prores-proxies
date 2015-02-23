#!/bin/bash

echo -n " * Converting frames to TIF ..."
for F in *.dng; do
    FN=$( echo "$F" | rev | cut -c5-10 | rev )
    echo -n " $FN"
    ( dcraw -c -q 0 "$F" | pnmtotiff -quiet -lzw > ${FN}.tif ) 2>&1 >> /dev/null
done
echo " [done]"

FIRSTFRAME=$( ls -1 *.tif | head -1 | cut -d. -f1 )
PREFIXLEN=$( ls -1 *.dng | head -1 | wc -c )
PREFIX=$( ls -1 *.dng | head -1 | cut -c1-$(( ${PREFIXLEN} - 12 )) )

echo " * Outputting to ${PREFIX}.mov ... "
ffmpeg -f image2 -r 23.97 -start_number ${FIRSTFRAME} -i %6d.tif -y -vcodec prores -profile 3 "${PREFIX}.mov" && rm *.tif

