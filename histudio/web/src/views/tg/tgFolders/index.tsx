import {memo, useState, useRef, useEffect, useCallback, useMemo} from "react"
import { FormattedMessage } from 'react-intl';
// import { useNavigate } from 'react-router-dom';
import MainCard from 'ui-component/cards/MainCard';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import {
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    Checkbox,
    // Chip,
    Pagination,
    Tooltip,
    IconButton
} from '@mui/material';
import DetailsIcon from '@mui/icons-material/Details';
import ModeEditIcon from '@mui/icons-material/ModeEdit';
import DeleteIcon from '@mui/icons-material/Delete';
import { useDispatch, useSelector } from 'store';
import { useHeightComponent} from 'utils/tools';
// import { createFilterOptions } from '@mui/material/Autocomplete';
// import { openSnackbar } from 'store/slices/snackbar';
import styles from './index.module.scss';
import SearchForm from './searchFrom';

import { getTgFoldersListAction } from 'store/slices/tg';
import axios from 'utils/axios';
import { columns } from './config';
import {tgFoldersDelete, tgFoldersView} from "../../../server/tg";
import {openSnackbar} from "../../../store/slices/snackbar";
import FormDialog from "../tgFolders/formDialog";
import useConfirm from "../../../hooks/useConfirm";
import {
    handleAsync
} from 'utils/tools'
// 账号分组
const TgFolders = () => {
    const [selected, setSelected] = useState<any>([]); // 多选
    const [rows, setrows] = useState([]); // table rows 数据
    const [paramsPayload, setParamsPayload] = useState({
        page: 1,
        pageSize: 10,
        id: undefined,
        folderName: undefined,
        createdAt: undefined
    }); // 分页
    const [searchForm, setSearchForm] = useState([]); // search Form
    const [pagetionTotle, setPagetionTotle] = useState(0); // total/ 弹窗控制
    const boxRef: any = useRef();
    const dispatch = useDispatch();
    const confirm = useConfirm(); // 弹窗
    // const navigate = useNavigate();
    const { tgFoldersList } = useSelector((state) => state.tg);
    let { height: boxHeight } = useHeightComponent(boxRef);

    useEffect(() => {
        getTableListActionFN();
        console.log('tgFoldersList', tgFoldersList?.data?.list);
    }, [dispatch, paramsPayload]);
    // 数据赋值
    useEffect(() => {
        // getTgSearchParams();
        setrows(tgFoldersList?.data?.list || []);
        setPagetionTotle(tgFoldersList?.data?.totalCount);
    }, [tgFoldersList]);
    // 网络请求
    useEffect(() => {
        getTgSearchParams();
    }, []);

    // tgUser 表格数据
    const getTableListActionFN = async () => {
        await dispatch(getTgFoldersListAction(paramsPayload));
    };
    // 分组选择请求
    const getTgSearchParams = async () => {
        try {
            const res = await axios.get(`/tg/tgFolders/list`);
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
            temp = value;
        } else if (key === 'isOnline') {
            temp = value;
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
    // 分页数量
    const PaginationCount = (count: number) => {
        return typeof count === 'number' ? Math.ceil(count / 10) : 1;
    }

    // 子传父 searchForm
    const handleSearchFormData = (obj: any) => {
        setParamsPayload({ ...paramsPayload, ...obj, page: 1 });
    };
    const [formDialogConfig, setFormDialogConfig] = useState<any>({
        title: '',
        edit: false,
        selectCheck: [],
        dialogType: undefined,

        params: undefined,
        renderField: undefined
    })
    // sendMsg
    const sendMsg = (msg: any = '~~', type: String = 'success') => {
        dispatch(openSnackbar({
            open: true,
            message: msg,
            variant: 'alert',
            alert: {
                color: type
            },
            close: false,
            anchorOrigin: {
                vertical: 'top',
                horizontal: 'center'
            },
            severity: type
        }))
    }
    // 弹窗开启
    const onBtnOpenList = useCallback(async (active: String, value: any = undefined) => {
        switch (active) {
            case 'Add':
                setFormDialogConfig({
                    ...formDialogConfig,
                    edit: true,
                    title: '添加',
                    dialogType: 'editForm',
                    type: 'Add',
                    params: { folderList: searchForm },
                });
                break
            case 'Edit':
                const { res, error } = await handleAsync(() => tgFoldersView({ id: value.id }))
                if (error) {
                    return sendMsg(error.message || '回显数据获取失败', 'error')
                }
                setFormDialogConfig({
                    ...formDialogConfig,
                    edit: true,
                    title: '编辑',
                    dialogType: 'editForm',
                    type: 'Edit',
                    params: { folderList: searchForm, echo: res?.data },
                });
                break
            default:
                break;
        }
    }, [formDialogConfig, searchForm])
    // 弹窗关闭
    const onBtnCloseList = useCallback((type: String, value: any) => {
        switch (type) {
            case 'Add':
                setFormDialogConfig({ ...formDialogConfig, edit: value, title: '', dialogType: '', type: '', prams: undefined });
                getTableListActionFN()
                break
            case 'Edit':
                setFormDialogConfig({ ...formDialogConfig, edit: value, title: '', dialogType: '', type: '', prams: undefined });
                getTableListActionFN()
                break
            default:
                break;
        }
    }, [formDialogConfig]);
    const memoizedFormDialogConfig = useMemo(() => formDialogConfig, [formDialogConfig]);
    const onExecuteClick = async (type: String, row: any = undefined, status: any = undefined) => {
        if (type === 'delete') {
            confirm('警告', `是否执行批量删除操作,该操作会批量删除数据且无法恢复，请谨慎操作。`)
                .then(async (result) => {
                    if (result) {
                        // 执行
                        const { res, error } = await handleAsync(() => tgFoldersDelete({ id: selected }))
                        if (error) {
                            return sendMsg(error?.message || '删除失败', 'error')
                        }
                        getTableListActionFN()
                        sendMsg(res.message)
                    } else {
                        console.log('Cancelled!');
                    }
                })
                .catch((error) => {
                    console.error('Error:', error);
                });
            return
        }
    }

    return (
        // <div>批量操作任务</div>
        <MainCard title={<FormattedMessageTitle />} content={true}>
            <div className={styles.box} ref={boxRef}>
                <div className={styles.searchTop}>
                    <SearchForm top100Films={searchForm} handleSearchFormData={handleSearchFormData} />
                </div>
                <div className={styles.btnList}>
                    <Stack direction="row" spacing={2}>
                        <Button size="small" variant="contained" onClick={e => onBtnOpenList('Add')}>
                            添加
                        </Button>
                        <Button size="small" variant="contained" disabled={selected.length > 0 ? false : true} onClick={e => onExecuteClick('delete')}>
                            批量删除
                        </Button>
                        {/*<Button size="small" variant="contained" disabled={true}>*/}
                        {/*    导出*/}
                        {/*</Button>*/}
                    </Stack>
                </div>
                <TableContainer
                    component={Paper}
                    style={{ maxHeight: `calc(${boxHeight - 150}px)`, borderTop: '1px solid #eaeaea', borderBottom: '1px solid #eaeaea' }}
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
                                                    <>
                                                        <IconButton onClick={e => onBtnOpenList('Edit', row)}>
                                                            <Tooltip title='编辑' placement="top">
                                                                <ModeEditIcon style={{ color: 'rgb(3, 106, 129)', fontSize: '18px' }} />
                                                            </Tooltip>
                                                        </IconButton>
                                                        <IconButton style={{ marginLeft: '5px' }}  >
                                                            <Tooltip title='删除' placement="top">
                                                                <DeleteIcon style={{ color: 'rgb(159, 86, 108)', fontSize: '18px' }} />
                                                            </Tooltip>
                                                        </IconButton>
                                                        <IconButton style={{ marginLeft: '5px' }}  >
                                                            <Tooltip title='详情' placement="top">
                                                                <DetailsIcon style={{ color: 'rgb(3, 106, 129)', fontSize: '18px' }} />
                                                            </Tooltip>
                                                        </IconButton>
                                                    </>
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
                {pagetionTotle !== 0 ? (
                    <>
                        <div className={styles.paginations}>
                            <div>共 {pagetionTotle} 条</div>
                            <Pagination count={PaginationCount(pagetionTotle)} color="primary" onChange={onPaginationChange} />
                        </div>
                    </>
                ) : (
                    ''
                )}
            </div>
            <FormDialog open={memoizedFormDialogConfig.edit} config={memoizedFormDialogConfig} onChangeDialogStatus={onBtnCloseList} />
        </MainCard>
    )
}
// 标题 tg
const FormattedMessageTitle = () => {
    return (
        <div className={styles.FormattedMessageTitle}>
            <FormattedMessage id="setting.cron.tg-folders" />
            {/* <div>
                <Button variant="outlined">登录</Button>
            </div> */}
        </div>
    );
};
export default memo(TgFolders)