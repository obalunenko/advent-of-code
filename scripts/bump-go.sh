#!/usr/bin/env bash

readonly CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
readonly ROOT_DIR="$(dirname "$CURRENT_DIR")"
readonly GO_MOD_FILE="${ROOT_DIR}/go.mod"

function main() {
  echo "Updating Go version:"

  local current_go_version="$(extract_current_version)"
  echo " - Current: ${current_go_version}"
  local escaped_current_go_version="$(echo "${current_go_version}" | sed 's/\./\\./g')"

  local go_version="${1}"
  local escaped_go_version="$(echo "${go_version}" | sed 's/\./\\./g')"
  echo " - New: ${go_version}"

  # bump mod files in all the modules
  for mod_file in $(find "${ROOT_DIR}" -name "go.mod" -not -path "${ROOT_DIR}/vendor/*" -not -path "${ROOT_DIR}/.git/*"); do
    bump_mod_file "${mod_file}" "${escaped_current_go_version}" "${escaped_go_version}"
  done

  # bump markdown files
  for f in $(find "${ROOT_DIR}" -name "*.md"); do
    bump_golang_docker_images "${f}" "${escaped_current_go_version}" "${escaped_go_version}"
  done

  # bump github action workflows
  for f in $(find "${ROOT_DIR}/.github/workflows" -name "*.yml"); do
    bump_ci_matrix "${f}" "${escaped_current_go_version}" "${escaped_go_version}"
  done
}

# it will replace the 'go-version: [${old_go_version}, 1.x]' with 'go-version: [${new_go_version}, 1.x]' in the given file
function bump_ci_matrix() {
  local file="${1}"
  local old_go_version="${2}"
  local new_go_version="${3}"

    sed "s/go-version: \[${old_go_version}/go-version: \[${new_go_version}/g" ${file} > ${file}.tmp
    mv ${file}.tmp ${file}
}

# it will replace the 'golang:${old_go_version}' with 'golang:${new_go_version}' in the given file
function bump_golang_docker_images() {
  local file="${1}"
  local old_go_version="${2}"
  local new_go_version="${3}"

    sed "s/golang:${old_go_version}/golang:${new_go_version}/g" ${file} > ${file}.tmp
    mv ${file}.tmp ${file}

}

# it will replace the 'go ${old_go_version}' with 'go ${new_go_version}' in the given go.mod file
function bump_mod_file() {
  local go_mod_file="${1}"
  local old_go_version="${2}"
  local new_go_version="${3}"

    sed "s/^go ${old_go_version}/go ${new_go_version}/g" ${go_mod_file} > ${go_mod_file}.tmp
    mv ${go_mod_file}.tmp ${go_mod_file}

}

# This function reads the reaper.go file and extracts the current version.
function extract_current_version() {
  cat "${GO_MOD_FILE}" | grep '^go .*' | sed 's/^go //g' | head -n 1
}

main "$@"