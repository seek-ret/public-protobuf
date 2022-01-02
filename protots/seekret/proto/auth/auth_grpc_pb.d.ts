// GENERATED CODE -- DO NOT EDIT!

// package: seekret.auth
// file: seekret/proto/auth/auth.proto

import * as seekret_proto_auth_auth_pb from "../../../seekret/proto/auth/auth_pb";
import * as grpc from "@grpc/grpc-js";

interface IAuthService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  clientAuthenticate: grpc.MethodDefinition<seekret_proto_auth_auth_pb.ClientAuthenticateRequest, seekret_proto_auth_auth_pb.ClientAuthenticateResponse>;
}

export const AuthService: IAuthService;

export interface IAuthServer extends grpc.UntypedServiceImplementation {
  clientAuthenticate: grpc.handleUnaryCall<seekret_proto_auth_auth_pb.ClientAuthenticateRequest, seekret_proto_auth_auth_pb.ClientAuthenticateResponse>;
}

export class AuthClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  clientAuthenticate(argument: seekret_proto_auth_auth_pb.ClientAuthenticateRequest, callback: grpc.requestCallback<seekret_proto_auth_auth_pb.ClientAuthenticateResponse>): grpc.ClientUnaryCall;
  clientAuthenticate(argument: seekret_proto_auth_auth_pb.ClientAuthenticateRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<seekret_proto_auth_auth_pb.ClientAuthenticateResponse>): grpc.ClientUnaryCall;
  clientAuthenticate(argument: seekret_proto_auth_auth_pb.ClientAuthenticateRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<seekret_proto_auth_auth_pb.ClientAuthenticateResponse>): grpc.ClientUnaryCall;
}
