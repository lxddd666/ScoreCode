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

// third-party
import { useFormik, FormikProvider } from 'formik';
import * as yup from 'yup';
import { useDispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';
import { ParamForm } from 'utils/generalInterface';

// ==============================|| PROFILE 1 - MY ACCOUNT ||============================== //

interface inputTemplate {
    geoAmapWebKey: ParamForm;
}

const GeoLocation = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const intl = useIntl();
    const settingGeoWebKey = intl.formatMessage({ id: 'setting.geo.webKey' });

    const [inputFields, setInputFields] = useState<ParamForm[]>([
        {
            type: 'password',
            name: 'geoAmapWebKey',
            label: intl.formatMessage({ id: 'setting.geo.webKey' }),
            desc:  intl.formatMessage({ id: 'setting.geo.webKey.desc' }),
        }
    ]);

    useEffect(() => {
        let newTemplate: inputTemplate = {
            geoAmapWebKey: { value:  intl.formatMessage({ id: 'setting.geo.webKey' }) , desc :  intl.formatMessage({ id: 'setting.geo.webKey.desc' })}
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
    }, [settingGeoWebKey]);

    const validationSchema = yup.object({});

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            geoAmapWebKey: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                let newValues = {
                    group: 'geo',
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
            const res = await Request({ url: `admin/config/get?${new URLSearchParams({ group: 'geo' })}`, method: 'GET' }, dispatch, intl);

            setIsLoading(false);
            if (res?.data?.data?.list) {
                inputFields?.map((item: any, idx: number) => {
                    if (res?.data?.data?.list && Object.keys(res?.data?.data?.list)?.length > 0 && item?.type != 'horizontalLineText') {
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
                    <SubCard title={intl.formatMessage({ id: 'setting.geographicalLocation' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} dispatch={dispatch} />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </div>
    );
};

export default GeoLocation;
