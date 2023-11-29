import { Grid } from '@mui/material';
import Skeleton from '@mui/material/Skeleton';
import { gridSpacing } from 'store/constant';

type Prop = {
    rows?: number;
    colsWidth?: number[];
};

const SkeletonLoader = ({ rows = 5, colsWidth = [0.5, 1.5, 10] }: Prop) => {
    const renderSkeleton = () => {
        return Array.from({ length: rows }, (_, index: any) => (
            <Grid key={index} container spacing={gridSpacing} paddingX="1rem">
                {colsWidth.map((width, index: any) => {
                    return (
                        <Grid key={index} item xs={width}>
                            <Skeleton variant="rounded" animation="wave" sx={{ my: 2 }} height={30} />
                        </Grid>
                    );
                })}
            </Grid>
        ));
    };

    return (
        <Grid container spacing={2}>
            {renderSkeleton()}
        </Grid>
    );
};
export default SkeletonLoader;
