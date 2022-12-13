/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "tictactoe.tictactoe.game";

export enum GameStatus {
  OPEN = 0,
  IN_PROGRESS = 1,
  COMPLETE = 2,
  UNRECOGNIZED = -1,
}

export function gameStatusFromJSON(object: any): GameStatus {
  switch (object) {
    case 0:
    case "OPEN":
      return GameStatus.OPEN;
    case 1:
    case "IN_PROGRESS":
      return GameStatus.IN_PROGRESS;
    case 2:
    case "COMPLETE":
      return GameStatus.COMPLETE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return GameStatus.UNRECOGNIZED;
  }
}

export function gameStatusToJSON(object: GameStatus): string {
  switch (object) {
    case GameStatus.OPEN:
      return "OPEN";
    case GameStatus.IN_PROGRESS:
      return "IN_PROGRESS";
    case GameStatus.COMPLETE:
      return "COMPLETE";
    case GameStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Game {
  id: number;
  status: GameStatus;
  challenger: string;
  opponent: string;
  playerX: string;
  playerO: string;
}

function createBaseGame(): Game {
  return { id: 0, status: 0, challenger: "", opponent: "", playerX: "", playerO: "" };
}

export const Game = {
  encode(message: Game, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint32(message.id);
    }
    if (message.status !== 0) {
      writer.uint32(16).int32(message.status);
    }
    if (message.challenger !== "") {
      writer.uint32(26).string(message.challenger);
    }
    if (message.opponent !== "") {
      writer.uint32(34).string(message.opponent);
    }
    if (message.playerX !== "") {
      writer.uint32(42).string(message.playerX);
    }
    if (message.playerO !== "") {
      writer.uint32(50).string(message.playerO);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Game {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGame();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.uint32();
          break;
        case 2:
          message.status = reader.int32() as any;
          break;
        case 3:
          message.challenger = reader.string();
          break;
        case 4:
          message.opponent = reader.string();
          break;
        case 5:
          message.playerX = reader.string();
          break;
        case 6:
          message.playerO = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Game {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      status: isSet(object.status) ? gameStatusFromJSON(object.status) : 0,
      challenger: isSet(object.challenger) ? String(object.challenger) : "",
      opponent: isSet(object.opponent) ? String(object.opponent) : "",
      playerX: isSet(object.playerX) ? String(object.playerX) : "",
      playerO: isSet(object.playerO) ? String(object.playerO) : "",
    };
  },

  toJSON(message: Game): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.status !== undefined && (obj.status = gameStatusToJSON(message.status));
    message.challenger !== undefined && (obj.challenger = message.challenger);
    message.opponent !== undefined && (obj.opponent = message.opponent);
    message.playerX !== undefined && (obj.playerX = message.playerX);
    message.playerO !== undefined && (obj.playerO = message.playerO);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Game>, I>>(object: I): Game {
    const message = createBaseGame();
    message.id = object.id ?? 0;
    message.status = object.status ?? 0;
    message.challenger = object.challenger ?? "";
    message.opponent = object.opponent ?? "";
    message.playerX = object.playerX ?? "";
    message.playerO = object.playerO ?? "";
    return message;
  },
};

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
