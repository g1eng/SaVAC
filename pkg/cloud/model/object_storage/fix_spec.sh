#!/bin/bash -ex
# This script make temporary fixation for the original spec,
# which lacks pagenation functionality for API key-related operations.

SRC=${1:?no source file specified}
DST=${2:?no destination file specified}


{
  l=1
  while read -r line ; do
  
  case $l in
  1467) echo '
      "PermissionKeyID": {
        "description": "Permission Key ID",
        "type": "string",
        "example": "XYZ123"
      },' 
        echo "$line"
   ;;
  1770|1804) echo "$line" | sed s/PermissionID/PermissionKeyID/ ;;
  *)  echo "$line";;
  esac
  l=$((l+1))
  
  done < $SRC
  echo "}"
} | sed s/\{site_name\}/isk01/g > $DST
