#!/bin/sh

# for Go
export GOOS=linux
export GOARCH=arm
export GOARM=7

# variables
FOLDER=KoboFileServer
APP=kobofileserver
SCRIPT=kobofileserver.sh

ZIP=/c/Program\ Files/7-Zip/7z.exe
RELEASE=Release.zip

# clear old folder
rm -f "$RELEASE"
rm -rf "$FOLDER"
mkdir "$FOLDER"

# build app
cd ..
go build -o "$APP" .

if [ "$?" != "0" ]; then
    echo "Please check code."
    exit 1
fi

# copy files
cd scripts

cp ../"$APP" "$FOLDER"
cp -r ../web "$FOLDER"

cp "$SCRIPT" "$FOLDER"
sed -i "s/\r//g" "$FOLDER"/"$SCRIPT"

# create zip
"$ZIP" a "$RELEASE" "$FOLDER"

if [ "$?" != "0" ]; then
    echo "Please zip folder manually."
    exit 1
fi

echo "Done"