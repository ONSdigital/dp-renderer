#!/bin/bash -eux

export cwd=$(pwd)

pushd $cwd/dp-renderer
  make audit
popd