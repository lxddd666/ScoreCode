// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';
// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['loginlog'] = {
    error:null,
    loginlogList:  null,
    loginlogview: null

};

const slice = createSlice({
    name: 'loginlog',
    initialState,
    reducers: {
        // HAS ERROR
        hasError(state, action) {
            state.error = action.payload;
        },

        // GET POST LIST
        getLoginlogList(state, action) {
            state.loginlogList = action.payload;
        },

        // GET LONGINLOG VIEW
        getLoginlogView(state, action) {
            state.loginlogview = action.payload;
        },
    }
});

// Reducer
export default slice.reducer;

export function getLoginlogList(queryParam?: String) {
    return async () => {
        try {
            const response = await axios.get(`/admin/loginLog/list${queryParam ? '?' + queryParam : ''}`);
            dispatch(slice.actions.getLoginlogList(response.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}

export function getLoginlogView(id: number) {
    return async () => {
        try {
            const response = await axios.get(`/admin/loginLog/view/?id=${id}`);
            dispatch(slice.actions.getLoginlogView(response.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}

