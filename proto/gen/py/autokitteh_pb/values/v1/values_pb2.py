# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: autokitteh/values/v1/values.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from autokitteh_pb.program.v1 import module_pb2 as autokitteh_dot_program_dot_v1_dot_module__pb2
from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!autokitteh/values/v1/values.proto\x12\x14\x61utokitteh.values.v1\x1a\"autokitteh/program/v1/module.proto\x1a\x1b\x62uf/validate/validate.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\t\n\x07Nothing\"\x16\n\x06String\x12\x0c\n\x01v\x18\x01 \x01(\tR\x01v\"\x17\n\x07Integer\x12\x0c\n\x01v\x18\x01 \x01(\x03R\x01v\"\x15\n\x05\x46loat\x12\x0c\n\x01v\x18\x01 \x01(\x01R\x01v\"\x17\n\x07\x42oolean\x12\x0c\n\x01v\x18\x01 \x01(\x08R\x01v\"&\n\x06Symbol\x12\x1c\n\x04name\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x04name\"A\n\x04List\x12\x39\n\x02vs\x18\x01 \x03(\x0b\x32\x1b.autokitteh.values.v1.ValueB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\x02vs\"@\n\x03Set\x12\x39\n\x02vs\x18\x01 \x03(\x0b\x32\x1b.autokitteh.values.v1.ValueB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\x02vs\"\x15\n\x05\x42ytes\x12\x0c\n\x01v\x18\x01 \x01(\x0cR\x01v\"\xbb\x01\n\x04\x44ict\x12\x43\n\x05items\x18\x01 \x03(\x0b\x32\x1f.autokitteh.values.v1.Dict.ItemB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\x05items\x1an\n\x04Item\x12\x32\n\x01k\x18\x01 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x01k\x12\x32\n\x01v\x18\x02 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x01v\"9\n\x04Time\x12\x31\n\x01v\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x01v\"<\n\x08\x44uration\x12\x30\n\x01v\x18\x01 \x01(\x0b\x32\x19.google.protobuf.DurationB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x01v\"\xf0\x01\n\x06Struct\x12\x38\n\x04\x63tor\x18\x01 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x04\x63tor\x12T\n\x06\x66ields\x18\x02 \x03(\x0b\x32(.autokitteh.values.v1.Struct.FieldsEntryB\x12\xfa\xf7\x18\x0e\x9a\x01\x0b\"\x04r\x02\x10\x01*\x03\xc8\x01\x01R\x06\x66ields\x1aV\n\x0b\x46ieldsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x31\n\x05value\x18\x02 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueR\x05value:\x02\x38\x01\"\xd8\x01\n\x06Module\x12\x1c\n\x04name\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x04name\x12W\n\x07members\x18\x02 \x03(\x0b\x32).autokitteh.values.v1.Module.MembersEntryB\x12\xfa\xf7\x18\x0e\x9a\x01\x0b\"\x04r\x02\x10\x01*\x03\xc8\x01\x01R\x07members\x1aW\n\x0cMembersEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x31\n\x05value\x18\x02 \x01(\x0b\x32\x1b.autokitteh.values.v1.ValueR\x05value:\x02\x38\x01\"\xbb\x01\n\x08\x46unction\x12)\n\x0b\x65xecutor_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\nexecutorId\x12\x1c\n\x04name\x18\x02 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x04name\x12<\n\x04\x64\x65sc\x18\x03 \x01(\x0b\x32\x1f.autokitteh.program.v1.FunctionB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x04\x64\x65sc\x12\x12\n\x04\x64\x61ta\x18\x04 \x01(\x0cR\x04\x64\x61ta\x12\x14\n\x05\x66lags\x18\x05 \x03(\tR\x05\x66lags\"\xd3\x06\n\x05Value\x12\x39\n\x07nothing\x18\x01 \x01(\x0b\x32\x1d.autokitteh.values.v1.NothingH\x00R\x07nothing\x12\x39\n\x07\x62oolean\x18\x02 \x01(\x0b\x32\x1d.autokitteh.values.v1.BooleanH\x00R\x07\x62oolean\x12\x36\n\x06string\x18\x03 \x01(\x0b\x32\x1c.autokitteh.values.v1.StringH\x00R\x06string\x12\x39\n\x07integer\x18\x04 \x01(\x0b\x32\x1d.autokitteh.values.v1.IntegerH\x00R\x07integer\x12\x33\n\x05\x66loat\x18\x05 \x01(\x0b\x32\x1b.autokitteh.values.v1.FloatH\x00R\x05\x66loat\x12\x30\n\x04list\x18\x06 \x01(\x0b\x32\x1a.autokitteh.values.v1.ListH\x00R\x04list\x12-\n\x03set\x18\x07 \x01(\x0b\x32\x19.autokitteh.values.v1.SetH\x00R\x03set\x12\x30\n\x04\x64ict\x18\x08 \x01(\x0b\x32\x1a.autokitteh.values.v1.DictH\x00R\x04\x64ict\x12\x33\n\x05\x62ytes\x18\t \x01(\x0b\x32\x1b.autokitteh.values.v1.BytesH\x00R\x05\x62ytes\x12\x30\n\x04time\x18\n \x01(\x0b\x32\x1a.autokitteh.values.v1.TimeH\x00R\x04time\x12<\n\x08\x64uration\x18\x0b \x01(\x0b\x32\x1e.autokitteh.values.v1.DurationH\x00R\x08\x64uration\x12\x36\n\x06struct\x18\x0c \x01(\x0b\x32\x1c.autokitteh.values.v1.StructH\x00R\x06struct\x12\x36\n\x06module\x18\r \x01(\x0b\x32\x1c.autokitteh.values.v1.ModuleH\x00R\x06module\x12\x36\n\x06symbol\x18\x0e \x01(\x0b\x32\x1c.autokitteh.values.v1.SymbolH\x00R\x06symbol\x12<\n\x08\x66unction\x18\x0f \x01(\x0b\x32\x1e.autokitteh.values.v1.FunctionH\x00R\x08\x66unctionB\x0e\n\x04type\x12\x06\xfa\xf7\x18\x02\x08\x01\x42\xe2\x01\n\x18\x63om.autokitteh.values.v1B\x0bValuesProtoP\x01ZGgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/values/v1;valuesv1\xa2\x02\x03\x41VX\xaa\x02\x14\x41utokitteh.Values.V1\xca\x02\x14\x41utokitteh\\Values\\V1\xe2\x02 Autokitteh\\Values\\V1\\GPBMetadata\xea\x02\x16\x41utokitteh::Values::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'autokitteh.values.v1.values_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\030com.autokitteh.values.v1B\013ValuesProtoP\001ZGgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/values/v1;valuesv1\242\002\003AVX\252\002\024Autokitteh.Values.V1\312\002\024Autokitteh\\Values\\V1\342\002 Autokitteh\\Values\\V1\\GPBMetadata\352\002\026Autokitteh::Values::V1'
  _SYMBOL.fields_by_name['name']._options = None
  _SYMBOL.fields_by_name['name']._serialized_options = b'\372\367\030\004r\002\020\001'
  _LIST.fields_by_name['vs']._options = None
  _LIST.fields_by_name['vs']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _SET.fields_by_name['vs']._options = None
  _SET.fields_by_name['vs']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _DICT_ITEM.fields_by_name['k']._options = None
  _DICT_ITEM.fields_by_name['k']._serialized_options = b'\372\367\030\003\310\001\001'
  _DICT_ITEM.fields_by_name['v']._options = None
  _DICT_ITEM.fields_by_name['v']._serialized_options = b'\372\367\030\003\310\001\001'
  _DICT.fields_by_name['items']._options = None
  _DICT.fields_by_name['items']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _TIME.fields_by_name['v']._options = None
  _TIME.fields_by_name['v']._serialized_options = b'\372\367\030\003\310\001\001'
  _DURATION.fields_by_name['v']._options = None
  _DURATION.fields_by_name['v']._serialized_options = b'\372\367\030\003\310\001\001'
  _STRUCT_FIELDSENTRY._options = None
  _STRUCT_FIELDSENTRY._serialized_options = b'8\001'
  _STRUCT.fields_by_name['ctor']._options = None
  _STRUCT.fields_by_name['ctor']._serialized_options = b'\372\367\030\003\310\001\001'
  _STRUCT.fields_by_name['fields']._options = None
  _STRUCT.fields_by_name['fields']._serialized_options = b'\372\367\030\016\232\001\013\"\004r\002\020\001*\003\310\001\001'
  _MODULE_MEMBERSENTRY._options = None
  _MODULE_MEMBERSENTRY._serialized_options = b'8\001'
  _MODULE.fields_by_name['name']._options = None
  _MODULE.fields_by_name['name']._serialized_options = b'\372\367\030\004r\002\020\001'
  _MODULE.fields_by_name['members']._options = None
  _MODULE.fields_by_name['members']._serialized_options = b'\372\367\030\016\232\001\013\"\004r\002\020\001*\003\310\001\001'
  _FUNCTION.fields_by_name['executor_id']._options = None
  _FUNCTION.fields_by_name['executor_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _FUNCTION.fields_by_name['name']._options = None
  _FUNCTION.fields_by_name['name']._serialized_options = b'\372\367\030\004r\002\020\001'
  _FUNCTION.fields_by_name['desc']._options = None
  _FUNCTION.fields_by_name['desc']._serialized_options = b'\372\367\030\003\310\001\001'
  _VALUE.oneofs_by_name['type']._options = None
  _VALUE.oneofs_by_name['type']._serialized_options = b'\372\367\030\002\010\001'
  _globals['_NOTHING']._serialized_start=189
  _globals['_NOTHING']._serialized_end=198
  _globals['_STRING']._serialized_start=200
  _globals['_STRING']._serialized_end=222
  _globals['_INTEGER']._serialized_start=224
  _globals['_INTEGER']._serialized_end=247
  _globals['_FLOAT']._serialized_start=249
  _globals['_FLOAT']._serialized_end=270
  _globals['_BOOLEAN']._serialized_start=272
  _globals['_BOOLEAN']._serialized_end=295
  _globals['_SYMBOL']._serialized_start=297
  _globals['_SYMBOL']._serialized_end=335
  _globals['_LIST']._serialized_start=337
  _globals['_LIST']._serialized_end=402
  _globals['_SET']._serialized_start=404
  _globals['_SET']._serialized_end=468
  _globals['_BYTES']._serialized_start=470
  _globals['_BYTES']._serialized_end=491
  _globals['_DICT']._serialized_start=494
  _globals['_DICT']._serialized_end=681
  _globals['_DICT_ITEM']._serialized_start=571
  _globals['_DICT_ITEM']._serialized_end=681
  _globals['_TIME']._serialized_start=683
  _globals['_TIME']._serialized_end=740
  _globals['_DURATION']._serialized_start=742
  _globals['_DURATION']._serialized_end=802
  _globals['_STRUCT']._serialized_start=805
  _globals['_STRUCT']._serialized_end=1045
  _globals['_STRUCT_FIELDSENTRY']._serialized_start=959
  _globals['_STRUCT_FIELDSENTRY']._serialized_end=1045
  _globals['_MODULE']._serialized_start=1048
  _globals['_MODULE']._serialized_end=1264
  _globals['_MODULE_MEMBERSENTRY']._serialized_start=1177
  _globals['_MODULE_MEMBERSENTRY']._serialized_end=1264
  _globals['_FUNCTION']._serialized_start=1267
  _globals['_FUNCTION']._serialized_end=1454
  _globals['_VALUE']._serialized_start=1457
  _globals['_VALUE']._serialized_end=2308
# @@protoc_insertion_point(module_scope)
