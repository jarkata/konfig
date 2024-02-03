#!/bin/bash

##Gitee

msg=$1
if [ -z "$msg" ];then
    echo "msg not emtpy"
    exit 0;
fi

git add .
git commit -m "$msg"
git push origin master