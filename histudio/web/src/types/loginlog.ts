import { Response } from './response';
import { ResponseList } from './response';

export interface LoginlogStateProps {
    error: object | string | null;
    loginlogList: ResponseList<LoginlogListData> | null;
    loginlogview: Response<LoginlogViewData> | null | undefined;
}

export interface LoginlogListData {
    id: number;
    reqId: string;
    memberId: number;
    username: string;
    loginAt: string;
    errMsg: string;
    status: number;
    createdAt: string;
    updatedAt: string;
    sysLogId: number;
    sysLogIp: string;
    sysLogProvinceId: number;
    sysLogCityId: number;
    sysLogErrorCode: number;
    sysLogUserAgent: string;
    cityLabel: string;
    os: string;
    browser: string;
}

export interface LoginlogViewData {
    id: number;
    reqId: string;
    memberId: number;
    username: string;
    response: object;
    loginAt: string;
    loginIp: string;
    errMsg: string;
    status: number;
    createdAt: string;
    updatedAt: string;
}
