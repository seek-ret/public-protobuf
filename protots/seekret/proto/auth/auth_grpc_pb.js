// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var seekret_proto_auth_auth_pb = require('../../../seekret/proto/auth/auth_pb.js');

function serialize_seekret_auth_ClientAuthenticateRequest(arg) {
  if (!(arg instanceof seekret_proto_auth_auth_pb.ClientAuthenticateRequest)) {
    throw new Error('Expected argument of type seekret.auth.ClientAuthenticateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_seekret_auth_ClientAuthenticateRequest(buffer_arg) {
  return seekret_proto_auth_auth_pb.ClientAuthenticateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_seekret_auth_ClientAuthenticateResponse(arg) {
  if (!(arg instanceof seekret_proto_auth_auth_pb.ClientAuthenticateResponse)) {
    throw new Error('Expected argument of type seekret.auth.ClientAuthenticateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_seekret_auth_ClientAuthenticateResponse(buffer_arg) {
  return seekret_proto_auth_auth_pb.ClientAuthenticateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthService = exports.AuthService = {
  clientAuthenticate: {
    path: '/seekret.auth.Auth/ClientAuthenticate',
    requestStream: false,
    responseStream: false,
    requestType: seekret_proto_auth_auth_pb.ClientAuthenticateRequest,
    responseType: seekret_proto_auth_auth_pb.ClientAuthenticateResponse,
    requestSerialize: serialize_seekret_auth_ClientAuthenticateRequest,
    requestDeserialize: deserialize_seekret_auth_ClientAuthenticateRequest,
    responseSerialize: serialize_seekret_auth_ClientAuthenticateResponse,
    responseDeserialize: deserialize_seekret_auth_ClientAuthenticateResponse,
  },
};

exports.AuthClient = grpc.makeGenericClientConstructor(AuthService);
