#!/usr/bin/env bash

if [[ $# -eq 0 ]] ; then
	for dir in */; do \
		(cd $dir && make image)
	done
else
	for dir in "$@"; do \
		(cd $dir && make image)
	done
fi
docker system prune -f
