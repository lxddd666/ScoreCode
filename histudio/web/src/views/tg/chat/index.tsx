import { memo, useEffect, useMemo } from 'react';
import { useState, useRef } from 'react';
import Avatar from '@mui/material/Avatar';
// import { deepOrange } from '@mui/material/colors';
// import ChatBubbleOutlineIcon from '@mui/icons-material/ChatBubbleOutline';
import Paper from '@mui/material/Paper';
import InputBase from '@mui/material/InputBase';
import IconButton from '@mui/material/IconButton';
import SearchIcon from '@mui/icons-material/Search';
// import DragIndicatorIcon from '@mui/icons-material/DragIndicator';
import TextField from '@mui/material/TextField';
import FilterRoundedIcon from '@mui/icons-material/FilterRounded';
import SendRoundedIcon from '@mui/icons-material/SendRounded';
import AttachFileIcon from '@mui/icons-material/AttachFile';
// import CampaignIcon from '@mui/icons-material/Campaign';
import AutoStoriesIcon from '@mui/icons-material/AutoStories';
import BallotIcon from '@mui/icons-material/Ballot';
import AssignmentIndIcon from '@mui/icons-material/AssignmentInd';
import ReadMoreIcon from '@mui/icons-material/ReadMore';
// import KeyboardVoiceIcon from '@mui/icons-material/KeyboardVoice';
import { openSnackbar } from 'store/slices/snackbar';
import BlurOnIcon from '@mui/icons-material/BlurOn';
import { styled } from '@mui/material/styles';
import {
    Badge,
    Button,
    Stack,
    Divider,
    List,
    ListItem,
    ListItemAvatar,
    ListItemText,
    Typography,
    Tooltip
} from '@mui/material';

import axios from 'utils/axios';
import { useDispatch, useSelector, shallowEqual } from 'store';
// import { getUserList } from 'store/slices/user';
import { getTgUserListAction } from 'store/slices/tg';

import { useParams } from "react-router-dom";
import { getTgArtsFoldersAction, getTgFoldersMessageAction, getTgFoldersMeeageHistoryAction } from 'store/slices/tg';
import styles from './index.module.scss';

import MessageBody from './messageBody';
import UserDetails from './UserDetails'

import { useScroll, timeAgo } from 'utils/tools';



