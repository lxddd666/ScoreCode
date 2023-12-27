import { memo, useEffect, useRef, useState, useCallback } from 'react';
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
    Autocomplete,
    Avatar,
    Tooltip,
    IconButton
} from '@mui/material';
// import DeleteIcon from '@mui/icons-material/Delete';
import ChatIcon from '@mui/icons-material/Chat';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import DeleteIcon from '@mui/icons-material/Delete';
import { styled } from '@mui/material/styles';
import Badge from '@mui/material/Badge';
import { useDispatch, useSelector, shallowEqual } from 'store';
import { useHeightComponent } from 'utils/tools';
import { createFilterOptions } from '@mui/material/Autocomplete';
import { openSnackbar } from 'store/slices/snackbar';
import styles from './index.module.scss';
import SearchForm from './searchFrom';
import FileUpload from './upload';
import SubmitDialog from './submitDialog';
import AnimateButton from 'ui-component/extended/AnimateButton';

import { getTgUserListAction } from 'store/slices/tg';

import axios from 'utils/axios';
import { columns, accountStatus, isOnline } from './config';
import FormDialog from './formDialog'

import { timeAgo } from 'utils/tools'
import {
    tgUnUserBindUser,
    tgUserBindUnProxy,
    tgUserBatchLoginOut,
    tgUserAllDelete
} from 'server/tg'
import useConfirm from 'hooks/useConfirm'
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
    const [handleSubmitOpen, setHandleSubmitOpen] = useState(false); // 弹窗控制
    const [handleSubmitOpenConfig, setHandleSubmitOpenConfig] = useState({
        title: ''
    });
    const [formDialogConfig, setFormDialogConfig] = useState<any>({
        title: '',
        edit: false,
        selectCheck: [],
        dialogType: undefined,

        params: undefined,
        renderField: undefined
    })
    const [checkListIsDisable, setCheckListIsDisable] = useState(true)
    const boxRef: any = useRef();
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const confirm = useConfirm(); // 弹窗
    const { tgUserList } = useSelector((state) => state.tg, shallowEqual);

    let { height: boxHeight } = useHeightComponent(boxRef);

    useEffect(() => {
        getTableListActionFN();
        // console.log('tgUserList', tgUserList?.data?.list);
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
    const getTableListActionFN = async () => {
        await dispatch(getTgUserListAction(paramsPayload));
    };
    // 分组选择请求
    const getTgSearchParams = async () => {
        try {
            const res = await axios.get(`/tg/tgFolders/list`, {
                params: {
                    page: 1,
                    pageSize: 9999
                }
            });
            // console.log('tg分组选择请求', res);
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
            setCheckListIsDisable(false)
            return;
        }
        setSelected([]);
        setCheckListIsDisable(true)
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
        // console.log(newSelected, selected);

        if (newSelected.length > 0) {
            setCheckListIsDisable(false)
        } else {
            setCheckListIsDisable(true)
        }
    };
    // id筛选
    const isSelected = (id: any) => selected.indexOf(id) !== -1;

    const renderTable = (value: any, key: any, item: any) => {
        // console.log(value, key, item);

        let temp: any = '';
        if (key === 'firstName') {
            temp = <div className={styles.tablesColumns}>
                <div className={styles.avatars}>
                    <StyledBadge
                        overlap="circular"
                        anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
                        variant="dot"
                        badgeColor={item.isOnline === 1 ? '#44b700' : 'red'}
                    >
                        <Avatar alt="Remy Sharp" src={item.photo}>
                            {item.lastName?.charAt(0)?.toUpperCase()}
                        </Avatar>
                    </StyledBadge>
                </div>
                <div className={styles.info}>
                    <div className={styles.titles}>
                        <p>{item.firstName}</p>
                        <p style={{ marginLeft: '5px' }}>{item.lastName}</p>
                    </div>
                    <div style={{ fontSize: '12px' }}>phone:{item.phone}</div>
                </div>
            </div>
        }
        else if (key === 'accountStatus') {
            temp = <Chip label={accountStatus(value)} color="primary" />;
        } else if (key === 'isOnline') {
            temp = <Chip label={isOnline(value)} sx={{ bgcolor: `${item.isOnline === 1 ? '#44b700' : 'red'}`, color: 'white' }} />;
        } else if (key === 'lastLoginTime') {
            temp = timeAgo(value)
        } else {
            temp = value;
        }
        return temp
    };

    // 分页事件
    const pageRef = useRef(1);
    const onPaginationChange = (event: object, page: number) => {
        pageRef.current = page;

        setParamsPayload({ ...paramsPayload, page: pageRef.current });
    };
    // 分页数量
    const PaginationCount = (count: number) => {
        return typeof count === 'number' ? Math.ceil(count / 10) : 1;
    };

    // 子传父 searchForm
    const handleSearchFormData = (obj: any) => {
        setParamsPayload({ ...paramsPayload, ...obj, page: 1 });
    };
    const handleSetImportOpenDialog = (type: String, value: any) => {
        // setImportOpenDialog(value);
        onBtnCloseList(type, value);
    };

    // 聊天室跳转
    const chatRoomToNavica = (rows: any) => {
        // console.log(rows);
        navigate(`/tg/chat/index/${rows.id}`);
        // navigate(`/tg/chat/index/1`);
    };
    // 弹窗开启
    const handleSubmitOpenCallback = useCallback(() => {
        setHandleSubmitOpen(true);
        setHandleSubmitOpenConfig({ ...handleSubmitOpenConfig, title: '手机验证码登录' });
    }, [handleSubmitOpen]);
    const onBtnOpenList = (active: String, value: any = undefined) => {
        let columns = []
        switch (active) {
            case 'import':
                setImportOpenDialog(true);
                break;
            case 'iphone':
                handleSubmitOpenCallback();
                break;
            case 'bindUser':
                // 绑定员工开启
                columns = [
                    {
                        id: 'username',
                        numeric: false,
                        label: '用户名',
                        align: 'center'
                    },
                    {
                        id: 'roleName',
                        numeric: false,
                        label: '绑定角色',
                        align: 'center'
                    },
                    {
                        id: 'orgName',
                        numeric: false,
                        label: '所属公司',
                        align: 'center'
                    },
                    {
                        id: 'status',
                        numeric: false,
                        label: '状态',
                        align: 'center'
                    },
                    {
                        id: 'lastActiveAt',
                        numeric: false,
                        label: '最近活跃',
                        align: 'center'
                    }
                ];
                setFormDialogConfig({ ...formDialogConfig, edit: true, title: '绑定用户', selectCheck: selected, columns, type: 'bindUser' });
                break;
            case 'bindProxy':
                // 绑定员工开启
                columns = [
                    {
                        id: 'address',
                        numeric: false,
                        label: '代理地址',
                        align: 'center'
                    },
                    {
                        id: 'type',
                        numeric: false,
                        label: '代理类型',
                        align: 'center'
                    },
                    {
                        id: 'maxConnections',
                        numeric: false,
                        label: '最大连接数',
                        align: 'center'
                    },
                    {
                        id: 'connectedCount',
                        numeric: false,
                        label: '已连接数',
                        align: 'center'
                    },
                    {
                        id: 'assignedCount',
                        numeric: false,
                        label: '已分配数',
                        align: 'center'
                    },
                    {
                        id: 'longTermCount',
                        numeric: false,
                        label: '长期未登录',
                        align: 'center'
                    },
                    {
                        id: 'region',
                        numeric: false,
                        label: '地区',
                        align: 'center'
                    },
                    {
                        id: 'delay',
                        numeric: false,
                        label: '延迟',
                        align: 'center'
                    },
                    {
                        id: 'status',
                        numeric: false,
                        label: '状态',
                        align: 'center'
                    }
                ];
                setFormDialogConfig({ ...formDialogConfig, edit: true, title: '绑定代理', selectCheck: selected, columns, type: 'bindProxy' });
                break;
            case 'Edit':
                let renderField = {
                    username: '用户名',
                    firstName: '姓氏',
                    lastName: '名字',
                    phone: '手机号码',
                    bio: '个性签名',
                    comment: '备注'
                }
                setFormDialogConfig({
                    ...formDialogConfig,
                    edit: true,
                    title: '绑定用户',
                    dialogType: 'editForm',
                    type: 'Edit',
                    params: value,
                    renderField
                });
                break
            default:
                break;
        }
    }
    // 弹窗关闭
    const handleSubmitCloseCallback = useCallback((value: any) => {
        setHandleSubmitOpen(value);
    }, []);
    const onBtnCloseList = (type: String, value: any) => {
        // console.log(type, value);

        switch (type) {
            case 'import':
                setImportOpenDialog(value);
                break;
            case 'iphone':
                handleSubmitCloseCallback(value);
                break;
            case 'bindUser':
                // 绑定员工关闭
                setFormDialogConfig({ ...formDialogConfig, edit: value, title: '', selectCheck: [], type: '' });
                getTableListActionFN()
                break;
            case 'bindProxy':
                // 绑定代理关闭
                setFormDialogConfig({ ...formDialogConfig, edit: value, title: '', selectCheck: [], type: '' });
                getTableListActionFN()
                break;
            case 'Edit':
                setFormDialogConfig({ ...formDialogConfig, edit: value, title: '', dialogType: '', type: '', prams: undefined });
                getTableListActionFN()
                break
            default:
                break;
        }
    };

    // 解绑员工操作
    const onUnBindUserClick = () => {
        // console.log('解绑员工操作', selected);
        confirm('警告', `是否解绑选中的 ${selected.length} 个员工,确定之后不可取消,请谨慎操作。`)
            .then(async (result) => {
                if (result) {
                    try {
                        const res = await tgUnUserBindUser({ ids: selected })
                        console.log('解绑成功', res);
                    } catch (error) {
                        console.log('解绑失败', error);
                    }
                } else {
                    console.log('Cancelled!');
                }
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }
    // 解绑代理操作
    const onUnBindProxyClick = () => {
        // console.log('解绑员工操作', selected);
        confirm('警告', `是否解绑选中的 ${selected.length} 个代理,确定之后不可取消,请谨慎操作。`)
            .then(async (result) => {
                if (result) {
                    try {
                        const res = await tgUserBindUnProxy({ ids: selected })
                        console.log('解绑成功', res);
                        getTableListActionFN()
                    } catch (error) {
                        console.log('解绑失败', error);
                    }
                } else {
                    console.log('Cancelled!');
                }
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }
    // 批量上线 下线 操作
    const onBatchLoginOutClick = (type: String) => {
        // console.log('解绑员工操作', selected);
        confirm('警告', `是否${type === 'batchLogin' ? '上线' : '下线'}选中的 ${selected.length} 个数据,确定之后不可取消,请谨慎操作。`)
            .then(async (result) => {
                if (result) {
                    try {
                        const res = await tgUserBatchLoginOut({ ids: selected }, type)
                        console.log(' 批量上线 下线 操作', res);

                        getTableListActionFN()
                    } catch (error) {
                        console.log('批量上下线失败', error);
                    }
                } else {
                    console.log('Cancelled!');
                }
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }
    // 批量删除 操作
    const onUserAllDeleteClick = (id: any = undefined) => {
        confirm('警告', `是否批量删除选中的 ${selected.length} 个数据,确定之后不可取消,请谨慎操作。`)
            .then(async (result) => {
                if (result) {
                    try {
                        const res = await tgUserAllDelete({ id: selected })
                        console.log('批量删除', res);
                        getTableListActionFN()
                        setSelected([]);
                        setCheckListIsDisable(true)
                    } catch (error) {
                        console.log('批量删除失败', error);
                    }
                } else {
                    console.log('Cancelled!');
                }
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }
    // 单个删除 操作
    const onUserSingleDeleteClick = (id: any) => {
        confirm('警告', `是否删除选中的 1 个数据,确定之后不可取消,请谨慎操作。`)
            .then(async (result) => {
                if (result) {
                    try {
                        const res = await tgUserAllDelete({ id: id })
                        console.log('批量删除', res);
                        getTableListActionFN()
                        setSelected([]);
                        setCheckListIsDisable(true)
                    } catch (error) {
                        console.log('批量删除失败', error);
                    }
                } else {
                    console.log('Cancelled!');
                }
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }


    return (
        <MainCard title={<FormattedMessageTitle />} content={true}>
            <div className={styles.box} ref={boxRef}>
                <div className={styles.searchTop}>
                    <SearchForm top100Films={searchForm} handleSearchFormData={handleSearchFormData} />
                </div>
                <div className={styles.btnList}>
                    <Stack direction="row" spacing={2}>
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" disabled={checkListIsDisable} color="error" onClick={onUserAllDeleteClick}>
                                批量删除
                            </Button>
                        </AnimateButton>
                        {/* <AnimateButton type="slide">
                            <Button size="small" variant="contained" >
                                导出
                            </Button>
                        </AnimateButton> */}
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" disabled={checkListIsDisable} onClick={(e) => onBtnOpenList('bindUser')} color="secondary">
                                绑定员工
                            </Button>
                        </AnimateButton>
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" disabled={checkListIsDisable} onClick={onUnBindUserClick} color="secondary">
                                解绑员工
                            </Button>
                        </AnimateButton>
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" disabled={checkListIsDisable} onClick={(e) => onBtnOpenList('bindProxy')} color="success">
                                绑定代理
                            </Button>
                        </AnimateButton>
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" disabled={checkListIsDisable} onClick={onUnBindProxyClick} color="success">
                                解绑代理
                            </Button>
                        </AnimateButton>
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" disabled={checkListIsDisable} onClick={(e) => onBatchLoginOutClick('batchLogin')} color="warning">
                                批量上线
                            </Button>
                        </AnimateButton>
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" disabled={checkListIsDisable} onClick={(e) => onBatchLoginOutClick('batchLogout')} color="warning">
                                批量下线
                            </Button>
                        </AnimateButton>
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" onClick={(e) => onBtnOpenList('import')}>
                                导入
                            </Button>
                        </AnimateButton>
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" onClick={(e) => onBtnOpenList('iphone')}>
                                手机验证码登录
                            </Button>
                        </AnimateButton>
                    </Stack>
                </div>
                <TableContainer
                    component={Paper}
                    style={{ maxHeight: `calc(${boxHeight - 270}px)`, borderTop: '1px solid #eaeaea', borderBottom: '1px solid #eaeaea' }}
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
                                                {renderTable(row[item.key], item.key, row)}

                                                {item.key === 'active' ? (
                                                    <div style={item.key === 'active' ? { width: '220px' } : {}}>
                                                        <IconButton aria-label="delete" onClick={(e) => chatRoomToNavica(row)}>
                                                            <Tooltip title='聊天室' placement="top">
                                                                <ChatIcon style={{ color: 'rgb(3, 106, 129)', fontSize: '18px' }} />
                                                            </Tooltip>
                                                        </IconButton>
                                                        {/* <Button size="small" variant="contained">
                                                            聊天室
                                                        </Button> */}
                                                        <IconButton style={{ marginLeft: '5px' }} onClick={(e) => onBtnOpenList('Edit', row)}>
                                                            <Tooltip title='编辑' placement="top">
                                                                <ModeEditIcon style={{ color: 'rgb(3, 106, 129)', fontSize: '18px' }} />
                                                            </Tooltip>
                                                        </IconButton>
                                                        <IconButton style={{ marginLeft: '5px' }} onClick={(e) => onUserSingleDeleteClick(row.id)} >
                                                            <Tooltip title='删除' placement="top">
                                                                <DeleteIcon style={{ color: 'rgb(159, 86, 108)', fontSize: '18px' }} />
                                                            </Tooltip>
                                                        </IconButton>
                                                    </div>
                                                ) : (
                                                    ''
                                                )}
                                            </TableCell>
                                        );
                                    })}

                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
                {pagetionTotle && pagetionTotle !== 0 && (
                    <>
                        <div className={styles.paginations}>
                            <div>共 {pagetionTotle} 条</div>
                            <Pagination count={PaginationCount(pagetionTotle)} color="primary" onChange={onPaginationChange} />
                        </div>
                    </>
                )}
            </div>

            <ImportOpenDialog importOpenDialog={importOpenDialog} data={searchForm} setImportOpenDialog={handleSetImportOpenDialog} getTgUserListActionFN={getTableListActionFN} getTgSearchParams={getTgSearchParams} />

            <SubmitDialog open={handleSubmitOpen} config={handleSubmitOpenConfig} setOpenChangeDialog={handleSetImportOpenDialog} />

            <FormDialog open={formDialogConfig.edit} config={formDialogConfig} onChangeDialogStatus={onBtnCloseList} />
        </MainCard>
    );
};
// 标题 tg
const FormattedMessageTitle = () => {
    return (
        <div className={styles.FormattedMessageTitle}>
            <FormattedMessage id="setting.cron.teleg-tg" />
            {/* <div>
                <Button variant="outlined">登录</Button>
            </div> */}
        </div>
    );
};
// 导入弹窗
const filter = createFilterOptions();
const ImportOpenDialog = (props: any) => {
    const { importOpenDialog, setImportOpenDialog, data, getTgUserListActionFN, getTgSearchParams } = props;

    const [value, setValue] = useState<any>('');
    const [open, toggleOpen] = useState(false);
    const [dialogValue, setDialogValue] = useState({
        folderName: ''
    });
    const [selectedFile, setSelectedFile] = useState(null);
    const dispatch = useDispatch();

    // dialog 弹出关闭
    const handleImportClose = () => {
        setImportOpenDialog('import', false);
        setValue('')
    };
    // dialog 提交
    const handleImportSubmit = (event: any) => {
        event.preventDefault();
        // setImportOpenDialog(false);

        if (!selectedFile) {
            alert('只能上传zip格式的文件，请重新上传');
            return;
        }
        console.log('value提交', value);

        const formData = new FormData();
        formData.append('file', selectedFile);
        formData.append('folderId', value?.value);
        // axios('http://10.8.12.88:8001/tg/tgUser/importSession', {
        axios('/tg/tgUser/importSession', {
            method: 'POST',
            transformRequest: [function (data, headers: any) {
                // 去除post请求默认的Content-Type
                // console.log(data, headers);
                // delete headers.post['Content-Type']
                return data
            }],
            data: formData,
        }).then(res => {
            // 处理响应
            console.log('res上传成功', res);
            setImportOpenDialog('import', false);
            setValue('')
            getTgUserListActionFN()
        }).catch(err => {
            console.log('res上传失败', err);
            dispatch(openSnackbar({
                open: true,
                message: err.message || '上传失败',
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
        })
    };

    const selectedFileChange = (file: any) => {
        setSelectedFile(file)
    }

    const handleClose = () => {
        setDialogValue({
            folderName: ''
        });

        toggleOpen(false);
    };

    // 导入分组选择
    const onAutocompleteChange = (event: any, newValue: any) => {
        console.log('Autocomplete', event.currentTarget, newValue);

        if (typeof newValue === 'string') {
            // timeout to avoid instant validation of the dialog's form.
            setTimeout(() => {
                toggleOpen(true);
                setDialogValue({
                    folderName: newValue
                });
            });
        } else if (newValue && newValue.inputValue) {
            toggleOpen(true);
            setDialogValue({
                folderName: newValue.inputValue
            });
        } else {
            setValue(newValue);
        }
    };

    // 添加分组名称
    const handleSubmit = (event: any) => {
        event.preventDefault();
        let obj: any = {
            folderName: dialogValue.folderName
        };
        console.log('添加分组名称', dialogValue.folderName);

        axios.post('/tg/tgFolders/edit', {
            ...obj
        }).then(({ data }) => {
            console.log('添加分组', data);

            // setValue('');
            setValue('')
            getTgSearchParams()
            handleClose();
        }).catch(err => {
            console.log('添加分组失败');

        })

    };
    // 下载模板
    const handleDownload = () => {
        let a = document.createElement("a");
        a.href = "./static/session.zip";
        a.download = "session.zip";
        a.style.display = "none";
        document.body.appendChild(a);
        a.click();
        a.remove();

    }
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
                <DialogTitle>导入 <Button onClick={handleDownload}>下载模板</Button></DialogTitle>
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
                                renderOption={(props, option) => <li {...props} key={option.value}>{option.title}</li>}
                                sx={{ width: 300 }}
                                freeSolo
                                renderInput={(params) => <TextField {...params} label="分组选择" />}
                                style={{ width: '100%', margin: '10px 0' }}
                            />
                        </div>
                        <div className={styles.formBox}>
                            <p className={styles.formTitle}>上传文件：</p>
                            <FileUpload selectedFileChange={selectedFileChange} />
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
                            value={dialogValue.folderName}
                            onChange={(event) =>
                                setDialogValue({
                                    ...dialogValue,
                                    folderName: event.target.value
                                })
                            }
                            label="分组名称"
                            type="text"
                            variant="outlined"
                            style={{ width: '100%' }}
                        />
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose}>取消</Button>
                        <Button type="submit">添加</Button>
                    </DialogActions>
                </form>
            </Dialog>
        </>
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
export default memo(TgUser);
