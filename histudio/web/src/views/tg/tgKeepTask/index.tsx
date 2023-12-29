import { memo, useState, useRef, useEffect, useCallback, useMemo } from "react"
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
    Chip,
    Pagination,
    Menu,
    MenuItem,
    Fade
} from '@mui/material';
import AnimateButton from 'ui-component/extended/AnimateButton';
import { useDispatch, useSelector } from 'store';
import { useHeightComponent } from 'utils/tools';
// import { createFilterOptions } from '@mui/material/Autocomplete';
// import { openSnackbar } from 'store/slices/snackbar';
import styles from './index.module.scss';
import SearchForm from './searchFrom';

import { getTgKeepTaskListAction } from 'store/slices/tg';
import { getScriptGroupListAction } from 'store/slices/script';
import { openSnackbar } from 'store/slices/snackbar';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
// import axios from 'utils/axios';
import { columns } from './config';
import FormDialog from './formDialog'
import {
    tgFoldersList,
    tgUserList,
    tgDictDataOptions,
    tgtgKeepTaskEditEcho,
    tgKeepTaskExecute,
    tgKeepTaskExecuteOne,
    tgKeepTaskAllDelete
} from 'server/tg'
import {
    handleAsync
} from 'utils/tools'
import useConfirm from 'hooks/useConfirm'

