import { useEffect, useState } from 'react';
import { styled, useTheme } from '@mui/material/styles';
import { Avatar, Box, Grid, Typography } from '@mui/material';
import MainCard from 'ui-component/cards/MainCard';
import SkeletonTotalVisitCard from 'ui-component/cards/Skeleton/TotalVisitCard';
import ArrowUpwardIcon from '@mui/icons-material/ArrowUpward';
import Request from 'utils/request'; // Make sure the path to your request utility is correct
import { useDispatch } from 'react-redux'; // Assuming you're using Redux
import { openSnackbar } from 'store/slices/snackbar'; // Verify the path for your snackbar slice
import { FormattedMessage } from 'react-intl';
import { useIntl } from 'react-intl';

const CardWrapper = styled(MainCard)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? theme.palette.dark.dark : theme.palette.secondary.dark,
    color: '#fff',
    overflow: 'hidden',
    position: 'relative',

    '&::after, &::before': {
        content: '""',
        position: 'absolute',
        width: 210,
        height: 210,
        background: `linear-gradient(140.9deg, ${
            theme.palette.mode === 'dark' ? theme.palette.secondary.dark : theme.palette.secondary[800]
        } -14.02%, rgba(144, 202, 249, 0) 70.50%)`,
        borderRadius: '50%'
    },

    '&::after': {
        top: -85,
        right: -95
    },

    '&::before': {
        top: -125,
        right: -15,
        opacity: 0.5,
        [theme.breakpoints.down('sm')]: {
            top: -155,
            right: -70
        }
    }
}));

interface TotalVisitProps {
    isLoading: boolean;
}

function TotalVisitCard({ isLoading }: TotalVisitProps) {
    const theme = useTheme();
    const intl = useIntl();

    const dispatch = useDispatch();
    const [visits, setVisitData] = useState({
        visits: {
            dayVisits: null,
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
                <SkeletonTotalVisitCard />
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
                                            color:
                                                theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.secondary[200]
                                        }}
                                    >
                                        <FormattedMessage id="dashboard.totalVisitCard" />
                                    </Typography>
                                </Grid>
                                <Grid item>
                                    <Typography
                                        sx={{
                                            fontSize: '1rem',
                                            fontWeight: 500,
                                            color:
                                                theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.secondary[200]
                                        }}
                                    >
                                        Day
                                    </Typography>
                                </Grid>
                            </Grid>
                            <Grid item>
                                <Typography sx={{ fontSize: '2.125rem', fontWeight: 500, mr: 1, mt: 1.75, mb: 0.75 }}>
                                    {visits.visits ? visits.visits.dayVisits : 'N/A'}
                                </Typography>
                            </Grid>
                            <Grid item>
                                <Avatar
                                    sx={{
                                        cursor: 'pointer',
                                        backgroundColor: theme.palette.secondary[200],
                                        color: theme.palette.secondary.dark,
                                        ...theme.typography.smallAvatar
                                    }}
                                >
                                    <ArrowUpwardIcon fontSize="inherit" sx={{ transform: 'rotate3d(1, 1, 1, 45deg)' }} />
                                </Avatar>
                            </Grid>
                            <Grid container justifyContent="space-between">
                                <Grid item>
                                    <Typography
                                        sx={{
                                            fontSize: '1.125rem',
                                            fontWeight: 500,
                                            color:
                                                theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.secondary[200]
                                        }}
                                    >
                                        Rise: {visits.visits ? visits.visits.rise : 'N/A'}
                                    </Typography>
                                </Grid>
                                <Grid item>
                                    <Typography
                                        sx={{
                                            fontSize: '1.125rem',
                                            fontWeight: 500,
                                            color:
                                                theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.secondary[200]
                                        }}
                                    >
                                        Decline: {visits.visits ? visits.visits.decline : 'N/A'}
                                    </Typography>
                                </Grid>
                            </Grid>
                            <Grid item>
                                <Typography
                                    sx={{
                                        fontSize: '1.125rem',
                                        fontWeight: 500,
                                        color: theme.palette.mode === 'dark' ? theme.palette.text.secondary : theme.palette.secondary[200]
                                    }}
                                >
                                    Amount: {visits.visits ? visits.visits.amount : 'N/A'}
                                </Typography>
                            </Grid>
                        </Grid>
                    </Box>
                </CardWrapper>
            )}
        </>
    );
}

export default TotalVisitCard;
