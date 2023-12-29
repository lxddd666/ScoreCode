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

import AdapterDateFns from '@date-io/date-fns';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
// import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import { DatePicker } from '@mui/x-date-pickers';
import dayjs from 'dayjs';

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));

const SearchForm = (props: any) => {
    const { handleSearchFormData } = props;
    const [formData, setFormData] = useState<any>({
        name: undefined,
        createdAt: undefined,
    })

    // 搜索按钮
    const onSearchClick = (e: any) => {
        console.log(e.target.value, formData);
        let obj = { ...formData, createdAt: [formData.startTime, formData.end], startTime: undefined, end: undefined }
        handleSearchFormData(obj);
    };
    // 重置按钮
    const onResetClick = (e: any) => {
        let obj = {
            name: undefined,
            createdAt: undefined,
        }
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
                            value={formData.name || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    name: event.target.value
                                })
                            }
                            label="请输入话术组名"
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
                    <Grid>
                        <Item>
                            <LocalizationProvider dateAdapter={AdapterDateFns}>
                                <div style={{ display: 'flex', flexDirection: 'row', justifyContent: 'center' }}>
                                    <DatePicker
                                        label="请输入开始时间"
                                        value={formData.startTime || null}
                                        inputFormat="yyyy/MM/dd"
                                        onChange={(newValue: any) => {
                                            console.log(dayjs(newValue).format('YYYY-MM-DD'));

                                            setFormData({
                                                ...formData,
                                                startTime: dayjs(newValue).format('YYYY-MM-DD')
                                            })
                                        }
                                        }
                                        InputProps={{
                                            startAdornment: (
                                                <InputAdornment position="start">
                                                    <SearchIcon />
                                                </InputAdornment>
                                            ),
                                        }}
                                        renderInput={(params: any) => <TextField
                                            sx={{ width: 300 }}
                                            autoFocus
                                            variant="outlined"
                                            type="text"
                                            margin="dense"
                                            size="small"


                                            {...params} />}
                                    />
                                    <div style={{ display: 'flex', alignItems: 'center', margin: '0 10px' }}> ~ </div>
                                    <DatePicker
                                        label="请输入结束时间"
                                        value={formData.end || null}
                                        inputFormat="yyyy/MM/dd"
                                        onChange={(newValue: any) => {
                                            console.log(dayjs(newValue).format('YYYY-MM-DD'));

                                            setFormData({
                                                ...formData,
                                                end: dayjs(newValue).format('YYYY-MM-DD')
                                            })
                                        }
                                        }
                                        InputProps={{
                                            startAdornment: (
                                                <InputAdornment position="start">
                                                    <SearchIcon />
                                                </InputAdornment>
                                            ),
                                        }}
                                        renderInput={(params: any) => <TextField
                                            sx={{ width: 300 }}
                                            autoFocus
                                            variant="outlined"
                                            type="text"
                                            margin="dense"
                                            size="small"


                                            {...params} />}
                                    />
                                </div>
                            </LocalizationProvider>
                        </Item>
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
