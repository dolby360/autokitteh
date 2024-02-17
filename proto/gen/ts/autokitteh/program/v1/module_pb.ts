// @generated by protoc-gen-es v1.5.1 with parameter "target=ts"
// @generated from file autokitteh/program/v1/module.proto (package autokitteh.program.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message autokitteh.program.v1.Module
 */
export class Module extends Message<Module> {
  /**
   * @generated from field: map<string, autokitteh.program.v1.Function> functions = 1;
   */
  functions: { [key: string]: Function } = {};

  /**
   * @generated from field: map<string, autokitteh.program.v1.Variable> variables = 2;
   */
  variables: { [key: string]: Variable } = {};

  constructor(data?: PartialMessage<Module>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.program.v1.Module";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "functions", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Function} },
    { no: 2, name: "variables", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Variable} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Module {
    return new Module().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Module {
    return new Module().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Module {
    return new Module().fromJsonString(jsonString, options);
  }

  static equals(a: Module | PlainMessage<Module> | undefined, b: Module | PlainMessage<Module> | undefined): boolean {
    return proto3.util.equals(Module, a, b);
  }
}

/**
 * @generated from message autokitteh.program.v1.Variable
 */
export class Variable extends Message<Variable> {
  /**
   * @generated from field: string description = 1;
   */
  description = "";

  constructor(data?: PartialMessage<Variable>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.program.v1.Variable";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Variable {
    return new Variable().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Variable {
    return new Variable().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Variable {
    return new Variable().fromJsonString(jsonString, options);
  }

  static equals(a: Variable | PlainMessage<Variable> | undefined, b: Variable | PlainMessage<Variable> | undefined): boolean {
    return proto3.util.equals(Variable, a, b);
  }
}

/**
 * @generated from message autokitteh.program.v1.Function
 */
export class Function extends Message<Function> {
  /**
   * @generated from field: string description = 1;
   */
  description = "";

  /**
   * @generated from field: string documentation_url = 2;
   */
  documentationUrl = "";

  /**
   * @generated from field: repeated autokitteh.program.v1.FunctionField input = 3;
   */
  input: FunctionField[] = [];

  /**
   * @generated from field: repeated autokitteh.program.v1.FunctionField output = 4;
   */
  output: FunctionField[] = [];

  /**
   * @generated from field: repeated autokitteh.program.v1.Example examples = 5;
   */
  examples: Example[] = [];

  /**
   * @generated from field: string deprecated_message = 6;
   */
  deprecatedMessage = "";

  constructor(data?: PartialMessage<Function>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.program.v1.Function";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "documentation_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "input", kind: "message", T: FunctionField, repeated: true },
    { no: 4, name: "output", kind: "message", T: FunctionField, repeated: true },
    { no: 5, name: "examples", kind: "message", T: Example, repeated: true },
    { no: 6, name: "deprecated_message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Function {
    return new Function().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Function {
    return new Function().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Function {
    return new Function().fromJsonString(jsonString, options);
  }

  static equals(a: Function | PlainMessage<Function> | undefined, b: Function | PlainMessage<Function> | undefined): boolean {
    return proto3.util.equals(Function, a, b);
  }
}

/**
 * @generated from message autokitteh.program.v1.FunctionField
 */
export class FunctionField extends Message<FunctionField> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string description = 2;
   */
  description = "";

  /**
   * Flexible informative annotation, not parsed.
   *
   * @generated from field: string type = 3;
   */
  type = "";

  /**
   * @generated from field: bool optional = 4;
   */
  optional = false;

  /**
   * @generated from field: string default_value = 5;
   */
  defaultValue = "";

  /**
   * @generated from field: bool kwarg = 6;
   */
  kwarg = false;

  /**
   * @generated from field: repeated autokitteh.program.v1.Example examples = 7;
   */
  examples: Example[] = [];

  constructor(data?: PartialMessage<FunctionField>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.program.v1.FunctionField";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "optional", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "default_value", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "kwarg", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "examples", kind: "message", T: Example, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FunctionField {
    return new FunctionField().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FunctionField {
    return new FunctionField().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FunctionField {
    return new FunctionField().fromJsonString(jsonString, options);
  }

  static equals(a: FunctionField | PlainMessage<FunctionField> | undefined, b: FunctionField | PlainMessage<FunctionField> | undefined): boolean {
    return proto3.util.equals(FunctionField, a, b);
  }
}

/**
 * @generated from message autokitteh.program.v1.Example
 */
export class Example extends Message<Example> {
  /**
   * @generated from field: string code = 1;
   */
  code = "";

  /**
   * Optional.
   *
   * @generated from field: string explanation = 2;
   */
  explanation = "";

  constructor(data?: PartialMessage<Example>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.program.v1.Example";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "explanation", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Example {
    return new Example().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Example {
    return new Example().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Example {
    return new Example().fromJsonString(jsonString, options);
  }

  static equals(a: Example | PlainMessage<Example> | undefined, b: Example | PlainMessage<Example> | undefined): boolean {
    return proto3.util.equals(Example, a, b);
  }
}

