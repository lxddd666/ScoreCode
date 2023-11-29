import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import { useDispatch } from 'store';
import * as CryptoJS from 'crypto-js';
import { debounce } from 'lodash';

// material-ui
import { useTheme, styled } from '@mui/material/styles';
import {
    Box,
    Button,
    Checkbox,
    FormControl,
    FormControlLabel,
    FormHelperText,
    Grid,
    IconButton,
    InputAdornment,
    InputLabel,
    OutlinedInput,
    Typography
} from '@mui/material';

// third party
import * as Yup from 'yup';
import { Formik } from 'formik';

// project imports
import AnimateButton from 'ui-component/extended/AnimateButton';
import useAuth from 'hooks/useAuth';
import useScriptRef from 'hooks/useScriptRef';
import { Spin } from 'antd';
import { LoadingOutlined } from '@ant-design/icons';

// assets
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import RefreshIcon from '@mui/icons-material/Refresh';

// utils
import axios from 'axios';

//env
import envRef from 'environment';

// popup
import { openSnackbar } from 'store/slices/snackbar';

// ===============================|| JWT LOGIN ||=============================== //

const GridWrapper = styled(Grid)(({ theme }) => ({
    '@media (min-width: 600px)': {
        paddingRight: '3px'
    }
}));

const CaptchaImgWrapper = styled(Grid)(({ theme }) => ({
    '@media (min-width: 600px)': {
        paddingLeft: '3px',
        textAlign: 'right'
    }
}));

