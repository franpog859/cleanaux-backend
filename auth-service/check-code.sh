#!/usr/bin/env bash

echo "Preparing shell..."
set -o errexit ; set -o nounset
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${CURRENT_DIR}

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

##
# DEP ENSURE
##
echo "Checking dep ensure..."
dep ensure -v --vendor-only
ensureResult=$?
if [ ${ensureResult} != 0 ]; then
	echo -e "${RED}✗ dep ensure -v --vendor-only ${NC}\n$ensureResult${NC}"
	exit 1
else echo -e "${GREEN}√ dep ensure -v --vendor-only ${NC}"
fi

##
# DEP STATUS
##
echo "Checking dep status..."
depResult=$(dep status -v)
if [ $? != 0 ]; then
	echo -e "${RED}✗ dep status -v\n $depResult${NC}"
	exit 1
else echo -e "${GREEN}√ dep status -v ${NC}"
fi

##
# GO TEST
##
echo "Checking go test..."
go test -race ./...
if [ $? != 0 ]; then
	echo -e "${RED}✗ go test -race ./...\n ${NC}"
	exit 1
else echo -e "${GREEN}√ go test -race ./... ${NC}"
fi

##
#  GO LINT
##
echo "Checking go build golint..."
go build -o golint-vendored ./vendor/github.com/golang/lint/golint
buildLintResult=$?
if [ ${buildLintResult} != 0 ]; then
	echo -e "${RED}✗ go build lint$ {NC}\n$buildLintResult${NC}"
	exit 1
fi

echo "Checking golint..."
filesToCheck=$(find . -type f -name "*.go" | egrep -v "\/vendor\/|_*/automock/|_*/testdata/|/pkg\/|_*export_test.go")
golintResult=$(echo "${filesToCheck}" | xargs -L1 ./golint-vendored)
rm golint-vendored

if [ $(echo ${#golintResult}) != 0 ]; then
	echo -e "${RED}✗ golint \n$golintResult${NC}"
	exit 1
else echo -e "${GREEN}√ golint ${NC}"
fi

##
# GO FMT
##
echo "Checking go fmt..."
filesToCheck=$(find . -type f -name "*.go" | egrep -v "\/vendor\/|_*/automock/|_*/testdata/|/pkg\/|_*export_test.go")
goFmtResult=$(echo "${filesToCheck}" | xargs -L1 go fmt)
if [ $(echo ${#goFmtResult}) != 0 ]
	then
    	echo -e "${RED}✗ go fmt ${NC}\n$goFmtResult${NC}"
    	exit 1;
	else echo -e "${GREEN}√ go fmt ${NC}"
fi

##
# GO IMPORTS
##
echo "Checking go build goimports..."
go build -o goimports-vendored ./vendor/golang.org/x/tools/cmd/goimports
buildGoImportResult=$?
if [ ${buildGoImportResult} != 0 ]; then
	echo -e "${RED}✗ go build goimports ${NC}\n$buildGoImportResult${NC}"
	exit 1
fi

echo "Checking goimports..."
filesToCheck=$(find . -type f -name "*.go" | egrep -v "\/vendor\/|_*/automock/|_*/testdata/|/pkg\/|_*export_test.go")
goImportsResult=$(echo "${filesToCheck}" | xargs -L1 ./goimports-vendored -w -l)
rm goimports-vendored

if [ $(echo ${#goImportsResult}) != 0 ]; then
	echo -e "${RED}✗ goimports ${NC}\n$goImportsResult${NC}"
	exit 1
else echo -e "${GREEN}√ goimports ${NC}"
fi

##
# GO VET
##
echo "Checking go vet..."
packagesToVet=("./cmd/...")
for vPackage in "${packagesToVet[@]}"; do
	vetResult=$(go vet ${vPackage})
	if [ $(echo ${#vetResult}) != 0 ]; then
		echo -e "${RED}✗ go vet ${vPackage} ${NC}\n$vetResult${NC}"
		exit 1
	else echo -e "${GREEN}√ go vet ${vPackage} ${NC}"
	fi
done

##
# GO BUILD
##
echo "Checking go build..."
go build -a -installsuffix cgo -o built-service ./cmd/authservice
goBuildResult=$?
if [ ${goBuildResult} != 0 ]; then
	echo -e "${RED}✗ go build ${NC}\n$goBuildResult${NC}"
	exit 1
else echo -e "${GREEN}√ go build ${NC}"
fi
rm built-service