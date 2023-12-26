
// material-ui
import { memo, useEffect, useState, useRef } from "react"

import {
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions, Button, Grid, TextField, Stack,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Chip,
    Avatar,
    Radio,
    Pagination
} from '@mui/material';

import AutorenewIcon from '@mui/icons-material/Autorenew';
import SearchIcon from '@mui/icons-material/Search';

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
    tgUserBindUser,
    tgUserBindProxy,
    tgUserEdit
} from 'server/tg'

import { getProxyListAction } from "store/slices/org";

const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));
const FormDialog = (props: any) => {
    const { open, config, onChangeDialogStatus } = props
    const [dialogValue, setDialogValue] = useState<any>({})

    console.log('dormDialog', open, config);


    // 提交表单
    const handleSubmit = () => {
        console.log('提交表单', dialogValue, config.selectCheck);
        // onChangeDialogStatus('bind', false)
        if (!dialogValue) {
            return alert('提交失败，请选择用户')
        }
        if (config.type === 'bindUser') {
            let data = {
                memberId: dialogValue,
                ids: config.selectCheck
            }

            tgUserBindUser(data).then(res => {
                onChangeDialogStatus(config.type, false)
            }).catch(err => {
                console.log('失败', err);

            })
        }
        if (config.type === 'bindProxy') {
            let data = {
                proxyId: dialogValue,
                ids: config.selectCheck
            }

            tgUserBindProxy(data).then(res => {
                onChangeDialogStatus(config.type, false)
            }).catch(err => {
                console.log('失败', err);

            })
        }
        if (config.type === 'Edit') {
            let data = {
                ...dialogValue
            }

            tgUserEdit(data).then(res => {
                onChangeDialogStatus(config.type, false)
            }).catch(err => {
                console.log('失败', err);

            })
        }
    }
    // 关闭弹窗
    const handleClose = () => {
        // console.log('关闭弹窗');
        onChangeDialogStatus(config.type, false)
    }
    const changeSubmitValue = (value: any) => {
        setDialogValue(value)
    }
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
                                />
                                :
                                <EditTable changeSubmitValue={changeSubmitValue} columns={config.columns} type={config.type} />
                        }

                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose}>取消</Button>
                        <Button onClick={handleSubmit}>提交</Button>
                    </DialogActions>

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

const validationSchema = yup.object({
    username: yup.string().trim('Enter a valid username').required('username is required'),
    phone: yup.string().trim('Enter a valid phone').required('phone is required'),
});

const EditForm = (props: any) => {
    const { row, renderField, changeSubmitValue } = props
    // console.log('EditForm', row);

    useEffect(() => {
        console.log(row.id);

        if (row.id) {
            changeSubmitValue({
                id: row.id,
                ...formik.values
            })
        }
    }, [])

    const formik = useFormik({
        initialValues: {
            username: row.username || '',
            firstName: row.firstName || '',
            lastName: row.lastName || '',
            phone: row.phone || '',
            bio: row.bio || '',
            comment: row.comment || ''
        },
        validationSchema,
        onSubmit: (values) => {
            console.log('value', values,);
        }
    });

    const inputChange = (event: any) => {
        formik.handleChange(event)
        // console.log(event.target.value, formik.values);
        changeSubmitValue({
            id: row.id,
            ...formik.values
        })
    }


    return (
        <form onSubmit={formik.handleSubmit} style={{ marginTop: '10px' }}>

            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="username"
                        name="username"
                        label={renderField.username}
                        value={formik.values.username}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.username && Boolean(formik.errors.username)}
                        helperText={formik.touched.username && Boolean(formik.errors.username)}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="firstName"
                        name="firstName"
                        label={renderField.firstName}
                        value={formik.values.firstName}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="lastName"
                        name="lastName"
                        label={renderField.lastName}
                        value={formik.values.lastName}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="phone"
                        name="phone"
                        label={renderField.phone}
                        value={formik.values.phone}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.username && Boolean(formik.errors.username)}
                        helperText={formik.touched.username && Boolean(formik.errors.username)}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="bio"
                        name="bio"
                        label={renderField.bio}
                        value={formik.values.bio}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="comment"
                        name="comment"
                        label={renderField.comment}
                        value={formik.values.comment}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                    />
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