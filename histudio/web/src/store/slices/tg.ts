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
    tgUserList: []
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
        }
    }
});

// Reducer
export default slice.reducer;

// 异步请求提交----------------------------------------------------------------------

export function getTgUserListAction(queryParam?: any) {
    return async () => {
        try {
            const res = await axios.get(`/tg/tgUser/list`,{
                params:queryParam
            });
            // console.log('tgUser列表',res)
            dispatch(slice.actions.emitTgUserList(res.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
