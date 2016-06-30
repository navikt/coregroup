#!/bin/bash

BASEDIR=$(dirname $0)
DOCKERDIR=$BASEDIR/docker
DISTDIR=$DOCKERDIR/dist

rm -rf $DOCKERDIR
mkdir -p $DISTDIR

cp -r dist $DOCKERDIR
cp Dockerfile $DOCKERDIR
