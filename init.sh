#!/bin/bash
set -eux

sudo mkdir -p $1/data/{lib,log}

# 注意: mac电脑中, chmod命令, 必须先写-R再写777
sudo chmod -R 777 $1/data