const Chat = () => {
    const [loginInfo, setLoginInfo] = useState<any>({})
    const [tabsIndex, setTabsIndex] = useState(0); // tab index
    const [textValue, setTextValue] = useState('') // 发送消息
    const [mockMessageList, setMockMessageList] = useState<any>([ // 历史消息
        {
            msgId: 1,
            message: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
            time: new Date(),
            out: true
        },
        {
            msgId: 2,
            message: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
            time: new Date(),
            out: false
        },
        {
            msgId: 3,
            message: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
            time: new Date(),
            out: true
        },
        {
            msgId: 4,
            message: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
            time: new Date(),
            out: true
        }
    ]);
    const [tabsSideList, setTabsSideList] = useState<any>([ // 会话分组
        // { id: 0, name: '消息' },
        // { id: 1, name: '群组' },
        // { id: 2, name: '群聊' },
        // { id: 3, name: '个人' }
    ])
    const [tabsUserList, setTabsUserList] = useState([]) // 会话分组列表
    const [inputDisable, setInputDisable] = useState(false) // 输入框 disable
    const [showHideUserList, setShowHideUserList] = useState(true)
    const [showHideHistoryList, setShowHideHistoryList] = useState(true)
    const [showHideUserInfo, setShowHideUserInfo] = useState(false)
    const [currentUserHistoryTgId, setCurrentUserHistoryTgId] = useState('')
    const [userLists, setUserLists] = useState<any>([])
    const [ids, setIds] = useState('')
    const textRef: any = useRef();
    const divRef: any = useRef(null);
    const dispatch = useDispatch();
    const { tgArtsFolders, tgFoldersMessageList, tgFoldersMeeageHistoryList } = useSelector((state) => state.tg, shallowEqual)
    // const { userList } = useSelector((state) => state.user);
    const { tgUserList } = useSelector((state) => state.tg, shallowEqual);
    const { id } = useParams()
    // console.log('dispatchdispatchdispatch', tgFoldersMessageList);
    const { scrollInfo } = useScroll(divRef)

    // divRef.current = divRef?.current?.scrollHeight;
    useEffect(() => {
        console.log('scrollInfo', scrollInfo, divRef?.current?.scrollHeight);
        divRef.current.scrollTop = scrollInfo.scrollHeight;
    }, [mockMessageList, divRef?.current?.scrollHeight]);
    // tg 登录
    useEffect(() => {
        tgArtsLogin()
        // fetchBindUser()
        getTableListActionFN()
    }, [])
    // 获取分类及消息队列
    useEffect(() => {
        if (loginInfo.phone && loginInfo.phone !== '') {
            // console.log('dispatch', loginInfo);
            dispatch(getTgArtsFoldersAction({ account: loginInfo.phone }))
            dispatch(getTgFoldersMessageAction({ account: loginInfo.phone }))
            setInputDisable(true)
        }
    }, [dispatch, loginInfo.phone])
    // 获取后台用户列表
    // const fetchBindUser = async (page = 1, value: any = undefined,) => {
    //     const updatedParams = [
    //         `page=${page}`,
    //         `pageSize=${9999}`,
    //         `${value ? `username=${value}` : ''}`,
    //     ];
    //     await dispatch(getUserList(updatedParams.filter((query) => !query.endsWith('=') && query !== 'status=0').join('&')))
    // }
    // tg 账号
    const getTableListActionFN = async () => {
        await dispatch(getTgUserListAction({
            page: 1,
            pageSize: 9999,
        }));
    };

    useEffect(() => {
        console.log('111', tabsSideList);
        setTabsSideList(tgArtsFolders?.data?.Elems || [])
    }, [tgArtsFolders])
    let memoizedList = useMemo(() => tgFoldersMessageList?.data?.list || [], [tgFoldersMessageList]);
    useEffect(() => {
        console.log('222', tabsUserList);
        setTabsUserList(memoizedList)
    }, [tabsUserList, memoizedList])
    useEffect(() => {
        console.log('333', tgFoldersMeeageHistoryList?.data?.list);
        divRef.current.scrollTop = 20
        setMockMessageList(tgFoldersMeeageHistoryList || [])
    }, [tgFoldersMeeageHistoryList])
    // useEffect(() => {
    //     if (userList?.data?.list) {
    //         setUserLists(userList?.data?.list || [])
    //     }
    // }, [userList])
    useEffect(() => {
        if (tgUserList?.data?.list) {
            setUserLists(tgUserList?.data?.list || [])
        }
    }, [tgUserList])
    useEffect(() => {
        if (ids && ids !== '') {
            tgArtsLogin(ids)
        }
    }, [ids])

    const tgArtsLogin = async (ids: any = undefined) => {
        console.log('111');

        try {
            let idV: any = ids ? ids : id
            const { data } = await axios.post('/tg/arts/login', {
                id: idV
            })
            console.log('routes', idV, data);
            if (data.code !== 0) {
                return dispatch(openSnackbar({
                    open: true,
                    message: data?.message || '出错了~~~',
                    variant: 'alert',
                    alert: {
                        color: 'error'
                    },
                    close: false,
                    anchorOrigin: {
                        vertical: 'top',
                        horizontal: 'center'
                    }
                }))
            }
            setLoginInfo(data.data)

            dispatch(openSnackbar({
                open: true,
                message: data?.data?.comment || data?.message || '登录成功~~~',
                variant: 'alert',
                alert: {
                    color: 'success'
                },
                close: false,
                anchorOrigin: {
                    vertical: 'top',
                    horizontal: 'center'
                }
            }))
        } catch (error: any) {
            console.log('登陆错误', error);
            dispatch(openSnackbar({
                open: true,
                message: error?.message || '登陆错误~~~',
                variant: 'alert',
                alert: {
                    color: 'error'
                },
                close: false,
                anchorOrigin: {
                    vertical: 'top',
                    horizontal: 'center'
                }
            }))
        }
        // dispatch(getTgArtsFoldersAction({ account: loginInfo.phone }))
        console.log('111');
    }

    // 会话 点击
    const onTabsClick = (index: any): any => {
        setTabsIndex(index);
    };
    // 消息输入
    const onTextValueChange = (event: any) => {
        setTextValue(event.target.value)
    }
    // 发送消息
    const onSendSubmit = async () => {
        if (!inputDisable) {
            return
        }
        // console.log('提交', textRef,textValue);
        // let obj = {
        //     id: String(new Date().getTime()),
        //     msg: textRef?.current?.value,
        //     flag: 'm'
        // };
        let textMsg = []
        textMsg.push(textRef?.current?.value);
        let obj = {
            msgId: String(new Date().getTime()),
            out: true,
            message: textRef?.current?.value
        }

        let tempList = [...mockMessageList, obj];

        setMockMessageList(tempList);

        try {
            const res = await axios.post('/tg/arts/sendMsg', {
                account: loginInfo.phone,
                receiver: currentUserHistoryTgId,
                // textMsg: textMsg.push(textRef?.current?.value)
                textMsg: textMsg
            })
            console.log('res', res);

            // if (res.code === 0) {

            // }
            setTextValue('')
        } catch (error) {
            tempList.pop()
            setMockMessageList(tempList);
        }
    };

    // 用户聊天点击
    const onUserClick = (item: any) => {
        console.log('111', item);
        setCurrentUserHistoryTgId(item.tgId)
        dispatch(getTgFoldersMeeageHistoryAction({
            account: loginInfo.phone,
            contact: item.tgId,
            offsetId: item.topMessage + 1,
        }))
        // console.log('scrollInfo', scrollInfo, divRef?.current?.scrollHeight);
    }

    // 员工列表显示隐藏
    const onShowHideUserListClick = () => {
        setShowHideUserList(!showHideUserList)
    }
    // history 列表显示隐藏
    const onShowHideHistoryListClick = () => {
        setShowHideHistoryList(!showHideHistoryList)
    }
    // history 列表显示隐藏
    const onShowHideUserInfoClick = () => {
        setShowHideUserInfo(!showHideUserInfo)
    }
    // tg 账号点击
    const tgClick = (item: any) => {
        console.log('item', item);
        setIds(item.id)
    }
    return (
        <div className={styles.chat}>
            {
                showHideUserList && <div className={styles.side}>
                    <div className={styles.avatars}>
                        <div className={styles.avata}>
                            <StyledBadge
                                overlap="circular"
                                anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
                                variant="dot"
                                badgeColor={loginInfo?.isOnline === 1 ? '#44b700' : 'red'}
                            >
                                <Avatar alt="Remy Sharp" src={loginInfo?.photo}>
                                    {loginInfo?.lastName?.charAt(0)?.toUpperCase()}
                                </Avatar>
                            </StyledBadge>
                            <div style={{ display: 'flex', alignItems: 'center', marginLeft: '5px' }}>
                                {loginInfo?.username || '未知'}
                            </div>
                        </div>

                        <IconButton aria-label="delete" size="small">
                            <BlurOnIcon fontSize="small" />
                        </IconButton>
                    </div>
                    <Divider />
                    <div className={styles.list}>
                        {userLists && userLists.map((item: any, index: Number) => (
                            <div
                                key={item.id}
                                className={`${styles.item} ${tabsIndex === index ? styles.itemActive : ''} `}
                                onClick={(event) => onTabsClick(index)}
                            >
                                <List sx={{ width: '100%', maxWidth: 300, bgcolor: 'background.paper' }}>

                                    <ListItem alignItems="flex-start"
                                        secondaryAction={
                                            <Tooltip title="切换账号" placement="top">
                                                <IconButton color="primary" aria-label="upload picture" component="span" onClick={e => tgClick(item)}>
                                                    <ReadMoreIcon />
                                                </IconButton>
                                            </Tooltip>
                                        }>
                                        <ListItemAvatar>
                                            <StyledBadge
                                                overlap="circular"
                                                anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
                                                variant="dot"
                                                badgeColor={item?.isOnline === 1 ? '#44b700' : 'red'}
                                            >
                                                <Avatar alt="Remy Sharp" src={item?.avatar}>
                                                    {item?.username?.charAt(0)?.toUpperCase()}
                                                </Avatar>
                                            </StyledBadge>
                                        </ListItemAvatar>
                                        <ListItemText
                                            primary={<span style={{ maxWidth: '150px', textOverflow: 'ellipsis', overflow: 'hidden', display: 'inline-block' }}>{item?.firstName} {item?.lastName}</span>}
                                            secondary={
                                                <>
                                                    <Typography
                                                        sx={{ display: 'inline' }}
                                                        component="span"
                                                        variant="body2"
                                                        color="text.primary"
                                                    >
                                                        (+1){item?.phone || '-'}
                                                    </Typography>
                                                </>
                                            }
                                        />
                                    </ListItem>
                                </List>
                                {/* <ChatBubbleOutlineIcon />
                                <span style={{ marginTop: '5px' }}>{item.Title || '全部消息'}</span> */}
                            </div>
                        ))}
                    </div>
                </div>
            }
            {
                showHideHistoryList && <div className={styles.userList}>
                    <div className={styles.ipt}>
                        <Paper
                            component="form"
                            sx={{
                                p: '2px 4px',
                                display: 'flex',
                                alignItems: 'center',
                                width: '85%',
                                border: '1px solid rgb(127, 127, 127)'
                            }}
                        >
                            <InputBase sx={{ ml: 1, flex: 1 }} placeholder="请输入" inputProps={{ 'aria-label': 'search google maps' }} />
                            <IconButton type="button" sx={{ p: '10px' }} aria-label="search">
                                <SearchIcon />
                            </IconButton>
                        </Paper>
                        <BlurOnIcon className={styles.dragin} onClick={onShowHideUserListClick} />
                    </div>
                    <div className={styles.list}>
                        {tabsUserList && tabsUserList.map((item: any, index: Number): any => {
                            return (
                                <div className={styles.item} key={item.id} onClick={e => onUserClick(item)}>
                                    <div>
                                        <Avatar
                                            variant="rounded"
                                            src="https://gw.alipayobjects.com/zos/antfincdn/x43I27A55%26/photo-1438109491414-7198515b166b.webp"
                                        >
                                            N
                                        </Avatar>
                                    </div>
                                    <div className={styles.lineFont}>
                                        <div className={styles.name}>{(item.firstName !== '' && item.firstName + ' ' + item.lastName) || (item.username !== '' && item.username) || item.title || '~'}</div>
                                        <div className={styles.line}>{item.last.message || ''}</div>
                                    </div>
                                </div>
                            );
                        })}
                    </div>
                </div>
            }
            <div className={styles.messageBody}>
                <div className={styles.messageTop}>
                    <div className={styles.messageTopLeft}>
                        <BlurOnIcon style={{ marginRight: '10px' }} onClick={onShowHideHistoryListClick} />
                        <Avatar alt=" " src="" >N</Avatar>
                        <div className={styles.name}>
                            <Badge color="secondary" variant="dot" invisible={true}>
                                <div>{loginInfo.firstName} {loginInfo.lastName}</div>
                            </Badge>

                            <div className={styles.userD}>
                                <div style={{ fontSize: '12px' }}>@{loginInfo.username || '未知'}</div>
                                <div style={{ fontSize: '12px', marginLeft: '10px' }}>{timeAgo(loginInfo.lastLoginTime)}</div>
                            </div>
                        </div>
                    </div>
                    <div className={styles.messageTopRight}>
                        {/* <div className={styles.item}>
                            <FilterRoundedIcon style={{ fontSize: '18px' }} />
                        </div> */}
                        {/* <div className={styles.item}>
                            <AttachFileIcon style={{ fontSize: '18px' }} />
                        </div> */}
                        {/* <div className={styles.item}>
                            <KeyboardVoiceIcon style={{ fontSize: '18px' }} />
                        </div> */}
                        {/* <div >
                                <AssignmentIndIcon style={{ fontSize: '18px' }} />
                            </div> */}
                        <Stack direction="row" spacing={2}>

                            {/* <div className={styles.item} style={{ fontSize: '18px', width: "120px" }}>
                            聊天素材
                        </div> */}
                            <Button variant="outlined" startIcon={<AutoStoriesIcon />}>
                                自定义话术
                            </Button>
                            <Button variant="outlined" startIcon={<BallotIcon />}>
                                话术库
                            </Button>
                            <Button variant="outlined" onClick={onShowHideUserInfoClick} style={{ display: 'flex', justifyContent: 'flex-end' }} startIcon={<AssignmentIndIcon />} />
                        </Stack>
                    </div>
                </div>
                <div className={styles.messageBodyInfo} ref={divRef}>
                    <MessageBody messageList={mockMessageList} key={Math.random() * 10} />
                </div>
                <div className={styles.messageSend}>
                    <div className={styles.utilsInfo}>
                        <div className={styles.item}>
                            <FilterRoundedIcon style={{ fontSize: '15px' }} />
                        </div>
                        <div className={styles.item}>
                            <AttachFileIcon style={{ fontSize: '15px' }} />
                        </div>
                        {/* <div className={styles.item}>
                            <KeyboardVoiceIcon style={{ fontSize: '15px' }} />
                        </div> */}
                        {/* <div className={styles.item}>
                            <CampaignIcon style={{ fontSize: '15px' }} />
                        </div> */}
                    </div>
                    <div className={styles.send}>
                        {/* <Paper
                            component="form"
                            sx={{
                                p: '2px 4px',
                                display: 'flex',
                                alignItems: 'center',
                                width: '90%',
                                border: '1px solid rgb(127, 127, 127)'
                            }}
                        > */}
                        <TextField
                            id="outlined-multiline-flexible"
                            placeholder="输入格式:/话术关键 弹出话术选择框"
                            multiline
                            minRows={3}
                            maxRows={3}
                            style={{ width: '90%' }}
                            inputRef={textRef}
                            value={textValue}
                            onChange={onTextValueChange}
                        />
                        {/* </Paper> */}
                        <div className={styles.dragin}>
                            <SendRoundedIcon color={'primary'} onClick={onSendSubmit} />
                        </div>
                    </div>
                </div>
            </div>

            {
                showHideUserInfo && <UserDetails user={loginInfo} />
            }
        </div>
    );
};
interface StyledBadgeProps {
    badgeColor?: string; // 这是你的自定义属性
}
const StyledBadge = styled(Badge)<StyledBadgeProps>(({ theme, badgeColor }) => ({
    '& .MuiBadge-badge': {
        backgroundColor: badgeColor || '#44b700',
        color: badgeColor || '#44b700',
        boxShadow: `0 0 0 2px ${theme.palette.background.paper}`,
        '&::after': {
            position: 'absolute',
            top: 0,
            left: 0,
            width: '100%',
            height: '100%',
            borderRadius: '50%',
            animation: 'ripple 1.2s infinite ease-in-out',
            border: '1px solid currentColor',
            content: '""',
        },
    },
    '@keyframes ripple': {
        '0%': {
            transform: 'scale(.8)',
            opacity: 1,
        },
        '100%': {
            transform: 'scale(2.4)',
            opacity: 0,
        },
    },
}));
export default memo(Chat);
