import React from 'react';
import { useIntl } from 'react-intl';
import { useDispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';
import useWebSocket from 'hooks/useWebSocket';

// material-ui
import { useTheme, styled } from '@mui/material/styles';
import { Avatar, Badge, BadgeProps, Divider, Grid, List, ListItem, ListItemAvatar, ListItemText, Typography } from '@mui/material';
import { Spin } from 'antd';

// assets
import NotificationImportantTwoToneIcon from '@mui/icons-material/NotificationImportantTwoTone';
import CampaignTwoToneIcon from '@mui/icons-material/CampaignTwoTone';
import QuestionAnswerTwoToneIcon from '@mui/icons-material/QuestionAnswerTwoTone';
import FolderOffTwoToneIcon from '@mui/icons-material/FolderOffTwoTone';
import Chip from 'ui-component/extended/Chip';

// API
import envRef from 'environment';
import axiosServices from 'utils/axios';

// styles
const ListItemWrapper = styled('div')(({ theme }) => ({
    cursor: 'pointer',
    padding: 16,
    '&:hover': {
        background: theme.palette.mode === 'dark' ? theme.palette.dark.main : theme.palette.primary.light
    },
    '& .MuiListItem-root': {
        padding: 0
    }
}));

type PropType = {
    data: any;
    status: number;
    typeValue: number;
};

// ==============================|| NOTIFICATION LIST ITEM ||============================== //

const NotificationList = ({ data, status, typeValue }: PropType) => {
    const theme = useTheme();
    const intl = useIntl();
    const dispatch = useDispatch();
    const { getNotificationData } = useWebSocket();
    const defaultErrorMessage = intl.formatMessage({ id: 'auth-register.default-error' });

    const [isLoading, setIsLoading] = React.useState<boolean>(false);

    const StyledBadge = styled(Badge)<BadgeProps>(({ theme }) => ({
        '& .MuiBadge-badge': {
            height: '16px',
            minWidth: '16px',
            right: 4,
            top: 4,
            zIndex: 1,
            border: `1px solid ${theme.palette.background.paper}`,
            padding: '0 2px',
            fontSize: '0.5rem'
        }
    }));

    function handleTag(tag: number) {
        switch (tag) {
            case 0:
                return <></>;
            case 1:
                return <Chip label={intl.formatMessage({ id: 'general.normal' })} size="medium" chipcolor="primary" />;
            case 2:
                return <Chip label={intl.formatMessage({ id: 'general.urgent' })} size="medium" chipcolor="error" />;
            case 3:
                return <Chip label={intl.formatMessage({ id: 'general.important' })} size="medium" chipcolor="warning" />;
            case 4:
                return <Chip label={intl.formatMessage({ id: 'general.remind' })} size="medium" chipcolor="success" />;
            case 5:
                return <Chip label={intl.formatMessage({ id: 'general.secondary' })} size="medium" chipcolor="" />;
            default:
                return <></>;
        }
    }

    function handleAvatar(type: number, isRead: boolean) {
        switch (type) {
            case 1:
                return <NotificationImportantTwoToneIcon />;
            case 2:
                return <CampaignTwoToneIcon />;
            case 3:
                return <QuestionAnswerTwoToneIcon />;
            default:
                return;
        }
    }

    async function handleRead(id: number) {
        try {
            setIsLoading(true);
            await axiosServices
                .post(`${envRef?.API_URL}admin/notice/upRead`, { id: id }, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        getNotificationData();
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
            setIsLoading(false);
        }
    }

    function renderData(data: any) {
        let filteredData: any = data.filter((obj: any) => obj.type === typeValue + 1);
        if (status === 8) {
            filteredData = filteredData.filter((obj: any) => obj.isRead);
        } else if (status === 9) {
            filteredData = filteredData.filter((obj: any) => !obj.isRead);
        }
        return (
            <>
                {isLoading && (
                    <>
                        <Spin style={{ position: 'absolute', left: '50%', zIndex: 1 }} />
                        <Grid sx={{ marginTop: '1rem', marginBottom: '3rem' }}>&nbsp;</Grid>
                    </>
                )}
                {!isLoading &&
                    filteredData.length > 0 &&
                    filteredData.map((obj: any, index: number) => {
                        return (
                            <ListItemWrapper
                                key={index}
                                onClick={() => {
                                    if (!obj.isRead) handleRead(obj.id);
                                }}
                            >
                                <ListItem alignItems="center">
                                    <ListItemAvatar>
                                        <StyledBadge badgeContent={obj.isRead ? undefined : ''} color="error">
                                            <Avatar alt="avatar" src={obj.senderAvatar !== '' ? obj.senderAvatar : undefined}>
                                                {handleAvatar(obj.type, obj.isRead)}
                                            </Avatar>
                                        </StyledBadge>
                                    </ListItemAvatar>
                                    <ListItemText primary={obj.title} />
                                    <ListItemText>
                                        <Grid container justifyContent="flex-end">
                                            <Grid item xs={12}>
                                                <Typography variant="body1">{obj.createdAt}</Typography>
                                            </Grid>
                                        </Grid>
                                    </ListItemText>
                                </ListItem>
                                <Grid container direction="column" className="list-container">
                                    <Grid item xs={12} sx={{ pb: 2 }}>
                                        <Typography variant="subtitle2" dangerouslySetInnerHTML={{ __html: obj.content }} />
                                    </Grid>
                                    <Grid item xs={12}>
                                        <Grid container>
                                            <Grid item>{handleTag(obj.tag)}</Grid>
                                        </Grid>
                                    </Grid>
                                </Grid>
                                {index < data.length - 1 && <Divider sx={{ paddingBottom: '1rem' }} />}
                            </ListItemWrapper>
                        );
                    })}
                {!isLoading && filteredData.length === 0 && (
                    <Grid textAlign="center" padding="1rem">
                        <FolderOffTwoToneIcon sx={{ verticalAlign: 'bottom' }} />
                        {intl.formatMessage({ id: 'general.no-records' })}
                    </Grid>
                )}
            </>
        );
    }

    return (
        <List
            sx={{
                width: '100%',
                py: 0,
                borderRadius: '10px',
                [theme.breakpoints.down('md')]: {
                    maxWidth: 320
                },
                '& .MuiListItemSecondaryAction-root': {
                    top: 22
                },
                '& .MuiDivider-root': {
                    my: 0
                },
                '& .list-container': {
                    pl: 7
                }
            }}
        >
            {data.list && renderData(data.list)}

            {/* <Divider />
            <ListItemWrapper>
                <ListItem alignItems="center">
                    <ListItemAvatar>
                        <Avatar
                            sx={{
                                color: theme.palette.success.dark,
                                backgroundColor: theme.palette.mode === 'dark' ? theme.palette.dark.main : theme.palette.success.light,
                                border: theme.palette.mode === 'dark' ? '1px solid' : 'none',
                                borderColor: theme.palette.success.main
                            }}
                        >
                            <IconBuildingStore stroke={1.5} size="20px" />
                        </Avatar>
                    </ListItemAvatar>
                    <ListItemText primary={<Typography variant="subtitle1">Store Verification Done</Typography>} />
                    <ListItemSecondaryAction>
                        <Grid container justifyContent="flex-end">
                            <Grid item xs={12}>
                                <Typography variant="caption" display="block" gutterBottom>
                                    2 min ago
                                </Typography>
                            </Grid>
                        </Grid>
                    </ListItemSecondaryAction>
                </ListItem>
                <Grid container direction="column" className="list-container">
                    <Grid item xs={12} sx={{ pb: 2 }}>
                        <Typography variant="subtitle2">We have successfully received your request.</Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid container>
                            <Grid item>
                                <Chip label="Unread" sx={chipErrorSX} />
                            </Grid>
                        </Grid>
                    </Grid>
                </Grid>
            </ListItemWrapper>
            <Divider />
            <ListItemWrapper>
                <ListItem alignItems="center">
                    <ListItemAvatar>
                        <Avatar
                            sx={{
                                color: theme.palette.primary.dark,
                                backgroundColor: theme.palette.mode === 'dark' ? theme.palette.dark.main : theme.palette.primary.light,
                                border: theme.palette.mode === 'dark' ? '1px solid' : 'none',
                                borderColor: theme.palette.primary.main
                            }}
                        >
                            <IconMailbox stroke={1.5} size="20px" />
                        </Avatar>
                    </ListItemAvatar>
                    <ListItemText primary={<Typography variant="subtitle1">Check Your Mail.</Typography>} />
                    <ListItemSecondaryAction>
                        <Grid container justifyContent="flex-end">
                            <Grid item>
                                <Typography variant="caption" display="block" gutterBottom>
                                    2 min ago
                                </Typography>
                            </Grid>
                        </Grid>
                    </ListItemSecondaryAction>
                </ListItem>
                <Grid container direction="column" className="list-container">
                    <Grid item xs={12} sx={{ pb: 2 }}>
                        <Typography variant="subtitle2">All done! Now check your inbox as you&apos;re in for a sweet treat!</Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid container>
                            <Grid item>
                                <Button variant="contained" disableElevation endIcon={<IconBrandTelegram stroke={1.5} size="20px" />}>
                                    Mail
                                </Button>
                            </Grid>
                        </Grid>
                    </Grid>
                </Grid>
            </ListItemWrapper>
            <Divider />
            <ListItemWrapper>
                <ListItem alignItems="center">
                    <ListItemAvatar>
                        <Avatar alt="John Doe" src={User1} />
                    </ListItemAvatar>
                    <ListItemText primary={<Typography variant="subtitle1">John Doe</Typography>} />
                    <ListItemSecondaryAction>
                        <Grid container justifyContent="flex-end">
                            <Grid item xs={12}>
                                <Typography variant="caption" display="block" gutterBottom>
                                    2 min ago
                                </Typography>
                            </Grid>
                        </Grid>
                    </ListItemSecondaryAction>
                </ListItem>
                <Grid container direction="column" className="list-container">
                    <Grid item xs={12} sx={{ pb: 2 }}>
                        <Typography component="span" variant="subtitle2">
                            Uploaded two file on &nbsp;
                            <Typography component="span" variant="h6">
                                21 Jan 2020
                            </Typography>
                        </Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid container>
                            <Grid item xs={12}>
                                <Card
                                    sx={{
                                        backgroundColor:
                                            theme.palette.mode === 'dark' ? theme.palette.dark.main : theme.palette.secondary.light
                                    }}
                                >
                                    <CardContent>
                                        <Grid container direction="column">
                                            <Grid item xs={12}>
                                                <Stack direction="row" spacing={2}>
                                                    <IconPhoto stroke={1.5} size="20px" />
                                                    <Typography variant="subtitle1">demo.jpg</Typography>
                                                </Stack>
                                            </Grid>
                                        </Grid>
                                    </CardContent>
                                </Card>
                            </Grid>
                        </Grid>
                    </Grid>
                </Grid>
            </ListItemWrapper>
            <Divider />
            <ListItemWrapper>
                <ListItem alignItems="center">
                    <ListItemAvatar>
                        <Avatar alt="John Doe" src={User1} />
                    </ListItemAvatar>
                    <ListItemText primary={<Typography variant="subtitle1">John Doe</Typography>} />
                    <ListItemSecondaryAction>
                        <Grid container justifyContent="flex-end">
                            <Grid item xs={12}>
                                <Typography variant="caption" display="block" gutterBottom>
                                    2 min ago
                                </Typography>
                            </Grid>
                        </Grid>
                    </ListItemSecondaryAction>
                </ListItem>
                <Grid container direction="column" className="list-container">
                    <Grid item xs={12} sx={{ pb: 2 }}>
                        <Typography variant="subtitle2">It is a long established fact that a reader will be distracted</Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid container>
                            <Grid item>
                                <Chip label="Confirmation of Account." sx={chipSuccessSX} />
                            </Grid>
                        </Grid>
                    </Grid>
                </Grid>
            </ListItemWrapper> */}
        </List>
    );
};

export default NotificationList;
