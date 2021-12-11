#!/usr/bin/env bash
# 打包脚本

package()
{
  echo "start to package: $(date)"
  if [[ ! -f ./dirichlet ]];then
    echo "binary not exist"
    exit 1
  fi

  # clear
  if [[ -d ./opt ]];then
    rm -rf ./opt
  fi

  mkdir ./opt

  # generate
  cp dirichlet ./opt
  cp -r ./conf ./opt
  tar -czvf dirichlet.tar.gz ./opt
}

clean()
{
   rm -rf ./opt
}

package
clean