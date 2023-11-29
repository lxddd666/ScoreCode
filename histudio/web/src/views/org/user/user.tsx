import * as React from 'react';

// locale
import { FormattedMessage, useIntl } from 'react-intl';

// material-ui
import { useTheme } from '@mui/material/styles';
import {
    Avatar,
    Box,
    Button,
    CardContent,
    Checkbox,
    Dialog,
    Divider,
    FormControl,
    FormHelperText,
    Grid,
    IconButton,
    InputAdornment,
    InputLabel,
    Menu,
    MenuItem,
    OutlinedInput,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TablePagination,
    TableRow,
    TableSortLabel,
    TextField,
    Toolbar,
    Tooltip,
    Typography
} from '@mui/material';
import { visuallyHidden } from '@mui/utils';

// project imports
import { NestedSubOption, FlatOption } from 'types/option';
import { UserListData } from 'types/user';

import { ResponseList } from 'types/response';
import { useDispatch, useSelector } from 'store';
import { getUserList, getAdminInfo } from 'store/slices/user';

// assets
import DeleteIcon from '@mui/icons-material/DeleteTwoTone';
import SearchIcon from '@mui/icons-material/SearchTwoTone';
import EditIcon from '@mui/icons-material/EditTwoTone';
import ToggleOnIcon from '@mui/icons-material/ToggleOnTwoTone';
import ToggleOffIcon from '@mui/icons-material/ToggleOffOutlined';
import FolderOffTwoToneIcon from '@mui/icons-material/FolderOffTwoTone';
import CancelTwoToneIcon from '@mui/icons-material/CancelTwoTone';
import FirstPageIcon from '@mui/icons-material/FirstPage';
import KeyboardArrowLeft from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRight from '@mui/icons-material/KeyboardArrowRight';
import LastPageIcon from '@mui/icons-material/LastPage';

import { GetComparator, EnhancedTableHeadProps, HeadCell, ArrangementOrder, KeyedObject } from 'types';
import logo from 'assets/images/general/logo.png';

// ui-components
import Chip from 'ui-component/extended/Chip';
import MainCard from 'ui-component/cards/MainCard';
import ExpandButton from 'ui-component/searchbar/ExpandButton';
import ResetButton from 'ui-component/searchbar/ResetButton';
import SearchButton from 'ui-component/searchbar/SearchButton';
import AddButton from 'ui-component/searchbar/AddButton';
import QrButton from 'ui-component/searchbar/QrButton';
import ExpandableSelect from 'ui-component/general/ExpandableSelect';
import GeneralDialog from 'ui-component/general/GeneralDialog';
import CustomDateTimeRangePicker from 'ui-component/general/CustomDateTimeRangePicker';

// API
import axiosServices from 'utils/axios';

// Crypto JS
import * as CryptoJS from 'crypto-js';

// env
import envRef from 'environment';
import { isEqual } from 'lodash';

// third party
import * as Yup from 'yup';
import { Formik, FormikValues } from 'formik';
import { QRCodeSVG } from 'qrcode.react';
import copy from 'clipboard-copy';
import generator from 'generate-password-ts';

// project imports
import AnimateButton from 'ui-component/extended/AnimateButton';
import useScriptRef from 'hooks/useScriptRef';
import { openSnackbar } from 'store/slices/snackbar';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import MultiSelect from 'ui-component/general/MultiSelect';
import PrefixRadio from 'ui-component/general/PrefixRadio';
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';

import flattenNestedObjects from 'utils/flattenNestedObjects';
import { gridSpacing } from 'store/constant';
import DynamicPermissionComponent from 'ui-component/general/DynamicPermissionComponent';
import { AccessPermissions, statusFlatOptions } from 'constant/general';

// START handle last active date
function handleLastActiveDate(lastActiveDate: string) {
    const currentDateTime = new Date().getTime();
    const lastActiveDateTime = new Date(lastActiveDate).getTime();

    if (lastActiveDate) {
        const msDiff = Math.abs(currentDateTime - lastActiveDateTime);
        const diffMinutes = Math.floor(msDiff / (1000 * 60));
        const diffHours = Math.floor(msDiff / (1000 * 60 * 60));
        const diffDays = Math.floor(msDiff / (1000 * 60 * 60 * 24));
        if (msDiff >= 0 && msDiff <= 1000 * 60 * 1)
            return (
                <TableCell align="center">
                    <FormattedMessage id="user.just-now" />
                </TableCell>
            );
        else if (msDiff > 1000 * 60 * 1 && msDiff < 1000 * 60 * 60)
            return (
                <TableCell align="center">
                    {diffMinutes} <FormattedMessage id="user.minutes-before" />
                </TableCell>
            );
        else if (msDiff >= 1000 * 60 * 60 && msDiff < 1000 * 60 * 60 * 24)
            return (
                <TableCell align="center">
                    {diffHours} <FormattedMessage id="user.hours-before" />
                </TableCell>
            );
        else if (msDiff >= 1000 * 60 * 60 * 24 && msDiff < 1000 * 60 * 60 * 24 * 30)
            return (
                <TableCell align="center">
                    {diffDays} <FormattedMessage id="user.days-before" />
                </TableCell>
            );
        else return <TableCell align="center">{lastActiveDate}</TableCell>;
    }
    return (
        <TableCell align="center">
            <FormattedMessage id="user.never-login" />
        </TableCell>
    );
}
// END handle last active date

// START table sort
function descendingComparator(a: KeyedObject, b: KeyedObject, orderBy: string) {
    if (b[orderBy] < a[orderBy]) {
        return -1;
    }
    if (b[orderBy] > a[orderBy]) {
        return 1;
    }
    return 0;
}

const getComparator: GetComparator = (order, orderBy) =>
    order === 'desc' ? (a, b) => descendingComparator(a, b, orderBy) : (a, b) => -descendingComparator(a, b, orderBy);

function stableSort(array: UserListData[], comparator: (a: UserListData, b: UserListData) => number) {
    const stabilizedThis = array.map((el: UserListData, index: number) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0] as UserListData, b[0] as UserListData);
        if (order !== 0) return order;
        return (a[1] as number) - (b[1] as number);
    });
    return stabilizedThis.map((el) => el[0]);
}
// END table sort

