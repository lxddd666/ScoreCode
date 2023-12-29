// material-ui
import { Pagination, CardActions, CardMedia, Divider, Table, TableBody, TableCell, TableContainer, TableHead, TableRow } from '@mui/material';

// third party
import PerfectScrollbar from 'react-perfect-scrollbar';

// project imports
import MainCard from 'ui-component/cards/MainCard';

// assets
import Flag1 from 'assets/images/widget/australia.jpg';
import Flag2 from 'assets/images/widget/brazil.jpg';
import Flag3 from 'assets/images/widget/germany.jpg';
import Flag4 from 'assets/images/widget/uk.jpg';
import Flag5 from 'assets/images/widget/usa.jpg';

// table data

const rows = [
    { id: 1, name: 'History', image: Flag1, subject: 'Germany', dept: 'Angelina', date: '2023-12-25', lastLoginTime: 1233, comment: '56.23%' },
    { id: 1, name: 'History', image: Flag2, subject: 'Germany', dept: 'Angelina', date: '2023-12-25', lastLoginTime: 1233, comment: '56.23%' },
    { id: 1, name: 'History', image: Flag3, subject: 'Germany', dept: 'Angelina', date: '2023-12-25', lastLoginTime: 1233, comment: '56.23%' },
    { id: 1, name: 'History', image: Flag4, subject: 'Germany', dept: 'Angelina', date: '2023-12-25', lastLoginTime: 1233, comment: '56.23%' },
    { id: 1, name: 'History', image: Flag5, subject: 'Germany', dept: 'Angelina', date: '2023-12-25', lastLoginTime: 1233, comment: '56.23%' },
    { id: 1, name: 'History', image: Flag1, subject: 'Germany', dept: 'Angelina', date: '2023-12-25', lastLoginTime: 1233, comment: '56.23%' },
    { id: 1, name: 'History', image: Flag1, subject: 'Germany', dept: 'Angelina', date: '2023-12-25', lastLoginTime: 1233, comment: '56.23%' },
    { id: 1, name: 'History', image: Flag1, subject: 'Germany', dept: 'Angelina', date: '2023-12-25', lastLoginTime: 1233, comment: '56.23%' },
    { id: 1, name: 'History', image: Flag1, subject: 'Germany', dept: 'Angelina', date: '2023-12-25', lastLoginTime: 1233, comment: '56.23%' },
];

// =========================|| DASHBOARD ANALYTICS - LATEST CUSTOMERS TABLE CARD ||========================= //
const columns = [
    {
        title: '公司ID',
        key: 'id'
    },
    {
        title: '公司名称',
        key: 'name'
    },
    {
        title: '国家',
        key: 'image'
    },
    {
        title: '注册时间',
        key: 'subject'
    },
    {
        title: '代理端口数',
        key: 'dept'
    },
    {
        title: '代理端口是否可用',
        key: 'date'
    },
    {
        title: '员工数',
        key: 'lastLoginTime'
    },
    {
        title: '账号数',
        key: 'comment'
    }
];



const LatestCustomerTableCard = (props: any) => {
    const { title } = props

    const renderTable = (value: any, key: any) => {
        let temp: any = ''
        if (key === 'image') {
            temp = <CardMedia component="img" image={value} title="image" sx={{ width: 30, height: 'auto' }} />
        } else {
            temp = value
        }
        return temp
    };
    return (
        <MainCard title={title} content={false}>
            <PerfectScrollbar style={{ height: 635, padding: 0 }}>
                <TableContainer>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell sx={{ pl: 3 }}>序号</TableCell>
                                {
                                    columns.map((item: any) => {
                                        return (
                                            <TableCell key={item.key} sx={{ pl: 3 }}>{item.title}</TableCell>
                                        )
                                    })
                                }
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {rows.map((row: any, index) => (
                                <TableRow hover key={index}>
                                    <TableCell sx={{ pl: 3 }}>
                                        {index + 1}
                                    </TableCell>
                                    {
                                        columns.map((item: any) => {
                                            return (
                                                <TableCell key={item.key} sx={{ pl: 3 }}>

                                                    {renderTable(row[item.key], item.key)}
                                                </TableCell>
                                            )
                                        })
                                    }
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </TableContainer>
            </PerfectScrollbar>

            <Divider />
            <CardActions sx={{ justifyContent: 'flex-end' }}>
                <Pagination count={10} color="primary" />
            </CardActions>
        </MainCard>
    )
}

export default LatestCustomerTableCard;
