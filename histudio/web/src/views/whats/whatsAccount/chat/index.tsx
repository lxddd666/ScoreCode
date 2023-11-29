import React from 'react';
import { useIntl } from 'react-intl';
import { useParams } from 'react-router-dom';

// ui component
import MainCard from 'ui-component/cards/MainCard';
import Loader from 'ui-component/Loader';

// Lib
import PerfectScrollbar from 'react-perfect-scrollbar';
import EmojiPicker, { SkinTones, EmojiClickData } from 'emoji-picker-react';

// store
import { drawerWidth, gridSpacing } from 'store/constant';
import { useDispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';

// MUI
import {
    Avatar,
    Box,
    Card,
    CardContent,
    CardMedia,
    ClickAwayListener,
    Divider,
    Grid,
    IconButton,
    InputAdornment,
    Menu,
    MenuItem,
    OutlinedInput,
    Popper,
    styled,
    Theme,
    Typography,
    useMediaQuery,
    useTheme
} from '@mui/material';

import AttachmentTwoToneIcon from '@mui/icons-material/AttachmentTwoTone';
import MenuRoundedIcon from '@mui/icons-material/MenuRounded';
import MoreHorizTwoToneIcon from '@mui/icons-material/MoreHorizTwoTone';
import ErrorTwoToneIcon from '@mui/icons-material/ErrorTwoTone';
// import VideoCallTwoToneIcon from '@mui/icons-material/VideoCallTwoTone';
// import CallTwoToneIcon from '@mui/icons-material/CallTwoTone';
import SendTwoToneIcon from '@mui/icons-material/SendTwoTone';
import MoodTwoToneIcon from '@mui/icons-material/MoodTwoTone';
import ChatDrawer from './drawer';
import HighlightOffTwoToneIcon from '@mui/icons-material/HighlightOffTwoTone';
import AvatarStatus from './AvatarStatus';
import SubCard from 'ui-component/cards/SubCard';
import PinDropTwoToneIcon from '@mui/icons-material/PinDropTwoTone';
import PhoneTwoToneIcon from '@mui/icons-material/PhoneTwoTone';
import EmailTwoToneIcon from '@mui/icons-material/EmailTwoTone';

import envRef from 'environment';
import axiosServices from 'utils/axios';

import { defaultErrorMessage } from 'constant/general';

const Main = styled('main', { shouldForwardProp: (prop: string) => prop !== 'open' })(
    ({ theme, open }: { theme: Theme; open: boolean }) => ({
        flexGrow: 1,
        paddingLeft: open ? theme.spacing(3) : 0,
        transition: theme.transitions.create('margin', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.shorter
        }),
        marginLeft: `-${drawerWidth}px`,
        [theme.breakpoints.down('lg')]: {
            paddingLeft: 0,
            marginLeft: 0
        },
        ...(open && {
            transition: theme.transitions.create('margin', {
                easing: theme.transitions.easing.easeOut,
                duration: theme.transitions.duration.shorter
            }),
            marginLeft: 0
        })
    })
);

