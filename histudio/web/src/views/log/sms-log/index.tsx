import * as React from 'react';
// locale
import { FormattedMessage, useIntl } from 'react-intl';

// material-ui
import { useTheme } from '@mui/material/styles';
import {
    Box,
    CardContent,
    Checkbox,
    Grid,
    IconButton,
    InputAdornment,
    MenuItem,
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
    Typography,
} from '@mui/material';
import { visuallyHidden } from '@mui/utils';

// project imports
import { FlatOption } from 'types/option';
import { SmslogListData } from 'types/smslog';
import { ResponseList } from 'types/response';
import { useDispatch, useSelector } from 'store';
import { getSmslogList } from 'store/slices/smslog';
// assets
import DeleteIcon from '@mui/icons-material/DeleteTwoTone';
import SearchIcon from '@mui/icons-material/SearchTwoTone';
import FolderOffTwoToneIcon from '@mui/icons-material/FolderOffTwoTone';
import FirstPageIcon from '@mui/icons-material/FirstPage';
import KeyboardArrowLeft from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRight from '@mui/icons-material/KeyboardArrowRight';
import LastPageIcon from '@mui/icons-material/LastPage';

import { GetComparator, EnhancedTableHeadProps, HeadCell, ArrangementOrder, KeyedObject } from 'types';

// ui-components
import Chip from 'ui-component/extended/Chip';
import MainCard from 'ui-component/cards/MainCard';
import ExpandButton from 'ui-component/searchbar/ExpandButton';
import ResetButton from 'ui-component/searchbar/ResetButton';
import SearchButton from 'ui-component/searchbar/SearchButton';
import CustomDateTimeRangePicker from 'ui-component/general/CustomDateTimeRangePicker';
import GeneralDialog from 'ui-component/general/GeneralDialog';

// API
import axiosServices from 'utils/axios';
// env
import envRef from 'environment';
import { isEqual } from 'lodash';
import { gridSpacing } from 'store/constant';
import { openSnackbar } from 'store/slices/snackbar';
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';

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

