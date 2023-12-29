
// material-ui
import { memo, useEffect, useState, useRef, useCallback } from "react"

import {
    Dialog,
    DialogTitle,
    DialogContent,
    Button, Grid, TextField, Stack,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Chip,
    Avatar,
    Radio,
    Pagination,
    FormControl,
    InputLabel,
    Select,
    OutlinedInput,
    MenuItem,
    // Checkbox,
    ListItemText,
    FormHelperText,
    // ButtonGroup,
    // Autocomplete
} from '@mui/material';
// import ClearIcon from '@mui/icons-material/Clear';
import AutorenewIcon from '@mui/icons-material/Autorenew';
import SearchIcon from '@mui/icons-material/Search';
import AnimateButton from 'ui-component/extended/AnimateButton';
import LoadingButton from '@mui/lab/LoadingButton';
import { openSnackbar } from 'store/slices/snackbar';

import { gridSpacing } from 'store/constant';
import Paper from '@mui/material/Paper';
import { styled } from '@mui/material/styles';
import { useSelector, useDispatch } from 'store';
import { getUserList } from 'store/slices/user';
import Badge from '@mui/material/Badge';
import styles from './index.module.scss';

import { useFormik } from 'formik';
import * as yup from 'yup';

import { timeAgo } from 'utils/tools'
import {
    // tgtgKeepTaskEdit,
    tgIncreaseFansCronCheckChannel,
    tgIncreaseFansCronChannelIncreaseFanDetail,
    tgIncreaseFansCronEdit
} from 'server/tg'

import { getProxyListAction } from "store/slices/org";
import {
    handleAsync
} from 'utils/tools'

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));
const createValidationSchema = (requiredFields: any) => {
    return yup.object().shape({
        taskName: yup.string().trim('Enter a valid username').test('is-required', 'Cron is required', (value) => {
            return requiredFields.taskName ? !!value : true;
        }),
        // cron: yup.string().trim('Enter a valid cron').required('cron is required'),
        // actions: yup.array().of(
        //     yup.number().required('Action is required')
        // ).required('At least one action is required')
        //     .min(1, 'At least one action is required'),
        accounts: yup.array().of(
            yup.number().required('Action is required')
        ).test(
            'is-accounts-required',
            'At least one account is required',
            (value) => !requiredFields.accounts || (Array.isArray(value) && value?.length > 0)
        ),
        folderId: yup.string().trim('Enter a valid cron').test('is-required', 'Cron is required', (value) => {
            return requiredFields.folderId ? !!value : true;
        }),
        // scriptGroup: yup.string().trim('Enter a valid scriptGroup').required('scriptGroup is required'),
    });
}

const FormDialog = (props: any) => {
    const { open, config, onChangeDialogStatus } = props
    const [dialogValue, setDialogValue] = useState<any>({})
    const [formikValue, setFormikValue] = useState<any>({
        id: undefined,
        channel: '',
        taskName: '',
        channelMemberCount: '',
        fansCount: '',
        folderId: '',
        channelId: '',
        dayCount: '',
        executedPlan: '',
    })
    const [requiredFields] = useState({
        accounts: false,
        folderId: false,
        channel: false,
        taskName: false,
    })
    const dispatch = useDispatch()
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
    console.log('dormDialog', open, config);
    useEffect(() => {
        setFormikValue({
            ...formikValue,
            ...config.params?.echo
        })
    }, [config.params])
    const validationSchema = createValidationSchema(requiredFields)
    const formik = useFormik({
        initialValues: formikValue,
        validationSchema,
        enableReinitialize: true,
        onSubmit: async (values) => {
            console.log('value', values);
            if (config.type === 'Add') {
                try {
                    const res: any = await tgIncreaseFansCronEdit(values)
                    console.log(res);
                    onChangeDialogStatus(config.type, false)
                    sendMsg(res.message)
                } catch (error: any) {
                    console.log('error', error);
                    sendMsg(error.message || '添加失败', 'error')
                }

            }
            if (config.type === 'Edit') {
                try {
                    const res: any = await tgIncreaseFansCronEdit(values)
                    console.log(res);
                    onChangeDialogStatus(config.type, false)
                    sendMsg(res.message)
                } catch (error: any) {
                    console.log('error', error);
                    sendMsg(error.message || '编辑失败', 'error')
                }

            }
        }
    });

    // 关闭弹窗
    const handleClose = () => {
        // console.log('关闭弹窗');
        onChangeDialogStatus(config.type, false)
        formik.resetForm();
        setFormikValue({
            id: undefined,
            channel: '',
            taskName: '',
            channelMemberCount: '',
            fansCount: '',
            folderId: '',
            channelId: '',
            dayCount: '',
            executedPlan: '',
        })
    }
    const changeSubmitValue = useCallback((value: any) => {
        setDialogValue(value)
    }, [dialogValue])
    return (
        <>
            {
                open && (<Dialog open={open} onClose={(event: any, reason: any) => {
                    if (reason !== 'backdropClick' && reason !== 'escapeKeyDown') {
                        // handleImportClose();
                    }
                }}
                    maxWidth="sm"
                    sx={{
                        '& .MuiDialog-paper': { width: '80%' }, // 80% 的宽度
                        '& .css-meoh0q-MuiPaper-root-MuiDialog-paper': { maxWidth: '1000px' }
                    }}
                    fullWidth={true}>
                    <DialogTitle>{config.title}</DialogTitle>
                    <DialogContent>
                        {
                            config.dialogType && config.dialogType === 'editForm'
                                ?
                                <EditForm changeSubmitValue={changeSubmitValue}
                                    row={config.params}
                                    renderField={config.renderField}
                                    formik={formik}
                                    handleClose={handleClose}
                                />
                                :
                                <EditTable changeSubmitValue={changeSubmitValue} columns={config.columns} type={config.type} />
                        }

                    </DialogContent>
                    {/* <DialogActions>
                        <Button onClick={handleClose}>取消</Button>
                        <Button type="submit">提交</Button>
                    </DialogActions> */}

                    {/* </form> */}
                </Dialog>)
            }
        </>
    )
}


