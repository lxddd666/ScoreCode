import { ResponseList } from './response';

export interface CronStateProps {
    error: object | string | null;
    postList: ResponseList<PostListData> | null;
}

export interface PostListData {
    id: number;
    groupId: number;
    name: string;
    params: string;
    pattern: string;
    policy: number;
    count: number;
    sort: number;
    remark: string;
    status: number;
    createdAt: string;
    updatedAt: string;
    groupName: number;
}