// table header options
const headCells: HeadCell[] = [
    {
        id: 'id',
        numeric: false,
        label: 'user.user-id',
        align: 'center'
    },
    {
        id: 'username',
        numeric: false,
        label: 'user.username',
        align: 'center'
    },
    {
        id: 'realName',
        numeric: false,
        label: 'user.real-name',
        align: 'center'
    },
    {
        id: 'avatar',
        numeric: false,
        label: 'user.avatar',
        align: 'center'
    },
    {
        id: 'mobile',
        numeric: true,
        label: 'user.mobile-number',
        align: 'center'
    },
    {
        id: 'roleName',
        numeric: false,
        label: 'user.role-name',
        align: 'center'
    },
    {
        id: 'deptName',
        numeric: false,
        label: 'user.department-name',
        align: 'center'
    },
    {
        id: 'status',
        numeric: false,
        label: 'general.status',
        align: 'center'
    },
    {
        id: 'lastActiveAt',
        numeric: false,
        label: 'user.last-active-date',
        align: 'center'
    },
    {
        id: 'createdAt',
        numeric: false,
        label: 'user.created-date',
        align: 'center'
    }
];

// ==============================|| TABLE HEADER ||============================== //

interface CustomListEnhancedTableHeadProps extends EnhancedTableHeadProps {
    selected: number[];
    handleDeleteModal: (selected: number[]) => void;
}

