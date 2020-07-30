#!/usr/bin/env bash
echo 'Running migrations...'
/go-cloud-native/migrate up > /dev/null 2>&1 &

echo 'Deleting mysql-client...'
apk del mysql-client

echo 'Start application...'
/go-cloud-native/app
