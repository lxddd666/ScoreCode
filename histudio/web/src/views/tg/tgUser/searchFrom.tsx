import { memo, useState } from 'react';
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Autocomplete, { createFilterOptions } from '@mui/material/Autocomplete';
import AutorenewIcon from '@mui/icons-material/Autorenew';
import SearchIcon from '@mui/icons-material/Search';
import styles from './searchForm.module.scss';
const filter = createFilterOptions();
const SearchForm = (props: any) => {
    const { top100Films, handleSearchFormData } = props;
    const [value, setValue] = useState(null);
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

    const onAutocompleteChange = (event: any, newValue: any) => {
        console.log('Autocomplete', event.currentTarget, newValue);

        if (typeof newValue === 'string') {
            // timeout to avoid instant validation of the dialog's form.
            setTimeout(() => {
                toggleOpen(true);
                setDialogValue({
                    title: newValue,
                    year: ''
                });
            });
        } else if (newValue && newValue.inputValue) {
            toggleOpen(true);
            setDialogValue({
                title: newValue.inputValue,
                year: ''
            });
        } else {
            setValue(newValue);
        }
    };

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
        // console.log(e.target.value,value);
        handleSearchFormData(value);
    };
    // 重置按钮
    const onResetClick = (e: any) => {
        setValue(null);
        handleSearchFormData({});
    };
    return (
        <>
            <div className={styles.searchForm}>
                <Autocomplete
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
                    renderInput={(params) => <TextField {...params} label="分组选择" />}
                />
                <Stack direction="row" spacing={2} style={{ marginLeft: '10px' }}>
                    <Button size="small" variant="outlined" startIcon={<SearchIcon />} onClick={onSearchClick}>
                        查询
                    </Button>
                    <Button size="small" variant="outlined" startIcon={<AutorenewIcon />} onClick={onResetClick}>
                        重置
                    </Button>
                </Stack>
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
