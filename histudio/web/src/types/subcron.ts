import { ResponseList } from './response';

export interface SubcronStateProps {
    error: object | string | null;
    postList: ResponseList<PostListData> | null;
}

export interface PostListData {
    id: number;
    pid: number;
    name: string;
    isDefault: number;
    sort: number;
    remark: string;
    status: number;
    createdAt: string;
    updatedAt: string;
}
