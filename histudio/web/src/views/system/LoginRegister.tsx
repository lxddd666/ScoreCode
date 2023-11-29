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
    horizontalLineTextSetting: ParamForm;
    loginCaptchaSwitch: ParamForm;
    loginRegisterSwitch: ParamForm;
    loginForceInvite: ParamForm;
    loginAutoOpenId: ParamForm;
    horizontalLineTextRegDefaultConfigure: ParamForm;
    loginAvatar: ParamForm;
    loginRoleId: ParamForm;
    loginDeptId: ParamForm;
    loginPostIds: ParamForm;
    horizontalLineTextProtocolConfigure: ParamForm;
    loginProtocol: ParamForm;
    loginPolicy: ParamForm;
}

const LoginRegister = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const intl = useIntl();
    const copyRight = intl.formatMessage({ id: 'setting.general.copyRight' });
    const RadioOption = [
        { label: intl.formatMessage({ id: 'general.on' }), value: 1 },
        { label: intl.formatMessage({ id: 'general.off' }), value: 2 }
    ];

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
            formik.setFieldValue('loginAvatar', response?.data?.data?.fileUrl);
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
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextSetting',
            label: intl.formatMessage({ id: 'setting.loginRegister.setting' })
        },
        {
            type: 'radio',
            name: 'loginCaptchaSwitch',
            label: intl.formatMessage({ id: 'setting.loginRegister.switchVerification' }),
            options: RadioOption
        },
        {
            type: 'radio',
            name: 'loginRegisterSwitch',
            label: intl.formatMessage({ id: 'setting.loginRegister.switchRegister' }),
            options: RadioOption
        },
        {
            type: 'radio',
            name: 'loginForceInvite',
            label: intl.formatMessage({ id: 'setting.loginRegister.forceInvite' }),
            options: RadioOption,
            desc: intl.formatMessage({ id: 'setting.loginRegister.forceInvite-desc' })
        },
        {
            type: 'radio',
            name: 'loginAutoOpenId',
            label: intl.formatMessage({ id: 'setting.loginRegister.getOpenId' }),
            options: RadioOption,
            desc: intl.formatMessage({ id: 'setting.loginRegister.getOpenId-desc' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextRegDefaultConfigure',
            label: intl.formatMessage({ id: 'setting.loginRegister.regDefaultConfigure' })
        },
        {
            type: 'upload',
            name: 'loginAvatar',
            label: intl.formatMessage({ id: 'setting.loginRegister.defaultProfileImg' }),
            fileChange: uploadChange
        },
        {
            type: 'select',
            name: 'loginRoleId',
            label: intl.formatMessage({ id: 'setting.loginRegister.defaultRegisterRole' }),
            options: []
        },
        {
            type: 'select',
            name: 'loginDeptId',
            label: intl.formatMessage({ id: 'setting.loginRegister.defaultRegisterDept' }),
            options: []
        },
        {
            type: 'multiselect',
            name: 'loginPostIds',
            label: intl.formatMessage({ id: 'setting.loginRegister.defaultRegisterPosition' }),
            options: [],
            valueKey: 'id',
            labelKey: 'name'
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextProtocolConfigure',
            label: intl.formatMessage({ id: 'setting.loginRegister.protocolConfigure' })
        },
        {
            type: 'quillEditor',
            name: 'loginProtocol',
            label: intl.formatMessage({ id: 'setting.loginRegister.userProtocol' })
        },
        {
            type: 'quillEditor',
            name: 'loginPolicy',
            label: intl.formatMessage({ id: 'setting.loginRegister.privacyPolicy' })
        }
    ]);

    useEffect(() => {
        let newTemplate: inputTemplate = {
            horizontalLineTextSetting: { label: intl.formatMessage({ id: 'setting.loginRegister.setting' }) },
            loginCaptchaSwitch: { value: intl.formatMessage({ id: 'setting.loginRegister.switchVerification' }) },
            loginRegisterSwitch: {
                value: intl.formatMessage({ id: 'setting.loginRegister.switchRegister' })
            },
            loginForceInvite: {
                value: intl.formatMessage({ id: 'setting.loginRegister.forceInvite' }),
                desc: intl.formatMessage({ id: 'setting.loginRegister.forceInvite-desc' })
            },
            loginAutoOpenId: {
                value: intl.formatMessage({ id: 'setting.loginRegister.getOpenId' }),
                desc: intl.formatMessage({ id: 'setting.loginRegister.getOpenId-desc' })
            },
            horizontalLineTextRegDefaultConfigure: { label: intl.formatMessage({ id: 'setting.loginRegister.regDefaultConfigure' }) },
            loginAvatar: { value: intl.formatMessage({ id: 'setting.loginRegister.defaultProfileImg' }) },
            loginRoleId: { value: intl.formatMessage({ id: 'setting.loginRegister.defaultRegisterRole' }) },
            loginDeptId: { value: intl.formatMessage({ id: 'setting.loginRegister.defaultRegisterDept' }) },
            loginPostIds: { value: intl.formatMessage({ id: 'setting.loginRegister.defaultRegisterPosition' }) },
            horizontalLineTextProtocolConfigure: { label: intl.formatMessage({ id: 'setting.loginRegister.protocolConfigure' }) },
            loginProtocol: { value: intl.formatMessage({ id: 'setting.loginRegister.userProtocol' }) },
            loginPolicy: { value: intl.formatMessage({ id: 'setting.loginRegister.privacyPolicy' }) }
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
    }, [copyRight]);

    const validationSchema = yup.object({});

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            loginCaptchaSwitch: '',
            loginRegisterSwitch: '',
            loginForceInvite: '',
            loginAutoOpenId: '',
            loginAvatar: '',
            loginRoleId: '',
            loginDeptId: '',
            loginPostIds: [],
            loginProtocol: '',
            loginPolicy: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                let newValues = {
                    group: 'login',
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
            const loginInfoReuqest = Request(
                { url: `admin/config/get?${new URLSearchParams({ group: 'login' })}`, method: 'GET' },
                dispatch,
                intl
            );
            const roleInfoRequest = Request(
                { url: `admin/role/list?${new URLSearchParams({ pageSize: '100' })}`, method: 'GET' },
                dispatch,
                intl
            );
            const departmentInfoRequest = Request(
                { url: `admin/dept/option?${new URLSearchParams({ pageSize: '100' })}`, method: 'GET' },
                dispatch,
                intl
            );
            const postInfoRequest = Request(
                { url: `admin/post/list?${new URLSearchParams({ pageSize: '100' })}`, method: 'GET' },
                dispatch,
                intl
            );

            const [res, resRole, resDept, resPost] = await Promise.all([
                loginInfoReuqest,
                roleInfoRequest,
                departmentInfoRequest,
                postInfoRequest
            ]);

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

                    /* Options */
                    if (item?.options) {
                        if (item?.name === 'loginRoleId') {
                            item.options = resRole?.data?.data?.list;
                        }

                        if (item?.name === 'loginDeptId') {
                            item.options = resDept?.data?.data?.list;
                        }

                        if (item?.name === 'loginPostIds') {
                            item.options = resPost?.data?.data?.list;
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
                    <SubCard title={intl.formatMessage({ id: 'setting.loginRegister' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} dispatch={dispatch}/>
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </div>
    );
};

export default LoginRegister;
