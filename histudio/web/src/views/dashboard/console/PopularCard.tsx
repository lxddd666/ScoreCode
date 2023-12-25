import React from 'react';

// material-ui
import { useTheme } from '@mui/material/styles';
import { Avatar, CardContent, Grid, Menu, MenuItem, Typography } from '@mui/material';

// project imports
import BajajAreaChartCard from './BajajAreaChartCard';
import MainCard from 'ui-component/cards/MainCard';
import SkeletonPopularCard from 'ui-component/cards/Skeleton/PopularCard';
import { gridSpacing } from 'store/constant';

// assets
// import ChevronRightOutlinedIcon from '@mui/icons-material/ChevronRightOutlined';
import MoreHorizOutlinedIcon from '@mui/icons-material/MoreHorizOutlined';
// import KeyboardArrowUpOutlinedIcon from '@mui/icons-material/KeyboardArrowUpOutlined';
// import KeyboardArrowDownOutlinedIcon from '@mui/icons-material/KeyboardArrowDownOutlined';

// ==============================|| DASHBOARD DEFAULT - POPULAR CARD ||============================== //

interface PopularCardProps {
    isLoading: boolean;
}

const PopularCard = ({ isLoading }: PopularCardProps) => {
    const theme = useTheme();

    const [anchorEl, setAnchorEl] = React.useState<Element | ((element: Element) => Element) | null | undefined>(null);

    const top10 = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

    const handleClick = (event: React.SyntheticEvent) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    return (
        <>
            {isLoading ? (
                <SkeletonPopularCard />
            ) : (
                <MainCard content={false}>
                    <CardContent>
                        <Grid container spacing={gridSpacing}>
                            <Grid item xs={12}>
                                <Grid container alignContent="center" justifyContent="space-between">
                                    <Grid item>
                                        <Typography variant="h4">优秀员工（TOP10）</Typography>
                                    </Grid>
                                    <Grid item>
                                        <MoreHorizOutlinedIcon
                                            fontSize="small"
                                            sx={{
                                                color: theme.palette.primary[200],
                                                cursor: 'pointer'
                                            }}
                                            aria-controls="menu-popular-card"
                                            aria-haspopup="true"
                                            onClick={handleClick}
                                        />
                                        <Menu
                                            id="menu-popular-card"
                                            anchorEl={anchorEl}
                                            keepMounted
                                            open={Boolean(anchorEl)}
                                            onClose={handleClose}
                                            variant="selectedMenu"
                                            anchorOrigin={{
                                                vertical: 'bottom',
                                                horizontal: 'right'
                                            }}
                                            transformOrigin={{
                                                vertical: 'top',
                                                horizontal: 'right'
                                            }}
                                        >
                                            <MenuItem onClick={handleClose}> Today</MenuItem>
                                            <MenuItem onClick={handleClose}> This Month</MenuItem>
                                            <MenuItem onClick={handleClose}> This Year </MenuItem>
                                        </Menu>
                                    </Grid>
                                </Grid>
                            </Grid>
                            <Grid item xs={12} sx={{ pt: '16px !important' }}>
                                <BajajAreaChartCard />
                            </Grid>
                            {
                                top10 && top10.map(item => {
                                    return (<Grid item xs={12}>
                                        <div style={{ display: 'flex', flexDirection: 'row' }}>
                                            <Grid item>
                                                <div style={{ width: '50px', display: 'flex', height: '100%', alignItems: 'center' }}> top {item}</div>
                                            </Grid>
                                            <Grid item>
                                                <Avatar
                                                    variant="rounded"
                                                    sx={{
                                                        width: 40,
                                                        height: 40,
                                                        borderRadius: '5px',
                                                        backgroundColor: theme.palette.success.light,
                                                        color: theme.palette.success.dark,
                                                        marginRight: '10px'
                                                        // ml: 2
                                                    }}
                                                    src='http://grata.gen-code.top/grata/attachment/2023-10-24/cwgjs0v4sn0wlocejb.gif'
                                                >

                                                </Avatar>
                                            </Grid>
                                            <Grid container direction="column">
                                                <Grid item>
                                                    <Grid container alignItems="center" justifyContent="space-between">
                                                        <Grid item>
                                                            <Typography variant="subtitle1" color="inherit">
                                                                卢 学东
                                                            </Typography>
                                                        </Grid>
                                                        <Grid item>
                                                            <Grid container alignItems="center" justifyContent="space-between">
                                                                <Grid item>
                                                                    <Typography variant="subtitle1" color="inherit">
                                                                        $1839.00
                                                                    </Typography>
                                                                </Grid>

                                                            </Grid>
                                                        </Grid>
                                                    </Grid>
                                                </Grid>
                                                <Grid item>
                                                    <Typography variant="subtitle2" sx={{ color: 'success.dark' }}>
                                                        10% Profit
                                                    </Typography>
                                                </Grid>
                                            </Grid>

                                        </div>

                                    </Grid>)
                                })
                            }

                        </Grid>
                    </CardContent>
                    {/* <CardActions sx={{ p: 1.25, pt: 0, justifyContent: 'center' }}>
                        <Button size="small" disableElevation>
                            View All
                            <ChevronRightOutlinedIcon />
                        </Button>
                    </CardActions> */}
                </MainCard>
            )}
        </>
    );
};

export default PopularCard;
