
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
    Checkbox,
    ListItemText,
    FormHelperText,
    ButtonGroup,
    // Autocomplete
} from '@mui/material';
// import ClearIcon from '@mui/icons-material/Clear';
import AutorenewIcon from '@mui/icons-material/Autorenew';
import SearchIcon from '@mui/icons-material/Search';
import AnimateButton from 'ui-component/extended/AnimateButton';

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
    tgtgKeepTaskEdit,

} from 'server/tg'

import { getProxyListAction } from "store/slices/org";

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));
const createValidationSchema = (requiredFields: any) => {
    return yup.object().shape({
        taskName: yup.string().trim('Enter a valid username').required('username is required'),
        cron: yup.string().trim('Enter a valid cron').required('cron is required'),
        actions: yup.array().of(
            yup.number().required('Action is required')
        ).required('At least one action is required')
            .min(1, 'At least one action is required'),
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
        scriptGroup: yup.string().trim('Enter a valid scriptGroup').required('scriptGroup is required'),
    });
}

const FormDialog = (props: any) => {
    const { open, config, onChangeDialogStatus } = props
    const [dialogValue, setDialogValue] = useState<any>({})
    const [formikValue, setFormikValue] = useState<any>({
        id: undefined,
        taskName: '',
        cron: '',
        actions: [],
        accounts: [],
        folderId: '',
        scriptGroup: ''
    })
    const [requiredFields] = useState({
        accounts: false,
        folderId: false,
    })

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
                    const res = await tgtgKeepTaskEdit(values)
                    console.log(res);
                    onChangeDialogStatus(config.type, false)
                } catch (error) {
                    console.log('error', error);

                }

            }
            if (config.type === 'Edit') {
                try {
                    const res = await tgtgKeepTaskEdit(values)
                    console.log(res);
                    onChangeDialogStatus(config.type, false)
                } catch (error) {
                    console.log('error', error);

                }

            }
        }
    });
    // 提交表单
    // const handleSubmit = () => {
    //     formik.handleSubmit()
    //     console.log('提交表单', dialogValue);
    //     // onChangeDialogStatus('bind', false)
    //     if (!dialogValue) {
    //         return alert('提交失败，请选择用户')
    //     }
    //     if (config.type === 'Edit') {
    //         let data = {
    //             ...dialogValue
    //         }

    //         tgUserEdit(data).then(res => {
    //             onChangeDialogStatus(config.type, false)
    //         }).catch(err => {
    //             console.log('失败', err);

    //         })
    //     }
    // }
    // 关闭弹窗
    const handleClose = () => {
        // console.log('关闭弹窗');
        onChangeDialogStatus(config.type, false)
        formik.resetForm();
        setFormikValue({
            id: undefined,
            taskName: '',
            cron: '',
            actions: [],
            accounts: [],
            folderId: '',
            scriptGroup: ''
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
    // return <Tooltip title={temp} placement="top-start">
    //     <p>{temp}</p>
    // </Tooltip>;
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

const ITEM_HEIGHT = 48;
const ITEM_PADDING_TOP = 8;
const MenuProps = {
    PaperProps: {
        style: {
            maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
            width: 250,
        },
    },
};
const EditForm = (props: any) => {
    const { row, changeSubmitValue, formik, handleClose } = props
    // console.log('EditForm', formik.initialValues);


    const inputChange = (event: any) => {
        formik.handleChange(event)
        changeSubmitValue({
            id: row?.id || undefined,
            ...formik.values
        })
    }

    const selectBtn = (event: any) => {
        const { value } = event.target
        console.log(value, value === '50', value === 50);

        let tempArr: [] = []
        if (value === '50') {
            tempArr = row?.accountsList?.slice(0, 50)
            let arr: any = []
            tempArr?.map((item: any) => {
                arr.push(String(item.id))
            })
            formik.setFieldValue('accounts', arr);
        }
        if (value === 'clear') {
            formik.setFieldValue('accounts', []);
        }
    }

    return (
        <form onSubmit={formik.handleSubmit} style={{ marginTop: '10px' }}>

            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="taskName"
                        name="taskName"
                        label='任务名称'
                        value={formik.values.taskName}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.taskName && Boolean(formik.errors.taskName)}
                        helperText={formik.touched.taskName && formik.errors.taskName}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="cron"
                        name="cron"
                        label='表达式'
                        value={formik.values.cron}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.cron && Boolean(formik.errors.cron)}
                        helperText={formik.touched.cron && formik.errors.cron}
                    />
                </Grid>
                <Grid item xs={12}>
                    {/* <Autocomplete
                        options={row?.keepActionsList?.keep_action} // 你的选项数组
                        getOptionLabel={(option: any) => option.label || ''} // 如何获取选项的标签
                        value={formik.values.actions || ''}
                        onChange={(event, value) => {
                            console.log(value);

                            formik.setFieldValue('actions', value.value);
                        }}
                        renderInput={(params) => <TextField
                            {...params}
                            fullWidth
                            id="actions"
                            name="actions"
                            onChange={inputChange}
                            label="养号动作"
                            onBlur={formik.handleBlur}
                            error={formik.touched.actions && Boolean(formik.errors.actions)}
                            helperText={formik.touched.actions && formik.errors.actions} />}
                    /> */}
                    <FormControl sx={{ width: '100%' }} error={Boolean(formik.errors.actions)}>
                        <InputLabel id="demo-multiple-checkbox-label">养号动作</InputLabel>
                        <Select
                            labelId="demo-multiple-checkbox-label"
                            id="demo-multiple-checkbox"
                            multiple
                            value={formik.values.actions || []}
                            onChange={(event) => {
                                const value = event.target.value;
                                formik.setFieldValue('actions', value);
                            }}
                            input={<OutlinedInput label="养号动作" />}
                            renderValue={(selected) => selected
                                .map((value: any) => row?.keepActionsList?.keep_action.find((item: any) => item.value === value)?.label)
                                .join(', ')}
                            MenuProps={MenuProps}

                        >
                            {row?.keepActionsList?.keep_action.map((item: any) => (
                                <MenuItem key={item.key} value={item.value}>
                                    <Checkbox checked={formik.values.actions?.includes(item.value)} />
                                    <ListItemText primary={item.label} />
                                </MenuItem>
                            ))}
                        </Select>
                        {formik.touched.actions && formik.errors.actions ? (
                            <FormHelperText>{formik.errors.actions}</FormHelperText>
                        ) : null}
                    </FormControl>
                </Grid>
                {
                    formik.values.accounts?.length === 0 && <Grid item xs={12}>
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
                }
                {(formik.values.folderId === 0 || formik.values.folderId === '')&& <Grid item xs={12}>
                    <FormControl sx={{ width: '100%' }} error={Boolean(formik.errors.accounts)}>
                        <InputLabel id="demo-multiple-checkbox-label">账号</InputLabel>
                        <Select
                            labelId="accounts"
                            id="accounts"
                            multiple
                            value={formik.values.accounts || []}
                            onChange={(event) => {
                                const value = event.target.value;
                                formik.setFieldValue('accounts', value);
                            }}
                            input={<OutlinedInput label="账号" />}
                            renderValue={(selected) => selected
                                .map((value: any) => row?.accountsList?.find((item: any) => item.id === value)?.phone)
                                .join(', ')}
                            MenuProps={MenuProps}

                        >
                            {row?.accountsList?.map((item: any) => (
                                <MenuItem key={item.id} value={item.id}>
                                    <Checkbox checked={formik.values.accounts?.includes(item.id)} />
                                    <ListItemText primary={`${item.phone}-${item.firstName} ${item.lastName}`} />
                                </MenuItem>
                            ))}
                        </Select>
                        {formik.touched.accounts && formik.errors.accounts ? (
                            <FormHelperText>{formik.errors.accounts}</FormHelperText>
                        ) : null}
                    </FormControl>
                    <ButtonGroup size="small" aria-label="small button group" sx={{ display: 'flex', justifyContent: 'flex-end' }}>
                        <Button value="50" onClick={selectBtn}>选择50个</Button>
                        <Button value="100">选择100个</Button>
                        <Button value="200">选择200个</Button>
                        <Button value="clear" onClick={selectBtn}>清除</Button>
                    </ButtonGroup>
                </Grid>}
                <Grid item xs={12}>
                    <FormControl sx={{ width: '100%' }} error={Boolean(formik.errors.scriptGroup)}>
                        <InputLabel id="demo-multiple-checkbox-label">话术分组</InputLabel>
                        <Select
                            labelId="scriptGroup"
                            id="scriptGroup"
                            value={formik.values.scriptGroup || ''}
                            onChange={(event) => {
                                const value = event.target.value;
                                formik.setFieldValue('scriptGroup', value);
                            }}
                            input={<OutlinedInput label="话术分组" />}
                        >
                            {row?.scriptList?.map((item: any) => (
                                <MenuItem key={item.id} value={item.id}>
                                    <ListItemText primary={`${item.name}`} />
                                </MenuItem>
                            ))}
                        </Select>
                        {formik.touched.scriptGroup && formik.errors.scriptGroup ? (
                            <FormHelperText>{formik.errors.scriptGroup}</FormHelperText>
                        ) : null}
                    </FormControl>
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