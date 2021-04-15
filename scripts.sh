
function startdb {
  id=$(_find_instance_id)
  if [ -z "$id" ]
  then
    docker run --rm -d -p 8529:8529 -e ARANGO_ROOT_PASSWORD=abc123 \
      --name arango arangodb
  fi
}

function killdb {
  id=$(_find_instance_id)
  if [ -n "$id" ]
  then
    docker kill "$id"
  fi
}

function _find_instance_id {
  id=$(docker ps | grep arango | awk '{ print $1 }')
  echo "$id"
}

function startserver {
  make build-grpc-server
  ./dist/grpc-server
}
