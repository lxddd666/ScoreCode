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
    InputAdornment,
    InputLabel,
    OutlinedInput,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    // TablePagination,
    TableRow,
    TextField,
    Tooltip,
    Typography
} from '@mui/material';

// project imports
import { ResponseList } from 'types/response';
import { DeptListData } from 'types/user';
import { useDispatch, useSelector } from 'store';
import { getDeptList } from 'store/slices/user';

// assets
import SearchIcon from '@mui/icons-material/SearchTwoTone';
import FolderOffTwoToneIcon from '@mui/icons-material/FolderOffTwoTone';
import CancelTwoToneIcon from '@mui/icons-material/CancelTwoTone';
import HelpTwoToneIcon from '@mui/icons-material/HelpTwoTone';

import { HeadCell } from 'types';

// ui-components
import AnimateButton from 'ui-component/extended/AnimateButton';
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';
import MainCard from 'ui-component/cards/MainCard';
import ExpandButton from 'ui-component/searchbar/ExpandButton';
import ResetButton from 'ui-component/searchbar/ResetButton';
import SearchButton from 'ui-component/searchbar/SearchButton';
import AddButton from 'ui-component/searchbar/AddButton';
import ExpandableSelect from 'ui-component/general/ExpandableSelect';
import GeneralDialog from 'ui-component/general/GeneralDialog';
import CustomDateTimeRangePicker from 'ui-component/general/CustomDateTimeRangePicker';
import ExpandableTableRow from 'ui-component/general/ExpandableTableRow';
import PrefixRadio from 'ui-component/general/PrefixRadio';

// API
import axiosServices from 'utils/axios';

// env
import envRef from 'environment';
import { isEqual } from 'lodash';

// third party
import * as Yup from 'yup';
import { Formik, FormikValues } from 'formik';

// project imports
import useScriptRef from 'hooks/useScriptRef';
import { openSnackbar } from 'store/slices/snackbar';
import { FieldToShow } from 'types/general';
import { statusFlatOptions } from 'constant/general';

