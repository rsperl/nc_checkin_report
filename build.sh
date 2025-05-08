#!/usr/bin/env bash

function build {
    local os=$1
    local arch=$2
    local bin_dir="bin/$os/$arch"
    mkdir -p "$bin_dir"
    local binary_name="$bin_dir/nc_checkin_report"
    if [ "$os" == "windows" ]; then
        binary_name="$binary_name.exe"
    fi
    echo "Building for $os/$arch"
    env GOOS=$os GOARCH=$arch go build -o "$binary_name" *.go
    echo "Creating tarball for $os/$arch"
    tar -czvf bin/nc_checkin_report_${os}_${arch}.tar.gz -C "$bin_dir" .
}

for arch in amd64 arm64; do
    for os in windows darwin; do
        echo "Building for $os/$arch"
        bin_dir="bin/$os/$arch"
        mkdir -p "$bin_dir"
        binary_name="$bin_dir/nc_checkin_report"
        if [ "$os" == "windows" ]; then
            binary_name="$binary_name.exe"
        fi
        build "$os" "$arch" &
    done
done

wait
echo "Build complete. Binaries are located in the bin directory."
