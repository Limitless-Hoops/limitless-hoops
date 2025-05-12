#!/bin/sh

HOSTS_FILE="/etc/hosts"
ENTRY="127.0.0.1 limitlesshoops.dev www.limitlesshoops.dev"

# Check if entry already exists
if grep -q "limitlesshoops.dev" "$HOSTS_FILE"; then
    echo "Entry already exists in $HOSTS_FILE"
else
    echo "Adding entry to $HOSTS_FILE"
    echo "$ENTRY" | sudo tee -a "$HOSTS_FILE" > /dev/null
    echo "Entry added successfully."
fi
