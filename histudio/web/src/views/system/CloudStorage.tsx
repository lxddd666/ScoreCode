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
    uploadDrive: ParamForm;
    horizontalLineTextUploadLimit: ParamForm;
    uploadImageSize: ParamForm;
    uploadImageType: ParamForm;
    uploadFileSize: ParamForm;
    uploadFileType: ParamForm;
    horizontalLineTextlocalStorage: ParamForm;
    uploadLocalPath: ParamForm;
    horizontalLineTextuCloud: ParamForm;
    uploadUCloudPublicKey: ParamForm;
    uploadUCloudPrivateKey: ParamForm;
    uploadUCloudPath: ParamForm;
    uploadUCloudBucketHost: ParamForm;
    uploadUCloudBucketName: ParamForm;
    uploadUCloudFileHost: ParamForm;
    uploadUCloudEndpoint: ParamForm;
    horizontalLineTextTencent: ParamForm;
    uploadCosSecretId: ParamForm;
    uploadCosSecretKey: ParamForm;
    uploadCosBucketURL: ParamForm;
    uploadCosPath: ParamForm;
    horizontalLineTextAliCloud: ParamForm;
    uploadOssSecretId: ParamForm;
    uploadOssSecretKey: ParamForm;
    uploadOssEndpoint: ParamForm;
    uploadOssPath: ParamForm;
    uploadOssBucket: ParamForm;
    uploadOssBucketURL: ParamForm;
    horizontalLineTextQiNiu: ParamForm;
    uploadQiNiuAccessKey: ParamForm;
    uploadQiNiuSecretKey: ParamForm;
    uploadQiNiuPath: ParamForm;
    uploadQiNiuBucket: ParamForm;
    uploadQiNiuDomain: ParamForm;
}

