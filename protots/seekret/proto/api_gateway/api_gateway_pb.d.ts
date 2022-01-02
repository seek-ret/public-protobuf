// package: seekret.api_gateway
// file: seekret/proto/api_gateway/api_gateway.proto

import * as jspb from "google-protobuf";

export class WorkspaceEntry extends jspb.Message {
  getBaseDir(): string;
  setBaseDir(value: string): void;

  getBucketName(): string;
  setBucketName(value: string): void;

  getEndpoint(): string;
  setEndpoint(value: string): void;

  getAccessKey(): string;
  setAccessKey(value: string): void;

  getSecretKey(): string;
  setSecretKey(value: string): void;

  getInsecure(): boolean;
  setInsecure(value: boolean): void;

  clearHostWhitelistList(): void;
  getHostWhitelistList(): Array<string>;
  setHostWhitelistList(value: Array<string>): void;
  addHostWhitelist(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WorkspaceEntry.AsObject;
  static toObject(includeInstance: boolean, msg: WorkspaceEntry): WorkspaceEntry.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WorkspaceEntry, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WorkspaceEntry;
  static deserializeBinaryFromReader(message: WorkspaceEntry, reader: jspb.BinaryReader): WorkspaceEntry;
}

export namespace WorkspaceEntry {
  export type AsObject = {
    baseDir: string,
    bucketName: string,
    endpoint: string,
    accessKey: string,
    secretKey: string,
    insecure: boolean,
    hostWhitelistList: Array<string>,
  }
}

export class FetchWorkspaceConfigurationRequest extends jspb.Message {
  getWorkspace(): string;
  setWorkspace(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FetchWorkspaceConfigurationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: FetchWorkspaceConfigurationRequest): FetchWorkspaceConfigurationRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: FetchWorkspaceConfigurationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FetchWorkspaceConfigurationRequest;
  static deserializeBinaryFromReader(message: FetchWorkspaceConfigurationRequest, reader: jspb.BinaryReader): FetchWorkspaceConfigurationRequest;
}

export namespace FetchWorkspaceConfigurationRequest {
  export type AsObject = {
    workspace: string,
  }
}

export class FetchWorkspaceConfigurationResponse extends jspb.Message {
  hasEntry(): boolean;
  clearEntry(): void;
  getEntry(): WorkspaceEntry | undefined;
  setEntry(value?: WorkspaceEntry): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FetchWorkspaceConfigurationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: FetchWorkspaceConfigurationResponse): FetchWorkspaceConfigurationResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: FetchWorkspaceConfigurationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FetchWorkspaceConfigurationResponse;
  static deserializeBinaryFromReader(message: FetchWorkspaceConfigurationResponse, reader: jspb.BinaryReader): FetchWorkspaceConfigurationResponse;
}

export namespace FetchWorkspaceConfigurationResponse {
  export type AsObject = {
    entry?: WorkspaceEntry.AsObject,
  }
}

