import React from 'react';
import { useIntl } from 'react-intl';
import { useDispatch, useSelector } from 'store';

// mui
import {
    Avatar,
    Badge,
    BadgeProps,
    Grid,
    IconButton,
    InputAdornment,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TablePagination,
    TableRow,
    TextField,
    Typography,
    styled
} from '@mui/material';
import { useTheme } from '@mui/material/styles';

import FolderOffTwoToneIcon from '@mui/icons-material/FolderOffTwoTone';
import SearchIcon from '@mui/icons-material/SearchTwoTone';
import FirstPageIcon from '@mui/icons-material/FirstPage';
import KeyboardArrowLeft from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRight from '@mui/icons-material/KeyboardArrowRight';
import LastPageIcon from '@mui/icons-material/LastPage';
import QuestionAnswerTwoToneIcon from '@mui/icons-material/QuestionAnswerTwoTone';

// ui-components
import SkeletonLoader from 'ui-component/cards/Skeleton/SkeletonLoader';

// API
import { getMessageList } from 'store/slices/user';
import { MessageListData } from 'types/user';
import { ResponseList } from 'types/response';
import { openSnackbar } from 'store/slices/snackbar';
import axiosServices from 'utils/axios';
import envRef from 'environment';
import { gridSpacing } from 'store/constant';
import Chip from 'ui-component/extended/Chip';

