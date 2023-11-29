import { ResponseList } from './response';

export interface WhatsStateProps {
    whatsAccountList: ResponseList<WhatsAccountListData> | null;
    error: object | string | null;
}

export interface WhatsAccountListData {
    id: number;
    account: string;
    nickName: string;
    avatar: string;
    accountStatus: number;
    isOnline: number;
    proxyAddress: string;
    lastLoginTime: string;
    comment: string;
    createdAt: string;
    updatedAt: string;
}
