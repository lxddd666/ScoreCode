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
const isOnlineArr = [
    {
        title: '在线',
        key: 1
    },
    {
        title: '离线',
        key: 2
    }
];
const SearchForm = (props: any) => {
    const { handleSearchFormData } = props;
    const [value, setValue] = useState<any>(null);
    const [formData, setFormData] = useState<any>({
        id: undefined,
        status: undefined,
        createdAt: undefined,
    })

    // 搜索按钮
    const onSearchClick = (e: any) => {
        console.log(e.target.value, formData);
        let obj = { folderId: value?.value ? value?.value : undefined, ...formData }
        handleSearchFormData(obj);
    };
    // 重置按钮
    const onResetClick = (e: any) => {
        // setValue({});
        // setFormData({
        //     username: undefined,
        //     firstName: undefined,
        //     lastName: undefined,
        //     phone: undefined,
        //     proxyAddress: undefined,
        //     accountStatus: undefined,
        //     isOnline: undefined
        // })
        // handleSearchFormData({});
        let obj = {
            id: undefined,
            status: undefined,
            createdAt: undefined,
        }
        setValue({});
        setFormData(obj)
        handleSearchFormData(obj);
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
                            value={formData.id || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    id: event.target.value
                                })
                            }
                            label="请输入ID"
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
                            value={formData.status || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    status: event.target.value
                                })
                            }
                            label="请输入任务状态"
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

                            {isOnlineArr.map((option) => (
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
