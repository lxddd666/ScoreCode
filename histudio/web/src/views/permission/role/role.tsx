import * as React from 'react';

// locale
import { FormattedMessage, useIntl } from 'react-intl';

// material-ui
import { useTheme } from '@mui/material/styles';
import {
    Box,
    Button,
    CardContent,
    Dialog,
    FormControl,
    FormHelperText,
    Grid,
    IconButton,
    InputLabel,
    OutlinedInput,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Tooltip,
    Typography
} from '@mui/material';
import FolderOffTwoToneIcon from '@mui/icons-material/FolderOffTwoTone';
import CancelTwoToneIcon from '@mui/icons-material/CancelTwoTone';
import HelpTwoToneIcon from '@mui/icons-material/HelpTwoTone';

// project imports
import { ResponseList } from 'types/response';
import { DataScopeListData, MenuListData, RoleListData } from 'types/user';
import { HeadCell } from 'types';
import { FieldToShow } from 'types/general';
import useScriptRef from 'hooks/useScriptRef';

// ui-components
import AnimateButton from 'ui-component/extended/AnimateButton';
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';
import MainCard from 'ui-component/cards/MainCard';
import AddButton from 'ui-component/searchbar/AddButton';
import ExpandableSelect from 'ui-component/general/ExpandableSelect';
import GeneralDialog from 'ui-component/general/GeneralDialog';
import ExpandableTableRow from 'ui-component/general/ExpandableTableRow';
import PrefixRadio from 'ui-component/general/PrefixRadio';
import ExpandableRadio from 'ui-component/general/ExpandableRadio';

// API
import axiosServices from 'utils/axios';
import envRef from 'environment';

// third party
import * as Yup from 'yup';
import { Formik, FormikValues } from 'formik';
import { isEqual } from 'lodash';

// project imports
import { statusFlatOptions } from 'constant/general';
import { gridSpacing } from 'store/constant';
import { openSnackbar } from 'store/slices/snackbar';
import { useDispatch, useSelector } from 'store';
import { getRoleList } from 'store/slices/user';

// table header options
const headCells: HeadCell[] = [
    {
        id: 'label',
        numeric: false,
        label: 'role.role-name',
        align: 'left'
    },
    {
        id: 'key',
        numeric: false,
        label: 'role.role-code',
        align: 'center'
    },
    {
        id: 'pid',
        numeric: false,
        label: 'role.default-role',
        align: 'center'
    },
    {
        id: 'sort',
        numeric: false,
        label: 'general.sort',
        align: 'center'
    },
    {
        id: 'remarks',
        numeric: true,
        label: 'general.remarks',
        align: 'center'
    },
    {
        id: 'status',
        numeric: false,
        label: 'general.status',
        align: 'center'
    },
    {
        id: 'createdAt',
        numeric: false,
        label: 'general.created-date',
        align: 'center'
    }
];

const fieldsToShow: FieldToShow[] = [
    {
        fieldName: 'key',
        fieldType: ''
    },
    {
        fieldName: 'pid',
        fieldType: 'isDefault'
    },
    {
        fieldName: 'sort',
        fieldType: ''
    },
    {
        fieldName: 'remark',
        fieldType: ''
    },
    {
        fieldName: 'status',
        fieldType: 'status'
    },
    {
        fieldName: 'createdAt',
        fieldType: ''
    }
];

// ==============================|| DEPARTMENT LIST ||============================== //

