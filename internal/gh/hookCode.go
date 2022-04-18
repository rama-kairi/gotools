package gh

const preCommitHook = `#!/bin/sh
# Running golangci-lint
printf "Running golangci-lint\n"
golangci-lint run --tests=0 ./...

# Running golangci-lint
printf "Running go mod tidy\n"
go mod tidy
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
# Pushing to main branch not allowed
if [ "$1" = "main" ]; then
  printf "❌ Pushing to main branch not allowed\n"
  exit 1
fi

exit 0
`

const postCommitHook = `#!/bin/sh
# Commit Success Message
echo "✅ -> Commit Succeeded...\n"

# Put Your Communication Code here

`
