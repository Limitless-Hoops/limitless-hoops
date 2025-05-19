#!/bin/bash

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

echo "🔎 Running tests..."
# Capture output and exit code
TEST_OUTPUT=$(go test ./... -json 2>&1 | gotestfmt)
EXIT_CODE=$?

# Print formatted output
echo "$TEST_OUTPUT"

# Show summary
if [ $EXIT_CODE -eq 0 ]; then
  echo -e "\n✅ ✅ ✅  ALL TESTS PASSED ✅ ✅ ✅"
else
  echo -e "\n❌ ❌ ❌  SOME TESTS FAILED ❌ ❌ ❌"
fi

exit $EXIT_CODE
