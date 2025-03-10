#! /bin/bash

TARGET_DIR="./connect"
FILES=("go.work" "go.work.sum")
SOURCE_DIR="./configs"

# Check if the file exists
for FILE_NAME in "${FILES[@]}"; do
    if [ -f "$TARGET_DIR/$FILE_NAME" ]; then
        echo "$FILE_NAME ファイルが存在します。削除します。"
        mv "$TARGET_DIR/$FILE_NAME" "$SOURCE_DIR/$FILE_NAME"
    elif [ -f "$SOURCE_DIR/$FILE_NAME" ]; then
        echo "$FILE_NAME が存在しません。コピーします。"
        mv "$SOURCE_DIR/$FILE_NAME" "$TARGET_DIR/$FILE_NAME"
    else
        echo "$TARGET_DIR/$FILE_NAME が存在しません。"
    fi
done
