#!/bin/bash
set -euo pipefail

CLIENT=$1
VERSION=$2

REPO="https://github.com/cybercryptio/d1-service-${CLIENT}.git"
TARGET="d1-${CLIENT}"

CURRENT_DIR=$(pwd)
CLIENT_DIR=$(realpath $TARGET)
CLIENT_PROTOBUF_DIR=$CLIENT_DIR/protobuf

# prepare target directories
rm -rf $CLIENT_DIR
mkdir $CLIENT_DIR
mkdir $CLIENT_PROTOBUF_DIR

# checkout service repo
rm -rf checkout
#git clone --branch $VERSION --depth=1 $REPO checkout 2> /dev/null
git clone --config advice.detachedHead=false --branch $VERSION --depth=1 $REPO checkout
#git clone --branch $VERSION --depth=1 $REPO checkout
cd checkout
SRC_DIR=$(pwd)

# copy client source files
cd $SRC_DIR/client
find . -name \*.go -exec cp --parents {} $CLIENT_DIR \;

# copy protobuf source files
cd $SRC_DIR/protobuf
find . -name \*.go -exec cp --parents {} $CLIENT_PROTOBUF_DIR \;

# clean up temporary checkout directory
cd $CURRENT_DIR
rm -rf checkout

# perform text replacements
REPLACEMENTS=(
    '"github.com/cybercryptio/d1-service-generic/client"=>"github.com/cybercryptio/d1-client-go/d1-generic"'
    '"github.com/cybercryptio/d1-service-generic/protobuf"=>"github.com/cybercryptio/d1-client-go/d1-generic/protobuf"'
    '"github.com/cybercryptio/d1-service-storage/client"=>"github.com/cybercryptio/d1-client-go/d1-storage"'
    '"github.com/cybercryptio/d1-service-storage/protobuf"=>"github.com/cybercryptio/d1-client-go/d1-storage/protobuf"'
)

for REPLACEMENT in "${REPLACEMENTS[@]}"; do
    OLD_VALUE=${REPLACEMENT%=>*} # drops substring from last occurrence of `=>` to end of string
    NEW_VALUE=${REPLACEMENT#*=>} # drops substring from start of string up to first occurrence of `=>`

    # NOTE: "|| true" is to suppress non-zero exit codes from grep when nothing is found
    FILES=$(grep --recursive --files-with-matches $OLD_VALUE $CLIENT_DIR || true)
    for FILE in $FILES; do
        sed -i "s|${OLD_VALUE}|${NEW_VALUE}|g" $FILE
    done
done

go mod tidy

GREEN='\033[0;32m'
NC='\033[0m' # No Color
printf "${GREEN}Client '${TARGET}' is now running version '${VERSION}'${NC}\n"
