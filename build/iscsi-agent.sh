#! /bin/bash -x

echo "Starting iscsid..."
iscsid -f &

# sleep one second to make sure iscsid is running in background
sleep 1

echo "Calling $@..."
exec "$@"
