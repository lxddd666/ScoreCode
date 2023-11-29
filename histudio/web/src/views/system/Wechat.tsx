import { useEffect, useState } from 'react';

// material-ui
import { Grid } from '@mui/material';

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
    horizontalLineTextPublicAcc: ParamForm;
    officialAccountAppId: ParamForm;
    officialAccountAppSecret: ParamForm;
    officialAccountToken: ParamForm;
    officialAccountEncodingAESKey: ParamForm;
    horizontalLineTextOpenPlatform: ParamForm;
    openPlatformAppId: ParamForm;
    openPlatformAppSecret: ParamForm;
    openPlatformToken: ParamForm;
    openPlatformEncodingAESKey: ParamForm;
}

const Wechat = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isLoading, setIsLoading] = useState(false);

    const intl = useIntl();
    const smtpServer = intl.formatMessage({ id: 'setting.mail.smtpServer' });

    const [inputFields, setInputFields] = useState<ParamForm[]>([
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextPublicAcc',
            label: intl.formatMessage({ id: 'setting.wechat.public.acc' })
        },
        {
            type: 'text',
            name: 'officialAccountAppId',
            label: intl.formatMessage({ id: 'setting.wechat.app.id' }),
            desc: intl.formatMessage({ id: 'setting.wechat.app.id.desc' })
        },
        {
            type: 'text',
            name: 'officialAccountAppSecret',
            label: intl.formatMessage({ id: 'setting.wechat.app.secret' }),
            desc: intl.formatMessage({ id: 'setting.wechat.app.secret.desc' })
        },
        {
            type: 'text',
            name: 'officialAccountToken',
            label: intl.formatMessage({ id: 'setting.wechat.app.token' }),
            desc: intl.formatMessage({ id: 'setting.wechat.app.token.desc' })
        },
        {
            type: 'text',
            name: 'officialAccountEncodingAESKey',
            label: intl.formatMessage({ id: 'setting.wechat.app.aes.encode' }),
            desc: intl.formatMessage({ id: 'setting.wechat.app.aes.encode.desc' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextOpenPlatform',
            label: intl.formatMessage({ id: 'setting.wechat.app.open.platform' })
        },
        {
            type: 'text',
            name: 'openPlatformAppId',
            label: intl.formatMessage({ id: 'setting.wechat.app.open.platform.appId' }),
            desc: intl.formatMessage({ id: 'setting.wechat.app.open.platform.appId.desc' })
        },
        {
            type: 'text',
            name: 'openPlatformAppSecret',
            label: intl.formatMessage({ id: 'setting.wechat.app.open.platform.appSecret' }),
            desc: intl.formatMessage({ id: 'setting.wechat.app.open.platform.appSecret.desc' })
        },
        {
            type: 'text',
            name: 'openPlatformToken',
            label: intl.formatMessage({ id: 'setting.wechat.app.open.platform.token' }),
            desc: intl.formatMessage({ id: 'setting.wechat.app.open.platform.token.desc' })
        },
        {
            type: 'text',
            name: 'openPlatformEncodingAESKey',
            label: intl.formatMessage({ id: 'setting.wechat.app.open.platform.encode.aes.key' }),
            desc: intl.formatMessage({ id: 'setting.wechat.app.open.platform.encode.aes.key.desc' })
        }
    ]);

    useEffect(() => {
        let newTemplate: inputTemplate = {
            horizontalLineTextPublicAcc: {
                label: intl.formatMessage({ id: 'setting.wechat.public.acc' })
            },
            officialAccountAppId: {
                value: intl.formatMessage({ id: 'setting.wechat.app.id' }),
                desc: intl.formatMessage({ id: 'setting.wechat.app.id.desc' })
            },
            officialAccountAppSecret: {
                value: intl.formatMessage({ id: 'setting.wechat.app.secret' }),
                desc: intl.formatMessage({ id: 'setting.wechat.app.secret.desc' })
            },
            officialAccountToken: {
                value: intl.formatMessage({ id: 'setting.wechat.app.token' }),
                desc: intl.formatMessage({ id: 'setting.wechat.app.token.desc' })
            },
            officialAccountEncodingAESKey: {
                value: intl.formatMessage({ id: 'setting.wechat.app.aes.encode' }),
                desc: intl.formatMessage({ id: 'setting.wechat.app.aes.encode.desc' })
            },
            horizontalLineTextOpenPlatform: {
                label: intl.formatMessage({ id: 'setting.wechat.app.open.platform' })
            },
            openPlatformAppId: {
                value: intl.formatMessage({ id: 'setting.wechat.app.open.platform.appId' }),
                desc: intl.formatMessage({ id: 'setting.wechat.app.open.platform.appId.desc' })
            },
            openPlatformAppSecret: {
                value: intl.formatMessage({ id: 'setting.wechat.app.open.platform.appSecret' }),
                desc: intl.formatMessage({ id: 'setting.wechat.app.open.platform.appSecret.desc' })
            },
            openPlatformToken: {
                value: intl.formatMessage({ id: 'setting.wechat.app.open.platform.token' }),
                desc: intl.formatMessage({ id: 'setting.wechat.app.open.platform.token.desc' })
            },
            openPlatformEncodingAESKey: {
                value: intl.formatMessage({ id: 'setting.wechat.app.open.platform.encode.aes.key' }),
                desc: intl.formatMessage({ id: 'setting.wechat.app.open.platform.encode.aes.key.desc' })
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
            });

        setInputFields([...inputFields]);

        formik.validateForm();
    }, [smtpServer]);

    const validationSchema = yup.object({});

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            officialAccountAppId: '',
            officialAccountAppSecret: '',
            officialAccountToken: '',
            officialAccountEncodingAESKey: '',
            openPlatformAppId: '',
            openPlatformAppSecret: "",
            openPlatformToken: '',
            openPlatformEncodingAESKey: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                let newValues = {
                    group: 'wechat',
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

            const res = await Request({ url: `admin/config/get?${new URLSearchParams({ group: 'wechat' })}`, method: 'GET' }, dispatch, intl);

            setIsLoading(false);
            if (res?.data?.data?.list) {
                inputFields?.map((item: any, idx: number) => {
                    if (res?.data?.data?.list && Object.keys(res?.data?.data?.list)?.length > 0 && item?.type != 'horizontalLineText') {
                        const userInfo = res?.data?.data?.list![item?.name];
                        if (userInfo) {
                            formik.setFieldValue(item?.name, userInfo);
                        } else {
                            formik.setFieldValue(item?.name, '');
                        }
                    }
                });
            }
        })();
    }, []);

    return (
        <div className="general-container">
            {isLoading ? <Loading isFixed={false} isTransprent={true} /> : undefined}
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <SubCard title={intl.formatMessage({ id: 'setting.wechatConfiguration' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </div>
    );
};

export default Wechat;
