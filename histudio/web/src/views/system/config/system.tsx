import React from 'react';

// material-ui
import { useTheme } from '@mui/material/styles';
import { Button, CardActions, CardContent, Divider, Grid, Tab, Tabs, Typography } from '@mui/material';

// project imports
import GeneralSetting from '../GeneralSetting';
import MailSetting from '../MailSetting';
import SmsSetting from '../SmsSetting';
import LoginRegister from '../LoginRegister';
import WithdrwalSetting from '../WithdrwalSetting';
import CloudStorage from '../CloudStorage';
import GeoLocation from '../GeoLocation';
import PaySetting from '../PaySetting';
import Wechat from '../Wechat';
import useConfig from 'hooks/useConfig';
import MainCard from 'ui-component/cards/MainCard';
import AnimateButton from 'ui-component/extended/AnimateButton';
import { gridSpacing } from 'store/constant';
import { useIntl } from 'react-intl';

// assets
import { TabsProps } from 'types';

// tabs
function TabPanel({ children, value, index, ...other }: TabsProps) {
    return (
        <div role="tabpanel" hidden={value !== index} id={`simple-tabpanel-${index}`} aria-labelledby={`simple-tab-${index}`} {...other}>
            {value === index && <div>{children}</div>}
        </div>
    );
}

function a11yProps(index: number) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`
    };
}

// ==============================|| PROFILE 1 - CHANGE PASSWORD ||============================== //

const System = () => {
    const intl = useIntl();
    const theme = useTheme();
    const { borderRadius } = useConfig();
    const [value, setValue] = React.useState<number>(0);

    const handleChange = (event: React.SyntheticEvent, newValue: number) => {
        setValue(newValue);
    };

    // tabs option
    const tabsOption = [
        {
            label: intl.formatMessage({ id: 'setting.basic' }),
            caption: intl.formatMessage({ id: 'setting.basicDesc' })
        },
        {
            label: intl.formatMessage({ id: 'setting.mail' }),
            caption: intl.formatMessage({ id: 'setting.mailDesc' })
        },
        {
            label: intl.formatMessage({ id: 'setting.sms' }),
            caption: intl.formatMessage({ id: 'setting.smsDesc' })
        },
        {
            label: intl.formatMessage({ id: 'setting.loginRegister' }),
            caption: intl.formatMessage({ id: 'setting.loginRegisterDesc' })
        },
        {
            label: intl.formatMessage({ id: 'setting.withdrawConfigure' }),
            caption: intl.formatMessage({ id: 'setting.withdrawConfigureDesc' })
        },
        {
            label: intl.formatMessage({ id: 'setting.cloudStorage' }),
            caption: intl.formatMessage({ id: 'setting.cloudStorageDesc' })
        },
        {
            label: intl.formatMessage({ id: 'setting.geographicalLocation' }),
            caption: intl.formatMessage({ id: 'setting.geographicalLocationDesc' })
        },
        {
            label: intl.formatMessage({ id: 'setting.paymentConfiguration' }),
            caption: intl.formatMessage({ id: 'setting.paymentConfigurationDesc' })
        },
        {
            label: intl.formatMessage({ id: 'setting.wechatConfiguration' }),
            caption: intl.formatMessage({ id: 'setting.wechatConfigurationDesc' })
        }
    ];

    return (
        <Grid container spacing={gridSpacing}>
            <Grid item xs={12}>
                <MainCard title="" content={false}>
                    <Grid container spacing={gridSpacing}>
                        <Grid item xs={12} lg={2}>
                            <CardContent>
                                <Tabs
                                    value={value}
                                    onChange={handleChange}
                                    orientation="vertical"
                                    variant="scrollable"
                                    sx={{
                                        '& .MuiTabs-flexContainer': {
                                            borderBottom: 'none'
                                        },
                                        '& button': {
                                            color: theme.palette.mode === 'dark' ? 'grey.600' : 'grey.900',
                                            minHeight: 'auto',
                                            minWidth: '100%',
                                            py: 1.5,
                                            px: 2,
                                            display: 'flex',
                                            flexDirection: 'row',
                                            alignItems: 'flex-start',
                                            textAlign: 'left',
                                            justifyContent: 'flex-start',
                                            borderRadius: `${borderRadius}px`
                                        },
                                        '& button.Mui-selected': {
                                            color: theme.palette.primary.main,
                                            background: theme.palette.mode === 'dark' ? theme.palette.dark.main : theme.palette.grey[50]
                                        },
                                        '& button > svg': {
                                            marginBottom: '0px !important',
                                            marginRight: 1.25,
                                            marginTop: 1.25,
                                            height: 20,
                                            width: 20
                                        },
                                        '& button > div > span': {
                                            display: 'block'
                                        },
                                        '& > div > span': {
                                            display: 'none'
                                        }
                                    }}
                                >
                                    {tabsOption.map((tab, index) => (
                                        <Tab
                                            key={index}
                                            label={
                                                <Grid container direction="column">
                                                    <Typography variant="subtitle1" color="inherit">
                                                        {tab.label}
                                                    </Typography>
                                                    <Typography component="div" variant="caption" sx={{ textTransform: 'capitalize' }}>
                                                        {tab.caption}
                                                    </Typography>
                                                </Grid>
                                            }
                                            {...a11yProps(index)}
                                        />
                                    ))}
                                </Tabs>
                            </CardContent>
                        </Grid>
                        <Grid item xs={12} lg={10}>
                            <CardContent
                                sx={{
                                    borderLeft: '1px solid',
                                    borderColor: theme.palette.mode === 'dark' ? theme.palette.background.default : theme.palette.grey[200],
                                    height: '100%'
                                }}
                            >
                                <TabPanel value={value} index={0}>
                                    <GeneralSetting />
                                </TabPanel>
                                <TabPanel value={value} index={1}>
                                    <MailSetting />
                                </TabPanel>
                                <TabPanel value={value} index={2}>
                                    <SmsSetting />
                                </TabPanel>
                                <TabPanel value={value} index={3}>
                                    <LoginRegister />
                                </TabPanel>
                                <TabPanel value={value} index={4}>
                                    <WithdrwalSetting />
                                </TabPanel>
                                <TabPanel value={value} index={5}>
                                    <CloudStorage />
                                </TabPanel>
                                <TabPanel value={value} index={6}>
                                    <GeoLocation />
                                </TabPanel>
                                <TabPanel value={value} index={7}>
                                    <PaySetting />
                                </TabPanel>
                                <TabPanel value={value} index={8}>
                                    <Wechat />
                                </TabPanel>
                            </CardContent>
                        </Grid>
                    </Grid>
                    <Divider />
                    <CardActions>
                        <Grid container justifyContent="space-between" spacing={0}>
                            <Grid item>
                                {value > 0 && (
                                    <AnimateButton>
                                        <Button variant="outlined" size="large" onClick={(e) => handleChange(e, value - 1)}>
                                            {intl.formatMessage({ id: 'general.back' })}
                                        </Button>
                                    </AnimateButton>
                                )}
                            </Grid>
                            {/*<Grid item>
                                {value < 3 && (
                                    <AnimateButton>
                                        <Button variant="contained" size="large" onClick={(e) => handleChange(e, 1 + value)}>
                                            Continue
                                        </Button>
                                    </AnimateButton>
                                )}
                            </Grid>*/}
                        </Grid>
                    </CardActions>
                </MainCard>
            </Grid>
        </Grid>
    );
};

export default System;
