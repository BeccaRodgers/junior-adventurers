#!/bin/bash

set -eu -o pipefail

cd "$(dirname "${BASH_SOURCE[0]}")"

find . -name '*_templ.go' -delete

go generate ./...