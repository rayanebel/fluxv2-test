#!/bin/bash

#!/usr/bin/env bash
# CI tool to compute Kustomization and HelmRelease based on git diff
set -euo pipefail

REFERENCE="${GITOPS_REFERENCE:-origin/main}"

echo "REFERENCE: $REFERENCE"