# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: autokitteh/program/v1/program.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#autokitteh/program/v1/program.proto\x12\x15\x61utokitteh.program.v1\x1a\x1b\x62uf/validate/validate.proto\"Z\n\x0c\x43odeLocation\x12\x12\n\x04path\x18\x01 \x01(\tR\x04path\x12\x10\n\x03row\x18\x02 \x01(\rR\x03row\x12\x10\n\x03\x63ol\x18\x03 \x01(\rR\x03\x63ol\x12\x12\n\x04name\x18\x04 \x01(\tR\x04name\"`\n\tCallFrame\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12?\n\x08location\x18\x02 \x01(\x0b\x32#.autokitteh.program.v1.CodeLocationR\x08location\"\xf2\x01\n\x05\x45rror\x12\"\n\x07message\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x07message\x12L\n\tcallstack\x18\x02 \x03(\x0b\x32 .autokitteh.program.v1.CallFrameB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\tcallstack\x12=\n\x05\x65xtra\x18\x03 \x03(\x0b\x32\'.autokitteh.program.v1.Error.ExtraEntryR\x05\x65xtra\x1a\x38\n\nExtraEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x42\xea\x01\n\x19\x63om.autokitteh.program.v1B\x0cProgramProtoP\x01ZIgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/program/v1;programv1\xa2\x02\x03\x41PX\xaa\x02\x15\x41utokitteh.Program.V1\xca\x02\x15\x41utokitteh\\Program\\V1\xe2\x02!Autokitteh\\Program\\V1\\GPBMetadata\xea\x02\x17\x41utokitteh::Program::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'autokitteh.program.v1.program_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\031com.autokitteh.program.v1B\014ProgramProtoP\001ZIgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/program/v1;programv1\242\002\003APX\252\002\025Autokitteh.Program.V1\312\002\025Autokitteh\\Program\\V1\342\002!Autokitteh\\Program\\V1\\GPBMetadata\352\002\027Autokitteh::Program::V1'
  _ERROR_EXTRAENTRY._options = None
  _ERROR_EXTRAENTRY._serialized_options = b'8\001'
  _ERROR.fields_by_name['message']._options = None
  _ERROR.fields_by_name['message']._serialized_options = b'\372\367\030\004r\002\020\001'
  _ERROR.fields_by_name['callstack']._options = None
  _ERROR.fields_by_name['callstack']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _globals['_CODELOCATION']._serialized_start=91
  _globals['_CODELOCATION']._serialized_end=181
  _globals['_CALLFRAME']._serialized_start=183
  _globals['_CALLFRAME']._serialized_end=279
  _globals['_ERROR']._serialized_start=282
  _globals['_ERROR']._serialized_end=524
  _globals['_ERROR_EXTRAENTRY']._serialized_start=468
  _globals['_ERROR_EXTRAENTRY']._serialized_end=524
# @@protoc_insertion_point(module_scope)