const JWTLogin = ({ loginProp, ...others }: { loginProp?: number }) => {
    const theme = useTheme();
    const secondaryMain = theme.palette.secondary.main;

    const antIcon = (
        <LoadingOutlined
            style={{
                fontSize: 20,
                marginLeft: 4,
                color: secondaryMain
            }}
            spin
        />
    );

    const { loginUser } = useAuth();
    const scriptedRef = useScriptRef();

    const [checked, setChecked] = React.useState(true);
    const [showPassword, setShowPassword] = React.useState(false);
    const [captchaId, setCaptchaId] = React.useState('');
    const [captchaImg, setCaptchaImg] = React.useState('');
    const [isUserSubmitting, setIsUserSubmitting] = React.useState(false);

    const handleClickShowPassword = () => {
        setShowPassword(!showPassword);
    };

    const handleMouseDownPassword = (event: React.MouseEvent) => {
        event.preventDefault()!;
    };

    const dispatch = useDispatch();

    const getCaptcha = () => {
        setIsUserSubmitting(true);
        getCaptchaFunc();
    };

    const getCaptchaFunc = debounce(async () => {
        await axios
            .get(`${envRef?.API_URL}admin/site/captcha`)
            .then(function (response) {
                setIsUserSubmitting(false);
                if (response?.data?.data) {
                    setCaptchaId(response?.data?.data?.cid);
                    setCaptchaImg(response?.data?.data?.base64);
                } else {
                    dispatch(
                        openSnackbar({
                            open: true,
                            message: '' + response?.data?.message || 'Something wrong. Please try again later.',
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
            })
            .catch(function (error) {
                setIsUserSubmitting(false);
                dispatch(
                    openSnackbar({
                        open: true,
                        message: '' + error?.message || 'Something wrong. Please try again later.',
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
    }, 3000);

    useEffect(() => {
        (async function () {
            {
                await axios
                    .get(`${envRef?.API_URL}admin/site/captcha`)
                    .then(function (response) {
                        if (response?.data?.data) {
                            setCaptchaId(response?.data?.data?.cid);
                            setCaptchaImg(response?.data?.data?.base64);
                        } else {
                            dispatch(
                                openSnackbar({
                                    open: true,
                                    message: '' + response?.data?.message || 'Something wrong. Please try again later.',
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
                    })
                    .catch(function (error) {
                        dispatch(
                            openSnackbar({
                                open: true,
                                message: '' + error?.message || 'Something wrong. Please try again later.',
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
            }
        })();
    }, []);

    return (
        <>
            <Formik
                initialValues={{
                    username: `${envRef?.USERNAME}`,
                    password: `${envRef?.PASSWORD}`,
                    captcha: '',
                    submit: null
                }}
                validationSchema={Yup.object().shape({
                    username: Yup.string().required('Email is required'),
                    password: Yup.string().max(255).required('Password is required'),
                    captcha: Yup.string().required('Captcha is required')
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

                        const response = await axios.post(
                            `${envRef?.API_URL}admin/site/accountLogin`,
                            { username: values.username, password: encryptedPassword, code: values.captcha, cid: captchaId },
                            {
                                headers: {
                                    'Content-Type': 'application/json'
                                },
                                withCredentials: true
                            }
                        );

                        if (response?.data?.code == 0) {
                            loginUser(response);
                            setStatus({ success: true });
                            setSubmitting(false);
                        } else {
                            setStatus({ success: false });
                            setErrors({ submit: '' + response?.data?.message });
                            getCaptcha();
                            setSubmitting(false);
                        }
                    } catch (err: any) {
                        if (scriptedRef.current) {
                            setStatus({ success: false });
                            setErrors({ submit: err.message || '' + err });
                            getCaptcha();
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
                            <InputLabel htmlFor="outlined-adornment-username-login">Username</InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-username-login"
                                type="text"
                                value={values.username}
                                name="username"
                                onBlur={handleBlur}
                                onChange={handleChange}
                                inputProps={{}}
                            />
                            {touched.username && errors.username && (
                                <FormHelperText error id="standard-weight-helper-text-username-login">
                                    {errors.username}
                                </FormHelperText>
                            )}
                        </FormControl>

                        <FormControl
                            fullWidth
                            error={Boolean(touched.password && errors.password)}
                            sx={{ ...theme.typography.customInput }}
                        >
                            <InputLabel htmlFor="outlined-adornment-password-login">Password</InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-password-login"
                                type={showPassword ? 'text' : 'password'}
                                value={values.password}
                                name="password"
                                onBlur={handleBlur}
                                onChange={handleChange}
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
                                label="Password"
                            />
                            {touched.password && errors.password && (
                                <FormHelperText error id="standard-weight-helper-text-password-login">
                                    {errors.password}
                                </FormHelperText>
                            )}
                        </FormControl>

                        <Grid container alignItems="center">
                            <GridWrapper style={{ flexWrap: 'nowrap' }} item xs={12} sm={8}>
                                <FormControl
                                    fullWidth
                                    error={Boolean(touched.captcha && errors.captcha)}
                                    sx={{ ...theme.typography.customInput }}
                                >
                                    <InputLabel htmlFor="outlined-adornment-captcha-login">Captcha</InputLabel>
                                    <OutlinedInput
                                        id="outlined-adornment-captcha-login"
                                        type="text"
                                        value={values.captcha}
                                        name="captcha"
                                        onBlur={handleBlur}
                                        onChange={handleChange}
                                        inputProps={{}}
                                    />
                                    {touched.captcha && errors.captcha && (
                                        <FormHelperText error id="standard-weight-helper-text-captcha-login">
                                            {errors.captcha}
                                        </FormHelperText>
                                    )}
                                </FormControl>
                            </GridWrapper>
                            <CaptchaImgWrapper
                                style={{ flexWrap: 'nowrap' }}
                                item
                                xs={12}
                                sm={4}
                                container
                                direction="row"
                                alignItems="center"
                            >
                                {captchaImg ? (
                                    <>
                                        <img src={captchaImg} style={{ flex: 1, maxWidth: 130 }} height={'auto'} />
                                        {isUserSubmitting ? (
                                            <Spin indicator={antIcon} />
                                        ) : (
                                            <RefreshIcon
                                                color="secondary"
                                                onClick={getCaptcha}
                                                style={{ cursor: 'pointer', width: '24px' }}
                                            />
                                        )}
                                    </>
                                ) : (
                                    <></>
                                )}
                            </CaptchaImgWrapper>
                        </Grid>

                        <Grid container alignItems="center" justifyContent="space-between">
                            <Grid item>
                                <FormControlLabel
                                    control={
                                        <Checkbox
                                            checked={checked}
                                            onChange={(event) => setChecked(event.target.checked)}
                                            name="checked"
                                            color="primary"
                                        />
                                    }
                                    label="Keep me logged in"
                                />
                            </Grid>
                            <Grid item>
                                <Typography
                                    variant="subtitle1"
                                    component={Link}
                                    to={
                                        loginProp
                                            ? `/pages/forgot-password/forgot-password${loginProp}`
                                            : '/pages/forgot-password/forgot-password3'
                                    }
                                    color="secondary"
                                    sx={{ textDecoration: 'none' }}
                                >
                                    Forgot Password?
                                </Typography>
                            </Grid>
                        </Grid>

                        {errors.submit && (
                            <Box sx={{ mt: 3 }}>
                                <FormHelperText error>{errors.submit}</FormHelperText>
                            </Box>
                        )}
                        <Box sx={{ mt: 2 }}>
                            <AnimateButton>
                                <Button color="secondary" disabled={isSubmitting} fullWidth size="large" type="submit" variant="contained">
                                    Sign In
                                </Button>
                            </AnimateButton>
                        </Box>
                    </form>
                )}
            </Formik>
        </>
    );
};

export default JWTLogin;
