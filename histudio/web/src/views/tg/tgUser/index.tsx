import React,{ memo, useState } from 'react';
import { Button } from '@mui/material';
// import { useNavigate } from 'react-router-dom';
import MainCard from 'ui-component/cards/MainCard';
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import AccountBoxIcon from '@mui/icons-material/AccountBox';
import AlignVerticalBottomIcon from '@mui/icons-material/AlignVerticalBottom';
import ContactlessIcon from '@mui/icons-material/Contactless';
import HeadsetMicIcon from '@mui/icons-material/HeadsetMic';
// import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
// import PropTypes from 'prop-types';
import { FormattedMessage } from 'react-intl';
import styles from './index.module.scss';

import TabPanel01 from './children/TabPanel01';
import TabPanel02 from './children/TabPanel02';

const TgUser = () => {
    // const navigate = useNavigate();
    const [value, setValue] = useState(0);

    const handleChange = (event: React.SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };

    // const chatRoomToNavica = () => {
    //     navigate('/tg/chat/index?id=1');
    // };

    return (
        <MainCard title={<FormattedMessageTitle />} content={true}>
            {/* <Button variant="outlined" onClick={chatRoomToNavica}>
                聊天室跳转测试
            </Button> */}
            <Box sx={{ flexGrow: 1, bgcolor: 'background.paper', display: 'flex', height: '100%' }}>
                <Tabs
                    orientation="vertical"
                    variant="scrollable"
                    value={value}
                    onChange={handleChange}
                    aria-label="Vertical tabs example"
                    sx={{ borderRight: 1, borderColor: 'divider' }}
                >
                    <Tab label="Item One" icon={<AccountBalanceIcon />} {...a11yProps(0)} />
                    <Tab label="Item One" icon={<AccountBoxIcon />} {...a11yProps(1)} />
                    <Tab label="Item One" icon={<HeadsetMicIcon />} {...a11yProps(2)} />
                    <Tab label="Item One" icon={<ContactlessIcon />} {...a11yProps(2)} />
                    <Tab label="Item One" icon={<AlignVerticalBottomIcon />} {...a11yProps(2)} />
                </Tabs>
                <TabPanel01 value={value} index={0}>
                    Item One
                </TabPanel01>
                <TabPanel02 value={value} index={1}>
                    Item Two
                </TabPanel02>
                <TabPanel01 value={value} index={2}>
                    Item Three
                </TabPanel01>
            </Box>
        </MainCard>
    );
};
const FormattedMessageTitle = () => {
    return (
        <div className={styles.FormattedMessageTitle}>
            <FormattedMessage id="setting.cron.teleg-tg" />
            <div>
                <Button className={styles.btn} variant="outlined">
                    首页
                </Button>
                <Button className={styles.btn} variant="outlined">
                    注册
                </Button>
                <Button variant="outlined">登录</Button>
            </div>
        </div>
    );
};

// TabPanel01.propTypes = {
//     children: PropTypes.node,
//     index: PropTypes.number.isRequired,
//     value: PropTypes.number.isRequired
// };

function a11yProps(index: any) {
    return {
        id: `vertical-tab-${index}`,
        'aria-controls': `vertical-tabpanel-${index}`
    };
}

export default memo(TgUser);
