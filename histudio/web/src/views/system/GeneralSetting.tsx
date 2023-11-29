import React, { useEffect, useState } from 'react';

// material-ui
import { Grid, Button } from '@mui/material';

// project imports
import SubCard from 'ui-component/cards/SubCard';
import { gridSpacing } from 'store/constant';
import GeneralForm from '../../ui-component/form';
import Request from 'utils/request';
import { useIntl } from 'react-intl';
import Loading from 'ui-component/Loading';

// third-party
import { useFormik, FormikProvider } from 'formik';
import * as yup from 'yup';
import { useDispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';
import { ParamForm } from 'utils/generalInterface';

import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { TransitionProps } from '@mui/material/transitions';
import Slide from '@mui/material/Slide';
import { AccessPermissions } from 'constant/general';
// ==============================|| PROFILE 1 - MY ACCOUNT ||============================== //

interface inputTemplate {
    basicName: ParamForm;
    basicLogo: ParamForm;
    basicDomain: ParamForm;
    basicWsAddr: ParamForm;
    basicSystemOpen: ParamForm;
    basicCloseText: ParamForm;
    basicIcpCode: ParamForm;
    basicCopyright: ParamForm;
}

const Transition = React.forwardRef(function Transition(
    props: TransitionProps & {
        children: React.ReactElement<any, any>;
    },
    ref: React.Ref<unknown>
) {
    return <Slide direction="down" ref={ref} {...props} />;
});

const GeneralSetting = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [open, setOpen] = React.useState(false);
    const intl = useIntl();
    const copyRight = intl.formatMessage({ id: 'setting.general.copyRight' });

    const handleClose = () => {
        setOpen(false);
    };

    const handleOk = () => {
        setOpen(false);
        const nameFieldValue = formik?.getFieldProps('basicSystemOpen')?.value;
        formik.setFieldValue('basicSystemOpen', !nameFieldValue);
    };

    const openDialog = (e: string) => {
        if (e === 'true') {
            setOpen(true);
        } else {
            handleOk();
        }
    };

    const fileRemove = async () => {
        formik.setFieldValue('basicLogo', '');
    };

    const uploadChange = async (data: any, name: string) => {
        const formData = new FormData();
        formData.append('type', '0');
        formData.append('file', data?.file);

        setIsSubmitting(true);

        const response = await Request(
            {
                url: 'admin/upload/file',
                method: 'POST',
                param: formData,
                header: {
                    'Content-Type': 'multipart/form-data',
                    uploadType: 'image'
                }
            },
            dispatch,
            intl
        );

        setIsSubmitting(false);

        if (response?.data?.code == 0) {
            formik.setFieldValue('basicLogo', response?.data?.data?.fileUrl);
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
        } else {
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
    };

    const [inputFields, setInputFields] = useState<ParamForm[]>([
        { type: 'text', name: 'basicName', label: intl.formatMessage({ id: 'setting.general.web' }), required: true },
        {
            type: 'upload',
            name: 'basicLogo',
            label: intl.formatMessage({ id: 'setting.general.webLogo' }),
            fileChange: uploadChange,
            fileRemove: fileRemove
        },
        {
            type: 'text',
            name: 'basicDomain',
            label: intl.formatMessage({ id: 'setting.general.webDomain' }),
            desc: intl.formatMessage({ id: 'setting.general.webDomainDesc' })
        },
        {
            type: 'text',
            name: 'basicWsAddr',
            label: intl.formatMessage({ id: 'setting.general.webSockertAddress' }),
            desc: intl.formatMessage({ id: 'setting.general.webSockertAddressDesc' })
        },
        { type: 'switch', name: 'basicSystemOpen', label: intl.formatMessage({ id: 'setting.general.webAllowVisit' }), change: openDialog },
        { type: 'textarea', name: 'basicCloseText', label: intl.formatMessage({ id: 'setting.general.webClose' }) },
        { type: 'text', name: 'basicIcpCode', label: intl.formatMessage({ id: 'setting.general.recordNumber' }) },
        { type: 'text', name: 'basicCopyright', label: intl.formatMessage({ id: 'setting.general.copyRight' }) }
    ]);

    useEffect(() => {
        let newTemplate: inputTemplate = {
            basicName: { value: intl.formatMessage({ id: 'setting.general.web' }) },
            basicLogo: { value: intl.formatMessage({ id: 'setting.general.webLogo' }) },
            basicDomain: {
                value: intl.formatMessage({ id: 'setting.general.webDomain' }),
                desc: intl.formatMessage({ id: 'setting.general.webDomainDesc' })
            },
            basicWsAddr: {
                value: intl.formatMessage({ id: 'setting.general.webSockertAddress' }),
                desc: intl.formatMessage({ id: 'setting.general.webSockertAddressDesc' })
            },
            basicSystemOpen: { value: intl.formatMessage({ id: 'setting.general.webAllowVisit' }) },
            basicCloseText: { value: intl.formatMessage({ id: 'setting.general.webClose' }) },
            basicIcpCode: { value: intl.formatMessage({ id: 'setting.general.recordNumber' }) },
            basicCopyright: { value: intl.formatMessage({ id: 'setting.general.copyRight' }) }
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
            });

        setInputFields([...inputFields]);
        formik.validateForm();
    }, [copyRight]);

    const validationSchema = yup.object({
        basicName: yup
            .string()
            .required(`${intl.formatMessage({ id: 'setting.general.web' })} ${intl.formatMessage({ id: 'error.required' })}`)
    });

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            basicName: '',
            basicLogo: '',
            basicDomain: '',
            basicWsAddr: '',
            basicSystemOpen: '',
            basicCloseText: '',
            basicIcpCode: '',
            basicCopyright: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                let newValues = {
                    group: 'basic',
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
            const res = await Request({ url: `admin/config/get?${new URLSearchParams({ group: 'basic' })}`, method: 'GET' }, dispatch, intl);
            setIsLoading(false);
            if (res?.data?.data?.list) {
                inputFields?.map((item: any, idx: number) => {
                    if (res?.data?.data?.list && Object.keys(res?.data?.data?.list)?.length > 0) {
                        //special case
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
                    <SubCard title={intl.formatMessage({ id: 'setting.basic' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm
                                formData={inputFields}
                                formik={formik}
                                isSubmitting={isSubmitting}
                                requiredPermissions={[AccessPermissions.CONFIG_UPDATE]}
                            />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>

            <Dialog
                open={open}
                TransitionComponent={Transition}
                keepMounted
                onClose={handleClose}
                aria-describedby="alert-dialog-slide-description"
            >
                <DialogTitle>{intl.formatMessage({ id: 'general.notice' })}</DialogTitle>
                <DialogContent>
                    <DialogContentText id="alert-dialog-slide-description">
                        {intl.formatMessage({ id: 'setting.general.webClose-alert' })}
                    </DialogContentText>
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>{intl.formatMessage({ id: 'general.cancel' })}</Button>
                    <Button onClick={handleOk}>{intl.formatMessage({ id: 'general.ok' })}</Button>
                </DialogActions>
            </Dialog>
        </div>
    );
};

export default GeneralSetting;
