// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';

// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['script'] = {
    error: null,
    scriptList:[]
};

const slice = createSlice({
    name: 'tg',
    initialState,
    reducers: {
        hasError(state, action) {
            state.error = action.payload;
        },
        emitScriptList(state, action) {
            state.scriptList = action.payload;
        },
    }
});

// Reducer
export default slice.reducer;

// 异步请求提交----------------------------------------------------------------------
//  公司/个人 话术管理 请求
export function getSysScriptListAction(type:Number,queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/admin/${type}/sysScript/list`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitScriptList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
//  公司/个人 话术分组 请求
export function getScriptGroupListAction(type:Number,queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/admin/${type}/scriptGroup/list`, {
                params: queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitScriptList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
