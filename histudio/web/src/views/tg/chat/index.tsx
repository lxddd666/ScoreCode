import { memo } from 'react';
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

import styles from './index.module.scss';

import MessageBody from './messageBody';

const Chat = () => {
    const [tabsIndex, setTabsIndex] = useState(0);
    const [textValue,setTextValue] = useState('')
    const [mockMessageList, setMockMessageList] = useState<any>([
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
    const textRef: any = useRef();
    const tabsSideList: any = [
        { id: 0, name: '消息' },
        { id: 1, name: '群组' },
        { id: 2, name: '群聊' },
        { id: 3, name: '个人' }
    ];
    const tabsUserList: any = [
        { id: 0, name: '消息1111111111' },
        { id: 1, name: '群组2222222222' },
        { id: 2, name: '群聊3333333333' },
        { id: 3, name: '个人4444444444' }
    ];
    const onTabsClick = (index: any): any => {
        setTabsIndex(index);
    };
    const onTextValueChange = (event:any) => {
        setTextValue(event.target.value)
    }
    const onSendSubmit = () => {
        // console.log('提交', textRef,textValue);
        let obj = {
            id: String(new Date().getTime()),
            msg: textRef?.current?.value,
            flag: 'm'
        };
        let tempList = [...mockMessageList, obj];
        // console.log(tempList);

        setMockMessageList(tempList);
        setTextValue('')
    };
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
                    {tabsSideList.map((item: any, index: Number) => (
                        <div
                            key={item.id}
                            className={`${styles.item} ${tabsIndex === index ? styles.itemActive : ''} `}
                            onClick={(event) => onTabsClick(index)}
                        >
                            <ChatBubbleOutlineIcon />
                            <span style={{ marginTop: '5px' }}>{item.name}</span>
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
                    {tabsUserList.map((item: any, index: Number): any => {
                        return (
                            <div className={styles.item} key={item.id}>
                                <div>
                                    <Avatar
                                        variant="rounded"
                                        src="https://gw.alipayobjects.com/zos/antfincdn/x43I27A55%26/photo-1438109491414-7198515b166b.webp"
                                    >
                                        N
                                    </Avatar>
                                </div>
                                <div className={styles.lineFont}>{item.name}</div>
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
                    <MessageBody messageList={mockMessageList} />
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
