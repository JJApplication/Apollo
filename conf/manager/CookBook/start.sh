#!/usr/bin/env bash

cd "${APP_ROOT}/${APP}" || exit "${APP_START_ERR}"

"${APP_ROOT}"/NoEngine/noengined -start "${APP}"
result=$?
docker restart NoEngine > /dev/null
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi