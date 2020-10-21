#!/bin/sh
rm -f nsqtracer_linux* nsqtracer_openbsd* nsqtracer_darwin* nsqtracer_android* nsqtracer_*.exe

echo "compiling linux/amd64..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC=gcc go build
mv nsqtracer nsqtracer_linux_amd64
echo "done."

echo "compiling linux/386..."
GOOS=linux GOARCH=386 CGO_ENABLED=1 go build
mv nsqtracer nsqtracer_linux_386
echo "done."

echo "compiling windows/386..."
GOOS=windows GOARCH=386 CGO_ENABLED=1 CC=i686-w64-mingw32-gcc go build
mv nsqtracer.exe nsqtracer_32.exe
echo "done."

echo "compiling windows/amd64..."
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build
mv nsqtracer.exe nsqtracer_64.exe
echo "done."

echo "compiling linux/arm..."
GOOS=linux GOARCH=arm go build
mv nsqtracer nsqtracer_linux_arm
echo "done"

echo "compiling linux/riscv64..."
GOOS=linux GOARCH=riscv64 go build
mv nsqtracer nsqtracer_linux_riscv64
echo "done"

echo "compiling linux/arm64..."
GOOS=linux GOARCH=arm64 go build
mv nsqtracer nsqtracer_linux_arm64
echo "done."

echo "compiling windows/arm..."
GOOS=windows GOARCH=arm go build
mv nsqtracer.exe nsqtracer_arm.exe
echo "done."

echo "compiling openbsd/386..."
GOOS=openbsd GOARCH=386 go build
mv nsqtracer nsqtracer_openbsd_386
echo "done."

echo "compiling openbsd/amd64..."
GOOS=openbsd GOARCH=amd64 go build
mv nsqtracer nsqtracer_openbsd_amd64
echo "done."

echo "compiling openbsd/arm..."
GOOS=openbsd GOARCH=arm go build
mv nsqtracer nsqtracer_openbsd_arm
echo "done."

echo "compiling openbsd/arm64..."
GOOS=openbsd GOARCH=arm64 go build
mv nsqtracer nsqtracer_openbsd_arm64
echo "done."

echo "compiling darwin/amd64..."
GOOS=darwin GOARCH=amd64 go build
mv nsqtracer nsqtracer_darwin_amd64
echo "done."

# echo "compiling darwin/arm64..."
# GOOS=darwin GOARCH=arm64 go build
# mv nsqtracer nsqtracer_darwin_arm64
# echo "done."

# echo "compiling android/386..."
# GOOS=android GOARCH=386 go build
# mv nsqtracer nsqtracer_android_386
# echo "done."

# echo "compiling android/amd64..."
# GOOS=android GOARCH=amd64 go build
# mv nsqtracer nsqtracer_android_amd64
# echo "done."

# echo "compiling android/arm..."
# GOOS=android GOARCH=arm go build
# mv nsqtracer nsqtracer_android_arm
# echo "done."

echo "compiling android/arm64..."
GOOS=android GOARCH=arm64 go build
mv nsqtracer nsqtracer_android_arm64
echo "done."
