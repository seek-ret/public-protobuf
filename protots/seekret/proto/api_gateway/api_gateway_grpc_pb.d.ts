// GENERATED CODE -- DO NOT EDIT!

// package: seekret.api_gateway
// file: seekret/proto/api_gateway/api_gateway.proto

import * as seekret_proto_api_gateway_api_gateway_pb from "../../../seekret/proto/api_gateway/api_gateway_pb";
import * as grpc from "@grpc/grpc-js";

interface IApiGatewayService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
  fetchWorkspaceConfiguration: grpc.MethodDefinition<seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationRequest, seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationResponse>;
}

export const ApiGatewayService: IApiGatewayService;

export interface IApiGatewayServer extends grpc.UntypedServiceImplementation {
  fetchWorkspaceConfiguration: grpc.handleUnaryCall<seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationRequest, seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationResponse>;
}

export class ApiGatewayClient extends grpc.Client {
  constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
  fetchWorkspaceConfiguration(argument: seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationRequest, callback: grpc.requestCallback<seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationResponse>): grpc.ClientUnaryCall;
  fetchWorkspaceConfiguration(argument: seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationRequest, metadataOrOptions: grpc.Metadata | grpc.CallOptions | null, callback: grpc.requestCallback<seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationResponse>): grpc.ClientUnaryCall;
  fetchWorkspaceConfiguration(argument: seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationRequest, metadata: grpc.Metadata | null, options: grpc.CallOptions | null, callback: grpc.requestCallback<seekret_proto_api_gateway_api_gateway_pb.FetchWorkspaceConfigurationResponse>): grpc.ClientUnaryCall;
}
