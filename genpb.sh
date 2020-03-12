#!/bin/sh

protoc rarticlepb/rarticle.proto --go_out=plugins=grpc:.