const CloudStorage = () => {
    const dispatch = useDispatch();
    const [isSubmitting, setIsSubmitting] = useState(false);
    const [isLoading, setIsLoading] = useState(false);

    const intl = useIntl();
    const smtpServer = intl.formatMessage({ id: 'setting.mail.smtpServer' });

    const [inputFields, setInputFields] = useState<ParamForm[]>([
        {
            type: 'select',
            name: 'uploadDrive',
            label: intl.formatMessage({ id: 'setting.cloud.defaultDriver' }),
            required: true,
            options: []
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextUploadLimit',
            label: intl.formatMessage({ id: 'setting.cloud.uploadLimit' })
        },
        {
            type: 'number',
            name: 'uploadImageSize',
            label: intl.formatMessage({ id: 'setting.cloud.uploadLimit-Image' }),
            InputAdornment: intl.formatMessage({ id: 'setting.cloud.uploadLimit-Image-Unit' })
        },
        {
            type: 'text',
            name: 'uploadImageType',
            label: intl.formatMessage({ id: 'setting.cloud.uploadLimit-Image-Restrict' })
        },
        {
            type: 'number',
            name: 'uploadFileSize',
            label: intl.formatMessage({ id: 'setting.cloud.fileLimit' }),
            InputAdornment: intl.formatMessage({ id: 'setting.cloud.fileLimit-Unit' })
        },
        {
            type: 'text',
            name: 'uploadFileType',
            label: intl.formatMessage({ id: 'setting.cloud.fileLimit-Restrict' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextlocalStorage',
            label: intl.formatMessage({ id: 'setting.cloud.localStorage' })
        },
        {
            type: 'text',
            name: 'uploadLocalPath',
            label: intl.formatMessage({ id: 'setting.cloud.localStorage-Path' }),
            desc: intl.formatMessage({ id: 'setting.cloud.localStorage-Path-desc' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextuCloud',
            label: intl.formatMessage({ id: 'setting.cloud.uCloud' })
        },
        {
            type: 'password',
            name: 'uploadUCloudPublicKey',
            label: intl.formatMessage({ id: 'setting.cloud.uCloud.public-key' }),
            desc: intl.formatMessage({ id: 'setting.cloud.uCloud.public-key-desc' })
        },
        {
            type: 'password',
            name: 'uploadUCloudPrivateKey',
            label: intl.formatMessage({ id: 'setting.cloud.uCloud.secret-key' })
        },
        {
            type: 'text',
            name: 'uploadUCloudPath',
            label: intl.formatMessage({ id: 'setting.cloud.uCloud.storage-path' }),
            desc: intl.formatMessage({ id: 'setting.cloud.uCloud.storage.path.desc' })
        },
        {
            type: 'text',
            name: 'uploadUCloudBucketHost',
            label: intl.formatMessage({ id: 'setting.cloud.uCloud.storage.regional.api' })
        },
        {
            type: 'text',
            name: 'uploadUCloudBucketName',
            label: intl.formatMessage({ id: 'setting.cloud.uCloud.bucket.name' }),
            desc: intl.formatMessage({ id: 'setting.cloud.uCloud.bucket.name.desc' })
        },
        {
            type: 'text',
            name: 'uploadUCloudFileHost',
            label: intl.formatMessage({ id: 'setting.cloud.uCloud.bucket.host' }),
            desc: intl.formatMessage({ id: 'setting.cloud.uCloud.bucket.host.desc' })
        },
        {
            type: 'text',
            name: 'uploadUCloudEndpoint',
            label: intl.formatMessage({ id: 'setting.cloud.uCloud.visit.domain.name' }),
            desc: intl.formatMessage({ id: 'setting.cloud.uCloud.visit.domain.name.desc' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextTencent',
            label: intl.formatMessage({ id: 'setting.cloud.tencent' })
        },
        {
            type: 'password',
            name: 'uploadCosSecretId',
            label: intl.formatMessage({ id: 'setting.cloud.tencent.secret.id' }),
            desc: intl.formatMessage({ id: 'setting.cloud.tencent.secret.id.desc' })
        },
        {
            type: 'password',
            name: 'uploadCosSecretKey',
            label: intl.formatMessage({ id: 'setting.cloud.tencent.secret.key' })
        },
        {
            type: 'text',
            name: 'uploadCosBucketURL',
            label: intl.formatMessage({ id: 'setting.cloud.tencent.storagePath' }),
            desc: intl.formatMessage({ id: 'setting.cloud.tencent.storagePath.desc' })
        },

        {
            type: 'text',
            name: 'uploadCosPath',
            label: intl.formatMessage({ id: 'setting.cloud.tencent.regional.api' }),
            desc: intl.formatMessage({ id: 'setting.cloud.tencent.regional.api.desc' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextAliCloud',
            label: intl.formatMessage({ id: 'setting.cloud.aliCloud' })
        },

        {
            type: 'password',
            name: 'uploadOssSecretId',
            label: intl.formatMessage({ id: 'setting.cloud.aliCloud.access-id' }),
            desc: intl.formatMessage({ id: 'setting.cloud.aliCloud.access.id.desc' })
        },
        {
            type: 'password',
            name: 'uploadOssSecretKey',
            label: intl.formatMessage({ id: 'setting.cloud.aliCloud.access.secret' })
        },
        {
            type: 'text',
            name: 'uploadOssEndpoint',
            label: intl.formatMessage({ id: 'setting.cloud.aliCloud.endPoint' }),
            desc: intl.formatMessage({ id: 'setting.cloud.aliCloud.endPoint-desc' })
        },
        {
            type: 'text',
            name: 'uploadOssPath',
            label: intl.formatMessage({ id: 'setting.cloud.aliCloud.storage.path' }),
            desc: intl.formatMessage({ id: 'setting.cloud.aliCloud.storage.path.desc' })
        },
        {
            type: 'text',
            name: 'uploadOssBucket',
            label: intl.formatMessage({ id: 'setting.cloud.aliCloud.storage.name' })
        },
        {
            type: 'text',
            name: 'uploadOssBucketURL',
            label: intl.formatMessage({ id: 'setting.cloud.aliCloud.bucket.domain.name' })
        },
        {
            type: 'horizontalLineText',
            name: 'horizontalLineTextQiNiu',
            label: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storage' })
        },
        {
            type: 'password',
            name: 'uploadQiNiuAccessKey',
            label: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storage.accesskey' }),
            desc: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storage.accesskey.desc' })
        },
        {
            type: 'password',
            name: 'uploadQiNiuSecretKey',
            label: intl.formatMessage({ id: 'setting.cloud.qiniucloud.secretkey' })
        },
        {
            type: 'text',
            name: 'uploadQiNiuPath',
            label: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storagePath' }),
            desc: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storagePath.desc' })
        },
        {
            type: 'text',
            name: 'uploadQiNiuBucket',
            label: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storage.name' })
        },

        {
            type: 'text',
            name: 'uploadQiNiuDomain',
            label: intl.formatMessage({ id: 'setting.cloud.qiniucloud.visit.domain.name' })
        }
    ]);

    useEffect(() => {
        let newTemplate: inputTemplate = {
            uploadDrive: {
                value: intl.formatMessage({ id: 'setting.cloud.defaultDriver' })
            },
            horizontalLineTextUploadLimit: {
                label: intl.formatMessage({ id: 'setting.cloud.uploadLimit' })
            },
            uploadImageSize: {
                value: intl.formatMessage({ id: 'setting.cloud.uploadLimit-Image' }),
                InputAdornment: intl.formatMessage({ id: 'setting.cloud.uploadLimit-Image-Unit' })
            },
            uploadImageType: {
                value: intl.formatMessage({ id: 'setting.cloud.uploadLimit-Image-Restrict' })
            },
            uploadFileSize: {
                value: intl.formatMessage({ id: 'setting.cloud.fileLimit' }),
                InputAdornment: intl.formatMessage({ id: 'setting.cloud.fileLimit-Unit' })
            },
            uploadFileType: {
                value: intl.formatMessage({ id: 'setting.cloud.fileLimit-Restrict' })
            },
            horizontalLineTextlocalStorage: {
                label: intl.formatMessage({ id: 'setting.cloud.localStorage' })
            },
            uploadLocalPath: {
                value: intl.formatMessage({ id: 'setting.cloud.localStorage-Path' }),
                desc: intl.formatMessage({ id: 'setting.cloud.localStorage-Path-desc' })
            },
            horizontalLineTextuCloud: {
                label: intl.formatMessage({ id: 'setting.cloud.uCloud' })
            },
            uploadUCloudPublicKey: {
                value: intl.formatMessage({ id: 'setting.cloud.uCloud.public-key' }),
                desc: intl.formatMessage({ id: 'setting.cloud.uCloud.public-key-desc' })
            },
            uploadUCloudPrivateKey: {
                value: intl.formatMessage({ id: 'setting.cloud.uCloud.secret-key' })
            },
            uploadUCloudPath: {
                value: intl.formatMessage({ id: 'setting.cloud.uCloud.storage-path' }),
                desc: intl.formatMessage({ id: 'setting.cloud.uCloud.storage.path.desc' })
            },
            uploadUCloudBucketHost: {
                value: intl.formatMessage({ id: 'setting.cloud.uCloud.storage.regional.api' })
            },
            uploadUCloudBucketName: {
                value: intl.formatMessage({ id: 'setting.cloud.uCloud.bucket.name' }),
                desc: intl.formatMessage({ id: 'setting.cloud.uCloud.bucket.name.desc' })
            },
            uploadUCloudFileHost: {
                value: intl.formatMessage({ id: 'setting.cloud.uCloud.bucket.host' }),
                desc: intl.formatMessage({ id: 'setting.cloud.uCloud.bucket.host.desc' })
            },
            uploadUCloudEndpoint: {
                value: intl.formatMessage({ id: 'setting.cloud.uCloud.visit.domain.name' }),
                desc: intl.formatMessage({ id: 'setting.cloud.uCloud.visit.domain.name.desc' })
            },
            horizontalLineTextTencent: {
                label: intl.formatMessage({ id: 'setting.cloud.tencent' })
            },
            uploadCosSecretId: {
                value: intl.formatMessage({ id: 'setting.cloud.tencent.secret.id' }),
                desc: intl.formatMessage({ id: 'setting.cloud.tencent.secret.id.desc' })
            },
            uploadCosSecretKey: {
                value: intl.formatMessage({ id: 'setting.cloud.tencent.secret.key' })
            },
            uploadCosBucketURL: {
                value: intl.formatMessage({ id: 'setting.cloud.tencent.storagePath' }),
                desc: intl.formatMessage({ id: 'setting.cloud.tencent.storagePath.desc' })
            },
            uploadCosPath: {
                value: intl.formatMessage({ id: 'setting.cloud.tencent.regional.api' }),
                desc: intl.formatMessage({ id: 'setting.cloud.tencent.regional.api.desc' })
            },
            horizontalLineTextAliCloud: {
                label: intl.formatMessage({ id: 'setting.cloud.aliCloud' })
            },
            uploadOssSecretId: {
                value: intl.formatMessage({ id: 'setting.cloud.aliCloud.access-id' }),
                desc: intl.formatMessage({ id: 'setting.cloud.aliCloud.access.id.desc' })
            },
            uploadOssSecretKey: {
                value: intl.formatMessage({ id: 'setting.cloud.aliCloud.access.secret' })
            },
            uploadOssEndpoint: {
                value: intl.formatMessage({ id: 'setting.cloud.aliCloud.endPoint' }),
                desc: intl.formatMessage({ id: 'setting.cloud.aliCloud.endPoint-desc' })
            },
            uploadOssPath: {
                value: intl.formatMessage({ id: 'setting.cloud.aliCloud.storage.path' }),
                desc: intl.formatMessage({ id: 'setting.cloud.aliCloud.storage.path.desc' })
            },
            uploadOssBucket: {
                value: intl.formatMessage({ id: 'setting.cloud.aliCloud.storage.name' })
            },
            uploadOssBucketURL: {
                value: intl.formatMessage({ id: 'setting.cloud.aliCloud.bucket.domain.name' })
            },
            horizontalLineTextQiNiu: {
                label: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storage' })
            },
            uploadQiNiuAccessKey: {
                value: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storage.accesskey' }),
                desc: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storage.accesskey.desc' })
            },
            uploadQiNiuSecretKey: {
                value: intl.formatMessage({ id: 'setting.cloud.qiniucloud.secretkey' })
            },
            uploadQiNiuPath: {
                value: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storagePath' }),
                desc: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storagePath.desc' })
            },
            uploadQiNiuBucket: {
                value: intl.formatMessage({ id: 'setting.cloud.qiniucloud.storage.name' })
            },
            uploadQiNiuDomain: {
                value: intl.formatMessage({ id: 'setting.cloud.qiniucloud.visit.domain.name' })
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

        formik.validateForm();
    }, [smtpServer]);

    const validationSchema = yup.object({
        uploadDrive: yup
            .string()
            .required(`${intl.formatMessage({ id: 'setting.cloud.defaultDriver' })} ${intl.formatMessage({ id: 'error.required' })}`)
    });

    const formik = useFormik({
        enableReinitialize: true,
        initialValues: {
            uploadDrive: '',
            uploadImageSize: '',
            uploadImageType: '',
            uploadFileSize: '',
            uploadFileType: '',
            uploadLocalPath: '',
            uploadUCloudPublicKey: '',
            uploadUCloudPrivateKey: '',
            uploadUCloudPath: '',
            uploadUCloudBucketHost: '',
            uploadUCloudBucketName: '',
            uploadUCloudFileHost: '',
            uploadUCloudEndpoint: '',
            uploadCosSecretId: '',
            uploadCosSecretKey: '',
            uploadCosBucketURL: '',
            uploadCosPath: '',
            uploadOssSecretId: '',
            uploadOssSecretKey: '',
            uploadOssEndpoint: '',
            uploadOssPath: '',
            uploadOssBucket: '',
            uploadOssBucketURL: '',
            uploadQiNiuAccessKey: '',
            uploadQiNiuSecretKey: '',
            uploadQiNiuPath: '',
            uploadQiNiuBucket: '',
            uploadQiNiuDomain: ''
        },
        validationSchema,
        onSubmit: async (values, { setErrors, setStatus, setSubmitting }) => {
            setIsSubmitting(true);
            try {
                let newValues = {
                    group: 'upload',
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
            let params = new URLSearchParams();
            params.append('types[]', 'config_upload_drive');

            const request1Promise = Request({ url: `admin/dictData/options?${params.toString()}`, method: 'GET' }, dispatch, intl);
            const request2Promise = Request(
                { url: `admin/config/get?${new URLSearchParams({ group: 'upload' })}`, method: 'GET' },
                dispatch,
                intl
            );

            setIsLoading(true);
            const [resOptions, res] = await Promise.all([request1Promise, request2Promise]);

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

                        /* Special case */
                        if (item?.name === 'uploadDrive') {
                            item.options = resOptions?.data?.data?.config_upload_drive;
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
                    <SubCard title={intl.formatMessage({ id: 'setting.cloudStorage' })}>
                        <FormikProvider value={formik}>
                            <GeneralForm formData={inputFields} formik={formik} isSubmitting={isSubmitting} />
                        </FormikProvider>
                    </SubCard>
                </Grid>
            </Grid>
        </div>
    );
};

export default CloudStorage;
