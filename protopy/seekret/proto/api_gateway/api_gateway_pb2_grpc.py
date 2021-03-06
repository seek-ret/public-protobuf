# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from seekret.proto.api_gateway import api_gateway_pb2 as seekret_dot_proto_dot_api__gateway_dot_api__gateway__pb2


class ApiGatewayStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.FetchWorkspaceConfiguration = channel.unary_unary(
                '/seekret.api_gateway.ApiGateway/FetchWorkspaceConfiguration',
                request_serializer=seekret_dot_proto_dot_api__gateway_dot_api__gateway__pb2.FetchWorkspaceConfigurationRequest.SerializeToString,
                response_deserializer=seekret_dot_proto_dot_api__gateway_dot_api__gateway__pb2.FetchWorkspaceConfigurationResponse.FromString,
                )


class ApiGatewayServicer(object):
    """Missing associated documentation comment in .proto file."""

    def FetchWorkspaceConfiguration(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ApiGatewayServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'FetchWorkspaceConfiguration': grpc.unary_unary_rpc_method_handler(
                    servicer.FetchWorkspaceConfiguration,
                    request_deserializer=seekret_dot_proto_dot_api__gateway_dot_api__gateway__pb2.FetchWorkspaceConfigurationRequest.FromString,
                    response_serializer=seekret_dot_proto_dot_api__gateway_dot_api__gateway__pb2.FetchWorkspaceConfigurationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'seekret.api_gateway.ApiGateway', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ApiGateway(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def FetchWorkspaceConfiguration(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/seekret.api_gateway.ApiGateway/FetchWorkspaceConfiguration',
            seekret_dot_proto_dot_api__gateway_dot_api__gateway__pb2.FetchWorkspaceConfigurationRequest.SerializeToString,
            seekret_dot_proto_dot_api__gateway_dot_api__gateway__pb2.FetchWorkspaceConfigurationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
