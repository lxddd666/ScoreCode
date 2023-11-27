import {cloneDeep} from "lodash-es";


export interface TChatItemParam {
    id: number;
    tgId: number;
    username: string;
    phone: number;
    type: number;
    avatar: string;
    firstName: string;
    lastName: string;
    title: string;
    message: string;
    date: number;
    last?: TMessage;
    topMessage: number;
    unreadCount: number;
    readInboxMaxID: number;
    readOutboxMaxID: number;
    msgList: TMessage[];

}

export interface TMessage {
    message: string;
    msgId: number;
    date: number;
    tgId: number;
    chatId: number;
    out: boolean;
}

export const defaultState = {
    id: 0,
    tgId: 0,
    username: '',
    phone: 0,
    type: 1,
    avatar: '',
    firstName: '',
    lastName: '',
    title: '',
    message: '',
    date: 0,
    topMessage: 0,
    unreadCount: 0,
    readInboxMaxID: 0,
    readOutboxMaxID: 0,
    msgList: [],
};


export function newState(state: TChatItemParam | null): TChatItemParam {
    if (state !== null) {
        return cloneDeep(state);
    }
    return cloneDeep(defaultState);
}


