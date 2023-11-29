import { useEffect, useState } from 'react';
import { styled, useTheme } from '@mui/material/styles';
import { Box, Grid, Typography } from '@mui/material';
import MainCard from 'ui-component/cards/MainCard';
import SkeletonCommissionCard from 'ui-component/cards/Skeleton/CommissionCard';
import Request from 'utils/request'; // Make sure the path to your request utility is correct
import { useDispatch } from 'react-redux'; // Assuming you're using Redux
import { openSnackbar } from 'store/slices/snackbar'; // Verify the path for your snackbar slice
import { useIntl } from 'react-intl';

const CardWrapper = styled(MainCard)(({ theme }) => ({
    overflow: 'hidden',
    position: 'relative',
    '&:after': {
        content: '""',
        position: 'absolute',
        width: 210,
        height: 210,
        background: `linear-gradient(210.04deg, ${theme.palette.warning.dark} -50.94%, rgba(144, 202, 249, 0) 83.49%)`,
        borderRadius: '50%',
        top: -30,
        right: -180
    },
    '&:before': {
        content: '""',
        position: 'absolute',
        width: 210,
        height: 210,
        background: `linear-gradient(140.9deg, ${theme.palette.warning.dark} -14.02%, rgba(144, 202, 249, 0) 70.50%)`,
        borderRadius: '50%',
        top: -160,
        right: -130
    }
}));

interface CommissionCardProps {
    isLoading: boolean;
}

function CommissionCard({ isLoading }: CommissionCardProps) {
    const theme = useTheme();
    const intl = useIntl();
    const dispatch = useDispatch();
    const [volume, setVisitData] = useState({
        volume: {
            weekLarge: null,
            rise: null,
            decline: null,
            amount: null
        }
    });

    async function fetchData() {
        try {
            const response = await Request(
                {
                    url: 'admin/console/stat',
                    method: 'GET'
                },
                dispatch,
                intl
            );

            if (response?.data?.data) {
                setVisitData(response.data.data);
            } else {
                dispatch(
                    openSnackbar({
                        open: true,
                        message: response?.data.message,
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
        } catch (error) {
            dispatch(
                openSnackbar({
                    open: true,
                    message: 'An error occurred while fetching data.',
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
    }

    useEffect(() => {
        if (!isLoading) {
            fetchData();
        }
    }, [isLoading, dispatch]);

    return (
        <>
            {isLoading ? (
                <SkeletonCommissionCard />
            ) : (
                <CardWrapper border={false} content={false}>
                    <Box sx={{ p: 2.25 }}>
                        <Grid container direction="column">
                            <Grid container justifyContent="space-between">
                                <Grid item>
                                    <Typography
                                        sx={{
                                            fontSize: '1rem',
                                            fontWeight: 500,
                                            color: theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.grey[500]
                                        }}
                                    >
                                        Total Commission
                                    </Typography>
                                </Grid>
                                <Grid item>
                                    <Typography
                                        sx={{
                                            fontSize: '1rem',
                                            fontWeight: 500,
                                            color: theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.grey[500]
                                        }}
                                    >
                                        Month
                                    </Typography>
                                </Grid>
                            </Grid>
                            <Grid item>
                                <Typography sx={{ fontSize: '2.125rem', fontWeight: 500, mr: 1, mt: 1.75, mb: 0.75 }}>
                                    {volume.volume ? volume.volume.weekLarge : 'N/A'}
                                </Typography>
                            </Grid>
                            <Grid container justifyContent="space-between">
                                <Grid item>
                                    <Typography
                                        sx={{
                                            fontSize: '1.125rem',
                                            fontWeight: 500,
                                            color: theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.grey[500]
                                        }}
                                    >
                                        Rise: {volume.volume ? volume.volume.rise : 'N/A'}
                                    </Typography>
                                </Grid>
                                <Grid item>
                                    <Typography
                                        sx={{
                                            fontSize: '1.125rem',
                                            fontWeight: 500,
                                            color: theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.grey[500]
                                        }}
                                    >
                                        Decline: {volume.volume ? volume.volume.decline : 'N/A'}
                                    </Typography>
                                </Grid>
                            </Grid>
                            <Grid item>
                                <Typography
                                    sx={{
                                        fontSize: '1.125rem',
                                        fontWeight: 500,
                                        color: theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.grey[500]
                                    }}
                                >
                                    Amount: {volume.volume ? volume.volume.amount : 'N/A'}
                                </Typography>
                            </Grid>
                        </Grid>
                    </Box>
                </CardWrapper>
            )}
        </>
    );
}

export default CommissionCard;
