// @generated by protoc-gen-es v1.5.1 with parameter "target=ts"
// @generated from file autokitteh/envs/v1/env.proto (package autokitteh.envs.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message autokitteh.envs.v1.Env
 */
export class Env extends Message<Env> {
  /**
   * @generated from field: string env_id = 1;
   */
  envId = "";

  /**
   * @generated from field: string project_id = 2;
   */
  projectId = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  constructor(data?: PartialMessage<Env>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.envs.v1.Env";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "env_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "project_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Env {
    return new Env().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Env {
    return new Env().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Env {
    return new Env().fromJsonString(jsonString, options);
  }

  static equals(a: Env | PlainMessage<Env> | undefined, b: Env | PlainMessage<Env> | undefined): boolean {
    return proto3.util.equals(Env, a, b);
  }
}

/**
 * @generated from message autokitteh.envs.v1.EnvVar
 */
export class EnvVar extends Message<EnvVar> {
  /**
   * @generated from field: string env_id = 1;
   */
  envId = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string value = 3;
   */
  value = "";

  /**
   * @generated from field: bool is_secret = 4;
   */
  isSecret = false;

  constructor(data?: PartialMessage<EnvVar>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.envs.v1.EnvVar";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "env_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "value", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "is_secret", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EnvVar {
    return new EnvVar().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EnvVar {
    return new EnvVar().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EnvVar {
    return new EnvVar().fromJsonString(jsonString, options);
  }

  static equals(a: EnvVar | PlainMessage<EnvVar> | undefined, b: EnvVar | PlainMessage<EnvVar> | undefined): boolean {
    return proto3.util.equals(EnvVar, a, b);
  }
}

