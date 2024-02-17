// @generated by protoc-gen-es v1.5.1 with parameter "target=ts"
// @generated from file autokitteh/sessions/v1/session.proto (package autokitteh.sessions.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Value } from "../../values/v1/values_pb.js";
import { CodeLocation, Error } from "../../program/v1/program_pb.js";

/**
 * TODO: Type might not be the best qualifier.
 *
 * @generated from enum autokitteh.sessions.v1.SessionStateType
 */
export enum SessionStateType {
  /**
   * @generated from enum value: SESSION_STATE_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SESSION_STATE_TYPE_CREATED = 1;
   */
  CREATED = 1,

  /**
   * @generated from enum value: SESSION_STATE_TYPE_RUNNING = 2;
   */
  RUNNING = 2,

  /**
   * @generated from enum value: SESSION_STATE_TYPE_ERROR = 3;
   */
  ERROR = 3,

  /**
   * @generated from enum value: SESSION_STATE_TYPE_COMPLETED = 4;
   */
  COMPLETED = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(SessionStateType)
proto3.util.setEnumType(SessionStateType, "autokitteh.sessions.v1.SessionStateType", [
  { no: 0, name: "SESSION_STATE_TYPE_UNSPECIFIED" },
  { no: 1, name: "SESSION_STATE_TYPE_CREATED" },
  { no: 2, name: "SESSION_STATE_TYPE_RUNNING" },
  { no: 3, name: "SESSION_STATE_TYPE_ERROR" },
  { no: 4, name: "SESSION_STATE_TYPE_COMPLETED" },
]);

/**
 * @generated from message autokitteh.sessions.v1.SessionState
 */
export class SessionState extends Message<SessionState> {
  /**
   * @generated from field: google.protobuf.Timestamp t = 1;
   */
  t?: Timestamp;

