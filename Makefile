SHELL := /usr/bin/env bash

APP_BIN := .bin/pos-go-api
HTTP_PORT ?= 8081
GO_TEST := GOCACHE=$${GOCACHE:-/tmp/go-build-cache} go test

.DEFAULT_GOAL := help

include make/help.mk
include make/format.mk
include make/test.mk
include make/audit.mk
include make/build.mk
include make/auth.mk
include make/db.mk
include make/git.mk
