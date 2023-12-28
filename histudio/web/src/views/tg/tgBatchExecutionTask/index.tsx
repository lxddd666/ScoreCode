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
    // Autocomplete
} from '@mui/material';
import { useDispatch, useSelector } from 'store';
import { useHeightComponent } from 'utils/tools';
import styles from './index.module.scss';
import SearchForm from './searchFrom';

import { getTgBatchExecutionTaskListAction } from 'store/slices/tg';
import axios from 'utils/axios';
import { columns } from './config';
import {tgBatchExecutionTaskDelete} from "../../../server/tg";
import AnimateButton from "../../../ui-component/extended/AnimateButton";
import useConfirm from "../../../hooks/useConfirm";
import FormDialog from "./formDialog";
import {openSnackbar} from "../../../store/slices/snackbar";

// 批量操作任务
const TgBatchExecutionTask = () => {
    const [selected, setSelected] = useState<any>([]); // 多选
    const [rows, setrows] = useState([]); // table rows 数据
    const [paramsPayload, setParamsPayload] = useState({
        page: 1,
        pageSize: 10,
        id: undefined,
        status: undefined,
        createdAt: undefined,
    }); // 分页
    const [searchForm, setSearchForm] = useState([]); // search Form
    const [pagetionTotle, setPagetionTotle] = useState(0); // total/ 弹窗控制
    const boxRef: any = useRef();
    const dispatch = useDispatch();
    // const navigate = useNavigate();
    const { tgBatchExecutionTaskList } = useSelector((state) => state.tg);
    const [checkListIsDisable, setCheckListIsDisable] = useState(true)
    const confirm = useConfirm(); // 弹窗
    let { height: boxHeight } = useHeightComponent(boxRef);

    useEffect(() => {
        getTableListActionFN();
        // console.log('getTgBatchExecutionTaskListAction', tgBatchExecutionTaskList?.data?.list);
    }, [dispatch, paramsPayload]);
    // 数据赋值
    useEffect(() => {
        // getTgSearchParams();
        setrows(tgBatchExecutionTaskList?.data?.list || []);
        setPagetionTotle(tgBatchExecutionTaskList?.data?.totalCount);
    }, [tgBatchExecutionTaskList]);
    // 网络请求
    useEffect(() => {
        getTgSearchParams();
    }, []);

    // tgUser 表格数据
    const getTableListActionFN = async () => {
        await dispatch(getTgBatchExecutionTaskListAction(paramsPayload));
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

        if (newSelected.length > 0) {
            setCheckListIsDisable(false)
        } else {
            setCheckListIsDisable(true)
        }
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
    const [formDialogConfig, setFormDialogConfig] = useState<any>({
        title: '',
        edit: false,
        selectCheck: [],
        dialogType: undefined,

        params: undefined,
        renderField: undefined
    })
    // 子传父 searchForm
    const handleSearchFormData = (obj: any) => {
        setParamsPayload({ ...paramsPayload, ...obj, page: 1 });
    };

    const PaginationCount = (count: number) => {
        return typeof count === 'number' ? Math.ceil(count / 10) : 1;
    }
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

    // 批量删除 操作
    const onUserAllDeleteClick = (id: any = undefined) => {
        confirm('警告', `是否批量删除选中的 ${selected.length} 个数据,确定之后不可取消,请谨慎操作。`)
            .then(async (result) => {
                if (result) {
                    try {
                        const res = await tgBatchExecutionTaskDelete({ id: selected })
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

    // 弹窗开启
    const onBtnOpenList = useCallback(async (active: String, value: any = undefined) => {
        // let columns = []
        switch (active) {
            case 'Add':
                setFormDialogConfig({
                    ...formDialogConfig,
                    edit: true,
                    title: '添加批量操作任务',
                    dialogType: 'editForm',
                    type: 'Add',
                    params: { ...searchForm}
                });
                break
            default:
                break;
        }
    }, [formDialogConfig])
    // 弹窗关闭
    const onBtnCloseList = useCallback((type: String, value: any) => {
        switch (type) {
            case 'Add':
                setFormDialogConfig({ ...formDialogConfig, edit: value, title: '', dialogType: '', type: '', prams: undefined });
                getTableListActionFN()
                sendMsg('添加成功')
                break;
        }
    }, [formDialogConfig]);
    const memoizedFormDialogConfig = useMemo(() => formDialogConfig, [formDialogConfig]);
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
                        <AnimateButton type="slide">
                            <Button size="small" variant="contained" disabled={checkListIsDisable} color="error" onClick={onUserAllDeleteClick}>
                                批量删除
                            </Button>
                        </AnimateButton>
                        <Button size="small" variant="contained" disabled={true}>
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
// 批量操作标题
const FormattedMessageTitle = () => {
    return (
        <div className={styles.FormattedMessageTitle}>
            <FormattedMessage id="setting.cron.tg-batch-execution-task" />
        </div>
    );
};
export default memo(TgBatchExecutionTask)