import { useEffect, useState, useRef } from 'react';

// material-ui
import { Grid, Button } from '@mui/material';
import Modal from '@mui/material/Modal';
import Box from '@mui/material/Box';
import CloseIcon from '@mui/icons-material/Close';
import TextField from '@mui/material/TextField';
import TextareaAutosize from '@mui/material/TextareaAutosize';
import AnimateButton from 'ui-component/extended/AnimateButton';
import { Spin } from 'antd';
import { LoadingOutlined } from '@ant-design/icons';

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
    smtpHost: ParamForm;
    smtpPort: ParamForm;
    smtpUser: ParamForm;
    smtpPass: ParamForm;
    smtpSendName: ParamForm;
    smtpAdminMailbox: ParamForm;
    smtpMinInterval: ParamForm;
    smtpMaxIpLimit: ParamForm;
    smtpCodeExpire: ParamForm;
    smtpTemplate: ParamForm;
    horizontalLineTextMailLimit: ParamForm;
}

const styleModal = {
    bgcolor: 'background.paper',
    boxShadow: 24,
    p: 3,
    width: '100%'
};

const antIcon = (
    <LoadingOutlined
        style={{
            fontSize: 20,
            marginLeft: 8
        }}
        spin
    />
);

const MailSetting = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [open, setOpen] = useState(false);
    const intl = useIntl();

    const textAreaRef = useRef<HTMLTextAreaElement | null>(null);
    const smtpServer = intl.formatMessage({ id: 'setting.mail.smtpServer' });

    const [inputFields, setInputFields] = useState<ParamForm[]>([
        {
            type: 'text',
            name: 'smtpHost',
            label: intl.formatMessage({ id: 'setting.mail.smtpServer' }),
            desc: intl.formatMessage({ id: 'setting.mail.smtpServerDesc' }),
            required: true
        },
        {
            type: 'text',
            name: 'smtpPort',
            label: intl.formatMessage({ id: 'setting.mail.smtpPort' }),
            desc: intl.formatMessage({ id: 'setting.mail.smtpPortDesc' })
        },
        {
            type: 'text',
            name: 'smtpUser',
            label: intl.formatMessage({ id: 'setting.mail.smtpUsername' }),
            desc: intl.formatMessage({ id: 'setting.mail.smtpUsernameDesc' })
        },
        {
            type: 'password',
            name: 'smtpPass',
            label: intl.formatMessage({ id: 'setting.mail.smtpPassword' }),
            desc: intl.formatMessage({ id: 'setting.mail.smtpPasswordDesc' })
        },
        { type: 'text', name: 'smtpSendName', label: intl.formatMessage({ id: 'setting.mail.smtpSender' }) },
        { type: 'text', name: 'smtpAdminMailbox', label: intl.formatMessage({ id: 'setting.mail.smtpAdminEmail' }) },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextMailLimit',
            label: intl.formatMessage({ id: 'setting.mail.smtpMailLimit' })
        },
        {
            type: 'number',
            name: 'smtpMinInterval',
            label: intl.formatMessage({ id: 'setting.mail.smtpMinSendInterval' }),
            desc: intl.formatMessage({ id: 'setting.mail.smtpMinSendInterval-desc' }),
            InputAdornment: intl.formatMessage({ id: 'setting.mail.smtpSecond' })
        },
        {
            type: 'number',
            name: 'smtpMaxIpLimit',
            label: intl.formatMessage({ id: 'setting.mail.smtpIpSendTimes' }),
            desc: intl.formatMessage({ id: 'setting.mail.smtpIpSendTimes-desc' })
        },
        {
            type: 'number',
            name: 'smtpCodeExpire',
            label: intl.formatMessage({ id: 'setting.mail.smtpVeriCodeValidate' }),
            InputAdornment: intl.formatMessage({ id: 'setting.mail.smtpSecond' })
        },
        {
            type: 'DynamicAddRemoveField',
            name: 'smtpTemplate',
            emptyButtonLabel: intl.formatMessage({ id: 'general.add' }),
            label: intl.formatMessage({ id: 'setting.mail.smtpMailTemplate' }),
            idx1: 'key',
            idx2: 'value'
        }
    ]);

    useEffect(() => {
        let newTemplate: inputTemplate = {
            smtpHost: {
                value: intl.formatMessage({ id: 'setting.mail.smtpServer' }),
                desc: intl.formatMessage({ id: 'setting.mail.smtpServerDesc' })
            },
            smtpPort: {
                value: intl.formatMessage({ id: 'setting.mail.smtpPort' }),
                desc: intl.formatMessage({ id: 'setting.mail.smtpPortDesc' })
            },
            smtpUser: {
                value: intl.formatMessage({ id: 'setting.mail.smtpUsername' }),
                desc: intl.formatMessage({ id: 'setting.mail.smtpUsernameDesc' })
            },
            smtpPass: {
                value: intl.formatMessage({ id: 'setting.mail.smtpPassword' }),
                desc: intl.formatMessage({ id: 'setting.mail.smtpPasswordDesc' })
            },
            smtpSendName: { value: intl.formatMessage({ id: 'setting.mail.smtpSender' }) },
            smtpAdminMailbox: { value: intl.formatMessage({ id: 'setting.mail.smtpAdminEmail' }) },
            smtpMinInterval: {
                value: intl.formatMessage({ id: 'setting.mail.smtpMinSendInterval' }),
                desc: intl.formatMessage({ id: 'setting.mail.smtpMinSendInterval-desc' }),
                InputAdornment: intl.formatMessage({ id: 'setting.mail.smtpSecond' })
            },
            smtpMaxIpLimit: {
                value: intl.formatMessage({ id: 'setting.mail.smtpIpSendTimes' }),
                desc: intl.formatMessage({ id: 'setting.mail.smtpIpSendTimes-desc' })
            },
            smtpCodeExpire: {
                value: intl.formatMessage({ id: 'setting.mail.smtpVeriCodeValidate' }),
                InputAdornment: intl.formatMessage({ id: 'setting.mail.smtpSecond' })
            },
            smtpTemplate: {
                value: intl.formatMessage({ id: 'setting.mail.smtpMailTemplate' }),
                emptyButtonLabel: intl.formatMessage({ id: 'general.add' })
            },
            horizontalLineTextMailLimit: { label: intl.formatMessage({ id: 'setting.mail.smtpMailLimit' }) }
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
        formik.validateForm();
    }, [smtpServer]);

    const validationSchema = yup.object({
        smtpHost: yup
            .string()
            .required(`${intl.formatMessage({ id: 'setting.mail.smtpServer' })} ${intl.formatMessage({ id: 'error.required' })}`)
    });

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            smtpHost: '',
            smtpPort: '',
            smtpUser: '',
            smtpPass: '',
            smtpSendName: '',
            smtpAdminMailbox: '',
            smtpMinInterval: '',
            smtpMaxIpLimit: '',
            smtpCodeExpire: '',
            smtpTemplate: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                let newValues = {
                    group: 'smtp',
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

    useEffect(() => {
        (async function () {
            setIsLoading(true);
            const res = await Request({ url: `admin/config/get?${new URLSearchParams({ group: 'smtp' })}`, method: 'GET' }, dispatch, intl);
            setIsLoading(false);
            if (res?.data?.data?.list) {
                inputFields?.map((item: any, idx: number) => {
                    if (res?.data?.data?.list && Object.keys(res?.data?.data?.list)?.length > 0 && item?.type != 'horizontalLineText') {
                        //special case
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
                    }
                });
            }
        })();
    }, []);

    const sendTestMail = () => {
        handleOpen();
    };

    const submitTestMail = async () => {
        if (textAreaRef.current) {
            setIsSubmitting(true);
            const res = await Request({ url: 'admin/ems/sendTest', method: 'POST', param: { to: textAreaRef.current?.value } }, dispatch, intl);
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

                    if (textAreaRef.current) {
                        textAreaRef.current.value = '';
                    }
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
        }
    };

    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);

    return (
        <div className="general-container">
            {isLoading ? <Loading isFixed={false} isTransprent={true} /> : undefined}
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <SubCard title={intl.formatMessage({ id: 'setting.mail' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm
                                formData={inputFields}
                                formik={formik}
                                isSubmitting={isSubmitting}
                                customButtonLabel={intl.formatMessage({ id: 'setting.mail.smtpSendTestMail' })}
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
                                <div className="general-modal-title">{intl.formatMessage({ id: 'setting.mail.smtpSendTestMail' })}</div>
                                <div className="general-modal-close" onClick={handleClose}>
                                    <CloseIcon />
                                </div>
                            </div>

                            <div className="general-modal-body mail-modal-body">
                                <Grid container spacing={gridSpacing}>
                                    <Grid item xs={12} sm={4}>
                                        <div className="general-modal-content-label">
                                            {intl.formatMessage({ id: 'security.receiveEmail' })}
                                        </div>
                                    </Grid>
                                    <Grid item xs={12} sm={8} style={{ display: 'flex' }}>
                                        <TextField
                                            fullWidth
                                            InputProps={{
                                                inputComponent: TextareaAutosize,
                                                inputProps: {
                                                    minRows: 5,
                                                    maxRows: 5,
                                                    ref: textAreaRef,
                                                    onFocus: () => {},
                                                    onBlur: () => {},
                                                    sx: { padding: 0, margin: '15.5px 14px', borderRadius: 0, resize: 'vertical' }
                                                }
                                            }}
                                        />
                                    </Grid>
                                </Grid>
                            </div>

                            <div className="general-modal-footer">
                                <AnimateButton>
                                    <Button
                                        variant="contained"
                                        style={{ opacity: isSubmitting ? 0.5 : 1, display: 'flex', marginRight: 5 }}
                                        color="primary"
                                        onClick={handleClose}
                                    >
                                        {intl.formatMessage({ id: 'general.cancel' })}
                                    </Button>
                                </AnimateButton>
                                <AnimateButton>
                                    <Button
                                        variant="contained"
                                        disabled={isSubmitting ?? true}
                                        style={{ opacity: isSubmitting ? 0.5 : 1, display: 'flex' }}
                                        color="secondary"
                                        onClick={submitTestMail}
                                    >
                                        {intl.formatMessage({ id: 'general.ok' })}
                                        {isSubmitting ? <Spin indicator={antIcon} /> : <></>}
                                    </Button>
                                </AnimateButton>
                            </div>
                        </Box>
                    </div>
                </div>
            </Modal>
        </div>
    );
};

export default MailSetting;
