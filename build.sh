#!/usr/bin/env bash

for arch in amd64 arm64; do
    for os in windows darwin; do
        echo "Building for $os/$arch"
        bin_dir="bin/$os/$arch"
        mkdir -p "$bin_dir"
        binary_name="$bin_dir/nc_checkin_report"
        if [ "$os" == "windows" ]; then
            binary_name="$binary_name.exe"
        fi
        env GOOS=$os GOARCH=$arch go build -o "$binary_name" *.go &
    done
done

wait
echo "Build complete. Binaries are located in the bin directory."
