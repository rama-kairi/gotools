package gh

const preCommitHook = `#!/bin/sh

# Creating a variable of all the files that have been changed
STAGED_GO_FILES=$(git diff --cached --name-only | grep ".go$")

# Checking if there are any staged files
if [[ "$STAGED_GO_FILES" = "" ]]; then
  exit 0
fi

PASS=true

# Looping through all the staged files
for FILE in $STAGED_GO_FILES
do
    golangci-lint run --tests=0 ./...
    if [ $? -ne 0 ]; then
        PASS=false
    fi
done

if ! $PASS; then
  printf "❌ Commit Failed with the above error, Please fix and Retry\n"
  exit 1
fi

exit 0
`

const prepareCommitMsgHook = `#!/bin/sh

# set -- $GIT_PARAMS

BRANCH_NAME=$(git symbolic-ref --short HEAD)

BRANCH_IN_COMMIT=0
if [ -f $1 ]; then
    BRANCH_IN_COMMIT=$(grep -c "\[$BRANCH_NAME\]" $1)
fi

if [ -n "$BRANCH_NAME" ] && ! [[ $BRANCH_IN_COMMIT -ge 1 ]]; then
  if [ -f $1 ]; then
    BRANCH_NAME="${BRANCH_NAME/\//\/}"
    sed -i.bak -e "1s@^@[$BRANCH_NAME] @" $1
  else
    echo "[$BRANCH_NAME] " > "$1"
  fi
fi

exit 0
`

const prePushHook = `#!/bin/sh
# Create dist folder if not exists
mkdir -p dist

# build the go binary
ENV_GOOS=linux go build -o dist/main
`

const postCommitHook = `#!/bin/sh
# Commit Success Message
echo "✅ -> Commit Succeeded...\n"

# Put Your Communication Code here

`
