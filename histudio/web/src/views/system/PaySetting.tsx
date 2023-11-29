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
    payDebug: ParamForm;
    horizontalLineTextAliPay: ParamForm;
    htmlcontentAliPay: ParamForm;
    payAliPayAppId: ParamForm;
    payAliPayPrivateKey: ParamForm;
    payAliPayAppCertPublicKey: ParamForm;
    payAliPayRootCert: ParamForm;
    payAliPayCertPublicKeyRSA2: ParamForm;
    horizontalLineTextWechatPay: ParamForm;
    payWxPayAppId: ParamForm;
    payWxPayMchId: ParamForm;
    payWxPaySerialNo: ParamForm;
    payWxPayAPIv3Key: ParamForm;
    payWxPayPrivateKey: ParamForm;
    horizontalLineTextQQPay: ParamForm;
    payQQPayAppId: ParamForm;
    payQQPayMchId: ParamForm;
    payQQPayApiKey: ParamForm;
}

const PaySetting = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isLoading, setIsLoading] = useState(false);

    const intl = useIntl();
    const smtpServer = intl.formatMessage({ id: 'setting.mail.smtpServer' });

    const handleSwitch = (e: string) => {
        if (e?.toLowerCase() === 'true') {
            formik.setFieldValue('payDebug', false);
        } else {
            formik.setFieldValue('payDebug', true);
        }
    };
    const [inputFields, setInputFields] = useState<ParamForm[]>([
        {
            type: 'switch',
            name: 'payDebug',
            label: intl.formatMessage({ id: 'setting.pay.debug' }),
            desc: intl.formatMessage({ id: 'setting.pay.debug.desc' }),
            change: handleSwitch
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextAliPay',
            label: intl.formatMessage({ id: 'setting.pay.aliPay' })
        },
        {
            type: 'htmlcontent',
            name: 'htmlcontentAliPay',
            value: intl.formatMessage({ id: 'setting.pay.aliPay.desc' }),
            style: {
                padding: '16px',
                marginBottom: '7px',
                borderRadius: '3px'
            }
        },
        {
            type: 'text',
            name: 'payAliPayAppId',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.appId' })
        },

        {
            type: 'text',
            name: 'payAliPayPrivateKey',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.private.key.path' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.private.key.path.desc' })
        },

        {
            type: 'text',
            name: 'payAliPayAppCertPublicKey',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.public.key' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.public.key.desc' })
        },
        {
            type: 'text',
            name: 'payAliPayRootCert',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.cert.root.path' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.cert.root.path.desc' })
        },
        {
            type: 'text',
            name: 'payAliPayCertPublicKeyRSA2',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.public.key.cert.path' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.public.key.cert.path.desc' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextWechatPay',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay' })
        },
        {
            type: 'text',
            name: 'payWxPayAppId',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.appId' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.appId.desc' })
        },

        {
            type: 'text',
            name: 'payWxPayMchId',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.merchant.id' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.merchant.id.desc' })
        },
        {
            type: 'text',
            name: 'payWxPaySerialNo',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.cert.serial.no' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.cert.serial.no.desc' })
        },
        {
            type: 'text',
            name: 'payWxPayAPIv3Key',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.api.v3.key' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.api.v3.key.desc' })
        },
        {
            type: 'text',
            name: 'payWxPayPrivateKey',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.private.key' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.private.key.desc' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextQQPay',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.pay' })
        },
        {
            type: 'text',
            name: 'payQQPayAppId',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.pay.app.id' })
        },
        {
            type: 'text',
            name: 'payQQPayMchId',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.merchant.id' })
        },
        {
            type: 'text',
            name: 'payQQPayApiKey',
            label: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.apiKey' }),
            desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.apiKey.desc' })
        }
    ]);

    useEffect(() => {
        let newTemplate: inputTemplate = {
            payDebug: {
                value: intl.formatMessage({ id: 'setting.pay.debug' }),
                desc: intl.formatMessage({ id: 'setting.pay.debug.desc' })
            },
            horizontalLineTextAliPay: { label: intl.formatMessage({ id: 'setting.pay.aliPay' }) },
            htmlcontentAliPay: { html: intl.formatMessage({ id: 'setting.pay.aliPay.desc' }) },
            payAliPayAppId: { value: intl.formatMessage({ id: 'setting.pay.aliPay.appId' }) },
            payAliPayPrivateKey: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.private.key.path' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.private.key.path.desc' })
            },
            payAliPayAppCertPublicKey: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.public.key' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.public.key.desc' })
            },
            payAliPayRootCert: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.cert.root.path' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.cert.root.path.desc' })
            },
            payAliPayCertPublicKeyRSA2: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.public.key.cert.path' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.public.key.cert.path.desc' })
            },
            horizontalLineTextWechatPay: { label: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay' }) },
            payWxPayAppId: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.appId' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.appId.desc' })
            },
            payWxPayMchId: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.merchant.id' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.merchant.id.desc' })
            },
            payWxPaySerialNo: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.cert.serial.no' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.cert.serial.no.desc' })
            },
            payWxPayAPIv3Key: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.api.v3.key' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.api.v3.key.desc' })
            },
            payWxPayPrivateKey: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.private.key' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.wechat.pay.private.key.desc' })
            },
            horizontalLineTextQQPay: {
                label: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.pay' })
            },
            payQQPayAppId: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.pay.app.id' })
            },
            payQQPayMchId: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.merchant.id' })
            },
            payQQPayApiKey: {
                value: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.apiKey' }),
                desc: intl.formatMessage({ id: 'setting.pay.aliPay.app.qq.apiKey.desc' })
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

                if(valueMatch?.html){
                    item.value = valueMatch?.html;
                }
            });

        setInputFields([...inputFields]);

        formik.validateForm();
    }, [smtpServer]);

    const validationSchema = yup.object({});

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            payDebug: '',
            payAliPayAppId: '',
            payAliPayPrivateKey: '',
            payAliPayAppCertPublicKey: '',
            payAliPayRootCert: '',
            payAliPayCertPublicKeyRSA2: '',
            payWxPayAppId: '',
            payWxPayMchId: '',
            payWxPaySerialNo: '',
            payWxPayAPIv3Key: '',
            payWxPayPrivateKey: '',
            payQQPayAppId: '',
            payQQPayMchId: '',
            payQQPayApiKey: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                let newValues = {
                    group: 'pay',
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

            const res = await Request({ url: `admin/config/get?${new URLSearchParams({ group: 'pay' })}`, method: 'GET' }, dispatch, intl);

            setIsLoading(false);
            if (res?.data?.data?.list) {
                inputFields?.map((item: any, idx: number) => {
                    if (
                        res?.data?.data?.list &&
                        Object.keys(res?.data?.data?.list)?.length > 0 &&
                        item?.type != 'horizontalLineText' &&
                        item?.type != 'htmlcontent'
                    ) {
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
                    <SubCard title={intl.formatMessage({ id: 'setting.paymentConfiguration' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </div>
    );
};

export default PaySetting;
