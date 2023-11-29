// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';

const initialState: DefaultRootStateProps['log'] = {
    error: null,
    logList: null,
    log: null
};

const slice = createSlice({
    name: 'log',
    initialState,
    reducers: {
        // HAS ERROR
        hasError(state, action) {
            state.error = action.payload;
        },
        // GET LOG LIST
        getLogList(state, action) {
            state.logList = action.payload;
        },
        // GET LOG VIEW
        getLogView(state, action) {
            state.log = action.payload;
        }
    }
});

// Reducer
export default slice.reducer;

export function getLogList(queryParam: String) {
    return async () => {
        try {
            const response = await axios.get(`/admin/log/list${queryParam ? '?' + queryParam : ''}`);
            dispatch(slice.actions.getLogList(response.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}

export function getLogView(id: number) {
    return async () => {
        try {
            const response = await axios.get(`/admin/log/view/?id=${id}`);
            dispatch(slice.actions.getLogView(response.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
