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

