import { useEffect, useState } from 'react';

// material-ui
import { Grid } from '@mui/material';

// project imports
// import EarningCard from '../EarningCard';
import ProjectTaskCard from './ProjectTaskCard';
import AgentCard from '../AgentCard';
// import CommissionCard from '../CommissionCard';
import PopularCard from '../PopularCard';
// import TotalOrderLineChartCard from '../TotalOrderLineChartCard';
// import TotalIncomeDarkCard from '../TotalIncomeDarkCard';
// import TotalIncomeLightCard from '../TotalIncomeLightCard';
import TotalGrowthBarChart from '../TotalGrowthBarChart';
// import LatestCustomerTableCard from './LatestCustomerTableCard'
import { gridSpacing } from 'store/constant';

// ==============================|| DEFAULT DASHBOARD ||============================== //

const Dashboard = () => {
    const [isLoading, setLoading] = useState(true);
    useEffect(() => {
        setLoading(false);
    }, []);

    return (
        <Grid container spacing={gridSpacing}>
            <Grid item lg={12} md={6} sm={6} xs={12}>
                {/* <TotalVisitCard isLoading={isLoading} /> */}
                <ProjectTaskCard />
            </Grid>
            <Grid item xs={8}>
                <Grid container spacing={gridSpacing}>
                    <Grid item lg={12} md={6} sm={6} xs={12}>
                        <AgentCard isLoading={isLoading} />
                    </Grid>
                </Grid>
                <Grid container spacing={gridSpacing} style={{ marginTop: '8px' }}>
                    <Grid item lg={12} xs={2} md={8}>
                        {/* <LatestCustomerTableCard /> */}
                        <TotalGrowthBarChart isLoading={isLoading}/>
                    </Grid>
                </Grid>
            </Grid>
            <Grid item xs={4}>
                <Grid item lg={12} md={6} sm={6} xs={12}>
                    <PopularCard isLoading={isLoading} />
                </Grid>
            </Grid>

        </Grid>
    );
};

export default Dashboard;
