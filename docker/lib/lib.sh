function log_info() {
  echo "$1"
}


function log_fatal() {
  log_info "$1"
  exit 1
}


function check_execuables() {
  which $1 >/dev/null || log_fatal "$1 not found in PATH"
}


function check_deps() {
  check_execuables go
  check_execuables yarn
  check_execuables docker
}


function build_backend() {
  mkdir -p ./app
  GOOS=linux GOARCH=arm64 go build -o ./app/boltdbwebeditor.linux.arm64 ../src/cli || log_fatal "Failed to build backend linux.arm64"
  GOOS=linux GOARCH=amd64 go build -o ./app/boltdbwebeditor.linux.amd64 ../src/cli || log_fatal "Failed to build backend linux.amd64"
  log_info "Build backend linux.arm64 and linux.arm64 successfully"
}


function build_frontend() {
  mkdir -p ./app/static
  cd ../src/react
  BUILD_PATH='../../docker/app/static' yarn build || log_fatal "Failed to build frontend"
  cd -
  log_info "Build frontend successfully"
}


function build_image() {
  docker build -t boltdbwebeditor/boltdbwebeditor:dev ./ || log_fatal "Failed to build image"
  log_info "Build image successfully"
}


function is_builder_instance_exist() {
  docker buildx ls 2>&1 | grep boltdbwebeditor >/dev/null
}


function create_builder_instance() {
  docker buildx create --name boltdbwebeditor --driver docker-container && docker buildx use boltdbwebeditor
}


function buildx_image() {
  is_builder_instance_exist || create_builder_instance || log_fatal "Failed to create boltdbwebeditor builder instance"
  docker buildx build --platform linux/arm64,linux/amd64 -t boltdbwebeditor/boltdbwebeditor:dev . --push || log_fatal "Failed to buildx and push boltdbwebeditor"
  log_info "Buildx and push image successfully"
}
