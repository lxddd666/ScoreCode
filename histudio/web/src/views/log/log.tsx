import * as React from 'react';
import { useNavigate } from 'react-router-dom';

// material-ui
import { useTheme } from '@mui/material/styles';
import {
    Box,
    // Button,
    CardContent,
    Checkbox,
    // Dialog,
    // Divider,
    // FormControl,
    // FormHelperText,
    Grid,
    IconButton,
    InputAdornment,
    // InputLabel,
    MenuItem,
    // OutlinedInput,
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

// ui-components
import Chip from 'ui-component/extended/Chip';
import MainCard from 'ui-component/cards/MainCard';
import ExpandButton from 'ui-component/searchbar/ExpandButton';
import ResetButton from 'ui-component/searchbar/ResetButton';
import SearchButton from 'ui-component/searchbar/SearchButton';
import ExportButton from 'ui-component/searchbar/ExportButton';
// import AddButton from 'ui-component/searchbar/AddButton';
// import ExpandableSelect from 'ui-component/general/ExpandableSelect';
import GeneralDialog from 'ui-component/general/GeneralDialog';
import CustomDateTimeRangePicker from 'ui-component/general/CustomDateTimeRangePicker';

// table assets
import DeleteIcon from '@mui/icons-material/DeleteTwoTone';
import SearchIcon from '@mui/icons-material/SearchTwoTone';
import EditIcon from '@mui/icons-material/EditTwoTone';
// import ToggleOnIcon from '@mui/icons-material/ToggleOnTwoTone';
// import ToggleOffIcon from '@mui/icons-material/ToggleOffOutlined';
import FolderOffTwoToneIcon from '@mui/icons-material/FolderOffTwoTone';
import FirstPageIcon from '@mui/icons-material/FirstPage';
import KeyboardArrowLeft from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRight from '@mui/icons-material/KeyboardArrowRight';
import LastPageIcon from '@mui/icons-material/LastPage';
// import CancelTwoToneIcon from '@mui/icons-material/CancelTwoTone';

import { GetComparator, EnhancedTableHeadProps, HeadCell, ArrangementOrder, KeyedObject } from 'types';
// import logo from 'assets/images/general/logo.png';

// locale
import { FormattedMessage, useIntl } from 'react-intl';

// API
import axiosServices from 'utils/axios';

// env
import envRef from 'environment';
import { isEqual } from 'lodash';

// third party
// import * as Yup from 'yup';
// import { Formik, FormikValues } from 'formik';

// project imports
import { FlatOption } from 'types/option';
import { Log } from 'types/log';
import { openSnackbar } from 'store/slices/snackbar';
import { ResponseList } from 'types/response';
import { useDispatch, useSelector } from 'store';
import { getLogList } from 'store/slices/log';
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';
import { gridSpacing } from 'store/constant';

// table header options
const headCells: HeadCell[] = [
    {
        id: 'module',
        numeric: false,
        label: 'general-log.module',
        align: 'left'
    },
    {
        id: 'user',
        numeric: false,
        label: 'general-log.user',
        align: 'center'
    },
    {
        id: 'method',
        numeric: false,
        label: 'general-log.request-type',
        align: 'center'
    },
    {
        id: 'url',
        numeric: false,
        label: 'general-log.request-path',
        align: 'center'
    },
    {
        id: 'ip',
        numeric: true,
        label: 'general-log.access-ip',
        align: 'center'
    },
    {
        id: 'status',
        numeric: false,
        label: 'general-log.status',
        align: 'center'
    },
    {
        id: 'takeUpTime',
        numeric: false,
        label: 'general-log.response-time',
        align: 'center'
    },
    {
        id: 'createdAt',
        numeric: false,
        label: 'general.created-date',
        align: 'center'
    }
];

// ==============================|| TABLE HEADER ||============================== //

interface CustomListEnhancedTableHeadProps extends EnhancedTableHeadProps {
    selected: number[];
    handleDeleteModal: (selected: number[]) => void;
}

const EnhancedTableHead = ({
    onSelectAllClick,
    order,
    orderBy,
    numSelected,
    rowCount,
    onRequestSort,
    selected,
    handleDeleteModal
}: CustomListEnhancedTableHeadProps) => {
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
};

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
        )}
    </Toolbar>
);

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

