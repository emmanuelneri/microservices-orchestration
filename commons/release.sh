#!/usr/bin/env bash

currentVersion=`git describe --abbrev=0 --tags 2>/dev/null`
echo "Current version: ${currentVersion}"

IFS='.' read -r -a array <<< "${currentVersion}"
minorRelease="$((${array[2]} + 1))"
newVersion="${array[0]}.${array[1]}.${minorRelease}"
echo "new version: ${newVersion}"

git tag ${newVersion}
git push --tags
