import { memo, useState } from 'react';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
// import Autocomplete, { createFilterOptions } from '@mui/material/Autocomplete';
import AutorenewIcon from '@mui/icons-material/Autorenew';
import SearchIcon from '@mui/icons-material/Search';
import styles from './searchForm.module.scss';
import { styled } from '@mui/material/styles';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import InputAdornment from '@mui/material/InputAdornment';
import MenuItem from '@mui/material/MenuItem';
// import { DatePicker } from '@mui/x-date-pickers';
import {accountStatusArr, isOnlineArr} from './config';
// import { LocalizationProvider } from '@mui/x-date-pickers';
// import AdapterDateFns from '@mui/lab/AdapterDateFns';
// import AdapterDateFns from '@mui/x-date-pickers/AdapterDateFns';

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));
// const filter = createFilterOptions();
const SearchForm = (props: any) => {
    const { top100Films, handleSearchFormData } = props;
    const [value, setValue] = useState<any>(null);
    const [formData, setFormData] = useState<any>({
        folderId: undefined,
        username: undefined,
        firstName: undefined,
        lastName: undefined,
        phone: undefined,
        proxyAddress: undefined,
        accountStatus: undefined,
        isOnline: undefined
    })
    const [open, toggleOpen] = useState(false);
    const [dialogValue, setDialogValue] = useState({
        title: '',
        year: ''
    });

    const handleClose = () => {
        setDialogValue({
            title: '',
            year: ''
        });

        toggleOpen(false);
    };

    // const onAutocompleteChange = (event: any, newValue: any) => {
    //     console.log('Autocomplete', event.currentTarget, newValue);

    //     if (typeof newValue === 'string') {
    //         // timeout to avoid instant validation of the dialog's form.
    //         setTimeout(() => {
    //             toggleOpen(true);
    //             setDialogValue({
    //                 title: newValue,
    //                 year: ''
    //             });
    //         });
    //     } else if (newValue && newValue.inputValue) {
    //         toggleOpen(true);
    //         setDialogValue({
    //             title: newValue.inputValue,
    //             year: ''
    //         });
    //     } else {
    //         setValue(newValue);
    //     }
    // };

    const handleSubmit = (event: any) => {
        event.preventDefault();
        let obj: any = {
            title: dialogValue.title,
            year: parseInt(dialogValue.year, 10)
        };
        setValue(obj);

        handleClose();
    };

    // 搜索按钮
    const onSearchClick = (e: any) => {
        console.log(e.target.value, formData);
        let obj = { folderId: value?.value ? value?.value : undefined, ...formData };
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
        handleSearchFormData({});
        let obj = {
            folderId: undefined,
            username: undefined,
            firstName: undefined,
            lastName: undefined,
            phone: undefined,
            proxyAddress: undefined,
            accountStatus: undefined,
            isOnline: undefined
        };
        setValue({});
        setFormData(obj);
        handleSearchFormData(obj);
    };
    return (
        <>
            <div className={styles.searchForm}>
                <Grid container spacing={0.3} alignItems="center">
                    <Grid item>
                        <Item> <TextField
                            className={styles.ipt}

                            autoFocus
                            sx={{ width: 300 }}
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.username || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    username: event.target.value
                                })
                            }
                            label="请输入用户名"
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
                        <Item><TextField

                            sx={{ width: 300 }}
                            autoFocus
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.firstName || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    firstName: event.target.value
                                })
                            }
                            label="请输入名字"
                            type="text"
                            variant="outlined"
                            size="small" InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon />
                                    </InputAdornment>
                                ),
                            }}
                        /></Item>
                    </Grid>
                    <Grid item >
                        <Item>
                            {/* <Autocomplete
                            size="small"
                            value={value}
                            onChange={onAutocompleteChange}
                            filterOptions={(options, params) => {
                                const filtered = filter(options, params);

                                if (params.inputValue !== '') {
                                    filtered.push({
                                        inputValue: params.inputValue,
                                        title: `Add "${params.inputValue}"`
                                    });
                                }

                                return filtered;
                            }}
                            // id="free-solo-dialog-demo"
                            id="controllable-states-demo"
                            options={top100Films}
                            getOptionLabel={(option: any) => {
                                // e.g value selected with enter, right from the input
                                if (typeof option === 'string') {
                                    return option;
                                }
                                if (option?.inputValue) {
                                    return option.inputValue;
                                }
                                return option.title;
                            }}
                            selectOnFocus
                            clearOnBlur
                            handleHomeEndKeys
                            renderOption={(props, option) => <li {...props}>{option.title}</li>}
                            sx={{ width: 300 }}
                            freeSolo

                            renderInput={(params) => <TextField {...params} label="分组选择" InputProps={{
                                ...params.InputProps,
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon />
                                    </InputAdornment>
                                ),
                            }} />}
                        /> */}
                            <TextField
                                select
                                sx={{ width: 300 }}
                                autoFocus
                                margin="dense"
                                id="standard-required"
                                inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                                value={formData.folderId || ''}
                                onChange={(event) =>
                                    setFormData({
                                        ...formData,
                                        folderId: event.target.value
                                    })
                                }
                                label="请输入分组选择"
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

                                {top100Films.map((option: any) => (
                                    <MenuItem key={option.value} value={option.value}>
                                        {option.title}
                                    </MenuItem>
                                ))}</TextField>
                        </Item>
                    </Grid>
                    <Grid item >
                        <Item> <TextField

                            sx={{ width: 300 }}
                            autoFocus
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.lastName || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    lastName: event.target.value
                                })
                            }
                            label="请输入姓氏"
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
                            value={formData.phone || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    phone: event.target.value
                                })
                            }
                            label="请输入手机号"
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
                            value={formData.accountStatus || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    accountStatus: event.target.value
                                })
                            }
                            label="请输入账号状态"
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

                            {accountStatusArr.map((option) => (
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
                            value={formData.isOnline || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    isOnline: event.target.value
                                })
                            }
                            label="请输入在线状态"
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
                        <Item> <TextField

                            sx={{ width: 300 }}
                            autoFocus
                            margin="dense"
                            id="standard-required"
                            inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                            value={formData.proxyAddress || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    proxyAddress: event.target.value
                                })
                            }
                            label="请输入代理地址"
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
                    {/* <Grid item >
                        <Item>
                            <LocalizationProvider dateAdapter={AdapterDateFns}><DatePicker
                                label="创建时间"
                                value={formData.date}
                                minDate={new Date('2017-01-01')}
                                onChange={(event: any) =>
                                    setFormData({
                                        ...formData,
                                        date: event.target.value
                                    })}
                                renderInput={(params: any) => <TextField {...params} />}
                            /> </LocalizationProvider></Item>
                    </Grid> */}
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

            <Dialog open={open} onClose={handleClose}>
                <form onSubmit={handleSubmit}>
                    <DialogTitle>添加分组名称</DialogTitle>
                    <DialogContent>
                        <TextField
                            autoFocus
                            margin="dense"
                            id="name"
                            value={dialogValue.title}
                            onChange={(event) =>
                                setDialogValue({
                                    ...dialogValue,
                                    title: event.target.value
                                })
                            }
                            label="title"
                            type="text"
                            variant="standard"
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose}>Cancel</Button>
                        <Button type="submit">Add</Button>
                    </DialogActions>
                </form>
            </Dialog>
        </>
    );
};

export default memo(SearchForm);
