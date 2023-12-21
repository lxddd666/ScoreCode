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
    orgList: []
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
        }
    }
});

// Reducer
export default slice.reducer;

// 异步请求提交----------------------------------------------------------------------
// Org请求
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