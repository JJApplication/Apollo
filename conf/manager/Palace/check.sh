#!/usr/bin/env bash

cd "${APP_ROOT}/${APP}" || exit "${APP_STATUS_ERR}"

"${APP_ROOT}"/NoEngine/noengined -stat "${APP}"
result=$?
if [[ $result != 0 ]];then
  exit "${APP_STATUS_ERR}"
else
  exit 0
fi

exit 0