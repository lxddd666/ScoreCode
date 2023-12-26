// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';

const initialState: DefaultRootStateProps['dict'] = {
    error: null,
    dictList: null
};

const slice = createSlice({
    name: 'dict',
    initialState,
    reducers: {
        // HAS ERROR
        hasError(state, action) {
            state.error = action.payload;
        },
        // 字典管理列表
        emitDictList(state, action) {
            state.dictList = action.payload;
        },
    }
});

// Reducer
export default slice.reducer;
//字典管理列表请求
export function getDictList(queryParam: { valueType: undefined; pageSize: number; page: number; id: undefined; type: undefined; value: undefined; status: undefined }) {
    return async () => {
        try {
            const response = await axios.get(`/admin/dictData/list`);
            dispatch(slice.actions.emitDictList(response.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}