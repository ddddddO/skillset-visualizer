#! /usr/bin/env bash

set -ex

CLUSTER=work-gcc

gcloud container clusters get-credentials ${CLUSTER}
