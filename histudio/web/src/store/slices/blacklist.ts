// third-party
import { createSlice } from '@reduxjs/toolkit';

// project imports
import axios from 'utils/axios';
import { dispatch } from '../index';

// types
import { DefaultRootStateProps } from 'types';

// ----------------------------------------------------------------------

const initialState: DefaultRootStateProps['blacklist'] = {
    error: null,
    List: {
        code: -1,
        message: '',
        data: {
            list:
                [
                    {
                        id: -1,
                        ip: '',
                        remark: '',
                        status: -1,
                        createdAt: '',
                        updatedAt: ''
                    }
                ],
            page: 1,
            pageSize: 10,
            pageCount: 0,
            totalCount: 0
        }
        ,
        error: {},
        timestamp: -1,
        traceID: ''
    }
};

const slice = createSlice({
    name: 'blacklist',
    initialState,
    reducers: {
        // HAS ERROR
        hasError(state, action) {
            state.error = action.payload;
        },

        // GET BLACK LIST
        getBlackList(state, action) {
            state.List = action.payload;
        }
    }
});

// Reducer
export default slice.reducer;

// ----------------------------------------------------------------------

export function getBlackList(queryParam: String) {
    return async () => {
        try {
            const response = await axios.get(`/admin/blacklist/list${queryParam ? '?' + queryParam : ''}`);
            dispatch(slice.actions.getBlackList(response.data));
        } catch (error) {
            dispatch(slice.actions.hasError(error));
        }
    };
}