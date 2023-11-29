import * as React from 'react';

// locale
import { FormattedMessage, useIntl } from 'react-intl';

// material-ui
import { useTheme } from '@mui/material/styles';
import {
    Link,
    Box,
    Button,
    CardContent,
    Checkbox,
    Dialog,
    FormControl,
    FormHelperText,
    Grid,
    IconButton,
    InputAdornment,
    InputLabel,
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
    Typography,
} from '@mui/material';
import { visuallyHidden } from '@mui/utils';

// project imports
import { NestedSubOption, FlatOption } from 'types/option';
import { PostListData } from 'types/cron';
import { ResponseList } from 'types/response';
import { useDispatch, useSelector } from 'store';
// cron list api
import { getPostList } from 'store/slices/cron';

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

// ui-components
import Chip from 'ui-component/extended/Chip';
import MainCard from 'ui-component/cards/MainCard';
import ExpandButton from 'ui-component/searchbar/ExpandButton';
import ResetButton from 'ui-component/searchbar/ResetButton';
import SearchButton from 'ui-component/searchbar/SearchButton';
import AddButton from 'ui-component/searchbar/AddButton';
import CronGroupButton from 'ui-component/searchbar/CronGroupButton';
import GeneralDialog from 'ui-component/general/GeneralDialog';

// API
import axiosServices from 'utils/axios';

import flattenNestedObjects from 'utils/flattenNestedObjects';
// env
import envRef from 'environment';
import { isEqual } from 'lodash';

// third party
import * as Yup from 'yup';
import { Formik, FormikValues } from 'formik';

// project imports
import AnimateButton from 'ui-component/extended/AnimateButton';
import useScriptRef from 'hooks/useScriptRef';
import { gridSpacing } from 'store/constant';
import { openSnackbar } from 'store/slices/snackbar';
import PrefixRadio from 'ui-component/general/PrefixRadio';
import ExpandableSelect from 'ui-component/general/ExpandableSelect';
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';
import { CronGroup } from './cronGroup';

const fieldStyle = {
    marginTop: `25%`
};

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

