#!/bin/bash
protoc --go_out=./pb --go-grpc_out=./pb coucoo.proto
