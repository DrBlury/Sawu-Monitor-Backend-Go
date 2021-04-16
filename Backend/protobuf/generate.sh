#!/bin/bash

protoc --proto_path=proto-files --go_out=../generated ./proto-files/sawu/*.proto