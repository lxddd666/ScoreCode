// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';

// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['tg'] = {
    error: null,
    tgUserList: [],
    tgBatchExecutionTaskList: [],
    tgFoldersList: [],
    tgIncreaseFansCronList: [],
    tgKeepTaskList: [],
    tgContactsList: [],
    tgMsgList: [],

    // 聊天室 会话分组
    tgArtsFolders: [],
    tgFoldersMessageList:[],
    tgFoldersMeeageHistoryList:[]
};

const slice = createSlice({
    name: 'tg',
    initialState,
    reducers: {
        hasError(state, action) {
            state.error = action.payload;
        },
        // HAS ERROR
        emitTgUserList(state, action) {
            state.tgUserList = action.payload;
        },
        // 批量操作 列表
        emitTatchExecutionTaskList(state, action) {
            state.tgBatchExecutionTaskList = action.payload;
        },
        // 账号分组 列表
        emitTgFoldersList(state, action) {
            state.tgFoldersList = action.payload;
        },
        // 长粉任务 列表
        emitTgIncreaseFansCronList(state, action) {
            state.tgIncreaseFansCronList = action.payload;
        },
        // 养号任务 列表
        emitTgKeepTaskList(state, action) {
            state.tgKeepTaskList = action.payload;
        },
        // 联系人管理 列表
        emitTgContactsList(state, action) {
            state.tgContactsList = action.payload;
        },
        // 消息记录 列表
        emitTgMsgList(state, action) {
            state.tgMsgList = action.payload;
        },
        // 聊天室 会话分组
        emitTgArtsFoldersList(state, action) {
            state.tgArtsFolders = action.payload;
        },
        // 聊天室 会话分组/消息队列
        emitTgFoldersMessageList(state, action) {
            state.tgFoldersMessageList = action.payload;
        },
        // 聊天室 会话分组/消息队列/聊天历史
        emitTgFoldersMeeageHistoryList(state, action) {
            state.tgFoldersMeeageHistoryList = action.payload;
        }
    }
});

// Reducer
export default slice.reducer;

// 异步请求提交----------------------------------------------------------------------
// tgUser 账号请求
export function getTgUserListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/tg/tgUser/list`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgUserList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
// tg 批量操作列表请求
export function getTgBatchExecutionTaskListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/tg/tgBatchExecutionTask/list`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTatchExecutionTaskList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
// tg 账号分组 列表请求
export function getTgFoldersListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/tg/tgFolders/list`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgFoldersList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
// tg 长粉任务 列表请求
export function getTgIncreaseFansCronListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/tg/tgIncreaseFansCron/list`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgIncreaseFansCronList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
// tg 养号任务 列表请求
export function getTgKeepTaskListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/tg/tgKeepTask/list`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgKeepTaskList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
// tg 联系人管理 列表请求
export function getTgContactsListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/tg/tgContacts/list`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgContactsList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
// tg 消息记录 列表请求
export function getTgMsgListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/tg/tgMsg/list`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgMsgList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
// tg 聊天室 会话分组 列表请求
export function getTgArtsFoldersAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/tg/arts/folders`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgArtsFoldersList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
// tg 聊天室 会话分组/消息队列 列表请求
export function getTgFoldersMessageAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.post(`/tg/arts/getDialogs`, {
                ...queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgFoldersMessageList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
// tg 聊天室 会话分组/消息队列/聊天历史 列表请求
export function getTgFoldersMeeageHistoryAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.post(`/tg/arts/getMsgHistory`, {
                ...queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgFoldersMeeageHistoryList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
