//import { lazy } from 'react';
import { lazy, useEffect, useState } from 'react';
import { useRoutes, Navigate, useLocation, RouteObject } from 'react-router-dom';

// routes
import LoginRoutes from './LoginRoutes';
import AuthenticationRoutes from './AuthenticationRoutes';

import { useDynamicRouteMenu, handleRoutes } from 'contexts/DynamicRouteMenuContext';
import useWebSocket from 'hooks/useWebSocket';

import AuthGuard from 'utils/route-guard/AuthGuard';
import MainLayout from 'layout/MainLayout';
import Loadable from 'ui-component/Loadable';

const WhatsChat = Loadable(lazy(() => import('views/whats/whatsAccount/chat')));

// ==============================|| ROUTING RENDER ||============================== //
const hardcodedRoutes: RouteObject = {
    path: '/',
    element: (
        <AuthGuard>
            <MainLayout />
        </AuthGuard>
    ),
    children: [
        {
            path: '/whats/whatsAccount/chat/:id',
            element: <WhatsChat />
        }
    ]
};

const initRouteObjects: RouteObject[] = [
    { path: '/', element: <Navigate to="/login" replace={true} /> },
    // TODO: need to remove this AuthenticationRoutes after changing to the correct path in pages
    AuthenticationRoutes,
    LoginRoutes,
    hardcodedRoutes
];

export default function ThemeRoutes() {
    const location = useLocation();
    const { stopHeartBeat } = useWebSocket();
    const { dynamicRoute } = useDynamicRouteMenu();

    async function stopAllHeartBeatEvent() {
        await stopHeartBeat();
    }

    // Clearing WebSocket Interval Functions on page change
    useEffect(() => {
        stopAllHeartBeatEvent();
    }, [location]);

    const [dynamicRoutes, setDynamicRoutes] = useState<RouteObject[]>([]);
    const [allRoute, setAllRoute] = useState<RouteObject[]>(initRouteObjects);

    const sessionRoute = sessionStorage.getItem('TAB_ROUTES');

    useEffect(() => {
        if (sessionRoute) setAllRoute([...initRouteObjects, ...dynamicRoutes]);
    }, [dynamicRoutes]);

    useEffect(() => {
        sessionRoute
            ? setDynamicRoutes([
                  {
                      path: '/',
                      element: (
                          <AuthGuard>
                              <MainLayout />
                          </AuthGuard>
                      ),
                      children: handleRoutes(JSON.parse(sessionRoute))
                  }
              ])
            : setDynamicRoutes([
                  {
                      path: '/',
                      element: (
                          <AuthGuard>
                              <MainLayout />
                          </AuthGuard>
                      ),
                      children: dynamicRoute
                  }
              ]);
    }, [dynamicRoute]);

    return useRoutes(allRoute);
}
