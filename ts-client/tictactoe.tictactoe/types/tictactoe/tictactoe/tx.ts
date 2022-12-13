/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "tictactoe.tictactoe";

export interface MsgCreateGame {
  creator: string;
}

export interface MsgCreateGameResponse {
  gameId: number;
}

export interface MsgAcceptGame {
  creator: string;
  gameId: number;
}

export interface MsgAcceptGameResponse {
}

export interface MsgPerformMove {
  creator: string;
  gameId: number;
  index: number;
}

export interface MsgPerformMoveResponse {
}

function createBaseMsgCreateGame(): MsgCreateGame {
  return { creator: "" };
}

export const MsgCreateGame = {
  encode(message: MsgCreateGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateGame {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateGame();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGame {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: MsgCreateGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateGame>, I>>(object: I): MsgCreateGame {
    const message = createBaseMsgCreateGame();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseMsgCreateGameResponse(): MsgCreateGameResponse {
  return { gameId: 0 };
}

export const MsgCreateGameResponse = {
  encode(message: MsgCreateGameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.gameId !== 0) {
      writer.uint32(8).uint32(message.gameId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateGameResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateGameResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.gameId = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGameResponse {
    return { gameId: isSet(object.gameId) ? Number(object.gameId) : 0 };
  },

  toJSON(message: MsgCreateGameResponse): unknown {
    const obj: any = {};
    message.gameId !== undefined && (obj.gameId = Math.round(message.gameId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateGameResponse>, I>>(object: I): MsgCreateGameResponse {
    const message = createBaseMsgCreateGameResponse();
    message.gameId = object.gameId ?? 0;
    return message;
  },
};

function createBaseMsgAcceptGame(): MsgAcceptGame {
  return { creator: "", gameId: 0 };
}

export const MsgAcceptGame = {
  encode(message: MsgAcceptGame, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameId !== 0) {
      writer.uint32(16).uint32(message.gameId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAcceptGame {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAcceptGame();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameId = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAcceptGame {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      gameId: isSet(object.gameId) ? Number(object.gameId) : 0,
    };
  },

  toJSON(message: MsgAcceptGame): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameId !== undefined && (obj.gameId = Math.round(message.gameId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAcceptGame>, I>>(object: I): MsgAcceptGame {
    const message = createBaseMsgAcceptGame();
    message.creator = object.creator ?? "";
    message.gameId = object.gameId ?? 0;
    return message;
  },
};

function createBaseMsgAcceptGameResponse(): MsgAcceptGameResponse {
  return {};
}

export const MsgAcceptGameResponse = {
  encode(_: MsgAcceptGameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAcceptGameResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAcceptGameResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAcceptGameResponse {
    return {};
  },

  toJSON(_: MsgAcceptGameResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAcceptGameResponse>, I>>(_: I): MsgAcceptGameResponse {
    const message = createBaseMsgAcceptGameResponse();
    return message;
  },
};

function createBaseMsgPerformMove(): MsgPerformMove {
  return { creator: "", gameId: 0, index: 0 };
}

export const MsgPerformMove = {
  encode(message: MsgPerformMove, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.gameId !== 0) {
      writer.uint32(16).uint32(message.gameId);
    }
    if (message.index !== 0) {
      writer.uint32(24).uint32(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPerformMove {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPerformMove();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.gameId = reader.uint32();
          break;
        case 3:
          message.index = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPerformMove {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      gameId: isSet(object.gameId) ? Number(object.gameId) : 0,
      index: isSet(object.index) ? Number(object.index) : 0,
    };
  },

  toJSON(message: MsgPerformMove): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.gameId !== undefined && (obj.gameId = Math.round(message.gameId));
    message.index !== undefined && (obj.index = Math.round(message.index));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPerformMove>, I>>(object: I): MsgPerformMove {
    const message = createBaseMsgPerformMove();
    message.creator = object.creator ?? "";
    message.gameId = object.gameId ?? 0;
    message.index = object.index ?? 0;
    return message;
  },
};

function createBaseMsgPerformMoveResponse(): MsgPerformMoveResponse {
  return {};
}

export const MsgPerformMoveResponse = {
  encode(_: MsgPerformMoveResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPerformMoveResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPerformMoveResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgPerformMoveResponse {
    return {};
  },

  toJSON(_: MsgPerformMoveResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPerformMoveResponse>, I>>(_: I): MsgPerformMoveResponse {
    const message = createBaseMsgPerformMoveResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse>;
  AcceptGame(request: MsgAcceptGame): Promise<MsgAcceptGameResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  PerformMove(request: MsgPerformMove): Promise<MsgPerformMoveResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateGame = this.CreateGame.bind(this);
    this.AcceptGame = this.AcceptGame.bind(this);
    this.PerformMove = this.PerformMove.bind(this);
  }
  CreateGame(request: MsgCreateGame): Promise<MsgCreateGameResponse> {
    const data = MsgCreateGame.encode(request).finish();
    const promise = this.rpc.request("tictactoe.tictactoe.Msg", "CreateGame", data);
    return promise.then((data) => MsgCreateGameResponse.decode(new _m0.Reader(data)));
  }

  AcceptGame(request: MsgAcceptGame): Promise<MsgAcceptGameResponse> {
    const data = MsgAcceptGame.encode(request).finish();
    const promise = this.rpc.request("tictactoe.tictactoe.Msg", "AcceptGame", data);
    return promise.then((data) => MsgAcceptGameResponse.decode(new _m0.Reader(data)));
  }

  PerformMove(request: MsgPerformMove): Promise<MsgPerformMoveResponse> {
    const data = MsgPerformMove.encode(request).finish();
    const promise = this.rpc.request("tictactoe.tictactoe.Msg", "PerformMove", data);
    return promise.then((data) => MsgPerformMoveResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
