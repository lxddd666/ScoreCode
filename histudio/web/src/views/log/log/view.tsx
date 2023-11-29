import React, { useEffect } from 'react'; //, useState

// material-ui
import { Grid, Typography, Divider } from '@mui/material';
import MainCard from '../../../ui-component/cards/MainCard';

// project imports
import { gridSpacing } from 'store/constant';
import { useIntl } from 'react-intl';
import JsonViewer from './json-viewer';
import { useDispatch, useSelector } from 'store';
import { openSnackbar } from 'store/slices/snackbar';
import { getLogView } from 'store/slices/log';
import { useParams } from 'react-router-dom';
import { Response } from 'types/response';
import { Log } from 'types/log';
import Chip from 'ui-component/extended/Chip';

// ==============================|| DEFAULT DASHBOARD ||============================== //

const View = () => {
    const intl = useIntl();
    const defaultErrorMessage = intl.formatMessage({ id: 'auth-register.default-error' });
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);
    const [res, setRes] = React.useState<Response<Log>>();
    const [logDetail, setLogDetail] = React.useState<Log>();
    const { log } = useSelector((state) => state.log);
    const logId = useParams().id;

    // const [isLoading, setLoading] = useState(true);
    // useEffect(() => {
    // setLoading(false);
    // }, []);

    useEffect(() => {
        setRes(log!);
    }, [log]);

    useEffect(() => {
        if (res?.data) {
            setLogDetail(res?.data);
            console.log(logDetail?.timestamp);
            // console.log(
            //     new Intl.DateTimeFormat('en-US', {
            //         year: 'numeric',
            //         month: '2-digit',
            //         day: '2-digit',
            //         hour: '2-digit',
            //         minute: '2-digit',
            //         second: '2-digit'
            //     }).format(logDetail?.timestamp! * 1000)
            // );
        }
    }, [res]);

    useEffect(() => {
        setLoading(false);
        fetchData();
    }, [dispatch]);

    const fetchData = async () => {
        try {
            setLoading(true);
            await dispatch(getLogView(+(logId ? logId : 0)));
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
                                        {intl.formatMessage({ id: 'general-log.request-type' }) + `:   `}
                                        {!loading && logDetail?.method}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.access-ip' }) + `:   `}
                                        {logDetail?.ip}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.response-datetime' }) + `:   `}
                                        {logDetail && logDetail.timestamp && timeStampToDate(logDetail.timestamp)}
                                    </Grid>
                                </Grid>
                                <Grid item md={12} lg={4} container spacing={1}>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.request-path' }) + `:   `} {logDetail?.url}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.ip-origin' }) + `:   `}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.request-time' }) + `:   `} {logDetail?.createdAt}
                                    </Grid>
                                </Grid>
                                <Grid item md={12} lg={4} container spacing={1}>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.response-time' }) + `:   `} {logDetail?.takeUpTime}
                                    </Grid>
                                    <Grid item xs={12}>
                                        {intl.formatMessage({ id: 'general-log.req-id' }) + `:   `} {logDetail?.reqId}
                                    </Grid>
                                    <Grid item xs={12}>
                                        <br />
                                    </Grid>
                                </Grid>
                            </Grid>
                        </MainCard>
                    </Grid>
                    <Grid item xs={12}>
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
                    </Grid>
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
                                        {intl.formatMessage({ id: 'general-log.error-code' }) + `:   `} {logDetail?.errorCode}
                                    </Grid>
                                    <Grid item md={6}>
                                        {intl.formatMessage({ id: 'general-log.error-hint' }) + `:   `}
                                        <Chip
                                            title={`${logDetail?.errorMsg}`}
                                            label={`${logDetail?.errorMsg}`}
                                            size="small"
                                            chipcolor={logDetail?.errorCode === 0 ? 'success' : 'orange'}
                                            // maxWidth="150px"
                                        />
                                    </Grid>
                                </Grid>
                            </Grid>
                        </MainCard>
                    </Grid>
                    <Grid item xs={12}>
                        <JsonViewer
                            title={intl.formatMessage({ id: 'general-log.stack-print' })}
                            jsonString={{}}
                            // jsonString={{
                            //     Accept: ['application/json, text/plain, */*'],
                            //     'Accept-Encoding': ['gzip, deflate'],
                            //     'Accept-Language': ['zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2'],
                            //     'Cache-Control': ['no-cache'],
                            //     Dnt: ['1'],
                            //     Pragma: ['no-cache'],
                            //     Referer: ['http://8.222.195.54:4885/'],
                            //     'User-Agent': ['Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0']
                            // }}
                        />
                    </Grid>
                    <Grid item xs={12}>
                        <JsonViewer
                            title={intl.formatMessage({ id: 'general-log.request-header' })}
                            jsonString={logDetail?.headerData}
                            // jsonString={{
                            //     Accept: ['application/json, text/plain, */*'],
                            //     'Accept-Encoding': ['gzip, deflate'],
                            //     'Accept-Language': ['zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2'],
                            //     'Cache-Control': ['no-cache'],
                            //     Dnt: ['1'],
                            //     Pragma: ['no-cache'],
                            //     Referer: ['http://8.222.195.54:4885/'],
                            //     'User-Agent': ['Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0']
                            // }}
                        />
                    </Grid>
                    <Grid item xs={12}>
                        <JsonViewer title={intl.formatMessage({ id: 'general-log.get-param' })} jsonString={logDetail?.getData} />
                    </Grid>
                    <Grid item xs={12}>
                        <JsonViewer title={intl.formatMessage({ id: 'general-log.post-param' })} jsonString={logDetail?.postData} />
                    </Grid>
                </Grid>
            </Grid>
        </Grid>
    );
};

export default View;
