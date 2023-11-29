import { useLocation, useNavigate } from 'react-router-dom';

// project imports
import useAuth from 'hooks/useAuth';
import { DASHBOARD_PATH } from 'config';
import { GuardProps } from 'types';
import { useEffect } from 'react';
import LoginRoutes from 'routes/LoginRoutes';

// ==============================|| GUEST GUARD ||============================== //

/**
 * Guest guard for routes having no auth required
 * @param {PropTypes.node} children children element/node
 */

const GuestGuard = ({ children }: GuardProps) => {
    const { isLoggedIn } = useAuth();
    const navigate = useNavigate();
    const location = useLocation();

    useEffect(() => {
        const isLoginPath = LoginRoutes.children.find((obj) => obj.path === location.pathname);
        if (isLoggedIn && isLoginPath !== undefined) {
            navigate(DASHBOARD_PATH, { replace: true });
        } else if (!isLoggedIn && isLoginPath === undefined) {
            navigate('/login', { replace: true });
        }
    }, [isLoggedIn, location]);

    return children;
};

export default GuestGuard;
