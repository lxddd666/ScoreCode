// ==============================|| DASHBOARD - TOTAL GROWTH BAR CHART ||============================== //
import { Props } from 'react-apexcharts';

const chartData: Props = {
    height: 500,
    type: 'bar',
    options: {
        chart: {
            id: 'bar-chart',
            stacked: true,
            toolbar: {
                show: true
            },
            zoom: {
                enabled: true
            }
        },
        responsive: [
            {
                breakpoint: 480,
                options: {
                    legend: {
                        position: 'bottom',
                        offsetX: -10,
                        offsetY: 0
                    }
                }
            }
        ],
        plotOptions: {
            bar: {
                horizontal: false,
                columnWidth: '50%'
            }
        },
        xaxis: {
            type: 'category',
            categories: ['一公司', '二公司', '三公司', '四公司', '五公司', '六公司', '七公司']
        },
        legend: {
            show: true,
            fontFamily: `'Roboto', sans-serif`,
            position: 'bottom',
            offsetX: 20,
            labels: {
                useSeriesColors: false
            },
            markers: {
                width: 16,
                height: 16,
                radius: 5
            },
            itemMargin: {
                horizontal: 15,
                vertical: 8
            }
        },
        fill: {
            type: 'solid'
        },
        dataLabels: {
            enabled: false
        },
        grid: {
            show: true
        }
    },
    series: [
        {
            name: '代理端口',
            data: [35, 125, 35, 35, 35, 80, 35]
        },
        {
            name: '员工数',
            data: [35, 15, 15, 35, 65, 40, 80,]
        },
        {
            name: '账号数',
            data: [35, 145, 35, 35, 20, 105, 100,]
        }
    ]
};
export default chartData;
