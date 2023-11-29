// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';
// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['smslog'] = {
    error:null,
    smslogList:  null
};

const slice = createSlice({
    name: 'smslog',
    initialState,
    reducers: {
        // HAS ERROR
        hasError(state, action) {
            state.error = action.payload;
        },

        // GET SERVELOG LIST
        getSmslogList(state, action) {
            state.smslogList = action.payload;
        },
    }
});

// Reducer
export default slice.reducer;

export function getSmslogList(queryParam?: String) {
    return async () => {
        try {
            const response = await axios.get(`/admin/smsLog/list${queryParam ? '?' + queryParam : ''}`);
            dispatch(slice.actions.getSmslogList(response.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}
