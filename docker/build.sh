function log_info() {
  echo "$1"
}

function log_fatal() {
  log_info "$1"
  exit 1
}

function check_file() {
  which $1 >/dev/null || log_fatal "$1 not found in PATH"
}

function build_backend() {
  mkdir -p ./app
  GOOS=linux GOARCH=arm64 go build -o ./app/boltdbwebeditor ../src/cli || log_fatal "Failed to build backend"
  log_info "Build backend successfully"
}


function build_frontend() {
  mkdir -p ./app/static
  cd ../src/react
  BUILD_PATH='../../docker/app/static' yarn build || log_fatal "Failed to build frontend"
  cd -
  log_info "Build frontend successfully"
}

function build_image() {
  docker build -t boltdbwebeditor/boltdbwebeditor ./ || log_fatal "Failed to build image"
  log_info "Build image successfully"
}

check_file go
check_file yarn

build_frontend
build_backend
build_image