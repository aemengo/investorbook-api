#!/usr/bin/env bash

dir=$(cd $(dirname $0) && pwd)

cleanup() {
  rm -f $dir/sampledata.sql
  docker rm -f $container_id
}

# trap cleanup exit

# cd $dir
# unzip sampledata.zip

# container_id=$(docker run --rm -v "$dir/sampledata.sql":/docker-entrypoint-initdb.d/1-init.sql -d -e POSTGRES_HOST_AUTH_METHOD=trust postgres:11)

# echo "Waiting for postgres..."
# sleep 30

go get github.com/onsi/ginkgo/ginkgo
export DATABASE_URI="postgres://postgres:postgres@localhost:5432/postgres"
~/go/bin/ginkgo integration