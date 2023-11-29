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
    Collapse,
    Dialog,
    FormControl,
    FormHelperText,
    Grid,
    IconButton,
    InputAdornment,
    InputLabel,
    MenuItem,
    OutlinedInput,
    Select,
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
import { WhatsAccountListData } from 'types/whats';

import { ResponseList } from 'types/response';
import { useDispatch, useSelector } from 'store';
import { getWhatsAccountList } from 'store/slices/whats';

// assets
import DeleteIcon from '@mui/icons-material/DeleteTwoTone';
import SearchIcon from '@mui/icons-material/SearchTwoTone';
import EditIcon from '@mui/icons-material/EditTwoTone';
import FolderOffTwoToneIcon from '@mui/icons-material/FolderOffTwoTone';
import CancelTwoToneIcon from '@mui/icons-material/CancelTwoTone';
import ContactPhoneTwoToneIcon from '@mui/icons-material/ContactPhoneTwoTone';
import SendTwoToneIcon from '@mui/icons-material/SendTwoTone';
import ExpandCircleDownTwoToneIcon from '@mui/icons-material/ExpandCircleDownTwoTone';
import FirstPageIcon from '@mui/icons-material/FirstPage';
import KeyboardArrowLeft from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRight from '@mui/icons-material/KeyboardArrowRight';
import LastPageIcon from '@mui/icons-material/LastPage';
import ChatBubbleTwoToneIcon from '@mui/icons-material/ChatBubbleTwoTone';

import { GetComparator, EnhancedTableHeadProps, HeadCell, ArrangementOrder, KeyedObject } from 'types';
import logo from 'assets/images/general/logo.png';

// ui-components
import Chip from 'ui-component/extended/Chip';
import MainCard from 'ui-component/cards/MainCard';
import ExpandButton from 'ui-component/searchbar/ExpandButton';
import ResetButton from 'ui-component/searchbar/ResetButton';
import SearchButton from 'ui-component/searchbar/SearchButton';
import AddButton from 'ui-component/searchbar/AddButton';
import GeneralDialog from 'ui-component/general/GeneralDialog';
import CustomDateTimeRangePicker from 'ui-component/general/CustomDateTimeRangePicker';

// API
import axiosServices from 'utils/axios';

// env
import envRef from 'environment';
import { isEqual } from 'lodash';

// third party
import * as Yup from 'yup';
import { Formik, FormikValues } from 'formik';

// project imports
import AnimateButton from 'ui-component/extended/AnimateButton';
import useScriptRef from 'hooks/useScriptRef';
import { openSnackbar } from 'store/slices/snackbar';
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';

import { gridSpacing } from 'store/constant';
import DynamicPermissionComponent from 'ui-component/general/DynamicPermissionComponent';
import { AccessPermissions, defaultErrorMessage } from 'constant/general';
import ImportButton from 'ui-component/searchbar/ImportButton';
import LoginUserButton from 'ui-component/searchbar/LoginUserButton';
import ImportDialog from 'ui-component/general/ImportDialog';
import { WhatsAccountHeaderValidation } from 'constant/validation';
import { Link } from 'react-router-dom';

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
                    <FormattedMessage id="general.just-now" />
                </TableCell>
            );
        else if (msDiff > 1000 * 60 * 1 && msDiff < 1000 * 60 * 60)
            return (
                <TableCell align="center">
                    {diffMinutes} <FormattedMessage id="general.minutes-before" />
                </TableCell>
            );
        else if (msDiff >= 1000 * 60 * 60 && msDiff < 1000 * 60 * 60 * 24)
            return (
                <TableCell align="center">
                    {diffHours} <FormattedMessage id="general.hours-before" />
                </TableCell>
            );
        else if (msDiff >= 1000 * 60 * 60 * 24 && msDiff < 1000 * 60 * 60 * 24 * 30)
            return (
                <TableCell align="center">
                    {diffDays} <FormattedMessage id="general.days-before" />
                </TableCell>
            );
        else return <TableCell align="center">{lastActiveDate}</TableCell>;
    }
    return (
        <TableCell align="center">
            <FormattedMessage id="general.never-login" />
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

function stableSort(array: WhatsAccountListData[], comparator: (a: WhatsAccountListData, b: WhatsAccountListData) => number) {
    const stabilizedThis = array.map((el: WhatsAccountListData, index: number) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0] as WhatsAccountListData, b[0] as WhatsAccountListData);
        if (order !== 0) return order;
        return (a[1] as number) - (b[1] as number);
    });
    return stabilizedThis.map((el) => el[0]);
}
// END table sort

// table header options
const headCells: HeadCell[] = [
    {
        id: 'account',
        numeric: false,
        label: 'whats.account-number',
        align: 'center'
    },
    {
        id: 'nickName',
        numeric: false,
        label: 'whats.account-nickname',
        align: 'center'
    },
    {
        id: 'avatar',
        numeric: false,
        label: 'whats.account-avatar',
        align: 'center'
    },
    {
        id: 'accountStatus',
        numeric: true,
        label: 'whats.account-status',
        align: 'center'
    },
    {
        id: 'isOnline',
        numeric: false,
        label: 'general.is-online',
        align: 'center'
    },
    {
        id: 'createdAt',
        numeric: false,
        label: 'general.last-active-date',
        align: 'center'
    },
    {
        id: 'proxyAddress',
        numeric: false,
        label: 'general.proxy-address',
        align: 'center'
    },
    {
        id: 'comment',
        numeric: false,
        label: 'general.remarks',
        align: 'center'
    }
];

