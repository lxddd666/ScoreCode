import { ResponseList, Response } from './response';

export interface LogStateProps {
    error: null;
    logList: ResponseList<Log> | null;
    log: Response<Log> | null;
}

export interface Log {
    id: number;
    reqId: string;
    appId: string;
    merchantId: number;
    memberId: number;
    method: string;
    module: string;
    url: string;
    getData: object;
    postData: object;
    headerData: object;
    ip: string;
    provinceId: number;
    cityId: number;
    errorCode: number;
    errorMsg: string;
    errorData: object;
    userAgent: string;
    takeUpTime: number;
    timestamp: number;
    status: number;
    createdAt: string;
    updatedAt: string;
    memberName: string;
    region: string;
    cityLabel: string;
}
