#!/bin/bash
# Expects the script is executed from the project root

# Save the pwd before we run anything
PROJECT_ROOT=`pwd`

# Derive the project name from the directory
PROJECT="$(basename $PROJECT_ROOT)"

# Detremine the output directory
OUT_DIR="$PROJECT_ROOT/bin"

# Run the project
cd $OUT_DIR
./$PROJECT.exe

EXIT_STATUS=$?

if [ $EXIT_STATUS == 0 ]; then
  echo "Run $PROJECT succeeded"
else
  echo "Run $PROJECT failed"
fi

cd $PROJECT_ROOT

exit $EXIT_STATUS