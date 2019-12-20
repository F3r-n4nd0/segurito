#!/bin/bash
set -e

cmd="$1"

until timeout 1 bash -c "cat < /dev/null > /dev/tcp/${RABBIT_MQ_HOST}/${RABBIT_MQ_PORT}"; do
  >&2 echo "Rabbit MQ not up yet on ${RABBIT_MQ_HOST}"
  sleep 1
done

echo "Rabbit MQ is up"
exec $cmd