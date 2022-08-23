#!/usr/bin/env bash

current_path=$(pwd)
bash "${current_path}"/stop.sh
echo "Apollo stopped"

bash "${current_path}"/start.sh
echo "Apollo started"

echo "done"