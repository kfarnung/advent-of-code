#!/bin/bash

cd "$(cd "${0%/*}" && pwd -P)";
emcc day15.c -s WASM=1 -O3 --js-transform './transform.sh' -o day15.js
