import { useState, SyntheticEvent } from 'react';
import { Link } from 'react-router-dom';

// material-ui
import { useTheme } from '@mui/material/styles';
import { Box, Grid, Tab, Tabs } from '@mui/material';

// project imports
import Profile from './Profile';
import MyAccount from './MyAccount';
import Security from './Security';
import MainCard from 'ui-component/cards/MainCard';
import Withdrawal from './Withdrawal';
import { gridSpacing } from 'store/constant';
import { useIntl } from 'react-intl';

// assets
import AccountCircleTwoToneIcon from '@mui/icons-material/AccountCircleTwoTone';
import LibraryBooksTwoToneIcon from '@mui/icons-material/LibraryBooksTwoTone';
import LockTwoToneIcon from '@mui/icons-material/LockTwoTone';
import PaidIcon from '@mui/icons-material/Paid';

// types
import { TabsProps } from 'types';

// tabs panel
function TabPanel({ children, value, index, ...other }: TabsProps) {
    return (
        <div role="tabpanel" hidden={value !== index} id={`simple-tabpanel-${index}`} aria-labelledby={`simple-tab-${index}`} {...other}>
            {value === index && <Box sx={{ p: 0 }}>{children}</Box>}
        </div>
    );
}

function a11yProps(index: number) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`
    };
}

// ==============================|| PROFILE 1 ||============================== //

const Profile1 = () => {
    const intl = useIntl();

    // tabs option
    const tabsOption = [
        {
            label: intl.formatMessage({ id: 'account.profile' }),
            icon: <AccountCircleTwoToneIcon sx={{ fontSize: '1.3rem' }} />
        },
        {
            label: intl.formatMessage({ id: 'account.account' }),
            icon: <LibraryBooksTwoToneIcon sx={{ fontSize: '1.3rem' }} />
        },
        {
            label: intl.formatMessage({ id: 'account.security' }),
            icon: <LockTwoToneIcon sx={{ fontSize: '1.3rem' }} />
        },
        {
            label: intl.formatMessage({ id: 'profile.withdrawal' }),
            icon: <PaidIcon sx={{ fontSize: '1.3rem' }} />
        }
    ];

    const theme = useTheme();

    const [value, setValue] = useState<number>(0);
    const handleChange = (event: SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };

    return (
        <MainCard>
            <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                    <Tabs
                        value={value}
                        indicatorColor="primary"
                        textColor="primary"
                        onChange={handleChange}
                        aria-label="simple tabs example"
                        variant="scrollable"
                        sx={{
                            mb: 3,
                            '& a': {
                                minHeight: 'auto',
                                minWidth: 10,
                                py: 1.5,
                                px: 1,
                                mr: 2.25,
                                color: theme.palette.mode === 'dark' ? 'grey.600' : 'grey.900',
                                display: 'flex',
                                flexDirection: 'row',
                                alignItems: 'center',
                                justifyContent: 'center'
                            },
                            '& a.Mui-selected': {
                                color: theme.palette.primary.main
                            },
                            '& .MuiTabs-indicator': {
                                bottom: 2
                            },
                            '& a > svg': {
                                marginBottom: '0px !important',
                                mr: 1.25
                            }
                        }}
                    >
                        {tabsOption.map((tab, index) => (
                            <Tab key={index} component={Link} to="#" icon={tab.icon} label={tab.label} {...a11yProps(index)} />
                        ))}
                    </Tabs>
                    <TabPanel value={value} index={0}>
                        <Profile />
                    </TabPanel>
                    <TabPanel value={value} index={1}>
                        <MyAccount />
                    </TabPanel>
                    <TabPanel value={value} index={2}>
                        <Security />
                    </TabPanel>
                    <TabPanel value={value} index={3}>
                        <Withdrawal />
                    </TabPanel>
                </Grid>
            </Grid>
        </MainCard>
    );
};

export default Profile1;
