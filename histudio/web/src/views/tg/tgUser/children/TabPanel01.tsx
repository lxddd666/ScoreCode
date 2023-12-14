// import Box from '@mui/material/Box';
import { Button } from '@mui/material';
import { memo } from 'react';
import PropTypes from 'prop-types';
import Paper from '@mui/material/Paper';
import InputBase from '@mui/material/InputBase';
// import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';
// import MenuIcon from '@mui/icons-material/Menu';
import SearchIcon from '@mui/icons-material/Search';
import { DataGrid } from '@mui/x-data-grid';
// import DirectionsIcon from '@mui/icons-material/Directions';
import styles from '../index.module.scss';
import { useNavigate } from 'react-router-dom';

const rows = [
    { id: 1, lastName: 'Snow', firstName: 'Jon', age: 35 },
    { id: 2, lastName: 'Lannister', firstName: 'Cersei', age: 42 },
    { id: 3, lastName: 'Lannister', firstName: 'Jaime', age: 45 },
    { id: 4, lastName: 'Stark', firstName: 'Arya', age: 16 },
    { id: 5, lastName: 'Targaryen', firstName: 'Daenerys', age: null },
    { id: 6, lastName: 'Melisandre', firstName: null, age: 150 },
    { id: 7, lastName: 'Clifford', firstName: 'Ferrara', age: 44 },
    { id: 8, lastName: 'Frances', firstName: 'Rossini', age: 36 },
    { id: 9, lastName: 'Roxie', firstName: 'Harvey', age: 65 }
];
const TabPanel01 = (props: any) => {
    const { value, index, ...other } = props;
    const navigate = useNavigate();

    const columns = [
        { field: 'id', headerName: 'ID', flex: 1 },
        { field: 'firstName', headerName: 'First name', flex: 1 },
        { field: 'lastName', headerName: 'Last name', flex: 1 },
        {
            field: 'age',
            headerName: 'Age',
            type: 'number',
            flex: 1
        },
        // {
        //     field: 'fullName',
        //     headerName: 'Full name',
        //     description: 'This column has a value getter and is not sortable.',
        //     sortable: false,
        //     valueGetter: (params: any) => `${params.row.firstName || ''} ${params.row.lastName || ''}`
        // },
        {
            field: 'action',
            headerName: 'Action',
            sortable: false,
            flex: 1,
            renderCell: (params: any) => {
                // 这里的renderCell用于渲染自定义的按钮
                const chatRoomToNavica = () => {
                    navigate('/tg/chat/index?id=1');
                };

                return (
                    <>
                        <Button variant="outlined" onClick={chatRoomToNavica}>
                            聊天室
                        </Button>
                    </>
                );
            }
        }
    ];
    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`vertical-tabpanel-${index}`}
            aria-labelledby={`vertical-tab-${index}`}
            {...other}
            className={styles.tab1}
        >
            {value === index && (
                <div className={styles.tabpane01}>
                    <div className={styles.left}>
                        <Paper
                            component="form"
                            sx={{
                                p: '2px 4px',
                                display: 'flex',
                                alignItems: 'center',
                                width: '90%',
                                border: '1px solid rgb(127, 127, 127)'
                            }}
                        >
                            <InputBase sx={{ ml: 1, flex: 1 }} placeholder="请输入" inputProps={{ 'aria-label': 'search google maps' }} />
                            <IconButton type="button" sx={{ p: '10px' }} aria-label="search">
                                <SearchIcon />
                            </IconButton>
                        </Paper>
                        <div className={styles.leftList}>
                            <Button variant="outlined" className={styles.itemBtn}>
                                登录管理
                            </Button>
                            <Button variant="outlined" className={styles.itemBtn}>
                                社交帐号登录
                            </Button>
                        </div>
                    </div>
                    <div className={styles.right}>
                        <DataGrid rows={rows} columns={columns} checkboxSelection autoHeight />
                    </div>
                </div>
            )}
        </div>
    );
};
TabPanel01.propTypes = {
    children: PropTypes.node,
    index: PropTypes.number.isRequired,
    value: PropTypes.number.isRequired
};
export default memo(TabPanel01);
