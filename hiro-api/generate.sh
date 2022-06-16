# For macOS

export HOMEBREW_NO_ENV_HINTS=1

# check if openapi-generator is installed
if ! [ -x "$(command -v openapi-generator)" ]; then
  echo 'Error: openapi-generator is not installed.' >&2
  echo 'Attempting to install openapi-generator with Homebrew...' >&2
  brew install openapi-generator
  if ! [ -x "$(command -v openapi-generator)" ]; then
    echo 'Error: openapi-generator is not available.' >&2
    exit 1
  fi
fi

# check if swagger-cli is installed
if ! [ -x "$(command -v swagger-cli)" ]; then
  echo 'Error: swagger-cli is not installed.' >&2
  echo 'Attempting to install swagger-cli with NPM...' >&2
  npm install -g @apidevtools/swagger-cli
  if ! [ -x "$(command -v swagger-cli)" ]; then
    echo 'Error: swagger-cli is not available.' >&2
    exit 1
  fi
fi

# check if json-dereference is installed
if ! [ -x "$(command -v json-dereference)" ]; then
  echo 'Error: json-dereference is not installed.' >&2
  echo 'Attempting to install json-dereference with NPM...' >&2
  npm install -g json-dereference-cli
  if ! [ -x "$(command -v json-dereference)" ]; then
    echo 'Error: json-dereference is not available.' >&2
    exit 1
  fi
fi

echo -e "Updating submodule..."
git submodule update --init --recursive

echo -e "Bundling OpenAPI spec..."
swagger-cli bundle -o spec.json stacks-api/docs/openapi.yaml

echo -e "Processing bundled spec..."
sed -i '' -e 's/STACKS_API_VERSION/v4.0.1/g' spec.json
sed -i '' -e 's/btc/BTC/g' spec.json
sed -i '' -e 's/stx/STX/g' spec.json
sed -i '' -e 's/nft/NFT/g' spec.json

echo "Dereferencing bundled spec..."
json-dereference -s spec.json -o spec.json

echo -e "Generating client..."
openapi-generator generate -i spec.json -g go -o ../api

cd ../api
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context