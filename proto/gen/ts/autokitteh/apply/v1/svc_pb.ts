// @generated by protoc-gen-es v1.5.1 with parameter "target=ts"
// @generated from file autokitteh/apply/v1/svc.proto (package autokitteh.apply.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message autokitteh.apply.v1.ApplyRequest
 */
export class ApplyRequest extends Message<ApplyRequest> {
  /**
   * @generated from field: string manifest = 1;
   */
  manifest = "";

  /**
   * @generated from field: string path = 2;
   */
  path = "";

  /**
   * @generated from field: string project_name = 3;
   */
  projectName = "";

  constructor(data?: PartialMessage<ApplyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.apply.v1.ApplyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "manifest", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "project_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApplyRequest {
    return new ApplyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApplyRequest {
    return new ApplyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApplyRequest {
    return new ApplyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ApplyRequest | PlainMessage<ApplyRequest> | undefined, b: ApplyRequest | PlainMessage<ApplyRequest> | undefined): boolean {
    return proto3.util.equals(ApplyRequest, a, b);
  }
}

/**
 * @generated from message autokitteh.apply.v1.ApplyResponse
 */
export class ApplyResponse extends Message<ApplyResponse> {
  /**
   * @generated from field: repeated string logs = 1;
   */
  logs: string[] = [];

  /**
   * @generated from field: repeated string project_ids = 2;
   */
  projectIds: string[] = [];

  constructor(data?: PartialMessage<ApplyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.apply.v1.ApplyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "logs", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "project_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApplyResponse {
    return new ApplyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApplyResponse {
    return new ApplyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApplyResponse {
    return new ApplyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ApplyResponse | PlainMessage<ApplyResponse> | undefined, b: ApplyResponse | PlainMessage<ApplyResponse> | undefined): boolean {
    return proto3.util.equals(ApplyResponse, a, b);
  }
}

