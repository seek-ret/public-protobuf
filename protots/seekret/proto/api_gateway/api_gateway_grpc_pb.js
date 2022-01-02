// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var seekret_proto_api_gateway_api_gateway_pb = require('../../../seekret/proto/api_gateway/api_gateway_pb.js');

function serialize_seekret_api_gateway_FetchWorkspaceConfigurationRequest(arg) {
  if (!(arg instanceof seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationRequest)) {
    throw new Error('Expected argument of type seekret.api_gateway.FetchWorkspaceConfigurationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_seekret_api_gateway_FetchWorkspaceConfigurationRequest(buffer_arg) {
  return seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_seekret_api_gateway_FetchWorkspaceConfigurationResponse(arg) {
  if (!(arg instanceof seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationResponse)) {
    throw new Error('Expected argument of type seekret.api_gateway.FetchWorkspaceConfigurationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_seekret_api_gateway_FetchWorkspaceConfigurationResponse(buffer_arg) {
  return seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ApiGatewayService = exports.ApiGatewayService = {
  fetchWorkspaceConfiguration: {
    path: '/seekret.api_gateway.ApiGateway/FetchWorkspaceConfiguration',
    requestStream: false,
    responseStream: false,
    requestType: seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationRequest,
    responseType: seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationResponse,
    requestSerialize: serialize_seekret_api_gateway_FetchWorkspaceConfigurationRequest,
    requestDeserialize: deserialize_seekret_api_gateway_FetchWorkspaceConfigurationRequest,
    responseSerialize: serialize_seekret_api_gateway_FetchWorkspaceConfigurationResponse,
    responseDeserialize: deserialize_seekret_api_gateway_FetchWorkspaceConfigurationResponse,
  },
};

exports.ApiGatewayClient = grpc.makeGenericClientConstructor(ApiGatewayService);
