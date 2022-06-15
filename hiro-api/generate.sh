# check if swagger-cli is installed
if ! [ -x "$(command -v swagger-cli)" ]; then
  echo 'Error: swagger-cli is not installed.' >&2
  exit 1
fi

# make sure the submodule is up to date
git submodule update --init --recursive

swagger-cli bundle -o BundledSpec.json stacks-api/docs/openapi.yaml