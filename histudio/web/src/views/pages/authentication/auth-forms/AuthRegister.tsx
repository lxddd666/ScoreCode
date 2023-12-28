import React, { useEffect, useState, useRef } from 'react';
import { useDispatch } from 'store';
import { Link, useNavigate, useParams } from 'react-router-dom';
import * as CryptoJS from 'crypto-js';

// material-ui
import { useTheme } from '@mui/material/styles';
import {
    Box,
    Button,
    Checkbox,
    Dialog,
    FormControl,
    FormControlLabel,
    FormHelperText,
    Grid,
    IconButton,
    InputAdornment,
    InputLabel,
    OutlinedInput,
    Typography,
    useMediaQuery
} from '@mui/material';
import CancelTwoToneIcon from '@mui/icons-material/CancelTwoTone';

// third party
import * as Yup from 'yup';
import { Formik } from 'formik';

// project imports
import AnimateButton from 'ui-component/extended/AnimateButton';
import useScriptRef from 'hooks/useScriptRef';
import { strengthColor, strengthIndicator } from 'utils/password-strength';
import { openSnackbar } from 'store/slices/snackbar';

// assets
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import { StringColorProps } from 'types';
// utils
import axios from 'axios';

// env
import envRef from 'environment';

// locale
import { useIntl } from 'react-intl';

// ===========================|| FIREBASE - REGISTER ||=========================== //

