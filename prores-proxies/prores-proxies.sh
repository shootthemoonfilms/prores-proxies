#!/bin/bash
#
#	prores-proxies.sh
#	@jbuchbinder
#
#	Convert all files in the current directory (or a specified set of
#	directories) to h.264 proxies in subdirectories of the directory(s)
#	containing the source files, called "proxy".
#

ffmpeg="/usr/bin/ffmpeg"
subdir="proxy"

files=""
if [ $# -gt 0 ]; then
	# Use args
	for d in $*; do
		files="$files $d/*.mov"
	done
else
	files="*.mov"
fi

for i in $files; do
	if [ -f "$i" ]; then
		o="$(dirname "${i}")/${subdir}/$(basename "${i}")"
		mkdir -p "$(dirname "$o")"
		echo "$(date)| Converting $i -> $o"
		time ${ffmpeg} \
			-y \
			-i "$i" \
			-vcodec libx264 -acodec aac -strict -2 -preset ultrafast \
			"${o}"
		echo "$(date)| Completed proxy : $o"
	else
		echo "$(date)| Skipping non-existing file ${i}"
	fi
done