function stableSort(array: SmslogListData[], comparator: (a: SmslogListData, b: SmslogListData) => number) {
    const stabilizedThis = array.map((el: SmslogListData, index: number) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0] as SmslogListData, b[0] as SmslogListData);
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
        label: 'general.id',
        align: 'center'
    },
    {
        id: 'event',
        numeric: false,
        label: 'smslog.event-template',
        align: 'center'
    },
    {
        id: 'mobile',
        numeric: false,
        label: 'smslog.phone-num',
        align: 'center'
    },
    {
        id: 'code',
        numeric: false,
        label: 'smslog.verify-code-sms-con',
        align: 'center'
    },
    {
        id: 'times',
        numeric: false,
        label: 'smslog.num-of-verify',
        align: 'center'
    },
    {
        id: 'ip',
        numeric: false,
        label: 'smslog.sender-ip',
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
        label: 'whats.send-time',
        align: 'center'
    },
    {
        id: 'updatedAt',
        numeric: false,
        label: 'loginlog.updated-at',
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
                    <TableCell
                        width={'15%'}
                        sortDirection={false}
                        align="center"
                        className="sticky"
                        sx={{ pr: 3, right: '0', minWidth: '200px' }}
                    >
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
        ) :(
            <Typography
                sx={{
                    color: (theme) => theme.palette.primary.light
                }}
                variant="h6"
                id="tableTitle"
            >
                <FormattedMessage id="position.position-management" />
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

// ==============================|| DATA LIST ||============================== //

const SmsLog = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);

    type SearchFields = {
        mobile: string;
        event: string;
        ip: string;
        status: number;
        created_at: Date[];  
    };
    const initSearchFields: SearchFields = { mobile: '', event: '', ip: '', status:0, created_at: [] };
    const [status, setStatus] = React.useState<number>(0);
    const [method, setMethod] = React.useState<string>("");

    const [order, setOrder] = React.useState<ArrangementOrder>('asc');
    const [orderBy, setOrderBy] = React.useState<string>('id');
    const [selected, setSelected] = React.useState<number[]>([]);
    const [page, setPage] = React.useState<number>(0);
    const [totalCount, setTotalCount] = React.useState<number>(0);
    const [pageCount, setPageCount] = React.useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = React.useState<number>(10);

    const [search, setSearch] = React.useState<SearchFields>(initSearchFields);
    
    const [rows, setRows] = React.useState<SmslogListData[]>();
    const [res, setRes] = React.useState<ResponseList<SmslogListData>>();
    const { smslogList } = useSelector((state) => state.smslog);
    
    const [isDataListEmpty, setIsDataListEmpty] = React.useState<boolean>(true);
    const [group, setgroup] = React.useState<number>(0);
    
    const [createdDate, setCreatedDate] = React.useState<Date[] | null>(null);
    const [reset, setReset] = React.useState<boolean>(false);

    const initialQueryParamString: string[] = [`page=1`, `pageSize=${rowsPerPage}`, `groupId=0`];
    const [queryParamString, setQueryParamString] = React.useState<String[]>(initialQueryParamString);

    React.useEffect(() => {
        const updatedParams = [
            `page=${page + 1}`,
            `pageSize=${rowsPerPage}`,
            `groupId=${group}`,
            ...queryParamString.filter((param) => !param.startsWith('page=') && !param.startsWith('pageSize=') && !param.startsWith('groupId='))
        ];
        setQueryParamString(updatedParams);
    }, [page, rowsPerPage, group]);

    React.useEffect(() => {
        setQueryParamString(findUpdatedParams());
    }, [search]);

    const findUpdatedParams = () => {
        const updatedParams = [...queryParamString.filter((param) => param.startsWith('page=') || param.startsWith('pageSize=') || param.startsWith('groupId='))];
        Object.entries(search).forEach(([key, value]) => {
            if (key === 'groupId') {
                setgroup(+value);
            } else if (key === 'createdAt') {
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
        console.log(`data: ${fetchData(queryParamString)}`)
    }, [dispatch]);

    const fetchData = async (queries: String[]) => {
        try {
            setLoading(true);
            setRows(undefined);
            await dispatch(getSmslogList(queries.filter((query) => !query.endsWith('=') && query !== 'status=0' && query !== 'policy=0').join('&')));
        } catch (error) {
            dispatch(
                openSnackbar({
                    open: true,
                    message: error ,
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
        setRes(smslogList!);
    }, [smslogList]);
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

    const handleSearchClick = () => {
        fetchData(queryParamString);
        setRes(smslogList!);
    };

    const handleResetClick = () => {
        setSearch(initSearchFields);
        setPage(0);
        setCreatedDate(null);
        setReset(!reset);
        fetchData(initialQueryParamString);
        setgroup(1);
    };

    const statusOptions = [
        {
            id: 0,
            value: 0,
            name: intl.formatMessage({ id: 'setting.blacklist.please-select-status' })
        },
        {
            id: 1,
            value: 1,
            name: intl.formatMessage({ id: 'smslog.unused' })
        },
        {
            id: 2,
            value: 2,
            name: intl.formatMessage({ id: 'smslog.used' })
        }
    ];

    const handleStatus = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const val = parseInt(event.target.value);
        setSearch({ ...search, status: val });
        setStatus(val);
    };

    // const handleExportClick = () => {
    //     ExportData(queryParamString);
    // };
    // const ExportData = (queries: String[]) => {
    //     axiosServices
    //         .get(
    //             `${envRef?.API_URL}/admin/smsLog/export?
    //                 ${queries
    //                     .filter((query) => {
    //                         return query === 'take_up_time=0' || query === 'error_code=1' || query.endsWith('=') ? false : true;
    //                     })
    //                     .join('&')
    //                 }`,
    //             { responseType: 'blob', headers: {} }
    //         )
    //         .then(function (response) {
    //             if (response) {
    //                 const outputFilename = `Log_${Date.now()}.xlsx`;

    //                 // If you want to download file automatically using link attribute.
    //                 const url = URL.createObjectURL(new Blob([response.data]));
    //                 const link = document.createElement('a');
    //                 link.href = url;
    //                 link.setAttribute('download', outputFilename);
    //                 document.body.appendChild(link);
    //                 link.click();

    //                 // OR you can save/write file locally.
    //                 // fs.writeFileSync(outputFilename, response.data);

    //                 dispatch(
    //                     openSnackbar({
    //                         open: true,
    //                         message: response?.data?.message,
    //                         variant: 'alert',
    //                         alert: {
    //                             color: 'success'
    //                         },
    //                         close: false
    //                     })
    //                 );
    //             } else {
    //                 dispatch(
    //                     openSnackbar({
    //                         open: true,
    //                         // message: response?.data?.message || defaultErrorMessage,
    //                         message: defaultErrorMessage,
    //                         variant: 'alert',
    //                         alert: {
    //                             color: 'error'
    //                         },
    //                         close: false,
    //                         anchorOrigin: {
    //                             vertical: 'top',
    //                             horizontal: 'center'
    //                         }
    //                     })
    //                 );
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
    // };

    const methodOptions: FlatOption[] = [
        {
            id: 0,
            valueString: ' ',
            name: intl.formatMessage({ id: 'smslog.please-select-event-template' })
        },
        {
            id: 1,
            valueString: 'login',
            name: intl.formatMessage({ id: 'general.login' })
        },
        {
            id: 2,
            valueString: 'register',
            name: intl.formatMessage({ id: 'register.sign-up' })
        },
        {
            id: 3,
            valueString: 'code',
            name: intl.formatMessage({ id: 'smslog.verification-code' })
        },
        {
            id: 4,
            valueString: 'bind',
            name: intl.formatMessage({ id: 'smslog.bind-phone-num' })
        },
        {
            id: 5,
            valueString: 'resetPwd',
            name: intl.formatMessage({ id: 'smslog.reset-pwd' })
        },
        {
            id: 6,
            valueString: 'cash',
            name: intl.formatMessage({ id: 'smslog.cash' })
        }
    ];

    const [isDeleteModalOpen, setIsDeleteModalOpen] = React.useState<boolean>(false);
    const [selectedIdsToDelete, setSelectedIdsToDelete] = React.useState<number[]>();
    const handleDeleteModal = (id: number[]) => {
        setSelectedIdsToDelete(id);
        setIsDeleteModalOpen(!isDeleteModalOpen);
    };

    function handleEvent(tag: string) {
        switch (tag) {
            case '':
                return <></>;
            case 'login':
                return <Chip label={intl.formatMessage({ id: 'general.login' })} size="small" chipcolor="success" />;
            case 'register':
                return <Chip label={intl.formatMessage({ id: 'register.sign-up' })} size="small" chipcolor="warning" />;
            case 'code':
                return <Chip label={intl.formatMessage({ id: 'smslog.verification-code' })} size="small" chipcolor="primary" />;
            case 'bind':
                return <Chip label={intl.formatMessage({ id: 'smslog.bind-phone-num' })} size="small" chipcolor="primary" />;
            case 'resetPwd':
                return <Chip label={intl.formatMessage({ id: 'smslog.reset-pwd' })} size="small" chipcolor="orange" />;
            case 'cash':
                return <Chip label={intl.formatMessage({ id: 'smslog.cash' })} size="small" chipcolor="info" />;
            default:
                return <></>;
        }
    }

    const handleDelete = (ids: number[]) => {
        axiosServices
            .post(`${envRef?.API_URL}admin/smsLog/delete`, { id: ids }, { headers: {} })
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
        const val = event?.target.value;

        if (searchParam === 'event') {
            setSearch({ ...search, [searchParam]: val! });
            setMethod(val!);
        } else {
            setSearch({ ...search, [searchParam]: val || '' });
        }
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
            `groupId=${group}`,
            ...queryParamString.filter((param) => !param.startsWith('page=') && !param.startsWith('pageSize=') && !param.startsWith('groupId='))
        ];
        fetchData(updatedParams);
    };

    const isSelected = (id: number) => selected.indexOf(id) !== -1;
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
                        // check if empty
                        ) : isDataListEmpty ? (
                            <TableRow>
                                <TableCell align="center" colSpan={headCells.length + 2}>
                                    <FolderOffTwoToneIcon sx={{ verticalAlign: 'bottom' }} />
                                    {intl.formatMessage({ id: 'general.no-records' })}
                                </TableCell>
                            </TableRow>
                        ) : (
                        // table sort
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
                                            align="center"
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
                                        <TableCell align="center">
                                            {handleEvent(row.event)}
                                        </TableCell>
                                        <TableCell align="center">
                                            {row.mobile}
                                        </TableCell>
                                        <TableCell align="center">
                                            {row.code}
                                        </TableCell>
                                        <TableCell align="center">
                                            {row.times}
                                        </TableCell>
                                        <TableCell align="center">
                                            {row.ip}
                                        </TableCell>
                                        <TableCell align="center">
                                            {row.status === 1 && (
                                                <Chip
                                                    label={intl.formatMessage({ id: 'smslog.unused' })}
                                                    size="small"
                                                    chipcolor="orange"
                                                />
                                            )}
                                            {row.status === 2 && (
                                                <Chip
                                                label={intl.formatMessage({ id: 'smslog.used' })}
                                                    size="small"
                                                    chipcolor="success"
                                                />
                                            )}
                                        </TableCell>
                                        <TableCell align="center">
                                            {row.createdAt || '-'}
                                        </TableCell>
                                        <TableCell align="center">
                                            {row.updatedAt || '-'}
                                        </TableCell>
                                        <TableCell align="center" className="sticky" sx={{ pr: 3, right: 0 }}>
                                           {/* <IconButton
                                                onClick={() => handleSmsModal(row.id)}
                                                color="secondary"
                                                size="medium"
                                                aria-label="Edit"
                                            >
                                                <EditIcon sx={{ fontSize: '1.3rem' }} />
                                            </IconButton> */}
                                            <IconButton
                                                onClick={() => handleDeleteModal([row.id])}
                                                color="secondary"
                                                size="medium"
                                                aria-label="Delete"
                                            >
                                                <DeleteIcon sx={{ fontSize: '1.3rem' }} />
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
            // Search Page Number input
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
        <MainCard title={<FormattedMessage id="servelog.server-log" />} content={false}>
            <CardContent>
                <Grid container alignItems="center" spacing={gridSpacing} justifyContent="flex-start">
                    <Grid item xs={12} sm={4} md={3}>
                        <TextField
                            select
                            fullWidth
                            onChange={(event) => {
                                handleSearch(event, 'event');
                            }}
                            value={method}
                            size="small"
                            label={intl.formatMessage({ id: 'setting.sms.eventTemplate' })}
                        >
                            {methodOptions.map((option) => {
                                return (
                                    <MenuItem key={option.id} value={option.valueString}>
                                        {option.name}
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
                            onChange={(event) => handleSearch(event, 'mobile')}
                            placeholder={intl.formatMessage({ id: 'smslog.phone-num' })}
                            value={search.mobile}
                            size="small"
                            // label={intl.formatMessage({ id: 'servelog.link-id' })}
                            label={intl.formatMessage({ id: 'smslog.phone-num' })}
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
                            onChange={(event) => handleSearch(event, 'ip')}
                            placeholder={intl.formatMessage({ id: 'smslog.sender-ip' })}
                            value={search.ip}
                            size="small"
                            label={intl.formatMessage({ id: 'smslog.sender-ip' })}
                        />
                    </Grid>
                    {expanded && (
                        <>
                            <Grid item xs={12} sm={3}>
                                <TextField
                                    select
                                    fullWidth
                                    onChange={(event) => {
                                        handleStatus(event);
                                        handleSearch(event, 'status');
                                    }}
                                    value={status}
                                    size="small"
                                    label={intl.formatMessage({ id: 'setting.blacklist.status' })}
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
                        {/* <ExportButton onClick={handleExportClick} /> */}
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
                        // list page num
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

export default SmsLog;