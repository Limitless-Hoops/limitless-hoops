#!/bin/bash

MODULE_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

echo "🧠 Running go vet..."
(cd "$MODULE_ROOT" && go vet ./...)
if [ $? -ne 0 ]; then
  echo -e "\n❌ go vet failed. Fix issues before continuing."
  exit 1
fi

echo "🕵️ Running staticcheck..."
if ! command -v staticcheck &> /dev/null; then
  echo "⚠️ staticcheck not found. Skipping."
else
  (cd "$MODULE_ROOT" && staticcheck ./...)
  if [ $? -ne 0 ]; then
    echo -e "\n❌ staticcheck failed. Fix issues before continuing."
    exit 1
  fi
fi

echo "🔎 Running tests..."

# Run tests from the tests dir
TEST_OUTPUT=$(go test ./... -json 2>&1 | gotestfmt)
EXIT_CODE=$?

echo "$TEST_OUTPUT"

if [ $EXIT_CODE -eq 0 ]; then
  echo -e "\n✅ ✅ ✅  ALL TESTS PASSED ✅ ✅ ✅"
else
  echo -e "\n❌ ❌ ❌  SOME TESTS FAILED ❌ ❌ ❌"
fi

exit $EXIT_CODE
