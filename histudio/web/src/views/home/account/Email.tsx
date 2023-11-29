import { useEffect } from 'react';
import useState from 'react-usestateref';
// material-ui
import { Grid } from '@mui/material';

// project imports
import SubCard from 'ui-component/cards/SubCard';
import { gridSpacing } from 'store/constant';
import GeneralForm from '../../../ui-component/form';
import { ParamForm } from 'utils/generalInterface';
import Request from 'utils/request';
import { setAdminInfo, setAdminInfoError } from 'store/slices/user';
import Loading from 'ui-component/Loading';

// thrid party
import { useFormik, FormikProvider } from 'formik';
import * as yup from 'yup';
import { useDispatch, useSelector } from 'store';
import { openSnackbar } from 'store/slices/snackbar';
import { useIntl } from 'react-intl';

// ==============================|| PROFILE 1 - CHANGE PASSWORD ||============================== //

interface jsonFormat {
    value?: string;
    label?: string;
    optionButtonLabel?: string;
}

interface jsonItem {
    email: jsonFormat;
    mailLabel: jsonFormat;
    code: jsonFormat;
}

const Email = () => {
    const intl = useIntl();
    const [isLoading, setIsLoading] = useState(false);
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isRequest, setIsRequest] = useState<boolean>(false);
    const [timeRemaining, setTimeRemaining, timeRemainingRef] = useState(0); // initial time in seconds
    const userState = useSelector((state) => state.user);

    const [validation, setValidation] = useState(
        yup.object({
            email: yup
                .string()
                .required(`${intl.formatMessage({ id: 'security.email' })} ${intl.formatMessage({ id: 'error.required' })}`)
                .email()
        })
    );

    const mail = intl.formatMessage({ id: 'security.email' });

    const dispatch = useDispatch();

    const [inputFields, setInputFields] = useState<ParamForm[]>([
        { type: 'text', name: 'email', label: intl.formatMessage({ id: 'security.email' }) }
    ]);

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            email: ''
        },
        validationSchema: validation,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                const response = await Request({ url: 'admin/member/updateEmail', method: 'POST', param: values }, dispatch, intl);
                setIsSubmitting(false);
                if (response?.data?.code == 0) {
                    dispatch(
                        openSnackbar({
                            open: true,
                            message: response?.data?.message,
                            variant: 'alert',
                            alert: {
                                color: 'success'
                            },
                            close: false,
                            anchorOrigin: {
                                vertical: 'top',
                                horizontal: 'center'
                            }
                        })
                    );

                    setTimeout(() => {
                        window?.location?.reload();
                    }, 1500);
                }
                // FAILED
                else {
                    dispatch(
                        openSnackbar({
                            open: true,
                            message: response?.data?.message || 'Something wrong. Please try again later',
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
            } catch (err: any) {
                setIsSubmitting(false);
                dispatch(
                    openSnackbar({
                        open: true,
                        message: '' + err,
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
        }
    });

    useEffect(() => {
        (async function () {
            setIsLoading(true);
            try {
                const response = await Request({ url: 'admin/member/info', method: 'GET' }, dispatch, intl);
                setIsLoading(false);
                if (response?.data?.data?.email) {
                    setValidation(
                        yup.object({
                            code: yup
                                .number()
                                .required(
                                    `${intl.formatMessage({ id: 'security.verificationCodeEmail' })} ${intl.formatMessage({
                                        id: 'error.required'
                                    })}`
                                ),
                            email: yup
                                .string()
                                .required(`${intl.formatMessage({ id: 'security.email' })} ${intl.formatMessage({ id: 'error.required' })}`)
                                .email()
                        })
                    );
                    formik.setFieldValue('code', '');
                    inputFields.unshift({
                        type: 'textLabel',
                        name: 'mailLabel',
                        label: `${intl.formatMessage({ id: 'security.receiveEmail' })}: ${response?.data?.data?.email}`
                    });
                    inputFields.unshift({
                        type: 'number',
                        name: 'code',
                        label: intl.formatMessage({ id: 'security.verificationCodeEmail' }),
                        optionButton: true,
                        optionButtonLabel: intl.formatMessage({ id: 'security.requestVerificationCodeEmail' }),
                        optionButtonFunc: sendOTP,
                        optionButtonTimer: timeRemaining
                    });
                    setInputFields([...inputFields]);
                }
                dispatch(setAdminInfo(response?.data?.data));
            } catch (error) {
                setIsLoading(false);
                dispatch(setAdminInfoError(error));
            }
        })();
    }, []);

    useEffect(() => {
        formik.validateForm();
    }, [validation]);

    useEffect(() => {
        const { adminInfo } = userState;
        let newTemplate: jsonItem = {
            email: { value: intl.formatMessage({ id: 'security.email' }) },
            mailLabel: { value: `${intl.formatMessage({ id: 'security.receiveEmail' })} : ${adminInfo?.email}` },
            code: {
                value: intl.formatMessage({ id: 'security.verificationCodeEmail' }),
                optionButtonLabel: intl.formatMessage({ id: 'security.requestVerificationCodeEmail' })
            }
        };

        inputFields &&
            inputFields?.map((item: any, index: number) => {
                const valueMatch = newTemplate[item?.name as keyof typeof newTemplate];

                if (valueMatch) {
                    if (item.label && valueMatch?.value) {
                        item.label = valueMatch?.value;
                    }

                    if (item?.optionButtonLabel && valueMatch?.optionButtonLabel) {
                        const optionButtonLabelNew = valueMatch['optionButtonLabel' as keyof typeof valueMatch];
                        item.optionButtonLabel = optionButtonLabelNew;
                    }
                }
            });
        setInputFields([...inputFields]);

        if (adminInfo?.email) {
            setValidation(
                yup.object({
                    code: yup
                        .number()
                        .required(
                            `${intl.formatMessage({ id: 'security.verificationCodeEmail' })} ${intl.formatMessage({
                                id: 'error.required'
                            })}`
                        ),
                    email: yup
                        .string()
                        .required(`${intl.formatMessage({ id: 'security.email' })} ${intl.formatMessage({ id: 'error.required' })}`)
                        .email()
                })
            );
        } else {
            setValidation(
                yup.object({
                    email: yup
                        .string()
                        .required(`${intl.formatMessage({ id: 'security.email' })} ${intl.formatMessage({ id: 'error.required' })}`)
                        .email()
                })
            );
        }
    }, [mail]);

    const startTimer = () => {
        let interval = setInterval(async () => {
            let findIndex = inputFields.findIndex((x) => x?.name === 'code');
            if (findIndex >= 0) {
                inputFields[findIndex].optionButtonTimer = timeRemainingRef.current;
                setInputFields([...inputFields]);
            }

            if (timeRemainingRef.current > 0) {
                await setTimeRemaining(timeRemainingRef.current - 1);
            } else {
                clearInterval(interval);
            }
        }, 1000);
    };

    const sendOTP = async () => {
        setIsRequest(true);
        try {
            const response = await Request({ url: 'admin/ems/sendBind', method: 'POST' }, dispatch, intl);
            setIsRequest(false);
            if (response?.data?.code == 0) {
                setTimeRemaining(60);
                startTimer();
                dispatch(
                    openSnackbar({
                        open: true,
                        message: response?.data?.message,
                        variant: 'alert',
                        alert: {
                            color: 'success'
                        },
                        close: false,
                        anchorOrigin: {
                            vertical: 'top',
                            horizontal: 'center'
                        }
                    })
                );
            }
            // FAILED
            else {
                dispatch(
                    openSnackbar({
                        open: true,
                        message: response?.data?.message,
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
        } catch (err: any) {
            setIsRequest(false);
            dispatch(
                openSnackbar({
                    open: true,
                    message: '' + err,
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

    return (
        <>
            {isLoading ? <Loading isTransprent={true}/> : undefined}
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <SubCard title={intl.formatMessage({ id: 'security.email' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} isRequest={isRequest} />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </>
    );
};

export default Email;
