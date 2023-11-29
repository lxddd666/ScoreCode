import { useEffect } from 'react';
import useState from 'react-usestateref';

// material-ui
import { Grid } from '@mui/material';

// project imports
import SubCard from 'ui-component/cards/SubCard';
import { gridSpacing } from 'store/constant';
import GeneralForm from '../../../ui-component/form';
import { useIntl } from 'react-intl';

// thrid party
import { useFormik, FormikProvider } from 'formik';
import * as yup from 'yup';
import { useDispatch, useSelector } from 'store';
import { openSnackbar } from 'store/slices/snackbar';

import Request from 'utils/request';
import { setAdminInfo, setAdminInfoError } from 'store/slices/user';
import Loading from 'ui-component/Loading';
import { ParamForm } from 'utils/generalInterface';

// ==============================|| PROFILE 1 - CHANGE Phone No ||============================== //
interface jsonFormat {
    value?: string;
    label?: string;
    optionButtonLabel?: string;
}

interface jsonItem {
    mobile: jsonFormat;
    phoneLabel: jsonFormat;
    code: jsonFormat;
}

const PhoneNumber = () => {
    const intl = useIntl();
    const userState = useSelector((state) => state.user);
    const phoneNo = intl.formatMessage({ id: 'security.phoneNumber' });

    const [isLoading, setIsLoading] = useState<boolean>(false);
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
    const [isRequest, setIsRequest] = useState<boolean>(false);
    const [timeRemaining, setTimeRemaining, timeRemainingRef] = useState(0); // initial time in seconds

    const [validation, setValidation] = useState(
        yup.object({
            mobile: yup
                .number()
                .required(`${intl.formatMessage({ id: 'security.phoneNumber' })} ${intl.formatMessage({ id: 'error.required' })}`)
        })
    );

    const dispatch = useDispatch();

    const [inputFields, setInputFields] = useState<ParamForm[]>([
        { type: 'number', name: 'mobile', label: intl.formatMessage({ id: 'security.phoneNumber' }) }
    ]);

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
            const response = await Request({ url: 'admin/sms/sendBind', method: 'POST' }, dispatch, intl);
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

    useEffect(() => {
        (async function () {
            setIsLoading(true);
            try {
                const response = await Request({ url: 'admin/member/info', method: 'GET' }, dispatch, intl);
                setIsLoading(false);
                if (response?.data?.data?.mobile) {
                    setValidation(
                        yup.object({
                            code: yup.number().required(
                                `${intl.formatMessage({ id: 'security.verificationCode' })} ${intl.formatMessage({
                                    id: 'error.required'
                                })}`
                            ),
                            mobile: yup
                                .number()
                                .required(
                                    `${intl.formatMessage({ id: 'security.phoneNumber' })} ${intl.formatMessage({ id: 'error.required' })}`
                                )
                        })
                    );
                    formik.setFieldValue('code', '');
                    inputFields.unshift({
                        type: 'textLabel',
                        name: 'phoneLabel',
                        label: `${intl.formatMessage({ id: 'security.receiveNumber' })}: ${response?.data?.data?.mobile}`
                    });
                    inputFields.unshift({
                        type: 'number',
                        name: 'code',
                        label: `${intl.formatMessage({ id: 'security.verificationCode' })}`,
                        optionButton: true,
                        optionButtonLabel: intl.formatMessage({ id: 'security.requestOtp' }),
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
            mobile: { value: intl.formatMessage({ id: 'security.phoneNumber' }) },
            phoneLabel: { value: `${intl.formatMessage({ id: 'security.receiveNumber' })} : ${adminInfo?.mobile}` },
            code: {
                value: intl.formatMessage({ id: 'security.verificationCode' }),
                optionButtonLabel: intl.formatMessage({ id: 'security.requestOtp' })
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

        if (adminInfo?.mobile) {
            setValidation(
                yup.object({
                    code: yup
                        .number()
                        .required(
                            `${intl.formatMessage({ id: 'security.verificationCode' })} ${intl.formatMessage({ id: 'error.required' })}`
                        ),
                    mobile: yup
                        .number()
                        .required(`${intl.formatMessage({ id: 'security.phoneNumber' })} ${intl.formatMessage({ id: 'error.required' })}`)
                })
            );
        } else {
            setValidation(
                yup.object({
                    mobile: yup
                        .number()
                        .required(`${intl.formatMessage({ id: 'security.phoneNumber' })} ${intl.formatMessage({ id: 'error.required' })}`)
                })
            );
        }
    }, [phoneNo]);

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            mobile: ''
        },
        validationSchema: validation,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                const response = await Request({ url: 'admin/member/updateMobile', method: 'POST', param: values }, dispatch, intl);
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

    return (
        <>
            {isLoading ? <Loading isTransprent={true}/> : undefined}
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <SubCard title={intl.formatMessage({ id: 'security.phoneNumber' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} isRequest={isRequest} />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </>
    );
};

export default PhoneNumber;
