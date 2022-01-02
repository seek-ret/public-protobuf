// package: seekret.auth
// file: seekret/proto/auth/auth.proto

import * as jspb from "google-protobuf";

export class ClientCredentials extends jspb.Message {
  getClientId(): string;
  setClientId(value: string): void;

  getClientSecret(): string;
  setClientSecret(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientCredentials.AsObject;
  static toObject(includeInstance: boolean, msg: ClientCredentials): ClientCredentials.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientCredentials, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientCredentials;
  static deserializeBinaryFromReader(message: ClientCredentials, reader: jspb.BinaryReader): ClientCredentials;
}

export namespace ClientCredentials {
  export type AsObject = {
    clientId: string,
    clientSecret: string,
  }
}

export class JWTDetails extends jspb.Message {
  getAccessToken(): string;
  setAccessToken(value: string): void;

  getRefreshToken(): string;
  setRefreshToken(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JWTDetails.AsObject;
  static toObject(includeInstance: boolean, msg: JWTDetails): JWTDetails.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JWTDetails, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JWTDetails;
  static deserializeBinaryFromReader(message: JWTDetails, reader: jspb.BinaryReader): JWTDetails;
}

export namespace JWTDetails {
  export type AsObject = {
    accessToken: string,
    refreshToken: string,
  }
}

export class ClientAuthenticateRequest extends jspb.Message {
  hasClientCredentials(): boolean;
  clearClientCredentials(): void;
  getClientCredentials(): ClientCredentials | undefined;
  setClientCredentials(value?: ClientCredentials): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientAuthenticateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ClientAuthenticateRequest): ClientAuthenticateRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientAuthenticateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientAuthenticateRequest;
  static deserializeBinaryFromReader(message: ClientAuthenticateRequest, reader: jspb.BinaryReader): ClientAuthenticateRequest;
}

export namespace ClientAuthenticateRequest {
  export type AsObject = {
    clientCredentials?: ClientCredentials.AsObject,
  }
}

export class ClientAuthenticateResponse extends jspb.Message {
  hasDetails(): boolean;
  clearDetails(): void;
  getDetails(): JWTDetails | undefined;
  setDetails(value?: JWTDetails): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientAuthenticateResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ClientAuthenticateResponse): ClientAuthenticateResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientAuthenticateResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientAuthenticateResponse;
  static deserializeBinaryFromReader(message: ClientAuthenticateResponse, reader: jspb.BinaryReader): ClientAuthenticateResponse;
}

export namespace ClientAuthenticateResponse {
  export type AsObject = {
    details?: JWTDetails.AsObject,
  }
}

