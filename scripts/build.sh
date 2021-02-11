#!/bin/bash
# Expects the script is executed from the project root
# Ref: https://gist.github.com/nate/2048566#file-build-sh-L47

# Save the pwd before we run anything
PROJECT_ROOT=`pwd`

# Determine the build script's actual directory
BUILD_DIR="./cmd/vfs"

# Derive the project name from the directory
PROJECT="$(basename $PROJECT_ROOT)"

# Detremine the output directory
OUT_DIR="$PROJECT_ROOT/bin"

# Build the project
cd $BUILD_DIR
mkdir -p $OUT_DIR
go build -o $OUT_DIR/$PROJECT.exe main.go

EXIT_STATUS=$?

if [ $EXIT_STATUS == 0 ]; then
  echo "Build succeeded"
else
  echo "Build failed"
fi

cd $PROJECT_ROOT

exit $EXIT_STATUS