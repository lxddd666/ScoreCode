// import { Response } from './response';
import { ResponseList } from './response';

export interface SmslogStateProps {
    error: object | string | null;
    smslogList: ResponseList<SmslogListData> | null;
}

export interface SmslogListData {
    id: number;
    event: string;
    mobile: string;
    code: string;
    times: number;
    ip: string;
    status: number;
    createdAt: string;
    updatedAt: string;
}
