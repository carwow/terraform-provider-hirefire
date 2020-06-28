#!/bin/sh

set -o nounset
set -o errexit
set -o pipefail

# uname os and arch mappings:
# https://github.com/goreleaser/godownloader/blob/e64d0375716bf060e73ee7248cfd48c49a0b1173/shellfn.go#L62-L84

uname_os() {
  uname -s | tr '[:upper:]' '[:lower:]'
}

uname_arch() {
  arch=$(uname -m)
  case $arch in
    x86_64) arch="amd64" ;;
    x86)    arch="386"   ;;
    i686)   arch="386"   ;;
    i386)   arch="386"   ;;
  esac
  echo ${arch}
}


if [ -z ${1+x} ]; then
	echo "Requires version."
	echo "Usage: $0 [VERSION]"
	exit 1
fi

VERSION=$1
OS=linux #$(uname_os)
ARCH=amd64 #$(uname_arch)


cd ~/.terraform.d/plugins

test -e terraform-provider-hirefire_v${VERSION} && exit || true

wget -O terraform-provider-hirefire_v${VERSION} \
  https://github.com/carwow/terraform-provider-hirefire/releases/download/v${VERSION}/terraform-provider-hirefire_v${VERSION}_${OS}_${ARCH}

chmod +x terraform-provider-hirefire_v${VERSION}
