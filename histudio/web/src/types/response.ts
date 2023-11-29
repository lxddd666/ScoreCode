export interface Response<T> {
    code: number;
    message: string;
    data: T | null;
    error: object;
    timestamp: number;
    traceID: string;
}

export interface ResponseList<T> {
    code: number;
    message: string;
    data: ResponseListInfo<T> | null;
    error: object;
    timestamp: number;
    traceID: string;
}

export interface ResponseListInfo<T> {
    list: T[] | null;
    page: number;
    pageSize: number;
    pageCount: number;
    totalCount: number;
}