const Chat = () => {
    const { id } = useParams();
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();

    const matchDownSM = useMediaQuery(theme.breakpoints.down('lg'));

    const [userData, setUserData] = React.useState<any>(null);

    const [anchorEl, setAnchorEl] = React.useState<Element | ((element: Element) => Element) | null | undefined>(null);
    const handleClickSort = (event: React.MouseEvent<HTMLButtonElement> | undefined) => {
        setAnchorEl(event?.currentTarget);
    };

    const handleCloseSort = () => {
        setAnchorEl(null);
    };

    const scrollRef = React.useRef();

    React.useLayoutEffect(() => {
        if (scrollRef?.current) {
            // @ts-ignore
            scrollRef.current.scrollIntoView();
        }
    });

    // toggle sidebar
    const [openChatDrawer, setOpenChatDrawer] = React.useState(true);
    const handleDrawerOpen = () => {
        setOpenChatDrawer((prevState) => !prevState);
    };

    // close sidebar when widow size below 'md' breakpoint
    React.useEffect(() => {
        setOpenChatDrawer(!matchDownSM);
    }, [matchDownSM]);

    React.useEffect(() => {
        if (id) getData(parseInt(id));
    }, []);

    async function getData(id: number) {
        await axiosServices
            .get(`${envRef?.API_URL}whats/whatsAccount/view?id=${id}`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setUserData(response.data.data);
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

    // handle message to send
    const [message, setMessage] = React.useState<string>('');
    const handleOnSend = () => {
        // const d = new Date();
        setMessage('');
        // const newMessage = {
        //     from: 'User1',
        //     to: user.name,
        //     text: message,
        //     time: d.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
        // };
        // setData((prevState) => [...prevState, newMessage]);
    };

    const [isUserInfoShow, setIsUserInfoShow] = React.useState<boolean>(false);

    const handleUserInfoShow = () => {
        setIsUserInfoShow(!isUserInfoShow);
    };

    // enter to send message
    const handleEnter = (event: React.KeyboardEvent<HTMLDivElement> | undefined) => {
        if (event?.key !== 'Enter') {
            return;
        }
        handleOnSend();
    };

    // handle emoji
    const onEmojiClick = (emojiObject: EmojiClickData, event: MouseEvent) => {
        setMessage((prev) => prev + emojiObject.emoji);
    };

    const [anchorElEmoji, setAnchorElEmoji] = React.useState<any>(); /** No single type can cater for all elements */
    const handleOnEmojiButtonClick = (event: React.MouseEvent<HTMLButtonElement> | undefined) => {
        setAnchorElEmoji(anchorElEmoji ? null : event?.currentTarget);
    };

    const emojiOpen = Boolean(anchorElEmoji);
    const emojiId = emojiOpen ? 'simple-popper' : undefined;
    const handleCloseEmoji = () => {
        setAnchorElEmoji(null);
    };

    if (!userData) return <Loader />;

    return (
        <MainCard title={intl.formatMessage({ id: 'general.chat' })}>
            <Box sx={{ display: 'flex' }}>
                <ChatDrawer openChatDrawer={openChatDrawer} handleDrawerOpen={handleDrawerOpen} userData={userData} />
                <Main theme={theme} open={openChatDrawer}>
                    <Grid container spacing={gridSpacing}>
                        <Grid item xs zeroMinWidth /* sx={{ display: emailDetails ? { xs: 'none', sm: 'flex' } : 'flex' }} */>
                            <MainCard
                                sx={{
                                    bgcolor: theme.palette.mode === 'dark' ? 'dark.main' : 'grey.50'
                                }}
                            >
                                <Grid container spacing={gridSpacing}>
                                    <Grid item xs={12}>
                                        <Grid container alignItems="center" spacing={0.5}>
                                            <Grid item>
                                                <IconButton onClick={handleDrawerOpen} size="large" aria-label="chat menu collapse">
                                                    <MenuRoundedIcon />
                                                </IconButton>
                                            </Grid>
                                            <Grid item>
                                                <Grid container spacing={2} alignItems="center" sx={{ flexWrap: 'nowrap' }}>
                                                    <Grid item>
                                                        <Avatar alt={`avatar_${userData.id || 0}`} src={userData.avatar || '-'} />
                                                    </Grid>
                                                    <Grid item sm zeroMinWidth>
                                                        <Grid container spacing={0} alignItems="center">
                                                            <Grid item xs={12}>
                                                                <Typography variant="h4" component="div">
                                                                    {userData.isOnline && (
                                                                        <AvatarStatus status={userData.isOnline || '1'} />
                                                                    )}
                                                                </Typography>
                                                            </Grid>
                                                            <Grid item xs={12}>
                                                                <Typography variant="subtitle2">{userData.lastLoginTime || '-'}</Typography>
                                                            </Grid>
                                                        </Grid>
                                                    </Grid>
                                                </Grid>
                                            </Grid>
                                            <Grid item sm zeroMinWidth />
                                            {/* <Grid item>
                                            <IconButton size="large" aria-label="chat user call">
                                                <CallTwoToneIcon />
                                            </IconButton>
                                        </Grid>
                                        <Grid item>
                                            <IconButton size="large" aria-label="chat user video call">
                                                <VideoCallTwoToneIcon />
                                            </IconButton>
                                        </Grid> */}
                                            <Grid item>
                                                <IconButton onClick={handleUserInfoShow} size="large" aria-label="chat user information">
                                                    <ErrorTwoToneIcon />
                                                </IconButton>
                                            </Grid>
                                            <Grid item>
                                                <IconButton onClick={handleClickSort} size="large" aria-label="chat user details change">
                                                    <MoreHorizTwoToneIcon />
                                                </IconButton>
                                                <Menu
                                                    id="simple-menu"
                                                    anchorEl={anchorEl}
                                                    keepMounted
                                                    open={Boolean(anchorEl)}
                                                    onClose={handleCloseSort}
                                                    anchorOrigin={{
                                                        vertical: 'bottom',
                                                        horizontal: 'right'
                                                    }}
                                                    transformOrigin={{
                                                        vertical: 'top',
                                                        horizontal: 'right'
                                                    }}
                                                >
                                                    <MenuItem onClick={handleCloseSort}>Name</MenuItem>
                                                    <MenuItem onClick={handleCloseSort}>Date</MenuItem>
                                                    <MenuItem onClick={handleCloseSort}>Ratting</MenuItem>
                                                    <MenuItem onClick={handleCloseSort}>Unread</MenuItem>
                                                </Menu>
                                            </Grid>
                                        </Grid>
                                        <Divider sx={{ mt: theme.spacing(2) }} />
                                    </Grid>
                                    <PerfectScrollbar
                                        style={{ width: '100%', height: 'calc(30vmax - 20px)', overflowX: 'hidden', minHeight: '20vmax' }}
                                    >
                                        <CardContent>
                                            {/* <ChartHistory theme={theme} user={user} data={data} /> */}
                                            {/* @ts-ignore */}
                                            <span ref={scrollRef} />
                                        </CardContent>
                                    </PerfectScrollbar>
                                    <Grid item xs={12}>
                                        <Grid container spacing={1} alignItems="center">
                                            <Grid item>
                                                <IconButton size="large" aria-label="attachment file">
                                                    <AttachmentTwoToneIcon />
                                                </IconButton>
                                                <IconButton
                                                    ref={anchorElEmoji}
                                                    aria-describedby={emojiId}
                                                    onClick={handleOnEmojiButtonClick}
                                                    size="large"
                                                    aria-label="emoji"
                                                >
                                                    <MoodTwoToneIcon />
                                                </IconButton>
                                            </Grid>
                                            <Grid item xs={12} sm zeroMinWidth>
                                                <OutlinedInput
                                                    id="message-send"
                                                    fullWidth
                                                    value={message}
                                                    onChange={(e) => setMessage(e.target.value)}
                                                    onKeyPress={handleEnter}
                                                    placeholder="Type a Message"
                                                    endAdornment={
                                                        <InputAdornment position="end">
                                                            <IconButton
                                                                disableRipple
                                                                color="primary"
                                                                onClick={handleOnSend}
                                                                aria-label="send message"
                                                            >
                                                                <SendTwoToneIcon />
                                                            </IconButton>
                                                        </InputAdornment>
                                                    }
                                                    aria-describedby="search-helper-text"
                                                    inputProps={{ 'aria-label': 'weight' }}
                                                />
                                            </Grid>
                                        </Grid>
                                    </Grid>
                                </Grid>
                                <Popper
                                    id={emojiId}
                                    open={emojiOpen}
                                    anchorEl={anchorElEmoji}
                                    disablePortal
                                    style={{ zIndex: 1200 }}
                                    modifiers={[
                                        {
                                            name: 'offset',
                                            options: {
                                                offset: [-20, 20]
                                            }
                                        }
                                    ]}
                                >
                                    <ClickAwayListener onClickAway={handleCloseEmoji}>
                                        <MainCard elevation={8} content={false}>
                                            <EmojiPicker
                                                onEmojiClick={onEmojiClick}
                                                defaultSkinTone={SkinTones.DARK}
                                                autoFocusSearch={false}
                                            />
                                        </MainCard>
                                    </ClickAwayListener>
                                </Popper>
                            </MainCard>
                        </Grid>

                        {isUserInfoShow && (
                            <Grid item sx={{ margin: { xs: '0 auto', md: 'initial' } }}>
                                <Box sx={{ display: { xs: 'block', sm: 'none', textAlign: 'right' } }}>
                                    <IconButton onClick={handleUserInfoShow} sx={{ mb: -5 }} size="large">
                                        <HighlightOffTwoToneIcon />
                                    </IconButton>
                                </Box>
                                <Grid container spacing={gridSpacing} sx={{ width: '100%', maxWidth: 300 }}>
                                    <Grid item xs={12}>
                                        <Card>
                                            <CardContent
                                                sx={{
                                                    textAlign: 'center',
                                                    background:
                                                        theme.palette.mode === 'dark'
                                                            ? theme.palette.dark.main
                                                            : theme.palette.primary.light
                                                }}
                                            >
                                                <Grid container spacing={1}>
                                                    <Grid item xs={12}>
                                                        <Avatar
                                                            alt={'contact.name'}
                                                            src={`contact.avatar`}
                                                            sx={{
                                                                m: '0 auto',
                                                                width: 130,
                                                                height: 130,
                                                                border: '1px solid',
                                                                borderColor: theme.palette.primary.main,
                                                                p: 1,
                                                                bgcolor: 'transparent'
                                                            }}
                                                        />
                                                    </Grid>
                                                    <Grid item xs={12}>
                                                        <AvatarStatus status={'contact.online_status'} />
                                                        <Typography variant="caption" component="div">
                                                            {'contact.online_status'}
                                                        </Typography>
                                                    </Grid>
                                                    <Grid item xs={12}>
                                                        <Typography variant="h5" component="div">
                                                            {'contact.name'}
                                                        </Typography>
                                                    </Grid>
                                                    <Grid item xs={12}>
                                                        <Typography variant="body2" component="div">
                                                            {'contact.role'}
                                                        </Typography>
                                                    </Grid>
                                                </Grid>
                                            </CardContent>
                                        </Card>
                                    </Grid>
                                    <Grid item xs={12}>
                                        <SubCard
                                            sx={{
                                                background: theme.palette.mode === 'dark' ? theme.palette.dark.main : theme.palette.grey[50]
                                            }}
                                        >
                                            <Grid container spacing={2}>
                                                <Grid item xs={12}>
                                                    <Typography variant="h5" component="div">
                                                        Information
                                                    </Typography>
                                                </Grid>
                                                <Grid item xs={12}>
                                                    <Grid container spacing={1}>
                                                        <Grid item xs={12}>
                                                            <Typography variant="body2">
                                                                <PinDropTwoToneIcon
                                                                    sx={{ verticalAlign: 'sub', fontSize: '1.125rem', mr: 0.625 }}
                                                                />{' '}
                                                                32188 Sips Parkways, U.S
                                                            </Typography>
                                                        </Grid>
                                                        <Grid item xs={12}>
                                                            <Typography variant="body2">
                                                                <PhoneTwoToneIcon
                                                                    sx={{ verticalAlign: 'sub', fontSize: '1.125rem', mr: 0.625 }}
                                                                />{' '}
                                                                995-250-1803
                                                            </Typography>
                                                        </Grid>
                                                        <Grid item xs={12}>
                                                            <Typography variant="body2">
                                                                <EmailTwoToneIcon
                                                                    sx={{ verticalAlign: 'sub', fontSize: '1.125rem', mr: 0.625 }}
                                                                />{' '}
                                                                O’Keefe@codedtheme.com
                                                            </Typography>
                                                        </Grid>
                                                    </Grid>
                                                </Grid>
                                                <Grid item xs={12}>
                                                    <Divider />
                                                </Grid>
                                                <Grid item xs={12}>
                                                    <Typography variant="h5" component="div">
                                                        Attachment
                                                    </Typography>
                                                </Grid>
                                                <Grid item xs={12}>
                                                    <Grid container spacing={1}>
                                                        <Grid item xs={12}>
                                                            <Grid container spacing={1}>
                                                                <Grid item>
                                                                    <CardMedia
                                                                        component="img"
                                                                        image={'images1'}
                                                                        title="image"
                                                                        sx={{ width: 42, height: 42 }}
                                                                    />
                                                                </Grid>
                                                                <Grid item xs zeroMinWidth>
                                                                    <Grid container spacing={0}>
                                                                        <Grid item xs={12}>
                                                                            <Typography variant="h6">File Name.jpg</Typography>
                                                                        </Grid>
                                                                        <Grid item xs={12}>
                                                                            <Typography variant="caption">4/16/2021 07:47:03</Typography>
                                                                        </Grid>
                                                                    </Grid>
                                                                </Grid>
                                                            </Grid>
                                                        </Grid>
                                                        <Grid item xs={12}>
                                                            <Grid container spacing={1}>
                                                                <Grid item>
                                                                    <CardMedia
                                                                        component="img"
                                                                        image={'images2'}
                                                                        title="image"
                                                                        sx={{ width: 42, height: 42 }}
                                                                    />
                                                                </Grid>
                                                                <Grid item xs zeroMinWidth>
                                                                    <Grid container spacing={0}>
                                                                        <Grid item xs={12}>
                                                                            <Typography variant="h6">File Name.ai</Typography>
                                                                        </Grid>
                                                                        <Grid item xs={12}>
                                                                            <Typography variant="caption">4/16/2021 07:47:03</Typography>
                                                                        </Grid>
                                                                    </Grid>
                                                                </Grid>
                                                            </Grid>
                                                        </Grid>
                                                        <Grid item xs={12}>
                                                            <Grid container spacing={1}>
                                                                <Grid item>
                                                                    <CardMedia
                                                                        component="img"
                                                                        image={'images3'}
                                                                        title="image"
                                                                        sx={{ width: 42, height: 42 }}
                                                                    />
                                                                </Grid>
                                                                <Grid item xs zeroMinWidth>
                                                                    <Grid container spacing={0}>
                                                                        <Grid item xs={12}>
                                                                            <Typography variant="h6">File Name.pdf</Typography>
                                                                        </Grid>
                                                                        <Grid item xs={12}>
                                                                            <Typography variant="caption">4/16/2021 07:47:03</Typography>
                                                                        </Grid>
                                                                    </Grid>
                                                                </Grid>
                                                            </Grid>
                                                        </Grid>
                                                    </Grid>
                                                </Grid>
                                            </Grid>
                                        </SubCard>
                                    </Grid>
                                </Grid>
                            </Grid>
                        )}
                    </Grid>
                </Main>
            </Box>
        </MainCard>
    );
};

export default Chat;
