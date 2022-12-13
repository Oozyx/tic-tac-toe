import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAcceptGame } from "./types/tictactoe/tictactoe/tx";
import { MsgCreateGame } from "./types/tictactoe/tictactoe/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/tictactoe.tictactoe.MsgAcceptGame", MsgAcceptGame],
    ["/tictactoe.tictactoe.MsgCreateGame", MsgCreateGame],
    
];

export { msgTypes }