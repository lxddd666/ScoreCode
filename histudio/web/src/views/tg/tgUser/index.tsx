import { memo, useEffect, useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import MainCard from 'ui-component/cards/MainCard';
import { FormattedMessage } from 'react-intl';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import {
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    TextField,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    Checkbox,
    Chip,
    Pagination,
    Autocomplete
} from '@mui/material';
import { useDispatch, useSelector } from 'store';
import { useHeightComponent } from 'utils/tools';
import { createFilterOptions } from '@mui/material/Autocomplete';
import styles from './index.module.scss';
import SearchForm from './searchFrom';
import  FileUpload from './upload'

import { getTgUserListAction } from 'store/slices/tg';
import axios from 'utils/axios';
import { columns, accountStatus, isOnline } from './conig';

const TgUser = () => {
    const [selected, setSelected] = useState<any>([]); // 多选
    const [rows, setrows] = useState([]); // table rows 数据
    const [paramsPayload, setParamsPayload] = useState({
        page: 1,
        pageSize: 10,
        folderId: undefined
    }); // 分页
    const [searchForm, setSearchForm] = useState([]); // search Form
    const [pagetionTotle, setPagetionTotle] = useState(0); // total
    const [importOpenDialog, setImportOpenDialog] = useState(false);
    const boxRef: any = useRef();
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const { tgUserList } = useSelector((state) => state.tg);

    let { height: boxHeight } = useHeightComponent(boxRef);

    useEffect(() => {
        getTgUserListActionFN();
        console.log('tgUserList', tgUserList?.data?.list);
    }, [dispatch, paramsPayload]);
    // 数据赋值
    useEffect(() => {
        // getTgSearchParams();
        setrows(tgUserList?.data?.list || []);
        setPagetionTotle(tgUserList?.data?.totalCount);
    }, [tgUserList]);
    // 网络请求
    useEffect(() => {
        getTgSearchParams();
    }, []);

    // tgUser 表格数据
    const getTgUserListActionFN = async () => {
        await dispatch(getTgUserListAction(paramsPayload));
    };
    // 分组选择请求
    const getTgSearchParams = async () => {
        try {
            const res = await axios.get(`/tg/tgFolders/list`);
            console.log('tg分组选择请求', res);
            let arr: any = [];
            res?.data?.data?.list.map((item: any) => {
                arr.push({
                    // title:item.folderName,
                    title: item.folderName,
                    value: item.id
                });
            });
            setSearchForm(arr);
        } catch (error) {
            console.log('分组数据请求失败');
        }
    };
    // table多选all操作
    const handleSelectAllClick = (event: any) => {
        if (event.target.checked) {
            const newSelecteds = rows.map((n: any) => n.id);
            setSelected(newSelecteds);
            return;
        }
        setSelected([]);
    };
    // table多选点击操作
    const handleClick = (event: any, id: any) => {
        const selectedIndex = selected.indexOf(id);
        let newSelected: any = [];

        if (selectedIndex === -1) {
            newSelected = newSelected.concat(selected, id);
        } else if (selectedIndex === 0) {
            newSelected = newSelected.concat(selected.slice(1));
        } else if (selectedIndex === selected.length - 1) {
            // } else if (selectedIndex === selected.length) {
            newSelected = newSelected.concat(selected.slice(0, -1));
        } else if (selectedIndex > 0) {
            newSelected = newSelected.concat(selected.slice(0, selectedIndex), selected.slice(selectedIndex + 1));
        }

        setSelected(newSelected);
    };
    // id筛选
    const isSelected = (id: any) => selected.indexOf(id) !== -1;

    const renderTable = (value: any, key: any) => {
        let temp: any = '';
        if (key === 'accountStatus') {
            temp = <Chip label={accountStatus(value)} color="primary" />;
        } else if (key === 'isOnline') {
            temp = <Chip label={isOnline(value)} color="primary" />;
        } else {
            temp = value;
        }
        return temp;
    };

    // 分页事件
    const pageRef = useRef(1);
    const onPaginationChange = (event: object, page: number) => {
        pageRef.current = page;

        setParamsPayload({ ...paramsPayload, page: pageRef.current });
    };

    // 子传父 searchForm
    const handleSearchFormData = (obj: any) => {
        setParamsPayload({ ...paramsPayload, folderId: obj?.value, page: 1 });
    };
    const handleSetImportOpenDialog = (value: any) => {
        setImportOpenDialog(value);
    };

    // 聊天室跳转
    const chatRoomToNavica = (rows: any) => {
        // console.log(rows);
        navigate(`/tg/chat/index?id=${rows.id}`);
    };
    const onBtnList = (active: String) => {
        switch (active) {
            case 'import':
                setImportOpenDialog(true);
                break;
            default:
                break;
        }
    };

    return (
        <MainCard title={<FormattedMessageTitle />} content={true}>
            <div className={styles.box} ref={boxRef}>
                <div className={styles.searchTop}>
                    <SearchForm top100Films={searchForm} handleSearchFormData={handleSearchFormData} />
                </div>
                <div className={styles.btnList}>
                    <Stack direction="row" spacing={2}>
                        <Button size="small" variant="contained" onClick={(e) => onBtnList('import')}>
                            导入
                        </Button>
                        <Button size="small" variant="contained">
                            导出
                        </Button>
                    </Stack>
                </div>
                <TableContainer
                    component={Paper}
                    style={{ maxHeight: `calc(${boxHeight - 170}px)`, borderTop: '1px solid #eaeaea', borderBottom: '1px solid #eaeaea' }}
                >
                    <Table aria-label="simple table" sx={{ border: 1, borderColor: 'divider' }} stickyHeader={true}>
                        <TableHead>
                            <TableRow>
                                <TableCell padding="checkbox">
                                    <Checkbox
                                        indeterminate={selected.length > 0 && selected.length < rows.length}
                                        checked={rows.length > 0 && selected.length === rows.length}
                                        onChange={handleSelectAllClick}
                                        inputProps={{ 'aria-label': 'select all desserts' }}
                                    />
                                </TableCell>
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.title}>
                                            {item.title}
                                        </TableCell>
                                    );
                                })}
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {rows.map((row: any) => (
                                <TableRow
                                    key={row.id}
                                    hover
                                    onClick={(event) => handleClick(event, row.id)}
                                    role="checkbox"
                                    aria-checked={isSelected(row.id)}
                                    tabIndex={-1}
                                    selected={isSelected(row.id)}
                                >
                                    <TableCell padding="checkbox">
                                        <Checkbox
                                            checked={isSelected(row.id)}
                                            inputProps={{ 'aria-labelledby': `enhanced-table-checkbox-${row.id}` }}
                                        />
                                    </TableCell>
                                    {columns.map((item) => {
                                        return (
                                            <TableCell align="center" key={item.key}>
                                                {renderTable(row[item.key], item.key)}

                                                {/* {item.key === 'accountStatus' ? <Chip label={accountStatus(row[item.key])} color="primary" />:''}
                                                {item.key === 'isOnline' ? <Chip label={isOnline(row[item.key])} color="primary" /> : ''} */}
                                                {item.key === 'active' ? (
                                                    <Button size="small" variant="contained" onClick={(e) => chatRoomToNavica(row)}>
                                                        聊天室
                                                    </Button>
                                                ) : (
                                                    ''
                                                )}
                                            </TableCell>
                                        );
                                    })}
                                    {/* <TableCell align="center">{row.memberUsername}</TableCell>
                                    <TableCell align="center">{row.username}</TableCell>
                                    <TableCell align="center">{row.firstName}</TableCell>
                                    <TableCell align="center">{row.phone}</TableCell>
                                    <TableCell align="center">{row.folderId}</TableCell>
                                    <TableCell align="center">{row.lastName}</TableCell>
                                    <TableCell align="center">{row.accountStatus}</TableCell>
                                    <TableCell align="center">{row.isOnline}</TableCell>
                                    <TableCell align="center">{row.proxyAddress}</TableCell>
                                    <TableCell align="center">{row.lastLoginTime}</TableCell>
                                    <TableCell align="center">{row.comment}</TableCell>
                                    <TableCell align="center">{row.createdAt}</TableCell>
                                    <TableCell align="center">{row.updatedAt}</TableCell> */}
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
                {pagetionTotle !== 0 ? (
                    <>
                        <div className={styles.paginations}>
                            <div>共 {pagetionTotle} 条</div>
                            <Pagination count={10} color="primary" onChange={onPaginationChange} />
                        </div>
                    </>
                ) : (
                    ''
                )}
            </div>

            <ImportOpenDialog importOpenDialog={importOpenDialog} data={searchForm} setImportOpenDialog={handleSetImportOpenDialog} />
        </MainCard>
    );
};
// 标题 tg
const FormattedMessageTitle = () => {
    return (
        <div className={styles.FormattedMessageTitle}>
            <FormattedMessage id="setting.cron.teleg-tg" />
            <div>
                <Button variant="outlined">登录</Button>
            </div>
        </div>
    );
};
// 导入弹窗
const filter = createFilterOptions();
const ImportOpenDialog = (props: any) => {
    const { importOpenDialog, setImportOpenDialog, data } = props;

    const [value, setValue] = useState(null);
    const [open, toggleOpen] = useState(false);
    const [dialogValue, setDialogValue] = useState({
        title: '',
        year: ''
    });

    // dialog 弹出关闭
    const handleImportClose = () => {
        setImportOpenDialog(false);
    };
    // dialog 提交
    const handleImportSubmit = (event: any) => {
        event.preventDefault();
        setImportOpenDialog(false);
    };

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
    return (
        <>
            <Dialog
                open={importOpenDialog}
                onClose={(event: any, reason: any) => {
                    if (reason !== 'backdropClick' && reason !== 'escapeKeyDown') {
                        handleImportClose();
                    }
                }}
                disableEscapeKeyDown={true}
            >
                <DialogTitle>导入</DialogTitle>
                <DialogContent>
                    <div className={styles.dialog}>
                       <div className={styles.formBox}>
                        <p className={styles.formTitle}>分组选择：</p>
                       <Autocomplete
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
                            options={data}
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
                            style={{ width: '100%', margin: '10px 0' }}
                        />
                       </div>
                       <div className={styles.formBox}>
                        <p className={styles.formTitle}>分组选择：</p>
                        <FileUpload />
                       </div>
                    </div>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleImportClose}>取消</Button>
                    <Button onClick={handleImportSubmit}>提交</Button>
                </DialogActions>
            </Dialog>

            <Dialog open={open} onClose={handleClose}>
                <form onSubmit={handleSubmit}>
                    <DialogTitle>添加分组名称</DialogTitle>
                    <DialogContent>
                        <TextField
                            autoFocus
                            margin="dense"
                            id="outlined-basic"
                            value={dialogValue.title}
                            onChange={(event) =>
                                setDialogValue({
                                    ...dialogValue,
                                    title: event.target.value
                                })
                            }
                            label="分组名称"
                            type="text"
                            variant="outlined"
                            style={{width:'100%'}}
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
export default memo(TgUser);
