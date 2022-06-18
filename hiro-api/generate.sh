# for macOS
export HOMEBREW_NO_ENV_HINTS=1

echo -e "Updating submodule..."
git submodule update --init --recursive

echo -e "Generating client..."
go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0 --package api --config oapi.cfg.yaml stacks-api/docs/openapi.yaml > ../api.gen.go
