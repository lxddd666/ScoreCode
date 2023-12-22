
// material-ui
import { memo, useState } from "react"

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
} from '@mui/material';

import AutorenewIcon from '@mui/icons-material/Autorenew';
import SearchIcon from '@mui/icons-material/Search';

import { gridSpacing } from 'store/constant';
import Paper from '@mui/material/Paper';
import { styled } from '@mui/material/styles';
const Item = styled(Paper)(({ theme }) => ({
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));
const FormDialog = (props: any) => {
    const { open, config, onChangeDialogStatus } = props
    const [dialogValue] = useState<any>({})
    // 提交表单
    const handleSubmit = () => {
        console.log('提交表单', dialogValue);
        // onChangeDialogStatus('bind', false)
    }
    // 关闭弹窗
    const handleClose = () => {
        console.log('关闭弹窗');
        onChangeDialogStatus('bind', false)
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
                        <Edit />
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

// table 表格
export const columns = [
    {
        title: '所属用户',
        key: 'memberUsername'
    },
    {
        title: '用户名',
        key: 'username'
    },
    {
        title: '用户信息',
        key: 'firstName'
    },
    // {
    //     title: '姓氏',
    //     key: 'lastName'
    // },
    // {
    //     title: '手机号',
    //     key: 'phone'
    // },
    // {
    //     title: '账号头像',
    //     key: 'photo'
    // },
    {
        title: '账号状态',
        key: 'accountStatus'
    },
    {
        title: '是否在线',
        key: 'isOnline'
    },
    {
        title: '代理地址',
        key: 'proxyAddress'
    },
    {
        title: '上次登录时间',
        key: 'lastLoginTime'
    },
    {
        title: '备注',
        key: 'comment'
    },
    {
        title: '创建时间',
        key: 'createdAt'
    },
    {
        title: '更新时间',
        key: 'updatedAt'
    },
    {
        title: '操作',
        key: 'active'
    }
];
const Edit = (props: any) => {
    const [formData, setFormData] = useState<any>({
        folderId: undefined,
        username: undefined,
        firstName: undefined,
        lastName: undefined,
        phone: undefined,
        proxyAddress: undefined,
        accountStatus: undefined,
        isOnline: undefined,
        memberId: undefined
    })

    // 搜索按钮
    const onSearchClick = (e: any) => {
        // console.log(e.target.value, formData);
        // let obj = { folderId: value?.value ? value?.value : undefined, ...formData };
        // handleSearchFormData(obj);
    };
    // 重置按钮
    const onResetClick = (e: any) => {
        // let obj = {
        //     folderId: undefined,
        //     username: undefined,
        //     firstName: undefined,
        //     lastName: undefined,
        //     phone: undefined,
        //     proxyAddress: undefined,
        //     accountStatus: undefined,
        //     isOnline: undefined,
        //     memberId: undefined
        // };
        // setValue({});
        // setFormData(obj);
        // handleSearchFormData(obj);
    };
    return (
        <div style={{height:'600px'}}>
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
            <div style={{ width: '100%',maxHeight:'400px' }}>
                <TableContainer
                    component={Paper}
                >
                    <Table aria-label="simple table" sx={{ border: 1, borderColor: 'divider' }} stickyHeader={true}>
                        <TableHead>
                            <TableRow>
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
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>


                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>

                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>
                            <TableRow
                                hover
                                tabIndex={-1}
                            >
                                {columns.map((item) => {
                                    return (
                                        <TableCell align="center" key={item.key}>
                                            11111
                                        </TableCell>
                                    );
                                })}

                            </TableRow>

                        </TableBody>
                    </Table>
                </TableContainer>
            </div>
        </div>
    )
}

export default memo(FormDialog)