// 养号任务
const TgKeepTask = () => {
    const [selected, setSelected] = useState<any>([]); // 多选
    const [rows, setrows] = useState([]); // table rows 数据
    const [paramsPayload, setParamsPayload] = useState({
        page: 1,
        pageSize: 10,
        taskName: undefined,
        actions: undefined,
        accounts: undefined,
        status: undefined,
        createdAt: undefined,
    }); // 分页
    const [formDialogConfig, setFormDialogConfig] = useState<any>({
        title: '',
        edit: false,
        selectCheck: [],
        dialogType: undefined,

        params: undefined,
        renderField: undefined
    })
    const [searchForm, setSearchForm] = useState({
        folderList: [],
        keepActionsList: {},
        accountsList: [],
        scriptList: []
    }); // search Form
    const [pagetionTotle, setPagetionTotle] = useState(0); // total/ 弹窗控制
    const [anchorEl, setAnchorEl] = useState(null); // menu 菜单
    const boxRef: any = useRef();
    const dispatch = useDispatch();
    // const navigate = useNavigate();
    const { tgKeepTaskList } = useSelector((state) => state.tg);
    const { scriptList } = useSelector((state) => state.script);
    const confirm = useConfirm(); // 弹窗
    let { height: boxHeight } = useHeightComponent(boxRef);

    useEffect(() => {
        getTableListActionFN();
        getScriptGroupListActionFN()
        // console.log('tgKeepTaskList', tgKeepTaskList?.data?.list);
    }, [dispatch, paramsPayload]);
    // 数据赋值
    useEffect(() => {
        // getTgSearchParams();
        setrows(tgKeepTaskList?.data?.list || []);
        setPagetionTotle(tgKeepTaskList?.data?.totalCount);
    }, [tgKeepTaskList]);
    // 网络请求
    useEffect(() => {
        getTgSearchParams();
    }, []);

    // tgUser 表格数据
    const getTableListActionFN = async () => {
        await dispatch(getTgKeepTaskListAction(paramsPayload));
    };
    // 话术分组
    const getScriptGroupListActionFN = async () => {
        await dispatch(getScriptGroupListAction(2));
    };
    // 查询条件请求 分组选择/账号/养号任务
    const getTgSearchParams = async () => {
        let promises = [
            await tgFoldersList({ page: 1, pageSize: 999 }),
            await tgUserList({ page: 1, pageSize: 999 }),
            await tgDictDataOptions(['keep_action', 'sys_job_status'])
        ];

        let storedResults: any = [];
        await Promise.allSettled(promises)
            .then((results) => {
                results.forEach((result: any) => {
                    if (result.status === 'fulfilled') {
                        // console.log('Promise fulfilled: ', result.value);
                        storedResults.push(result.value.data)
                    } else if (result.status === 'rejected') {
                        console.log('Promise rejected: ', result.reason);
                    }
                });
            });


        // console.log('storedResults', storedResults);

        let folderListArr: any = [{
            title: '请选择',
            value: ''
        }];
        storedResults[0]?.data?.list.map((item: any) => {
            folderListArr.push({
                title: item.folderName,
                value: item.id
            });
        });

        await setSearchForm({
            ...searchForm,
            folderList: folderListArr,
            accountsList: storedResults[1]?.list,
            keepActionsList: storedResults[2]
        });
    };
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
            }
        }))
    }

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
        if (key === 'status') {
            temp = <Chip color={value === 1 ? "secondary" : 'warning'} label={value === 1 ? '运行中' : '已结束'} />;
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

    // 弹窗开启
    const onBtnOpenList = useCallback(async (active: String, value: any = undefined) => {
        // let columns = []
        switch (active) {
            case 'Add':
                setFormDialogConfig({
                    ...formDialogConfig,
                    edit: true,
                    title: '添加任务',
                    dialogType: 'editForm',
                    type: 'Add',
                    params: { ...searchForm, scriptList: scriptList?.data?.list },
                });
                break
            case 'Edit':
                const res: any = await tgtgKeepTaskEditEcho({ id: value.id });
                // console.log(res.data);

                setFormDialogConfig({
                    ...formDialogConfig,
                    edit: true,
                    title: '编辑任务',
                    dialogType: 'editForm',
                    type: 'Edit',
                    params: { ...searchForm, scriptList: scriptList?.data?.list, echo: { ...res.data, status: 2 } },
                });
                break
            default:
                break;
        }
    }, [formDialogConfig, scriptList, searchForm])
    // 弹窗关闭
    const onBtnCloseList = useCallback((type: String, value: any) => {
        // console.log(type, value);
        switch (type) {
            case 'Add':
                setFormDialogConfig({ ...formDialogConfig, edit: value, title: '', dialogType: '', type: '', prams: undefined });
                getTableListActionFN()
                // sendMsg('添加成功')
                break
            case 'Edit':
                setFormDialogConfig({ ...formDialogConfig, edit: value, title: '', dialogType: '', type: '', prams: undefined });
                getTableListActionFN()
                // sendMsg('编辑成功')
                break
            default:
                break;
        }
    }, [formDialogConfig]);
    const memoizedFormDialogConfig = useMemo(() => formDialogConfig, [formDialogConfig]);

    /**
     * 执行 执行一次 
     * @param type 类型
     * @param row 执行id
     * @param status 要改变的状态
     */
    const onExecuteClick = async (type: String, row: any = undefined, status: any = undefined) => {
        if (type === 'execute') {
            const { res, error } = await handleAsync(() => tgKeepTaskExecute({ id: row.id, status: status }))
            if (error) {
                return sendMsg('执行失败', 'error')
            }
            console.log('res执行成功', res);
            getTableListActionFN()
            sendMsg('执行成功')
            return
        }
        if (type === 'executeOne') {
            confirm('警告', `是否仅执行一次该养号 [ ${row?.taskName || ''} ] 任务。`)
                .then(async (result) => {
                    if (result) {
                        // 执行
                        const { res, error } = await handleAsync(() => tgKeepTaskExecuteOne({ id: row.id }))
                        if (error) {
                            return sendMsg('执行失败', 'error')
                        }
                        console.log('res执行成功', res);
                        getTableListActionFN()
                        sendMsg('执行成功')
                    } else {
                        console.log('Cancelled!');
                    }
                })
                .catch((error) => {
                    console.error('Error:', error);
                });
            return
        }
        if (type === 'delete') {
            confirm('警告', `是否执行批量删除操作,该操作会批量删除数据且无法恢复，请谨慎操作。`)
                .then(async (result) => {
                    if (result) {
                        // 执行
                        const { res, error } = await handleAsync(() => tgKeepTaskAllDelete({ id: selected }))
                        if (error) {
                            return sendMsg('删除失败', 'error')
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
        if (type === 'deleteSingle') {
            confirm('警告', `是否执行批量删除操作,该操作会批量删除数据且无法恢复，请谨慎操作。`)
                .then(async (result) => {
                    if (result) {
                        // 执行
                        const { res, error } = await handleAsync(() => tgKeepTaskAllDelete({ id: row.id }))
                        if (error) {
                            return sendMsg('删除失败', 'error')
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


    const opens = Boolean(anchorEl);
    const handleMenuClick = (event: any) => {
        setAnchorEl(event.currentTarget);
    };
    const handleMenuClose = () => {
        setAnchorEl(null);
    };
    return (
        // <div>批量操作任务</div>
        <MainCard title={<FormattedMessageTitle />} content={true}>
            <div className={styles.box} ref={boxRef}>
                <div className={styles.searchTop}>
                    <SearchForm searchForm={searchForm} handleSearchFormData={handleSearchFormData} />
                </div>
                <div className={styles.btnList}>
                    <Stack direction="row" spacing={2}>
                        <Button size="small" variant="contained" onClick={e => onBtnOpenList('Add')}>
                            添加
                        </Button>
                        <Button size="small" variant="contained" disabled={selected.length > 0 ? false : true} onClick={e => onExecuteClick('delete')}>
                            批量删除
                        </Button>
                        {/* <Button size="small" variant="contained" disabled={true}>
                            导出
                        </Button> */}
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
                                        <TableCell sx={{ maxWidth: 200 }} align="center" key={item.title}>
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
                                                    <div style={item.key === 'active' ? { display: 'flex', flexDirection: 'row', justifyContent: 'center' } : {}}>
                                                        <AnimateButton scale={{
                                                            hover: 1.1,
                                                            tap: 0.9
                                                        }}>
                                                            <Button size="small" variant="contained" onClick={e => onBtnOpenList('Edit', row)}>
                                                                编辑
                                                            </Button>
                                                        </AnimateButton>
                                                        <AnimateButton scale={{
                                                            hover: 1.1,
                                                            tap: 0.9
                                                        }}>
                                                            {
                                                                row.status === 2 ? <Button size="small"
                                                                    variant="contained"
                                                                    color="error"
                                                                    style={{ marginLeft: '5px' }}
                                                                    onClick={e => onExecuteClick('execute', row, 1)}>
                                                                    执行
                                                                </Button>
                                                                    :
                                                                    <Button size="small"
                                                                        variant="contained"
                                                                        color="error"
                                                                        style={{ marginLeft: '5px' }}
                                                                        onClick={e => onExecuteClick('execute', row, 2)}>
                                                                        暂停
                                                                    </Button>
                                                            }
                                                        </AnimateButton>
                                                        <AnimateButton scale={{
                                                            hover: 1.1,
                                                            tap: 0.9
                                                        }}>
                                                            <Button size="small"
                                                                variant="contained"
                                                                color="error"
                                                                style={{ marginLeft: '5px' }}
                                                                onClick={e => onExecuteClick('executeOne', row)}>
                                                                执行一次
                                                            </Button>
                                                        </AnimateButton>
                                                        <Button size="small"
                                                            variant="outlined"
                                                            color="secondary"
                                                            style={{ marginLeft: '5px' }}
                                                            id="basic-menu"
                                                            aria-controls="basic-menu"
                                                            aria-haspopup="true"
                                                            aria-expanded={opens ? 'true' : undefined}
                                                            onClick={handleMenuClick}
                                                            endIcon={<KeyboardArrowDownIcon />}>
                                                            更多
                                                        </Button>
                                                        <Menu
                                                            anchorEl={anchorEl}
                                                            open={opens}
                                                            onClose={handleMenuClose}
                                                            TransitionComponent={Fade}
                                                        >
                                                            <MenuItem onClick={handleMenuClose}>详情</MenuItem>
                                                            <MenuItem onClick={handleMenuClose}>删除</MenuItem>
                                                        </Menu>
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
            <FormattedMessage id="setting.cron.tg-keep-task" />
            {/* <div>
                <Button variant="outlined">登录</Button>
            </div> */}
        </div>
    );
};

export default memo(TgKeepTask)