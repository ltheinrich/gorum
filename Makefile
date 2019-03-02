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

.PHONY: build
build: clean fmt test buildng assetsng assetsgo buildgo sign

.PHONY: install
install: gorum
	cp ${BINARY} ${INSTALL_DIR}

.PHONY: run
run: assetsgo buildgo
	clear
	./${BINARY}

.PHONY: serve
serve:
	clear
	cd ${WEB_DIR} && \
	ng serve

.PHONY: fmt
fmt:
	go fmt ${GO_FILES}

.PHONY: test
test:
	go vet -v ${GO_FILES}
	go test -v -race ${GO_FILES}

.PHONY: buildng
buildng:
	cd ${WEB_DIR} && \
	yarn install && \
	ng build --prod

.PHONY: assetsng
assetsng:
	cd ${WEBASSETS_DIR} && \
	go-bindata -o ${WEBASSETS_FILE} -pkg ${WEBASSETS_PKG} .

.PHONY: assetsgo
assetsgo:
	cd ${ASSETS_DIR} && \
	go-bindata -o ${ASSETS_FILE} -pkg ${ASSETS_PKG} .

.PHONY: buildgo
buildgo:
	go build ${LDFLAGS} -o ${BINARY}

.PHONY: sign
sign:
	gpg2 -a --detach-sign ${BINARY}

.PHONY: installgo
installgo:
	go install ${LDFLAGS}

.PHONY: clean
clean:
	if [ -d ${WEB_DIST} ] ; then rm -rf ${WEB_DIST} ; fi
	if [ -d ${WEB_MODULES} ] ; then rm -rf ${WEB_MODULES} ; fi
	if [ -f ${WEB_LOCK} ] ; then rm -f ${WEB_LOCK} ; fi
	if [ -f ${BINARY} ] ; then rm -f ${BINARY} ; fi
	if [ -f ${BINARY}.asc ] ; then rm -f ${BINARY}.asc ; fi
