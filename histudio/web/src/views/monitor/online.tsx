import * as React from 'react';

// locale
import { FormattedMessage, useIntl } from 'react-intl';

// material-ui
import { useTheme } from '@mui/material/styles';
import {
    Avatar,
    Box,
    CardContent,
    Grid,
    IconButton,
    InputAdornment,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TablePagination,
    TableRow,
    TableSortLabel,
    TextField,
    Tooltip,
    Typography
} from '@mui/material';
import { visuallyHidden } from '@mui/utils';

// project imports
import { UserOnlineListData } from 'types/user';
import { ResponseList } from 'types/response';
import { useDispatch, useSelector } from 'store';
import { getUserOnlineList } from 'store/slices/user';

// assets
import PersonOffTwoToneIcon from '@mui/icons-material/PersonOffTwoTone';
import SearchIcon from '@mui/icons-material/SearchTwoTone';
import FolderOffTwoToneIcon from '@mui/icons-material/FolderOffTwoTone';
import FirstPageIcon from '@mui/icons-material/FirstPage';
import KeyboardArrowLeft from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRight from '@mui/icons-material/KeyboardArrowRight';
import LastPageIcon from '@mui/icons-material/LastPage';
import logo from 'assets/images/general/logo.png';

import { GetComparator, EnhancedTableHeadProps, HeadCell, ArrangementOrder, KeyedObject } from 'types';

// ui-components
import Chip from 'ui-component/extended/Chip';
import MainCard from 'ui-component/cards/MainCard';
import ExpandButton from 'ui-component/searchbar/ExpandButton';
import ResetButton from 'ui-component/searchbar/ResetButton';
import SearchButton from 'ui-component/searchbar/SearchButton';
import GeneralDialog from 'ui-component/general/GeneralDialog';
import CustomDateTimeRangePicker from 'ui-component/general/CustomDateTimeRangePicker';

// API
import axiosServices from 'utils/axios';

// env
import envRef from 'environment';
import { isEqual } from 'lodash';

// project imports
import { openSnackbar } from 'store/slices/snackbar';
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';
import { format } from 'date-fns';
import { gridSpacing } from 'store/constant';

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

function stableSort(array: UserOnlineListData[], comparator: (a: UserOnlineListData, b: UserOnlineListData) => number) {
    const stabilizedThis = array.map((el: UserOnlineListData, index: number) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0] as UserOnlineListData, b[0] as UserOnlineListData);
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
        label: 'general.session-id',
        align: 'center'
    },
    {
        id: 'app',
        numeric: false,
        label: 'general.login-application',
        align: 'center'
    },
    {
        id: 'userId',
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
        id: 'avatar',
        numeric: false,
        label: 'user.avatar',
        align: 'center'
    },
    {
        id: 'ip',
        numeric: true,
        label: 'general.login-ip',
        align: 'center'
    },
    {
        id: 'browser',
        numeric: true,
        label: 'general.browser',
        align: 'center'
    },
    {
        id: 'os',
        numeric: true,
        label: 'general.os',
        align: 'center'
    },
    {
        id: 'heartbeatTime',
        numeric: true,
        label: 'general.last-active',
        align: 'center'
    },
    {
        id: 'firstTime',
        numeric: true,
        label: 'general.last-login',
        align: 'center'
    }
];

// ==============================|| TABLE HEADER ||============================== //

