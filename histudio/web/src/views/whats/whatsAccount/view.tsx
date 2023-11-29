import React from 'react';
import { FormattedMessage, useIntl } from 'react-intl';
import { useParams } from 'react-router-dom';

import { Avatar, CardContent, Grid } from '@mui/material';

import MainCard from 'ui-component/cards/MainCard';
import Chip from 'ui-component/extended/Chip';

import { gridSpacing } from 'store/constant';
import { WhatsAccountListData } from 'types/whats';

import envRef from 'environment';
import axiosServices from 'utils/axios';

import { useDispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';

import { defaultErrorMessage } from 'constant/general';

const View = () => {
    const { id } = useParams();
    const intl = useIntl();
    const dispatch = useDispatch();

    const [data, setData] = React.useState<WhatsAccountListData | null>(null);

    const [whatsAccountStatusData, setWhatsAccountStatusData] = React.useState<any[]>([]);
    const [whatsAccountLoginStatusData, setWhatsAccountLoginStatusData] = React.useState<any[]>([]);

    React.useEffect(() => {
        getOptions();
        if (id) getData(parseInt(id));
    }, [id]);

    async function getOptions() {
        await axiosServices
            .get(`${envRef?.API_URL}admin/dictData/options?types[]=account_status&types[]=login_status`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setWhatsAccountStatusData(response.data.data.account_status);
                    setWhatsAccountLoginStatusData(response.data.data.login_status);
                }
            })
            .catch(function (error) {
                dispatch(
                    openSnackbar({
                        open: true,
                        message: error?.message || defaultErrorMessage,
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
            });
    }

    async function getData(id: number) {
        await axiosServices
            .get(`${envRef?.API_URL}whats/whatsAccount/view?id=${id}`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setData(response.data.data);
                }
            })
            .catch(function (error) {
                dispatch(
                    openSnackbar({
                        open: true,
                        message: error?.message || defaultErrorMessage,
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
            });
    }

    return (
        <MainCard title={<FormattedMessage id="whats.whats-account-management-details" />} content={false}>
            <CardContent>
                <Grid container alignItems="center" spacing={gridSpacing} justifyContent="flex-start">
                    <Grid item container xs={12} sm={6} md={3} alignItems="center">
                        <Grid item xs={6}>
                            {intl.formatMessage({ id: 'whats.account-number' })}
                        </Grid>
                        <Grid item xs={6}>
                            {data?.account || '-'}
                        </Grid>
                    </Grid>
                    <Grid item container xs={12} sm={6} md={3} alignItems="center">
                        <Grid item xs={6}>
                            {intl.formatMessage({ id: 'whats.account-nickname' })}
                        </Grid>
                        <Grid item xs={6}>
                            {data?.nickName || '-'}
                        </Grid>
                    </Grid>
                    <Grid item container xs={12} sm={6} md={3} alignItems="center">
                        <Grid item xs={6}>
                            {intl.formatMessage({ id: 'whats.account-avatar' })}
                        </Grid>
                        <Grid item xs={6}>
                            {data?.avatar === '' ? '-' : <Avatar src={data?.avatar} alt={'avatar.' + data?.id} />}
                        </Grid>
                    </Grid>
                    <Grid item container xs={12} sm={6} md={3} alignItems="center">
                        <Grid item xs={6}>
                            {intl.formatMessage({ id: 'whats.account-status' })}
                        </Grid>
                        <Grid item xs={6}>
                            <Chip
                                label={
                                    whatsAccountStatusData.length > 0
                                        ? whatsAccountStatusData.find((statusData) => statusData.value === data?.accountStatus).label
                                        : '-'
                                }
                                size="small"
                                chipcolor={
                                    whatsAccountStatusData.length > 0
                                        ? whatsAccountStatusData.find((statusData) => statusData.value === data?.accountStatus).listClass
                                        : 'default'
                                }
                            />
                        </Grid>
                    </Grid>
                    <Grid item container xs={12} sm={6} md={3} alignItems="center">
                        <Grid item xs={6}>
                            {intl.formatMessage({ id: 'general.is-online' })}
                        </Grid>
                        <Grid item xs={6}>
                            {whatsAccountLoginStatusData.length > 0
                                ? whatsAccountLoginStatusData.find((statusData) => statusData.value === data?.isOnline).label
                                : '-'}
                        </Grid>
                    </Grid>
                    <Grid item container xs={12} sm={6} md={3} alignItems="center">
                        <Grid item xs={6}>
                            {intl.formatMessage({ id: 'general.proxy-address' })}
                        </Grid>
                        <Grid item xs={6}>
                            {data?.proxyAddress || '-'}
                        </Grid>
                    </Grid>
                    <Grid item container xs={12} sm={6} alignItems="center">
                        <Grid item xs={6} md={3}>
                            {intl.formatMessage({ id: 'general.remarks' })}
                        </Grid>
                        <Grid item xs={6} md={9}>
                            {data?.comment || '-'}
                        </Grid>
                    </Grid>
                </Grid>
            </CardContent>
        </MainCard>
    );
};

export default View;
