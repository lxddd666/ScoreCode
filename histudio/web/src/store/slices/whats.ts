// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';

// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['whats'] = {
    error: null,
    whatsAccountList: null
};

const slice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        // HAS ERROR
        hasError(state, action) {
            state.error = action.payload;
        },

        // GET ROLE LIST
        getWhatsAccountList(state, action) {
            state.whatsAccountList = action.payload;
        }
    }
});

// Reducer
export default slice.reducer;

// ----------------------------------------------------------------------

export function getWhatsAccountList(queryParam?: String) {
    return async () => {
        try {
            const response = await axios.get(`/whats/whatsAccount/list${queryParam ? '?' + queryParam : ''}`);
            dispatch(slice.actions.getWhatsAccountList(response.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
