#!/bin/bash

BASEDIR=$(dirname $0)
DOCKERDIR=$BASEDIR/docker
DISTDIR=$DOCKERDIR/dist

rm -rf $DOCKERDIR
mkdir -p $DISTDIR

cp coregroups $DISTDIR
cp coregroups.json $DISTDIR
cp Dockerfile $DOCKERDIR