function EnhancedTableHead({ order, orderBy, onRequestSort }: EnhancedTableHeadProps) {
    const createSortHandler = (property: string) => (event: React.SyntheticEvent<Element, Event>) => {
        onRequestSort(event, property);
    };
    const intl = useIntl();

    return (
        <TableHead>
            <TableRow>
                {headCells.map((headCell) => (
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
                <TableCell sortDirection={false} align="center" className="sticky" sx={{ pr: 3, right: '0' }}>
                    {intl.formatMessage({ id: 'general.control' })}
                </TableCell>
            </TableRow>
        </TableHead>
    );
}

// ==============================|| DATA LIST ||============================== //

const OnlineUser = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);

    type SearchFields = {
        userId: string;
        username: string;
        ip: string;
        firstTime: Date[];
    };
    const initSearchFields: SearchFields = { userId: '', username: '', ip: '', firstTime: [] };
    const [order, setOrder] = React.useState<ArrangementOrder>('asc');
    const [orderBy, setOrderBy] = React.useState<string>('id');
    const [page, setPage] = React.useState<number>(0);
    const [totalCount, setTotalCount] = React.useState<number>(0);
    const [pageCount, setPageCount] = React.useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = React.useState<number>(10);
    const [search, setSearch] = React.useState<SearchFields>(initSearchFields);
    const [rows, setRows] = React.useState<UserOnlineListData[]>();
    const [res, setRes] = React.useState<ResponseList<UserOnlineListData>>();
    const { userOnlineList } = useSelector((state) => state.user);
    const [isDataListEmpty, setIsDataListEmpty] = React.useState<boolean>(true);

    const [firstTime, setFirstTime] = React.useState<Date[] | null>(null);
    const [reset, setReset] = React.useState<boolean>(false);

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
            if (key === 'firstTime') {
                if (firstTime) {
                    updatedParams.push(`${key}[]=${firstTime[0].getTime()}`);
                    updatedParams.push(`${key}[]=${firstTime[1].getTime()}`);
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
            await dispatch(getUserOnlineList(queries.filter((query) => !query.endsWith('=')).join('&')));
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
        setRes(userOnlineList!);
    }, [userOnlineList]);
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
        setPage(0);
        const updatedParams = [
            `page=1`,
            `pageSize=${rowsPerPage}`,
            ...queryParamString.filter((param) => !param.startsWith('page=') && !param.startsWith('pageSize='))
        ];
        setQueryParamString(updatedParams);
        fetchData(updatedParams);
        setRes(userOnlineList!);
    };

    const handleResetClick = () => {
        setPage(0);
        setSearch(initSearchFields);
        setFirstTime(null);
        setReset(!reset);
        fetchData(initialQueryParamString);
    };

    const [isKickModalOpen, setIsKickModalOpen] = React.useState<boolean>(false);
    const [selectedIdsToKick, setSelectedIdsToKick] = React.useState<string>();
    const handleKickModal = (id: string) => {
        setSelectedIdsToKick(id);
        setIsKickModalOpen(!isKickModalOpen);
    };

    const handleKick = (id: string) => {
        axiosServices
            .post(`${envRef?.API_URL}admin/monitor/userOffline`, { id: id }, { headers: {} })
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
        setIsKickModalOpen(!isKickModalOpen);
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

    const defaultErrorMessage = intl.formatMessage({ id: 'auth-register.default-error' });

    function handleHeartbeatTime(lastActive: number) {
        const currentDateTime = new Date().getTime();
        const lastActiveTimestamp = lastActive * 1000;

        if (lastActive > 0) {
            const msDiff = Math.abs(currentDateTime - lastActiveTimestamp);
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
            else return <TableCell align="center">{lastActive}</TableCell>;
        }
        return <TableCell align="center"></TableCell>;
    }

    function handleFirstTime(firstTime: number) {
        const formattedDateTime = format(new Date(firstTime * 1000), 'yyyy-MM-dd HH:mm:ss');

        if (firstTime > 0) {
            return <TableCell align="center">{formattedDateTime}</TableCell>;
        }
        return <TableCell align="center"></TableCell>;
    }

    function renderTable() {
        return (
            <TableContainer>
                <Table sx={{ minWidth: 750 }} aria-labelledby="tableTitle">
                    <EnhancedTableHead
                        numSelected={0}
                        order={order}
                        orderBy={orderBy}
                        onRequestSort={handleRequestSort}
                        rowCount={rows?.length || 0}
                        onSelectAllClick={() => {
                            return;
                        }}
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
                                const labelId = `enhanced - table - checkbox - ${index} `;

                                return (
                                    <TableRow hover role="checkbox" tabIndex={-1} key={index}>
                                        <TableCell align="center" component="th" id={labelId} scope="row" sx={{ cursor: 'pointer' }}>
                                            <Typography
                                                variant="subtitle1"
                                                sx={{ color: theme.palette.mode === 'dark' ? 'grey.600' : 'grey.900' }}
                                            >
                                                <Chip label={row.id || ''} size="small" chipcolor="info" />
                                            </Typography>
                                        </TableCell>
                                        <TableCell align="center">{row.app}</TableCell>
                                        <TableCell align="center">{row.userId}</TableCell>
                                        <TableCell align="center">{row.username}</TableCell>
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

                                        <TableCell align="center">{row.ip || '-'}</TableCell>
                                        <TableCell align="center">{row.browser || '-'}</TableCell>
                                        <TableCell align="center">{row.os || '-'}</TableCell>
                                        {handleHeartbeatTime(row.heartbeatTime || 0)}
                                        {handleFirstTime(row.firstTime || 0)}
                                        <TableCell align="center" className="sticky" sx={{ pr: 3, right: 0 }}>
                                            <IconButton
                                                onClick={() => handleKickModal(row.id)}
                                                color="secondary"
                                                size="medium"
                                                aria-label="Kick"
                                            >
                                                <Tooltip title={intl.formatMessage({ id: 'system.kick-user' })}>
                                                    <PersonOffTwoToneIcon sx={{ fontSize: '1.3rem' }} />
                                                </Tooltip>
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
        <MainCard title={<FormattedMessage id="system.online-user" />} content={false}>
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
                            onChange={(event) => handleSearch(event, 'userId')}
                            placeholder={intl.formatMessage({ id: 'user.search-user-id' })}
                            value={search.userId}
                            size="small"
                            label={intl.formatMessage({ id: 'user.user-id' })}
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
                            onChange={(event) => handleSearch(event, 'ip')}
                            placeholder={intl.formatMessage({ id: 'system.search-ip' })}
                            value={search.ip}
                            size="small"
                            label={intl.formatMessage({ id: 'general.login-ip' })}
                        />
                    </Grid>
                    {expanded && (
                        <>
                            <Grid item xs={12} sm={8} md={4}>
                                <CustomDateTimeRangePicker
                                    reset={reset}
                                    onSelectChange={(e) => {
                                        if (e) {
                                            setFirstTime(e);
                                            setSearch({ ...search, firstTime: e });
                                        } else {
                                            setFirstTime(null);
                                            setSearch({ ...search, firstTime: [] });
                                        }
                                    }}
                                    label={intl.formatMessage({ id: 'general.last-login' })}
                                />
                            </Grid>
                        </>
                    )}
                    <Grid item xs={12} sm={4} md={3} sx={{ display: 'flex', justifyContent: 'flex-end', marginLeft: 'auto' }}>
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
                confirmFunction={() => handleKick(selectedIdsToKick!)}
                isOpen={isKickModalOpen}
                setIsOpen={() => setIsKickModalOpen(!isKickModalOpen)}
                id="kick-confirm-modal"
                type="warning"
                content={intl.formatMessage({ id: 'general.confirm-kick-content' })}
            />
        </MainCard>
    );
};

export default OnlineUser;
