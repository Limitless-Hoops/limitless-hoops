#!/bin/bash

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
