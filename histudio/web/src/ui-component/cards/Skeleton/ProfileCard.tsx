// material-ui
import { Card, CardContent, Grid } from '@mui/material';
import Skeleton from '@mui/material/Skeleton';

// ==============================|| SKELETON - EARNING CARD ||============================== //

const ProfileCard = () => (
    <Card>
        <CardContent>
            <Grid container direction="column">
                <Grid item>
                    <Grid container justifyContent="space-between">
                        <Grid item>
                            <Skeleton variant="rounded" width={44} height={44} />
                        </Grid>
                        <Grid item>
                      
                        </Grid>
                    </Grid>
                </Grid>
                <Grid item>
                    <Skeleton variant="rounded" sx={{ my: 2 }} height={40} />
                </Grid>
                <Grid item>
                    <Skeleton variant="rounded" sx={{ my: 2 }} height={30} />
                </Grid>
                <Grid item>
                    <Skeleton variant="rounded" sx={{ my: 2 }} height={30} />
                </Grid>
                <Grid item>
                    <Skeleton variant="rounded" sx={{ my: 2 }} height={30} />
                </Grid>
                <Grid item>
                    <Skeleton variant="rounded" sx={{ my: 2 }} height={30} />
                </Grid>
            </Grid>
        </CardContent>
    </Card>
);

export default ProfileCard;
