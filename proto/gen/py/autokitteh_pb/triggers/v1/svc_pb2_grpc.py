# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from autokitteh_pb.triggers.v1 import svc_pb2 as autokitteh_dot_triggers_dot_v1_dot_svc__pb2


class TriggersServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Create = channel.unary_unary(
                '/autokitteh.triggers.v1.TriggersService/Create',
                request_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.CreateRequest.SerializeToString,
                response_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.CreateResponse.FromString,
                )
        self.Update = channel.unary_unary(
                '/autokitteh.triggers.v1.TriggersService/Update',
                request_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.UpdateRequest.SerializeToString,
                response_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.UpdateResponse.FromString,
                )
        self.Delete = channel.unary_unary(
                '/autokitteh.triggers.v1.TriggersService/Delete',
                request_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.DeleteRequest.SerializeToString,
                response_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.DeleteResponse.FromString,
                )
        self.Get = channel.unary_unary(
                '/autokitteh.triggers.v1.TriggersService/Get',
                request_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.GetRequest.SerializeToString,
                response_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.GetResponse.FromString,
                )
        self.List = channel.unary_unary(
                '/autokitteh.triggers.v1.TriggersService/List',
                request_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.ListRequest.SerializeToString,
                response_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.ListResponse.FromString,
                )


class TriggersServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

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


def add_TriggersServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.CreateRequest.FromString,
                    response_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.CreateResponse.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.UpdateRequest.FromString,
                    response_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.UpdateResponse.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.DeleteRequest.FromString,
                    response_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.DeleteResponse.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.GetRequest.FromString,
                    response_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.GetResponse.SerializeToString,
            ),
            'List': grpc.unary_unary_rpc_method_handler(
                    servicer.List,
                    request_deserializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.ListRequest.FromString,
                    response_serializer=autokitteh_dot_triggers_dot_v1_dot_svc__pb2.ListResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'autokitteh.triggers.v1.TriggersService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TriggersService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.triggers.v1.TriggersService/Create',
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.CreateRequest.SerializeToString,
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.CreateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.triggers.v1.TriggersService/Update',
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.UpdateRequest.SerializeToString,
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.UpdateResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/autokitteh.triggers.v1.TriggersService/Delete',
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.DeleteRequest.SerializeToString,
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.DeleteResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/autokitteh.triggers.v1.TriggersService/Get',
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.GetRequest.SerializeToString,
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.GetResponse.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/autokitteh.triggers.v1.TriggersService/List',
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.ListRequest.SerializeToString,
            autokitteh_dot_triggers_dot_v1_dot_svc__pb2.ListResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)