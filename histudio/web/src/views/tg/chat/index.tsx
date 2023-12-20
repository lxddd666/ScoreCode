import { memo, useEffect, useMemo } from 'react';
import { useState, useRef } from 'react';
import Avatar from '@mui/material/Avatar';
import { deepOrange } from '@mui/material/colors';
import ChatBubbleOutlineIcon from '@mui/icons-material/ChatBubbleOutline';
import Paper from '@mui/material/Paper';
import InputBase from '@mui/material/InputBase';
import IconButton from '@mui/material/IconButton';
import SearchIcon from '@mui/icons-material/Search';
import DragIndicatorIcon from '@mui/icons-material/DragIndicator';
import TextField from '@mui/material/TextField';
import FilterRoundedIcon from '@mui/icons-material/FilterRounded';
import SendRoundedIcon from '@mui/icons-material/SendRounded';
import AttachFileIcon from '@mui/icons-material/AttachFile';
import CampaignIcon from '@mui/icons-material/Campaign';
import KeyboardVoiceIcon from '@mui/icons-material/KeyboardVoice';
import { openSnackbar } from 'store/slices/snackbar';

import axios from 'utils/axios';
import { useDispatch, useSelector, shallowEqual } from 'store';

import { useParams } from "react-router-dom";
import { getTgArtsFoldersAction, getTgFoldersMessageAction, getTgFoldersMeeageHistoryAction } from 'store/slices/tg';
import styles from './index.module.scss';

import MessageBody from './messageBody';



const Chat = () => {
    const [loginInfo, setLoginInfo] = useState<any>({})
    const [tabsIndex, setTabsIndex] = useState(0); // tab index
    const [textValue, setTextValue] = useState('') // 发送消息
    const [mockMessageList, setMockMessageList] = useState<any>([ // 历史消息
        {
            id: String(new Date().getTime()),
            msg: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
            time: new Date(),
            flag: 'o'
        },
        {
            id: String(new Date().getTime()),
            msg: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
            time: new Date(),
            flag: 'm'
        },
        {
            id: String(new Date().getTime()),
            msg: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
            time: new Date(),
            flag: 'o'
        },
        {
            id: String(new Date().getTime()),
            msg: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
            time: new Date(),
            flag: 'o'
        }
    ]);
    const [tabsSideList, setTabsSideList] = useState<any>([ // 会话分组
        // { id: 0, name: '消息' },
        // { id: 1, name: '群组' },
        // { id: 2, name: '群聊' },
        // { id: 3, name: '个人' }
    ])
    const [tabsUserList, setTabsUserList] = useState([]) // 会话分组列表
    const textRef: any = useRef();
    const dispatch = useDispatch();
    const { tgArtsFolders, tgFoldersMessageList, tgFoldersMeeageHistoryList } = useSelector((state) => state.tg, shallowEqual)
    const { id } = useParams()

    // tg 登录
    useEffect(() => {
        tgArtsLogin()
    }, [])
    useEffect(() => {

        if (loginInfo.phone && loginInfo.phone !== '') {
            console.log('dispatch', loginInfo);
            dispatch(getTgArtsFoldersAction({ account: loginInfo.phone }))
            dispatch(getTgFoldersMessageAction({ account: loginInfo.phone }))
        }
    }, [dispatch, loginInfo.phone])
    useEffect(() => {
        console.log('111');
        setTabsSideList(tgArtsFolders?.data?.Elems || [])
    }, [tgArtsFolders])
    const memoizedList = useMemo(() => tgFoldersMessageList?.data?.list || [], [tgFoldersMessageList]);
    useEffect(() => {
        console.log('222', tabsUserList);
        setTabsUserList(memoizedList)
    }, [tabsUserList])
    useEffect(() => {
        console.log('333');
        setMockMessageList(tgFoldersMeeageHistoryList?.data?.list || [])
    }, [tgFoldersMeeageHistoryList])
    const tgArtsLogin = async () => {
        console.log('111');

        try {
            const { data } = await axios.post('/tg/arts/login', {
                id: id
            })
            console.log('routes', id, data);
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
                message: data?.data?.comment || '出错了~~~',
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




        } catch (error) {
            console.log('登陆错误', error);

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
            out: false,
            message: textRef?.current?.value
        }

        let tempList = [...mockMessageList, obj];

        setMockMessageList(tempList);

        try {
            const res = await axios.post('/tg/arts/sendMsg', {
                account: loginInfo.phone,
                receiver: loginInfo.tgId,
                textMsg: textMsg.push(textRef?.current?.value)
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
        dispatch(getTgFoldersMeeageHistoryAction({
            account: loginInfo.phone,
            contact: loginInfo.tgId,
            // offsetId: item.topMessage + 1,
        }))
    }
    return (
        <div className={styles.chat}>
            <div className={styles.side}>
                <div className={styles.avatars}>
                    <Avatar
                        sx={{ bgcolor: deepOrange[500], width: 60, height: 60 }}
                        variant="rounded"
                        src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"
                    >
                        N
                    </Avatar>
                </div>
                <div className={styles.list}>
                    {tabsSideList && tabsSideList.map((item: any, index: Number) => (
                        <div
                            key={item.ID}
                            className={`${styles.item} ${tabsIndex === index ? styles.itemActive : ''} `}
                            onClick={(event) => onTabsClick(index)}
                        >
                            <ChatBubbleOutlineIcon />
                            <span style={{ marginTop: '5px' }}>{item.Title || '全部消息'}</span>
                        </div>
                    ))}
                </div>
            </div>
            <div className={styles.userList}>
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
                    <DragIndicatorIcon className={styles.dragin} />
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
            <div className={styles.messageBody}>
                <div className={styles.messageTop}>
                    <div>
                        <Avatar alt=" " src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png" />
                    </div>
                    <div className={styles.messageTopRight}>
                        <div className={styles.item}>
                            <FilterRoundedIcon style={{ fontSize: '18px' }} />
                        </div>
                        <div className={styles.item}>
                            <AttachFileIcon style={{ fontSize: '18px' }} />
                        </div>
                        <div className={styles.item}>
                            <KeyboardVoiceIcon style={{ fontSize: '18px' }} />
                        </div>
                        <div className={styles.item}>
                            <CampaignIcon style={{ fontSize: '18px' }} />
                        </div>
                    </div>
                </div>
                <div className={styles.messageBodyInfo}>
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
                        <div className={styles.item}>
                            <KeyboardVoiceIcon style={{ fontSize: '15px' }} />
                        </div>
                        <div className={styles.item}>
                            <CampaignIcon style={{ fontSize: '15px' }} />
                        </div>
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
        </div>
    );
};

export default memo(Chat);
