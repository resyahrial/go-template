# Start the HTTP service locally
env="dev"
if [[ ! -z "$1" ]]
then
  env=$1
fi

source files/etc/env/env.sh
go run ./cmd/app-http/*.go -env=$env