import React from 'react';
import { useIntl } from 'react-intl';

import Chart, { Props } from 'react-apexcharts';

import {
    Card,
    CardContent,
    Divider,
    Grid,
    Icon,
    InputLabel,
    LinearProgress,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Typography,
    useTheme
} from '@mui/material';
import MemoryTwoToneIcon from '@mui/icons-material/MemoryTwoTone';
import AppsTwoToneIcon from '@mui/icons-material/AppsTwoTone';
import PieChartTwoToneIcon from '@mui/icons-material/PieChartTwoTone';
import AnalyticsTwoToneIcon from '@mui/icons-material/AnalyticsTwoTone';
import useWebSocket from 'hooks/useWebSocket';
import MainCard from 'ui-component/cards/MainCard';
import { EventOptions } from 'contexts/WebSocketContext';
import TotalLineChartCard from 'ui-component/cards/TotalLineChartCard';

const icons = {
    HardwareChip: MemoryTwoToneIcon,
    AppsSharp: AppsTwoToneIcon,
    PieChart: PieChartTwoToneIcon,
    Analytics: AnalyticsTwoToneIcon
};

const SystemMonitoring = () => {
    const theme = useTheme();
    const { monitorRunInfo, monitorTrends, startHeartBeat, lastMessage } = useWebSocket();
    const [trends, setTrends] = React.useState<any>(null);
    // const [trendsNet, setTrendsNet] = React.useState<any>(null);
    const [runInfo, setRunInfo] = React.useState<any>(null);
    const intl = useIntl();

    const [trendsDataHeadChartData, setTrendsDataHeadChartData] = React.useState<number[]>([0, 0, 0, 0, 0, 0, 0, 0, 0, 0]);
    const [trendsDataNetChartUp, setTrendsDataNetChartUp] = React.useState<number[]>([0, 0, 0, 0, 0, 0, 0, 0, 0, 0]);
    const [trendsDataNetChartDown, setTrendsDataNetChartDown] = React.useState<number[]>([0, 0, 0, 0, 0, 0, 0, 0, 0, 0]);
    const [trendsDataNetChartDate, setTrendsDataNetChartDate] = React.useState<string[]>([
        '0',
        '0',
        '0',
        '0',
        '0',
        '0',
        '0',
        '0',
        '0',
        '0'
    ]);

    // Chart Data for trends.data.head
    const dataHeadChartData: Props = {
        type: 'area',
        height: 150,
        options: {
            theme: theme.palette,
            chart: {
                id: 'data-net-chart',
                toolbar: {
                    show: false
                },
                zoom: {
                    enabled: false
                },
                sparkline: {
                    enabled: false
                }
            },
            dataLabels: {
                enabled: false
            },
            fill: {
                type: 'solid',
                colors: [theme.palette.primary.light]
            },
            stroke: {
                curve: 'smooth',
                width: 3
            },
            labels: trendsDataNetChartDate,
            xaxis: {
                labels: {
                    show: false
                },
                tooltip: {
                    enabled: true
                }
            },
            yaxis: {
                min: 0,
                max: Math.max(...trendsDataHeadChartData) === 0 ? 2.5 : Math.max(...trendsDataHeadChartData) + 0.5,
                labels: {
                    show: false
                },
                tooltip: {
                    enabled: true
                }
            },
            tooltip: {
                theme: theme.palette.mode,
                fixed: {
                    enabled: false
                },
                shared: false,
                x: {
                    show: false
                },
                // y: {
                //     formatter(val, opts) {
                //         return `${intl.formatMessage({ id: 'system.cpu-load-ratio' })} ${val}KB`;
                //     }
                // },
                custom: function ({ series, seriesIndex, dataPointIndex }) {
                    return ``;
                },
                marker: {
                    show: true
                }
            }
        },
        series: [
            {
                name: '',
                data: trendsDataHeadChartData
            }
        ]
    };

    // Chart data for trends.data.net
    const dataNetChartData: Props = {
        height: '300px',
        type: 'area',
        options: {
            theme: theme.palette,
            chart: {
                id: 'data-net-chart',
                toolbar: {
                    show: false
                },
                zoom: {
                    enabled: false
                },
                sparkline: {
                    enabled: false
                }
            },
            colors: [theme.palette.primary.main, theme.palette.secondary.main],
            dataLabels: {
                enabled: true
            },
            stroke: {
                curve: 'smooth',
                width: 2
            },
            fill: {
                type: 'gradient',
                gradient: {
                    shadeIntensity: 1,
                    opacityFrom: 0.8,
                    opacityTo: 0.2,
                    stops: [0, 80, 100]
                }
            },
            legend: {
                show: true
            },
            labels: trendsDataNetChartDate,
            xaxis: {
                type: 'category'
            },
            yaxis: {
                title: {
                    text: 'Y-Axis Label'
                },
                min: 0,
                max: Math.max(...trendsDataNetChartUp) + 10,
                labels: {
                    show: true,
                    offsetX: -20
                },
                tooltip: {
                    enabled: true
                }
            },
            tooltip: {
                theme: theme.palette.mode,
                fixed: {
                    enabled: false
                },
                x: {
                    show: true
                },
                y: {
                    formatter(val, opts) {
                        return val + 'KB';
                    }
                },
                marker: {
                    show: false
                }
            }
        },
        series: [
            {
                name: intl.formatMessage({ id: 'system.traffic-upload' }),
                data: trendsDataNetChartUp
            },
            {
                name: intl.formatMessage({ id: 'system.traffic-download' }),
                data: trendsDataNetChartDown
            }
        ]
    };

    async function initMonitorSystem() {
        // init display
        await monitorRunInfo();
        await monitorTrends();

        // init start heartbeat event
        await startHeartBeat(2, [EventOptions.MonitorRunInfo, EventOptions.MonitorTrends]);
    }

    React.useEffect(() => {
        initMonitorSystem();
    }, []);

    function handleLastMessage() {
        if (lastMessage) {
            const data = JSON.parse(lastMessage.data);
            switch (data.event) {
                case 'admin/monitor/runInfo':
                    setRunInfo(data);
                    return;
                case 'admin/monitor/trends':
                    setTrends(data);
                    return;
            }
        }
    }

    React.useEffect(() => {
        handleLastMessage();
    }, [lastMessage]);

    React.useEffect(() => {}, [runInfo]);

    React.useEffect(() => {
        if (trends) {
            const analytic = trends.data.head.find((item: any) => item.iconClass === 'Analytics');
            setTrendsDataHeadChartData([...trendsDataHeadChartData, parseFloat(analytic.data)]);
            setTrendsDataNetChartUp(trends.data.net.map((net: any) => net.up));
            setTrendsDataNetChartDown(trends.data.net.map((net: any) => net.down));
            setTrendsDataNetChartDate(trends.data.net.map((net: any) => net.time));
        }
    }, [trends]);

    React.useEffect(() => {
        if (trendsDataHeadChartData.length > 10) {
            const updated = trendsDataHeadChartData.slice(1);
            setTrendsDataHeadChartData(updated);
        }
    }, [trendsDataHeadChartData]);

    const handleDataHead = () => {
        if (trends !== null) {
            return (
                <Grid container alignItems="stretch" spacing={2}>
                    {trends.data.head.map((head: any, index: number) => (
                        <Grid key={index} item xs={12} sm={6} md={3}>
                            <Card
                                variant="outlined"
                                square
                                style={{ padding: '1rem', height: '100%', display: 'flex', flexDirection: 'column' }}
                            >
                                <Grid
                                    container
                                    sx={{
                                        alignItems: 'center',
                                        justifyContent: 'space-between'
                                    }}
                                    spacing={2}
                                >
                                    <Typography variant="h5" paddingLeft="1rem">
                                        {head.title}
                                    </Typography>
                                    <Icon component={icons[`${head.iconClass}` as keyof typeof icons]} sx={{ alignSelf: 'flex-end' }} />
                                </Grid>
                                <Divider />
                                {head.iconClass === 'HardwareChip' && (
                                    <Grid paddingY="1rem" sx={{ flex: '1' }}>
                                        <Typography marginBottom="1rem">{head.data}</Typography>
                                        <Typography>{head.extra.data}</Typography>
                                        <Typography>{head.extra.data1}</Typography>
                                    </Grid>
                                )}
                                {head.iconClass === 'AppsSharp' && (
                                    <Grid paddingY="1rem" sx={{ flex: '1' }}>
                                        <Typography marginBottom="1rem">{head.data}</Typography>
                                        <Typography>
                                            {intl.formatMessage({ id: 'system.used-memory' })}
                                            {head.extra.data}
                                        </Typography>
                                        <Typography>
                                            {intl.formatMessage({ id: 'system.remaining-memory' })}
                                            {head.extra.data1}
                                        </Typography>
                                    </Grid>
                                )}
                                {head.iconClass === 'PieChart' && (
                                    <Grid paddingY="1rem" sx={{ flex: '1' }}>
                                        <Typography marginBottom="1rem">{head.data}</Typography>
                                        <InputLabel htmlFor="progress-bar">{head.extra.data}%</InputLabel>
                                        <LinearProgress id="progress-bar" variant="determinate" value={parseInt(head.extra.data)} />
                                    </Grid>
                                )}
                                {head.iconClass === 'Analytics' && (
                                    <Grid paddingY="1rem" sx={{ flex: '1' }}>
                                        <TotalLineChartCard chartData={dataHeadChartData} value={head.data} title="" percentage="" />
                                    </Grid>
                                )}
                                <Divider />
                                <Grid
                                    container
                                    sx={{
                                        alignItems: 'center',
                                        justifyContent: 'space-between',
                                        width: '100%',
                                        margin: '0'
                                    }}
                                    spacing={2}
                                >
                                    <Typography>{head.bottomTitle}</Typography>
                                    <Typography>{head.totalSum}</Typography>
                                </Grid>
                            </Card>
                        </Grid>
                    ))}
                </Grid>
            );
        }
    };

    const handleDataNet = () => {
        return (
            <Grid paddingTop="1rem">
                <Card variant="outlined" square style={{ padding: '1rem', height: '100%', display: 'flex', flexDirection: 'column' }}>
                    <Typography variant="h4">{intl.formatMessage({ id: 'system.traffic' })}</Typography>
                    <Divider />
                    <Grid
                        container
                        sx={{
                            alignItems: 'center',
                            justifyContent: 'space-between',
                            paddingY: '1rem'
                        }}
                        spacing={2}
                    >
                        <Grid item>
                            <Typography>{trends?.data.net.length > 0 ? trends.data.net[trends.data.net.length - 1].up : 0} KB</Typography>
                            <Typography>{intl.formatMessage({ id: 'system.traffic-upload' })}</Typography>
                        </Grid>
                        <Grid item>
                            <Typography>{trends?.data.net.length > 0 ? trends.data.net[trends.data.net.length - 1].down : 0} KB</Typography>
                            <Typography>{intl.formatMessage({ id: 'system.traffic-download' })}</Typography>
                        </Grid>
                        <Grid item>
                            <Typography>
                                {trends?.data.net.length > 0 ? trends.data.net[trends.data.net.length - 1].bytesSent : 0}
                            </Typography>
                            <Typography>{intl.formatMessage({ id: 'system.bytes-sent' })}</Typography>
                        </Grid>
                        <Grid item>
                            <Typography>
                                {trends?.data.net.length > 0 ? trends.data.net[trends.data.net.length - 1].bytesRecv : 0}
                            </Typography>
                            <Typography>{intl.formatMessage({ id: 'system.bytes-received' })}</Typography>
                        </Grid>
                    </Grid>
                    <Grid>
                        <Chart {...dataNetChartData} />
                    </Grid>
                </Card>
            </Grid>
        );
    };

    const handleRunInfo = () => {
        function formatDuration(runTime: number) {
            const days = Math.floor(runTime / 86400);
            runTime %= 86400;
            const hours = Math.floor(runTime / 3600);
            runTime %= 3600;
            const minutes = Math.floor(runTime / 60);
            const seconds = runTime % 60;

            return `
            ${days ? `${days} ${intl.formatMessage({ id: 'general.days' })}` : ''}
            ${hours ? `${hours} ${intl.formatMessage({ id: 'general.hours' })}` : ''}
            ${minutes ? `${minutes} ${intl.formatMessage({ id: 'general.minutes' })}` : ''}
            ${seconds ? `${seconds} ${intl.formatMessage({ id: 'general.seconds' })}` : ''}
            `;
        }

        return (
            <>
                <Grid paddingTop="1rem">
                    <Card variant="outlined" square style={{ padding: '1rem', height: '100%', display: 'flex', flexDirection: 'column' }}>
                        <Typography variant="h4" sx={{ paddingBottom: '1rem' }}>
                            {intl.formatMessage({ id: 'system.server-info' })}
                        </Typography>
                        <Divider />
                        <TableContainer>
                            <Table className="borderTable">
                                <TableHead>
                                    <TableRow>
                                        <TableCell width="25%">{intl.formatMessage({ id: 'system.server-name' })}</TableCell>
                                        <TableCell width="25%">{intl.formatMessage({ id: 'system.server-os' })}</TableCell>
                                        <TableCell width="50%">{intl.formatMessage({ id: 'system.server-ip' })}</TableCell>
                                    </TableRow>
                                </TableHead>
                                <TableBody>
                                    <TableRow>
                                        <TableCell>{runInfo ? runInfo.data.hostname : ''}</TableCell>
                                        <TableCell>{runInfo ? runInfo.data.os : ''}</TableCell>
                                        <TableCell>
                                            {runInfo ? runInfo.data.public_ip : ''} / {runInfo ? runInfo.data.intranet_ip : ''}
                                        </TableCell>
                                    </TableRow>
                                    <TableRow>
                                        <TableCell variant="head" colSpan={3}>
                                            {intl.formatMessage({ id: 'system.server-architecture' })}
                                        </TableCell>
                                    </TableRow>
                                    <TableRow>
                                        <TableCell colSpan={3}>{runInfo ? runInfo.data.arch : ''}</TableCell>
                                    </TableRow>
                                </TableBody>
                            </Table>
                        </TableContainer>
                    </Card>
                </Grid>
                <Grid paddingTop="1rem">
                    <Card variant="outlined" square style={{ padding: '1rem', height: '100%', display: 'flex', flexDirection: 'column' }}>
                        <Typography variant="h4" sx={{ paddingBottom: '1rem' }}>
                            {intl.formatMessage({ id: 'system.running-info' })}
                        </Typography>
                        <Divider />
                        <TableContainer>
                            <Table className="borderTable">
                                <TableHead>
                                    <TableRow>
                                        <TableCell variant="head" width="25%">
                                            {intl.formatMessage({ id: 'system.go-version' })}
                                        </TableCell>
                                        <TableCell variant="head" width="25%">
                                            {intl.formatMessage({ id: 'system.grata-version' })}
                                        </TableCell>
                                        <TableCell variant="head" width="50%">
                                            {intl.formatMessage({ id: 'system.boot-time' })}
                                        </TableCell>
                                    </TableRow>
                                </TableHead>
                                <TableBody>
                                    <TableRow>
                                        <TableCell>{runInfo ? runInfo.data.version : ''}</TableCell>
                                        <TableCell>{runInfo ? runInfo.data.hgVersion : ''}</TableCell>
                                        <TableCell>{runInfo ? runInfo.data.startTime : ''}</TableCell>
                                    </TableRow>
                                    <TableRow>
                                        <TableCell variant="head">{intl.formatMessage({ id: 'system.runtime' })}</TableCell>
                                        <TableCell variant="head">{intl.formatMessage({ id: 'system.execution-path' })}</TableCell>
                                        <TableCell variant="head">{intl.formatMessage({ id: 'system.goroutine-amount' })}</TableCell>
                                    </TableRow>
                                    <TableRow>
                                        <TableCell>{runInfo ? formatDuration(runInfo.data.runTime) : ''}</TableCell>
                                        <TableCell>{runInfo ? runInfo.data.pwd : ''}</TableCell>
                                        <TableCell>{runInfo ? runInfo.data.goroutine : ''}</TableCell>
                                    </TableRow>
                                    <TableRow>
                                        <TableCell variant="head">{intl.formatMessage({ id: 'system.runtime-memory' })}</TableCell>
                                        <TableCell variant="head" colSpan={2}>
                                            {intl.formatMessage({ id: 'system.hard-disk-used' })}
                                        </TableCell>
                                    </TableRow>
                                    <TableRow>
                                        <TableCell>{runInfo ? runInfo.data.goMem : ''}</TableCell>
                                        <TableCell colSpan={2}>{runInfo ? runInfo.data.goSize : ''}</TableCell>
                                    </TableRow>
                                </TableBody>
                            </Table>
                        </TableContainer>
                    </Card>
                </Grid>
            </>
        );
    };

    return (
        <>
            <MainCard title="System Monitoring" content={false}>
                <CardContent>
                    {handleDataHead()}
                    {handleDataNet()}
                    {handleRunInfo()}
                </CardContent>
            </MainCard>
        </>
    );
};

export default SystemMonitoring;
