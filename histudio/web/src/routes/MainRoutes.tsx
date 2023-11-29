/*  ==============================|| TODO:PENDING TO REMOVE, NOT USING ANYMORE ||============================== */
import { RouteObject } from 'react-router-dom';

// project imports
import MainLayout from 'layout/MainLayout';

import AuthGuard from 'utils/route-guard/AuthGuard';
/*  ==============================|| CHECKED ||==============================

// const DashboardDefault = Loadable(lazy(() => import('views/dashboard/Default')));
// const User = Loadable(lazy(() => import('views/organization/user')));
// const Department = Loadable(lazy(() => import('views/organization/department')));
// const Position = Loadable(lazy(() => import('views/organization/position')));
// const Security = Loadable(lazy(() => import('views/settings/Security')));
// const BlackList = Loadable(lazy(() => import('views/settings/blacklist')));
// const GeneralLog = Loadable(lazy(() => import('views/logs/general-log')));
// const SystemMonitoring = Loadable(lazy(() => import('views/pages/monitoring/SystemMonitoring')));
// const OnlineUser = Loadable(lazy(() => import('views/pages/monitoring/OnlineUser')));
// const Inbox = Loadable(lazy(() => import('views/pages/inbox')));
// const Profile = Loadable(lazy(() => import('views/pages/profile')));
*/

// ==============================|| MAIN ROUTING ||============================== //
const MainRoutes = (): RouteObject => {
    const routeObject: RouteObject = {
        path: '/',
        element: (
            <AuthGuard>
                <MainLayout />
            </AuthGuard>
        ),
        children: []
        // [
        //     /* ==================== START DASHBOARD  ==================== */
        //     {
        //         path: '/dashboard/console',
        //         element: <DashboardDefault />
        //     },
        //     /* ==================== END DASHBOARD  ==================== */

        //     /* ==================== START TELEGRAM  ==================== */
        //     /* ==================== END TELEGRAM  ==================== */

        //     /* ==================== START WHATSAPP  ==================== */
        //     /* ==================== END WHATSAPP  ==================== */

        //     /* ==================== START ORGANIZATION  ==================== */
        //     {
        //         path: '/org/user',
        //         element: <User />
        //     },
        //     {
        //         path: '/org/dept',
        //         element: <Department />
        //     },
        //     {
        //         path: '/org/post',
        //         element: <Position />
        //     },
        //     /* ==================== END ORGANIZATION  ==================== */

        //     /* ==================== START SYSTEM  ==================== */
        //     {
        //         path: '/system/config',
        //         element: <Security />
        //     },
        //     {
        //         path: '/system/blacklist',
        //         element: <BlackList />
        //     },
        //     /* ==================== END SYSTEM  ==================== */

        //     /* ==================== START PERMISSION  ==================== */
        //     /* ==================== END PERMISSION  ==================== */

        //     /* ==================== START LOG  ==================== */
        //     {
        //         path: '/log/log',
        //         element: <GeneralLog />
        //     },
        //     /* ==================== END LOG  ==================== */

        //     /* ==================== START MONITOR  ==================== */
        //     {
        //         path: '/monitor/serve_monitor',
        //         element: <SystemMonitoring />
        //     },
        //     {
        //         path: '/monitor/online',
        //         element: <OnlineUser />
        //     },
        //     /* ==================== END MONITOR  ==================== */

        //     /* ==================== START ASSET  ==================== */
        //     /* ==================== END ASSET  ==================== */

        //     /* ==================== START SYSTEM UTIL ==================== */
        //     /* ==================== END SYSTEM UTIL ==================== */

        //     /* ==================== START DEV TOOL ==================== */
        //     /* ==================== END DEV TOOL ==================== */

        //     /* ==================== START PLUGIN ==================== */
        //     /* ==================== END PLUGIN ==================== */

        //     /* ==================== START MESSAGE ==================== */
        //     /* ==================== END MESSAGE ==================== */

        //     /* ==================== START HOME ==================== */
        //     {
        //         path: '/home/account',
        //         element: <DynamicLoader path={'pages/profile'} />
        //     },
        //     /* ==================== START SPECIAL HANDLING INBOX  ==================== */
        //     {
        //         path: '/home/message',
        //         element: <Navigate to="/home/message/1" />
        //     },
        //     {
        //         path: '/home/message/:type',
        //         element: <Inbox />
        //     },
        //     /* ==================== END SPECIAL HANDLING INBOX  ==================== */
        //     /* ==================== END HOME ==================== */

        //     /* ==================== START DOCUMENT ==================== */
        //     /* ==================== END DOCUMENT ==================== */

        //     /* ==================== START ABOUT ==================== */
        //     /* ==================== END ABOUT ==================== */

        //     /* ==================== HANDLING PAGE NOT FOUND  ==================== */
        //     {
        //         path: '*',
        //         element: <MaintenanceUnderConstruction />
        //     },
        //     /* ==================== HANDLING RESTRICTED PAGE  ==================== */
        //     {
        //         path: '/restricted',
        //         element: (
        //             <ProtectedRoute element={<Restricted />} requiredPermissions={requiredPermissions} userPermissions={userPermissions} />
        //         )
        //     }
        // ]
    };
    return routeObject;
};

export default MainRoutes;
