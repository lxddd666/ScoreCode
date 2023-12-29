// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';

// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['org'] = {
    error: null,
    orgList: [],
    proxyList: []
};

const slice = createSlice({
    name: 'org',
    initialState,
    reducers: {
        hasError(state, action) {
            state.error = action.payload;
        },
        //公司信息列表
        emitOrgList(state, action) {
            state.orgList = action.payload;
        },
        //代理管理列表
        emitProxyList(state, action) {
            state.proxyList = action.payload;
        }
    }
});

// Reducer
export default slice.reducer;

// 异步请求提交----------------------------------------------------------------------
// 公司列表请求
export function getOrgListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/admin/org/list`, {
                params: queryParam
            });
            dispatch(slice.actions.emitOrgList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
//代理管理列表请求
export function getProxyListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/admin/sysProxy/list`, {
                params: queryParam
            });
            dispatch(slice.actions.emitProxyList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}

