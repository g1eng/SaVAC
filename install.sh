#!/bin/sh -ex

throw(){ echo "$*" >&2 ; exit 1; }

OS=`uname -a | awk '{print $1}' | tr A-Z  a-z`

case $OS in
darwin|linux|freebsd) :;;
*) throw "$OS is not supported for this install script. Install it manually." ;;
esac

if $OS = "linux"; then
	ARCH=`uname -a | awk '{print $(NF-1)}'`
else
	ARCH=`uname -a | awk '{print $NF}'`
fi

case $ARCH in
amd64|x86_64|arm64) :;;
*) throw "$ARCH is not supported for this install script. Build it manually.";;
esac

[ $ARCH = x86_64 ] && ARCH=amd64

curl -fSL -o savac.zip https://github.com/g1eng/savac-old-mainline/releases/download/v0.4.1/savac-${OS}-${ARCH}.zip
unzip -d savac
test -d $PREFIX/bin || {
  mkdir -v -p $PREFIX/bin
  chmod u+x savac
  mv -v savac $PREFIX/bin
}