const Role = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);

    const [selectedRole, setSelectedRole] = React.useState<RoleListData | null>();
    const [rows, setRows] = React.useState<RoleListData[]>();
    const [addModalOptions, setAddModalOptions] = React.useState<RoleListData[]>();
    const [res, setRes] = React.useState<ResponseList<RoleListData>>();
    const [menuList, setMenuList] = React.useState<ResponseList<MenuListData>>();
    const [dataScopeList, setDataScopeList] = React.useState<ResponseList<DataScopeListData>>();
    const { roleList } = useSelector((state) => state.user);
    const [isDataListEmpty, setIsDataListEmpty] = React.useState<boolean>(true);
    const scriptedRef = useScriptRef();
    const FormikInitialValuesTemplate: FormikValues = {
        id: 0,
        customDept: [],
        dataScope: 1,
        key: '',
        level: 1,
        name: '',
        pid: 0,
        remark: '',
        sort: 0,
        status: 1,
        tree: '',
        submit: null
    };
    const [formikInitialValues, setFormikInitialValues] = React.useState<FormikValues>(FormikInitialValuesTemplate);
    const [selectedPerms, setSelectedPerms] = React.useState<number[]>([]);
    const [triggerCheckAll, setTriggerCheckAll] = React.useState<boolean>(false);
    const [triggerExpandAll, setTriggerExpandAll] = React.useState<boolean>(false);
    const [selectedDataScope, setSelectedDataScope] = React.useState<number>();
    React.useEffect(() => {
        fetchData();
    }, [dispatch]);

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
        await axiosServices
            .get(`${envRef?.API_URL}admin/role/dataScope/select`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setDataScopeList(response.data.data.list);
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
    }

    React.useEffect(() => {
        getOptions();
    }, []);

    const fetchData = async () => {
        try {
            setLoading(true);
            setIsDataListEmpty(true);
            setRows(undefined);
            await dispatch(getRoleList(['pageSize=100', 'page=1'].join('&')));
        } catch (error) {
            dispatch(
                openSnackbar({
                    open: true,
                    message: error || defaultErrorMessage,
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
        } finally {
            setLoading(false);
        }
    };

    React.useEffect(() => {
        setRes(roleList!);
    }, [roleList]);
    React.useEffect(() => {
        setRows(res?.data?.list ? res.data.list : []);
    }, [res]);
    React.useEffect(() => {
        setIsDataListEmpty(isEqual(rows, []));
        // const appendedModalOptions: RoleListData[] = [
        //     {
        //         id: 0,
        //         orgId: 0,
        //         pid: 0,
        //         name: intl.formatMessage({ id: 'department.top-department' }),
        //         code: 'top',
        //         type: 'dept',
        //         leader: 'admin',
        //         phone: '15888888888',
        //         email: '163@qq.com',
        //         level: 3,
        //         tree: '',
        //         sort: 1,
        //         status: 1,
        //         createdAt: '',
        //         updatedAt: '',
        //         label: intl.formatMessage({ id: 'department.top-department' }),
        //         value: 0,
        //         children: rows || null
        //     }
        // ];
        if (rows) {
            const appendedModalOptions: RoleListData[] = [
                {
                    id: 0,
                    name: intl.formatMessage({ id: 'role.top-role' }),
                    key: 'top',
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
                    orgAdmin: 2,
                    label: intl.formatMessage({ id: 'role.top-role' }),
                    value: 0,
                    children: null
                },
                ...rows
            ];
            setAddModalOptions(appendedModalOptions);
        } else setAddModalOptions(rows);
    }, [rows]);

    const [isAddModalOpen, setIsAddModalOpen] = React.useState<boolean>(false);

    const handleAddModal = (id?: number, type?: 'add' | 'edit') => {
        if (type === 'edit') {
            axiosServices
                .get(`${envRef?.API_URL}admin/role/view?id=${id}`, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        const selectedFormikValues: FormikValues = {
                            id: response.data.data.id,
                            name: response.data.data.name,
                            key: response.data.data.key,
                            dataScope: response.data.data.dataScope,
                            customDept: response.data.data.customDept,
                            pid: response.data.data.pid,
                            level: response.data.data.level,
                            tree: response.data.data.tree,
                            remark: response.data.data.remark,
                            sort: response.data.data.sort,
                            status: response.data.data.status,
                            createdAt: response.data.data.createdAt,
                            updatedAt: response.data.data.updatedAt,
                            orgAdmin: response.data.data.orgAdmin,
                            label: response.data.data.label,
                            value: response.data.data.value,
                            children: response.data.data.children,
                            submit: null
                        };
                        setFormikInitialValues(selectedFormikValues);
                    }
                    setIsAddModalOpen(!isAddModalOpen);
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
        } else {
            if (id) {
                const updatedFormikInitialValues: FormikValues = {
                    ...FormikInitialValuesTemplate,
                    pid: id
                };
                setFormikInitialValues(updatedFormikInitialValues);
            } else setFormikInitialValues(FormikInitialValuesTemplate);
            setIsAddModalOpen(!isAddModalOpen);
        }
    };

    const [isDeleteModalOpen, setIsDeleteModalOpen] = React.useState<boolean>(false);
    const [isMenuPermissionModalOpen, setIsMenuPermissionModalOpen] = React.useState<boolean>(false);
    const [isDataPermissionModalOpen, setIsDataPermissionModalOpen] = React.useState<boolean>(false);
    const [menuPermissionInitialValue, setMenuPermissionInitialValue] = React.useState<any>(null);
    const [dataScopeInitialValue, setDataScopeInitialValue] = React.useState<any>(null);
    const [selectedIdsToDelete, setSelectedIdsToDelete] = React.useState<number[]>([]);
    const handleDeleteModal = (id: number[]) => {
        setIsDeleteModalOpen(!isDeleteModalOpen);
        setSelectedIdsToDelete(id);
    };

    const handleDelete = (ids: number[]) => {
        axiosServices
            .post(`${envRef?.API_URL}admin/role/delete`, { id: ids[0] }, { headers: {} })
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
                    fetchData();
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
        setIsDeleteModalOpen(!isDeleteModalOpen);
    };

    async function handleMenuPermissionModal(selectedRole: RoleListData | null) {
        setSelectedRole(selectedRole);
        if (selectedRole) {
            await axiosServices
                .get(`${envRef?.API_URL}admin/role/getPermissions?id=${selectedRole.id}`, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        setMenuPermissionInitialValue(response.data.data.menuIds);
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
        }
        setIsMenuPermissionModalOpen(!isMenuPermissionModalOpen);
    }

    async function handleChangeMenuPermission() {
        await axiosServices
            .post(`${envRef?.API_URL}admin/role/updatePermissions`, { id: selectedRole!.id, menuIds: selectedPerms }, { headers: {} })
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
                    fetchData();
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
    }

    function handleDataPermissionModal(selectedRole: RoleListData | null) {
        setSelectedRole(selectedRole);
        setDataScopeInitialValue([selectedRole?.dataScope]);
        setIsDataPermissionModalOpen(!isDataPermissionModalOpen);
    }

    async function handleChangeDataScope() {
        await axiosServices
            .post(`${envRef?.API_URL}admin/role/dataScope/edit`, { id: selectedRole!.id, dataScope: selectedDataScope }, { headers: {} })
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
                    fetchData();
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
    }

    const defaultErrorMessage = intl.formatMessage({ id: 'auth-register.default-error' });

    function renderTable() {
        return (
            <TableContainer>
                <Table sx={{ minWidth: 750 }} aria-labelledby="tableTitle">
                    <TableHead>
                        <TableRow>
                            {headCells.map((headCell) => (
                                <TableCell
                                    size="medium"
                                    key={headCell.id}
                                    align={headCell.align}
                                    padding={headCell.disablePadding ? 'none' : 'normal'}
                                >
                                    {headCell.id === 'name' ? (
                                        <>
                                            {intl.formatMessage({ id: headCell.label })}
                                            <Tooltip title={intl.formatMessage({ id: 'general.department-tooltip' })} placement="top">
                                                <IconButton size="small">
                                                    <HelpTwoToneIcon />
                                                </IconButton>
                                            </Tooltip>
                                        </>
                                    ) : (
                                        intl.formatMessage({ id: headCell.label })
                                    )}
                                </TableCell>
                            ))}
                            <TableCell align="center" className="sticky" sx={{ pr: 3, right: '0' }}>
                                {intl.formatMessage({ id: 'general.control' })}
                            </TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {loading ? (
                            <TableRow>
                                <TableCell align="center" colSpan={headCells.length + 1}>
                                    <SkeletonLoader />
                                </TableCell>
                            </TableRow>
                        ) : isDataListEmpty ? (
                            <TableRow>
                                <TableCell align="center" colSpan={headCells.length + 1}>
                                    <FolderOffTwoToneIcon sx={{ verticalAlign: 'bottom' }} />
                                    {intl.formatMessage({ id: 'general.no-records' })}
                                </TableCell>
                            </TableRow>
                        ) : (
                            <ExpandableTableRow
                                data={rows || []}
                                firstRowDisableControl={true}
                                handleAddModal={handleAddModal}
                                handleDeleteModal={handleDeleteModal}
                                handleMenuPermissionModal={handleMenuPermissionModal}
                                handleDataPermissionModal={handleDataPermissionModal}
                                fieldsToShow={fieldsToShow}
                            />
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
        );
    }

    return (
        <MainCard title={<FormattedMessage id="role.role-management" />} content={false}>
            <CardContent>
                <Grid container alignItems="center" spacing={2}>
                    <Grid item xs={12} sm={4} md={3} sx={{ display: 'flex', justifyContent: 'flex-end', marginLeft: 'auto' }}>
                        <AddButton onClick={() => handleAddModal()} tooltipTitle={intl.formatMessage({ id: 'role.add-role' })} />
                    </Grid>
                </Grid>
            </CardContent>

            {renderTable()}

            {/* <TablePagination
                rowsPerPageOptions={[10, 15, 20, 30, 50, 100]}
                component="div"
                count={totalCount}
                rowsPerPage={rowsPerPage}
                page={page}
                onPageChange={() => undefined}
                onRowsPerPageChange={undefined}
                labelRowsPerPage="" // Set an empty string to hide the rows per page label
            /> */}
            <Dialog
                id="addRoleModal"
                className="hideBackdrop"
                onClose={() => handleAddModal()}
                open={isAddModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                <Grid container spacing={2}>
                    <Grid item sm={6} textAlign="left" sx={{ alignSelf: 'center' }}>
                        <Typography>
                            {formikInitialValues['id'] === 0
                                ? intl.formatMessage({ id: 'role.add-role' })
                                : intl.formatMessage({ id: 'role.edit-role' })}
                        </Typography>
                    </Grid>
                    <Grid item sm={6} textAlign="right">
                        <IconButton onClick={() => handleAddModal()} sx={{ alignSelf: 'end' }}>
                            <CancelTwoToneIcon />
                        </IconButton>
                    </Grid>
                </Grid>

                {isAddModalOpen && (
                    <>
                        <Formik
                            initialValues={formikInitialValues}
                            validationSchema={Yup.object().shape({
                                name: Yup.string()
                                    .max(255)
                                    .required(intl.formatMessage({ id: 'validation.role-name-required' })),
                                key: Yup.string()
                                    .max(255)
                                    .required(intl.formatMessage({ id: 'validation.role-code-required' }))
                            })}
                            onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                                try {
                                    await axiosServices
                                        .post(
                                            `${envRef?.API_URL}admin/role/edit`,
                                            {
                                                id: values.id || 0,
                                                name: values.name,
                                                key: values.key,
                                                dataScope: values.dataScope || 1,
                                                customDept: values.customDept || [],
                                                pid: values.pid || 0,
                                                level: values.level || 1,
                                                tree: values.tree || '',
                                                remark: values.remark || '',
                                                sort: values.sort || 0,
                                                status: values.status,
                                                createdAt: values.createdAt,
                                                updatedAt: values.updatedAt,
                                                orgAdmin: values.orgAdmin,
                                                label: values.label,
                                                value: values.value
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
                                                handleAddModal();
                                                fetchData();
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
                                setFieldTouched
                            }) => (
                                <form noValidate onSubmit={handleSubmit}>
                                    <Grid container alignItems="center" spacing={2}>
                                        <Grid item className="expandableSelectContainer" xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.pid && errors.pid)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <ExpandableSelect
                                                    label={intl.formatMessage({ id: 'role.upper-role' })}
                                                    id="pid"
                                                    options={addModalOptions}
                                                    onSelectChange={(e) => setFieldValue('pid', e[0])}
                                                    initialValue={[values.pid]}
                                                />
                                                {touched.pid && errors.pid && (
                                                    <FormHelperText error id="standard-weight-helper-text-pid-add">
                                                        {errors.pid.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.name && errors.name)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-role-name-add">
                                                    {intl.formatMessage({ id: 'role.role-name' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-role-name-add"
                                                    type="text"
                                                    value={values.name}
                                                    name="name"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.name && errors.name && (
                                                    <FormHelperText error id="standard-weight-helper-text-role-name-add">
                                                        {errors.name.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.key && errors.key)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-role-code-add">
                                                    {intl.formatMessage({ id: 'role.role-code' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-role-code-add"
                                                    type="text"
                                                    value={values.key}
                                                    name="key"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.key && errors.key && (
                                                    <FormHelperText error id="standard-weight-helper-text-role-code-add">
                                                        {errors.key.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <PrefixRadio
                                                options={statusFlatOptions}
                                                label={intl.formatMessage({ id: 'general.status' })}
                                                id="status"
                                                onSelectChange={(e) => setFieldValue('status', e)}
                                                initialValue={values.status}
                                            />
                                        </Grid>
                                        <Grid item xs={12}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.remark && errors.remark)}
                                                sx={{ ...theme.typography.customInput }}
                                                className="MultiLineInput"
                                            >
                                                <InputLabel htmlFor="outlined-adornment-remark-add">
                                                    {intl.formatMessage({ id: 'general.remarks' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    multiline
                                                    id="outlined-adornment-remark-add"
                                                    type="text"
                                                    value={values.remark}
                                                    name="remark"
                                                    rows={1}
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{
                                                        style: { marginTop: '4%', whiteSpace: 'break-spaces' }
                                                    }}
                                                />
                                                {touched.remark && errors.remark && (
                                                    <FormHelperText error id="standard-weight-helper-text-remark-add">
                                                        {errors.remark.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                    </Grid>

                                    {errors.submit && (
                                        <Box sx={{ mt: 3 }}>
                                            <FormHelperText error>{errors.submit.toString()}</FormHelperText>
                                        </Box>
                                    )}
                                    <Grid container justifyContent="flex-end" spacing={2}>
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
                                        <Grid item xs={12} sm={4} md={3}>
                                            <Box sx={{ mt: 2 }}>
                                                <AnimateButton>
                                                    <Button
                                                        fullWidth
                                                        size="large"
                                                        variant="outlined"
                                                        color="primary"
                                                        onClick={() => handleAddModal()}
                                                    >
                                                        {intl.formatMessage({ id: 'general.cancel' })}
                                                    </Button>
                                                </AnimateButton>
                                            </Box>
                                        </Grid>
                                    </Grid>
                                </form>
                            )}
                        </Formik>
                    </>
                )}
            </Dialog>

            <GeneralDialog
                confirmFunction={() => handleDelete(selectedIdsToDelete)}
                isOpen={isDeleteModalOpen}
                setIsOpen={() => setIsDeleteModalOpen(!isDeleteModalOpen)}
                id="delete-confirm-modal"
                type="warning"
                content={intl.formatMessage({ id: 'general.confirm-delete-content' })}
            />

            <Dialog
                id="menu-permission-modal"
                className="hideBackdrop"
                onClose={() => handleMenuPermissionModal(null)}
                open={isMenuPermissionModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                <Grid container spacing={gridSpacing}>
                    <Grid item xs={6} textAlign="left" sx={{ alignSelf: 'center', display: 'flex', alignItems: 'center' }}>
                        <Typography sx={{ fontWeight: '800' }}>
                            {intl.formatMessage({ id: 'general.assign-menu-permission' })} - {selectedRole?.label}
                        </Typography>
                    </Grid>
                    <Grid item xs={6} textAlign="right">
                        <IconButton onClick={() => handleMenuPermissionModal(null)} sx={{ alignSelf: 'end' }}>
                            <CancelTwoToneIcon />
                        </IconButton>
                    </Grid>
                </Grid>

                <Grid container spacing={gridSpacing} sx={{ overflowY: 'scroll', maxHeight: '100%', marginY: '2rem' }}>
                    <Grid className="expandableSelectContainer" item sm={12} textAlign="center" padding="1rem">
                        <ExpandableRadio
                            id="menuIds"
                            options={menuList}
                            onSelectChange={(e) => setSelectedPerms(e)}
                            multiSelect={true}
                            valueFieldName="key"
                            displayFieldName="label"
                            initialValue={menuPermissionInitialValue}
                            triggerExpandAll={triggerExpandAll}
                            triggerCheckAll={triggerCheckAll}
                        />
                    </Grid>
                </Grid>

                <Grid container spacing={gridSpacing} justifyContent="flex-end">
                    <Grid item xs={6} sm={3} textAlign="center">
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
                    <Grid item xs={6} sm={3} textAlign="center">
                        <Button
                            fullWidth
                            color="info"
                            variant="outlined"
                            onClick={() => setTriggerCheckAll(!triggerCheckAll)}
                            sx={{ alignSelf: 'end' }}
                        >
                            {triggerCheckAll
                                ? intl.formatMessage({ id: 'general.unselect-all' })
                                : intl.formatMessage({ id: 'general.select-all' })}
                        </Button>
                    </Grid>

                    <Grid item xs={6} sm={3} textAlign="center">
                        <Button
                            fullWidth
                            color="primary"
                            variant="contained"
                            onClick={() => {
                                handleChangeMenuPermission();
                                setIsMenuPermissionModalOpen(!isMenuPermissionModalOpen);
                            }}
                            sx={{ alignSelf: 'end' }}
                        >
                            {intl.formatMessage({ id: 'general.submit' })}
                        </Button>
                    </Grid>

                    <Grid item xs={6} sm={3} textAlign="center">
                        <Button
                            color="secondary"
                            fullWidth
                            variant="outlined"
                            onClick={() => handleMenuPermissionModal(null)}
                            sx={{ alignSelf: 'end' }}
                        >
                            {intl.formatMessage({ id: 'general.cancel' })}
                        </Button>
                    </Grid>
                </Grid>
            </Dialog>

            <Dialog
                id="data-permission-modal"
                className="hideBackdrop"
                onClose={() => handleDataPermissionModal(null)}
                open={isDataPermissionModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                <Grid container spacing={gridSpacing}>
                    <Grid item xs={6} textAlign="left" sx={{ alignSelf: 'center', display: 'flex', alignItems: 'center' }}>
                        <Typography sx={{ fontWeight: '800' }}>
                            {intl.formatMessage({ id: 'general.assign-data-permission' })} - {selectedRole?.label}
                        </Typography>
                    </Grid>
                    <Grid item xs={6} textAlign="right">
                        <IconButton onClick={() => handleDataPermissionModal(null)} sx={{ alignSelf: 'end' }}>
                            <CancelTwoToneIcon />
                        </IconButton>
                    </Grid>
                </Grid>

                <Grid container spacing={gridSpacing} sx={{ marginY: '2rem' }}>
                    <Grid className="expandableSelectContainer" item sm={12} textAlign="center" padding="1rem">
                        <ExpandableSelect
                            label={intl.formatMessage({ id: 'general.data-scope' })}
                            id="dataScope"
                            options={dataScopeList}
                            onSelectChange={(e) => setSelectedDataScope(e[0])}
                            valueFieldName="value"
                            displayFieldName="label"
                            initialValue={dataScopeInitialValue}
                            filterStatus={false}
                            disableParent={true}
                        />
                    </Grid>
                </Grid>

                <Grid container spacing={gridSpacing} justifyContent="flex-end">
                    <Grid item xs={6} sm={3} textAlign="center">
                        <Button
                            fullWidth
                            color="primary"
                            variant="contained"
                            onClick={() => {
                                handleChangeDataScope();
                                setIsDataPermissionModalOpen(!isDataPermissionModalOpen);
                            }}
                            sx={{ alignSelf: 'end' }}
                        >
                            {intl.formatMessage({ id: 'general.submit' })}
                        </Button>
                    </Grid>

                    <Grid item xs={6} sm={3} textAlign="center">
                        <Button
                            color="secondary"
                            fullWidth
                            variant="outlined"
                            onClick={() => handleDataPermissionModal(null)}
                            sx={{ alignSelf: 'end' }}
                        >
                            {intl.formatMessage({ id: 'general.cancel' })}
                        </Button>
                    </Grid>
                </Grid>
            </Dialog>
        </MainCard>
    );
};

export default Role;
