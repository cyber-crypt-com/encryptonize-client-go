#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

SCRIPT=$(basename $0)
CLIENT=$1
VERSION=$2

CURRENT_DIR=$(pwd)
CHECKOUT_DIR=$(mktemp -d)

finish() {
    # clean up temporary checkout directory
    rm -rf $CHECKOUT_DIR
    # return to previous working directory
    cd $CURRENT_DIR
}
trap finish EXIT

REPO="https://github.com/cybercryptio/d1-service-${CLIENT}.git"
TARGET="d1-${CLIENT}"

CLIENT_DIR=$(realpath $TARGET)
CLIENT_PROTOBUF_DIR=$CLIENT_DIR/protobuf

# prepare target directories
rm -rf $CLIENT_DIR
mkdir $CLIENT_DIR
mkdir $CLIENT_PROTOBUF_DIR

# checkout service repo
rm -rf $CHECKOUT_DIR
git clone $REPO $CHECKOUT_DIR
cd $CHECKOUT_DIR
git fetch --quiet --tags
git checkout --quiet "$VERSION"
COMMIT_ID=$(git rev-parse HEAD)
SRC_DIR=$(pwd)

# remove_copyright_comments reads text from STDIN,
# removes all copyright comments and returns it on STDOUT
remove_copyright_comments() {
    local MODE_KEEP='keep'
    local MODE_SKIP='skip'
    local MODE=$MODE_KEEP
    local IFS=''

    while read LINE; do
        if [[ $MODE == $MODE_KEEP ]]; then
            if [[ "$LINE" =~ '// Copyright'.*'CYBERCRYPT' ]]; then
                MODE=$MODE_SKIP
            else
                echo "$LINE"
            fi
        elif [[ "$LINE" != "" ]]; then
            MODE=$MODE_KEEP
            echo "$LINE"
        fi
    done
}

# replace reads text from STDIN, replaces all occurrences
# of $1 with $2, and returns it on STDOUT
replace() {
    OLD_VALUE=$1
    NEW_VALUE=$2
    sed -e "s|${OLD_VALUE}|${NEW_VALUE}|g"
}

# process_source_file takes a go source file (specified as a path through $1)
# applies the following adjustments and returns it on STDOUT:
# - Prepends an Apache license header
# - Prepends a DO NOT EDIT reminder with source information
# - Removes existing copyright comments
# - Replaces various Go imports to make it work from the new location
process_source_file() {
    FILE=$1

    # output an Apache license header and a "DO NOT EDIT" reminder.
    cat << EOF
// Copyright 2022 CYBERCRYPT
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by ${SCRIPT}. DO NOT EDIT.
// version: ${VERSION}
// source: ${REPO}
// commit: ${COMMIT_ID}

EOF

    cat $FILE \
    | remove_copyright_comments \
    | replace \
        'd1-service-generic/client' \
        'd1-client-go/d1-generic' \
    | replace \
        'd1-service-generic/protobuf' \
        'd1-client-go/d1-generic/protobuf' \
    | replace \
        'd1-service-storage/client' \
        'd1-client-go/d1-storage' \
    | replace \
        'd1-service-storage/protobuf' \
        'd1-client-go/d1-storage/protobuf'
}

# copy and process client source files
cd $SRC_DIR/client
GO_FILES=$(find . -name \*.go)
for GO_FILE in $GO_FILES; do
    SRC_PATH=$(realpath $GO_FILE)
    DST_PATH=$(realpath $CLIENT_DIR/$GO_FILE)
    process_source_file $SRC_PATH > $DST_PATH
done

# copy protobuf source files
cd $SRC_DIR/protobuf
GO_FILES=$(find . -name \*.go)
for GO_FILE in $GO_FILES; do
    SRC_PATH=$(realpath $GO_FILE)
    DST_PATH=$(realpath $CLIENT_PROTOBUF_DIR/$GO_FILE)
    cp $SRC_PATH $DST_PATH
done

cd $CURRENT_DIR
go mod tidy

COLOR_GREEN='\033[0;32m'
COLOR_NONE='\033[0m'
printf "${COLOR_GREEN}Client '${TARGET}' is now running version '${VERSION} (${COMMIT_ID})'${COLOR_NONE}\n"
