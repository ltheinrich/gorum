NAME=gorum
BINARY=${NAME}
INSTALL_DIR=/usr/local/bin

VERSION=v1.0.14
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/ltheinrich/gorum/cmd.Version=${VERSION} -X github.com/ltheinrich/gorum/cmd.BuildTime=${BUILD_TIME}"
GO_FILES=./...

WEB_DIR=web
WEB_DIST=${WEB_DIR}/dist
WEB_MODULES=${WEB_DIR}/node_modules
WEB_LOCK=${WEB_DIR}/yarn.lock

ASSETS_DIR=assets
ASSETS_PKG=assets
ASSETS_FILE=../internal/pkg/assets/assets.go

WEBASSETS_DIR=${WEB_DIST}/gorum
WEBASSETS_PKG=webassets
WEBASSETS_FILE=../../../internal/pkg/webassets/webassets.go

build: clean fmt test buildng assetsng assetsgo buildgo sign

install: gorum
	cp ${BINARY} ${INSTALL_DIR}

run: assetsgo buildgo
	clear
	./${BINARY}

serve:
	clear
	cd ${WEB_DIR} && \
	ng serve

fmt:
	go fmt ${GO_FILES}

test:
	go vet -v ${GO_FILES}
	go test -v -race ${GO_FILES}

buildng:
	cd ${WEB_DIR} && \
	yarn install && \
	ng build --prod

assetsng:
	cd ${WEBASSETS_DIR} && \
	go-bindata -o ${WEBASSETS_FILE} -pkg ${WEBASSETS_PKG} .

assetsgo:
	cd ${ASSETS_DIR} && \
	go-bindata -o ${ASSETS_FILE} -pkg ${ASSETS_PKG} .

buildgo:
	go build ${LDFLAGS} -o ${BINARY}

sign:
	gpg2 -a --detach-sign ${BINARY}

installgo:
	go install ${LDFLAGS}

clean:
	if [ -d ${WEB_DIST} ] ; then rm -rf ${WEB_DIST} ; fi
	if [ -d ${WEB_MODULES} ] ; then rm -rf ${WEB_MODULES} ; fi
	if [ -f ${WEB_LOCK} ] ; then rm -f ${WEB_LOCK} ; fi
	if [ -f ${BINARY} ] ; then rm -f ${BINARY} ; fi
	if [ -f ${BINARY}.asc ] ; then rm -f ${BINARY}.asc ; fi
