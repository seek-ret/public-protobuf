// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var seekret_proto_auth_auth_pb = require('../../../seekret/proto/auth/auth_pb.js');

function serialize_seekret_authentication_AuthenticateRequest(arg) {
  if (!(arg instanceof seekret_proto_auth_auth_pb.AuthenticateRequest)) {
    throw new Error('Expected argument of type seekret.authentication.AuthenticateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_seekret_authentication_AuthenticateRequest(buffer_arg) {
  return seekret_proto_auth_auth_pb.AuthenticateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_seekret_authentication_AuthenticateResponse(arg) {
  if (!(arg instanceof seekret_proto_auth_auth_pb.AuthenticateResponse)) {
    throw new Error('Expected argument of type seekret.authentication.AuthenticateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_seekret_authentication_AuthenticateResponse(buffer_arg) {
  return seekret_proto_auth_auth_pb.AuthenticateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthService = exports.AuthService = {
  authenticate: {
    path: '/seekret.authentication.Auth/Authenticate',
    requestStream: false,
    responseStream: false,
    requestType: seekret_proto_auth_auth_pb.AuthenticateRequest,
    responseType: seekret_proto_auth_auth_pb.AuthenticateResponse,
    requestSerialize: serialize_seekret_authentication_AuthenticateRequest,
    requestDeserialize: deserialize_seekret_authentication_AuthenticateRequest,
    responseSerialize: serialize_seekret_authentication_AuthenticateResponse,
    responseDeserialize: deserialize_seekret_authentication_AuthenticateResponse,
  },
};

exports.AuthClient = grpc.makeGenericClientConstructor(AuthService);
