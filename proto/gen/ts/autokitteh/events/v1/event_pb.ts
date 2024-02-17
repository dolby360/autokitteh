// @generated by protoc-gen-es v1.5.1 with parameter "target=ts"
// @generated from file autokitteh/events/v1/event.proto (package autokitteh.events.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";
import { Value } from "../../values/v1/values_pb.js";

/**
 * @generated from enum autokitteh.events.v1.EventState
 */
export enum EventState {
  /**
   * @generated from enum value: EVENT_STATE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: EVENT_STATE_SAVED = 1;
   */
  SAVED = 1,

  /**
   * @generated from enum value: EVENT_STATE_PROCESSING = 2;
   */
  PROCESSING = 2,

  /**
   * @generated from enum value: EVENT_STATE_COMPLETED = 3;
   */
  COMPLETED = 3,

  /**
   * @generated from enum value: EVENT_STATE_FAILED = 4;
   */
  FAILED = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(EventState)
proto3.util.setEnumType(EventState, "autokitteh.events.v1.EventState", [
  { no: 0, name: "EVENT_STATE_UNSPECIFIED" },
  { no: 1, name: "EVENT_STATE_SAVED" },
  { no: 2, name: "EVENT_STATE_PROCESSING" },
  { no: 3, name: "EVENT_STATE_COMPLETED" },
  { no: 4, name: "EVENT_STATE_FAILED" },
]);

/**
 * @generated from message autokitteh.events.v1.EventRecord
 */
export class EventRecord extends Message<EventRecord> {
  /**
   * @generated from field: uint32 seq = 1;
   */
  seq = 0;

  /**
   * @generated from field: string event_id = 2;
   */
  eventId = "";

  /**
   * @generated from field: autokitteh.events.v1.EventState state = 3;
   */
  state = EventState.UNSPECIFIED;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 4;
   */
  createdAt?: Timestamp;

  constructor(data?: PartialMessage<EventRecord>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.events.v1.EventRecord";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "seq", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 2, name: "event_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "state", kind: "enum", T: proto3.getEnumType(EventState) },
    { no: 4, name: "created_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventRecord {
    return new EventRecord().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventRecord {
    return new EventRecord().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventRecord {
    return new EventRecord().fromJsonString(jsonString, options);
  }

  static equals(a: EventRecord | PlainMessage<EventRecord> | undefined, b: EventRecord | PlainMessage<EventRecord> | undefined): boolean {
    return proto3.util.equals(EventRecord, a, b);
  }
}

/**
 * @generated from message autokitteh.events.v1.Event
 */
export class Event extends Message<Event> {
  /**
   * @generated from field: string event_id = 1;
   */
  eventId = "";

  /**
   * @generated from field: string integration_id = 2;
   */
  integrationId = "";

  /**
   * TODO: think of a name that does not hint authn.
   *
   * @generated from field: string integration_token = 3;
   */
  integrationToken = "";

  /**
   * @generated from field: string original_event_id = 4;
   */
  originalEventId = "";

  /**
   * @generated from field: string event_type = 5;
   */
  eventType = "";

  /**
   * @generated from field: map<string, autokitteh.values.v1.Value> data = 6;
   */
  data: { [key: string]: Value } = {};

  /**
   * @generated from field: map<string, string> memo = 7;
   */
  memo: { [key: string]: string } = {};

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 8;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: uint64 seq = 9;
   */
  seq = protoInt64.zero;

  constructor(data?: PartialMessage<Event>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "autokitteh.events.v1.Event";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "event_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "integration_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "integration_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "original_event_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "event_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "data", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Value} },
    { no: 7, name: "memo", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 8, name: "created_at", kind: "message", T: Timestamp },
    { no: 9, name: "seq", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event {
    return new Event().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event {
    return new Event().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event {
    return new Event().fromJsonString(jsonString, options);
  }

  static equals(a: Event | PlainMessage<Event> | undefined, b: Event | PlainMessage<Event> | undefined): boolean {
    return proto3.util.equals(Event, a, b);
  }
}

