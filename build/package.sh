#!/usr/bin/env bash
# 打包脚本

package()
{
  echo "start to package: $(date)"
  if [[ ! -f ./apollo ]];then
    echo "binary not exist"
    exit 1
  fi

  # clear
  if [[ -d ./opt ]];then
    rm -rf ./opt
  fi

  mkdir ./opt

  # generate
  cp ./apollo ./opt
  cp -r ./conf ./opt
  cp -r ./web ./opt
  tar -czvf dirichlet.tar.gz ./opt
}

clean()
{
   rm -rf ./opt
}

package
clean