import React, { createContext, useEffect, useReducer } from 'react';

// third-party
import { Chance } from 'chance';
import jwtDecode from 'jwt-decode';

// reducer - state management
import { LOGIN, LOGOUT } from 'store/actions';
import accountReducer from 'store/accountReducer';

// project imports
import Loader from 'ui-component/Loader';
import axios from 'utils/axios';

// types
import { KeyedObject } from 'types';
import { InitialLoginContextProps, JWTContextType } from 'types/auth';

// env
import envRef from 'environment';

// utils
import Request from 'utils/request';

import { useIntl } from 'react-intl';
import { useDispatch } from 'store';

const chance = new Chance();

// constant
const initialState: InitialLoginContextProps = {
    isLoggedIn: false,
    isInitialized: false,
    user: null
};

const verifyToken: (st: string) => boolean = (serviceToken) => {
    if (!serviceToken) {
        return false;
    }
    const decoded: KeyedObject = jwtDecode(serviceToken);

    const expiredDate = new Date(decoded?.loginAt);
    expiredDate.setDate(expiredDate.getDate() + envRef?.token?.expired);

    /**
     * Property 'exp' does not exist on type '<T = unknown>(token: string, options?: JwtDecodeOptions | undefined) => T'.
     */
    return decoded?.exp ? decoded.exp > Date.now() / 1000 : expiredDate?.getTime() > new Date()?.getTime();
};

const setSession = (serviceToken?: string | null) => {
    if (serviceToken) {
        localStorage.setItem('serviceToken', serviceToken);
        axios.defaults.headers.common.Authorization = `Bearer ${serviceToken}`;
    } else {
        localStorage.removeItem('serviceToken');
        delete axios.defaults.headers.common.Authorization;
    }
};

// ==============================|| JWT CONTEXT & PROVIDER ||============================== //
const JWTContext = createContext<JWTContextType | null>(null);

export const JWTProvider = ({ children }: { children: React.ReactElement }) => {
    const [state, dispatch] = useReducer(accountReducer, initialState);
    const [userInfo, setUserInfo] = React.useState<any>(null);
    const intl = useIntl()
    const dispatchFunc = useDispatch();

    const init = async () => {
        try {
            const serviceToken = window.localStorage.getItem('serviceToken');
            if (serviceToken && verifyToken(serviceToken)) {
                setSession(serviceToken);

                const response = await Request({ url: 'admin/member/info', method: 'GET' }, dispatchFunc, intl);
                const user = response?.data?.data;
                dispatch({
                    type: LOGIN,
                    payload: {
                        isLoggedIn: true,
                        user
                    }
                });
                setUserInfo(user);
            } else {
                dispatch({
                    type: LOGOUT
                });
            }
        } catch (err) {
            console.error(err);
            dispatch({
                type: LOGOUT
            });
        }
    };

    useEffect(() => {
        init();
    }, []);

    const login = async (email: string, password: string) => {
        const response = await axios.post('/api/account/login', { email, password });
        const { serviceToken, user } = response.data;
        setSession(serviceToken);
        dispatch({
            type: LOGIN,
            payload: {
                isLoggedIn: true,
                user
            }
        });
    };

    const loginUser = async (response: any) => {
        setSession(response?.data?.data?.token);
        dispatch({
            type: LOGIN,
            payload: {
                isLoggedIn: true,
                user: response?.data?.data
            }
        });
        init();
    };

    const register = async (email: string, password: string, firstName: string, lastName: string) => {
        // todo: this flow need to be recode as it not verified
        const id = chance.bb_pin();
        const response = await axios.post('/api/account/register', {
            id,
            email,
            password,
            firstName,
            lastName
        });
        let users = response.data;

        if (window.localStorage.getItem('users') !== undefined && window.localStorage.getItem('users') !== null) {
            const localUsers = window.localStorage.getItem('users');
            users = [
                ...JSON.parse(localUsers!),
                {
                    id,
                    email,
                    password,
                    name: `${firstName} ${lastName}`
                }
            ];
        }

        window.localStorage.setItem('users', JSON.stringify(users));
    };

    const logout = () => {
        setSession(null);
        sessionStorage.clear();
        localStorage.clear();
        dispatch({ type: LOGOUT });
    };

    const resetPassword = async (email: string) => {};

    const updateProfile = () => {};

    if (state.isInitialized !== undefined && !state.isInitialized) {
        return <Loader />;
    }

    return (
        <JWTContext.Provider value={{ ...state, login, logout, register, resetPassword, updateProfile, loginUser, userInfo }}>
            {children}
        </JWTContext.Provider>
    );
};

export default JWTContext;
