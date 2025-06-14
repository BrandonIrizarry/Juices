#!/usr/bin/env bash

BINS=('sitegen' 'server')

for bin in "${BINS[@]}"
do
    go build -o ./tmp/"$bin" ./cmd/"$bin"
done

