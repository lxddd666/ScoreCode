import axios from 'axios';

import envRef from 'environment';

import { openSnackbar } from 'store/slices/snackbar';

import { LOGOUT } from 'store/actions';

interface requestParam {
    url: string;
    method: string;
    header?: { [key: string]: any };
    param?: { [key: string]: any };
}

async function Request({ url, method = 'POST', header, param }: requestParam, dispatch: any, intl? : any) {
    try {
        const serviceToken = await window.localStorage.getItem('serviceToken');
        let axiosParam = {
            method,
            url: `${envRef.API_URL}${url}`,
            data: param,
            headers: {
                Authorization: serviceToken,
                'Content-Type': 'application/json'
            }
        };

        if (header && Object.keys(header)?.length > 0) {
            axiosParam['headers'] = { ...axiosParam['headers'], ...header };
        }

        if (method != 'GET') {
            axiosParam['data'] = param;
        }

        const res = await axios(axiosParam);

        /* Token invalid or expired */
        if (res?.data?.code && res.data.code == 61) {
            if (typeof dispatch !== 'undefined') {
                dispatch(
                    openSnackbar({
                        open: true,
                        message: typeof intl !== "undefined" ? intl?.formatMessage({ id: 'general.session.expired' }) : 'Session expired. Please login again.',
                        variant: 'alert',
                        alert: {
                            color: 'error'
                        },
                        close: false,
                        anchorOrigin: {
                            vertical: 'top',
                            horizontal: 'center'
                        }
                    })
                );
            }

            setTimeout(async () => {
                localStorage.removeItem('serviceToken');
                await dispatch({ type: LOGOUT });
                window.location.href = '/login';
            }, 2000);
        } else {
            return res;
        }
    } catch (err) {
        if (typeof dispatch !== 'undefined') {
            dispatch(
                openSnackbar({
                    open: true,
                    message: typeof intl !== "undefined" ? intl?.formatMessage({ id: 'general.api.error' }) : '' + err,
                    variant: 'alert',
                    alert: {
                        color: 'error'
                    },
                    close: false,
                    anchorOrigin: {
                        vertical: 'top',
                        horizontal: 'center'
                    }
                })
            );
        }
        return {data: {msg : typeof intl !== "undefined" ? intl?.formatMessage({ id: 'general.api.error' }) : '' + err}}
    }
}

export default Request;
