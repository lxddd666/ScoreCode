import { memo, useState } from 'react';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import AutorenewIcon from '@mui/icons-material/Autorenew';
import SearchIcon from '@mui/icons-material/Search';
import styles from './searchForm.module.scss';
import { styled } from '@mui/material/styles';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import InputAdornment from '@mui/material/InputAdornment';
import MenuItem from '@mui/material/MenuItem';
import {readArr, sendStatusArr} from "./config";
// import { DatePicker } from '@mui/x-date-pickers';

// import { LocalizationProvider } from '@mui/x-date-pickers';
// import AdapterDateFns from '@mui/lab/AdapterDateFns';
// import AdapterDateFns from '@mui/x-date-pickers/AdapterDateFns';

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));

const SearchForm = (props: any) => {
    const { handleSearchFormData } = props;
    const [value, setValue] = useState<any>(null);
    const [formData, setFormData] = useState<any>({
        reqId: undefined,
        createdAt: undefined,
        initiator: undefined,
        sender: undefined,
        msgType: undefined,
        sendTime: undefined,
        read: undefined,
        receiver: undefined
    })

    // 搜索按钮
    const onSearchClick = (e: any) => {
        console.log(e.target.value, formData);
        let obj = { folderId: value?.value ? value?.value : undefined, ...formData }
        handleSearchFormData(obj);
    };
    // 重置按钮
    const onResetClick = (e: any) => {

        let obj = {
            reqId: undefined,
            createdAt: undefined,
            initiator: undefined,
            sender: undefined,
            msgType: undefined,
            sendTime: undefined,
            read: undefined,
            sendStatus: undefined,
            receiver: undefined
        }
        setValue({});
        setFormData(obj)
        handleSearchFormData(obj);
        // console.log(formData);

    };
    return (
        <>
            <div className={styles.searchForm}>
                <Grid container spacing={0.3} alignItems="center">

                    <Grid item >
                        <Item> <TextField

                            sx={{ width: 300 }}
                            autoFocus
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.initiator || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    initiator: event.target.value
                                })
                            }
                            label="请输入聊天发起人"
                            type="text"
                            variant="outlined"
                            size="small"
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon />
                                    </InputAdornment>
                                ),
                            }}
                        /></Item>
                    </Grid>
                    <Grid item >
                        <Item> <TextField

                            sx={{ width: 300 }}
                            autoFocus
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.sender || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    sender: event.target.value
                                })
                            }
                            label="请输入发送人"
                            type="text"
                            variant="outlined"
                            size="small"
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon />
                                    </InputAdornment>
                                ),
                            }}
                        /></Item>
                    </Grid>
                    <Grid item >
                        <Item> <TextField

                            sx={{ width: 300 }}
                            autoFocus
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.receiver || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    receiver: event.target.value
                                })
                            }
                            label="请输入接收人"
                            type="text"
                            variant="outlined"
                            size="small"
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon />
                                    </InputAdornment>
                                ),
                            }}
                        /></Item>
                    </Grid>
                    <Grid item >
                        <Item> <TextField

                            sx={{ width: 300 }}
                            autoFocus
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.reqId || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    reqId: event.target.value
                                })
                            }
                            label="请输入请求id"
                            type="text"
                            variant="outlined"
                            size="small"
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon />
                                    </InputAdornment>
                                ),
                            }}
                        /></Item>
                    </Grid>
                    <Grid item >
                        <Item> <TextField
                            select
                            sx={{ width: 300 }}
                            autoFocus
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.read || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    read: event.target.value
                                })
                            }
                            label="请输入是否已读"
                            type="text"
                            variant="outlined"
                            size="small"
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon />
                                    </InputAdornment>
                                ),
                            }}
                        >

                            {readArr.map((option) => (
                                <MenuItem key={option.key} value={option.key}>
                                    {option.title}
                                </MenuItem>
                            ))}</TextField></Item>
                    </Grid>
                    <Grid item >
                        <Item> <TextField
                            select
                            sx={{ width: 300 }}
                            autoFocus
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.sendStatus || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    sendStatus: event.target.value
                                })
                            }
                            label="请输入发送状态"
                            type="text"
                            variant="outlined"
                            size="small"
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon />
                                    </InputAdornment>
                                ),
                            }}
                        >

                            {sendStatusArr.map((option) => (
                                <MenuItem key={option.key} value={option.key}>
                                    {option.title}
                                </MenuItem>
                            ))}</TextField></Item>
                    </Grid>


                    <Grid item >
                        <Item><Stack direction="row" spacing={2} style={{ marginLeft: '10px', height: '30px' }}>
                            <Button size="small" variant="outlined" startIcon={<SearchIcon />} onClick={onSearchClick}>
                                查询
                            </Button>
                            <Button size="small" variant="outlined" startIcon={<AutorenewIcon />} onClick={onResetClick}>
                                重置
                            </Button>
                        </Stack></Item>
                    </Grid>
                </Grid>






            </div>
        </>
    );
};

export default memo(SearchForm);
