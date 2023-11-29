import { GenericCardProps } from 'types';
import { ResponseList } from './response';
import { PostDataType, UserProfile, Profile } from 'types/user-profile';

export interface FollowerCardProps {
    avatar: string;
    follow: number;
    location: string;
    name: string;
}

export interface FriendRequestCardProps extends Profile {
    mutual: number;
}

export interface FriendsCardProps {
    avatar: string;
    location: string;
    name: string;
}

export interface UserProfileCardProps extends UserProfile {
    profile: string;
}

export interface UserSimpleCardProps {
    avatar: string;
    name: string;
    status: string;
}

export interface UserStateProps {
    usersS1: UserProfile[];
    usersS2: UserProfileStyle2[];
    followers: FollowerCardProps[];
    friendRequests: FriendRequestCardProps[];
    friends: FriendsCardProps[];
    gallery: GenericCardProps[];
    posts: PostDataType[];
    detailCards: UserProfile[];
    simpleCards: UserSimpleCardProps[];
    profileCards: UserProfileCardProps[];
    userList: ResponseList<UserListData> | null;
    deptList: ResponseList<DeptListData> | null;
    postList: ResponseList<PostListData> | null;
    roleList: ResponseList<RoleListData> | null;
    userOnlineList: ResponseList<UserOnlineListData> | null;
    messageList: ResponseList<MessageListData> | null;
    error: object | string | null;
    adminInfo: { [key: string]: any };
}

export type UserProfileStyle2 = {
    image: string;
    name: string;
    designation: string;
    badgeStatus: string;
    subContent: string;
    email: string;
    phone: string;
    location: string;
    progressValue: string;
};

export interface credentialInfo {
    username: string;
    password: string;
    code: string;
    cid: string;
}

export interface UserListData {
    id: number;
    deptId: number;
    roleId: number;
    realName: string;
    username: string;
    passwordHash: string;
    salt: string;
    passwordResetToken: string;
    integral: number;
    balance: number;
    avatar: string;
    sex: number;
    qq: string;
    email: string;
    mobile: string;
    birthday: string;
    cityId: number;
    address: string;
    pid: number;
    level: number;
    tree: string;
    inviteCode: string;
    cash: {};
    lastActiveAt: string;
    remark: string;
    status: number;
    createdAt: string;
    updatedAt: string;
    deptName: string;
    roleName: string;
    postIds: [number];
}

export interface DeptListData {
    id: number;
    orgId: number;
    pid: number;
    name: string;
    code: string;
    type: string;
    leader: string;
    phone: string;
    email: string;
    level: number;
    tree: string;
    sort: number;
    status: number;
    createdAt: string;
    updatedAt: string;
    label: string;
    value: number;
    children: DeptListData[] | null;
}

export interface PostListData {
    id: number;
    code: string;
    name: string;
    remark: string;
    sort: number;
    status: number;
    createdAt: string;
    updatedAt: string;
    orgId: number;
}

export interface UserOnlineListData {
    id: string;
    ip: string;
    os: string;
    browser: string;
    firstTime: number;
    heartbeatTime: number;
    app: string;
    userId: number;
    username: string;
    avatar: string;
}

export interface MessageListData {
    id: number;
    type: number;
    title: string;
    content: string;
    tag: number;
    sort: number;
    createdBy: number;
    createdAt: string;
    senderAvatar: string;
    isRead: boolean;
}

export interface RoleListData {
    id: number;
    name: string;
    key: string;
    dataScope: number;
    customDept: {};
    pid: number;
    level: number;
    tree: string;
    remark: string;
    sort: number;
    status: number;
    createdAt: string;
    updatedAt: string;
    orgAdmin: number;
    label: string;
    value: number;
    children: RoleListData[] | null;
}

export interface MenuListData {
    id: number;
    name: string;
    key: number;
    dataScope: number;
    customDept: {};
    pid: number;
    level: number;
    tree: string;
    remark: string;
    sort: number;
    status: number;
    createdAt: string;
    updatedAt: string;
    orgAdmin: number;
    label: string;
    value: number;
    children: RoleListData[] | null;
}

export interface DataScopeListData {
    type: string;
    label: string;
    key: number;
    value: number;
    children:
        | [
              {
                  label: string;
                  value: number;
              }
          ]
        | null;
}
