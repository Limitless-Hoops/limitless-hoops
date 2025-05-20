#!/bin/bash

# Ensure Go-installed tools are available
export PATH="$PATH:$HOME/go/bin"

# Find the Go module root (directory containing go.mod)
find_go_mod_root() {
  local dir="$PWD"
  while [ "$dir" != "/" ]; do
    if [ -f "$dir/go.mod" ]; then
      echo "$dir"
      return
    fi
    dir="$(dirname "$dir")"
  done
  echo "❌ go.mod not found. Are you in a Go project?" >&2
  exit 1
}

MODULE_ROOT="$(find_go_mod_root)"
cd "$MODULE_ROOT" || exit 1

echo "🧠 Running go vet..."
go vet ./...
if [ $? -ne 0 ]; then
  echo -e "\n❌ go vet failed. Fix issues before continuing."
  exit 1
fi

echo "🕵️ Running staticcheck..."
if ! command -v staticcheck &> /dev/null; then
  echo "⚠️ staticcheck not found. Skipping."
else
  staticcheck ./...
  if [ $? -ne 0 ]; then
    echo -e "\n❌ staticcheck failed. Fix issues before continuing."
    exit 1
  fi
fi

echo "🔎 Running tests with coverage..."
COVERAGE_FILE="coverage.out"
go test ./... -coverprofile=$COVERAGE_FILE -json 2>&1 | gotestfmt
EXIT_CODE=$?

# Display coverage summary
if [ -f "$COVERAGE_FILE" ]; then
  echo -e "\n📊 Test Coverage Summary:"
  go tool cover -func=$COVERAGE_FILE | grep total:
  echo
  # Optional: generate HTML report
  go tool cover -html=$COVERAGE_FILE -o coverage.html
  rm $COVERAGE_FILE
fi

if [ $EXIT_CODE -eq 0 ]; then
  echo -e "\n✅ ✅ ✅  ALL TESTS PASSED ✅ ✅ ✅"
else
  echo -e "\n❌ ❌ ❌  SOME TESTS FAILED ❌ ❌ ❌"
fi

exit $EXIT_CODE
