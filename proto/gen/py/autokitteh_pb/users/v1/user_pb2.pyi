from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class User(_message.Message):
    __slots__ = ["user_id", "email", "display_name", "disabled", "default_org_id"]
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    DISPLAY_NAME_FIELD_NUMBER: _ClassVar[int]
    DISABLED_FIELD_NUMBER: _ClassVar[int]
    DEFAULT_ORG_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    email: str
    display_name: str
    disabled: bool
    default_org_id: str
    def __init__(self, user_id: _Optional[str] = ..., email: _Optional[str] = ..., display_name: _Optional[str] = ..., disabled: bool = ..., default_org_id: _Optional[str] = ...) -> None: ...
