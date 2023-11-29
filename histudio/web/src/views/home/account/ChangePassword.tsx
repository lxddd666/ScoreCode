import { useEffect, useState } from 'react';

// material-ui
import { Grid } from '@mui/material';

// project imports
import SubCard from 'ui-component/cards/SubCard';
import { gridSpacing } from 'store/constant';
import GeneralForm from '../../../ui-component/form';
import Request from 'utils/request';
import { useIntl } from 'react-intl';

// thrid party
import { useFormik, FormikProvider } from 'formik';
import * as yup from 'yup';
import { useDispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';

// ==============================|| PROFILE 1 - CHANGE PASSWORD ||============================== //

const ChangePassword = () => {
    const dispatch = useDispatch();
    const intl = useIntl();
    const currentPassword = intl.formatMessage({ id: 'security.currentPassword' });

    useEffect(() => {
        formik?.validateForm();
    }, [currentPassword]);

    const [isSubmitting, setIsSubmitting] = useState(false);

    const inputFields = [
        { type: 'password', name: 'oldPassword', label: intl.formatMessage({ id: 'security.currentPassword' }) },
        { type: 'password', name: 'newPassword', label: intl.formatMessage({ id: 'security.newPassword' }) }
    ];

    const validationSchema = yup.object({
        oldPassword: yup
            .string()
            .required(`${intl.formatMessage({ id: 'security.currentPassword' })} ${intl.formatMessage({ id: 'error.required' })}`),
        newPassword: yup
            .string()
            .required(`${intl.formatMessage({ id: 'security.newPassword' })} ${intl.formatMessage({ id: 'error.required' })}`) //.matches(, 'Is not in correct format')
    });

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            oldPassword: '',
            newPassword: '',
            submit: null
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                const response = await Request({ url: 'admin/member/updatePwd', method: 'POST', param: values }, dispatch , intl);
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
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <SubCard title={intl.formatMessage({ id: 'security.changePassword' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </>
    );
};

export default ChangePassword;
