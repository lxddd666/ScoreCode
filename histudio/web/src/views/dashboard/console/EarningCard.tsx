import { useEffect, useState } from 'react';

// material-ui
import { styled, useTheme } from '@mui/material/styles';
import { Avatar, Box, Grid, Typography } from '@mui/material';

// project imports
import MainCard from 'ui-component/cards/MainCard';
import SkeletonEarningCard from 'ui-component/cards/Skeleton/EarningCard';

// assets
import EarningIcon from 'assets/images/icons/earning.svg';
import ArrowUpwardIcon from '@mui/icons-material/ArrowUpward';
import Request from 'utils/request'; // Make sure the path to your request utility is correct
import { useDispatch } from 'react-redux'; // Assuming you're using Redux
import { openSnackbar } from 'store/slices/snackbar'; // Verify the path for your snackbar slice
import ProgressBar from 'ui-component/ProgressBar';

// thrid party
import { useIntl } from 'react-intl';

const CardWrapper = styled(MainCard)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? theme.palette.dark.dark : theme.palette.primary.dark,
    color: '#fff',
    overflow: 'hidden',
    position: 'relative',

    '&::after, &::before': {
        content: '""',
        position: 'absolute',
        width: 210,
        height: 210,
        background: `linear-gradient(140.9deg, ${
            theme.palette.mode === 'dark' ? theme.palette.primary.dark : theme.palette.primary[200]
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

// ===========================|| DASHBOARD DEFAULT - EARNING CARD ||=========================== //

interface EarningCardProps {
    isLoading: boolean;
}

const EarningCard = ({ isLoading }: EarningCardProps) => {
    const theme = useTheme();
    const intl = useIntl();
    const dispatch = useDispatch();
    const [saleroom, setVisitData] = useState({
        saleroom: {
            weekSaleroom: null,
            amount: null,
            degree: null
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
                <SkeletonEarningCard />
            ) : (
                <CardWrapper border={false} content={false}>
                    <Box sx={{ p: 2.25 }}>
                        <Grid container direction="column">
                            <Grid item>
                                <Grid container justifyContent="space-between">
                                    <Grid item>
                                        <Avatar
                                            variant="rounded"
                                            sx={{
                                                ...theme.typography.commonAvatar,
                                                ...theme.typography.largeAvatar,
                                                backgroundColor:
                                                    theme.palette.mode === 'dark' ? theme.palette.dark.main : theme.palette.primary[200],
                                                mt: 1
                                            }}
                                        >
                                            <img src={EarningIcon} alt="Notification" />
                                        </Avatar>
                                    </Grid>
                                    <Grid item>
                                        <Typography
                                            sx={{
                                                fontSize: '1rem',
                                                fontWeight: 500,
                                                color:
                                                    theme.palette.mode === 'dark' ? theme.palette.text.primary : theme.palette.primary[200]
                                            }}
                                        >
                                            Week
                                        </Typography>
                                    </Grid>
                                </Grid>
                            </Grid>
                            <Grid item>
                                <Grid container alignItems="center">
                                    <Grid item>
                                        <Typography sx={{ fontSize: '2.125rem', fontWeight: 500, mr: 1, mt: 1.75, mb: 0.75 }}>
                                            ${saleroom.saleroom ? saleroom.saleroom.weekSaleroom : 'N/A'}
                                        </Typography>
                                    </Grid>
                                    <Grid item>
                                        <Avatar
                                            sx={{
                                                cursor: 'pointer',
                                                ...theme.typography.smallAvatar,
                                                backgroundColor: theme.palette.primary[200],
                                                color: theme.palette.primary.dark
                                            }}
                                        >
                                            <ArrowUpwardIcon fontSize="inherit" sx={{ transform: 'rotate3d(1, 1, 1, 45deg)' }} />
                                        </Avatar>
                                    </Grid>
                                </Grid>
                            </Grid>
                            <Grid item sx={{ mb: 1.25 }}>
                                <Typography
                                    sx={{
                                        fontSize: '1rem',
                                        fontWeight: 500,
                                        color: theme.palette.mode === 'dark' ? theme.palette.text.primary : theme.palette.primary[200]
                                    }}
                                >
                                    <ProgressBar value={saleroom.saleroom ? saleroom.saleroom.degree ?? 0 : 0} />
                                </Typography>
                            </Grid>
                            <Grid item sx={{ mb: 1.25 }}>
                                <Typography
                                    sx={{
                                        fontSize: '1rem',
                                        fontWeight: 500,
                                        color: theme.palette.mode === 'dark' ? theme.palette.text.primary : theme.palette.primary[200]
                                    }}
                                >
                                    Total Earning: {saleroom.saleroom ? saleroom.saleroom.amount : 'N/A'}
                                </Typography>
                            </Grid>
                        </Grid>
                    </Box>
                </CardWrapper>
            )}
        </>
    );
};

export default EarningCard;
