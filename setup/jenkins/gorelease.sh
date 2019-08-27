

VERSION=0.106.0

if [[ -z "${VERSION}" ]]; then
    echo Version is required
    exit 1
fi

workPath=$(cd `dirname $0`; pwd)
GORELEASER_PATH=$workPath/goreleaser/bin

GORELEASER_BIN=$GORELEASER_PATH/goreleaser
if [ -f $GORELEASER_BIN ];then
    PATH=$PATH:$GORELEASER_PATH
    INSTALLED_VERSION=`goreleaser --version 2>&1`
    if [[ "$INSTALLED_VERSION" == *"$VERSION"* ]]; then
        echo ${INSTALLED_VERSION}
        exit 0
    fi
fi



curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh -s -- -b $GORELEASER_PATH v${VERSION}
