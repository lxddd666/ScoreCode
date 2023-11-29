import React, { useEffect } from 'react'; //, useState

// material-ui
import { Grid, Typography, Divider } from '@mui/material';
import MainCard from '../../../ui-component/cards/MainCard';

// project imports
import { gridSpacing } from 'store/constant';
import { useIntl } from 'react-intl';
import JsonViewer from '../log/json-viewer';
import { useDispatch, useSelector } from 'store';
import { openSnackbar } from 'store/slices/snackbar';
import { getLoginlogView } from 'store/slices/loginlog';
import { useParams } from 'react-router-dom';
import { Response } from 'types/response';
import { LoginlogViewData } from 'types/loginlog';
import Chip from 'ui-component/extended/Chip';

// ==============================|| DEFAULT DASHBOARD ||============================== //

const LoginLogView = () => {
    const intl = useIntl();
    const defaultErrorMessage = intl.formatMessage({ id: 'auth-register.default-error' });
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);
    const [res, setRes] = React.useState<Response<LoginlogViewData>>();
    const [logDetail, setLogDetail] = React.useState<LoginlogViewData>();
    const { loginlogview } = useSelector((state) => state.loginlog);
    const logId = useParams().id;

    useEffect(() => {
        setRes(loginlogview!);
    }, [loginlogview]);

    useEffect(() => {
        if (res?.data) {
            setLogDetail(res?.data);
        }
    }, [res]);

    useEffect(() => {
        setLoading(false);
        fetchData();
    }, [dispatch]);

    const fetchData = async () => {
        try {
            setLoading(true);
            await dispatch(getLoginlogView(+(logId ? logId : 0)));
        } catch (error) {
            dispatch(
                openSnackbar({
                    open: true,
                    message: error || defaultErrorMessage,
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
        } finally {
            setLoading(false);
        }
    };

    const timeStampToDate = (timeStamp: number) => {
        let d = new Intl.DateTimeFormat('en-US', {
            year: 'numeric',
            month: '2-digit',
            day: '2-digit',
            hour: '2-digit',
            minute: '2-digit',
            second: '2-digit',
            hour12: false
        })
            .formatToParts(timeStamp * 1000) // * 1000 to convert from UNIX timestamp to javascript timestamp
            .reduce((acc, part) => {
                acc[part.type] = part.value;
                return acc;
            }, [] as any);
        console.log(d);
        return `${d.year}-${d.month}-${d.day} ${d.hour}:${d.minute} ${d.second}`;
    };

    return (
        // hardcode -- for api doesn't return
        <Grid container spacing={gridSpacing}>
            <Grid item xs={12}>
                <Grid container spacing={gridSpacing}>
                    <Grid item xs={12}>
                        <MainCard>
                            <Grid container spacing={2} style={{ whiteSpace: 'pre-wrap' }}>
                                <Grid item xs={12}>
                                    <Typography variant="h4">{intl.formatMessage({ id: 'general-log.log-detail' })}</Typography>
                                </Grid>
                                <Grid item xs={12}>
                                    <Divider />
                                </Grid>
                                <Grid item md={12} lg={4} container spacing={1}>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'user.username' }) + `:   `}
                                        {!loading && logDetail?.username}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.request-type' }) + `:   --`}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.access-ip' }) + `:   `}
                                        {!loading && logDetail?.loginIp}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.response-datetime' }) + `:   `}
                                        {res && res.timestamp && timeStampToDate(res.timestamp)}
                                    </Grid>
                                </Grid>
                                <Grid item md={12} lg={4} container spacing={1}>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.request-path' }) + `:   --`} 
                                        {/* {logDetail?.url} */}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.ip-origin' }) + `:   --`}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.request-time' }) + `:   `} 
                                        {logDetail?.createdAt}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'loginlog.updated-at' }) + `:   `} 
                                        {logDetail?.updatedAt}
                                    </Grid>
                                </Grid>
                                <Grid item md={12} lg={4} container spacing={1}>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.response-time' }) + `:   --`} 
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.req-id' }) + `:   `} 
                                        {logDetail?.reqId}
                                    </Grid>
                                    <Grid item xs={12}>
                                        <br />
                                    </Grid>
                                </Grid>
                            </Grid>
                        </MainCard>
                    </Grid>
                    {/* <Grid item xs={12}>
                        <MainCard>
                            <Grid container spacing={2}>
                                <Grid item xs={12}>
                                    <Typography variant="h4">{intl.formatMessage({ id: 'general-log.access-agent' })}</Typography>
                                </Grid>
                                <Grid item xs={12}>
                                    <Divider />
                                </Grid>
                                <Grid item xs={12}>
                                    {logDetail?.userAgent}
                                </Grid>
                            </Grid>
                        </MainCard>
                    </Grid> */}
                    <Grid item xs={12}>
                        <MainCard>
                            <Grid container spacing={2}>
                                <Grid item xs={12}>
                                    <Typography variant="h4">{intl.formatMessage({ id: 'general-log.error-message' })}</Typography>
                                </Grid>
                                <Grid item xs={12}>
                                    <Divider />
                                </Grid>
                                <Grid item container spacing={2}>
                                    <Grid item md={6}>
                                        {intl.formatMessage({ id: 'general-log.error-code' }) + `:   `} {logDetail?.status}
                                    </Grid>
                                    <Grid item md={6}>
                                        {intl.formatMessage({ id: 'general-log.error-hint' }) + `:   `}
                                        <Chip
                                            title={`${logDetail?.status}`}
                                            label={logDetail?.status === 1 ? intl.formatMessage({ id: 'loginlog.successful-operation' }) 
                                            : intl.formatMessage({ id: 'loginlog.operation-failed' })}
                                            size="small"
                                            chipcolor={logDetail?.status === 1 ? 'success' : 'orange'}
                                            // maxWidth="150px"
                                        />
                                    </Grid>
                                </Grid>
                            </Grid>
                        </MainCard>
                    </Grid>
                    <Grid item xs={12}>
                        <JsonViewer title={intl.formatMessage({ id: 'loginlog.response' })} jsonString={logDetail?.response} />
                    </Grid>
                </Grid>
            </Grid>
        </Grid>
    );
};

export default LoginLogView;
