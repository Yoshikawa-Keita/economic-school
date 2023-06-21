#!/bin/bash

echo "start batch"

# Call the ranking creation script
echo "Creating rankings..."
./create_user_ranking.sh

# Example of a process to delete unused users
echo "Deleting unused users..."
# TODO

echo "Batch processes completed."
