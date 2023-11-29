export interface blacklistStateProps {
    List: BlackList1 | null;
    error: object | string | null;
}

export interface BlackListData {
    id: number;
    ip: string;
    remark: string;
    status: number;
    createdAt: string;
    updatedAt: string;
}
export interface BlackList1 {
    code: number;
    message: string;
    data: BlackListInfo;
    error: object;
    timestamp: number;
    traceID: string;
}

export interface BlackListInfo {
    list: BlackListData[],
    page: number,
    pageSize: number,
    pageCount: number,
    totalCount: number
}
// END API RESPONSE

// START GENERAL
// export interface NestedSubOption {
//     id: number;
//     name: string;
//     value: number;
//     expanded: boolean;
//     children: NestedSubOption[] | null; // for hierarchy view
//     label?: string;
//     [key: string]: any;
// }

// export interface FlatOption {
//     id: number; //value = id
//     name: string;
// }
// // END GENERAL