// ==============================|| TABLE HEADER ||============================== //

interface CustomListEnhancedTableHeadProps extends EnhancedTableHeadProps {
    selected: number[];
    handleDeleteModal: (selected: number[]) => void;
    handleLoginModal: (ids: number[]) => void;
}

function EnhancedTableHead({
    onSelectAllClick,
    order,
    orderBy,
    numSelected,
    rowCount,
    onRequestSort,
    selected,
    handleDeleteModal,
    handleLoginModal
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
                        <EnhancedTableToolbar
                            handleDeleteModal={handleDeleteModal}
                            numSelected={selected.length}
                            selected={selected}
                            handleLoginModal={handleLoginModal}
                        />
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
    handleLoginModal: (ids: number[]) => void;
};

const EnhancedTableToolbar = ({ numSelected, selected, handleDeleteModal, handleLoginModal }: CustomEnhancedTableToolbarProps) => (
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
                <FormattedMessage id="whats.whats-account-management" />
            </Typography>
        )}
        <Box sx={{ flexGrow: 2 }} />
        {numSelected > 0 && (
            <>
                <LoginUserButton onClick={(ids) => handleLoginModal(ids || selected)} tooltipTitle="general.login" />
                <Tooltip title={<FormattedMessage id="general.delete" />}>
                    <IconButton
                        onClick={() => {
                            handleDeleteModal(selected);
                        }}
                        size="large"
                    >
                        <DeleteIcon fontSize="small" />
                    </IconButton>
                </Tooltip>
            </>
        )}
    </Toolbar>
);

// ==============================|| USER LIST ||============================== //

