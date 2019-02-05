CC ?= gcc
CFLAGS ?= -g -Wall

SRC_DIR = src
SRC = $(wildcard ${SRC_DIR}/*)

BUILD_DIR = build
BUILD_FILE = ${BUILD_DIR}/nmv

build: ${SRC}
	mkdir -p ${BUILD_DIR}
	${CC} ${CFLAGS} ${SRC} -o ${BUILD_FILE}
