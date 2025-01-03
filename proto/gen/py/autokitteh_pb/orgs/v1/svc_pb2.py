# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: autokitteh/orgs/v1/svc.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from autokitteh_pb.orgs.v1 import org_pb2 as autokitteh_dot_orgs_dot_v1_dot_org__pb2
from buf.validate import validate_pb2 as buf_dot_validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x61utokitteh/orgs/v1/svc.proto\x12\x12\x61utokitteh.orgs.v1\x1a\x1c\x61utokitteh/orgs/v1/org.proto\x1a\x1b\x62uf/validate/validate.proto\"C\n\rCreateRequest\x12\x32\n\x03org\x18\x01 \x01(\x0b\x32\x17.autokitteh.orgs.v1.OrgB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x03org\"1\n\x0e\x43reateResponse\x12\x1f\n\x06org_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x05orgId\"#\n\nGetRequest\x12\x15\n\x06org_id\x18\x01 \x01(\tR\x05orgId\"8\n\x0bGetResponse\x12)\n\x03org\x18\x01 \x01(\x0b\x32\x17.autokitteh.orgs.v1.OrgR\x03org\"C\n\rUpdateRequest\x12\x32\n\x03org\x18\x01 \x01(\x0b\x32\x17.autokitteh.orgs.v1.OrgB\x07\xfa\xf7\x18\x03\xc8\x01\x01R\x03org\"\x10\n\x0eUpdateResponse\"V\n\x10\x41\x64\x64MemberRequest\x12!\n\x07user_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x06userId\x12\x1f\n\x06org_id\x18\x02 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x05orgId\"\x13\n\x11\x41\x64\x64MemberResponse\"Y\n\x13RemoveMemberRequest\x12!\n\x07user_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x06userId\x12\x1f\n\x06org_id\x18\x02 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x05orgId\"\x16\n\x14RemoveMemberResponse\"5\n\x12ListMembersRequest\x12\x1f\n\x06org_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x05orgId\"?\n\x13ListMembersResponse\x12(\n\x08user_ids\x18\x01 \x03(\tB\r\xfa\xf7\x18\t\x92\x01\x06\"\x04r\x02\x10\x01R\x07userIds\"U\n\x0fIsMemberRequest\x12!\n\x07user_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x06userId\x12\x1f\n\x06org_id\x18\x02 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x05orgId\"/\n\x10IsMemberResponse\x12\x1b\n\tis_member\x18\x01 \x01(\x08R\x08isMember\":\n\x15GetOrgsForUserRequest\x12!\n\x07user_id\x18\x01 \x01(\tB\x08\xfa\xf7\x18\x04r\x02\x10\x01R\x06userId\"S\n\x16GetOrgsForUserResponse\x12\x39\n\x04orgs\x18\x01 \x03(\x0b\x32\x17.autokitteh.orgs.v1.OrgB\x0c\xfa\xf7\x18\x08\x92\x01\x05\"\x03\xc8\x01\x01R\x04orgs2\xd4\x05\n\x0bOrgsService\x12O\n\x06\x43reate\x12!.autokitteh.orgs.v1.CreateRequest\x1a\".autokitteh.orgs.v1.CreateResponse\x12\x46\n\x03Get\x12\x1e.autokitteh.orgs.v1.GetRequest\x1a\x1f.autokitteh.orgs.v1.GetResponse\x12O\n\x06Update\x12!.autokitteh.orgs.v1.UpdateRequest\x1a\".autokitteh.orgs.v1.UpdateResponse\x12X\n\tAddMember\x12$.autokitteh.orgs.v1.AddMemberRequest\x1a%.autokitteh.orgs.v1.AddMemberResponse\x12\x61\n\x0cRemoveMember\x12\'.autokitteh.orgs.v1.RemoveMemberRequest\x1a(.autokitteh.orgs.v1.RemoveMemberResponse\x12^\n\x0bListMembers\x12&.autokitteh.orgs.v1.ListMembersRequest\x1a\'.autokitteh.orgs.v1.ListMembersResponse\x12U\n\x08IsMember\x12#.autokitteh.orgs.v1.IsMemberRequest\x1a$.autokitteh.orgs.v1.IsMemberResponse\x12g\n\x0eGetOrgsForUser\x12).autokitteh.orgs.v1.GetOrgsForUserRequest\x1a*.autokitteh.orgs.v1.GetOrgsForUserResponseB\xd1\x01\n\x16\x63om.autokitteh.orgs.v1B\x08SvcProtoP\x01ZCgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/orgs/v1;orgsv1\xa2\x02\x03\x41OX\xaa\x02\x12\x41utokitteh.Orgs.V1\xca\x02\x12\x41utokitteh\\Orgs\\V1\xe2\x02\x1e\x41utokitteh\\Orgs\\V1\\GPBMetadata\xea\x02\x14\x41utokitteh::Orgs::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'autokitteh.orgs.v1.svc_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\026com.autokitteh.orgs.v1B\010SvcProtoP\001ZCgo.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/orgs/v1;orgsv1\242\002\003AOX\252\002\022Autokitteh.Orgs.V1\312\002\022Autokitteh\\Orgs\\V1\342\002\036Autokitteh\\Orgs\\V1\\GPBMetadata\352\002\024Autokitteh::Orgs::V1'
  _CREATEREQUEST.fields_by_name['org']._options = None
  _CREATEREQUEST.fields_by_name['org']._serialized_options = b'\372\367\030\003\310\001\001'
  _CREATERESPONSE.fields_by_name['org_id']._options = None
  _CREATERESPONSE.fields_by_name['org_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _UPDATEREQUEST.fields_by_name['org']._options = None
  _UPDATEREQUEST.fields_by_name['org']._serialized_options = b'\372\367\030\003\310\001\001'
  _ADDMEMBERREQUEST.fields_by_name['user_id']._options = None
  _ADDMEMBERREQUEST.fields_by_name['user_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _ADDMEMBERREQUEST.fields_by_name['org_id']._options = None
  _ADDMEMBERREQUEST.fields_by_name['org_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _REMOVEMEMBERREQUEST.fields_by_name['user_id']._options = None
  _REMOVEMEMBERREQUEST.fields_by_name['user_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _REMOVEMEMBERREQUEST.fields_by_name['org_id']._options = None
  _REMOVEMEMBERREQUEST.fields_by_name['org_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _LISTMEMBERSREQUEST.fields_by_name['org_id']._options = None
  _LISTMEMBERSREQUEST.fields_by_name['org_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _LISTMEMBERSRESPONSE.fields_by_name['user_ids']._options = None
  _LISTMEMBERSRESPONSE.fields_by_name['user_ids']._serialized_options = b'\372\367\030\t\222\001\006\"\004r\002\020\001'
  _ISMEMBERREQUEST.fields_by_name['user_id']._options = None
  _ISMEMBERREQUEST.fields_by_name['user_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _ISMEMBERREQUEST.fields_by_name['org_id']._options = None
  _ISMEMBERREQUEST.fields_by_name['org_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _GETORGSFORUSERREQUEST.fields_by_name['user_id']._options = None
  _GETORGSFORUSERREQUEST.fields_by_name['user_id']._serialized_options = b'\372\367\030\004r\002\020\001'
  _GETORGSFORUSERRESPONSE.fields_by_name['orgs']._options = None
  _GETORGSFORUSERRESPONSE.fields_by_name['orgs']._serialized_options = b'\372\367\030\010\222\001\005\"\003\310\001\001'
  _globals['_CREATEREQUEST']._serialized_start=111
  _globals['_CREATEREQUEST']._serialized_end=178
  _globals['_CREATERESPONSE']._serialized_start=180
  _globals['_CREATERESPONSE']._serialized_end=229
  _globals['_GETREQUEST']._serialized_start=231
  _globals['_GETREQUEST']._serialized_end=266
  _globals['_GETRESPONSE']._serialized_start=268
  _globals['_GETRESPONSE']._serialized_end=324
  _globals['_UPDATEREQUEST']._serialized_start=326
  _globals['_UPDATEREQUEST']._serialized_end=393
  _globals['_UPDATERESPONSE']._serialized_start=395
  _globals['_UPDATERESPONSE']._serialized_end=411
  _globals['_ADDMEMBERREQUEST']._serialized_start=413
  _globals['_ADDMEMBERREQUEST']._serialized_end=499
  _globals['_ADDMEMBERRESPONSE']._serialized_start=501
  _globals['_ADDMEMBERRESPONSE']._serialized_end=520
  _globals['_REMOVEMEMBERREQUEST']._serialized_start=522
  _globals['_REMOVEMEMBERREQUEST']._serialized_end=611
  _globals['_REMOVEMEMBERRESPONSE']._serialized_start=613
  _globals['_REMOVEMEMBERRESPONSE']._serialized_end=635
  _globals['_LISTMEMBERSREQUEST']._serialized_start=637
  _globals['_LISTMEMBERSREQUEST']._serialized_end=690
  _globals['_LISTMEMBERSRESPONSE']._serialized_start=692
  _globals['_LISTMEMBERSRESPONSE']._serialized_end=755
  _globals['_ISMEMBERREQUEST']._serialized_start=757
  _globals['_ISMEMBERREQUEST']._serialized_end=842
  _globals['_ISMEMBERRESPONSE']._serialized_start=844
  _globals['_ISMEMBERRESPONSE']._serialized_end=891
  _globals['_GETORGSFORUSERREQUEST']._serialized_start=893
  _globals['_GETORGSFORUSERREQUEST']._serialized_end=951
  _globals['_GETORGSFORUSERRESPONSE']._serialized_start=953
  _globals['_GETORGSFORUSERRESPONSE']._serialized_end=1036
  _globals['_ORGSSERVICE']._serialized_start=1039
  _globals['_ORGSSERVICE']._serialized_end=1763
# @@protoc_insertion_point(module_scope)
