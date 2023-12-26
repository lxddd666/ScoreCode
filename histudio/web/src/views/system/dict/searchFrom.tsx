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
        id: undefined,
        type: undefined,
        label: undefined,
        value: undefined,
        valueType: undefined,
        status: undefined,
        active: undefined
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
            id: undefined,
            type: undefined,
            label: undefined,
            value: undefined,
            valueType: undefined,
            status: undefined,
            active: undefined
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
                            value={formData.type || ''}
                            onChange={(event) =>
                                setFormData({
                                    ...formData,
                                    type: event.target.value
                                })
                            }
                            label="请输入标签名称"
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
