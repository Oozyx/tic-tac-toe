import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgPerformMove } from "./types/tictactoe/tictactoe/tx";
import { MsgCreateGame } from "./types/tictactoe/tictactoe/tx";
import { MsgAcceptGame } from "./types/tictactoe/tictactoe/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/tictactoe.tictactoe.MsgPerformMove", MsgPerformMove],
    ["/tictactoe.tictactoe.MsgCreateGame", MsgCreateGame],
    ["/tictactoe.tictactoe.MsgAcceptGame", MsgAcceptGame],
    
];

export { msgTypes }