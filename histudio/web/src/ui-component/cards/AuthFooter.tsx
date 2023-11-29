// material-ui
import { Link, Typography, Stack } from '@mui/material';

// env
import envRef from 'environment';

// ==============================|| FOOTER - AUTHENTICATION 2 & 3 ||============================== //

const AuthFooter = () => (
    <Stack direction="row" justifyContent="space-between">
        <Typography variant="subtitle2" component={Link} target="_blank" underline="hover"></Typography>
        <Typography variant="subtitle2" component={Link} target="_blank" underline="none">
            &copy; {envRef?.web?.name}
        </Typography>
    </Stack>
);

export default AuthFooter;
