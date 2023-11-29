// import { Response } from './response';
import { ResponseList } from './response';

export interface ServelogStateProps {
    error: object | string | null;
    servelogList: ResponseList<ServelogListData> | null;
}

export interface ServelogListData {
    id: number;
    traceId: string;
    levelFormat: string;
    content: string;
    stack: string;
    line: string;
    triggerNs: number;
    createdAt: string;
    updatedAt: string;
    sysLogId: number;
}