// table header options

const renderTable = (value: any, key: any, item: any) => {
    // console.log(value, key, item);

    let temp: any = '';
    if (key === 'username') {
        temp = <div className={styles.tablesColumns}>
            <div className={styles.avatars}>
                <StyledBadge
                    overlap="circular"
                    anchorOrigin={{ vertical: 'bottom', horizontal: 'right' }}
                    variant="dot"
                    badgeColor={item.status === 1 ? '#44b700' : 'red'}
                >
                    {/* <Avatar alt="Remy Sharp" src="https://berrydashboard.io/assets/avatar-1-8ab8bc8e.png"> */}
                    <Avatar alt="Remy Sharp" src={item.avatar}>
                        {item?.username?.charAt(0).toUpperCase()}
                    </Avatar>
                </StyledBadge>
            </div>
            <div className={styles.info}>
                <div className={styles.titles}>
                    <p>{item.username}</p>
                    <p style={{ marginLeft: '5px' }}>{item.realName}</p>
                </div>
                <div style={{ fontSize: '12px' }}>email:{item.email}</div>
            </div>
        </div>
    }
    else if (key === 'status') {
        temp = <Chip label={value === 1 ? '正常' : '异常'} color="primary" sx={{ bgcolor: `${item.status === 1 ? '#44b700' : 'red'}`, color: 'white' }} />;
    } else if (key === 'lastActiveAt') {
        temp = timeAgo(value)
    } else {
        temp = value;
    }
    return temp
};
const EditTable = (props: any) => {
    const { changeSubmitValue, columns, type } = props
    const dispatch = useDispatch();
    const [formData, setFormData] = useState<any>({
        username: undefined,
    })
    // 用于跟踪选中行的ID
    const [selectedId, setSelectedId] = useState(null);
    const { userList } = useSelector((state) => state.user);
    const { proxyList } = useSelector((state) => state.org);
    const [userListRow, setUserListRow] = useState<any>([])
    const [pagetionTotle, setPagetionTotle] = useState(0); // total
    useEffect(() => {
        if (type === 'bindUser') {
            fetchBindUser()
        }
        if (type === 'bindProxy') {
            fetchBindProxy()
        }

    }, [])
    const fetchBindUser = async (page = 1, value: any = undefined,) => {
        const updatedParams = [
            `page=${page}`,
            `pageSize=${10}`,
            `${value ? `username=${value}` : ''}`,
        ];
        await dispatch(getUserList(updatedParams.filter((query) => !query.endsWith('=') && query !== 'status=0').join('&')))
    }
    const fetchBindProxy = async (page = 1, value: any = undefined,) => {
        const updatedParams = {
            page: page,
            pageSize: 10,
            address: value
        }
        await dispatch(getProxyListAction(updatedParams))
    }
    useEffect(() => {
        if (userList?.data?.list) {
            setUserListRow(userList?.data?.list || [])
            setPagetionTotle(userList?.data?.totalCount || 0)
        }
    }, [userList])
    useEffect(() => {
        if (proxyList?.data?.list) {
            setUserListRow(proxyList?.data?.list || [])
            setPagetionTotle(proxyList?.data?.totalCount || 0)
        }
    }, [proxyList])

    const handleSelectRow = (id: any) => {
        // console.log(id);
        setSelectedId(id);
        changeSubmitValue(id)
    };

    // 搜索按钮
    const onSearchClick = (e: any) => {
        console.log(e.target.value, formData);

        if (type === 'bindUser') {
            fetchBindUser(1, formData.username)
        }
        if (type === 'bindProxy') {
            fetchBindProxy(1, formData.username)
        }

    };
    // 重置按钮
    const onResetClick = (e: any) => {
        let obj = {
            username: undefined,
        };
        // setValue({});
        setFormData(obj);
        if (type === 'bindUser') {
            fetchBindUser()
        }
        if (type === 'bindProxy') {
            fetchBindProxy()
        }
    };
    // 分页事件
    const pageRef = useRef(1);
    const onPaginationChange = (event: object, page: number) => {
        pageRef.current = page;
        if (type === 'bindUser') {
            fetchBindUser(pageRef.current)
        }
        if (type === 'bindProxy') {
            fetchBindProxy(pageRef.current)
        }
    };
    // 分页数量
    const PaginationCount = (count: number) => {
        return typeof count === 'number' ? Math.ceil(count / 10) : 1;
    };
    return (
        <div style={{ height: '600px' }}>
            <Grid container spacing={gridSpacing} alignItems="center">
                <Grid item xs={3}>
                    <Item>
                        <TextField
                            autoFocus
                            sx={{ width: '100%' }}
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
                        />
                    </Item>
                </Grid>
                <Grid item xs={3}>
                    <Item>
                        <Stack direction="row" spacing={2}>
                            <Button size="small" variant="outlined" startIcon={<SearchIcon />} onClick={onSearchClick}>
                                查询
                            </Button>
                            <Button size="small" variant="outlined" startIcon={<AutorenewIcon />} onClick={onResetClick}>
                                重置
                            </Button>
                        </Stack>
                    </Item>
                </Grid>

            </Grid>
            <div style={{ width: '100%', maxHeight: '400px' }}>
                <TableContainer
                    component={Paper}
                    style={{ maxHeight: 480 }}
                >
                    <Table aria-label="simple table" sx={{ border: 1, borderColor: 'divider' }} stickyHeader={true}>
                        <TableHead>
                            <TableRow>
                                <TableCell>#</TableCell>
                                {columns.map((item: any) => {
                                    return (
                                        <TableCell align="center" key={item.id}>
                                            {item.label}
                                        </TableCell>
                                    );
                                })}
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {userListRow && userListRow.map((row: any) => {
                                return (
                                    <TableRow
                                        key={row.id}
                                        selected={selectedId === row.id}
                                        onClick={() => handleSelectRow(row.id)}
                                        hover
                                    >
                                        <TableCell padding="checkbox">
                                            <Radio
                                                checked={selectedId === row.id}
                                                onChange={() => handleSelectRow(row.id)}
                                            />
                                        </TableCell>
                                        {columns.map((item: any) => {
                                            return (
                                                <TableCell align="center" key={item.label}>
                                                    {renderTable(row[item.id], item.id, row)}
                                                </TableCell>
                                            );
                                        })}

                                    </TableRow>
                                )
                            })}


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
        </div>
    )
}

