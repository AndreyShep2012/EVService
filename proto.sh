#!/bin/sh

protoc -I=ShippingService/protobuf/proto --go_out=plugins=grpc:ShippingService/protobuf/pb ShippingService/protobuf/proto/*.proto