const WhatsAccount = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);

    type SearchFields = {
        account: string;
        accountStatus: number;
        proxyAddress: string;
        created_at: Date[];
    };
    const initSearchFields: SearchFields = { account: '', accountStatus: 0, proxyAddress: '', created_at: [] };

    const [order, setOrder] = React.useState<ArrangementOrder>('asc');
    const [orderBy, setOrderBy] = React.useState<string>('id');
    const [selected, setSelected] = React.useState<number[]>([]);
    const [selectedRow, setSelectedRow] = React.useState<WhatsAccountListData | null>();
    const [page, setPage] = React.useState<number>(0);
    const [totalCount, setTotalCount] = React.useState<number>(0);
    const [pageCount, setPageCount] = React.useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = React.useState<number>(10);
    const [search, setSearch] = React.useState<SearchFields>(initSearchFields);
    const [rows, setRows] = React.useState<WhatsAccountListData[]>();
    const [res, setRes] = React.useState<ResponseList<WhatsAccountListData>>();
    const { whatsAccountList } = useSelector((state) => state.whats);
    const [accountStatus, setAccountStatus] = React.useState<number>(-1);
    const [isDataListEmpty, setIsDataListEmpty] = React.useState<boolean>(true);
    const scriptedRef = useScriptRef();
    const FormikInitialValuesTemplate: FormikValues = {
        id: 0,
        account: '',
        nickName: '',
        avatar: '',
        accountStatus: 0,
        isOnline: 1,
        proxyAddress: '',
        comment: '',
        encryption: '',
        deletedAt: '',
        createdAt: '',
        updatedAt: '',
        submit: null
    };
    const [formikInitialValues, setFormikInitialValues] = React.useState<FormikValues>(FormikInitialValuesTemplate);
    const [createdDate, setCreatedDate] = React.useState<Date[] | null>(null);
    const [reset, setReset] = React.useState<boolean>(false);

    const [whatsAccountStatusData, setWhatsAccountStatusData] = React.useState<any[]>([]);
    const [whatsAccountLoginStatusData, setWhatsAccountLoginStatusData] = React.useState<any[]>([]);

    React.useEffect(() => {
        getOptions();
    }, []);

    async function getOptions() {
        await axiosServices
            .get(`${envRef?.API_URL}admin/dictData/options?types[]=account_status&types[]=login_status`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setWhatsAccountStatusData(response.data.data.account_status);
                    setWhatsAccountLoginStatusData(response.data.data.login_status);
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

    const initialQueryParamString: string[] = [`page=1`, `pageSize=${rowsPerPage}`];
    const [queryParamString, setQueryParamString] = React.useState<String[]>(initialQueryParamString);

    React.useEffect(() => {
        const updatedParams = [
            `page=${page + 1}`,
            `pageSize=${rowsPerPage}`,
            ...queryParamString.filter((param) => !param.startsWith('page=') && !param.startsWith('pageSize='))
        ];
        setQueryParamString(updatedParams);
    }, [page, rowsPerPage]);

    React.useEffect(() => {
        setQueryParamString(findUpdatedParams());
    }, [search]);

    const findUpdatedParams = () => {
        const updatedParams = [...queryParamString.filter((param) => param.startsWith('page=') || param.startsWith('pageSize='))];
        Object.entries(search).forEach(([key, value]) => {
            if (key === 'created_at') {
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
            await dispatch(
                getWhatsAccountList(queries.filter((query) => !query.endsWith('=') && !query.endsWith('accountStatus=-1')).join('&'))
            );
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
        setRes(whatsAccountList!);
    }, [whatsAccountList]);
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
    const [isImportModalOpen, setIsImportModalOpen] = React.useState<boolean>(false);

    const handleAddModal = async (id?: number) => {
        if (id) {
            await axiosServices
                .get(`${envRef?.API_URL}whats/whatsAccount/view?id=${id}`, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        const selectedFormikValues: FormikValues = {
                            id: response.data.data.id,
                            account: response.data.data.account,
                            nickName: response.data.data.nickName,
                            avatar: response.data.data.avatar,
                            accountStatus: response.data.data.accountStatus,
                            isOnline: response.data.data.isOnline,
                            proxyAddress: response.data.data.proxyAddress,
                            comment: response.data.data.comment,
                            encryption: response.data.data.encryption,
                            deletedAt: response.data.data.deletedAt,
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
            setIsAddModalOpen(!isAddModalOpen);
            setFormikInitialValues(FormikInitialValuesTemplate);
        }
    };

    const handleImportModal = () => {
        setIsImportModalOpen(!isImportModalOpen);
    };

    const handleSearchClick = () => {
        fetchData(queryParamString);
        setRes(whatsAccountList!);
    };

    const handleResetClick = () => {
        setSearch(initSearchFields);
        setPage(0);
        setCreatedDate(null);
        setReset(!reset);
        setAccountStatus(-1);
        fetchData(initialQueryParamString);
    };

    const handleStatus = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const val = parseInt(event.target.value);
        setSearch({ ...search, accountStatus: val });
        setAccountStatus(val);
    };

    const [isDeleteModalOpen, setIsDeleteModalOpen] = React.useState<boolean>(false);
    const [isLoginModalOpen, setIsLoginModalOpen] = React.useState<boolean>(false);
    const [isSyncContactModalOpen, setIsSyncContactModalOpen] = React.useState<boolean>(false);
    const [isSendMessageModalOpen, setIsSendMessageModalOpen] = React.useState<boolean>(false);
    const [isSendVCardModalOpen, setIsSendVCardModalOpen] = React.useState<boolean>(false);
    const [isLogoutModalOpen, setIsLogoutModalOpen] = React.useState<boolean>(false);
    const [selectedUser, setSelectedUser] = React.useState<number[]>();
    const [isMoreControlOpen, setIsMoreControlOpen] = React.useState<number>(-1);
    const handleDeleteModal = (id: number[]) => {
        setSelectedUser(id);
        setIsDeleteModalOpen(!isDeleteModalOpen);
    };

    const handleDelete = async (ids: number[]) => {
        await axiosServices
            .post(`${envRef?.API_URL}whats/whatsAccount/delete`, { id: ids }, { headers: {} })
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
            ...queryParamString.filter((param) => !param.startsWith('page=') && !param.startsWith('pageSize='))
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
            ...queryParamString.filter((param) => !param.startsWith('page=') && !param.startsWith('pageSize='))
        ];
        fetchData(updatedParams);
    };

    const isSelected = (id: number) => selected.indexOf(id) !== -1;

    const handleLoginModal = (id: number[]) => {
        setSelectedUser(id);
        setIsLoginModalOpen(!isLoginModalOpen);
    };

    async function handleLoginUser(ids: number[]) {
        await axiosServices
            .post(`${envRef?.API_URL}whats/whats/login`, { ids: ids }, { headers: {} })
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
        setIsLoginModalOpen(!isLoginModalOpen);
    }

    function handleSyncContactModal(id?: number[]) {
        setSelected(id || []);
        setIsSyncContactModalOpen(!isSyncContactModalOpen);
    }

    function handleSendMessageModal(row?: WhatsAccountListData) {
        setSelectedRow(row || null);
        setIsSendMessageModalOpen(!isSendMessageModalOpen);
    }

    function handleSendVCardModal(row?: WhatsAccountListData) {
        setSelectedRow(row || null);
        setIsSendVCardModalOpen(!isSendVCardModalOpen);
    }

    function handleLogoutModal(row: WhatsAccountListData) {
        setSelectedRow(row);
        setIsLogoutModalOpen(!isLogoutModalOpen);
    }

    async function handleLogoutUser() {
        if (selectedRow) {
            await axiosServices
                .post(
                    `${envRef?.API_URL}whats/whats/logout`,
                    { logoutDetail: [{ account: selectedRow.account, proxy: selectedRow.proxyAddress }] },
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
            setIsLogoutModalOpen(!isLogoutModalOpen);
        }
    }

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
                        handleLoginModal={handleLoginModal}
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
                                            // onClick={(event) => handleClick(event, row.id)}
                                            // sx={{ cursor: 'pointer' }}
                                        >
                                            <Typography
                                                variant="subtitle1"
                                                sx={{ color: theme.palette.mode === 'dark' ? 'grey.600' : 'grey.900' }}
                                            >
                                                {row.account || '-'}
                                            </Typography>
                                        </TableCell>
                                        <TableCell align="center">{row.nickName || '-'}</TableCell>
                                        <TableCell align="center">
                                            {row.avatar === '/src/assets/images/logo.png' ? (
                                                <Avatar src={logo} alt={'avatar.' + row.id} />
                                            ) : row.avatar === '' ? (
                                                '-'
                                            ) : (
                                                <Avatar src={row.avatar} alt={'avatar.' + row.id} />
                                            )}
                                        </TableCell>
                                        <TableCell align="center">
                                            <Chip
                                                label={
                                                    whatsAccountStatusData.find((statusData) => statusData.value === row.accountStatus)
                                                        .label
                                                }
                                                size="small"
                                                chipcolor={
                                                    whatsAccountStatusData.find((statusData) => statusData.value === row.accountStatus)
                                                        .listClass || 'default'
                                                }
                                            />
                                        </TableCell>
                                        <TableCell align="center">
                                            {whatsAccountLoginStatusData.find((statusData) => statusData.value === row.isOnline).label}
                                        </TableCell>
                                        {handleLastActiveDate(row.lastLoginTime)}
                                        <TableCell align="center">{row.proxyAddress || '-'}</TableCell>
                                        <TableCell align="center">{row.comment || '-'}</TableCell>
                                        <TableCell align="center" className="sticky" sx={{ pr: 3, right: 0 }}>
                                            <Grid container spacing={gridSpacing}>
                                                <Grid item xs={6} md={4} lg={2}>
                                                    <IconButton
                                                        onClick={() => handleAddModal(row.id)}
                                                        color="secondary"
                                                        size="medium"
                                                        aria-label="Edit"
                                                    >
                                                        <EditIcon sx={{ fontSize: '1.3rem' }} />
                                                    </IconButton>
                                                </Grid>
                                                <Grid item xs={6} md={4} lg={2}>
                                                    <IconButton
                                                        onClick={() => handleDeleteModal([row.id])}
                                                        color="secondary"
                                                        size="medium"
                                                        aria-label="Delete"
                                                    >
                                                        <DeleteIcon sx={{ fontSize: '1.3rem' }} />
                                                    </IconButton>
                                                </Grid>
                                                <Grid item xs={6} md={4} lg={2}>
                                                    <Link to={`/whats/whatsAccount/chat/${row.id}`}>
                                                        <IconButton color="secondary" size="medium" aria-label="Chat">
                                                            <ChatBubbleTwoToneIcon sx={{ fontSize: '1.3rem' }} />
                                                        </IconButton>
                                                    </Link>
                                                </Grid>
                                                <Grid item xs={6} md={4} lg={2}>
                                                    <IconButton
                                                        onClick={() => handleSendMessageModal(row)}
                                                        color="secondary"
                                                        size="medium"
                                                        aria-label="Send Message"
                                                    >
                                                        <SendTwoToneIcon sx={{ fontSize: '1.3rem' }} />
                                                    </IconButton>
                                                </Grid>
                                                <Grid item xs={6} md={4} lg={2}>
                                                    <IconButton
                                                        onClick={() => handleSendVCardModal(row)}
                                                        color="secondary"
                                                        size="medium"
                                                        aria-label="Send Contact"
                                                    >
                                                        <ContactPhoneTwoToneIcon sx={{ fontSize: '1.3rem' }} />
                                                    </IconButton>
                                                </Grid>
                                                <Grid item xs={6} md={4} lg={2}>
                                                    <IconButton
                                                        color="secondary"
                                                        size="medium"
                                                        aria-label="More Control"
                                                        onClick={() => {
                                                            row.id === isMoreControlOpen
                                                                ? setIsMoreControlOpen(-1)
                                                                : setIsMoreControlOpen(row.id);
                                                        }}
                                                    >
                                                        <ExpandCircleDownTwoToneIcon />
                                                    </IconButton>
                                                </Grid>
                                                <Grid item xs={12} sx={{ display: isMoreControlOpen === row.id ? 'display' : 'none' }}>
                                                    <Collapse in={isMoreControlOpen === row.id}>
                                                        <Grid container spacing={1}>
                                                            <Grid item xs={12}>
                                                                <Link to={`/whats/whatsAccount/view/${row.id}`}>
                                                                    <Button
                                                                        fullWidth
                                                                        variant="outlined"
                                                                        color="secondary"
                                                                        size="medium"
                                                                        aria-label="View Detail"
                                                                    >
                                                                        {intl.formatMessage({ id: 'general.view-details' })}
                                                                    </Button>
                                                                </Link>
                                                            </Grid>
                                                            <Grid item xs={12}>
                                                                <Button
                                                                    onClick={() => handleSyncContactModal([row.id])}
                                                                    fullWidth
                                                                    variant="outlined"
                                                                    color="secondary"
                                                                    size="medium"
                                                                    aria-label="Sync Contacts"
                                                                >
                                                                    {intl.formatMessage({ id: 'general.sync-contacts' })}
                                                                </Button>
                                                            </Grid>
                                                            <Grid item xs={12}>
                                                                <Button
                                                                    onClick={() => handleLogoutModal(row)}
                                                                    fullWidth
                                                                    variant="outlined"
                                                                    color="secondary"
                                                                    size="medium"
                                                                    aria-label="Logout"
                                                                >
                                                                    {intl.formatMessage({ id: 'general.logout' })}
                                                                </Button>
                                                            </Grid>
                                                        </Grid>
                                                    </Collapse>
                                                </Grid>
                                            </Grid>
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

    async function handleImport(data?: any) {
        if (data) {
            await axiosServices
                .post(`${envRef?.API_URL}whats/whatsAccount/upload`, { list: data }, { headers: {} })
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
            setIsImportModalOpen(!isImportModalOpen);
        }
    }

    return (
        <MainCard title={<FormattedMessage id="whats.whats-account-management" />} content={false}>
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
                            onChange={(event) => handleSearch(event, 'account')}
                            placeholder={intl.formatMessage({ id: 'whats.search-account-number' })}
                            value={search.account}
                            size="small"
                            label={intl.formatMessage({ id: 'whats.account-number' })}
                        />
                    </Grid>
                    <Grid item xs={12} sm={4} md={3}>
                        <TextField
                            select
                            fullWidth
                            onChange={(event) => {
                                handleStatus(event);
                                handleSearch(event, 'accountStatus');
                            }}
                            value={accountStatus}
                            size="small"
                            label={intl.formatMessage({ id: 'whats.account-status' })}
                        >
                            {[
                                {
                                    id: -1,
                                    label: intl.formatMessage({ id: 'whats.select-account-status' }),
                                    value: -1
                                },
                                ...whatsAccountStatusData
                            ].map((option) => {
                                return (
                                    <MenuItem key={option.id} value={option.value}>
                                        {option.label}
                                    </MenuItem>
                                );
                            })}
                        </TextField>
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
                            onChange={(event) => handleSearch(event, 'proxyAddress')}
                            placeholder={intl.formatMessage({ id: 'whats.search-proxy-address' })}
                            value={search.proxyAddress}
                            size="small"
                            label={intl.formatMessage({ id: 'general.proxy-address' })}
                        />
                    </Grid>
                    {expanded && (
                        <>
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
                                    label={intl.formatMessage({ id: 'general.created-date' })}
                                />
                            </Grid>
                        </>
                    )}
                    <Grid item xs={12} sm={4} md={3} sx={{ display: 'flex', justifyContent: 'flex-end', marginLeft: 'auto' }}>
                        <AddButton onClick={() => handleAddModal()} tooltipTitle={intl.formatMessage({ id: 'whats.add-account' })} />
                        <SearchButton onClick={handleSearchClick} />
                        <ResetButton onClick={handleResetClick} />
                        <ImportButton onClick={() => handleImportModal()} tooltipTitle={intl.formatMessage({ id: 'general.import' })} />
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
                                ? intl.formatMessage({ id: 'general.add' })
                                : intl.formatMessage({ id: 'general.edit' })}
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
                                account: Yup.string()
                                    .max(255)
                                    .required(intl.formatMessage({ id: 'validation.account-number-required' }))
                            })}
                            onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                                try {
                                    await axiosServices
                                        .post(
                                            `${envRef?.API_URL}whats/whatsAccount/edit`,
                                            {
                                                id: values.id || 0,
                                                account: values.account,
                                                nickName: values.nickName,
                                                avatar: values.avatar || '',
                                                accountStatus: values.accountStatus || 0,
                                                isOnline: values.isOnline || 1,
                                                proxyAddress: values.proxyAddress || '',
                                                comment: values.comment || '',
                                                encryption: values.encryption || '',
                                                deletedAt: values.deletedAt || '',
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
                                                error={Boolean(touched.account && errors.account)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-account-number-add">
                                                    {intl.formatMessage({ id: 'whats.account-number' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-account-number-add"
                                                    type="text"
                                                    value={values.account}
                                                    name="account"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.account && errors.account && (
                                                    <FormHelperText error id="standard-weight-helper-text-account-number-add">
                                                        {errors.account.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.nickName && errors.nickName)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-nickname-add">
                                                    {intl.formatMessage({ id: 'whats.account-nickname' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-nickname-add"
                                                    type="text"
                                                    value={values.nickName}
                                                    name="nickName"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.nickName && errors.nickName && (
                                                    <FormHelperText error id="standard-weight-helper-text-nickname-add">
                                                        {errors.nickName.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                    </Grid>

                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                        <Grid item xs={12}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.avatar && errors.avatar)}
                                                sx={{ ...theme.typography.customInput }}
                                                className="MultiLineInput"
                                            >
                                                <InputLabel htmlFor="outlined-adornment-avatar-add">
                                                    {intl.formatMessage({ id: 'whats.account-avatar' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    multiline
                                                    id="outlined-adornment-avatar-add"
                                                    type="avatar"
                                                    value={values.avatar}
                                                    name="avatar"
                                                    rows={1}
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{
                                                        style: { marginTop: '4%', whiteSpace: 'break-spaces' }
                                                    }}
                                                />
                                                {touched.avatar && errors.avatar && (
                                                    <FormHelperText error id="standard-weight-helper-text-avatar-add">
                                                        {errors.avatar.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                        </Grid>
                                    </Grid>
                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                        <Grid item xs={12} sm={6}>
                                            <InputLabel htmlFor="select-account-status">
                                                {intl.formatMessage({ id: 'whats.account-status' })}
                                            </InputLabel>
                                            <Select
                                                fullWidth
                                                id="select-account-status"
                                                value={values.accountStatus}
                                                onChange={(e) => {
                                                    setFieldValue('accountStatus', e.target.value);
                                                }}
                                            >
                                                {whatsAccountStatusData.map((data: any, index: number) => {
                                                    return (
                                                        <MenuItem key={index} value={data.value}>
                                                            {data.label}
                                                        </MenuItem>
                                                    );
                                                })}
                                            </Select>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <InputLabel htmlFor="select-account-login-status">
                                                {intl.formatMessage({ id: 'general.is-online' })}
                                            </InputLabel>
                                            <Select
                                                fullWidth
                                                id="select-account-login-status"
                                                value={values.isOnline}
                                                onChange={(e) => {
                                                    setFieldValue('isOnline', e.target.value);
                                                }}
                                            >
                                                {whatsAccountLoginStatusData.map((data: any, index: number) => {
                                                    return (
                                                        <MenuItem key={index} value={data.value}>
                                                            {data.label}
                                                        </MenuItem>
                                                    );
                                                })}
                                            </Select>
                                        </Grid>
                                    </Grid>

                                    <Grid container alignItems="center" spacing={gridSpacing}>
                                        <Grid item xs={12}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.comment && errors.comment)}
                                                sx={{ ...theme.typography.customInput }}
                                                className="MultiLineInput"
                                            >
                                                <InputLabel htmlFor="outlined-adornment-remark-add">
                                                    {intl.formatMessage({ id: 'general.remark' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    multiline
                                                    id="outlined-adornment-remark-add"
                                                    type="text"
                                                    value={values.comment}
                                                    name="comment"
                                                    rows={1}
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{
                                                        style: { marginTop: '4%', whiteSpace: 'break-spaces' }
                                                    }}
                                                />
                                                {touched.comment && errors.comment && (
                                                    <FormHelperText error id="standard-weight-helper-text-remark-add">
                                                        {errors.comment.toString()}
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
            <Dialog
                id="syncContactModal"
                className="hideBackdrop"
                maxWidth="sm"
                fullWidth
                onClose={() => handleSyncContactModal()}
                open={isSyncContactModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                <Grid container spacing={gridSpacing}>
                    <Grid item sm={6} textAlign="left" sx={{ alignSelf: 'center' }}>
                        <Typography>{intl.formatMessage({ id: 'general.sync-contacts' })}</Typography>
                    </Grid>
                    <Grid item sm={6} textAlign="right">
                        <IconButton onClick={() => handleSyncContactModal()} sx={{ alignSelf: 'end' }}>
                            <CancelTwoToneIcon />
                        </IconButton>
                    </Grid>
                </Grid>

                <>
                    <Formik
                        initialValues={{ account: selected[0], contacts: '', submit: null }}
                        onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                            try {
                                await axiosServices
                                    .post(
                                        `${envRef?.API_URL}whats/whats/syncContact`,
                                        {
                                            account: values.account,
                                            contacts: values.contacts.split(',') || ''
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
                                            handleSyncContactModal();
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
                                    <Grid item xs={12}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.account && errors.account)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-sender">
                                                {intl.formatMessage({ id: 'whats.sender' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                disabled={true}
                                                id="outlined-adornment-sender"
                                                type="text"
                                                value={values.account}
                                                name="account"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.account && errors.account && (
                                                <FormHelperText error id="standard-weight-helper-text-sender">
                                                    {errors.account.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={12}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.contacts && errors.contacts)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-contacts-add">
                                                {intl.formatMessage({ id: 'whats.contacts-to-add' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-contacts-add"
                                                type="text"
                                                value={values.contacts}
                                                name="contacts"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.contacts && errors.contacts && (
                                                <FormHelperText error id="standard-weight-helper-text-contacts-add">
                                                    {errors.contacts.toString()}
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
                                                    onClick={() => handleSyncContactModal()}
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
            </Dialog>
            <Dialog
                id="sendMessageModal"
                className="hideBackdrop"
                maxWidth="sm"
                fullWidth
                onClose={() => handleSendMessageModal()}
                open={isSendMessageModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                <Grid container spacing={gridSpacing}>
                    <Grid item sm={6} textAlign="left" sx={{ alignSelf: 'center' }}>
                        <Typography>{intl.formatMessage({ id: 'general.send-message' })}</Typography>
                    </Grid>
                    <Grid item sm={6} textAlign="right">
                        <IconButton onClick={() => handleSendMessageModal()} sx={{ alignSelf: 'end' }}>
                            <CancelTwoToneIcon />
                        </IconButton>
                    </Grid>
                </Grid>

                <>
                    <Formik
                        initialValues={{ sender: selectedRow?.account, receiver: '', textMsg: '', submit: null }}
                        validationSchema={Yup.object().shape({
                            sender: Yup.string()
                                .max(255)
                                .required(intl.formatMessage({ id: 'validation.sender-required' })),
                            receiver: Yup.string()
                                .max(255)
                                .required(intl.formatMessage({ id: 'validation.receiver-required' })),
                            textMsg: Yup.string()
                                .max(255)
                                .required(intl.formatMessage({ id: 'validation.text-message-required' }))
                        })}
                        onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                            try {
                                await axiosServices
                                    .post(
                                        `${envRef?.API_URL}whats/whats/sendMsg`,
                                        {
                                            sender: values.sender,
                                            receiver: values.receiver,
                                            textMsg: [values.textMsg]
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
                                            handleSendMessageModal();
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
                                    <Grid item xs={12}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.sender && errors.sender)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-sender">
                                                {intl.formatMessage({ id: 'whats.sender' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                disabled={true}
                                                id="outlined-adornment-sender"
                                                type="text"
                                                value={values.sender}
                                                name="sender"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.sender && errors.sender && (
                                                <FormHelperText error id="standard-weight-helper-text-sender">
                                                    {errors.sender.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={12}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.receiver && errors.receiver)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-receiver">
                                                {intl.formatMessage({ id: 'whats.receiver' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-receiver"
                                                type="text"
                                                value={values.receiver}
                                                name="receiver"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.receiver && errors.receiver && (
                                                <FormHelperText error id="standard-weight-helper-text-receiver">
                                                    {errors.receiver.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={12}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.textMsg && errors.textMsg)}
                                            sx={{ ...theme.typography.customInput }}
                                            className="MultiLineInput"
                                        >
                                            <InputLabel htmlFor="outlined-adornment-message">
                                                {intl.formatMessage({ id: 'whats.text-message' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                multiline
                                                id="outlined-adornment-message"
                                                type="text"
                                                value={values.textMsg}
                                                name="textMsg"
                                                rows={1}
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{
                                                    style: { marginTop: '4%', whiteSpace: 'break-spaces' }
                                                }}
                                            />
                                            {touched.textMsg && errors.textMsg && (
                                                <FormHelperText error id="standard-weight-helper-text-message">
                                                    {errors.textMsg.toString()}
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
                                                    onClick={() => handleSendMessageModal()}
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
            </Dialog>
            <Dialog
                id="sendVCardModal"
                className="hideBackdrop"
                maxWidth="sm"
                fullWidth
                onClose={() => handleSendVCardModal()}
                open={isSendVCardModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '1.5rem 2rem' }, backgroundColor: '#f5f5f5' }}
            >
                <Grid container spacing={gridSpacing}>
                    <Grid item sm={6} textAlign="left" sx={{ alignSelf: 'center' }}>
                        <Typography>{intl.formatMessage({ id: 'general.send-v-card' })}</Typography>
                    </Grid>
                    <Grid item sm={6} textAlign="right">
                        <IconButton onClick={() => handleSendVCardModal()} sx={{ alignSelf: 'end' }}>
                            <CancelTwoToneIcon />
                        </IconButton>
                    </Grid>
                </Grid>

                <>
                    <Formik
                        initialValues={{
                            sender: selectedRow?.account,
                            receiver: '',
                            version: '',
                            prodid: '',
                            fn: '',
                            tel: '',
                            org: '',
                            xwabizname: '',
                            end: '',
                            /* Shown in grata but not passed into API */
                            // displayName: '',
                            family: '',
                            prefixes: '',
                            language: '',
                            submit: null
                        }}
                        validationSchema={Yup.object().shape({
                            sender: Yup.string()
                                .max(255)
                                .required(intl.formatMessage({ id: 'validation.sender-required' })),
                            receiver: Yup.string()
                                .max(255)
                                .required(intl.formatMessage({ id: 'validation.receiver-required' })),
                            tel: Yup.string()
                                .max(255)
                                .required(intl.formatMessage({ id: 'validation.telephone-required' }))
                        })}
                        onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                            try {
                                await axiosServices
                                    .post(
                                        `${envRef?.API_URL}whats/whats/sendVcardMsg`,
                                        {
                                            sender: values.sender,
                                            receiver: values.receiver,
                                            vcard: {
                                                end: values.end,
                                                family: values.family,
                                                fn: values.fn,
                                                language: values.language,
                                                org: values.org,
                                                prefixes: values.prefixes,
                                                prodid: values.prodid,
                                                tel: values.tel,
                                                version: values.version,
                                                // displayName: values.displayName,
                                                xwabizname: values.xwabizname
                                            }
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
                                            handleSendVCardModal();
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
                                    <Grid item xs={12}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.sender && errors.sender)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-sender">
                                                {intl.formatMessage({ id: 'whats.sender' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                disabled={true}
                                                id="outlined-adornment-sender"
                                                type="text"
                                                value={values.sender}
                                                name="sender"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.sender && errors.sender && (
                                                <FormHelperText error id="standard-weight-helper-text-sender">
                                                    {errors.sender.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={12}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.receiver && errors.receiver)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-receiver">
                                                {intl.formatMessage({ id: 'whats.receiver' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-receiver"
                                                type="text"
                                                value={values.receiver}
                                                name="receiver"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.receiver && errors.receiver && (
                                                <FormHelperText error id="standard-weight-helper-text-receiver">
                                                    {errors.receiver.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>

                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.version && errors.version)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-version">
                                                {intl.formatMessage({ id: 'general.version' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-version"
                                                type="text"
                                                value={values.version}
                                                name="version"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.version && errors.version && (
                                                <FormHelperText error id="standard-weight-helper-text-version">
                                                    {errors.version.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.prodid && errors.prodid)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-prod-id">
                                                {intl.formatMessage({ id: 'whats.v-card-app' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-prod-id"
                                                type="text"
                                                value={values.prodid}
                                                name="prodid"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.prodid && errors.prodid && (
                                                <FormHelperText error id="standard-weight-helper-text-prod-id">
                                                    {errors.prodid.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>

                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.fn && errors.fn)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-name">
                                                {intl.formatMessage({ id: 'general.name' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-name"
                                                type="text"
                                                value={values.fn}
                                                name="fn"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.fn && errors.fn && (
                                                <FormHelperText error id="standard-weight-helper-text-name">
                                                    {errors.fn.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.tel && errors.tel)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-tel">
                                                {intl.formatMessage({ id: 'general.telephone' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-tel"
                                                type="text"
                                                value={values.tel}
                                                name="tel"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.tel && errors.tel && (
                                                <FormHelperText error id="standard-weight-helper-text-tel">
                                                    {errors.tel.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>

                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.org && errors.org)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-org">
                                                {intl.formatMessage({ id: 'general.organization' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-org"
                                                type="text"
                                                value={values.org}
                                                name="org"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.org && errors.org && (
                                                <FormHelperText error id="standard-weight-helper-text-org">
                                                    {errors.org.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.xwabizname && errors.xwabizname)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-custom-name">
                                                {intl.formatMessage({ id: 'general.custom-name' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-custom-name"
                                                type="text"
                                                value={values.xwabizname}
                                                name="xwabizname"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                                placeholder={intl.formatMessage({ id: 'general.custom-name-tooltip' })}
                                            />
                                            {touched.xwabizname && errors.xwabizname && (
                                                <FormHelperText error id="standard-weight-helper-text-custom-name">
                                                    {errors.xwabizname.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>

                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.end && errors.end)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-end">
                                                {intl.formatMessage({ id: 'whats.v-card-ending' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-end"
                                                type="text"
                                                value={values.end}
                                                name="end"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.end && errors.end && (
                                                <FormHelperText error id="standard-weight-helper-text-end">
                                                    {errors.end.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>
                                    {/* <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.displayName && errors.displayName)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-display-name">
                                                {intl.formatMessage({ id: 'general.display-name' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-display-name"
                                                type="text"
                                                value={values.displayName}
                                                name="displayName"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                placeholder={intl.formatMessage({ id: 'general.display-name-tooltip' })}
                                                inputProps={{}}
                                            />
                                            {touched.displayName && errors.displayName && (
                                                <FormHelperText error id="standard-weight-helper-text-display-name">
                                                    {errors.displayName.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid> */}
                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.family && errors.family)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-family">
                                                {intl.formatMessage({ id: 'general.family' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-family"
                                                type="text"
                                                value={values.family}
                                                name="family"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.family && errors.family && (
                                                <FormHelperText error id="standard-weight-helper-text-family">
                                                    {errors.family.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>

                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.prefixes && errors.prefixes)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-name-prefix">
                                                {intl.formatMessage({ id: 'general.name-prefix' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-name-prefix"
                                                type="text"
                                                value={values.prefixes}
                                                name="prefixes"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                placeholder={intl.formatMessage({ id: 'general.name-prefix-tooltip' })}
                                                inputProps={{}}
                                            />
                                            {touched.prefixes && errors.prefixes && (
                                                <FormHelperText error id="standard-weight-helper-text-name-prefix">
                                                    {errors.prefixes.toString()}
                                                </FormHelperText>
                                            )}
                                        </FormControl>
                                    </Grid>
                                    <Grid item xs={12} sm={6}>
                                        <FormControl
                                            fullWidth
                                            error={Boolean(touched.language && errors.language)}
                                            sx={{ ...theme.typography.customInput }}
                                        >
                                            <InputLabel htmlFor="outlined-adornment-language">
                                                {intl.formatMessage({ id: 'general.language' })}
                                            </InputLabel>
                                            <OutlinedInput
                                                id="outlined-adornment-language"
                                                type="text"
                                                value={values.language}
                                                name="language"
                                                onBlur={handleBlur}
                                                onChange={handleChange}
                                                inputProps={{}}
                                            />
                                            {touched.language && errors.language && (
                                                <FormHelperText error id="standard-weight-helper-text-language">
                                                    {errors.language.toString()}
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
                                                    onClick={() => handleSendVCardModal()}
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
            </Dialog>
            <GeneralDialog
                confirmFunction={() => handleDelete(selectedUser!)}
                isOpen={isDeleteModalOpen}
                setIsOpen={() => setIsDeleteModalOpen(!isDeleteModalOpen)}
                id="deleteConfirmModal"
                type="warning"
                content={intl.formatMessage({ id: 'general.confirm-delete-content' })}
            />
            <GeneralDialog
                confirmFunction={() => handleLoginUser(selectedUser!)}
                isOpen={isLoginModalOpen}
                setIsOpen={() => setIsLoginModalOpen(!isLoginModalOpen)}
                id="loginConfirmModal"
                type="warning"
                content={intl.formatMessage({ id: 'general.confirm-login-content' })}
            />
            <GeneralDialog
                confirmFunction={() => handleLogoutUser()}
                isOpen={isLogoutModalOpen}
                setIsOpen={() => setIsLogoutModalOpen(!isLogoutModalOpen)}
                id="logoutConfirmModal"
                type="warning"
                content={intl.formatMessage({ id: 'general.confirm-logout-content' })}
            />
            <ImportDialog
                isOpen={isImportModalOpen}
                setIsOpen={() => setIsImportModalOpen(!isImportModalOpen)}
                id="importModal"
                checkHeader={WhatsAccountHeaderValidation}
                importTemplatePath="/documents/whatsAccountTemplate.xlsx"
                importTemplateFileName="whatsAccountTemplate.xlsx"
                confirmFunction={handleImport}
                valueFields={['account', 'identify', 'privateKey', 'privateMsgKey', 'publicKey', 'publicMsgKey']}
            />
        </MainCard>
    );
};

export default WhatsAccount;
