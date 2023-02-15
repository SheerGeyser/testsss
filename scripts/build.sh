#!/bin/sh

DATE=$(date +%FT%T%z)
VERSION=$(git describe --abbrev=0 --tags)
CURRENT_BRANCH=$(git symbolic-ref --short -q HEAD)

if [ -z "$APP" ]; then
  APP=secretary-bot
fi

if [ -z "$PACKAGE" ]; then
  PACKAGE=gitlab.sima-land.ru/dev-dep/secretary-bot
fi

if [ -n "$CI_COMMIT_SHA" ]; then
  CURRENT_BRANCH=$CI_COMMIT_SHA
fi

if [ -z "$RELEASE" ]; then
  RELEASE=dev
fi

if [ -z "$VERSION" ]; then
  VERSION=0.0.0
fi

SEMVER='gitlab.sima-land.ru/pac/semver'
LDFLAGS="-w -s -X $SEMVER.Date=$DATE -X $SEMVER.Version=$VERSION -X $SEMVER.Env=$RELEASE -X $SEMVER.Branch=$CURRENT_BRANCH"

if [ "$RELEASE" = "dev" ]; then
  echo "$APP", "$PACKAGE"
  cd src && go build -ldflags "$LDFLAGS" -v -o ../bin/"$APP" .
else
  cd src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${LDFLAGS}" -v -o ../bin/"$APP" .
fi
