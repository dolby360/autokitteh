# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from autokitteh_pb.builds.v1 import svc_pb2 as autokitteh_dot_builds_dot_v1_dot_svc__pb2


class BuildsServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Get = channel.unary_unary(
                '/autokitteh.builds.v1.BuildsService/Get',
                request_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.GetRequest.SerializeToString,
                response_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.GetResponse.FromString,
                )
        self.List = channel.unary_unary(
                '/autokitteh.builds.v1.BuildsService/List',
                request_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.ListRequest.SerializeToString,
                response_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.ListResponse.FromString,
                )
        self.Save = channel.unary_unary(
                '/autokitteh.builds.v1.BuildsService/Save',
                request_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.SaveRequest.SerializeToString,
                response_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.SaveResponse.FromString,
                )
        self.Remove = channel.unary_unary(
                '/autokitteh.builds.v1.BuildsService/Remove',
                request_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.RemoveRequest.SerializeToString,
                response_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.RemoveResponse.FromString,
                )
        self.Download = channel.unary_unary(
                '/autokitteh.builds.v1.BuildsService/Download',
                request_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.DownloadRequest.SerializeToString,
                response_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.DownloadResponse.FromString,
                )


class BuildsServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def List(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Save(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Remove(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Download(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_BuildsServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.GetRequest.FromString,
                    response_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.GetResponse.SerializeToString,
            ),
            'List': grpc.unary_unary_rpc_method_handler(
                    servicer.List,
                    request_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.ListRequest.FromString,
                    response_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.ListResponse.SerializeToString,
            ),
            'Save': grpc.unary_unary_rpc_method_handler(
                    servicer.Save,
                    request_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.SaveRequest.FromString,
                    response_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.SaveResponse.SerializeToString,
            ),
            'Remove': grpc.unary_unary_rpc_method_handler(
                    servicer.Remove,
                    request_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.RemoveRequest.FromString,
                    response_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.RemoveResponse.SerializeToString,
            ),
            'Download': grpc.unary_unary_rpc_method_handler(
                    servicer.Download,
                    request_deserializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.DownloadRequest.FromString,
                    response_serializer=autokitteh_dot_builds_dot_v1_dot_svc__pb2.DownloadResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'autokitteh.builds.v1.BuildsService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class BuildsService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.builds.v1.BuildsService/Get',
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.GetRequest.SerializeToString,
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.GetResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def List(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.builds.v1.BuildsService/List',
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.ListRequest.SerializeToString,
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.ListResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Save(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.builds.v1.BuildsService/Save',
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.SaveRequest.SerializeToString,
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.SaveResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Remove(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.builds.v1.BuildsService/Remove',
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.RemoveRequest.SerializeToString,
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.RemoveResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Download(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.builds.v1.BuildsService/Download',
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.DownloadRequest.SerializeToString,
            autokitteh_dot_builds_dot_v1_dot_svc__pb2.DownloadResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
