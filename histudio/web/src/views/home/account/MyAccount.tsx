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
import { gender } from '../../../constant/general';
import { getAdminInfo } from 'store/slices/user';
import { country, province, city } from 'constant/country';

// ==============================|| PROFILE 1 - MY ACCOUNT ||============================== //

const MyAccount = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const intl = useIntl();
    const accountName = intl.formatMessage({ id: 'myaccount.address' });

    const removeImage = async (key: string, id: string) => {
        formik.setFieldValue('avatar', '');
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
            formik.setFieldValue('avatar', response?.data?.data?.fileUrl);
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

    const countryChange = (e: string) => {
        formik.setFieldValue('province', '');
        formik.setFieldValue('cityId', '');

        const indexProvince = inputFields.findIndex((x: any) => x?.name == 'province');
        const indexCityId = inputFields.findIndex((x: any) => x?.name == 'cityId');

        if (indexCityId >= 0) {
            inputFields[indexCityId]['options'] = [];
        }

        if (indexProvince >= 0) {
            const selectProvince = province![e as keyof typeof province];
            inputFields[indexProvince]['options'] = selectProvince;
        }
        setInputFields([...inputFields]);
    };

    const provinceChange = (e: string) => {
        formik.setFieldValue('cityId', '');
        const indexCityId = inputFields.findIndex((x: any) => x?.name == 'cityId');
        if (indexCityId >= 0) {
            const selectProvince = city![e as keyof typeof city];
            inputFields[indexCityId]['options'] = selectProvince;
        }
    };

    const userState = useSelector((state) => state.user);
    const [inputFields, setInputFields]: Array<any> = useState([
        {
            type: 'upload',
            name: 'avatar',
            label: intl.formatMessage({ id: 'myaccount.avatar' }),
            fileChange: uploadChange,
            fileRemove: removeImage
        },
        { type: 'text', name: 'realName', label: intl.formatMessage({ id: 'myaccount.name' }) },
        { type: 'text', name: 'qq', label: intl.formatMessage({ id: 'myaccount.qq' }) },
        { type: 'datetime', name: 'birthday', label: intl.formatMessage({ id: 'myaccount.dob' }) },
        { type: 'radio', name: 'sex', label: intl.formatMessage({ id: 'myaccount.gender' }), options: gender },
        {
            type: 'select',
            name: 'country',
            label: intl.formatMessage({ id: 'myaccount.country' }),
            onChange: countryChange,
            options: country
        },
        { type: 'select', name: 'province', label: intl.formatMessage({ id: 'myaccount.province' }), onChange: provinceChange },
        { type: 'select', name: 'cityId', label: intl.formatMessage({ id: 'myaccount.city' }) },
        { type: 'textarea', name: 'address', label: intl.formatMessage({ id: 'myaccount.address' }) }
    ]);

    useEffect(() => {
        let newTemplate = {
            avatar: { value: intl.formatMessage({ id: 'myaccount.avatar' }) },
            realName: { value: intl.formatMessage({ id: 'myaccount.name' }) },
            qq: { value: intl.formatMessage({ id: 'myaccount.qq' }) },
            birthday: { value: intl.formatMessage({ id: 'myaccount.dob' }) },
            sex: { value: intl.formatMessage({ id: 'myaccount.gender' }) },
            country: { value: intl.formatMessage({ id: 'myaccount.country' }) },
            province: { value: intl.formatMessage({ id: 'myaccount.province' }) },
            cityId: { value: intl.formatMessage({ id: 'myaccount.city' }) },
            address: { value: intl.formatMessage({ id: 'myaccount.address' }) }
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
    }, [accountName]);

    const validationSchema = yup.object({
        avatar: yup.string().required(`${intl.formatMessage({ id: 'myaccount.avatar' })} ${intl.formatMessage({ id: 'error.required' })}`),
        realName: yup.string().required(`${intl.formatMessage({ id: 'myaccount.name' })} ${intl.formatMessage({ id: 'error.required' })}`), //.matches(, 'Is not in correct format')
        qq: yup.string().required(`${intl.formatMessage({ id: 'myaccount.qq' })} ${intl.formatMessage({ id: 'error.required' })}`),
        birthday: yup
            .string()
            .required(`${intl.formatMessage({ id: 'myaccount.dob' })} ${intl.formatMessage({ id: 'error.required' })}`)
            .matches(/^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$/, 'Date must be in format dd/MM/yyyy'),
        sex: yup.string().required(`${intl.formatMessage({ id: 'myaccount.gender' })} ${intl.formatMessage({ id: 'error.required' })}`),
        address: yup
            .string()
            .required(`${intl.formatMessage({ id: 'myaccount.address' })} ${intl.formatMessage({ id: 'error.required' })}`),
        cityId: yup.string().required(`${intl.formatMessage({ id: 'myaccount.city' })} ${intl.formatMessage({ id: 'error.required' })}`)
    });

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            avatar: '',
            realName: '',
            qq: '',
            birthday: '',
            sex: '',
            address: '',
            cityId: '',
            submit: null,
            country: '',
            province: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                const response = await Request({ url: 'admin/member/updateProfile', method: 'POST', param: values }, dispatch, intl);
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
            if (adminInfo && Object.keys(adminInfo)?.length > 0) {
                //special case
                const userInfo = adminInfo![item?.name];
                if (userInfo) {
                    if (item?.name == 'cityId') {
                        const indexCityId = inputFields.findIndex((x: any) => x?.name == 'cityId');
                        const indexProvince = inputFields.findIndex((x: any) => x?.name == 'province');
                        for (var x in city) {
                            const citySelected = city![x as keyof typeof city];
                            for (let i = 0; i < citySelected?.length; i++) {
                                if (citySelected[i]?.value == userInfo) {
                                    const provinceList = Object.keys(province);
                                    for (let provinceItem = 0; provinceItem < provinceList.length; provinceItem++) {
                                        const provinceSelected = province![provinceList[provinceItem] as keyof typeof province];
                                        for (let j = 0; j < provinceSelected?.length; j++) {
                                            if (provinceSelected[j]?.value == x) {
                                                inputFields[indexProvince]['options'] = provinceSelected;
                                                for (let countryItem = 0; countryItem < country.length; countryItem++) {
                                                    if (country[countryItem]?.value == provinceList[provinceItem]) {
                                                        formik.setFieldValue('country', country[countryItem]?.value);
                                                        break;
                                                    }
                                                }
                                                break;
                                            }
                                        }
                                    }
                                    formik.setFieldValue('province', x);
                                    inputFields[indexCityId]['options'] = citySelected;
                                    break;
                                }
                            }
                        }
                    }
                    formik.setFieldValue(item?.name, userInfo);
                } else {
                    formik.setFieldValue(item?.name, '');
                }
            }
        });
    }, [userState]);

    return (
        <>
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <SubCard title={intl.formatMessage({ id: 'myaccount.title' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </>
    );
};

export default MyAccount;
