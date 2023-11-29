import React, { useEffect, useRef, useState } from 'react';
import { useIntl } from 'react-intl';
import useWebSocket from 'hooks/useWebSocket';

// material-ui
import { useTheme } from '@mui/material/styles';
import {
    Avatar,
    Badge,
    BadgeProps,
    Box,
    Button,
    CardActions,
    ClickAwayListener,
    Divider,
    Grid,
    Paper,
    Popper,
    Stack,
    Tab,
    Tabs,
    TextField,
    Typography,
    styled,
    useMediaQuery
} from '@mui/material';

// third-party
import PerfectScrollbar from 'react-perfect-scrollbar';

// project imports
import MainCard from 'ui-component/cards/MainCard';
import Transitions from 'ui-component/extended/Transitions';
import NotificationList from './NotificationList';

// assets
import { IconBell } from '@tabler/icons';
import NotificationImportantTwoToneIcon from '@mui/icons-material/NotificationImportantTwoTone';
import CampaignTwoToneIcon from '@mui/icons-material/CampaignTwoTone';
import QuestionAnswerTwoToneIcon from '@mui/icons-material/QuestionAnswerTwoTone';

// API
import axiosServices from 'utils/axios';
import envRef from 'environment';
import { Spin } from 'antd';
import { Link } from 'react-router-dom';

// ==============================|| NOTIFICATION ||============================== //

