#!/usr/bin/env bash

# How to run manually:
# docker build --pull --rm -f "contrib/devtools/Dockerfile" -t cosmossdk-proto:latest "contrib/devtools"
# docker run --rm -v $(pwd):/workspace --workdir /workspace cosmossdk-proto sh ./scripts/protocgen.sh

echo "Formatting protobuf files"
find ./ -name "*.proto" -exec clang-format -i {} \;

set -e

home=$PWD

echo "Generating proto code"
proto_dirs=$(find ./ -name 'buf.yaml' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  echo "Generating proto code for $dir"

  cd $dir
  # check if buf.gen.gogo.yaml exists in the proto directory
  if [ -f "buf.gen.gogo.yaml" ]; then
      for file in $(find . -maxdepth 5 -name '*.proto'); do
        echo "$file"
        # this regex checks if a proto file has its go_package set to cosmossdk.io/api/...
        # gogo proto files SHOULD ONLY be generated if this is false
        # we don't want gogo proto to run for proto files which are natively built for google.golang.org/protobuf
        if grep -q "option go_package" "$file" && grep -H -o -c 'option go_package.*cosmossdk.io/api' "$file" | grep -q ':0$'; then
          buf generate --template buf.gen.gogo.yaml $file
        fi
    done
  fi

  cd $home
done

# move generated files to the right places
cp -r github.com/forbole/juno/v5/cosmos-sdk ./
cp -r github.com/forbole/juno/v5/ibc-go ./
rm -rf github.com

# Comment out all the proto registrations.
# We assume that the codec provided by the client has registered the types.
find . -name "*.pb.go" | while read -r file; do
    sed -i '/proto\.Register/s/^/\/\//' "$file"
done

go mod tidy
