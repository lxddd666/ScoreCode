
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
    ListItemText,
    FormHelperText,
} from '@mui/material';
// import ClearIcon from '@mui/icons-material/Clear';
import AutorenewIcon from '@mui/icons-material/Autorenew';
import SearchIcon from '@mui/icons-material/Search';
import AnimateButton from 'ui-component/extended/AnimateButton';

import { gridSpacing } from 'store/constant';
import Paper from '@mui/material/Paper';
import { styled } from '@mui/material/styles';
import Badge from '@mui/material/Badge';
import styles from './index.module.scss';

import { useFormik } from 'formik';
import * as yup from 'yup';

import { timeAgo } from 'utils/tools'
import {
    tgBatchExecutionTaskEdit
} from 'server/tg'


const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));
const createValidationSchema = (requiredFields: any) => {
    return yup.object().shape({
        orgId: yup.string().trim('Enter a valid orgId'),

    });
}

const FormDialog = (props: any) => {
    const { open, config, onChangeDialogStatus } = props
    const [dialogValue, setDialogValue] = useState<any>({})
    const [formikValue, setFormikValue] = useState<any>({
        id: undefined,
        orgId:undefined,
        taskName: '',
        cron: '',
        action: [],
        status: '',
        comment: ''
    })
    const [requiredFields] = useState({
        id: false,
        orgId:false
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
                    const res = await tgBatchExecutionTaskEdit(values)
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
            orgId:undefined,
            taskName: '',
            cron: '',
            action: [],
            status: '',
            comment: ''
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
    const { changeSubmitValue, columns } = props
    const [formData, setFormData] = useState<any>({
        username: undefined,
    })
    // 用于跟踪选中行的ID
    const [selectedId, setSelectedId] = useState(null);
    const [userListRow,] = useState<any>([])
    const [pagetionTotle] = useState(0); // total


    const handleSelectRow = (id: any) => {
        // console.log(id);
        setSelectedId(id);
        changeSubmitValue(id)
    };

    // 搜索按钮
    const onSearchClick = (e: any) => {
        console.log(e.target.value, formData);


    };
    // 重置按钮
    const onResetClick = (e: any) => {
        let obj = {
            username: undefined,
        };
        // setValue({});
        setFormData(obj);
    };
    // 分页事件
    const pageRef = useRef(1);
    const onPaginationChange = (event: object, page: number) => {
        pageRef.current = page;
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
    const { row, changeSubmitValue, formik, handleClose } = props
    // console.log('EditForm', formik.initialValues);


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
                        fullWidth
                        id="orgId"
                        name="orgId"
                        label='组织ID'
                        value={formik.values.orgId}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.orgId && Boolean(formik.errors.orgId)}
                        helperText={formik.touched.orgId && formik.errors.orgId}
                    />
                </Grid>
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
                        id="action"
                        name="action"
                        label='操作动作'
                        value={formik.values.action}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.action && Boolean(formik.errors.action)}
                        helperText={formik.touched.action && formik.errors.action}
                    />
                </Grid>
                <Grid item xs={12}>
                    <FormControl sx={{ width: '100%' }} error={Boolean(formik.errors.taskStatus)}>
                        <InputLabel id="demo-multiple-checkbox-label">任务状态</InputLabel>
                        <Select
                            labelId="status"
                            id="status"
                            value={formik.values.status || ''}
                            onChange={(event) => {
                                const value = event.target.value;
                                formik.setFieldValue('status', value);
                            }}
                            input={<OutlinedInput label="任务状态" />}
                        >
                            {row?.statusList?.map((item: any) => (
                                <MenuItem key={item.value} value={item.value}>
                                    <ListItemText primary={item.title} />
                                </MenuItem>
                            ))}
                            <MenuItem key="normal" value="normal">
                                <ListItemText primary="正常" />
                            </MenuItem>
                            <MenuItem key="inactive" value="inactive">
                                <ListItemText primary="停用" />
                            </MenuItem>
                        </Select>
                        {formik.touched.status && formik.errors.status ? (
                            <FormHelperText>{formik.errors.status}</FormHelperText>
                        ) : null}
                    </FormControl>
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        fullWidth
                        id="id"
                        name="id"
                        label='账号id'
                        value={formik.values.action}
                        onChange={inputChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.action && Boolean(formik.errors.action)}
                        helperText={formik.touched.action && formik.errors.action}
                    />
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
                <Grid container spacing={gridSpacing}>
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