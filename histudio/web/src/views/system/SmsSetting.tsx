import { useEffect, useState } from 'react';

// material-ui
import { Grid } from '@mui/material';
import Modal from '@mui/material/Modal';
import Box from '@mui/material/Box';
import CloseIcon from '@mui/icons-material/Close';

// project imports
import SubCard from 'ui-component/cards/SubCard';
import { gridSpacing } from 'store/constant';
import GeneralForm from '../../ui-component/form';
import Request from 'utils/request';
import { useIntl } from 'react-intl';
import Loading from 'ui-component/Loading';
import { ParamForm } from 'utils/generalInterface';

// third-party
import { useFormik, FormikProvider } from 'formik';
import * as yup from 'yup';
import { useDispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';

// ==============================|| PROFILE 1 - MY ACCOUNT ||============================== //

interface inputTemplate {
    smsDrive: ParamForm;
    horizontalLineTextSendRestrict: ParamForm;
    smsMinInterval: ParamForm;
    smsMaxIpLimit: ParamForm;
    smsCodeExpire: ParamForm;
    horizontalLineTextAliCloud: ParamForm;
    smsAliYunAccessKeyID: ParamForm;
    smsAliYunAccessKeySecret: ParamForm;
    smsAliYunSign: ParamForm;
    smsAliYunTemplate: ParamForm;
    horizontalLineTextTencentCloud: ParamForm;
    smsTencentSecretId: ParamForm;
    smsTencentSecretKey: ParamForm;
    smsTencentEndpoint: ParamForm;
    smsTencentRegion: ParamForm;
    smsTencentAppId: ParamForm;
    smsTencentSign: ParamForm;
    smsTencentTemplate: ParamForm;
}

interface inputModalTemplate {
    event: ParamForm;
    mobile: ParamForm;
    code: ParamForm;
}

const styleModal = {
    bgcolor: 'background.paper',
    boxShadow: 24,
    p: 3,
    width: '100%'
};

const SmsSetting = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [open, setOpen] = useState(false);
    const intl = useIntl();
    const smtpServer = intl.formatMessage({ id: 'setting.mail.smtpServer' });

    const [inputFieldsModal, setInputFieldsModal] = useState<ParamForm[]>([
        {
            type: 'select',
            name: 'event',
            label: intl.formatMessage({ id: 'setting.sms.eventTemplate' }),
            options: []
        },
        {
            type: 'text',
            name: 'mobile',
            label: intl.formatMessage({ id: 'setting.sms.phoneNumber' })
        },
        {
            type: 'text',
            name: 'code',
            label: intl.formatMessage({ id: 'setting.sms.verfiCode' })
        }
    ]);

    const [inputFields, setInputFields] = useState<ParamForm[]>([
        {
            type: 'select',
            name: 'smsDrive',
            label: intl.formatMessage({ id: 'setting.sms.defaultDriver' }),
            required: true,
            options: []
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextSendRestrict',
            label: intl.formatMessage({ id: 'setting.sms.sendRestrict' })
        },
        {
            type: 'number',
            name: 'smsMinInterval',
            label: intl.formatMessage({ id: 'setting.sms.minSendInterval' }),
            desc: intl.formatMessage({ id: 'setting.sms.minSendInterval-desc' }),
            InputAdornment: intl.formatMessage({ id: 'setting.sms-second' })
        },
        {
            type: 'number',
            name: 'smsMaxIpLimit',
            label: intl.formatMessage({ id: 'setting.sms.ipSendTimes' }),
            desc: intl.formatMessage({ id: 'setting.sms.ipSendTimes-desc' })
        },
        {
            type: 'number',
            name: 'smsCodeExpire',
            label: intl.formatMessage({ id: 'setting.sms.veriCodeValidate' }),
            InputAdornment: intl.formatMessage({ id: 'setting.sms-second' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextAliCloud',
            label: intl.formatMessage({ id: 'setting.sms.aliCloud' })
        },
        {
            type: 'text',
            name: 'smsAliYunAccessKeyID',
            label: intl.formatMessage({ id: 'setting.sms.aliAccessKeyId' }),
            desc: intl.formatMessage({ id: 'setting.sms.aliAccessKeyIdDesc' })
        },
        {
            type: 'password',
            name: 'smsAliYunAccessKeySecret',
            label: intl.formatMessage({ id: 'setting.sms.aliAccessKeySecret' })
        },
        {
            type: 'text',
            name: 'smsAliYunSign',
            label: intl.formatMessage({ id: 'setting.sms.aliSign' }),
            desc: intl.formatMessage({ id: 'setting.sms.aliSignDesc' })
        },
        {
            type: 'DynamicAddRemoveField',
            name: 'smsAliYunTemplate',
            emptyButtonLabel: intl.formatMessage({ id: 'general.add' }),
            label: intl.formatMessage({ id: 'setting.sms.aliTemplate' }),
            idx1: 'key',
            idx2: 'value'
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextTencentCloud',
            label: intl.formatMessage({ id: 'setting.sms.tencentCloud' })
        },
        {
            type: 'text',
            name: 'smsTencentSecretId',
            label: intl.formatMessage({ id: 'setting.sms.tencentSecretId' }),
            desc: intl.formatMessage({ id: 'setting.sms.tencentSecretId-desc' })
        },
        {
            type: 'password',
            name: 'smsTencentSecretKey',
            label: intl.formatMessage({ id: 'setting.sms.tencentSecretKey' })
        },
        {
            type: 'text',
            name: 'smsTencentEndpoint',
            label: intl.formatMessage({ id: 'setting.sms.tencentAccessDomainName' }),
            desc: intl.formatMessage({ id: 'setting.sms.tencentAccessDomainName-desc' })
        },
        {
            type: 'text',
            name: 'smsTencentRegion',
            label: intl.formatMessage({ id: 'setting.sms.tencentGeoInfo' }),
            desc: intl.formatMessage({ id: 'setting.sms.tencentGeoInfo-desc' })
        },
        {
            type: 'text',
            name: 'smsTencentAppId',
            label: intl.formatMessage({ id: 'setting.sms.tencentAppId' }),
            desc: intl.formatMessage({ id: 'setting.sms.tencentAppId-desc' })
        },
        {
            type: 'text',
            name: 'smsTencentSign',
            label: intl.formatMessage({ id: 'setting.sms.tencentSign' }),
            desc: intl.formatMessage({ id: 'setting.sms.tencentSign-desc' })
        },
        {
            type: 'DynamicAddRemoveField',
            name: 'smsTencentTemplate',
            emptyButtonLabel: intl.formatMessage({ id: 'general.add' }),
            label: intl.formatMessage({ id: 'setting.sms.tencentTemplate' }),
            idx1: 'key',
            idx2: 'value'
        }
    ]);

    useEffect(() => {
        let newModalTemplete: inputModalTemplate = {
            event: {
                value: intl.formatMessage({ id: 'setting.sms.eventTemplate' })
            },
            mobile: {
                value: intl.formatMessage({ id: 'setting.sms.phoneNumber' })
            },
            code: {
                value: intl.formatMessage({ id: 'setting.sms.verfiCode' })
            }
        };

        let newTemplate: inputTemplate = {
            smsDrive: {
                value: intl.formatMessage({ id: 'setting.sms.defaultDriver' })
            },
            horizontalLineTextSendRestrict: {
                label: intl.formatMessage({ id: 'setting.sms.sendRestrict' })
            },
            smsMinInterval: {
                value: intl.formatMessage({ id: 'setting.sms.minSendInterval' }),
                desc: intl.formatMessage({ id: 'setting.sms.minSendInterval-desc' }),
                InputAdornment: intl.formatMessage({ id: 'setting.sms-second' })
            },
            smsMaxIpLimit: {
                value: intl.formatMessage({ id: 'setting.sms.ipSendTimes' }),
                desc: intl.formatMessage({ id: 'setting.sms.ipSendTimes-desc' })
            },
            smsCodeExpire: {
                value: intl.formatMessage({ id: 'setting.sms.veriCodeValidate' }),
                InputAdornment: intl.formatMessage({ id: 'setting.sms-second' })
            },
            horizontalLineTextAliCloud: {
                label: intl.formatMessage({ id: 'setting.sms.aliCloud' })
            },
            smsAliYunAccessKeyID: {
                value: intl.formatMessage({ id: 'setting.sms.aliAccessKeyId' }),
                desc: intl.formatMessage({ id: 'setting.sms.aliAccessKeyIdDesc' })
            },
            smsAliYunAccessKeySecret: {
                value: intl.formatMessage({ id: 'setting.sms.aliAccessKeySecret' })
            },
            smsAliYunSign: {
                value: intl.formatMessage({ id: 'setting.sms.aliSign' }),
                desc: intl.formatMessage({ id: 'setting.sms.aliSignDesc' })
            },
            smsAliYunTemplate: {
                value: intl.formatMessage({ id: 'setting.sms.aliTemplate' }),
                emptyButtonLabel: intl.formatMessage({ id: 'general.add' })
            },
            horizontalLineTextTencentCloud: {
                label: intl.formatMessage({ id: 'setting.sms.tencentCloud' })
            },
            smsTencentSecretId: {
                value: intl.formatMessage({ id: 'setting.sms.tencentSecretId' }),
                desc: intl.formatMessage({ id: 'setting.sms.tencentSecretId-desc' })
            },
            smsTencentSecretKey: {
                value: intl.formatMessage({ id: 'setting.sms.tencentSecretKey' })
            },
            smsTencentEndpoint: {
                value: intl.formatMessage({ id: 'setting.sms.tencentAccessDomainName' }),
                desc: intl.formatMessage({ id: 'setting.sms.tencentAccessDomainName-desc' })
            },
            smsTencentRegion: {
                value: intl.formatMessage({ id: 'setting.sms.tencentGeoInfo' }),
                desc: intl.formatMessage({ id: 'setting.sms.tencentGeoInfo-desc' })
            },
            smsTencentAppId: {
                value: intl.formatMessage({ id: 'setting.sms.tencentAppId' }),
                desc: intl.formatMessage({ id: 'setting.sms.tencentAppId-desc' })
            },
            smsTencentSign: {
                value: intl.formatMessage({ id: 'setting.sms.tencentSign' }),
                desc: intl.formatMessage({ id: 'setting.sms.tencentSign-desc' })
            },
            smsTencentTemplate: {
                value: intl.formatMessage({ id: 'setting.sms.tencentTemplate' }),
                emptyButtonLabel: intl.formatMessage({ id: 'general.add' })
            }
        };

        inputFields &&
            inputFields?.map((item: any, index: number) => {
                const valueMatch = newTemplate[item?.name as keyof typeof newTemplate];
                if (valueMatch && valueMatch?.value) {
                    item.label = valueMatch?.value;
                }

                if (valueMatch && valueMatch?.desc) {
                    item.desc = valueMatch?.desc;
                }

                if (item?.type == 'horizontalLineText') {
                    item.label = valueMatch?.label;
                }

                if (valueMatch && valueMatch?.InputAdornment) {
                    item.InputAdornment = valueMatch?.InputAdornment;
                }

                if (valueMatch && valueMatch?.emptyButtonLabel && item?.emptyButtonLabel) {
                    item.emptyButtonLabel = valueMatch?.emptyButtonLabel;
                }
            });

        setInputFields([...inputFields]);

        inputFieldsModal &&
            inputFieldsModal?.map((item: any, index: number) => {
                const valueMatch = newModalTemplete[item?.name as keyof typeof newModalTemplete];
                if (valueMatch && valueMatch?.value) {
                    item.label = valueMatch?.value;
                }

                if (valueMatch && valueMatch?.desc) {
                    item.desc = valueMatch?.desc;
                }

                if (item?.type == 'horizontalLineText') {
                    item.label = valueMatch?.label;
                }

                if (valueMatch && valueMatch?.InputAdornment) {
                    item.InputAdornment = valueMatch?.InputAdornment;
                }
            });

        setInputFieldsModal([...inputFieldsModal]);

        formik.validateForm();
    }, [smtpServer]);

    const validationSchema = yup.object({
        smsDrive: yup
            .string()
            .required(`${intl.formatMessage({ id: 'setting.sms.defaultDriver' })} ${intl.formatMessage({ id: 'error.required' })}`)
    });

    const validationModalSchema = yup.object({
        event: yup
            .string()
            .required(`${intl.formatMessage({ id: 'setting.sms.eventTemplate' })} ${intl.formatMessage({ id: 'error.required' })}`),
        mobile: yup
            .string()
            .required(`${intl.formatMessage({ id: 'setting.sms.phoneNumber' })} ${intl.formatMessage({ id: 'error.required' })}`),
        code: yup
            .string()
            .required(`${intl.formatMessage({ id: 'setting.sms.verfiCode' })} ${intl.formatMessage({ id: 'error.required' })}`)
    });

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            smsDrive: '',
            smsMinInterval: '',
            smsMaxIpLimit: '',
            smsCodeExpire: '',
            smsAliYunAccessKeyID: '',
            smsAliYunAccessKeySecret: '',
            smsAliYunSign: '',
            smsAliYunTemplate: '',
            smsTencentSecretId: '',
            smsTencentSecretKey: '',
            smsTencentEndpoint: '',
            smsTencentRegion: '',
            smsTencentAppId: '',
            smsTencentSign: '',
            smsTencentTemplate: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                let newValues = {
                    group: 'sms',
                    list: values
                };

                const response = await Request({ url: 'admin/config/update', method: 'POST', param: newValues }, dispatch, intl);
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

    const formikModal = useFormik({
        enableReinitialize: true,
        initialValues: {
            event: '',
            mobile: '',
            code: ''
        },
        validationSchema: validationModalSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                setIsSubmitting(true);
                const res = await Request({ url: 'admin/sms/sendTest', method: 'POST', param: values }, dispatch, intl);
                setIsSubmitting(false);
                if (res?.data?.code == 0) {
                    dispatch(
                        openSnackbar({
                            open: true,
                            message: res?.data?.message,
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
                        handleClose();
                        inputFieldsModal?.length > 0 &&
                            inputFieldsModal.map((data: any, index: number) => {
                                formikModal.setFieldValue(data?.name, '');
                            });

                            formikModal?.validateForm();
                    }, 1500);
                } else {
                    dispatch(
                        openSnackbar({
                            open: true,
                            message: res?.data?.message,
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
            let params = new URLSearchParams();
            params.append('types[]', 'config_sms_template');
            params.append('types[]', 'config_sms_drive');

            const request1Promise = Request({ url: `admin/dictData/options?${params.toString()}`, method: 'GET' }, dispatch, intl);
            const request2Promise = Request({ url: `admin/config/get?${new URLSearchParams({ group: 'sms' })}`, method: 'GET' }, dispatch, intl);

            setIsLoading(true);
            const [resOptions, res] = await Promise.all([request1Promise, request2Promise]);

            setIsLoading(false);
            if (res?.data?.data?.list) {
                inputFields?.map((item: any, idx: number) => {
                    if (res?.data?.data?.list && Object.keys(res?.data?.data?.list)?.length > 0 && item?.type != 'horizontalLineText') {
                        const userInfo = res?.data?.data?.list![item?.name];
                        if (userInfo) {
                            if (item?.type == 'DynamicAddRemoveField') {
                                try {
                                    formik.setFieldValue(item?.name, JSON.parse(userInfo));
                                } catch (err) {
                                    console.error('err ', err);
                                }
                            } else {
                                formik.setFieldValue(item?.name, userInfo);
                            }
                        } else {
                            formik.setFieldValue(item?.name, '');
                        }

                        /* Special case */
                        if (item?.name === 'smsDrive') {
                            item.options = resOptions?.data?.data?.config_sms_drive;
                        }
                    }
                });
            }

            if (resOptions?.data?.data?.config_sms_template) {
                inputFieldsModal?.map((item: any, idx: number) => {
                    if (item?.name === 'event') {
                        item.options = resOptions?.data?.data?.config_sms_template;
                    }
                });
            }
        })();
    }, []);

    const sendTestMail = () => {
        handleOpen();
    };

    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);

    return (
        <div className="general-container">
            {isLoading ? <Loading isFixed={false} isTransprent={true} /> : undefined}
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <SubCard title={intl.formatMessage({ id: 'setting.sms' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm
                                formData={inputFields}
                                formik={formik}
                                isSubmitting={isSubmitting}
                                customButtonLabel={intl.formatMessage({ id: 'setting.sms.sendTestSMS' })}
                                customButtonFunc={sendTestMail}
                            />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
            <Modal
                className=""
                open={open}
                onClose={handleClose}
                aria-labelledby="modal-modal-title"
                aria-describedby="modal-modal-description"
            >
                <div className="general-modal-background">
                    <div className="general-modal-dialog">
                        <Box sx={styleModal}>
                            <div className="general-modal-header">
                                <div className="general-modal-title">{intl.formatMessage({ id: 'setting.sms.sendTestSMS' })}</div>
                                <div className="general-modal-close" onClick={handleClose}>
                                    <CloseIcon />
                                </div>
                            </div>

                            <div className="general-modal-body">
                                <FormikProvider value={formikModal}>
                                    <GeneralForm
                                        formData={inputFieldsModal}
                                        formik={formikModal}
                                        isSubmitting={isSubmitting}
                                        showCancel={true}
                                        cancelFunc={handleClose}
                                    />
                                </FormikProvider>
                            </div>
                        </Box>
                    </div>
                </div>
            </Modal>
        </div>
    );
};

export default SmsSetting;