  /**
   * @generated from oneof autokitteh.sessions.v1.SessionState.states
   */
  states: {
    /**
     * @generated from field: autokitteh.sessions.v1.SessionState.Created created = 10;
     */
    value: SessionState_Created;
    case: "created";
  } | {
    /**
     * @generated from field: autokitteh.sessions.v1.SessionState.Running running = 11;
     */
    value: SessionState_Running;
    case: "running";
  } | {
    /**
     * @generated from field: autokitteh.sessions.v1.SessionState.Error error = 12;
     */
    value: SessionState_Error;
    case: "error";
  } | {
    /**
     * @generated from field: autokitteh.sessions.v1.SessionState.Completed completed = 13;
     */
    value: SessionState_Completed;
    case: "completed";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<SessionState>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.SessionState";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "t", kind: "message", T: Timestamp },
    { no: 10, name: "created", kind: "message", T: SessionState_Created, oneof: "states" },
    { no: 11, name: "running", kind: "message", T: SessionState_Running, oneof: "states" },
    { no: 12, name: "error", kind: "message", T: SessionState_Error, oneof: "states" },
    { no: 13, name: "completed", kind: "message", T: SessionState_Completed, oneof: "states" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionState {
    return new SessionState().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionState {
    return new SessionState().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionState {
    return new SessionState().fromJsonString(jsonString, options);
  }

  static equals(a: SessionState | PlainMessage<SessionState> | undefined, b: SessionState | PlainMessage<SessionState> | undefined): boolean {
    return proto3.util.equals(SessionState, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.SessionState.Created
 */
export class SessionState_Created extends Message<SessionState_Created> {
  constructor(data?: PartialMessage<SessionState_Created>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.SessionState.Created";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionState_Created {
    return new SessionState_Created().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionState_Created {
    return new SessionState_Created().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionState_Created {
    return new SessionState_Created().fromJsonString(jsonString, options);
  }

  static equals(a: SessionState_Created | PlainMessage<SessionState_Created> | undefined, b: SessionState_Created | PlainMessage<SessionState_Created> | undefined): boolean {
    return proto3.util.equals(SessionState_Created, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.SessionState.Running
 */
export class SessionState_Running extends Message<SessionState_Running> {
  /**
   * @generated from field: string run_id = 1;
   */
  runId = "";

  /**
   * present if run is a Call.
   *
   * @generated from field: autokitteh.values.v1.Value call = 2;
   */
  call?: Value;

  constructor(data?: PartialMessage<SessionState_Running>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.SessionState.Running";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "run_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "call", kind: "message", T: Value },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionState_Running {
    return new SessionState_Running().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionState_Running {
    return new SessionState_Running().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionState_Running {
    return new SessionState_Running().fromJsonString(jsonString, options);
  }

  static equals(a: SessionState_Running | PlainMessage<SessionState_Running> | undefined, b: SessionState_Running | PlainMessage<SessionState_Running> | undefined): boolean {
    return proto3.util.equals(SessionState_Running, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.SessionState.Error
 */
export class SessionState_Error extends Message<SessionState_Error> {
  /**
   * @generated from field: repeated string prints = 1;
   */
  prints: string[] = [];

  /**
   * @generated from field: autokitteh.program.v1.Error error = 2;
   */
  error?: Error;

  constructor(data?: PartialMessage<SessionState_Error>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.SessionState.Error";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "prints", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "error", kind: "message", T: Error },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionState_Error {
    return new SessionState_Error().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionState_Error {
    return new SessionState_Error().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionState_Error {
    return new SessionState_Error().fromJsonString(jsonString, options);
  }

  static equals(a: SessionState_Error | PlainMessage<SessionState_Error> | undefined, b: SessionState_Error | PlainMessage<SessionState_Error> | undefined): boolean {
    return proto3.util.equals(SessionState_Error, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.SessionState.Completed
 */
export class SessionState_Completed extends Message<SessionState_Completed> {
  /**
   * @generated from field: repeated string prints = 1;
   */
  prints: string[] = [];

  /**
   * @generated from field: map<string, autokitteh.values.v1.Value> exports = 2;
   */
  exports: { [key: string]: Value } = {};

  /**
   * @generated from field: autokitteh.values.v1.Value return_value = 3;
   */
  returnValue?: Value;

  constructor(data?: PartialMessage<SessionState_Completed>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.SessionState.Completed";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "prints", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "exports", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Value} },
    { no: 3, name: "return_value", kind: "message", T: Value },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionState_Completed {
    return new SessionState_Completed().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionState_Completed {
    return new SessionState_Completed().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionState_Completed {
    return new SessionState_Completed().fromJsonString(jsonString, options);
  }

  static equals(a: SessionState_Completed | PlainMessage<SessionState_Completed> | undefined, b: SessionState_Completed | PlainMessage<SessionState_Completed> | undefined): boolean {
    return proto3.util.equals(SessionState_Completed, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.Call
 */
export class Call extends Message<Call> {
  /**
   * @generated from field: autokitteh.sessions.v1.Call.Spec spec = 1;
   */
  spec?: Call_Spec;

  /**
   * @generated from field: repeated autokitteh.sessions.v1.Call.Attempt attempts = 2;
   */
  attempts: Call_Attempt[] = [];

  constructor(data?: PartialMessage<Call>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.Call";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "spec", kind: "message", T: Call_Spec },
    { no: 2, name: "attempts", kind: "message", T: Call_Attempt, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Call {
    return new Call().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Call {
    return new Call().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Call {
    return new Call().fromJsonString(jsonString, options);
  }

  static equals(a: Call | PlainMessage<Call> | undefined, b: Call | PlainMessage<Call> | undefined): boolean {
    return proto3.util.equals(Call, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.Call.Spec
 */
export class Call_Spec extends Message<Call_Spec> {
  /**
   * @generated from field: autokitteh.values.v1.Value function = 1;
   */
  function?: Value;

  /**
   * @generated from field: repeated autokitteh.values.v1.Value args = 2;
   */
  args: Value[] = [];

  /**
   * @generated from field: map<string, autokitteh.values.v1.Value> kwargs = 3;
   */
  kwargs: { [key: string]: Value } = {};

  /**
   * @generated from field: uint32 seq = 4;
   */
  seq = 0;

  constructor(data?: PartialMessage<Call_Spec>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.Call.Spec";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "function", kind: "message", T: Value },
    { no: 2, name: "args", kind: "message", T: Value, repeated: true },
    { no: 3, name: "kwargs", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Value} },
    { no: 4, name: "seq", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Call_Spec {
    return new Call_Spec().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Call_Spec {
    return new Call_Spec().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Call_Spec {
    return new Call_Spec().fromJsonString(jsonString, options);
  }

  static equals(a: Call_Spec | PlainMessage<Call_Spec> | undefined, b: Call_Spec | PlainMessage<Call_Spec> | undefined): boolean {
    return proto3.util.equals(Call_Spec, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.Call.Attempt
 */
export class Call_Attempt extends Message<Call_Attempt> {
  /**
   * @generated from field: autokitteh.sessions.v1.Call.Attempt.Start start = 1;
   */
  start?: Call_Attempt_Start;

  /**
   * @generated from field: autokitteh.sessions.v1.Call.Attempt.Complete complete = 2;
   */
  complete?: Call_Attempt_Complete;

  constructor(data?: PartialMessage<Call_Attempt>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.Call.Attempt";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "start", kind: "message", T: Call_Attempt_Start },
    { no: 2, name: "complete", kind: "message", T: Call_Attempt_Complete },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Call_Attempt {
    return new Call_Attempt().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Call_Attempt {
    return new Call_Attempt().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Call_Attempt {
    return new Call_Attempt().fromJsonString(jsonString, options);
  }

  static equals(a: Call_Attempt | PlainMessage<Call_Attempt> | undefined, b: Call_Attempt | PlainMessage<Call_Attempt> | undefined): boolean {
    return proto3.util.equals(Call_Attempt, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.Call.Attempt.Result
 */
export class Call_Attempt_Result extends Message<Call_Attempt_Result> {
  /**
   * @generated from oneof autokitteh.sessions.v1.Call.Attempt.Result.result
   */
  result: {
    /**
     * @generated from field: autokitteh.values.v1.Value value = 10;
     */
    value: Value;
    case: "value";
  } | {
    /**
     * @generated from field: autokitteh.program.v1.Error error = 11;
     */
    value: Error;
    case: "error";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Call_Attempt_Result>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.Call.Attempt.Result";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 10, name: "value", kind: "message", T: Value, oneof: "result" },
    { no: 11, name: "error", kind: "message", T: Error, oneof: "result" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Call_Attempt_Result {
    return new Call_Attempt_Result().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Call_Attempt_Result {
    return new Call_Attempt_Result().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Call_Attempt_Result {
    return new Call_Attempt_Result().fromJsonString(jsonString, options);
  }

  static equals(a: Call_Attempt_Result | PlainMessage<Call_Attempt_Result> | undefined, b: Call_Attempt_Result | PlainMessage<Call_Attempt_Result> | undefined): boolean {
    return proto3.util.equals(Call_Attempt_Result, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.Call.Attempt.Start
 */
export class Call_Attempt_Start extends Message<Call_Attempt_Start> {
  /**
   * @generated from field: google.protobuf.Timestamp started_at = 1;
   */
  startedAt?: Timestamp;

  /**
   * @generated from field: uint32 num = 5;
   */
  num = 0;

  constructor(data?: PartialMessage<Call_Attempt_Start>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.Call.Attempt.Start";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "started_at", kind: "message", T: Timestamp },
    { no: 5, name: "num", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Call_Attempt_Start {
    return new Call_Attempt_Start().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Call_Attempt_Start {
    return new Call_Attempt_Start().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Call_Attempt_Start {
    return new Call_Attempt_Start().fromJsonString(jsonString, options);
  }

  static equals(a: Call_Attempt_Start | PlainMessage<Call_Attempt_Start> | undefined, b: Call_Attempt_Start | PlainMessage<Call_Attempt_Start> | undefined): boolean {
    return proto3.util.equals(Call_Attempt_Start, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.Call.Attempt.Complete
 */
export class Call_Attempt_Complete extends Message<Call_Attempt_Complete> {
  /**
   * @generated from field: google.protobuf.Timestamp completed_at = 1;
   */
  completedAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Duration retry_interval = 2;
   */
  retryInterval?: Duration;

  /**
   * @generated from field: bool is_last = 3;
   */
  isLast = false;

  /**
   * @generated from field: autokitteh.sessions.v1.Call.Attempt.Result result = 4;
   */
  result?: Call_Attempt_Result;

  constructor(data?: PartialMessage<Call_Attempt_Complete>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.Call.Attempt.Complete";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "completed_at", kind: "message", T: Timestamp },
    { no: 2, name: "retry_interval", kind: "message", T: Duration },
    { no: 3, name: "is_last", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "result", kind: "message", T: Call_Attempt_Result },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Call_Attempt_Complete {
    return new Call_Attempt_Complete().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Call_Attempt_Complete {
    return new Call_Attempt_Complete().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Call_Attempt_Complete {
    return new Call_Attempt_Complete().fromJsonString(jsonString, options);
  }

  static equals(a: Call_Attempt_Complete | PlainMessage<Call_Attempt_Complete> | undefined, b: Call_Attempt_Complete | PlainMessage<Call_Attempt_Complete> | undefined): boolean {
    return proto3.util.equals(Call_Attempt_Complete, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.SessionLogRecord
 */
export class SessionLogRecord extends Message<SessionLogRecord> {
  /**
   * @generated from field: google.protobuf.Timestamp t = 1;
   */
  t?: Timestamp;

  /**
   * @generated from oneof autokitteh.sessions.v1.SessionLogRecord.data
   */
  data: {
    /**
     * @generated from field: string print = 10;
     */
    value: string;
    case: "print";
  } | {
    /**
     * @generated from field: autokitteh.sessions.v1.Call.Spec call_spec = 11;
     */
    value: Call_Spec;
    case: "callSpec";
  } | {
    /**
     * @generated from field: autokitteh.sessions.v1.Call.Attempt.Start call_attempt_start = 12;
     */
    value: Call_Attempt_Start;
    case: "callAttemptStart";
  } | {
    /**
     * @generated from field: autokitteh.sessions.v1.Call.Attempt.Complete call_attempt_complete = 13;
     */
    value: Call_Attempt_Complete;
    case: "callAttemptComplete";
  } | {
    /**
     * @generated from field: autokitteh.sessions.v1.SessionState state = 14;
     */
    value: SessionState;
    case: "state";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<SessionLogRecord>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.SessionLogRecord";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "t", kind: "message", T: Timestamp },
    { no: 10, name: "print", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "data" },
    { no: 11, name: "call_spec", kind: "message", T: Call_Spec, oneof: "data" },
    { no: 12, name: "call_attempt_start", kind: "message", T: Call_Attempt_Start, oneof: "data" },
    { no: 13, name: "call_attempt_complete", kind: "message", T: Call_Attempt_Complete, oneof: "data" },
    { no: 14, name: "state", kind: "message", T: SessionState, oneof: "data" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionLogRecord {
    return new SessionLogRecord().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionLogRecord {
    return new SessionLogRecord().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionLogRecord {
    return new SessionLogRecord().fromJsonString(jsonString, options);
  }

  static equals(a: SessionLogRecord | PlainMessage<SessionLogRecord> | undefined, b: SessionLogRecord | PlainMessage<SessionLogRecord> | undefined): boolean {
    return proto3.util.equals(SessionLogRecord, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.SessionLog
 */
export class SessionLog extends Message<SessionLog> {
  /**
   * Chronological order: the last item is the latest.
   *
   * @generated from field: repeated autokitteh.sessions.v1.SessionLogRecord records = 1;
   */
  records: SessionLogRecord[] = [];

  constructor(data?: PartialMessage<SessionLog>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.SessionLog";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "records", kind: "message", T: SessionLogRecord, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionLog {
    return new SessionLog().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionLog {
    return new SessionLog().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionLog {
    return new SessionLog().fromJsonString(jsonString, options);
  }

  static equals(a: SessionLog | PlainMessage<SessionLog> | undefined, b: SessionLog | PlainMessage<SessionLog> | undefined): boolean {
    return proto3.util.equals(SessionLog, a, b);
  }
}

/**
 * @generated from message autokitteh.sessions.v1.Session
 */
export class Session extends Message<Session> {
  /**
   * @generated from field: string session_id = 1;
   */
  sessionId = "";

  /**
   * @generated from field: string deployment_id = 2;
   */
  deploymentId = "";

  /**
   * TODO: eliminate? move to inputs? or memo?
   *
   * @generated from field: string event_id = 3;
   */
  eventId = "";

  /**
   * @generated from field: autokitteh.program.v1.CodeLocation entrypoint = 4;
   */
  entrypoint?: CodeLocation;

  /**
   * @generated from field: map<string, autokitteh.values.v1.Value> inputs = 5;
   */
  inputs: { [key: string]: Value } = {};

  /**
   * @generated from field: string parent_session_id = 6;
   */
  parentSessionId = "";

  /**
   * @generated from field: map<string, string> memo = 7;
   */
  memo: { [key: string]: string } = {};

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 10;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 11;
   */
  updatedAt?: Timestamp;

  /**
   * @generated from field: autokitteh.sessions.v1.SessionStateType state = 12;
   */
  state = SessionStateType.UNSPECIFIED;

  constructor(data?: PartialMessage<Session>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.sessions.v1.Session";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "deployment_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "event_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "entrypoint", kind: "message", T: CodeLocation },
    { no: 5, name: "inputs", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Value} },
    { no: 6, name: "parent_session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "memo", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 10, name: "created_at", kind: "message", T: Timestamp },
    { no: 11, name: "updated_at", kind: "message", T: Timestamp },
    { no: 12, name: "state", kind: "enum", T: proto3.getEnumType(SessionStateType) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Session {
    return new Session().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Session {
    return new Session().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Session {
    return new Session().fromJsonString(jsonString, options);
  }

  static equals(a: Session | PlainMessage<Session> | undefined, b: Session | PlainMessage<Session> | undefined): boolean {
    return proto3.util.equals(Session, a, b);
  }
}

