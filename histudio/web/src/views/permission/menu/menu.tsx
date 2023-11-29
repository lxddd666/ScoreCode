import React from 'react';
import { useIntl } from 'react-intl';

import { useDispatch } from 'store';
import { openSnackbar } from 'store/slices/snackbar';
import { MenuListData } from 'types/user';

import useScriptRef from 'hooks/useScriptRef';

import { Box, Button, Divider, FormControl, FormHelperText, Grid, InputLabel, OutlinedInput, Typography, useTheme } from '@mui/material';
import EditTwoToneIcon from '@mui/icons-material/EditTwoTone';
import InfoTwoToneIcon from '@mui/icons-material/InfoTwoTone';
import AddCircleOutlineTwoToneIcon from '@mui/icons-material/AddCircleOutlineTwoTone';

import MainCard from 'ui-component/cards/MainCard';
import ExpandableRadio from 'ui-component/general/ExpandableRadio';
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';

import axiosServices from 'utils/axios';
import envRef from 'environment';
import { MenuFlatOptions, defaultErrorMessage, statusActiveFlatOptions, statusYesNoFlatOptions } from 'constant/general';
import { gridSpacing } from 'store/constant';

import { Formik, FormikValues } from 'formik';
import * as Yup from 'yup';
import PrefixRadio from 'ui-component/general/PrefixRadio';
import AnimateButton from 'ui-component/extended/AnimateButton';
import ExpandableSelect from 'ui-component/general/ExpandableSelect';
import GeneralDialog from 'ui-component/general/GeneralDialog';
import AntdPicker from 'ui-component/icon-picker/AntdPicker';