const NotificationSection = () => {
    const theme = useTheme();
    const matchesXs = useMediaQuery(theme.breakpoints.down('md'));
    const intl = useIntl();
    const { notificationData, getNotificationData } = useWebSocket();

    // notification status options
    const status = [
        {
            value: -1,
            label: intl.formatMessage({ id: 'general.all' })
        },
        {
            value: 9,
            label: intl.formatMessage({ id: 'general.unread' })
        },
        {
            value: 8,
            label: intl.formatMessage({ id: 'general.read' })
        }
    ];

    const tabsOption = [
        {
            label: intl.formatMessage({ id: 'inbox.notice' }),
            icon: <CampaignTwoToneIcon sx={{ fontSize: '1.3rem' }} />
        },
        {
            label: intl.formatMessage({ id: 'inbox.announcement' }),
            icon: <NotificationImportantTwoToneIcon sx={{ fontSize: '1.3rem' }} />
        },
        {
            label: intl.formatMessage({ id: 'inbox.private-message' }),
            icon: <QuestionAnswerTwoToneIcon sx={{ fontSize: '1.3rem' }} />
        }
    ];

    const StyledBadge = styled(Badge)<BadgeProps>(({ theme }) => ({
        '& .MuiBadge-badge': {
            height: '16px',
            minWidth: '16px',
            right: 3,
            top: 3,
            border: `1px solid ${theme.palette.background.paper}`,
            padding: '0 2px',
            fontSize: '0.5rem'
        }
    }));

    const [open, setOpen] = useState(false);
    const [value, setValue] = useState('-1');
    const [tabValue, setTabValue] = React.useState(0);
    const [rawData, setRawData] = useState<any>(notificationData);
    const [totalCount, setTotalCount] = useState<number>(0);
    const [isLoading, setIsLoading] = React.useState<boolean>(false);

    React.useEffect(() => {
        setRawData(notificationData);
        if (notificationData?.list)
            setTotalCount(notificationData?.letterCount + notificationData?.notifyCount + notificationData?.noticeCount);
    }, [notificationData]);

    React.useEffect(() => {
        getLatestData();
    }, [tabValue]);

    async function getLatestData() {
        try {
            setIsLoading(true);
            await getNotificationData();
        } catch (error) {
            console.error(error);
        } finally {
            setIsLoading(false);
        }
    }

    /**
     * anchorRef is used on different componets and specifying one type leads to other components throwing an error
     * */
    const anchorRef = useRef<any>(null);

    const handleToggle = () => {
        setOpen((prevOpen) => !prevOpen);
    };

    const handleClose = (event: React.MouseEvent<HTMLDivElement> | MouseEvent | TouchEvent) => {
        if (anchorRef.current && anchorRef.current.contains(event.target)) {
            return;
        }
        setOpen(false);
    };

    const prevOpen = useRef(open);
    useEffect(() => {
        if (prevOpen.current === true && open === false) {
            anchorRef.current.focus();
        }
        prevOpen.current = open;
    }, [open]);

    const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement> | undefined) => {
        event?.target.value && setValue(event?.target.value);
    };

    const handleTabChange = (event: React.SyntheticEvent, newValue: number) => {
        setTabValue(newValue);
    };

    async function handleReadAll() {
        try {
            setIsLoading(true);
            for (let i = 1; i < 4; i++) {
                await axiosServices
                    .post(`${envRef?.API_URL}admin/notice/readAll`, { type: i }, { headers: {} })
                    .then(function (response) {
                        if (response?.data?.code === 0) {
                            getNotificationData();
                        }
                    })
                    .catch(function (error) {
                        console.error(error);
                    });
            }
        } catch (error) {
            console.error(error);
        } finally {
            setIsLoading(false);
        }
    }

    function a11yProps(index: number) {
        return {
            id: `simple-tab-${index}`,
            'aria-controls': `simple-tabpanel-${index}`
        };
    }

    return (
        <>
            <Box
                sx={{
                    ml: 2,
                    mr: 3,
                    [theme.breakpoints.down('md')]: {
                        mr: 2
                    }
                }}
            >
                <Avatar
                    variant="rounded"
                    sx={{
                        ...theme.typography.commonAvatar,
                        ...theme.typography.mediumAvatar,
                        transition: 'all .2s ease-in-out',
                        background: theme.palette.mode === 'dark' ? theme.palette.dark.main : theme.palette.secondary.light,
                        color: theme.palette.mode === 'dark' ? theme.palette.warning.dark : theme.palette.secondary.dark,
                        '&[aria-controls="menu-list-grow"],&:hover': {
                            background: theme.palette.mode === 'dark' ? theme.palette.warning.dark : theme.palette.secondary.dark,
                            color: theme.palette.mode === 'dark' ? theme.palette.grey[800] : theme.palette.secondary.light
                        }
                    }}
                    ref={anchorRef}
                    aria-controls={open ? 'menu-list-grow' : undefined}
                    aria-haspopup="true"
                    onClick={handleToggle}
                    color="inherit"
                >
                    {totalCount > 0 && (
                        <StyledBadge badgeContent={totalCount > 9 ? '9+' : totalCount} color="error">
                            <IconBell stroke={1.5} size="20px" />
                        </StyledBadge>
                    )}
                    {totalCount === 0 && <IconBell stroke={1.5} size="20px" />}
                </Avatar>
            </Box>

            <Popper
                placement={matchesXs ? 'bottom' : 'bottom-end'}
                open={open}
                anchorEl={anchorRef.current}
                role={undefined}
                transition
                disablePortal
                modifiers={[
                    {
                        name: 'offset',
                        options: {
                            offset: [matchesXs ? 5 : 0, 20]
                        }
                    }
                ]}
            >
                {({ TransitionProps }) => (
                    <ClickAwayListener onClickAway={handleClose}>
                        <Transitions position={matchesXs ? 'top' : 'top-right'} in={open} {...TransitionProps}>
                            <Paper>
                                {open && (
                                    <MainCard border={false} elevation={16} content={false} boxShadow shadow={theme.shadows[16]}>
                                        <Grid container direction="column" spacing={2}>
                                            <Grid item xs={12}>
                                                <Grid container alignItems="center" justifyContent="space-between" sx={{ pt: 2, px: 2 }}>
                                                    <Grid item>
                                                        <Stack direction="row" spacing={2}>
                                                            <Typography variant="subtitle1">
                                                                {intl.formatMessage({ id: 'general.notification' })}
                                                            </Typography>
                                                        </Stack>
                                                    </Grid>
                                                    <Grid item>
                                                        <Button
                                                            size="small"
                                                            variant="outlined"
                                                            color="primary"
                                                            onClick={() => handleReadAll()}
                                                        >
                                                            {intl.formatMessage({ id: 'general.mark-all-as-read' })}
                                                        </Button>
                                                    </Grid>
                                                </Grid>
                                            </Grid>
                                            <Grid item xs={12}>
                                                <PerfectScrollbar
                                                    style={{ height: '100%', maxHeight: 'calc(100vh - 205px)', overflowX: 'hidden' }}
                                                >
                                                    <Grid container direction="column" spacing={2}>
                                                        <Grid item xs={12}>
                                                            <Box sx={{ px: 2, pt: 0.25 }}>
                                                                <TextField
                                                                    id="outlined-select-currency-native"
                                                                    select
                                                                    fullWidth
                                                                    value={value}
                                                                    onChange={handleChange}
                                                                    SelectProps={{
                                                                        native: true
                                                                    }}
                                                                >
                                                                    {status.map((option) => (
                                                                        <option key={option.value} value={option.value}>
                                                                            {option.label}
                                                                        </option>
                                                                    ))}
                                                                </TextField>
                                                            </Box>
                                                            <Box>
                                                                <Tabs
                                                                    value={tabValue}
                                                                    indicatorColor="primary"
                                                                    textColor="primary"
                                                                    onChange={handleTabChange}
                                                                    aria-label="simple tabs example"
                                                                    variant="fullWidth"
                                                                >
                                                                    {tabsOption.map((tab, index) => (
                                                                        <Tab
                                                                            key={index}
                                                                            component={Link}
                                                                            to="#"
                                                                            icon={tab.icon}
                                                                            label={tab.label}
                                                                            {...a11yProps(index)}
                                                                        />
                                                                    ))}
                                                                </Tabs>
                                                            </Box>
                                                        </Grid>
                                                    </Grid>
                                                    {isLoading && (
                                                        <>
                                                            <Spin
                                                                style={{ position: 'absolute', left: '50%', zIndex: 1, marginTop: '1rem' }}
                                                            />
                                                            <Grid sx={{ marginY: '2rem' }}>&nbsp;</Grid>
                                                        </>
                                                    )}
                                                    {!isLoading && (
                                                        <NotificationList data={rawData} status={parseInt(value)} typeValue={tabValue} />
                                                    )}
                                                </PerfectScrollbar>
                                            </Grid>
                                        </Grid>
                                        <Divider />
                                        <CardActions sx={{ p: 1.25, justifyContent: 'center' }}>
                                            <Link to={`/account/inbox/${tabValue + 1}`}>
                                                <Button
                                                    size="small"
                                                    disableElevation
                                                    onClick={() => {
                                                        setOpen(false);
                                                    }}
                                                >
                                                    {intl.formatMessage({ id: 'general.view-all' })}
                                                </Button>
                                            </Link>
                                        </CardActions>
                                    </MainCard>
                                )}
                            </Paper>
                        </Transitions>
                    </ClickAwayListener>
                )}
            </Popper>
        </>
    );
};

export default NotificationSection;