// const ITEM_HEIGHT = 48;
// const ITEM_PADDING_TOP = 8;
// const MenuProps = {
//     PaperProps: {
//         style: {
//             maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
//             width: 250,
//         },
//     },
// };
const EditForm = (props: any) => {
    const { row, formik, handleClose } = props
    const dispatch = useDispatch();

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

    // const inputChange = (event: any) => {
    //     formik.handleChange(event)
    //     changeSubmitValue({
    //         id: row?.id || undefined,
    //         ...formik.values
    //     })
    // }

    // 校验频道
    const checkBtn = async (type: String) => {
        console.log('fir', formik.values);

        if (type === 'check') {
            const { res, error } = await handleAsync(() => tgIncreaseFansCronCheckChannel({ ...formik.values }))
            if (error) {
                return sendMsg(error.message || '失败', 'error')
            }
            console.log('res执行成功', res);
            const info = res?.data?.channelMsg
            formik.setFieldValue('channelMemberCount', res?.data?.channelMsg?.channelMemberCount);
            formik.setFieldValue('channelId', res?.data?.channelMsg?.channelId);
            sendMsg(`频道有效,频道Title:${info.channelTitle},频道Id:${info.channelId},频道人数:${info.channelMemberCount}`)
        }
        if (type === 'fanDetail') {
            const { res, error } = await handleAsync(() => tgIncreaseFansCronChannelIncreaseFanDetail({ ...formik.values }))
            if (error) {
                return sendMsg(error.message || '失败', 'error')
            }
            console.log('res执行成功', res);
            const info = res?.data
            formik.setFieldValue('executedPlan', info.dailyIncreaseFan);
            sendMsg('已设置涨粉计划')
        }
    }

    return (
        <form onSubmit={formik.handleSubmit} style={{ marginTop: '10px' }}>

            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="channel"
                        name="channel"
                        label='频道地址'
                        value={formik.values.channel}
                        onChange={(event) => {
                            const value = event.target.value;
                            console.log(event.target.value);

                            formik.setFieldValue('channel', value);
                        }}
                        onBlur={formik.handleBlur}
                        InputProps={{
                            endAdornment: <LoadingButton style={{ width: '200px' }} variant="contained" onClick={e => checkBtn('check')} >
                                校验频道
                            </LoadingButton>,
                        }}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="taskName"
                        name="taskName"
                        label='任务名称'
                        value={formik.values.taskName}
                        onChange={(event) => {
                            const value = event.target.value;
                            formik.setFieldValue('taskName', value);
                        }}
                        onBlur={formik.handleBlur}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        type="number"
                        inputProps={{
                            min: 0, // 最小值
                            max: 10000, // 最大值
                            step: 1 // 步长
                        }}
                        fullWidth
                        id="channelMemberCount"
                        name="channelMemberCount"
                        label='频道当前粉丝数'
                        value={formik.values.channelMemberCount}
                        disabled={true}
                        onChange={(event) => {
                            const value = event.target.value;
                            formik.setFieldValue('channelMemberCount', value);
                        }}
                        onBlur={formik.handleBlur}
                    />
                </Grid>
                <Grid item xs={12}>
                    <FormControl sx={{ width: '100%' }} error={Boolean(formik.errors.folderId)}>
                        <InputLabel id="demo-multiple-checkbox-label">分组</InputLabel>
                        <Select
                            labelId="folderId"
                            id="folderId"
                            value={formik.values.folderId || ''}
                            onChange={(event) => {
                                const value = event.target.value;
                                formik.setFieldValue('folderId', value);
                            }}
                            input={<OutlinedInput label="分组" />}

                        >
                            {row?.folderList?.map((item: any) => (
                                <MenuItem key={item.value} value={item.value}>
                                    <ListItemText primary={item.title} />
                                </MenuItem>
                            ))}
                        </Select>
                        {formik.touched.folderId && formik.errors.folderId ? (
                            <FormHelperText>{formik.errors.folderId}</FormHelperText>
                        ) : null}
                    </FormControl>
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="channelId"
                        name="channelId"
                        label='频道ID'
                        disabled={true}
                        value={formik.values.channelId}
                        onChange={(event) => {
                            const value = event.target.value;
                            formik.setFieldValue('channelId', value);
                        }}
                        onBlur={formik.handleBlur}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        type="number"
                        inputProps={{
                            min: 0, // 最小值
                            max: 10000, // 最大值
                            step: 1 // 步长
                        }}
                        fullWidth
                        id="dayCount"
                        name="dayCount"
                        label='持续天数'
                        value={formik.values.dayCount}
                        onChange={(event) => {
                            const value = event.target.value;
                            formik.setFieldValue('dayCount', value);
                        }}
                        onBlur={formik.handleBlur}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="executedPlan"
                        name="executedPlan"
                        label='执行计划'
                        disabled={true}
                        value={formik.values.executedPlan}
                        onChange={(event) => {
                            const value = event.target.value;
                            formik.setFieldValue('executedPlan', value);
                        }}
                        onBlur={formik.handleBlur}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        type="number"
                        inputProps={{
                            min: 0, // 最小值
                            max: 10000, // 最大值
                            step: 1 // 步长
                        }}
                        fullWidth
                        id="fansCount"
                        name="fansCount"
                        label='涨粉数量'
                        value={formik.values.fansCount}
                        onChange={(event) => {
                            const value = event.target.value;
                            console.log(event.target.value);

                            formik.setFieldValue('fansCount', value);
                        }}
                        onBlur={formik.handleBlur}
                        InputProps={{
                            endAdornment: <>
                                <LoadingButton style={{ width: '200px' }} variant="contained" onClick={e => checkBtn('fanDetail')} >
                                    增长数量计算
                                </LoadingButton>
                            </>,
                        }}
                    />
                </Grid>
                <Grid item xs={12}>
                    <Stack direction="row" justifyContent="flex-end">
                        <AnimateButton>
                            <Button onClick={handleClose}>
                                取消
                            </Button>
                        </AnimateButton>
                        <AnimateButton>
                            <Button type="submit">
                                提交
                            </Button>
                        </AnimateButton>
                    </Stack>
                </Grid>
            </Grid>
        </form >
    )
}

interface StyledBadgeProps {
    badgeColor?: string; // 这是你的自定义属性
}
const StyledBadge = styled(Badge, {
    shouldForwardProp: (prop) => prop !== 'badgeColor',
})<StyledBadgeProps>(({ theme, badgeColor }) => ({
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
export default memo(FormDialog)