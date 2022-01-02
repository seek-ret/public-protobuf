// GENERATED CODE -- DO NOT EDIT!

// package: seekret.authentication
// file: seekret/proto/auth/auth.proto

import * as seekret_proto_auth_auth_pb from "../../../seekret/proto/auth/auth_pb";
import * as grpc from "@grpc/grpc-js";

interface IAuthService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  authenticate: grpc.MethodDefinition<seekret_proto_auth_auth_pb.AuthenticateRequest, seekret_proto_auth_auth_pb.AuthenticateResponse>;
}

export const AuthService: IAuthService;

export interface IAuthServer extends grpc.UntypedServiceImplementation {
  authenticate: grpc.handleUnaryCall<seekret_proto_auth_auth_pb.AuthenticateRequest, seekret_proto_auth_auth_pb.AuthenticateResponse>;
}

export class AuthClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  authenticate(argument: seekret_proto_auth_auth_pb.AuthenticateRequest, callback: grpc.requestCallback<seekret_proto_auth_auth_pb.AuthenticateResponse>): grpc.ClientUnaryCall;
  authenticate(argument: seekret_proto_auth_auth_pb.AuthenticateRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<seekret_proto_auth_auth_pb.AuthenticateResponse>): grpc.ClientUnaryCall;
  authenticate(argument: seekret_proto_auth_auth_pb.AuthenticateRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<seekret_proto_auth_auth_pb.AuthenticateResponse>): grpc.ClientUnaryCall;
}