const Menu = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();

    const [menuList, setMenuList] = React.useState<MenuListData[]>();
    const [triggerExpandAll, setTriggerExpandAll] = React.useState<boolean>(false);
    const [selectedValue, setSelectedValue] = React.useState<number | undefined>(undefined);
    const [isLoading, setIsLoading] = React.useState<boolean>(false);
    const [isDeleteModalOpen, setIsDeleteModalOpen] = React.useState<boolean>(false);

    const [mode, setMode] = React.useState<'add' | 'edit'>('edit');
    const scriptedRef = useScriptRef();

    const FormikInitialValuesTemplate: FormikValues = {
        id: 0,
        pid: 0,
        title: '',
        name: '',
        path: '',
        icon: '',
        type: 1,
        redirect: '',
        permissions: '',
        permissionName: '',
        component: '',
        alwaysShow: 0,
        activeMenu: '',
        isRoot: 0,
        isFrame: 0,
        frameSrc: '',
        keepAlive: 1,
        hidden: 0,
        affix: 0,
        level: 0,
        tree: '',
        sort: 10,
        remark: '',
        status: 1,
        createdAt: '',
        updatedAt: '',
        submit: null
    };
    const [formikInitialValues, setFormikInitialValues] = React.useState<FormikValues>(FormikInitialValuesTemplate);

    React.useEffect(() => {
        getOptions();
        setFormikInitialValues(FormikInitialValuesTemplate);
    }, []);

    async function getOptions() {
        // Get dept list
        await axiosServices
            .get(`${envRef?.API_URL}admin/menu/list`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setMenuList(response.data.data.list);
                }
            })
            .catch(function (error) {
                dispatch(
                    openSnackbar({
                        open: true,
                        message: error?.message || intl.formatMessage({ id: defaultErrorMessage }),
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
            });
    }

    function handleAdd(pid?: number, type?: number): void {
        setSelectedValue(pid || 0);
        setMode('add');
        if (pid !== undefined && type !== undefined) {
            setFormikInitialValues({ ...FormikInitialValuesTemplate, pid: pid, type: type });
        } else {
            setFormikInitialValues(FormikInitialValuesTemplate);
        }
    }

    async function handleEdit(id: number) {
        if (id !== undefined && id !== 0) {
            try {
                setIsLoading(true);
                setMode('edit');
                await axiosServices
                    .get(`${envRef?.API_URL}admin/menu/view?id=${id}`, { headers: {} })
                    .then(function (response) {
                        if (response?.data?.code === 0) {
                            setFormikInitialValues({
                                ...formikInitialValues,
                                id: response.data.data.view.id,
                                pid: response.data.data.view.pid,
                                title: response.data.data.view.title,
                                name: response.data.data.view.name,
                                path: response.data.data.view.path,
                                icon: response.data.data.view.icon,
                                type: response.data.data.view.type,
                                redirect: response.data.data.view.redirect,
                                permissions: response.data.data.view.permissions,
                                permissionName: response.data.data.view.permissionName,
                                component: response.data.data.view.component,
                                alwaysShow: response.data.data.view.alwaysShow,
                                activeMenu: response.data.data.view.activeMenu,
                                isRoot: response.data.data.view.isRoot,
                                isFrame: response.data.data.view.isFrame,
                                frameSrc: response.data.data.view.frameSrc,
                                keepAlive: response.data.data.view.keepAlive,
                                hidden: response.data.data.view.hidden,
                                affix: response.data.data.view.affix,
                                level: response.data.data.view.level,
                                tree: response.data.data.view.tree,
                                sort: response.data.data.view.sort,
                                remark: response.data.data.view.remark,
                                status: response.data.data.view.status,
                                createdAt: response.data.data.view.createdAt,
                                updatedAt: response.data.data.view.updatedAt,
                                submit: null
                            });
                        }
                    })
                    .catch(function (error) {
                        dispatch(
                            openSnackbar({
                                open: true,
                                message: error?.message || intl.formatMessage({ id: defaultErrorMessage }),
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
                    });
            } catch (e) {
                console.error(e);
            } finally {
                setIsLoading(false);
            }
        }
    }

    function handleDeleteModal() {
        setIsDeleteModalOpen(!isDeleteModalOpen);
    }

    async function handleDelete(id: number) {
        try {
            setIsLoading(true);
            await axiosServices
                .post(
                    `${envRef?.API_URL}admin/menu/delete`,
                    {
                        id: id
                    },
                    { headers: {} }
                )
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        dispatch(
                            openSnackbar({
                                open: true,
                                message: response?.data?.message,
                                variant: 'alert',
                                alert: {
                                    color: 'success'
                                },
                                close: false
                            })
                        );
                        setMenuList([]);
                        getOptions();
                    }
                })
                .catch(function (error) {
                    dispatch(
                        openSnackbar({
                            open: true,
                            message: error?.message || intl.formatMessage({ id: defaultErrorMessage }),
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
                });
        } catch (e) {
            console.error(e);
        } finally {
            handleDeleteModal();
            setIsLoading(false);
        }
    }

    return (
        <>
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <MainCard
                        title={intl.formatMessage({ id: 'menu-perm.menu-perm-title' })}
                        content={true}
                        children={intl.formatMessage({ id: 'menu-perm.menu-perm-content' })}
                    ></MainCard>
                </Grid>
            </Grid>
            <Grid container spacing={gridSpacing} marginTop="unset">
                <Grid item xs={12} sm={6} md={5}>
                    <MainCard>
                        <Grid container style={{ flexDirection: 'column', height: '100%', padding: '24px' }} spacing={gridSpacing}>
                            {menuList === undefined && <SkeletonLoader colsWidth={[12]} />}
                            {menuList !== undefined && (
                                <Grid container spacing={gridSpacing}>
                                    <Grid item xs={12} sm={4}>
                                        <Button
                                            fullWidth
                                            color="info"
                                            variant="outlined"
                                            sx={{ alignSelf: 'end' }}
                                            onClick={(e) => handleAdd()}
                                        >
                                            {intl.formatMessage({ id: 'menu-perm.add-menu' })}
                                        </Button>
                                    </Grid>
                                    <Grid item xs={12} sm={4}>
                                        <Button
                                            disabled={selectedValue === undefined || selectedValue === 0}
                                            fullWidth
                                            color="info"
                                            variant="outlined"
                                            sx={{ alignSelf: 'end' }}
                                            onClick={(e) => {
                                                handleAdd(selectedValue, 2);
                                            }}
                                        >
                                            {intl.formatMessage({ id: 'menu-perm.add-sub-menu' })}
                                        </Button>
                                    </Grid>
                                    <Grid item xs={12} sm={4}>
                                        <Button
                                            fullWidth
                                            color="info"
                                            variant="outlined"
                                            onClick={() => setTriggerExpandAll(!triggerExpandAll)}
                                            sx={{ alignSelf: 'end' }}
                                        >
                                            {triggerExpandAll
                                                ? intl.formatMessage({ id: 'general.collapse-all' })
                                                : intl.formatMessage({ id: 'general.expand-all' })}
                                        </Button>
                                    </Grid>
                                    <Grid item xs={12} sx={{ overflowY: 'scroll', height: 'calc(30vmax - 24px)', marginTop: '1rem' }}>
                                        <ExpandableRadio
                                            onSelectChange={(e) => {
                                                setSelectedValue(e[0]);
                                                handleEdit(e[0]);
                                            }}
                                            id="menuIds"
                                            options={menuList}
                                            initialValue={[selectedValue || 0]}
                                            valueFieldName="key"
                                            displayFieldName="label"
                                            triggerExpandAll={triggerExpandAll}
                                        />
                                    </Grid>
                                </Grid>
                            )}
                        </Grid>
                    </MainCard>
                </Grid>
                <Grid item xs={12} sm={6} md={7}>
                    <MainCard
                        style={{ display: 'flex', flexDirection: 'column', height: '100%' }}
                        title={
                            <Grid container alignItems="center">
                                {mode === 'edit' && (
                                    <>
                                        <EditTwoToneIcon />
                                        {intl.formatMessage({ id: 'menu-perm.edit-menu' })}
                                    </>
                                )}
                                {mode === 'add' && (
                                    <>
                                        <AddCircleOutlineTwoToneIcon />
                                        {intl.formatMessage({ id: 'menu-perm.add-menu' })}
                                    </>
                                )}
                            </Grid>
                        }
                    >
                        {isLoading ? (
                            <SkeletonLoader colsWidth={[6, 6]} />
                        ) : selectedValue === undefined ? (
                            <Grid container style={{ flex: 1 }} spacing={gridSpacing} alignItems="center">
                                <Grid item xs={12} textAlign="center">
                                    <InfoTwoToneIcon sx={{ fontSize: '5rem' }} color="info" />
                                    <Typography variant="h1">{intl.formatMessage({ id: 'general.hint' })}</Typography>
                                    <Typography variant="h4">{intl.formatMessage({ id: 'menu-perm.hint-content' })}</Typography>
                                </Grid>
                            </Grid>
                        ) : (
                            <Grid container style={{ flex: 1 }} spacing={gridSpacing} alignItems="center">
                                <Grid item xs={12}>
                                    <Formik
                                        enableReinitialize={true}
                                        initialValues={formikInitialValues}
                                        validationSchema={Yup.object().shape({})}
                                        onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                                            try {
                                                await axiosServices
                                                    .post(
                                                        `${envRef?.API_URL}admin/menu/edit`,
                                                        {
                                                            id: values.id,
                                                            pid: values.pid,
                                                            title: values.title,
                                                            name: values.name,
                                                            path: values.path,
                                                            icon: values.icon,
                                                            type: values.type,
                                                            redirect: values.redirect,
                                                            permissions: values.permissions,
                                                            permissionName: values.permissionName,
                                                            component: values.component,
                                                            alwaysShow: values.alwaysShow,
                                                            activeMenu: values.activeMenu,
                                                            isRoot: values.isRoot,
                                                            isFrame: values.isFrame,
                                                            frameSrc: values.frameSrc,
                                                            keepAlive: values.keepAlive,
                                                            hidden: values.hidden,
                                                            affix: values.affix,
                                                            level: values.level,
                                                            tree: values.tree,
                                                            sort: values.sort,
                                                            remark: values.remark,
                                                            status: values.status,
                                                            createdAt: values.createdAt,
                                                            updatedAt: values.updatedAt
                                                        },
                                                        { headers: {} }
                                                    )
                                                    .then(function (response) {
                                                        if (response?.data?.code === 0) {
                                                            dispatch(
                                                                openSnackbar({
                                                                    open: true,
                                                                    message: response?.data?.message,
                                                                    variant: 'alert',
                                                                    alert: {
                                                                        color: 'success'
                                                                    },
                                                                    close: false
                                                                })
                                                            );
                                                            setMenuList([]);
                                                            getOptions();
                                                        } else {
                                                            dispatch(
                                                                openSnackbar({
                                                                    open: true,
                                                                    message: response?.data?.message || defaultErrorMessage,
                                                                    variant: 'alert',
                                                                    alert: {
                                                                        color: 'error',
                                                                        severity: 'error'
                                                                    },
                                                                    close: false,
                                                                    anchorOrigin: {
                                                                        vertical: 'top',
                                                                        horizontal: 'center'
                                                                    }
                                                                })
                                                            );
                                                        }
                                                    })
                                                    .catch(function (error) {
                                                        dispatch(
                                                            openSnackbar({
                                                                open: true,
                                                                message: error?.message || defaultErrorMessage,
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
                                                    });
                                            } catch (err: any) {
                                                if (scriptedRef.current) {
                                                    setStatus({ success: false });
                                                    setErrors({ submit: err.message });
                                                    setSubmitting(false);
                                                }
                                            }
                                        }}
                                    >
                                        {({
                                            errors,
                                            handleBlur,
                                            handleChange,
                                            handleSubmit,
                                            isSubmitting,
                                            touched,
                                            values,
                                            setFieldValue,
                                            setFieldTouched,
                                            resetForm
                                        }) => (
                                            <form noValidate onSubmit={handleSubmit}>
                                                <Divider sx={{ paddingY: '1rem' }} textAlign="left">
                                                    {intl.formatMessage({ id: 'general.general-setting' })}
                                                </Divider>
                                                <Grid container alignItems="center" spacing={gridSpacing}>
                                                    <Grid item xs={12} sm={6}>
                                                        <PrefixRadio
                                                            options={MenuFlatOptions}
                                                            label={intl.formatMessage({ id: 'general.type' })}
                                                            id="type"
                                                            onSelectChange={(e) => setFieldValue('type', e)}
                                                            initialValue={values.type}
                                                        />
                                                    </Grid>
                                                    <Grid item className="expandableSelectContainer" xs={12} sm={6}>
                                                        <FormControl
                                                            fullWidth
                                                            error={Boolean(touched.roleId && errors.roleId)}
                                                            sx={{ ...theme.typography.customInput }}
                                                        >
                                                            <ExpandableSelect
                                                                label={intl.formatMessage({ id: 'user.role-name' })}
                                                                id="pid"
                                                                filterStatus={false}
                                                                options={[
                                                                    {
                                                                        id: 0,
                                                                        name: 'root',
                                                                        key: 0,
                                                                        dataScope: 1,
                                                                        customDept: {},
                                                                        pid: 0,
                                                                        level: 0,
                                                                        tree: '',
                                                                        remark: '',
                                                                        sort: 0,
                                                                        status: 1,
                                                                        createdAt: '',
                                                                        updatedAt: '',
                                                                        orgAdmin: 0,
                                                                        label: intl.formatMessage({ id: 'general.root' }),
                                                                        value: 0,
                                                                        children: null
                                                                    },
                                                                    ...menuList!
                                                                ]}
                                                                onSelectChange={(e) => {
                                                                    setFieldValue('pid', e[0]);
                                                                }}
                                                                initialValue={values.pid === undefined ? [0] : [values.pid]}
                                                                valueFieldName="key"
                                                                displayFieldName="label"
                                                                triggerExpandAll={false}
                                                                multiSelect={false}
                                                            />
                                                            {touched.roleId && errors.roleId && (
                                                                <FormHelperText error id="standard-weight-helper-text-role-add">
                                                                    {errors.roleId.toString()}
                                                                </FormHelperText>
                                                            )}
                                                        </FormControl>
                                                    </Grid>
                                                </Grid>
                                                <Grid container alignItems="center" spacing={gridSpacing}>
                                                    <Grid item xs={12} sm={6}>
                                                        <FormControl
                                                            fullWidth
                                                            error={Boolean(touched.title && errors.title)}
                                                            sx={{ ...theme.typography.customInput }}
                                                        >
                                                            <InputLabel htmlFor="outlined-adornment-directory-title-add">
                                                                {values.type === 1 &&
                                                                    intl.formatMessage({ id: 'menu-perm.directory-name' })}
                                                                {values.type === 2 && intl.formatMessage({ id: 'menu-perm.menu-name' })}
                                                                {values.type === 3 && intl.formatMessage({ id: 'menu-perm.button-name' })}
                                                            </InputLabel>
                                                            <OutlinedInput
                                                                id="outlined-adornment-directory-title-add"
                                                                type="text"
                                                                value={values.title}
                                                                name="title"
                                                                onBlur={handleBlur}
                                                                onChange={handleChange}
                                                                inputProps={{}}
                                                            />
                                                            {touched.title && errors.title && (
                                                                <FormHelperText error id="standard-weight-helper-text-directory-title-add">
                                                                    {errors.title.toString()}
                                                                </FormHelperText>
                                                            )}
                                                        </FormControl>
                                                    </Grid>

                                                    {/* TODO: ICON FIELD && tooltip ?*/}
                                                    {values.type !== 3 && (
                                                        <Grid item xs={12} sm={6}>
                                                            <FormControl
                                                                fullWidth
                                                                error={Boolean(touched.icon && errors.icon)}
                                                                sx={{ ...theme.typography.customInput }}
                                                            >
                                                                <AntdPicker
                                                                    initialValue={values.icon}
                                                                    onSelectChange={(e) => setFieldValue('icon', e)}
                                                                    label={intl.formatMessage({ id: 'menu-perm.menu-icon' })}
                                                                />
                                                                {touched.icon && errors.icon && (
                                                                    <FormHelperText error id="standard-weight-helper-text-menu-icon-add">
                                                                        {errors.icon.toString()}
                                                                    </FormHelperText>
                                                                )}
                                                            </FormControl>
                                                        </Grid>
                                                    )}
                                                </Grid>

                                                <Grid container alignItems="center" spacing={gridSpacing}>
                                                    {values.type !== 3 && (
                                                        <Grid item xs={12} sm={6}>
                                                            <FormControl
                                                                fullWidth
                                                                error={Boolean(touched.path && errors.path)}
                                                                sx={{ ...theme.typography.customInput }}
                                                            >
                                                                <InputLabel htmlFor="outlined-adornment-route-address-add">
                                                                    {intl.formatMessage({ id: 'menu-perm.route-address' })}
                                                                </InputLabel>
                                                                <OutlinedInput
                                                                    id="outlined-adornment-route-address-add"
                                                                    type="text"
                                                                    value={values.path}
                                                                    name="path"
                                                                    onBlur={handleBlur}
                                                                    onChange={handleChange}
                                                                    inputProps={{}}
                                                                />
                                                                {touched.path && errors.path && (
                                                                    <FormHelperText
                                                                        error
                                                                        id="standard-weight-helper-text-route-address-add"
                                                                    >
                                                                        {errors.path.toString()}
                                                                    </FormHelperText>
                                                                )}
                                                            </FormControl>
                                                        </Grid>
                                                    )}
                                                    <Grid item xs={12} sm={6}>
                                                        <FormControl
                                                            fullWidth
                                                            error={Boolean(touched.name && errors.name)}
                                                            sx={{ ...theme.typography.customInput }}
                                                        >
                                                            <InputLabel htmlFor="outlined-adornment-route-nickname-add">
                                                                {intl.formatMessage({ id: 'menu-perm.route-nickname' })}
                                                            </InputLabel>
                                                            <OutlinedInput
                                                                id="outlined-adornment-route-nickname-add"
                                                                type="text"
                                                                value={values.name}
                                                                name="name"
                                                                onBlur={handleBlur}
                                                                onChange={handleChange}
                                                                inputProps={{}}
                                                            />
                                                            {touched.name && errors.name && (
                                                                <FormHelperText error id="standard-weight-helper-text-route-nickname-add">
                                                                    {errors.name.toString()}
                                                                </FormHelperText>
                                                            )}
                                                        </FormControl>
                                                    </Grid>
                                                </Grid>

                                                {values.type !== 3 && (
                                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                                        <Grid item xs={12} sm={6}>
                                                            <FormControl
                                                                fullWidth
                                                                error={Boolean(touched.component && errors.component)}
                                                                sx={{ ...theme.typography.customInput }}
                                                            >
                                                                <InputLabel htmlFor="outlined-adornment-component-path-add">
                                                                    {intl.formatMessage({ id: 'menu-perm.component-path' })}
                                                                </InputLabel>
                                                                <OutlinedInput
                                                                    id="outlined-adornment-component-path-add"
                                                                    type="text"
                                                                    value={values.component}
                                                                    name="component"
                                                                    onBlur={handleBlur}
                                                                    onChange={handleChange}
                                                                    inputProps={{}}
                                                                />
                                                                {touched.component && errors.component && (
                                                                    <FormHelperText
                                                                        error
                                                                        id="standard-weight-helper-text-component-path-add"
                                                                    >
                                                                        {errors.component.toString()}
                                                                    </FormHelperText>
                                                                )}
                                                            </FormControl>
                                                        </Grid>
                                                        {values.type === 1 && (
                                                            <Grid item xs={12} sm={6}>
                                                                <FormControl
                                                                    fullWidth
                                                                    error={Boolean(touched.redirect && errors.redirect)}
                                                                    sx={{ ...theme.typography.customInput }}
                                                                >
                                                                    <InputLabel htmlFor="outlined-adornment-default-redirect-add">
                                                                        {intl.formatMessage({ id: 'menu-perm.default-redirect' })}
                                                                    </InputLabel>
                                                                    <OutlinedInput
                                                                        id="outlined-adornment-default-redirect-add"
                                                                        type="text"
                                                                        value={values.redirect}
                                                                        name="redirect"
                                                                        onBlur={handleBlur}
                                                                        onChange={handleChange}
                                                                        inputProps={{}}
                                                                    />
                                                                    {touched.redirect && errors.redirect && (
                                                                        <FormHelperText
                                                                            error
                                                                            id="standard-weight-helper-text-default-redirect-add"
                                                                        >
                                                                            {errors.redirect.toString()}
                                                                        </FormHelperText>
                                                                    )}
                                                                </FormControl>
                                                            </Grid>
                                                        )}
                                                    </Grid>
                                                )}

                                                <Divider sx={{ paddingY: '1rem' }} textAlign="left">
                                                    {intl.formatMessage({ id: 'general.function-setting' })}
                                                </Divider>

                                                <Grid container alignItems="center" spacing={gridSpacing}>
                                                    <Grid item xs={12}>
                                                        <FormControl
                                                            fullWidth
                                                            error={Boolean(touched.permissions && errors.permissions)}
                                                            sx={{ ...theme.typography.customInput }}
                                                        >
                                                            <InputLabel htmlFor="outlined-adornment-permissions-add">
                                                                {intl.formatMessage({ id: 'general.permissions' })}
                                                            </InputLabel>
                                                            <OutlinedInput
                                                                id="outlined-adornment-permissions-add"
                                                                type="text"
                                                                value={values.permissions}
                                                                name="permissions"
                                                                onBlur={handleBlur}
                                                                onChange={handleChange}
                                                                inputProps={{}}
                                                            />
                                                            {touched.permissions && errors.permissions && (
                                                                <FormHelperText error id="standard-weight-helper-text-permissions-add">
                                                                    {errors.permissions.toString()}
                                                                </FormHelperText>
                                                            )}
                                                        </FormControl>
                                                    </Grid>
                                                </Grid>

                                                <Grid container alignItems="center" spacing={gridSpacing}>
                                                    {values.type !== 3 && (
                                                        <Grid item xs={12} sm={6}>
                                                            <FormControl
                                                                fullWidth
                                                                error={Boolean(touched.activeMenu && errors.activeMenu)}
                                                                sx={{ ...theme.typography.customInput }}
                                                            >
                                                                <InputLabel htmlFor="outlined-adornment-highlight-route-add">
                                                                    {intl.formatMessage({ id: 'menu-perm.highlight-route' })}
                                                                </InputLabel>
                                                                <OutlinedInput
                                                                    id="outlined-adornment-highlight-route-add"
                                                                    type="text"
                                                                    value={values.activeMenu}
                                                                    name="activeMenu"
                                                                    onBlur={handleBlur}
                                                                    onChange={handleChange}
                                                                    inputProps={{}}
                                                                />
                                                                {touched.activeMenu && errors.activeMenu && (
                                                                    <FormHelperText
                                                                        error
                                                                        id="standard-weight-helper-text-highlight-route-add"
                                                                    >
                                                                        {errors.activeMenu.toString()}
                                                                    </FormHelperText>
                                                                )}
                                                            </FormControl>
                                                        </Grid>
                                                    )}
                                                    <Grid item xs={12} sm={6}>
                                                        <FormControl
                                                            fullWidth
                                                            error={Boolean(touched.sort && errors.sort)}
                                                            sx={{ ...theme.typography.customInput }}
                                                        >
                                                            <InputLabel htmlFor="outlined-adornment-menu-order-add">
                                                                {intl.formatMessage({ id: 'menu-perm.menu-order' })}
                                                            </InputLabel>
                                                            <OutlinedInput
                                                                id="outlined-adornment-menu-order-add"
                                                                type="number"
                                                                value={values.sort}
                                                                name="sort"
                                                                onBlur={handleBlur}
                                                                onChange={handleChange}
                                                                inputProps={{}}
                                                            />
                                                            {touched.sort && errors.sort && (
                                                                <FormHelperText error id="standard-weight-helper-text-menu-order-add">
                                                                    {errors.sort.toString()}
                                                                </FormHelperText>
                                                            )}
                                                        </FormControl>
                                                    </Grid>
                                                </Grid>

                                                <Grid container alignItems="center" spacing={gridSpacing}>
                                                    {values.type !== 3 && (
                                                        <Grid item sm={12} md={6} lg={4}>
                                                            <PrefixRadio
                                                                valueFieldName="value"
                                                                options={statusActiveFlatOptions}
                                                                label={intl.formatMessage({ id: 'menu-perm.root-route' })}
                                                                id="isRoot"
                                                                onSelectChange={(e) => setFieldValue('isRoot', e)}
                                                                initialValue={values.isRoot}
                                                            />
                                                        </Grid>
                                                    )}
                                                    {values.type !== 3 && (
                                                        <Grid item sm={12} md={6} lg={4}>
                                                            <PrefixRadio
                                                                valueFieldName="value"
                                                                options={statusActiveFlatOptions}
                                                                label={intl.formatMessage({ id: 'menu-perm.persist-page' })}
                                                                id="affix"
                                                                onSelectChange={(e) => setFieldValue('affix', e)}
                                                                initialValue={values.affix}
                                                            />
                                                        </Grid>
                                                    )}
                                                    {values.type !== 3 && (
                                                        <Grid item sm={12} md={6} lg={4}>
                                                            <PrefixRadio
                                                                valueFieldName="value"
                                                                options={statusActiveFlatOptions}
                                                                label={intl.formatMessage({ id: 'menu-perm.simplify-route' })}
                                                                id="alwaysShow"
                                                                onSelectChange={(e) => setFieldValue('alwaysShow', e)}
                                                                initialValue={values.alwaysShow}
                                                            />
                                                        </Grid>
                                                    )}
                                                    {values.type !== 3 && (
                                                        <Grid item sm={12} md={6} lg={4}>
                                                            <PrefixRadio
                                                                valueFieldName="value"
                                                                options={statusActiveFlatOptions}
                                                                label={intl.formatMessage({ id: 'menu-perm.cache-route' })}
                                                                id="keepAlive"
                                                                onSelectChange={(e) => setFieldValue('keepAlive', e)}
                                                                initialValue={values.keepAlive}
                                                            />
                                                        </Grid>
                                                    )}
                                                    {values.type !== 3 && (
                                                        <Grid item sm={12} md={6} lg={4}>
                                                            <PrefixRadio
                                                                valueFieldName="value"
                                                                options={statusYesNoFlatOptions}
                                                                label={intl.formatMessage({ id: 'general.hidden' })}
                                                                id="hidden"
                                                                onSelectChange={(e) => setFieldValue('hidden', e)}
                                                                initialValue={values.hidden}
                                                            />
                                                        </Grid>
                                                    )}
                                                    <Grid item sm={12} md={6} lg={4}>
                                                        <PrefixRadio
                                                            options={statusActiveFlatOptions}
                                                            label={intl.formatMessage({ id: 'general.status' })}
                                                            id="status"
                                                            valueFieldName="value"
                                                            onSelectChange={(e) => setFieldValue('status', e)}
                                                            initialValue={values.status}
                                                        />
                                                    </Grid>

                                                    {values.type !== 3 && (
                                                        <Grid item sm={12} md={6} lg={4}>
                                                            <PrefixRadio
                                                                options={statusActiveFlatOptions}
                                                                valueFieldName="value"
                                                                label={intl.formatMessage({ id: 'general.external-link' })}
                                                                id="isFrame"
                                                                onSelectChange={(e) => setFieldValue('isFrame', e)}
                                                                initialValue={values.isFrame}
                                                            />
                                                        </Grid>
                                                    )}
                                                    {values.type !== 3 && values.isFrame === 1 && (
                                                        <Grid item sm={12} md={6} lg={4}>
                                                            <FormControl
                                                                fullWidth
                                                                error={Boolean(touched.frameSrc && errors.frameSrc)}
                                                                sx={{ ...theme.typography.customInput }}
                                                            >
                                                                <InputLabel htmlFor="outlined-adornment-external-src-add">
                                                                    {intl.formatMessage({ id: 'general.external-address' })}
                                                                </InputLabel>
                                                                <OutlinedInput
                                                                    id="outlined-adornment-external-src-add"
                                                                    type="text"
                                                                    value={values.frameSrc}
                                                                    name="frameSrc"
                                                                    onBlur={handleBlur}
                                                                    onChange={handleChange}
                                                                    inputProps={{}}
                                                                />
                                                                {touched.frameSrc && errors.frameSrc && (
                                                                    <FormHelperText error id="standard-weight-helper-text-external-src-add">
                                                                        {errors.frameSrc.toString()}
                                                                    </FormHelperText>
                                                                )}
                                                            </FormControl>
                                                        </Grid>
                                                    )}
                                                </Grid>

                                                {errors.submit && (
                                                    <Box sx={{ mt: 3 }}>
                                                        <FormHelperText error>{errors.submit.toString()}</FormHelperText>
                                                    </Box>
                                                )}
                                                <Grid container justifyContent="flex-end" spacing={gridSpacing}>
                                                    <Grid item xs={12} sm={4} md={3}>
                                                        <Box sx={{ mt: 2 }}>
                                                            <AnimateButton>
                                                                <Button
                                                                    disableElevation
                                                                    disabled={isSubmitting}
                                                                    fullWidth
                                                                    size="large"
                                                                    variant="outlined"
                                                                    color="inherit"
                                                                    onClick={() => handleDeleteModal()}
                                                                >
                                                                    {intl.formatMessage({ id: 'general.delete' })}
                                                                </Button>
                                                            </AnimateButton>
                                                        </Box>
                                                    </Grid>
                                                    <Grid item xs={12} sm={4} md={3}>
                                                        <Box sx={{ mt: 2 }}>
                                                            <AnimateButton>
                                                                <Button
                                                                    disableElevation
                                                                    disabled={isSubmitting}
                                                                    fullWidth
                                                                    size="large"
                                                                    variant="outlined"
                                                                    color="secondary"
                                                                    onClick={() => resetForm()}
                                                                >
                                                                    {intl.formatMessage({ id: 'general.reset' })}
                                                                </Button>
                                                            </AnimateButton>
                                                        </Box>
                                                    </Grid>
                                                    <Grid item xs={12} sm={4} md={3}>
                                                        <Box sx={{ mt: 2 }}>
                                                            <AnimateButton>
                                                                <Button
                                                                    disableElevation
                                                                    disabled={isSubmitting}
                                                                    fullWidth
                                                                    size="large"
                                                                    type="submit"
                                                                    variant="contained"
                                                                    color="secondary"
                                                                >
                                                                    {intl.formatMessage({ id: 'general.confirm' })}
                                                                </Button>
                                                            </AnimateButton>
                                                        </Box>
                                                    </Grid>
                                                </Grid>
                                            </form>
                                        )}
                                    </Formik>
                                </Grid>
                            </Grid>
                        )}
                        <GeneralDialog
                            confirmFunction={() => handleDelete(selectedValue!)}
                            isOpen={isDeleteModalOpen}
                            setIsOpen={() => setIsDeleteModalOpen(!isDeleteModalOpen)}
                            id="delete-confirm-modal"
                            type="warning"
                            content={intl.formatMessage({ id: 'general.confirm-delete-content' })}
                        />
                    </MainCard>
                </Grid>
            </Grid>
        </>
    );
};

export default Menu;
