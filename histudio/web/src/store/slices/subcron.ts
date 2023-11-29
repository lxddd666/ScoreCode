// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';
// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['subcron'] = {
    error:null,
    postList:  null
};

const slice = createSlice({
    name: 'subcron',
    initialState,
    reducers: {
        // HAS ERROR
        hasError(state, action) {
            state.error = action.payload;
        },

        // GET POST LIST
        getPostList(state, action) {
            state.postList = action.payload;
        },
    }
});

// Reducer
export default slice.reducer;

export function getPostList(queryParam?: String) {
    return async () => {
        try {
            const response = await axios.get(`/admin/cronGroup/list${queryParam ? '?' + queryParam : ''}`);
            dispatch(slice.actions.getPostList(response.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}

