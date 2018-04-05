#!/bin/sh

go clean
git add -A
git commit -m "$1"
git push
