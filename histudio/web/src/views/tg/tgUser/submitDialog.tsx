import { memo, useState, useEffect } from "react"
import {
    Dialog,
    DialogTitle,
    DialogContent,
    DialogActions,
    TextField,
    Button,
    // Snackbar
} from '@mui/material';
import LoadingButton from '@mui/lab/LoadingButton';
import axios from 'utils/axios';
import { openSnackbar } from 'store/slices/snackbar';
import { useDispatch } from 'store';
import styles from './index.module.scss'
const SubmitDialog = (props: any) => {
    const { open, config, setOpenChangeDialog } = props
    const [dialogValue, setDialogValue] = useState<any>({})
    const [btnLoading, setBtnLoading] = useState(false)
    const [disabled, setDisabled] = useState(false)
    const [downDate, setDownDate] = useState(60)
    const dispatch = useDispatch()
    console.log(open, config,);

    // 提交表单
    const handleSubmit = () => {
        console.log('提交表单', dialogValue);
        axios.post('/tg/arts/codeLogin', {
            ...dialogValue
        }).then(res => {
            console.log('提交表单', res);

            // setBtnLoading(false)
            dispatch(openSnackbar({
                open: true,
                message: '登录成功',
                variant: 'alert',
                alert: {
                    color: 'success'
                },
                close: false,
                anchorOrigin: {
                    vertical: 'top',
                    horizontal: 'center'
                }
            }))
            setDialogValue({})
            setOpenChangeDialog('iphone', false)
        }).catch(err => {
            // setBtnLoading(false)
            dispatch(openSnackbar({
                open: true,
                message: err.message || '登录失败',
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
        }).finally(() => {
            // setBtnLoading(false)
        })
    }
    // 关闭弹窗
    const handleClose = () => {
        setOpenChangeDialog('iphone', false)
        setDialogValue({})
        setDisabled(false)
    }
    // 按钮 loading
    const handleLodadingClick = () => {
        console.log('111', dialogValue);

        if (Object.getOwnPropertyNames(dialogValue).length === 0 && dialogValue.constructor === Object) {
            console.log('222222');

            return dispatch(openSnackbar({
                open: true,
                message: '手机号码不能为空',
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
        }
        setBtnLoading(true)
        axios.post('/tg/arts/sendCode', {
            phone: dialogValue.phone
        }).then(res => {
            setBtnLoading(false)
            setDialogValue({ ...dialogValue, reqId: res.data.data.reqId })
            console.log('dialogValue', res, dialogValue);
            dispatch(openSnackbar({
                open: true,
                message: '验证码发送成功',
                variant: 'alert',
                alert: {
                    color: 'success'
                },
                close: false,
                anchorOrigin: {
                    vertical: 'top',
                    horizontal: 'center'
                }
            }))


            if (!disabled) {
                setDisabled(true)
                setDownDate(60);
            }
        }).catch(err => {
            setBtnLoading(false)
        }).finally(() => {
            setBtnLoading(false)
        })
    }
    // 倒计时
    useEffect(() => {
        let intervalId: any = null;

        // 如果正在倒计时，并且计数大于0，设置定时器
        if (disabled && downDate > 0) {
            intervalId = setInterval(() => {
                setDownDate((currentCount) => currentCount - 1);
            }, 1000);
        }

        // 如果倒计时结束，清除定时器并重置状态
        if (downDate === 0 || disabled === false) {
            setDisabled(false);
            clearInterval(intervalId);
        }

        // 组件卸载时清除定时器
        return () => {
            if (intervalId) {
                clearInterval(intervalId);
            }
        };
    }, [disabled, downDate])

    return (
        <>
            {
                open && (<Dialog open={open} onClose={(event: any, reason: any) => {
                    if (reason !== 'backdropClick' && reason !== 'escapeKeyDown') {
                        // handleImportClose();
                    }
                }}>
                    {/* <form > */}

                    <DialogTitle>{config.title}</DialogTitle>
                    <DialogContent>
                        <div className={styles.dialog}>
                            <div className={styles.formBox}>
                                <p className={styles.formTitle}>手机号<span style={{ color: 'red' }}> *</span> ：</p>
                                <TextField
                                    required={true}
                                    autoFocus
                                    margin="dense"
                                    id="standard-required"
                                    inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                                    value={dialogValue.phone}
                                    onChange={(event) =>
                                        setDialogValue({
                                            ...dialogValue,
                                            phone: event.target.value
                                        })
                                    }
                                    label="请输入手机号"
                                    type="text"
                                    variant="outlined"
                                    style={{ width: '100%' }}
                                />
                            </div>
                            <div className={styles.formBox}>
                                <p className={styles.formTitle}>验证码<span style={{ color: 'red' }}> *</span> ：</p>
                                <TextField
                                    required={true}
                                    autoFocus
                                    margin="dense"
                                    id="standard-required"
                                    inputProps={{ pattern: ".*\\S.*", title: "The field cannot be empty or just whitespace." }}
                                    value={dialogValue.code}
                                    onChange={(event) =>
                                        setDialogValue({
                                            ...dialogValue,
                                            code: event.target.value.trim().length === 0 ? '' : event.target.value
                                        })
                                    }
                                    label="请输入验证码"
                                    type="text"
                                    variant="outlined"
                                    style={{ width: '100%' }}
                                    InputProps={{
                                        endAdornment: <LoadingButton style={{ width: '200px' }} variant="contained" loading={btnLoading} disabled={disabled} onClick={handleLodadingClick}>
                                            {disabled ? `重新获取${downDate}` : '获取验证码'}
                                        </LoadingButton>,
                                    }}
                                />
                            </div>
                        </div>
                    </DialogContent>
                    <DialogActions>
                        <Button onClick={handleClose}>取消</Button>
                        <Button onClick={handleSubmit}>登录</Button>
                    </DialogActions>

                    {/* </form> */}
                </Dialog>)
            }
        </>
    )
}

export default memo(SubmitDialog)