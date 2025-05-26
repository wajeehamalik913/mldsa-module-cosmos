protoVer=0.16.0
protoImageName=ghcr.io/cosmos/proto-builder:$protoVer
docker run --rm -v $PWD:/workspace --workdir /workspace $protoImageName sh ./scripts/protocgen.sh

