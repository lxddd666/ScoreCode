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
  message: string;
  isRead: boolean;
  date: string;
  last?: TMessage;
}

export interface TMessage {
  sendMsg: string;
  reqId: number;
  sendTime: string;
  tgId: number;
  chatId: number;
  msgType: number;
  read: number;
  out: number;
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
  message: '',
  isRead: false,
  date: '',
};


export function newState(state: TChatItemParam | null): TChatItemParam {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}