function stableSort(array: Log[], comparator: (a: Log, b: Log) => number) {
    const stabilizedThis = array.map((el: Log, index: number) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0] as Log, b[0] as Log);
        if (order !== 0) return order;
        return (a[1] as number) - (b[1] as number);
    });
    return stabilizedThis.map((el) => el[0]);
}
// END table sort

// ==============================|| LOG LIST ||============================== //

const GeneralLog = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);
    const navigate = useNavigate();

    type SearchFields = {
        member_id: string;
        url: string;
        ip: string;
        method: string;
        created_at: Date[];
        take_up_time: number;
        error_code: number;
    };
    const initSearchFields: SearchFields = { member_id: '', url: '', ip: '', method: '', created_at: [], take_up_time: 0, error_code: 1 };

    const [method, setMethod] = React.useState<number>(0);

    // async function getOptions() {
    //     // Get position list
    //     await axiosServices
    //         .get(`${envRef?.API_URL}admin/post/list`, { headers: {} })
    //         .then(function (response) {
    //             if (response?.data?.code === 0) {
    //                 // setFlatPostOptions(response.data.data.list);
    //             }
    //         })
    //         .catch(function (error) {
    //             dispatch(
    //                 openSnackbar({
    //                     open: true,
    //                     message: error?.message || defaultErrorMessage,
    //                     variant: 'alert',
    //                     alert: {
    //                         color: 'error'
    //                     },
    //                     close: false,
    //                     anchorOrigin: {
    //                         vertical: 'top',
    //                         horizontal: 'center'
    //                     }
    //                 })
    //             );
    //         });

    //     // Get role list
    //     await axiosServices
    //         .get(`${envRef?.API_URL}admin/role/list`, { headers: {} })
    //         .then(function (response) {
    //             if (response?.data?.code === 0) {
    //                 // setRoleOptions(response.data.data.list);
    //             }
    //         })
    //         .catch(function (error) {
    //             dispatch(
    //                 openSnackbar({
    //                     open: true,
    //                     message: error?.message || defaultErrorMessage,
    //                     variant: 'alert',
    //                     alert: {
    //                         color: 'error'
    //                     },
    //                     close: false,
    //                     anchorOrigin: {
    //                         vertical: 'top',
    //                         horizontal: 'center'
    //                     }
    //                 })
    //             );
    //         });

    //     // Get dept list
    //     await axiosServices
    //         .get(`${envRef?.API_URL}admin/dept/option`, { headers: {} })
    //         .then(function (response) {
    //             if (response?.data?.code === 0) {
    //                 // setDeptOptions(response.data.data.list);
    //             }
    //         })
    //         .catch(function (error) {
    //             dispatch(
    //                 openSnackbar({
    //                     open: true,
    //                     message: error?.message || defaultErrorMessage,
    //                     variant: 'alert',
    //                     alert: {
    //                         color: 'error'
    //                     },
    //                     close: false,
    //                     anchorOrigin: {
    //                         vertical: 'top',
    //                         horizontal: 'center'
    //                     }
    //                 })
    //             );
    //         });
    // }

    // React.useEffect(() => {
    //     getOptions();
    // }, []);

    const [order, setOrder] = React.useState<ArrangementOrder>('asc');
    const [orderBy, setOrderBy] = React.useState<string>('id');
    const [selected, setSelected] = React.useState<number[]>([]);
    const [page, setPage] = React.useState<number>(0);
    const [totalCount, setTotalCount] = React.useState<number>(0);
    const [pageCount, setPageCount] = React.useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = React.useState<number>(10);
    const [search, setSearch] = React.useState<SearchFields>(initSearchFields);
    const [rows, setRows] = React.useState<Log[]>();
    const [res, setRes] = React.useState<ResponseList<Log>>();
    const { logList } = useSelector((state) => state.log);
    const [isDataListEmpty, setIsDataListEmpty] = React.useState<boolean>(true);
    // const scriptedRef = useScriptRef();
    // const [showPassword, setShowPassword] = React.useState(false);
    // const FormikInitialValuesTemplate: FormikValues = {
    //     id: 0,
    //     realName: '',
    //     username: '',
    //     roleId: 1,
    //     deptId: 100,
    //     postIds: [],
    //     password: '',
    //     mobile: '',
    //     email: '',
    //     sex: 1,
    //     status: 1,
    //     remark: '',
    //     submit: null
    // };
    // const [formikInitialValues, setFormikInitialValues] = React.useState<FormikValues>(FormikInitialValuesTemplate);
    const [createdDate, setCreatedDate] = React.useState<Date[] | null>(null);
    const [reset, setReset] = React.useState<boolean>(false);

    const initialQueryParamString: string[] = [`page=${page + 1}`, `pageSize=${rowsPerPage}`];
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
        const updatedParams = [
            ...queryParamString.filter((param) => param.startsWith('page=') || param.startsWith('pageSize=') || param.startsWith('roleId='))
        ];
        Object.entries(search).forEach(([key, value]) => {
            // if (key === 'roleId') {
            //     setRole(+value);
            // } else
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
                getLogList(
                    queries
                        .filter((query) => {
                            return query === 'take_up_time=0' || query === 'error_code=1' || query.endsWith('=') ? false : true;
                        })
                        .join('&')
                )
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
        setRes(logList!);
    }, [logList]);
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

    // const [isAddModalOpen, setIsAddModalOpen] = React.useState<boolean>(false);

    const navigateView = (id?: number) => {
        navigate(`/log/view/${id}`);
    };

    const handleSearchClick = () => {
        fetchData(queryParamString);
        setRes(logList!);
    };

    const handleResetClick = () => {
        setSearch(initSearchFields);
        setCreatedDate(null);
        setReset(!reset);
        setMethod(0);
        // setStatus(0);
        // setRole(-1);
        fetchData(initialQueryParamString);
    };

    const handleExportClick = () => {
        ExportData(queryParamString);
        // setRes(logList!);
    };

    const methodOptions: FlatOption[] = [
        {
            id: 0,
            name: intl.formatMessage({ id: 'general-log.select-request-type' })
        },
        {
            id: 1,
            name: 'GET'
        },
        {
            id: 2,
            name: 'POST'
        }
    ];

    const takeUpTimeOptions: FlatOption[] = [
        {
            id: 0,
            name: intl.formatMessage({ id: 'general-log.select-take-up-time' })
        },
        {
            id: 50,
            name: '50ms'
        },
        {
            id: 100,
            name: '100ms'
        },
        {
            id: 200,
            name: '200ms'
        },
        {
            id: 500,
            name: '500ms'
        }
    ];

    const errorCodeOptions: FlatOption[] = [
        {
            id: 1,
            name: intl.formatMessage({ id: 'general-log.select-error-code' })
        },
        {
            id: 0,
            name: '0 ' + intl.formatMessage({ id: 'general-log.error-code-success' })
        },
        {
            id: -1,
            name: '-1 ' + intl.formatMessage({ id: 'general-log.error-code-failed' })
        }
    ];

    const [isDeleteModalOpen, setIsDeleteModalOpen] = React.useState<boolean>(false);
    const [selectedIdsToDelete, setSelectedIdsToDelete] = React.useState<number[]>();
    const handleDeleteModal = (id: number[]) => {
        setSelectedIdsToDelete(id);
        setIsDeleteModalOpen(!isDeleteModalOpen);
    };

    const handleDelete = (ids: number[]) => {
        axiosServices
            .post(`${envRef?.API_URL}admin/log/delete`, { id: ids }, { headers: {} })
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

    const ExportData = (queries: String[]) => {
        axiosServices
            .get(
                `${envRef?.API_URL}/admin/log/export?${queries
                    .filter((query) => {
                        return query === 'take_up_time=0' || query === 'error_code=1' || query.endsWith('=') ? false : true;
                    })
                    .join('&')}`,
                { responseType: 'blob', headers: {} }
            )
            .then(function (response) {
                if (response) {
                    const outputFilename = `Log_${Date.now()}.xlsx`;

                    // If you want to download file automatically using link attribute.
                    const url = URL.createObjectURL(new Blob([response.data]));
                    const link = document.createElement('a');
                    link.href = url;
                    link.setAttribute('download', outputFilename);
                    document.body.appendChild(link);
                    link.click();

                    // OR you can save/write file locally.
                    // fs.writeFileSync(outputFilename, response.data);

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
                } else {
                    dispatch(
                        openSnackbar({
                            open: true,
                            // message: response?.data?.message || defaultErrorMessage,
                            message: defaultErrorMessage,
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

    const handleSearch = (event: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement> | undefined, searchParam: string) => {
        const val = event?.target.value;

        if (searchParam === 'method') {
            const methodId = parseInt(val!);
            const name = methodId > 0 ? methodOptions.find((m) => m.id === methodId)?.name : '';

            setSearch({ ...search, [searchParam]: name! });
            setMethod(methodId);
        } else if (['take_up_time', 'error_code'].indexOf(searchParam) > -1) {
            setSearch({ ...search, [searchParam]: val });
        } else {
            setSearch({ ...search, [searchParam]: val || '' });
        }

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
        //     setRows(logList!.data!.list!);
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

    // const mobileNumberRegExp = /^(\+?\d{0,4})?\s?-?\s?(\(?\d{3}\)?)\s?-?\s?(\(?\d{3}\)?)\s?-?\s?(\(?\d{4}\)?)?$/;
    const defaultErrorMessage = intl.formatMessage({ id: 'auth-register.default-error' });

    const renderTable = () => {
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
                                        <TableCell align="center">
                                            {
                                                <Chip
                                                    label={row.module}
                                                    size="small"
                                                    chipcolor={row.module.toLowerCase() == 'admin' ? 'info' : 'success'}
                                                />
                                            }
                                        </TableCell>
                                        <TableCell align="center">
                                            {row.memberName || '-'}
                                            {row.memberId > 0 ? `(${row.memberId})` : ''}
                                        </TableCell>
                                        <TableCell align="center">{row.method || '-'}</TableCell>
                                        <TableCell align="center">{row.url || '-'}</TableCell>
                                        <TableCell align="center">{row.ip || '-'}</TableCell>
                                        <TableCell align="center">
                                            {
                                                <Chip
                                                    title={`${row.errorMsg}(${row.errorCode})`}
                                                    label={`${row.errorMsg}(${row.errorCode})`}
                                                    size="small"
                                                    chipcolor={row.errorCode === 0 ? 'success' : 'orange'}
                                                    maxWidth="150px"
                                                />
                                            }
                                        </TableCell>
                                        <TableCell align="center">
                                            {row.takeUpTime || '-'}
                                            {' ms'}
                                        </TableCell>
                                        <TableCell align="center">{row.createdAt || '-'}</TableCell>
                                        <TableCell align="center" className="sticky" sx={{ pr: 3, right: 0 }}>
                                            <IconButton
                                                onClick={() => navigateView(row.id)}
                                                color="secondary"
                                                size="medium"
                                                aria-label="Edit"
                                            >
                                                <EditIcon sx={{ fontSize: '1.3rem' }} />
                                            </IconButton>
                                            <IconButton onClick={() => {}} color="secondary" size="medium" aria-label="Delete">
                                                <DeleteIcon
                                                    onClick={() => {
                                                        handleDeleteModal([row.id]);
                                                    }}
                                                    sx={{ fontSize: '1.3rem' }}
                                                />
                                            </IconButton>
                                        </TableCell>
                                    </TableRow>
                                );
                            })
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
        );
    };

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
        <MainCard title={<FormattedMessage id="general-log.general-log" />} content={false}>
            <CardContent>
                <Grid container alignItems="center" spacing={2} justifyContent="flex-start">
                    <Grid item xs={12} sm={3}>
                        <TextField
                            fullWidth
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon fontSize="small" />
                                    </InputAdornment>
                                )
                            }}
                            onChange={(event) => {
                                handleSearch(event, 'member_id');
                            }}
                            placeholder={intl.formatMessage({ id: 'general-log.user' })}
                            value={search.member_id}
                            size="small"
                            label={intl.formatMessage({ id: 'general-log.user' })}
                        />
                    </Grid>
                    <Grid item xs={12} sm={3}>
                        <TextField
                            fullWidth
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon fontSize="small" />
                                    </InputAdornment>
                                )
                            }}
                            onChange={(event) => handleSearch(event, 'url')}
                            placeholder={intl.formatMessage({ id: 'general-log.request-path' })}
                            value={search.url}
                            size="small"
                            label={intl.formatMessage({ id: 'general-log.request-path' })}
                        />
                    </Grid>
                    <Grid item xs={12} sm={3}>
                        <TextField
                            fullWidth
                            InputProps={{
                                startAdornment: (
                                    <InputAdornment position="start">
                                        <SearchIcon fontSize="small" />
                                    </InputAdornment>
                                )
                            }}
                            onChange={(event) => handleSearch(event, 'ip')}
                            placeholder={intl.formatMessage({ id: 'general-log.access-ip' })}
                            value={search.ip}
                            size="small"
                            label={intl.formatMessage({ id: 'general-log.access-ip' })}
                        />
                    </Grid>
                    {expanded && (
                        <>
                            <Grid item xs={12} sm={3}>
                                <TextField
                                    select
                                    fullWidth
                                    onChange={(event) => {
                                        // handleMethod(event)
                                        handleSearch(event, 'method');
                                    }}
                                    value={method}
                                    size="small"
                                    label={intl.formatMessage({ id: 'general-log.request-type' })}
                                >
                                    {methodOptions.map((option) => {
                                        return (
                                            <MenuItem key={option.id} value={option.id}>
                                                {option.name}
                                            </MenuItem>
                                        );
                                    })}
                                </TextField>
                            </Grid>
                            <Grid item xs={12} sm={6}>
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
                                    label={intl.formatMessage({ id: 'general-log.access-date' })}
                                    viewTime={true}
                                />
                            </Grid>
                            <Grid item xs={12} sm={3}>
                                <TextField
                                    select
                                    fullWidth
                                    onChange={(event) => {
                                        handleSearch(event, 'take_up_time');
                                    }}
                                    value={search.take_up_time}
                                    size="small"
                                    label={intl.formatMessage({ id: 'general-log.response-time' })}
                                >
                                    {takeUpTimeOptions.map((option) => {
                                        return (
                                            <MenuItem key={option.id} value={option.id}>
                                                {option.name}
                                            </MenuItem>
                                        );
                                    })}
                                </TextField>
                            </Grid>
                            <Grid item xs={12} sm={3}>
                                <TextField
                                    select
                                    fullWidth
                                    onChange={(event) => {
                                        handleSearch(event, 'error_code');
                                    }}
                                    value={search.error_code}
                                    size="small"
                                    label={intl.formatMessage({ id: 'general-log.status' })}
                                >
                                    {errorCodeOptions.map((option) => {
                                        return (
                                            <MenuItem key={option.id} value={option.id}>
                                                {option.name}
                                            </MenuItem>
                                        );
                                    })}
                                </TextField>
                            </Grid>
                        </>
                    )}
                    <Grid item xs={12} sm={3} md={3} sx={{ display: 'flex', justifyContent: 'flex-end', marginLeft: 'auto' }}>
                        {/* <AddButton onClick={() => handleAddModal()} tooltipTitle={intl.formatMessage({ id: 'user.add-user' })} /> */}
                        <ExportButton onClick={handleExportClick} />
                        <SearchButton onClick={handleSearchClick} />
                        <ResetButton onClick={handleResetClick} />
                        <ExpandButton onClick={handleExpandClick} transformValue={expanded ? 'rotate(180deg)' : 'rotate(0deg)'} />
                    </Grid>
                </Grid>
            </CardContent>

            {renderTable()}

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
            <GeneralDialog
                confirmFunction={() => handleDelete(selectedIdsToDelete!)}
                isOpen={isDeleteModalOpen}
                setIsOpen={() => setIsDeleteModalOpen(!isDeleteModalOpen)}
                id="delete-confirm-modal"
                type="warning"
                content={intl.formatMessage({ id: 'general.confirm-delete-content' })}
            />
        </MainCard>
    );
};

export default GeneralLog;
