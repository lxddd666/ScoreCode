import { useEffect, useState } from 'react';

// material-ui
import { Grid } from '@mui/material';

// project imports
import SubCard from 'ui-component/cards/SubCard';
import { gridSpacing } from 'store/constant';
import GeneralForm from '../../../ui-component/form';
import Request from 'utils/request';
import { useIntl } from 'react-intl';

// third-party
import { useFormik, FormikProvider } from 'formik';
import * as yup from 'yup';
import { useDispatch, useSelector } from 'store';
import { openSnackbar } from 'store/slices/snackbar';

// constant
import { getAdminInfo } from 'store/slices/user';

// ==============================|| PROFILE 1 - MY ACCOUNT ||============================== //

const Withdrawal = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const intl = useIntl();
    const profileWithdraw = intl.formatMessage({ id: 'profile.withdrawal' });

    const removeImage = async (key: string, id: string) => {
        formik.setFieldValue('payeeCode', '');
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
                    uploadType: 'default'
                }
            },
            dispatch,
            intl
        );

        setIsSubmitting(false);

        if (response?.data?.code == 0) {
            formik.setFieldValue('payeeCode', response?.data?.data?.fileUrl);
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

    const userState = useSelector((state) => state.user);
    const [inputFields, setInputFields]: Array<any> = useState([
        { type: 'text', name: 'name', label: intl.formatMessage({ id: 'profile.withdrawalName' }) },
        { type: 'text', name: 'account', label: intl.formatMessage({ id: 'profile.withdrawalAccount' }) },
        {
            type: 'upload',
            name: 'payeeCode',
            label: intl.formatMessage({ id: 'profile.withdrawalAccountImg' }),
            fileChange: uploadChange,
            fileRemove: removeImage
        },
        { type: 'password', name: 'password', label: intl.formatMessage({ id: 'profile.withdrawalPassword' }) }
    ]);

    useEffect(() => {
        let newTemplate = {
            name: { value: intl.formatMessage({ id: 'profile.withdrawalName' }) },
            account: { value: intl.formatMessage({ id: 'profile.withdrawalAccount' }) },
            payeeCode: { value: intl.formatMessage({ id: 'profile.withdrawalAccountImg' }) },
            password: { value: intl.formatMessage({ id: 'profile.withdrawalPassword' }) }
        };

        inputFields &&
            inputFields?.map((item: any, index: number) => {
                const valueMatch = newTemplate[item?.name as keyof typeof newTemplate];
                if (valueMatch && valueMatch?.value) {
                    item.label = valueMatch?.value;
                }
            });

        setInputFields([...inputFields]);
        formik.validateForm();
    }, [profileWithdraw]);

    const validationSchema = yup.object({
        name: yup
            .string()
            .required(`${intl.formatMessage({ id: 'profile.withdrawalName' })} ${intl.formatMessage({ id: 'error.required' })}`),
        account: yup
            .string()
            .required(`${intl.formatMessage({ id: 'profile.withdrawalAccount' })} ${intl.formatMessage({ id: 'error.required' })}`),
        payeeCode: yup
            .string()
            .required(`${intl.formatMessage({ id: 'profile.withdrawalAccountImg' })} ${intl.formatMessage({ id: 'error.required' })}`),
        password: yup
            .string()
            .required(`${intl.formatMessage({ id: 'profile.withdrawalPassword' })} ${intl.formatMessage({ id: 'error.required' })}`)
    });

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            name: '',
            account: '',
            payeeCode: '',
            password: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                const response = await Request({ url: 'admin/member/updateCash', method: 'POST', param: values }, dispatch, intl);
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
            await dispatch(getAdminInfo(intl));
        })();
    }, []);

    useEffect(() => {
        const { adminInfo } = userState;
        inputFields?.map((item: any, idx: number) => {
            if (adminInfo && adminInfo?.cash && Object.keys(adminInfo?.cash)?.length > 0) {
                //special case
                const userInfo = adminInfo?.cash![item?.name];
                if (userInfo) {
                    formik.setFieldValue(item?.name, userInfo);
                } else {
                    formik.setFieldValue(item?.name, '');
                }
            }else{
                formik.setFieldValue(item?.name, '');
            }
        });
    }, [userState]);

    return (
        <>
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <SubCard title={intl.formatMessage({ id: 'profile.withdrawal' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </>
    );
};

export default Withdrawal;
