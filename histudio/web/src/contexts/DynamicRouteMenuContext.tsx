import React, { createContext, useContext, useEffect, useState } from 'react';
import { NavItemType, NavItemTypeObject } from 'types';
import { RouteObject } from 'react-router-dom';
import DynamicLoader from '../utils/general/DynamicLoader';
import axiosServices from 'utils/axios';
import envRef from 'environment';
import useAuth from 'hooks/useAuth';

// Create a context to manage dynamic menu data
const DynamicMenuContext = createContext<any>(undefined);

export function useDynamicRouteMenu() {
    const context = useContext(DynamicMenuContext);
    if (context === undefined) {
        throw new Error('useDynamicRouteMenu must be used within a DynamicRouteMenuProvider');
    }
    return context;
}

const CustomIcon = (type: string) => {
    const icons = require(`@ant-design/icons`);
    const Component = icons[type];
    return <Component />;
};

export function handleMenus(data: any) {
    const menuItems: NavItemType[] = [];
    data.forEach((obj: any, key: number) => {
        const objChildren: NavItemType[] = [];

        obj.children.map((child: any) => {
            if (!child.path.endsWith('/:id?')) {
                objChildren.push({
                    id: child.name,
                    title: child.meta.title,
                    type: 'item',
                    url: `${obj.path}/${child.path}`
                });
            }
        });

        const navItem: NavItemType = {
            id: obj.name,
            title: obj.meta.title,
            type: 'group',
            icon: obj.meta.icon ? () => CustomIcon(obj.meta.icon) : undefined,
            children: objChildren
        };
        menuItems.push(navItem);
    });
    const finalMenu: NavItemTypeObject = { items: menuItems };
    return finalMenu;
}

export function handleRoutes(data: any) {
    const routeItems: RouteObject[] = [];

    function processChildren(children: any[], basePath: string) {
        children.forEach((child: any) => {
            const fullPath = `${basePath}/${child.path}`;

            if (child.component.endsWith('/index')) {
                routeItems.push({
                    path: fullPath,
                    element: <DynamicLoader path={child.component.slice(0, -6)} />
                });
            } else {
                routeItems.push({
                    path: fullPath,
                    element: <DynamicLoader path={child.component} />
                });
            }

            if (child.children) {
                processChildren(child.children, fullPath);
            }
        });
    }

    data.forEach((obj: any) => {
        processChildren(obj.children, obj.path);
    });

    return routeItems;
}

export const DynamicRouteMenuProvider = ({ children }: { children: React.ReactElement }) => {
    const sessionRoute = sessionStorage.getItem('TAB_ROUTES');
    const { isLoggedIn } = useAuth();

    const [dynamicRoute, setDynamicRoute] = useState<RouteObject[]>([]);
    const [dynamicMenu, setDynamicMenu] = useState<NavItemTypeObject | undefined>(
        sessionRoute ? handleMenus(JSON.parse(sessionRoute)) : undefined
    );

    async function getDynamicMenu() {
        if (sessionRoute) {
            setDynamicMenu(handleMenus(JSON.parse(sessionRoute)));
            setDynamicRoute(handleRoutes(JSON.parse(sessionRoute)));
        } else {
            await axiosServices
                .get(`${envRef?.API_URL}admin/role/dynamic`, { headers: {} })
                .then(function (response) {
                    if (response?.data?.code === 0) {
                        sessionStorage.setItem('TAB_ROUTES', JSON.stringify(response.data.data.list));
                        setDynamicMenu(handleMenus(response.data.data.list));
                        setDynamicRoute(handleRoutes(response.data.data.list));
                    }
                })
                .catch(function (error) {
                    console.error(error);
                });
        }
    }

    useEffect(() => {
        if (isLoggedIn) {
            getDynamicMenu();
        } else {
            setDynamicMenu(undefined);
        }
    }, [isLoggedIn]);

    return <DynamicMenuContext.Provider value={{ dynamicMenu, dynamicRoute }}>{children}</DynamicMenuContext.Provider>;
};
