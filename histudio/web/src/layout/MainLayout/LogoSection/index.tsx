import { Link as RouterLink } from 'react-router-dom';

// material-ui
import { Link } from '@mui/material';

// project imports
import { DASHBOARD_PATH } from 'config';

// assets
import logo from '../../../assets/images/general/logo.png';

// env
import envRef from 'environment';

// ==============================|| MAIN LOGO ||============================== //

const LogoSection = () => (
    <Link className="sidebar-menu-logo-link" component={RouterLink} to={DASHBOARD_PATH} aria-label="theme-logo">
        <div className="sidebar-menu-logo-container">
            <img src={logo} width={35} alt={'logo'} />
            <div className="sidebar-menu-title">{envRef?.web?.name}</div>
        </div>
    </Link>
);

export default LogoSection;