const PrivateMessage = () => {
    const intl = useIntl();
    const theme = useTheme();
    const dispatch = useDispatch();
    const defaultErrorMessage = intl.formatMessage({ id: 'auth-register.default-error' });

    const [loading, setLoading] = React.useState<boolean>(false);

    const [page, setPage] = React.useState<number>(0);
    const [totalCount, setTotalCount] = React.useState<number>(0);
    const [pageCount, setPageCount] = React.useState<number>(0);
    const [rowsPerPage, setRowsPerPage] = React.useState<number>(5);
    const [rows, setRows] = React.useState<MessageListData[]>();
    const [res, setRes] = React.useState<ResponseList<MessageListData>>();
    const { messageList } = useSelector((state) => state.user);

    const initialQueryParamString: string[] = [`page=1`, `pageSize=${rowsPerPage}`, `type=3`];
    const [queryParamString, setQueryParamString] = React.useState<String[]>(initialQueryParamString);

    React.useEffect(() => {
        setQueryParamString([`page=${page + 1}`, `pageSize=${rowsPerPage}`, `type=3`]);
    }, [page, rowsPerPage]);

    React.useEffect(() => {
        fetchData(queryParamString);
    }, [dispatch]);

    React.useEffect(() => {
        setRes(messageList!);
    }, [messageList]);

    React.useEffect(() => {
        setRows(res?.data?.list ? res.data.list : []);
    }, [res]);

    React.useEffect(() => {
        setPage(res?.data?.page ? res.data.page - 1 : 0);
        setRowsPerPage(res?.data?.pageSize ? res.data.pageSize : 5);
        setTotalCount(res?.data?.totalCount ? res.data.totalCount : 0);
        setPageCount(res?.data?.pageCount ? res.data.pageCount : 1);
    }, [rows]);

    const fetchData = async (queries: String[]) => {
        try {
            let proceedNext = false;
            setLoading(true);
            setRows(undefined);
            await axiosServices
                .get(`${envRef?.API_URL}admin/notice/pullMessages`, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0 && response?.data?.data.list) {
                        proceedNext = true;
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
            if (proceedNext) await dispatch(getMessageList(queries.join('&')));
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

    const handleChangePage = (event: React.MouseEvent<HTMLButtonElement, MouseEvent> | null, newPage: number) => {
        setPage(newPage);
        fetchData([`page=${newPage + 1}`, `pageSize=${rowsPerPage}`, `type=3`]);
    };

    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement> | undefined) => {
        const val = event ? parseInt(event.target.value) : 10;
        setRowsPerPage(val);
        setPage(page === 1 ? 0 : page);
        fetchData([`page=${page + 1}`, `pageSize=${val}`, `type=3`]);
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

                                            fetchData([
                                                `page=${inputValue > pageCount ? pageCount : inputValue}`,
                                                `pageSize=${rowsPerPage}`,
                                                `type=3`
                                            ]);
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
    function handleTag(tag: number) {
        switch (tag) {
            case 0:
                return <></>;
            case 1:
                return <Chip label={intl.formatMessage({ id: 'general.normal' })} size="medium" chipcolor="primary" />;
            case 2:
                return <Chip label={intl.formatMessage({ id: 'general.urgent' })} size="medium" chipcolor="error" />;
            case 3:
                return <Chip label={intl.formatMessage({ id: 'general.important' })} size="medium" chipcolor="warning" />;
            case 4:
                return <Chip label={intl.formatMessage({ id: 'general.remind' })} size="medium" chipcolor="success" />;
            case 5:
                return <Chip label={intl.formatMessage({ id: 'general.secondary' })} size="medium" chipcolor="" />;
            default:
                return <></>;
        }
    }

    async function handleRead(id: number) {
        try {
            setLoading(true);
            setRows(undefined);
            await axiosServices
                .post(`${envRef?.API_URL}admin/notice/upRead`, { id: id }, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        dispatch(
                            openSnackbar({
                                open: true,
                                message: response?.data?.message,
                                variant: 'alert',
                                alert: {
                                    color: 'success',
                                    severity: 'success'
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
    }

    const StyledBadge = styled(Badge)<BadgeProps>(({ theme }) => ({
        '& .MuiBadge-badge': {
            right: 3,
            top: 3,
            border: `2px solid ${theme.palette.background.paper}`,
            padding: '0 4px'
        }
    }));

    return (
        <>
            <TableContainer>
                <Table sx={{ minWidth: 750 }} aria-labelledby="tableTitle">
                    <TableBody>
                        {loading ? (
                            <TableRow>
                                <TableCell align="center">
                                    <SkeletonLoader />
                                </TableCell>
                            </TableRow>
                        ) : rows ? (
                            rows.map((row, index) => {
                                return (
                                    <TableRow hover tabIndex={-1} key={index}>
                                        <TableCell
                                            onClick={() => {
                                                if (!row.isRead) handleRead(row.id);
                                            }}
                                            sx={{ cursor: `${row.isRead ? 'inherit' : 'pointer'}` }}
                                        >
                                            <Grid container spacing={gridSpacing}>
                                                <Grid item xs={2}>
                                                    <StyledBadge badgeContent={row.isRead ? undefined : ''} color="error">
                                                        <Avatar sizes="medium">
                                                            <QuestionAnswerTwoToneIcon sx={{ fontSize: '1.3rem' }} />
                                                        </Avatar>
                                                    </StyledBadge>
                                                </Grid>
                                                <Grid item xs={8}>
                                                    <Typography variant="h4">{row.title}</Typography>
                                                    <Typography variant="body1" sx={{ paddingTop: '0.5rem', paddingBottom: '1rem' }}>
                                                        {row.createdAt}
                                                    </Typography>
                                                    <Typography
                                                        variant="body2"
                                                        dangerouslySetInnerHTML={{ __html: row.content }}
                                                    ></Typography>
                                                </Grid>
                                                <Grid item xs={2}>
                                                    {handleTag(row.tag)}
                                                </Grid>
                                            </Grid>
                                        </TableCell>
                                    </TableRow>
                                );
                            })
                        ) : (
                            <TableRow>
                                <TableCell align="center">
                                    <FolderOffTwoToneIcon sx={{ verticalAlign: 'bottom' }} />
                                    {intl.formatMessage({ id: 'general.no-records' })}
                                </TableCell>
                            </TableRow>
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
            <TablePagination
                rowsPerPageOptions={[5, 10, 20, 30, 50, 100]}
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
        </>
    );
};

export default PrivateMessage;
