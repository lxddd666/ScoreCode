import React, { SyntheticEvent } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useIntl } from 'react-intl';

// mui
import { Box, Grid, Tab, Tabs, Typography } from '@mui/material';
import { useTheme } from '@mui/material/styles';
// mui icons
import NotificationImportantTwoToneIcon from '@mui/icons-material/NotificationImportantTwoTone';
import CampaignTwoToneIcon from '@mui/icons-material/CampaignTwoTone';
import QuestionAnswerTwoToneIcon from '@mui/icons-material/QuestionAnswerTwoTone';

// ui-components
import MainCard from 'ui-component/cards/MainCard';

// projects
import Notice from './notice';
import Announcement from './announcement';
import PrivateMessage from './private_message';

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

const Message = () => {
    const theme = useTheme();
    const intl = useIntl();
    const location = useLocation();
    const queryParams = new URLSearchParams(location.search);
    const typeParam = queryParams.get('type');

    const [messageType, setMessageType] = React.useState<number>(typeParam ? parseInt(typeParam) : 1);

    const tabsOption = [
        {
            label: intl.formatMessage({ id: 'inbox.notice' }),
            icon: <CampaignTwoToneIcon sx={{ fontSize: '1.3rem' }} />
        },
        {
            label: intl.formatMessage({ id: 'inbox.announcement' }),
            icon: <NotificationImportantTwoToneIcon sx={{ fontSize: '1.3rem' }} />
        },
        {
            label: intl.formatMessage({ id: 'inbox.private-message' }),
            icon: <QuestionAnswerTwoToneIcon sx={{ fontSize: '1.3rem' }} />
        }
    ];

    const [value, setValue] = React.useState<number>((messageType || 1) - 1);

    React.useEffect(() => {
        setValue((messageType || 1) - 1);
    }, [messageType]);

    const handleChange = (event: SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };

    return (
        <>
            <MainCard title={intl.formatMessage({ id: 'menu-items.my-inbox' })}>
                <Typography sx={{ paddingBottom: '1rem' }}>{intl.formatMessage({ id: 'inbox.header-content' })}</Typography>
                <Grid container spacing={2}>
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
                                <Tab
                                    onClick={() => setMessageType(index + 1)}
                                    key={index}
                                    component={Link}
                                    to="#"
                                    icon={tab.icon}
                                    label={tab.label}
                                    {...a11yProps(index)}
                                />
                            ))}
                        </Tabs>
                        <TabPanel value={value} index={0}>
                            <Notice />
                        </TabPanel>
                        <TabPanel value={value} index={1}>
                            <Announcement />
                        </TabPanel>
                        <TabPanel value={value} index={2}>
                            <PrivateMessage />
                        </TabPanel>
                    </Grid>
                </Grid>
            </MainCard>
        </>
    );
};

export default Message;