const JWTRegister = ({ ...others }) => {
    const intl = useIntl();
    const theme = useTheme();
    const navigate = useNavigate();
    const scriptedRef = useScriptRef();
    const dispatch = useDispatch();
    const matchDownSM = useMediaQuery(theme.breakpoints.down('md'));
    const { inviteCode } = useParams();

    const defaultErrorMessage = intl.formatMessage({ id: 'auth-register.default-error' });
    // TODO - change url
    const tncFileUrl = '../../../../assets/documents/test.pdf';

    const [showPassword, setShowPassword] = useState(false);
    const [checked, setChecked] = useState(true);

    const [strength, setStrength] = useState(1);
    const [level, setLevel] = useState<StringColorProps>();

    const handleClickShowPassword = () => {
        setShowPassword(!showPassword);
    };

    const handleMouseDownPassword = (event: React.SyntheticEvent) => {
        event.preventDefault();
    };

    // START Modal handler
    const [isTncModalOpen, setIsTncModalOpen] = useState(false);

    const handleTncModal = () => {
        setIsTncModalOpen(!isTncModalOpen);
    };
    // END Modal handler

    const changePassword = (value: string) => {
        const temp = strengthIndicator(value);
        setStrength(temp);
        setLevel(strengthColor(temp));
    };

    // START OTP Timer Countdown
    const defaultOtpTimerSeconds = 60;
    const [otpTimer, setOtpTimer] = useState(0);
    const otpTimerRef = useRef(() => { });

    const sendOtpCode = (event: React.SyntheticEvent, mobile: string, email: string) => {
        event.preventDefault();
        try {
            axios
                .post(`${envRef?.API_URL}/admin/site/register/sendCode`, { mobile: mobile, email: email }, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        dispatch(
                            openSnackbar({
                                open: true,
                                message: response?.data?.message,
                                variant: 'alert',
                                alert: {
                                    color: 'success'
                                },
                                close: false
                            })
                        );
                        setOtpTimer(defaultOtpTimerSeconds);
                    } else {
                        dispatch(
                            openSnackbar({
                                open: true,
                                message: response?.data?.message || defaultErrorMessage,
                                variant: 'alert',
                                alert: {
                                    color: 'error',
                                    severity: 'error'
                                },
                                close: false,
                                anchorOrigin: {
                                    vertical: 'top',
                                    horizontal: 'center'
                                }
                            })
                        );
                    }
                });
        } catch (error: any) {
            dispatch(
                openSnackbar({
                    open: true,
                    message: error?.message || defaultErrorMessage,
                    variant: 'alert',
                    alert: {
                        color: 'error'
                    },
                    close: false,
                    anchorOrigin: {
                        vertical: 'top',
                        horizontal: 'center'
                    }
                })
            );
        }
    };

    useEffect(() => {
        otpTimerRef.current = () => {
            if (otpTimer != 0) {
                setOtpTimer(otpTimer - 1);
            }
        };
    });

    useEffect(() => {
        const timer = setInterval(() => {
            otpTimerRef.current();
        }, 1000);

        return () => {
            clearInterval(timer);
        };
    }, []);
    // END OTP Timer Countdown

    // Mobile Number Regular Expression
    const mobileNumberRegExp = /^(\+?\d{0,4})?\s?-?\s?(\(?\d{3}\)?)\s?-?\s?(\(?\d{3}\)?)\s?-?\s?(\(?\d{4}\)?)?$/;

    return (
        <>
            <Formik
                initialValues={{
                    username: '',
                    firstName: '',
                    lastName: '',
                    mobile: '',
                    email: '',
                    password: '',
                    code: '',
                    inviteCode,
                    tnc: '',
                    orgInfo: {
                        name: '',
                        code: '',
                        phone: '',
                        leader: '',
                        email: '',
                    },
                    submit: null
                }}
                validationSchema={Yup.object().shape({
                    username: Yup.string()
                        .max(255)
                        .required(intl.formatMessage({ id: 'auth-register.username-required' })),
                    firstName: Yup.string()
                        .max(255)
                        .required(intl.formatMessage({ id: 'auth-register.username-required' })),
                    lastName: Yup.string()
                        .max(255)
                        .required(intl.formatMessage({ id: 'auth-register.username-required' })),
                    code: Yup.string().required(intl.formatMessage({ id: 'auth-register.otp-required' })),
                    email: Yup.string()
                        .email(intl.formatMessage({ id: 'auth-register.email-invalid' }))
                        .max(255)
                        .required(intl.formatMessage({ id: 'auth-register.email-required' })),
                    mobile: Yup.string()
                        .matches(mobileNumberRegExp, intl.formatMessage({ id: 'auth-register.mobile-invalid' }))
                        .max(255)
                        .required(intl.formatMessage({ id: 'auth-register.mobile-required' })),
                    password: Yup.string()
                        .max(255)
                        .required(intl.formatMessage({ id: 'auth-register.password-required' }))
                })}
                onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                    try {
                        const text = CryptoJS.enc.Utf8.parse(values.password);
                        const key = CryptoJS.enc.Utf8.parse(envRef?.AES_KEY);

                        const encryptedPassword = CryptoJS.AES.encrypt(text, key, {
                            //iv: ivSpec, //yes I used password as iv too. Dont mind.
                            mode: CryptoJS.mode.ECB,
                            padding: CryptoJS.pad.Pkcs7,
                            keySize: 128 / 8
                        }).toString();

                        await axios
                            .post(
                                `${envRef?.API_URL}/admin/site/register`,
                                {
                                    username: values.username,
                                    firstName: values.firstName,
                                    lastName: values.lastName,
                                    password: encryptedPassword,
                                    mobile: values.mobile,
                                    code: values.code,
                                    inviteCode: values.inviteCode,
                                    email: values.email,
                                    orgInfo: values.orgInfo
                                },
                                { headers: {} }
                            )
                            .then(function (response) {
                                if (response?.data?.code == 0) {
                                    setStatus({ success: true });
                                    setSubmitting(false);
                                    dispatch(
                                        openSnackbar({
                                            open: true,
                                            message: response?.data?.message,
                                            variant: 'alert',
                                            alert: {
                                                color: 'success'
                                            },
                                            close: false
                                        })
                                    );

                                    setTimeout(() => {
                                        navigate('/login', { replace: true });
                                    }, 1500);
                                } else {
                                    dispatch(
                                        openSnackbar({
                                            open: true,
                                            message: response?.data?.message || defaultErrorMessage,
                                            variant: 'alert',
                                            alert: {
                                                color: 'error',
                                                severity: 'error'
                                            },
                                            close: false,
                                            anchorOrigin: {
                                                vertical: 'top',
                                                horizontal: 'center'
                                            }
                                        })
                                    );
                                }
                            })
                            .catch(function (error) {
                                dispatch(
                                    openSnackbar({
                                        open: true,
                                        message: error?.message || defaultErrorMessage,
                                        variant: 'alert',
                                        alert: {
                                            color: 'error'
                                        },
                                        close: false,
                                        anchorOrigin: {
                                            vertical: 'top',
                                            horizontal: 'center'
                                        }
                                    })
                                );
                            });
                    } catch (err: any) {
                        if (scriptedRef.current) {
                            setStatus({ success: false });
                            setErrors({ submit: err.message });
                            setSubmitting(false);
                        }
                    }
                }}
            >
                {({ errors, handleBlur, handleChange, handleSubmit, isSubmitting, touched, values }) => (
                    <form noValidate onSubmit={handleSubmit} {...others}>
                        <FormControl
                            fullWidth
                            error={Boolean(touched.username && errors.username)}
                            sx={{ ...theme.typography.customInput }}
                        >
                            <InputLabel htmlFor="outlined-adornment-username-register">
                                {intl.formatMessage({ id: 'auth-register.username' })}
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-username-register"
                                type="text"
                                value={values.username}
                                name="username"
                                onBlur={handleBlur}
                                onChange={handleChange}
                                inputProps={{}}
                            />
                            {touched.username && errors.username && (
                                <FormHelperText error id="standard-weight-helper-text-username-register">
                                    {errors.username}
                                </FormHelperText>
                            )}
                        </FormControl>

                        <FormControl
                            fullWidth
                            error={Boolean(touched.firstName && errors.firstName)}
                            sx={{ ...theme.typography.customInput }}
                        >
                            <InputLabel htmlFor="outlined-adornment-username-register">
                                请输入姓氏
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-username-register"
                                type="text"
                                value={values.firstName}
                                name="firstName"
                                onBlur={handleBlur}
                                onChange={handleChange}
                                inputProps={{}}
                            />
                            {touched.firstName && errors.firstName && (
                                <FormHelperText error id="standard-weight-helper-text-username-register">
                                    {errors.firstName}
                                </FormHelperText>
                            )}
                        </FormControl>

                        <FormControl
                            fullWidth
                            error={Boolean(touched.lastName && errors.lastName)}
                            sx={{ ...theme.typography.customInput }}
                        >
                            <InputLabel htmlFor="outlined-adornment-username-register">
                                请输入名字
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-username-register"
                                type="text"
                                value={values.lastName}
                                name="lastName"
                                onBlur={handleBlur}
                                onChange={handleChange}
                                inputProps={{}}
                            />
                            {touched.lastName && errors.lastName && (
                                <FormHelperText error id="standard-weight-helper-text-username-register">
                                    {errors.lastName}
                                </FormHelperText>
                            )}
                        </FormControl>

                        <FormControl fullWidth error={Boolean(touched.email && errors.email)} sx={{ ...theme.typography.customInput }}>
                            <InputLabel htmlFor="outlined-adornment-email-register">
                                {intl.formatMessage({ id: 'auth-register.email-address' })}
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-email-register"
                                type="email"
                                value={values.email}
                                name="email"
                                onBlur={handleBlur}
                                onChange={handleChange}
                                inputProps={{}}
                            />
                            {touched.email && errors.email && (
                                <FormHelperText error id="standard-weight-helper-text-email-register">
                                    {errors.email}
                                </FormHelperText>
                            )}
                        </FormControl>

                        <FormControl fullWidth error={Boolean(touched.mobile && errors.mobile)} sx={{ ...theme.typography.customInput }}>
                            <InputLabel htmlFor="outlined-adornment-mobile-register">
                                {intl.formatMessage({ id: 'auth-register.mobile-number' })}
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-mobile-register"
                                type="mobile"
                                value={values.mobile}
                                name="mobile"
                                onBlur={handleBlur}
                                onChange={handleChange}
                                inputProps={{}}
                            />
                            {touched.mobile && errors.mobile && (
                                <FormHelperText error id="standard-weight-helper-text-mobile-register">
                                    {errors.mobile}
                                </FormHelperText>
                            )}
                        </FormControl>

                        <Grid container spacing={matchDownSM ? 1 : 2} alignItems="center">
                            <Grid item xs={12} sm={8}>
                                <FormControl
                                    fullWidth
                                    error={Boolean(touched.code && errors.code)}
                                    sx={{ ...theme.typography.customInput }}
                                >
                                    <InputLabel htmlFor="outlined-adornment-code-register">
                                        {intl.formatMessage({ id: 'auth-register.otp-code' })}
                                    </InputLabel>
                                    <OutlinedInput
                                        id="outlined-adornment-code-register"
                                        type="code"
                                        value={values.code}
                                        name="code"
                                        onBlur={handleBlur}
                                        onChange={handleChange}
                                        inputProps={{}}
                                    />
                                </FormControl>
                            </Grid>
                            <Grid item xs={12} sm={4}>
                                <Button
                                    id="send-otp-button"
                                    disabled={otpTimer == 0 ? false : true}
                                    size="large"
                                    type="button"
                                    variant="outlined"
                                    color="inherit"
                                    fullWidth
                                    sx={{ textWrap: 'nowrap', color: theme.palette.secondary.main }}
                                    onClick={(e) => {
                                        sendOtpCode(e, values.mobile, values.email);
                                    }}
                                >
                                    <span>{otpTimer == 0 ? intl.formatMessage({ id: 'auth-register.send-otp' }) : otpTimer}</span>
                                </Button>
                            </Grid>
                        </Grid>

                        <FormControl
                            fullWidth
                            sx={{
                                display: errors.code ? 'block' : 'none'
                            }}
                        >
                            {touched.code && errors.code && (
                                <FormHelperText error id="standard-weight-helper-text-code-register">
                                    {errors.code}
                                </FormHelperText>
                            )}
                        </FormControl>

                        <FormControl
                            fullWidth
                            error={Boolean(touched.password && errors.password)}
                            sx={{ ...theme.typography.customInput }}
                        >
                            <InputLabel htmlFor="outlined-adornment-password-register">
                                {intl.formatMessage({ id: 'auth-register.password' })}
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-password-register"
                                type={showPassword ? 'text' : 'password'}
                                value={values.password}
                                name="password"
                                label="Password"
                                onBlur={handleBlur}
                                onChange={(e) => {
                                    handleChange(e);
                                    changePassword(e.target.value);
                                }}
                                endAdornment={
                                    <InputAdornment position="end">
                                        <IconButton
                                            aria-label="toggle password visibility"
                                            onClick={handleClickShowPassword}
                                            onMouseDown={handleMouseDownPassword}
                                            edge="end"
                                            size="large"
                                        >
                                            {showPassword ? <Visibility /> : <VisibilityOff />}
                                        </IconButton>
                                    </InputAdornment>
                                }
                                inputProps={{}}
                            />
                            {touched.password && errors.password && (
                                <FormHelperText error id="standard-weight-helper-text-password-register">
                                    {errors.password}
                                </FormHelperText>
                            )}
                        </FormControl>

                        {strength !== 0 && (
                            <FormControl fullWidth>
                                <Box sx={{ mb: 2 }}>
                                    <Grid container spacing={2} alignItems="center">
                                        <Grid item>
                                            <Box
                                                style={{ backgroundColor: level?.color }}
                                                sx={{ width: 85, height: 8, borderRadius: '7px' }}
                                            />
                                        </Grid>
                                        <Grid item>
                                            <Typography variant="subtitle1" fontSize="0.75rem">
                                                {level?.label}
                                            </Typography>
                                        </Grid>
                                    </Grid>
                                </Box>
                            </FormControl>
                        )}

                        <FormControl
                            fullWidth
                            error={Boolean(touched.inviteCode && errors.inviteCode)}
                            sx={{ ...theme.typography.customInput }}
                        >
                            <InputLabel htmlFor="outlined-adornment-inviteCode-register">
                                {intl.formatMessage({ id: 'auth-register.invite-code' })}
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-inviteCode-register"
                                type="text"
                                value={values.inviteCode}
                                name="inviteCode"
                                onBlur={handleBlur}
                                onChange={handleChange}
                                inputProps={{}}
                            // disabled={values.inviteCode ? true : false}
                            />
                            {touched.inviteCode && errors.inviteCode && (
                                <FormHelperText error id="standard-weight-helper-text-inviteCode-register">
                                    {errors.inviteCode}
                                </FormHelperText>
                            )}
                        </FormControl>




                        {
                            !values.inviteCode ? <>
                                <FormControl
                                    fullWidth
                                    sx={{ ...theme.typography.customInput }}
                                >
                                    <InputLabel htmlFor="outlined-adornment-username-register">
                                        请输入公司名称
                                    </InputLabel>
                                    <OutlinedInput
                                        id="outlined-adornment-username-register"
                                        type="text"
                                        value={values?.orgInfo?.name || ''}
                                        name="orgInfo.name"
                                        onBlur={handleBlur}
                                        onChange={handleChange}
                                        inputProps={{}}
                                    />
                                </FormControl>
                                <FormControl
                                    fullWidth
                                    sx={{ ...theme.typography.customInput }}
                                >
                                    <InputLabel htmlFor="outlined-adornment-username-register">
                                        请输入公司编码
                                    </InputLabel>
                                    <OutlinedInput
                                        id="outlined-adornment-username-register"
                                        type="text"
                                        value={values?.orgInfo?.code || ''}
                                        name="orgInfo.code"
                                        onBlur={handleBlur}
                                        onChange={handleChange}
                                        inputProps={{}}
                                    />
                                </FormControl>
                                <FormControl
                                    fullWidth
                                    sx={{ ...theme.typography.customInput }}
                                >
                                    <InputLabel htmlFor="outlined-adornment-username-register">
                                        请输入公司联系电话
                                    </InputLabel>
                                    <OutlinedInput
                                        id="outlined-adornment-username-register"
                                        type="text"
                                        value={values?.orgInfo?.phone || ''}
                                        name="orgInfo.phone"
                                        onBlur={handleBlur}
                                        onChange={handleChange}
                                        inputProps={{}}
                                    />
                                </FormControl>
                                <FormControl
                                    fullWidth
                                    sx={{ ...theme.typography.customInput }}
                                >
                                    <InputLabel htmlFor="outlined-adornment-username-register">
                                        请输入公司负责人
                                    </InputLabel>
                                    <OutlinedInput
                                        id="outlined-adornment-username-register"
                                        type="text"
                                        value={values?.orgInfo?.leader || ''}
                                        name="orgInfo.leader"
                                        onBlur={handleBlur}
                                        onChange={handleChange}
                                        inputProps={{}}
                                    />
                                </FormControl>
                                <FormControl
                                    fullWidth
                                    sx={{ ...theme.typography.customInput }}
                                >
                                    <InputLabel htmlFor="outlined-adornment-username-register">
                                        请输入公司邮箱
                                    </InputLabel>
                                    <OutlinedInput
                                        id="outlined-adornment-username-register"
                                        type="text"
                                        value={values?.orgInfo?.email || ''}
                                        name="orgInfo.email"
                                        onBlur={handleBlur}
                                        onChange={handleChange}
                                        inputProps={{}}
                                    />
                                </FormControl>
                            </> :
                                ''
                        }

                        <Grid container alignItems="center" justifyContent="space-between">
                            <Grid item>
                                <FormControlLabel
                                    control={
                                        <Checkbox
                                            checked={checked}
                                            onChange={(event) => {
                                                setChecked(event.target.checked);
                                                errors.tnc = event.target.checked ? '' : intl.formatMessage({ id: 'auth-register.tnc' });
                                            }}
                                            name="checked"
                                            sx={{ color: theme.palette.secondary.main }}
                                        />
                                    }
                                    label={
                                        <Typography variant="subtitle1" sx={{ textWrap: 'nowrap' }}>
                                            {intl.formatMessage({ id: 'auth-register.agree-with' })} &nbsp;
                                            <Typography variant="subtitle1" component={Link} to="#" onClick={handleTncModal}>
                                                {intl.formatMessage({ id: 'auth-register.terms-and-conditions' })}
                                            </Typography>
                                        </Typography>
                                    }
                                />
                                {errors.tnc && (
                                    <FormHelperText error id="standard-weight-helper-text-tnc-register">
                                        {errors.tnc}
                                    </FormHelperText>
                                )}
                            </Grid>
                        </Grid>

                        {errors.submit && (
                            <Box sx={{ mt: 3 }}>
                                <FormHelperText error>{errors.submit}</FormHelperText>
                            </Box>
                        )}

                        <Box sx={{ mt: 2 }}>
                            <AnimateButton>
                                <Button
                                    disableElevation
                                    disabled={isSubmitting}
                                    fullWidth
                                    size="large"
                                    type="submit"
                                    variant="contained"
                                    color="secondary"
                                >
                                    11{intl.formatMessage({ id: 'auth-register.sign-up' })}
                                </Button>
                            </AnimateButton>
                        </Box>
                    </form>
                )}
            </Formik>
            <Dialog
                maxWidth="md"
                fullWidth
                onClose={handleTncModal}
                open={isTncModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                <IconButton onClick={handleTncModal} sx={{ alignSelf: 'end' }}>
                    <CancelTwoToneIcon />
                </IconButton>
                {isTncModalOpen && <iframe src={tncFileUrl} style={{ width: '100%', minHeight: '500px' }} />}
            </Dialog>
        </>
    );
};

export default JWTRegister;
