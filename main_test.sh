#!/bin/bash
echo Test all programe of minio
mgo minio bucket ls
mgo minio bucket make apptest
mgo minio object putf apptest test.jpg -p ./test.jpg
mgo minio object ls apptest
mgo minio object stat apptest test.jpg
mgo minio object getf apptest test.jpg -p ./test1.jpg
mgo minio object rm apptest test.jpg
mgo minio object ls apptest
mgo minio bucket rm apptest
mgo minio bucket ls