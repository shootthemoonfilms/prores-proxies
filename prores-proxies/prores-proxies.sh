#!/bin/bash
#
#	prores-proxies.sh
#	@jbuchbinder
#
#	Convert all files in the current directory to h.264 proxies in
#	a subdirectory called "proxy".
#

ffmpeg="/usr/bin/ffmpeg"
subdir="proxy"

for i in *.mov *.MOV; do
	o="$(dirname "${i}")/${subdir}/$(basename "${i}")"
	mkdir -p "$(dirname "$o")"
	echo "$(date)| Converting $i -> $o"
	time ${ffmpeg} \
		-y \
		-i "$i" \
		-vcodec libx264 -acodec aac -strict -2 -preset ultrafast \
		"${o}"
done
