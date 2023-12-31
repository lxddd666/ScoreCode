
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
    // ButtonGroup,
    // Autocomplete
} from '@mui/material';
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
    tgFoldersEdit,
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
        folderName: yup.string().trim('Enter a valid folderName').required('folderName is required'),
        // id: yup.string().trim('Enter a valid id').required('id is required'),
        // accounts: yup.array().of(
        //     yup.number().required('Action is required')
        // ).test(
        //     'is-accounts-required',
        //     'At least one account is required',
        //     (value) => !requiredFields.accounts || (Array.isArray(value) && value?.length > 0)
        // ),
    });
}

const FormDialog = (props: any) => {
    const { open, config, onChangeDialogStatus } = props
    const [dialogValue, setDialogValue] = useState<any>({})
    const [formikValue, setFormikValue] = useState<any>({
        id: undefined,
        orgId: '',
        memberId: '',
        folderName: '',
        memberCount: '',
        accounts: [],
        comment: '',

    })
    const [requiredFields] = useState({
        accounts: false,
        folderName: false,
    })

    // console.log('dormDialog', open, config);
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
                    const res = await tgFoldersEdit(values)
                    console.log(res);
                    onChangeDialogStatus(config.type, false)
                } catch (error) {
                    console.log('error', error);

                }

            }
            if (config.type === 'Edit') {
                try {
                    const res = await tgFoldersEdit(values)
                    console.log(res);
                    onChangeDialogStatus(config.type, false)
                } catch (error) {
                    console.log('error', error);

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
            orgId: '',
            memberId: '',
            folderName: '',
            memberCount: '',
            accounts: [],
            comment: '',

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
                </Dialog>)
            }
        </>
    )
}


// table header options

const renderTable = (value: any, key: any, item: any) => {
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
                {/*<Grid item xs={3}>*/}
                {/*    <Item>*/}
                {/*        <Stack direction="row" spacing={2}>*/}
                {/*            <Button size="small" variant="outlined" startIcon={<SearchIcon />} onClick={onSearchClick}>*/}
                {/*                查询*/}
                {/*            </Button>*/}
                {/*            <Button size="small" variant="outlined" startIcon={<AutorenewIcon />} onClick={onResetClick}>*/}
                {/*                重置*/}
                {/*            </Button>*/}
                {/*        </Stack>*/}
                {/*    </Item>*/}
                {/*</Grid>*/}

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
    // console.log('EditForm', row);


    const inputChange = (event: any) => {
        formik.handleChange(event)
        changeSubmitValue({
            id: row?.id || undefined,
            ...formik.values
        })
    }

    return (
        <form onSubmit={formik.handleSubmit} style={{ marginTop: '10px' }}>

            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <TextField
                        type="number"
                        inputProps={{
                            min: 0, // 最小值
                            max: 10000, // 最大值
                            step: 1 // 步长
                        }}
                        fullWidth
                        id="orgId"
                        name="orgId"
                        label='组织ID'
                        value={formik.values.orgId}
                        onChange={(event) => {
                            const value = event.target.value;
                            formik.setFieldValue('orgId', value);
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
                        id="memberId"
                        name="memberId"
                        label='用户ID'
                        value={formik.values.memberId}
                        onChange={(event) => {
                            const value = event.target.value;
                            formik.setFieldValue('memberId', value);
                        }}
                        onBlur={formik.handleBlur}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="folderName"
                        name="folderName"
                        label='分组名称'
                        value={formik.values.folderName || ''}
                        onChange={(event) => {
                            const value = event.target.value;
                            formik.setFieldValue('folderName', value);
                        }}
                        onBlur={formik.handleBlur}
                        error={formik.touched.folderName && Boolean(formik.errors.folderName)}
                        helperText={formik.touched.folderName && formik.errors.folderName}
                    />
                </Grid>

                <Grid item xs={12}>
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
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="comment"
                        name="comment"
                        label='备注'
                        value={formik.values.comment}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.comment && Boolean(formik.errors.comment)}
                        helperText={formik.touched.comment && formik.errors.comment}
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
        </form>
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