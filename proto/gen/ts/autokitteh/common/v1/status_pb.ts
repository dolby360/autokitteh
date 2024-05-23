// @generated by protoc-gen-es v1.5.1 with parameter "target=ts"
// @generated from file autokitteh/common/v1/status.proto (package autokitteh.common.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message autokitteh.common.v1.Status
 */
export class Status extends Message<Status> {
  /**
   * @generated from field: autokitteh.common.v1.Status.Code code = 1;
   */
  code = Status_Code.UNSPECIFIED;

  /**
   * @generated from field: string message = 2;
   */
  message = "";

  constructor(data?: PartialMessage<Status>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.common.v1.Status";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "code", kind: "enum", T: proto3.getEnumType(Status_Code) },
    { no: 2, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Status {
    return new Status().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Status {
    return new Status().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Status {
    return new Status().fromJsonString(jsonString, options);
  }

  static equals(a: Status | PlainMessage<Status> | undefined, b: Status | PlainMessage<Status> | undefined): boolean {
    return proto3.util.equals(Status, a, b);
  }
}

/**
 * @generated from enum autokitteh.common.v1.Status.Code
 */
export enum Status_Code {
  /**
   * @generated from enum value: CODE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: CODE_OK = 1;
   */
  OK = 1,

  /**
   * @generated from enum value: CODE_WARNING = 2;
   */
  WARNING = 2,

  /**
   * @generated from enum value: CODE_ERROR = 3;
   */
  ERROR = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Status_Code)
proto3.util.setEnumType(Status_Code, "autokitteh.common.v1.Status.Code", [
  { no: 0, name: "CODE_UNSPECIFIED" },
  { no: 1, name: "CODE_OK" },
  { no: 2, name: "CODE_WARNING" },
  { no: 3, name: "CODE_ERROR" },
]);