// table header options
const headCells: HeadCell[] = [
    {
        id: 'name',
        numeric: false,
        label: 'department.department-name',
        align: 'left'
    },
    {
        id: 'code',
        numeric: false,
        label: 'department.department-code',
        align: 'center'
    },
    {
        id: 'leader',
        numeric: false,
        label: 'department.department-leader',
        align: 'center'
    },
    {
        id: 'phone',
        numeric: false,
        label: 'general.mobile-number',
        align: 'center'
    },
    {
        id: 'email',
        numeric: true,
        label: 'general.email',
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
        fieldName: 'code',
        fieldType: ''
    },
    {
        fieldName: 'leader',
        fieldType: ''
    },
    {
        fieldName: 'phone',
        fieldType: ''
    },
    {
        fieldName: 'email',
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

const Department = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);

    type SearchFields = {
        name: string;
        code: string;
        leader: string;
        created_at: Date[];
    };
    const initSearchFields: SearchFields = { name: '', code: '', leader: '', created_at: [] };

    // const [page, setPage] = React.useState<number>(0);
    // const [totalCount, setTotalCount] = React.useState<number>(0);
    // const [rowsPerPage, setRowsPerPage] = React.useState<number>(10);
    const [search, setSearch] = React.useState<SearchFields>(initSearchFields);
    const [rows, setRows] = React.useState<DeptListData[]>();
    const [addModalOptions, setAddModalOptions] = React.useState<DeptListData[]>();
    const [res, setRes] = React.useState<ResponseList<DeptListData>>();
    const { deptList } = useSelector((state) => state.user);
    const [isDataListEmpty, setIsDataListEmpty] = React.useState<boolean>(true);
    const scriptedRef = useScriptRef();
    const FormikInitialValuesTemplate: FormikValues = {
        id: 0,
        name: '',
        code: '',
        leader: '',
        phone: '',
        email: '',
        sort: 0,
        status: 1,
        pid: 0, // parent ID
        type: '',
        createdAt: '',
        updatedAt: '',
        submit: null
    };
    const [formikInitialValues, setFormikInitialValues] = React.useState<FormikValues>(FormikInitialValuesTemplate);
    const [createdDate, setCreatedDate] = React.useState<Date[] | null>(null);
    const [reset, setReset] = React.useState<boolean>(false);

    const [queryParamString, setQueryParamString] = React.useState<String[]>([]);

    React.useEffect(() => {
        setQueryParamString(findUpdatedParams());
    }, [search]);

    const findUpdatedParams = () => {
        const updatedParams: string[] = [];
        Object.entries(search).forEach(([key, value]) => {
            if (key === 'created_at' && createdDate) {
                updatedParams.push(`${key}[]=${createdDate[0].getTime()}`);
                updatedParams.push(`${key}[]=${createdDate[1].getTime()}`);
            } else {
                if (value !== '') {
                    updatedParams.push(`${key}=${value}`);
                }
            }
        });
        return updatedParams;
    };

    React.useEffect(() => {
        fetchData([]);
    }, [dispatch]);

    const fetchData = async (queries: String[]) => {
        try {
            setLoading(true);
            setIsDataListEmpty(true);
            setRows(undefined);
            await dispatch(getDeptList(queries.filter((query) => !query.endsWith('=') && query !== 'status=0').join('&')));
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
        setRes(deptList!);
    }, [deptList]);
    React.useEffect(() => {
        setRows(res?.data?.list ? res.data.list : []);
    }, [res]);
    React.useEffect(() => {
        setIsDataListEmpty(isEqual(rows, []));
        // setPage(res?.data?.page ? res.data.page - 1 : 0);
        // setRowsPerPage(res?.data?.pageSize ? res.data.pageSize : 10);
        // setTotalCount(res?.data?.totalCount ? res.data.totalCount : 0);

        const appendedModalOptions: DeptListData[] = [
            {
                id: 0,
                orgId: 0,
                pid: 0,
                name: intl.formatMessage({ id: 'department.top-department' }),
                code: 'top',
                type: 'dept',
                leader: 'admin',
                phone: '15888888888',
                email: '163@qq.com',
                level: 3,
                tree: '',
                sort: 1,
                status: 1,
                createdAt: '',
                updatedAt: '',
                label: intl.formatMessage({ id: 'department.top-department' }),
                value: 0,
                children: rows || null
            }
        ];
        setAddModalOptions(appendedModalOptions);
    }, [rows]);

    const [expanded, setExpanded] = React.useState<boolean>(false);
    const handleExpandClick = React.useCallback(() => {
        setExpanded((prevExpanded) => !prevExpanded);
    }, []);

    const [isAddModalOpen, setIsAddModalOpen] = React.useState<boolean>(false);

    const handleAddModal = (id?: number, type?: 'add' | 'edit') => {
        if (type === 'edit') {
            axiosServices
                .get(`${envRef?.API_URL}admin/dept/view?id=${id}`, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        const selectedFormikValues: FormikValues = {
                            id: response.data.data.id,
                            orgId: response.data.data.orgId,
                            pid: response.data.data.pid,
                            name: response.data.data.name,
                            code: response.data.data.code,
                            type: response.data.data.type,
                            leader: response.data.data.leader,
                            mobile: response.data.data.phone,
                            email: response.data.data.email,
                            level: response.data.data.level,
                            tree: response.data.data.tree,
                            sort: response.data.data.sort,
                            status: response.data.data.status,
                            createdAt: response.data.data.createdAt,
                            updatedAt: response.data.data.updatedAt,
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

    const handleSearchClick = () => {
        fetchData(queryParamString);
        setRes(deptList!);
    };

    const handleResetClick = () => {
        setSearch(initSearchFields);
        setCreatedDate(null);
        setReset(!reset);
        fetchData([]);
    };

    const handleActivate = (id: number, status: number) => {
        axiosServices
            .post(
                `${envRef?.API_URL}admin/dept/status`,
                {
                    id: id,
                    /*
                        Currently only two status - normal/disabled,
                        therefore handle in this way, may need to tune
                        if more statuses are added in the future
                    */
                    status: status === 1 ? 2 : 1
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
                    fetchData([]);
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
        dispatch(getDeptList());
    };

    const [isDeleteModalOpen, setIsDeleteModalOpen] = React.useState<boolean>(false);
    const [selectedUserToDelete, setSelectedUserToDelete] = React.useState<number[]>();
    const handleDeleteModal = (id: number[]) => {
        setSelectedUserToDelete(id);
        setIsDeleteModalOpen(!isDeleteModalOpen);
    };

    const handleDelete = (ids: number[]) => {
        axiosServices
            .post(`${envRef?.API_URL}admin/dept/delete`, { id: ids }, { headers: {} })
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
                    fetchData(queryParamString);
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
        setSelectedUserToDelete([]);
    };

    const handleSearch = (event: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement> | undefined, searchParam: string) => {
        const newString = event?.target.value;
        setSearch({ ...search, [searchParam]: newString || '' });
    };

    const mobileNumberRegExp = /^(\+?\d{0,4})?\s?-?\s?(\(?\d{3}\)?)\s?-?\s?(\(?\d{3}\)?)\s?-?\s?(\(?\d{4}\)?)?$/;
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
                                handleActivate={handleActivate}
                                handleAddModal={handleAddModal}
                                handleDeleteModal={handleDeleteModal}
                                fieldsToShow={fieldsToShow}
                            />
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
        );
    }

    return (
        <MainCard title={<FormattedMessage id="department.department-management" />} content={false}>
            <CardContent>
                <Grid container alignItems="center" spacing={2}>
                    <Grid item xs={12} sm={4} md={3}>
                        <TextField
                            fullWidth
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon fontSize="small" />
                                    </InputAdornment>
                                )
                            }}
                            onChange={(event) => handleSearch(event, 'name')}
                            placeholder={intl.formatMessage({ id: 'department.search-department-name' })}
                            value={search.name}
                            size="small"
                            label={intl.formatMessage({ id: 'department.department-name' })}
                        />
                    </Grid>
                    <Grid item xs={12} sm={4} md={3}>
                        <TextField
                            fullWidth
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon fontSize="small" />
                                    </InputAdornment>
                                )
                            }}
                            onChange={(event) => handleSearch(event, 'code')}
                            placeholder={intl.formatMessage({ id: 'department.search-department-code' })}
                            value={search.code}
                            size="small"
                            label={intl.formatMessage({ id: 'department.department-code' })}
                        />
                    </Grid>
                    <Grid item xs={12} sm={4} md={3}>
                        <TextField
                            fullWidth
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon fontSize="small" />
                                    </InputAdornment>
                                )
                            }}
                            onChange={(event) => handleSearch(event, 'leader')}
                            placeholder={intl.formatMessage({ id: 'department.search-department-leader' })}
                            value={search.leader}
                            size="small"
                            label={intl.formatMessage({ id: 'department.department-leader' })}
                        />
                    </Grid>

                    {expanded && (
                        <>
                            <Grid item xs={12} sm={8} md={6}>
                                <CustomDateTimeRangePicker
                                    reset={reset}
                                    onSelectChange={(e) => {
                                        setCreatedDate(e);
                                        setSearch({ ...search, created_at: e || [] });
                                    }}
                                    label={intl.formatMessage({ id: 'general.created-date' })}
                                />
                            </Grid>
                        </>
                    )}
                    <Grid item xs={12} sm={4} md={3} sx={{ display: 'flex', justifyContent: 'flex-end', marginLeft: 'auto' }}>
                        <AddButton
                            onClick={() => handleAddModal()}
                            tooltipTitle={intl.formatMessage({ id: 'department.add-department' })}
                        />
                        <SearchButton onClick={handleSearchClick} />
                        <ResetButton onClick={handleResetClick} />
                        <ExpandButton onClick={handleExpandClick} transformValue={expanded ? 'rotate(180deg)' : 'rotate(0deg)'} />
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
                id="addDeptModal"
                className="hideBackdrop"
                onClose={() => handleAddModal()}
                open={isAddModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                <Grid container spacing={2}>
                    <Grid item sm={6} textAlign="left" sx={{ alignSelf: 'center' }}>
                        <Typography>
                            {formikInitialValues['id'] === 0
                                ? intl.formatMessage({ id: 'department.add-department' })
                                : intl.formatMessage({ id: 'department.edit-department' })}
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
                                    .required(intl.formatMessage({ id: 'validation.department-name-required' })),
                                code: Yup.string()
                                    .max(255)
                                    .required(intl.formatMessage({ id: 'validation.department-code-required' })),
                                mobile: Yup.string()
                                    .matches(mobileNumberRegExp, intl.formatMessage({ id: 'auth-register.mobile-invalid' }))
                                    .max(255),
                                email: Yup.string()
                                    .email(intl.formatMessage({ id: 'auth-register.email-invalid' }))
                                    .max(255)
                            })}
                            onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                                try {
                                    await axiosServices
                                        .post(
                                            `${envRef?.API_URL}admin/dept/edit`,
                                            {
                                                id: values.id || 0,
                                                pid: values.pid,
                                                name: values.name,
                                                code: values.code,
                                                type: values.type || 'dept',
                                                leader: values.leader || '',
                                                phone: values.mobile || '',
                                                email: values.email || '',
                                                sort: values.sort || 0,
                                                status: values.status || 1,
                                                createdAt: values.createdAt || '',
                                                updatedAt: values.updatedAt || ''
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
                                                fetchData([]);
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
                                                    label={intl.formatMessage({ id: 'department.upper-department' })}
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
                                                <InputLabel htmlFor="outlined-adornment-dept-name-add">
                                                    {intl.formatMessage({ id: 'department.department-name' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-dept-name-add"
                                                    type="text"
                                                    value={values.name}
                                                    name="name"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.name && errors.name && (
                                                    <FormHelperText error id="standard-weight-helper-text-dept-name-add">
                                                        {errors.name.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.code && errors.code)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-dept-code-add">
                                                    {intl.formatMessage({ id: 'department.department-code' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-dept-code-add"
                                                    type="text"
                                                    value={values.code}
                                                    name="code"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.code && errors.code && (
                                                    <FormHelperText error id="standard-weight-helper-text-dept-code-add">
                                                        {errors.code.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.leader && errors.leader)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-dept-leader-add">
                                                    {intl.formatMessage({ id: 'department.department-leader' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-dept-leader-add"
                                                    type="text"
                                                    value={values.leader}
                                                    name="leader"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.leader && errors.leader && (
                                                    <FormHelperText error id="standard-weight-helper-text-dept-leader-add">
                                                        {errors.leader.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>

                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.mobile && errors.mobile)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-mobile-add">
                                                    {intl.formatMessage({ id: 'general.mobile-number' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-mobile-add"
                                                    type="mobile"
                                                    value={values.mobile}
                                                    name="mobile"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.mobile && errors.mobile && (
                                                    <FormHelperText error id="standard-weight-helper-text-mobile-add">
                                                        {errors.mobile.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>

                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.email && errors.email)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-email-add">
                                                    {intl.formatMessage({ id: 'general.email' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-email-add"
                                                    type="email"
                                                    value={values.email}
                                                    name="email"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.email && errors.email && (
                                                    <FormHelperText error id="standard-weight-helper-text-email-add">
                                                        {errors.email.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>

                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.sort && errors.sort)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-sort-add">
                                                    {intl.formatMessage({ id: 'general.sort' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-sort-add"
                                                    type="number"
                                                    value={values.sort}
                                                    name="sort"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.sort && errors.sort && (
                                                    <FormHelperText error id="standard-weight-helper-text-sort-add">
                                                        {errors.sort.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>

                                        <Grid item xs={12} sm={6}>
                                            <PrefixRadio
                                                options={statusFlatOptions}
                                                label={intl.formatMessage({ id: 'user.status' })}
                                                id="status"
                                                onSelectChange={(e) => setFieldValue('status', e)}
                                                initialValue={values.status}
                                            />
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
                confirmFunction={() => handleDelete(selectedUserToDelete!)}
                isOpen={isDeleteModalOpen}
                setIsOpen={() => setIsDeleteModalOpen(!isDeleteModalOpen)}
                id="delete-confirm-modal"
                type="warning"
                content={intl.formatMessage({ id: 'general.confirm-delete-content' })}
            />
        </MainCard>
    );
};

export default Department;