function EnhancedTableHead({
    onSelectAllClick,
    order,
    orderBy,
    numSelected,
    rowCount,
    onRequestSort,
    selected,
    handleDeleteModal
}: CustomListEnhancedTableHeadProps) {
    const createSortHandler = (property: string) => (event: React.SyntheticEvent<Element, Event>) => {
        onRequestSort(event, property);
    };
    const intl = useIntl();

    return (
        <TableHead>
            <TableRow>
                <TableCell padding="checkbox" sx={{ pl: 3 }}>
                    <Checkbox
                        color="primary"
                        indeterminate={numSelected > 0 && numSelected < rowCount}
                        checked={rowCount > 0 && numSelected === rowCount}
                        onChange={onSelectAllClick}
                        inputProps={{
                            'aria-label': 'select all desserts'
                        }}
                    />
                </TableCell>
                {numSelected > 0 && (
                    <TableCell padding="none" colSpan={headCells.length + 1}>
                        <EnhancedTableToolbar handleDeleteModal={handleDeleteModal} numSelected={selected.length} selected={selected} />
                    </TableCell>
                )}
                {numSelected <= 0 &&
                    headCells.map((headCell) => (
                        <TableCell
                            key={headCell.id}
                            align={headCell.align}
                            padding={headCell.disablePadding ? 'none' : 'normal'}
                            sortDirection={orderBy === headCell.id ? order : false}
                        >
                            {headCell.id === 'avatar' ? (
                                intl.formatMessage({ id: headCell.label })
                            ) : (
                                <TableSortLabel
                                    active={orderBy === headCell.id}
                                    direction={orderBy === headCell.id ? order : 'asc'}
                                    onClick={createSortHandler(headCell.id)}
                                >
                                    {intl.formatMessage({ id: headCell.label })}
                                    {orderBy === headCell.id ? (
                                        <Box component="span" sx={visuallyHidden}>
                                            {order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                                        </Box>
                                    ) : null}
                                </TableSortLabel>
                            )}
                        </TableCell>
                    ))}
                {numSelected <= 0 && (
                    <TableCell sortDirection={false} align="center" className="sticky" sx={{ pr: 3, right: '0' }}>
                        {intl.formatMessage({ id: 'general.control' })}
                    </TableCell>
                )}
            </TableRow>
        </TableHead>
    );
}

// ==============================|| TABLE HEADER TOOLBAR ||============================== //
type CustomEnhancedTableToolbarProps = {
    numSelected: number;
    selected: number[];
    handleDeleteModal: (selected: number[]) => void;
};
const EnhancedTableToolbar = ({ numSelected, selected, handleDeleteModal }: CustomEnhancedTableToolbarProps) => (
    <Toolbar
        sx={{
            p: 0,
            pl: 1,
            pr: 1,
            ...(numSelected > 0 && {
                color: (theme) => theme.palette.secondary.main
            })
        }}
    >
        {numSelected > 0 ? (
            <Typography
                sx={{
                    color: (theme) => theme.palette.primary.light
                }}
                variant="h4"
            >
                {numSelected} <FormattedMessage id="general.selected" />
            </Typography>
        ) : (
            <Typography
                sx={{
                    color: (theme) => theme.palette.primary.light
                }}
                variant="h6"
                id="tableTitle"
            >
                <FormattedMessage id="user.user-list" />
            </Typography>
        )}
        <Box sx={{ flexGrow: 2 }} />
        {numSelected > 0 && (
            <Tooltip title="Delete">
                <IconButton
                    onClick={() => {
                        handleDeleteModal(selected);
                    }}
                    size="large"
                >
                    <DeleteIcon fontSize="small" />
                </IconButton>
            </Tooltip>
        )}
    </Toolbar>
);

// ==============================|| USER LIST ||============================== //

const User = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);

    type SearchFields = {
        username: string;
        realName: string;
        mobile: string;
        email: string;
        status: number;
        roleId: number;
        created_at: Date[];
    };
    const initSearchFields: SearchFields = { username: '', realName: '', mobile: '', email: '', status: 0, roleId: -1, created_at: [] };

    const [flatPostOptions, setFlatPostOptions] = React.useState<FlatOption[]>([]);
    const [roleOptions, setRoleOptions] = React.useState<NestedSubOption[]>([]);
    const [deptOptions, setDeptOptions] = React.useState<NestedSubOption[]>([]);

    async function getOptions() {
        // Get position list
        await axiosServices
            .get(`${envRef?.API_URL}admin/post/list`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setFlatPostOptions(response.data.data.list);
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

        // Get role list
        await axiosServices
            .get(`${envRef?.API_URL}admin/role/list`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setRoleOptions(response.data.data.list);
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

        // Get dept list
        await axiosServices
            .get(`${envRef?.API_URL}admin/dept/option`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setDeptOptions(response.data.data.list);
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
        (async function () {
            getOptions();
            await dispatch(getAdminInfo(intl));
        })();
    }, []);

    const [order, setOrder] = React.useState<ArrangementOrder>('asc');
    const [orderBy, setOrderBy] = React.useState<string>('id');
    const [selected, setSelected] = React.useState<number[]>([]);
    const [page, setPage] = React.useState<number>(0);
    const [totalCount, setTotalCount] = React.useState<number>(0);
    const [pageCount, setPageCount] = React.useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = React.useState<number>(10);
    const [search, setSearch] = React.useState<SearchFields>(initSearchFields);
    const [rows, setRows] = React.useState<UserListData[]>();
    const [res, setRes] = React.useState<ResponseList<UserListData>>();
    const { userList, adminInfo } = useSelector((state) => state.user);
    const [status, setStatus] = React.useState<number>(0);
    const [role, setRole] = React.useState<number>(-1);
    const [isDataListEmpty, setIsDataListEmpty] = React.useState<boolean>(true);
    const scriptedRef = useScriptRef();
    const [showPassword, setShowPassword] = React.useState(false);
    const FormikInitialValuesTemplate: FormikValues = {
        id: 0,
        realName: '',
        username: '',
        roleId: 1,
        deptId: 100,
        postIds: [],
        password: '',
        mobile: '',
        email: '',
        sex: 1,
        status: 1,
        remark: '',
        submit: null
    };
    const [formikInitialValues, setFormikInitialValues] = React.useState<FormikValues>(FormikInitialValuesTemplate);
    const [createdDate, setCreatedDate] = React.useState<Date[] | null>(null);
    const [reset, setReset] = React.useState<boolean>(false);

    const handleClickShowPassword = () => {
        setShowPassword(!showPassword);
    };

    const handleMouseDownPassword = (event: React.SyntheticEvent) => {
        event.preventDefault();
    };
    const initialQueryParamString: string[] = [`page=1`, `pageSize=${rowsPerPage}`, `roleId=-1`];
    const [queryParamString, setQueryParamString] = React.useState<String[]>(initialQueryParamString);

    React.useEffect(() => {
        const updatedParams = [
            `page=${page + 1}`,
            `pageSize=${rowsPerPage}`,
            `roleId=${role}`,
            ...queryParamString.filter(
                (param) => !param.startsWith('page=') && !param.startsWith('pageSize=') && !param.startsWith('roleId=')
            )
        ];
        setQueryParamString(updatedParams);
    }, [page, rowsPerPage, role]);

    React.useEffect(() => {
        setQueryParamString(findUpdatedParams());
    }, [search]);

    const findUpdatedParams = () => {
        const updatedParams = [
            ...queryParamString.filter((param) => param.startsWith('page=') || param.startsWith('pageSize=') || param.startsWith('roleId='))
        ];
        Object.entries(search).forEach(([key, value]) => {
            if (key === 'roleId') {
                setRole(+value);
            } else if (key === 'created_at') {
                if (createdDate) {
                    updatedParams.push(`${key}[]=${createdDate[0].getTime()}`);
                    updatedParams.push(`${key}[]=${createdDate[1].getTime()}`);
                }
            } else {
                if (value !== '') {
                    updatedParams.push(`${key}=${value}`);
                }
            }
        });
        return updatedParams;
    };

    React.useEffect(() => {
        fetchData(queryParamString);
    }, [dispatch]);

    const fetchData = async (queries: String[]) => {
        try {
            setLoading(true);
            setRows(undefined);
            await dispatch(getUserList(queries.filter((query) => !query.endsWith('=') && query !== 'status=0').join('&')));
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
        setRes(userList!);
    }, [userList]);
    React.useEffect(() => {
        setRows(res?.data?.list ? res.data.list : []);
    }, [res]);
    React.useEffect(() => {
        setPage(res?.data?.page ? res.data.page - 1 : 0);
        setRowsPerPage(res?.data?.pageSize ? res.data.pageSize : 10);
        setIsDataListEmpty(isEqual(rows, []));
        setTotalCount(res?.data?.totalCount ? res.data.totalCount : 0);
        setPageCount(res?.data?.pageCount ? res.data.pageCount : 1);
    }, [rows]);

    const [expanded, setExpanded] = React.useState<boolean>(false);
    const handleExpandClick = React.useCallback(() => {
        setExpanded((prevExpanded) => !prevExpanded);
    }, []);

    const [isAddModalOpen, setIsAddModalOpen] = React.useState<boolean>(false);

    const decryptAES = (encryptedBase64: string) => {
        const key = CryptoJS.enc.Utf8.parse(envRef?.AES_KEY);
        const decrypted = CryptoJS.AES.decrypt(encryptedBase64, key, {
            mode: CryptoJS.mode.ECB,
            padding: CryptoJS.pad.Pkcs7,
            keySize: 128 / 8
        });
        if (decrypted) {
            try {
                const str = decrypted.toString(CryptoJS.enc.Utf8);
                if (str.length > 0) {
                    return str;
                }
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
            }
        }
    };

    const handleAddModal = async (id?: number) => {
        if (id) {
            await axiosServices
                .get(`${envRef?.API_URL}admin/member/view?id=${id}`, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        const selectedFormikValues: FormikValues = {
                            id: response.data.data.id,
                            realName: response.data.data.realName,
                            username: response.data.data.username,
                            roleId: response.data.data.roleId,
                            deptId: response.data.data.deptId,
                            postIds: response.data.data.postIds,
                            password: response.data.data.passwordHash === '' ? decryptAES(response.data.data.passwordHash) : '',
                            mobile: response.data.data.mobile,
                            email: response.data.data.email,
                            sex: response.data.data.sex,
                            status: response.data.data.status,
                            remark: response.data.data.remark,
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
            setIsAddModalOpen(!isAddModalOpen);
            setFormikInitialValues(FormikInitialValuesTemplate);
        }
    };

    const handleSearchClick = () => {
        fetchData(queryParamString);
        setRes(userList!);
    };

    const handleResetClick = () => {
        setSearch(initSearchFields);
        setPage(0);
        setCreatedDate(null);
        setReset(!reset);
        setStatus(0);
        setRole(-1);
        fetchData(initialQueryParamString);
    };

    const statusOptions = [
        {
            id: 0,
            value: 0,
            name: intl.formatMessage({ id: 'user.please-select-status' })
        },
        {
            id: 1,
            value: 1,
            name: intl.formatMessage({ id: 'general.normal' })
        },
        {
            id: 2,
            value: 2,
            name: intl.formatMessage({ id: 'general.disabled' })
        }
    ];

    const genderOptions: FlatOption[] = [
        {
            id: 1,
            name: 'user.male'
        },
        {
            id: 2,
            name: 'user.female'
        },
        {
            id: 3,
            name: 'user.unknown'
        }
    ];

    const handleStatus = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const val = parseInt(event.target.value);
        setSearch({ ...search, status: val });
        setStatus(val);
    };

    const handleRole = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const val = parseInt(event.target.value);
        setRole(val);
    };

    const handleActivate = async (id: number, status: number) => {
        await axiosServices
            .post(
                `${envRef?.API_URL}admin/member/status`,
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
    };

    const [isDeleteModalOpen, setIsDeleteModalOpen] = React.useState<boolean>(false);
    const [selectedUserToDelete, setSelectedUserToDelete] = React.useState<number[]>();
    const handleDeleteModal = (id: number[]) => {
        setSelectedUserToDelete(id);
        setIsDeleteModalOpen(!isDeleteModalOpen);
    };

    const handleDelete = async (ids: number[]) => {
        await axiosServices
            .post(`${envRef?.API_URL}admin/member/delete`, { id: ids }, { headers: {} })
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
        setSelected([]);
    };

    const handleSearch = (event: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement> | undefined, searchParam: string) => {
        const newString = event?.target.value;
        setSearch({ ...search, [searchParam]: newString || '' });
        // if (newString) {
        //     const newRows = rows?.filter((row: KeyedObject) => {
        //         let matches = true;
        //         let containsQuery = false;

        //         if (searchParam === 'status' || searchParam === 'roleId') {
        //             if (row[searchParam].toString().toLowerCase() === newString.toString().toLowerCase()) {
        //                 containsQuery = true;
        //             }
        //         } else {
        //             if (row[searchParam].toString().toLowerCase().includes(newString.toString().toLowerCase())) {
        //                 containsQuery = true;
        //             }
        //         }

        //         if (!containsQuery) {
        //             matches = false;
        //         }
        //         return matches;
        //     });
        //     setRows(newRows);
        // } else {
        //     setRows(userList!.data!.list!);
        // }
    };

    const handleRequestSort = (event: React.SyntheticEvent<Element, Event>, property: string) => {
        const isAsc = orderBy === property && order === 'asc';
        setOrder(isAsc ? 'desc' : 'asc');
        setOrderBy(property);
    };

    const handleSelectAllClick = (event: React.ChangeEvent<HTMLInputElement>) => {
        if (event.target.checked) {
            if (selected.length > 0) {
                setSelected([]);
            } else {
                const newSelectedId = rows?.map((n) => n.id);
                setSelected(newSelectedId || []);
            }
            return;
        }
        setSelected([]);
    };

    const handleClick = (event: React.MouseEvent<HTMLTableHeaderCellElement, MouseEvent>, id: number) => {
        const selectedIndex = selected.indexOf(id);
        let newSelected: number[] = [];

        if (selectedIndex === -1) {
            newSelected = newSelected.concat(selected, id);
        } else if (selectedIndex === 0) {
            newSelected = newSelected.concat(selected.slice(1));
        } else if (selectedIndex === selected.length - 1) {
            newSelected = newSelected.concat(selected.slice(0, -1));
        } else if (selectedIndex > 0) {
            newSelected = newSelected.concat(selected.slice(0, selectedIndex), selected.slice(selectedIndex + 1));
        }

        setSelected(newSelected);
    };

    const handleChangePage = (event: React.MouseEvent<HTMLButtonElement, MouseEvent> | null, newPage: number) => {
        setPage(newPage);
        const updatedParams = [
            `page=${newPage + 1}`,
            `pageSize=${rowsPerPage}`,
            `roleId=${role}`,
            ...queryParamString.filter(
                (param) => !param.startsWith('page=') && !param.startsWith('pageSize=') && !param.startsWith('roleId=')
            )
        ];
        fetchData(updatedParams);
    };

    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement> | undefined) => {
        const val = event ? parseInt(event.target.value) : 10;
        setRowsPerPage(val);
        setPage(page === 1 ? 0 : page);
        const updatedParams = [
            `page=${page + 1}`,
            `pageSize=${val}`,
            `roleId=${role}`,
            ...queryParamString.filter(
                (param) => !param.startsWith('page=') && !param.startsWith('pageSize=') && !param.startsWith('roleId=')
            )
        ];
        fetchData(updatedParams);
    };
    const [anchorElMap, setAnchorElMap] = React.useState<{ [key: number]: HTMLElement | null }>({});
    const handleMenuOpen = (event: React.MouseEvent<HTMLButtonElement>, id: number) => {
        setAnchorElMap((prev) => ({
            ...prev,
            [id]: event.currentTarget
        }));
    };
    const handleMenuClose = () => {
        setAnchorElMap({});
    };

    const [qr, setQr] = React.useState<string>('');
    const [isQrModalOpen, setIsQrModalOpen] = React.useState<boolean>(false);
    const handleQrClick = (inviteCode: string) => {
        setQr(inviteCode);
        setIsQrModalOpen(!isQrModalOpen);
        handleMenuClose();
    };

    const currentUrl = `${window.location.origin}`;

    const handleCopyUrl = () => {
        const urlToCopy = `${currentUrl}/register/inviteCode/${qr}`;
        copy(urlToCopy)
            .then(() => {
                alert('URL copied to clipboard!');
            })
            .catch((error) => {
                console.error('Copy failed:', error);
            });
    };
    const [isResetModalOpen, setIsResetModalOpen] = React.useState<boolean>(false);
    const [selectedUserToReset, setSelectedUserToReset] = React.useState<number>();
    const [newPassword, setnewPassword] = React.useState<string>('');
    const handleResetModal = (id: number) => {
        const newPassword = generator.generate({ length: 12 });
        setnewPassword(newPassword);
        setSelectedUserToReset(id);
        setIsResetModalOpen(!isResetModalOpen);
        handleMenuClose();
    };

    const handleReset = (ids: number, newPassword: string) => {
        axiosServices
            .post(`${envRef?.API_URL}admin/member/resetPwd`, { id: ids, password: newPassword }, { headers: {} })
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
        setIsResetModalOpen(!isResetModalOpen);
        setSelected([]);
    };

    const isSelected = (id: number) => selected.indexOf(id) !== -1;

    const mobileNumberRegExp = /^(\+?\d{0,4})?\s?-?\s?(\(?\d{3}\)?)\s?-?\s?(\(?\d{3}\)?)\s?-?\s?(\(?\d{4}\)?)?$/;
    const defaultErrorMessage = intl.formatMessage({ id: 'auth-register.default-error' });

    function renderTable() {
        return (
            <TableContainer>
                <Table sx={{ minWidth: 750 }} aria-labelledby="tableTitle">
                    <EnhancedTableHead
                        numSelected={selected.length}
                        order={order}
                        orderBy={orderBy}
                        onSelectAllClick={handleSelectAllClick}
                        onRequestSort={handleRequestSort}
                        rowCount={rows?.length || 0}
                        selected={selected}
                        handleDeleteModal={handleDeleteModal}
                    />
                    <TableBody>
                        {loading ? (
                            <TableRow>
                                <TableCell align="center" colSpan={headCells.length + 2}>
                                    <SkeletonLoader />
                                </TableCell>
                            </TableRow>
                        ) : isDataListEmpty ? (
                            <TableRow>
                                <TableCell align="center" colSpan={headCells.length + 2}>
                                    <FolderOffTwoToneIcon sx={{ verticalAlign: 'bottom' }} />
                                    {intl.formatMessage({ id: 'general.no-records' })}
                                </TableCell>
                            </TableRow>
                        ) : (
                            stableSort(rows ? rows : [], getComparator(order, orderBy)).map((row, index) => {
                                if (typeof row === 'number') return null;
                                const isItemSelected = isSelected(row.id);
                                const labelId = `enhanced - table - checkbox - ${index} `;

                                return (
                                    <TableRow
                                        hover
                                        role="checkbox"
                                        aria-checked={isItemSelected}
                                        tabIndex={-1}
                                        key={index}
                                        selected={isItemSelected}
                                    >
                                        <TableCell padding="checkbox" sx={{ pl: 3 }} onClick={(event) => handleClick(event, row.id)}>
                                            <Checkbox
                                                color="primary"
                                                checked={isItemSelected}
                                                inputProps={{
                                                    'aria-labelledby': labelId
                                                }}
                                            />
                                        </TableCell>
                                        <TableCell
                                            component="th"
                                            id={labelId}
                                            scope="row"
                                            onClick={(event) => handleClick(event, row.id)}
                                            sx={{ cursor: 'pointer' }}
                                        >
                                            <Typography
                                                variant="subtitle1"
                                                sx={{ color: theme.palette.mode === 'dark' ? 'grey.600' : 'grey.900' }}
                                            >
                                                {row.id || '-'}
                                            </Typography>
                                        </TableCell>
                                        <TableCell>{row.username}</TableCell>
                                        <TableCell>{row.realName}</TableCell>
                                        <TableCell align="center">
                                            {
                                                // Special handling avatar using HiSeven logo due to incompatible path
                                                // Currently displaying "-" for null avatar
                                                row.avatar === '/src/assets/images/logo.png' ? (
                                                    <Avatar src={logo} alt={'avatar.' + row.id} />
                                                ) : row.avatar === '' ? (
                                                    '-'
                                                ) : (
                                                    <Avatar src={row.avatar} alt={'avatar.' + row.id} />
                                                )
                                            }
                                        </TableCell>
                                        <TableCell align="center">{row.mobile || '-'}</TableCell>
                                        <TableCell align="center">{row.roleName || '-'}</TableCell>
                                        <TableCell align="center">{row.deptName || '-'}</TableCell>
                                        <TableCell align="center">
                                            {row.status === 1 && (
                                                <Chip
                                                    label={intl.formatMessage({ id: 'general.normal' })}
                                                    size="small"
                                                    chipcolor="success"
                                                />
                                            )}
                                            {row.status === 2 && (
                                                <Chip
                                                    label={intl.formatMessage({ id: 'general.disabled' })}
                                                    size="small"
                                                    chipcolor="orange"
                                                />
                                            )}
                                        </TableCell>
                                        {handleLastActiveDate(row.lastActiveAt)}
                                        <TableCell align="center">{row.createdAt || '-'}</TableCell>
                                        <TableCell align="center" className="sticky" sx={{ pr: 3, right: 0 }}>
                                            <DynamicPermissionComponent
                                                requiredPermissions={[AccessPermissions.MEMBER_STATUS]}
                                                children={
                                                    <IconButton
                                                        onClick={() => handleActivate(row.id, row.status)}
                                                        color="secondary"
                                                        size="medium"
                                                        aria-label="View"
                                                    >
                                                        {row.status === 1 ? (
                                                            <ToggleOnIcon sx={{ fontSize: '1.3rem' }} />
                                                        ) : (
                                                            <ToggleOffIcon sx={{ fontSize: '1.3rem' }} />
                                                        )}
                                                    </IconButton>
                                                }
                                            />
                                            <DynamicPermissionComponent
                                                requiredPermissions={[AccessPermissions.MEMBER_EDIT, AccessPermissions.MEMBER_VIEW]}
                                                children={
                                                    <IconButton
                                                        onClick={() => handleAddModal(row.id)}
                                                        color="secondary"
                                                        size="medium"
                                                        aria-label="Edit"
                                                    >
                                                        <EditIcon sx={{ fontSize: '1.3rem' }} />
                                                    </IconButton>
                                                }
                                            />

                                            <DynamicPermissionComponent
                                                requiredPermissions={[AccessPermissions.MEMBER_DELETE]}
                                                children={
                                                    <IconButton
                                                        onClick={() => handleDeleteModal([row.id])}
                                                        color="secondary"
                                                        size="medium"
                                                        aria-label="Delete"
                                                    >
                                                        <DeleteIcon sx={{ fontSize: '1.3rem' }} />
                                                    </IconButton>
                                                }
                                            />
                                            <div style={{ marginLeft: 'auto' }}>
                                                <Button
                                                    onClick={(event) => handleMenuOpen(event, row.id)}
                                                    aria-controls={`qrqr-${row.id}`}
                                                    aria-haspopup="true"
                                                    style={{ height: '30px', width: '120px' }}
                                                    size="small"
                                                    variant="contained"
                                                    color="secondary"
                                                >
                                                    {intl.formatMessage({ id: 'user.more' })}
                                                </Button>
                                                <Menu
                                                    id={`qrqr-${row.id}`} 
                                                    anchorEl={anchorElMap[row.id]}
                                                    keepMounted
                                                    open={Boolean(anchorElMap[row.id])} 
                                                    onClose={handleMenuClose}
                                                >
                                                    <MenuItem onClick={() => handleResetModal(row.id)}>
                                                        {intl.formatMessage({ id: 'user.pwd' })}
                                                    </MenuItem>
                                                    <MenuItem onClick={() => handleQrClick(row.inviteCode)}>
                                                        {intl.formatMessage({ id: 'user.qr' })}
                                                    </MenuItem>
                                                </Menu>
                                            </div>
                                        </TableCell>
                                    </TableRow>
                                );
                            })
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
        );
    }

    interface TablePaginationActionsProps {
        count: number;
        page: number;
        rowsPerPage: number;
        onPageChange: (event: React.MouseEvent<HTMLButtonElement>, newPage: number) => void;
    }

    function TablePaginationActions(props: TablePaginationActionsProps) {
        const { count, page, rowsPerPage, onPageChange } = props;
        const textFieldRef = React.useRef<HTMLInputElement | null>(null);

        const handleFirstPageButtonClick = (event: React.MouseEvent<HTMLButtonElement>) => {
            onPageChange(event, 0);
        };

        const handleBackButtonClick = (event: React.MouseEvent<HTMLButtonElement>) => {
            onPageChange(event, page - 1);
        };

        const handleNextButtonClick = (event: React.MouseEvent<HTMLButtonElement>) => {
            onPageChange(event, page + 1);
        };

        const handleLastPageButtonClick = (event: React.MouseEvent<HTMLButtonElement>) => {
            onPageChange(event, Math.max(0, Math.ceil(count / rowsPerPage) - 1));
        };

        return (
            <Grid
                container
                spacing={gridSpacing}
                alignItems="center"
                justifyContent="flex-end"
                marginLeft={1}
                className="customTablePagination"
            >
                <Grid item xs={12} sm={6} textAlign="end">
                    <TextField
                        fullWidth
                        type="number"
                        defaultValue={page + 1}
                        InputProps={{
                            endAdornment: (
                                <InputAdornment position="end">
                                    <IconButton
                                        onClick={() => {
                                            const inputValue = textFieldRef.current ? parseInt(textFieldRef.current.value) : 0;
                                            const updatedParams = [
                                                `page=${inputValue > pageCount ? pageCount : inputValue}`,
                                                `pageSize=${rowsPerPage}`,
                                                ...queryParamString.filter(
                                                    (param) => !param.startsWith('page=') && !param.startsWith('pageSize=')
                                                )
                                            ];
                                            fetchData(updatedParams);
                                        }}
                                        color="secondary"
                                        size="small"
                                        aria-label="Search"
                                    >
                                        <SearchIcon fontSize="small" />
                                    </IconButton>
                                </InputAdornment>
                            )
                        }}
                        placeholder={intl.formatMessage({ id: 'general.jump-to-page' })}
                        inputRef={textFieldRef}
                        size="small"
                        label={intl.formatMessage({ id: 'general.page-number' })}
                    />
                </Grid>
                <Grid item xs={12} sm={6} className="hidden-xs" textAlign="end" container justifyContent="space-evenly">
                    <IconButton
                        onClick={handleFirstPageButtonClick}
                        disabled={page === 0}
                        aria-label={intl.formatMessage({ id: 'general.first-page' })}
                    >
                        {theme.direction === 'rtl' ? <LastPageIcon /> : <FirstPageIcon />}
                    </IconButton>
                    <IconButton
                        onClick={handleBackButtonClick}
                        disabled={page === 0}
                        aria-label={intl.formatMessage({ id: 'general.previous-page' })}
                    >
                        {theme.direction === 'rtl' ? <KeyboardArrowRight /> : <KeyboardArrowLeft />}
                    </IconButton>
                    <IconButton
                        onClick={handleNextButtonClick}
                        disabled={page >= Math.ceil(count / rowsPerPage) - 1}
                        aria-label={intl.formatMessage({ id: 'general.next-page' })}
                    >
                        {theme.direction === 'rtl' ? <KeyboardArrowLeft /> : <KeyboardArrowRight />}
                    </IconButton>
                    <IconButton
                        onClick={handleLastPageButtonClick}
                        disabled={page >= Math.ceil(count / rowsPerPage) - 1}
                        aria-label={intl.formatMessage({ id: 'general.last-page' })}
                    >
                        {theme.direction === 'rtl' ? <FirstPageIcon /> : <LastPageIcon />}
                    </IconButton>
                </Grid>
            </Grid>
        );
    }

    return (
        <MainCard title={<FormattedMessage id="user.user-list" />} content={false}>
            <CardContent>
                <Grid container alignItems="center" spacing={gridSpacing} justifyContent="flex-start">
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
                            onChange={(event) => handleSearch(event, 'username')}
                            placeholder={intl.formatMessage({ id: 'user.search-username' })}
                            value={search.username}
                            size="small"
                            label={intl.formatMessage({ id: 'user.username' })}
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
                            onChange={(event) => handleSearch(event, 'realName')}
                            placeholder={intl.formatMessage({ id: 'user.search-real-name' })}
                            value={search.realName}
                            size="small"
                            label={intl.formatMessage({ id: 'user.real-name' })}
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
                            onChange={(event) => handleSearch(event, 'mobile')}
                            placeholder={intl.formatMessage({ id: 'user.search-mobile' })}
                            value={search.mobile}
                            size="small"
                            label={intl.formatMessage({ id: 'user.mobile-number' })}
                        />
                    </Grid>
                    {expanded && (
                        <>
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
                                    onChange={(event) => handleSearch(event, 'email')}
                                    placeholder={intl.formatMessage({ id: 'user.search-email' })}
                                    value={search.email}
                                    size="small"
                                    label={intl.formatMessage({ id: 'user.email' })}
                                />
                            </Grid>
                            <Grid item xs={12} sm={4} md={3}>
                                <TextField
                                    select
                                    fullWidth
                                    onChange={(event) => {
                                        handleStatus(event);
                                        handleSearch(event, 'status');
                                    }}
                                    value={status}
                                    size="small"
                                    label={intl.formatMessage({ id: 'general.status' })}
                                >
                                    {statusOptions.map((option) => {
                                        return (
                                            <MenuItem key={option.id} value={option.value}>
                                                {option.name}
                                            </MenuItem>
                                        );
                                    })}
                                </TextField>
                            </Grid>
                            <Grid item xs={12} sm={4} md={3}>
                                <TextField
                                    select
                                    fullWidth
                                    onChange={handleRole}
                                    value={role}
                                    size="small"
                                    label={intl.formatMessage({ id: 'user.role-name' })}
                                >
                                    <MenuItem key={-1} value={-1}>
                                        {intl.formatMessage({ id: 'user.all' })}
                                    </MenuItem>
                                    {flattenNestedObjects(roleOptions).map((option) => {
                                        return (
                                            <MenuItem key={option.id} value={option.id}>
                                                {option.name}
                                            </MenuItem>
                                        );
                                    })}
                                </TextField>
                            </Grid>
                            <Grid item xs={12} sm={8} md={6}>
                                <CustomDateTimeRangePicker
                                    reset={reset}
                                    onSelectChange={(e) => {
                                        if (e) {
                                            setCreatedDate(e);
                                            setSearch({ ...search, created_at: e });
                                        } else {
                                            setCreatedDate(null);
                                            setSearch({ ...search, created_at: [] });
                                        }
                                    }}
                                    label={intl.formatMessage({ id: 'user.created-date' })}
                                />
                            </Grid>
                        </>
                    )}
                    <Grid item xs={12} sm={4} md={3} sx={{ display: 'flex', justifyContent: 'flex-end', marginLeft: 'auto' }}>
                        <AddButton onClick={() => handleAddModal()} tooltipTitle={intl.formatMessage({ id: 'user.add-user' })} />
                        <QrButton onClick={() => handleQrClick(adminInfo?.inviteCode)} />
                        <SearchButton onClick={handleSearchClick} />
                        <ResetButton onClick={handleResetClick} />
                        <ExpandButton onClick={handleExpandClick} transformValue={expanded ? 'rotate(180deg)' : 'rotate(0deg)'} />
                    </Grid>
                </Grid>
            </CardContent>

            <DynamicPermissionComponent
                requiredPermissions={[
                    AccessPermissions.DEPT_LIST,
                    AccessPermissions.POST_LIST,
                    AccessPermissions.ROLE_LIST,
                    AccessPermissions.MEMBER_LIST,
                    AccessPermissions.DEPT_OPTION
                ]}
                children={renderTable()}
            />

            <TablePagination
                rowsPerPageOptions={[10, 15, 20, 30, 50, 100]}
                component="div"
                count={totalCount}
                rowsPerPage={rowsPerPage}
                sx={{
                    '& .MuiTablePagination-spacer': {
                        display: 'none'
                    },
                    '& .MuiToolbar-root.MuiToolbar-gutters.MuiToolbar-regular.MuiTablePagination-toolbar': {
                        justifyContent: 'flex-end'
                    },
                    alignItems: 'flex-end'
                }}
                page={page}
                onPageChange={handleChangePage}
                onRowsPerPageChange={handleChangeRowsPerPage}
                ActionsComponent={TablePaginationActions}
                labelRowsPerPage="" // Set an empty string to hide the rows per page label
                labelDisplayedRows={({ from, to, page, count }) => {
                    return (
                        <Typography className="hidden-xs" variant="caption">
                            {intl.formatMessage({ id: 'general.page-number' })}: {page + 1}
                            &nbsp;|&nbsp;
                            {from} - {to}
                        </Typography>
                    );
                }} // Customize the displayed rows label
            />
            <Dialog
                id="Qr"
                className="hideBackdrop"
                maxWidth="lg"
                onClick={() => handleQrClick(qr)}
                open={isQrModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1rem 1.5rem' }, backgroundColor: '#f5f5f5' }}
            >
                <Grid container spacing={gridSpacing}>
                    <Grid
                        item
                        sm={6}
                        textAlign="justify"
                        sx={{
                            alignSelf: 'center',
                            '& .MuiTypography-root': {
                                fontSize: '1.2rem' // Adjust the font size as needed
                            }
                        }}
                    >
                        <Typography>
                            <FormattedMessage id="user.qrDisplay" />
                        </Typography>
                    </Grid>
                    <Grid item sm={6} textAlign="right">
                        <IconButton onClick={() => handleQrClick(qr)} sx={{ alignSelf: 'end' }}>
                            <CancelTwoToneIcon />
                        </IconButton>
                    </Grid>
                </Grid>
                <Grid container spacing={gridSpacing}>
                    <Grid item sm={12} textAlign="center" sx={{ alignSelf: 'center' }}>
                        <div id="Qr">{isQrModalOpen && <QRCodeSVG value={`${currentUrl}/register/inviteCode/${qr}`} size={300} />}</div>
                    </Grid>
                    <Grid container justifyContent="center">
                    <OutlinedInput
                            id="test-input"
                            type="text"
                            value={`${currentUrl}/register/inviteCode/${qr}`}
                            name="inviteCodetest"
                            readOnly
                            style={{ width: '50%', height: '70%'}} 
                        />
                        <Grid item>
                            <Button variant="outlined" color="primary" onClick={handleCopyUrl}>
                                Copy URL
                            </Button>
                        </Grid>
                    </Grid>
                </Grid>
            </Dialog>
            <Dialog
                id="addUserModal"
                className="hideBackdrop"
                maxWidth="sm"
                fullWidth
                onClose={() => handleAddModal()}
                open={isAddModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                <Grid container spacing={gridSpacing}>
                    <Grid item sm={6} textAlign="left" sx={{ alignSelf: 'center' }}>
                        <Typography>
                            {formikInitialValues['id'] === 0
                                ? intl.formatMessage({ id: 'user.add-user' })
                                : intl.formatMessage({ id: 'user.edit-user' })}
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
                                username: Yup.string()
                                    .max(255)
                                    .required(intl.formatMessage({ id: 'auth-register.username-required' })),
                                roleId: Yup.number().required(intl.formatMessage({ id: 'validation.role-required' })),
                                deptId: Yup.number().required(intl.formatMessage({ id: 'validation.department-required' })),
                                postIds: Yup.array()
                                    .of(Yup.number())
                                    .test('required', intl.formatMessage({ id: 'validation.position-required' }), (value) => {
                                        return value!.length > 0;
                                    }),
                                password: Yup.string().max(255),
                                mobile: Yup.string()
                                    .matches(mobileNumberRegExp, intl.formatMessage({ id: 'auth-register.mobile-invalid' }))
                                    .max(255),
                                email: Yup.string()
                                    .email(intl.formatMessage({ id: 'auth-register.email-invalid' }))
                                    .max(255)
                            })}
                            onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                                try {
                                    // const text = CryptoJS.enc.Utf8.parse(values.password);
                                    // const key = CryptoJS.enc.Utf8.parse(envRef?.AES_KEY);

                                    // const encryptedPassword = CryptoJS.AES.encrypt(text, key, {
                                    //     mode: CryptoJS.mode.ECB,
                                    //     padding: CryptoJS.pad.Pkcs7,
                                    //     keySize: 128 / 8
                                    // }).toString();

                                    await axiosServices
                                        .post(
                                            `${envRef?.API_URL}admin/member/edit`,
                                            {
                                                username: values.username,
                                                realName: values.realName,
                                                password: values.password,
                                                mobile: values.mobile,
                                                email: values.email,
                                                deptId: values.deptId,
                                                roleId: values.roleId,
                                                postIds: values.postIds,
                                                sex: values.sex,
                                                status: values.status,
                                                remark: values.remark,
                                                /* 
                                                    below fields are hardcoded cuz sharing API with edit,
                                                    suspecting id=0 means adding. 
                                                */
                                                id: values.id | 0,
                                                leader: values.leader,
                                                createdAt: values.createdat,
                                                updatedAt: values.updatedAt,
                                                sort: values.sort | 0
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
                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.realName && errors.realName)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-real-name-add">
                                                    {intl.formatMessage({ id: 'user.real-name' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-real-name-add"
                                                    type="text"
                                                    value={values.realName}
                                                    name="realName"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.realName && errors.realName && (
                                                    <FormHelperText error id="standard-weight-helper-text-real-name-add">
                                                        {errors.realName.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.username && errors.username)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-username-add">
                                                    {intl.formatMessage({ id: 'user.username' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-username-add"
                                                    type="text"
                                                    value={values.username}
                                                    name="username"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.username && errors.username && (
                                                    <FormHelperText error id="standard-weight-helper-text-username-add">
                                                        {errors.username.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                    </Grid>

                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                        <Grid item className="expandableSelectContainer" xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.roleId && errors.roleId)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <ExpandableSelect
                                                    label={intl.formatMessage({ id: 'user.role-name' })}
                                                    id="role"
                                                    options={roleOptions}
                                                    onSelectChange={(e) => setFieldValue('roleId', e[0])}
                                                    initialValue={[values.roleId]}
                                                    valueFieldName="value"
                                                />
                                                {touched.roleId && errors.roleId && (
                                                    <FormHelperText error id="standard-weight-helper-text-role-add">
                                                        {errors.roleId.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                        <Grid item className="expandableSelectContainer" xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.deptId && errors.deptId)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <ExpandableSelect
                                                    label={intl.formatMessage({ id: 'user.department-name' })}
                                                    id="dept"
                                                    options={deptOptions}
                                                    onSelectChange={(e) => setFieldValue('deptId', e[0])}
                                                    initialValue={[values.deptId]}
                                                    valueFieldName="value"
                                                />
                                                {touched.deptId && errors.deptId && (
                                                    <FormHelperText error id="standard-weight-helper-text-role-add">
                                                        {errors.deptId.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                    </Grid>

                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={touched.postIds && Boolean(errors.postIds)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <MultiSelect
                                                    label={intl.formatMessage({ id: 'user.position' })}
                                                    id="position"
                                                    options={flatPostOptions}
                                                    size="medium"
                                                    onSelectChange={(e) => {
                                                        setFieldValue('postIds', e);
                                                        if (e.length > 0) setFieldTouched('postIds', true);
                                                    }}
                                                    initialValue={values.postIds}
                                                />
                                                {touched.postIds && errors.postIds && (
                                                    <FormHelperText error id="standard-weight-helper-text-role-add">
                                                        {errors.postIds.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.password && errors.password)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-password-add">
                                                    {intl.formatMessage({ id: 'user.password' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-password-add"
                                                    type={showPassword ? 'text' : 'password'}
                                                    value={values.password}
                                                    name="password"
                                                    label="Password"
                                                    onBlur={handleBlur}
                                                    onChange={(e) => {
                                                        handleChange(e);
                                                    }}
                                                    endAdornment={
                                                        <InputAdornment position="end">
                                                            <IconButton
                                                                aria-label="toggle password visibility"
                                                                onClick={handleClickShowPassword}
                                                                onMouseDown={handleMouseDownPassword}
                                                                edge="end"
                                                                size="small"
                                                            >
                                                                {showPassword ? <Visibility /> : <VisibilityOff />}
                                                            </IconButton>
                                                        </InputAdornment>
                                                    }
                                                    inputProps={{}}
                                                />
                                                {touched.password && errors.password && (
                                                    <FormHelperText error id="standard-weight-helper-text-password-add">
                                                        {errors.password.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                    </Grid>

                                    <Divider sx={{ paddingY: '1rem' }} textAlign="left">
                                        {intl.formatMessage({ id: 'user.fill-more-info' })}
                                    </Divider>

                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.mobile && errors.mobile)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-mobile-add">
                                                    {intl.formatMessage({ id: 'user.mobile-number' })}
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
                                                    {intl.formatMessage({ id: 'user.email' })}
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
                                    </Grid>

                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                        <Grid item xs={12} sm={6}>
                                            <PrefixRadio
                                                options={genderOptions}
                                                label={intl.formatMessage({ id: 'user.gender' })}
                                                id="sex"
                                                onSelectChange={(e) => setFieldValue('sex', e)}
                                                initialValue={values.sex}
                                            />
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
                                    </Grid>

                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                        <Grid item xs={12}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.remark && errors.remark)}
                                                sx={{ ...theme.typography.customInput }}
                                                className="MultiLineInput"
                                            >
                                                <InputLabel htmlFor="outlined-adornment-remark-add">
                                                    {intl.formatMessage({ id: 'user.remark' })}
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
                                    <Grid container justifyContent="flex-end" spacing={gridSpacing}>
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
            <GeneralDialog
                id="reset-confirm-modal"
                type="warning"
                title={intl.formatMessage({ id: 'user.confirm-reset' })}
                content={
                    intl.formatMessage({ id: 'user.confirm-reset-content' }) +
                    intl.formatMessage({ id: 'user.confirm-reset-content2' }) +
                    newPassword +
                    intl.formatMessage({ id: 'user.confirm-reset-content3' })
                } // Ensure this is a string
                isOpen={isResetModalOpen}
                setIsOpen={() => setIsResetModalOpen(!isResetModalOpen)}
                confirmFunction={() => handleReset(selectedUserToReset!, newPassword)}
            />
        </MainCard>
    );
};

export default User;
