NAME=gorum
BINARY=${NAME}
INSTALL_DIR=/usr/local/bin

NUM_VERSION=1.0.16
VERSION=v${NUM_VERSION}
BUILD_TIME=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/NathanNr/gorum/cmd.Version=${VERSION} -X github.com/NathanNr/gorum/cmd.BuildTime=${BUILD_TIME}"
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
build: clean fmt test buildng assetsng assetsgo buildgo sign deb

.PHONY: install
install: ${BINARY}
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
	strip ${BINARY}

.PHONY: sign
sign:
	gpg2 -a --detach-sign ${BINARY}

.PHONY: deb
deb:
	cp ${BINARY} ${NAME}_${NUM_VERSION}/usr/bin/gorum
	dpkg -b ${NAME}_${NUM_VERSION}

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
	if [ -f ${NAME}_${NUM_VERSION}.deb ] ; then rm -f ${NAME}_${NUM_VERSION}.deb ; fi
