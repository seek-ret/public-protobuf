// package: seekret.log_options
// file: seekret/proto/options.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_descriptor_pb from "google-protobuf/google/protobuf/descriptor_pb";

export class FieldOptions extends jspb.Message {
  hasHideFromLog(): boolean;
  clearHideFromLog(): void;
  getHideFromLog(): boolean | undefined;
  setHideFromLog(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FieldOptions.AsObject;
  static toObject(includeInstance: boolean, msg: FieldOptions): FieldOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: FieldOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FieldOptions;
  static deserializeBinaryFromReader(message: FieldOptions, reader: jspb.BinaryReader): FieldOptions;
}

export namespace FieldOptions {
  export type AsObject = {
    hideFromLog?: boolean,
  }
}

export class MessageOptions extends jspb.Message {
  hasHideFromLog(): boolean;
  clearHideFromLog(): void;
  getHideFromLog(): boolean | undefined;
  setHideFromLog(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageOptions.AsObject;
  static toObject(includeInstance: boolean, msg: MessageOptions): MessageOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MessageOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageOptions;
  static deserializeBinaryFromReader(message: MessageOptions, reader: jspb.BinaryReader): MessageOptions;
}

export namespace MessageOptions {
  export type AsObject = {
    hideFromLog?: boolean,
  }
}

export class MethodOptions extends jspb.Message {
  hasLogPayload(): boolean;
  clearLogPayload(): void;
  getLogPayload(): LogPayloadOptionMap[keyof LogPayloadOptionMap] | undefined;
  setLogPayload(value: LogPayloadOptionMap[keyof LogPayloadOptionMap]): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MethodOptions.AsObject;
  static toObject(includeInstance: boolean, msg: MethodOptions): MethodOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MethodOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MethodOptions;
  static deserializeBinaryFromReader(message: MethodOptions, reader: jspb.BinaryReader): MethodOptions;
}

export namespace MethodOptions {
  export type AsObject = {
    logPayload?: LogPayloadOptionMap[keyof LogPayloadOptionMap],
  }
}

  export const seekField: jspb.ExtensionFieldInfo<FieldOptions>;

  export const seekMessage: jspb.ExtensionFieldInfo<MessageOptions>;

  export const seekMethod: jspb.ExtensionFieldInfo<MethodOptions>;

export interface LogPayloadOptionMap {
  NONE: 0;
  REQUEST: 1;
  RESPONSE: 2;
  BOTH: 3;
}

export const LogPayloadOption: LogPayloadOptionMap;

