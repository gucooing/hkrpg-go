#!/bin/bash

FOLDER_NAME="resources"
ZIP_URL="https://github.alsl.xyz/https://github.com/gucooing/hkrpg-go-Resources"

if [ -d "$FOLDER_NAME" ]; then
    echo "目录 '$FOLDER_NAME' 已存在。"
else
    echo "目录 '$FOLDER_NAME' 不存在，正在下载并解压..."
    git clone --depth=1 $ZIP_URL resources
    echo "资源已下载并解压到 '$FOLDER_NAME'。"
fi

echo "RUN hkrpg-go"
./hkrpg-go