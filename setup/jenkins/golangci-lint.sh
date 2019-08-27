VERSION=1.15.0

if [[ -z "${VERSION}" ]]; then
    echo Version is required
    exit 1
fi

workPath=$(cd `dirname $0`; pwd)
GOLANGCI_LINT_PATH=$workPath/golangci-lint/bin

GOLANGCI_LINT_BIN=$GOLANGCI_LINT_PATH/golangci-lint
if [ -f $GOLANGCI_LINT_BIN ];then
    PATH=$PATH:$GOLANGCI_LINT_PATH
    INSTALLED_VERSION=`golangci-lint --version 2> /dev/null`
    if [[ "$INSTALLED_VERSION" == *"$VERSION"* ]]; then
        echo ${INSTALLED_VERSION}
        exit 0
    fi
fi

curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOLANGCI_LINT_PATH v${VERSION}
