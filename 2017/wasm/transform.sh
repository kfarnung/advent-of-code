#!/bin/bash

sed -i -e 's=nodePath\['"'"'normalize'"'"'\](filename)=nodePath['"'"'normalize'"'"'](__dirname + '"'"'/'"'"' + filename)=g' $1