function stableSort(array: PostListData[], comparator: (a: PostListData, b: PostListData) => number) {
    const stabilizedThis = array.map((el: PostListData, index: number) => [el, index]);
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0] as PostListData, b[0] as PostListData);
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
        id: 'assignTask',
        numeric: false,
        label: 'setting.cron.task-grouping',
        align: 'center'
    },
    {
        id: 'taskName',
        numeric: false,
        label: 'setting.cron.mission-name',
        align: 'center'
    },
    {
        id: 'participant',
        numeric: false,
        label: 'setting.cron.execute-param',
        align: 'center'
    },
    {
        id: 'planning',
        numeric: false,
        label: 'setting.cron.execute-strategy',
        align: 'center'
    },
    {
        id: 'conversation',
        numeric: false,
        label: 'setting.cron.expression',
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
        numeric: true,
        label: 'general.created-date',
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
        ) : (
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

const Cron = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();
    const [loading, setLoading] = React.useState<boolean>(false);

    type SearchFields = {
        name: string;
        policy: number;
        status: number;
        groupId: number;
    };
    const initSearchFields: SearchFields = { name: '', policy: 0, status: 0, groupId: 0 };
    const [isMultiStrategy, setIsMultiStrategy] = React.useState<boolean>(false);

    const [order, setOrder] = React.useState<ArrangementOrder>('asc');
    const [orderBy, setOrderBy] = React.useState<string>('id');
    const [selected, setSelected] = React.useState<number[]>([]);
    const [page, setPage] = React.useState<number>(0);
    const [totalCount, setTotalCount] = React.useState<number>(0);
    const [pageCount, setPageCount] = React.useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = React.useState<number>(10);

    const [search, setSearch] = React.useState<SearchFields>(initSearchFields);

    const [rows, setRows] = React.useState<PostListData[]>();
    const [res, setRes] = React.useState<ResponseList<PostListData>>();
    const { postList } = useSelector((state) => state.cron);

    const [status, setStatus] = React.useState<number>(0);
    const [isDataListEmpty, setIsDataListEmpty] = React.useState<boolean>(true);
    const [group, setgroup] = React.useState<number>(0);
    const scriptedRef = useScriptRef();
    
    // const [addModalOptions, setAddModalOptions] = React.useState<PostListData[]>();

    // setAddModalOptions([])
    const FormikInitialValuesTemplate: FormikValues = {
        id: 0,
        groupId: 0,
        name: '',
        params: '',
        pattern: 0,
        policy: 0,
        count: 0,
        sort: 0,
        remark: '',
        status: 0,
        createdAt: '',
        updatedAt: '',
        groupName: ''
    };
    const [formikInitialValues, setFormikInitialValues] = React.useState<FormikValues>(FormikInitialValuesTemplate);
    const [createdDate, setCreatedDate] = React.useState<Date[] | null>(null);
    const [reset, setReset] = React.useState<boolean>(false);

    const initialQueryParamString: string[] = [`page=1`, `pageSize=${rowsPerPage}`, `groupId=0`];
    const [queryParamString, setQueryParamString] = React.useState<String[]>(initialQueryParamString);

    const [dataScopeList, setDataScopeList] = React.useState<ResponseList<PostListData>>();
    const [groupOptions, setgroupOptions] = React.useState<NestedSubOption[]>([]);
    async function getOptions() {
        await axiosServices
            .get(`${envRef?.API_URL}admin/cronGroup/select`, { headers: {} })
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
            await axiosServices
            .get(`${envRef?.API_URL}admin/cronGroup/select`, { headers: {} })
            .then(function (response) {
                if (response?.data?.code === 0) {
                    setgroupOptions(response.data.data.list);
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
            await dispatch(getPostList(queries.filter((query) => !query.endsWith('=') && query !== 'status=0' && query !== 'policy=0').join('&')));
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
        setRes(postList!);
    }, [postList]);
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
    const [isTaskModalOpen, setIsTaskModalOpen] = React.useState<boolean>(false);
    
    const handleTaskModal = () => {
        setIsTaskModalOpen(!isTaskModalOpen);
    };
    
    const handleAddModal = (id?: number) => {
        if (id) {
            axiosServices
                .get(`${envRef?.API_URL}admin/cron/view?id=${id}`, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        const selectedFormikValues: FormikValues = {
                            id: response.data.data.id,
                            groupId: response.data.data.groupId,
                            name: response.data.data.name,
                            params: response.data.data.params,
                            pattern: response.data.data.pattern,
                            policy: response.data.data.policy,
                            count: response.data.data.count,
                            remark: response.data.data.remark,
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
            setIsAddModalOpen(!isAddModalOpen);
            setFormikInitialValues(FormikInitialValuesTemplate);
        }
    };

    const handleSearchClick = () => {
        fetchData(queryParamString);
        setRes(postList!);
    };

    const handleResetClick = () => {
        setSearch(initSearchFields);
        setPage(0);
        setCreatedDate(null);
        setReset(!reset);
        setStatus(0);
        fetchData(initialQueryParamString);
        setgroup(1);
    };

    const statusOptions = [
        {
            id: 0,
            value: 0,
            name: intl.formatMessage({ id: 'setting.cron.please-select-strategy' })
        },
        {
            id: 1,
            value: 1,
            name: intl.formatMessage({ id: 'setting.cron.parallel-strategy' })
        },
        {
            id: 2,
            value: 2,
            name: intl.formatMessage({ id: 'setting.cron.single-column-strategy' })
        },
        {
            id: 3,
            value: 3,
            name: intl.formatMessage({ id: 'setting.cron.one-shot-strategy' })
        },
        {
            id: 4,
            value: 4,
            name: intl.formatMessage({ id: 'setting.cron.multiple-strategy' })
        }
    ];

    const statusOptions2 = [
        {
            id: 0,
            value: 0,
            name: intl.formatMessage({ id: 'user.please-select-status' })
        },
        {
            id: 1,
            value: 1,
            name: intl.formatMessage({ id: 'setting.cron.running' })
        },
        {
            id: 2,
            value: 2,
            name: intl.formatMessage({ id: 'setting.cron.over' })
        }
    ];

    const statusFlatOptions: FlatOption[] = [
        {
            id: 1,
            name: intl.formatMessage({ id: 'setting.cron.parallel-strategy' })
        },
        {
            id: 2,
            name: intl.formatMessage({ id: 'setting.cron.single-column-strategy' })
        },
        {
            id: 3,
            name: intl.formatMessage({ id: 'setting.cron.one-shot-strategy' })
        },
        {
            id: 4,
            name: intl.formatMessage({ id: 'setting.cron.multiple-strategy' })
        }
    ];

    const statusFlatOptions2: FlatOption[] = [
        {
            id: 1,
            name: intl.formatMessage({ id: 'setting.cron.running' })
        },
        {
            id: 2,
            name: intl.formatMessage({ id: 'setting.cron.over' })
        }
    ];

    const handleStatus = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const val = parseInt(event.target.value);
        setSearch({ ...search, status: val });
        setStatus(val);
    };

    const handlegroup = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const val = parseInt(event.target.value);
        setgroup(val);
    };

    const handleActivate = (id: number, status: number) => {
        axiosServices
            .post(
                `${envRef?.API_URL}admin/cron/status`,
                {
                    id: id,
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
    const [selectedIdsToDelete, setSelectedIdsToDelete] = React.useState<number[]>();
    const handleDeleteModal = (id: number[]) => {
        setSelectedIdsToDelete(id);
        setIsDeleteModalOpen(!isDeleteModalOpen);
    };

    const handleDelete = (ids: number[]) => {
        axiosServices
            .post(`${envRef?.API_URL}admin/cron/delete`, { id: ids }, { headers: {} })
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
                                        <TableCell align="center">{row.groupName}</TableCell>
                                        <TableCell align="center">{row.name}</TableCell>
                                        <TableCell align="center">{row.params}</TableCell>

                                        <TableCell align="center">
                                            {row.policy === 1 && (
                                                <Chip
                                                    label={intl.formatMessage({ id: 'setting.cron.parallel-strategy' })}
                                                    size="small"
                                                    chipcolor="primary"
                                                />
                                            )}
                                            {row.policy === 2 && (
                                                <Chip
                                                    label={intl.formatMessage({ id: 'setting.cron.single-column-strategy' })}
                                                    size="small"
                                                    chipcolor="primary"
                                                />
                                            )}
                                             {row.policy === 3 && (
                                                <Chip
                                                    label={intl.formatMessage({ id: 'setting.cron.one-shot-strategy' })}
                                                    size="small"
                                                    chipcolor="primary"
                                                />
                                            )}
                                             {row.policy === 4 && (
                                                <Chip
                                                    label={intl.formatMessage({ id: 'setting.cron.multiple-strategy' })}
                                                    size="small"
                                                    chipcolor="primary"
                                                />
                                            )}
                                        </TableCell>
                                        <TableCell align="center">{row.pattern}</TableCell>
                                        <TableCell align="center">
                                            {row.status === 1 && (
                                                <Chip
                                                    // label={`运行中`}
                                                    label={intl.formatMessage({ id: 'setting.cron.running' })}
                                                    size="small"
                                                    chipcolor="success"
                                                />
                                            )}
                                            {row.status === 2 && (
                                                <Chip
                                                    label={intl.formatMessage({ id: 'setting.cron.over' })}
                                                    size="small"
                                                    chipcolor="orange"
                                                />
                                            )}
                                        </TableCell>
                                        <TableCell align="center">{row.createdAt || '-'}</TableCell>
                                        <TableCell align="center" className="sticky" sx={{ pr: 3, right: 0 }}>
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
                                            <IconButton
                                                onClick={() => handleAddModal(row.id)}
                                                color="secondary"
                                                size="medium"
                                                aria-label="Edit"
                                            >
                                                <EditIcon sx={{ fontSize: '1.3rem' }} />
                                            </IconButton>
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
        <MainCard title={<FormattedMessage id="setting.cron.scheduled-tasks" />} content={false}>
            <CardContent>
                <Grid container alignItems="center" spacing={gridSpacing} justifyContent="flex-start">
                    <Grid item xs={12} sm={4} md={3}>
                        <TextField
                            select
                            fullWidth
                            onChange={handlegroup}
                            value={group}
                            size="small"
                            label={intl.formatMessage({ id: 'setting.cron.task-grouping' })}
                        >
                            <MenuItem key={0} value={0}>
                                {intl.formatMessage({ id: 'user.all' })}
                            </MenuItem>
                            {flattenNestedObjects(groupOptions).map((option) => {
                                return (
                                    <MenuItem key={option.id} value={option.id}>
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
                            onChange={(event) => handleSearch(event, 'name')}
                            placeholder={intl.formatMessage({ id: 'position.search-position-name' })}
                            value={search.name}
                            size="small"
                            label={intl.formatMessage({ id: 'setting.cron.mission-name' })}
                        />
                    </Grid>
                    {expanded && (
                        <>
                            <Grid item xs={12} sm={3}>
                                <TextField
                                    select
                                    fullWidth
                                    onChange={(event) => {
                                        handleSearch(event, 'policy');
                                    }}
                                    value={search.policy}
                                    size="small"
                                    label={intl.formatMessage({ id: 'setting.cron.execute-strategy' })}
                                >
                                    {statusOptions.map((option) => {
                                        return (
                                            <MenuItem key={option.id} value={option.id}>
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
                                    onChange={(event) => {
                                        handleStatus(event);
                                        handleSearch(event, 'status');
                                    }}
                                    value={status}
                                    size="small"
                                    label={intl.formatMessage({ id: 'general.status' })}
                                >
                                    {statusOptions2.map((option) => {
                                        return (
                                            <MenuItem key={option.id} value={option.value}>
                                                {option.name}
                                            </MenuItem>
                                        );
                                    })}
                                </TextField>
                            </Grid>
                        </>
                    )}
                    <Grid item xs={12} sm={4} md={3} sx={{ display: 'flex', justifyContent: 'flex-end', marginLeft: 'auto' }}>
                        <AddButton onClick={() => handleAddModal()} tooltipTitle={intl.formatMessage({ id: 'setting.cron.add-task' })} />
                        <CronGroupButton onClick={() => handleTaskModal()} tooltipTitle={intl.formatMessage({ id: 'setting.cron.task-grouping' })} />
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
            
            {/* modal */}
            <Dialog
                id="addModal"
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
                                ? intl.formatMessage({ id: 'setting.cron.add-task' })
                                : intl.formatMessage({ id: 'position.edit-position' })}
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
                                    .required(intl.formatMessage({ id: 'setting.cron.task-name-cannot-empty' })),
                            })}
                            onSubmit={async (values, { setErrors, setStatus, setSubmitting }) => {
                                try {
                                    await axiosServices
                                        .post(
                                            `${envRef?.API_URL}admin/cron/edit`,
                                            {
                                                id: values.id || 0,
                                                groupId: values.groupId || 0,
                                                name: values.name,
                                                params: values.params,
                                                pattern: values.pattern,
                                                policy: values.policy,
                                                count: values.count || 0,
                                                sort: values.sort || 0,
                                                remark: values.remark || '',
                                                status: values.status,
                                                createdAt: values.createdAt || '',
                                                updatedAt: ''
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
                                        <Grid item className="expandableSelectContainer" xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <ExpandableSelect
                                                    label={intl.formatMessage({ id: 'setting.cron.task-grouping' })}
                                                    id="groupId"
                                                    options={dataScopeList}
                                                    onSelectChange={(e) => setFieldValue('groupId', e[0])}
                                                    initialValue={[values.groupId]}
                                                />
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                error={Boolean(touched.name && errors.name)}
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-code-add">
                                                    {intl.formatMessage({ id: 'setting.cron.mission-name' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-code-add"
                                                    type="text"
                                                    value={values.name}
                                                    name="name"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{}}
                                                />
                                                {touched.name && errors.name && (
                                                    <FormHelperText error>
                                                        {errors.name.toString()}
                                                    </FormHelperText>
                                                )}
                                            </FormControl>
                                            <FormattedMessage id="setting.cron.go-function-name" />
                                        </Grid>

                                        <Grid item xs={12} sm={6}>
                                            <FormControl
                                                fullWidth
                                                sx={{ ...theme.typography.customInput }}
                                                className="MultiLineInput"
                                            >
                                                <InputLabel htmlFor="outlined-adornment-remark-add">
                                                    {intl.formatMessage({ id: 'setting.cron.execute-param' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    multiline
                                                    id="outlined-adornment-remark-add"
                                                    type="remark"
                                                    value={values.params}
                                                    name="params"
                                                    rows={1}
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{
                                                        style: { marginTop: '4%', whiteSpace: 'break-spaces' }
                                                    }}
                                                />
                                            </FormControl>
                                        </Grid>
                                        <Grid item xs={12} sm={6}>
                                            <PrefixRadio
                                                options={statusFlatOptions}
                                                label={intl.formatMessage({ id: 'setting.cron.execute-strategy' })}
                                                id="policy"
                                                onSelectChange={(e) => {
                                                    e === 4 ? setIsMultiStrategy(true) : setIsMultiStrategy(false)
                                                    setFieldValue('policy', e)
                                                }}
                                                initialValue={values.policy}
                                            />
                                        </Grid>
                                        {/* if 多次策略 else hide */}
                                        { isMultiStrategy && 
                                            <Grid item xs={12} sm={6}>
                                                <FormControl
                                                    fullWidth
                                                    sx={{ ...theme.typography.customInput }}
                                                >
                                                    <InputLabel htmlFor="outlined-adornment-post-name-add">
                                                        {intl.formatMessage({ id: 'setting.cron.no-of-executions' })}
                                                    </InputLabel>
                                                    <OutlinedInput
                                                        id="outlined-adornment-post-name-add"
                                                        type="number"
                                                        value={values.count}
                                                        name="count"
                                                        onBlur={handleBlur}
                                                        onChange={handleChange}
                                                        inputProps={{}}
                                                    />
                                                </FormControl>
                                            </Grid>
                                        }

                                        <Grid item xs={12} sm={6} >
                                            <FormControl
                                                style={fieldStyle}
                                                fullWidth
                                                sx={{ ...theme.typography.customInput }}
                                            >
                                                <InputLabel htmlFor="outlined-adornment-code-add">
                                                    {intl.formatMessage({ id: 'setting.cron.timed-expression' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    id="outlined-adornment-post-name-add"
                                                    type="text"
                                                    value={values.pattern}
                                                    name="pattern"
                                                    onBlur={handleBlur}
                                                    onChange={handleChange}
                                                    inputProps={{
                                                       
                                                    }}
                                                />
                                            </FormControl>
                                            <FormattedMessage id="setting.cron.syntax-reference"/>
                                            <Link href="https://goframe.org/pages/viewpage.action?pageId=30736411" underline="none">
                                                {'https://goframe.org/pages/viewpage.action?pageId=30736411'}
                                            </Link>
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
                                                options={statusFlatOptions2}
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
                                                    {intl.formatMessage({ id: 'general.remark' })}
                                                </InputLabel>
                                                <OutlinedInput
                                                    multiline
                                                    id="outlined-adornment-remark-add"
                                                    type="remark"
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
            <Dialog
                id="cronGroupModal"
                className="hideBackdrop"
                fullWidth
                maxWidth="xl"
                onClose={() => handleTaskModal()}
                open={isTaskModalOpen}
                sx={{ '& .MuiDialog-paper': { p: '0rem 0rem' }, backgroundColor: 'none' }}
            >  
            {CronGroup()}
            </Dialog>
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

export default Cron;