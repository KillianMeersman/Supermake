#!/bin/bash

if [ $# -lt 1 ]; then
    echo "Please provide a semantic version, 'major', 'minor' or 'patch'"
    exit 1
fi


NEW_VERSION="$1"

git fetch --tags
LATEST_RELEASE_TAG=$(git tag | grep -iE 'v.*' | sort -V | tail -n 1)
NEW_RELEASE_MAJOR=$(echo ${LATEST_RELEASE_TAG//v} | cut -d. -f 1)
NEW_RELEASE_MINOR=$(echo ${LATEST_RELEASE_TAG//v} | cut -d. -f 2)
NEW_RELEASE_PATCH=$(echo ${LATEST_RELEASE_TAG//v} | cut -d. -f 3)

if [[ "${NEW_VERSION}" == 'major' ]]; then
    NEW_RELEASE_MAJOR=$(($NEW_RELEASE_MAJOR + 1))
    NEW_RELEASE_MINOR='0'
    NEW_RELEASE_PATCH='0'
elif [[ "${NEW_VERSION}" == 'minor' ]]; then
    NEW_RELEASE_MINOR=$(($NEW_RELEASE_MINOR + 1))
    NEW_RELEASE_PATCH='0'
elif [[ "${NEW_VERSION}" == 'patch' ]]; then
    NEW_RELEASE_PATCH=$(($NEW_RELEASE_PATCH + 1))
elif [[ "${NEW_VERSION}" =~ v.*\..*\..* ]]; then
    NEW_RELEASE_MAJOR=$(echo ${NEW_VERSION//v} | cut -d. -f 1)
    NEW_RELEASE_MINOR=$(echo ${NEW_VERSION//v} | cut -d. -f 2)
    NEW_RELEASE_PATCH=$(echo ${NEW_VERSION//v} | cut -d. -f 3)
else
    echo "Invalid argument. Please provide a semantic version, 'major', 'minor' or 'patch'"
    exit 1
fi

echo "Latest release tag is '${LATEST_RELEASE_TAG}'"
NEW_RELEASE="v${NEW_RELEASE_MAJOR}.${NEW_RELEASE_MINOR}.${NEW_RELEASE_PATCH}"
echo "Creating release '${NEW_RELEASE}'"

TAG_MESSAGE_FILE=$(mktemp)
echo "Version ${NEW_RELEASE_MAJOR}.${NEW_RELEASE_MINOR}.${NEW_RELEASE_PATCH}" > ${TAG_MESSAGE_FILE}
${EDITOR} ${TAG_MESSAGE_FILE}

git tag -a ${NEW_RELEASE} -F ${TAG_MESSAGE_FILE}

while true; do
    read -p "Release tag created, push, delete or keep? [p/d/k]: " choice
    case $choice in
        [Pp]*) git push origin ${NEW_RELEASE} ; break ;;
        [Dd]*) git tag -d ${NEW_RELEASE} ; break ;;
        [Kk]*) break ;;
    esac
done
