#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-renderer
  make lint
popd