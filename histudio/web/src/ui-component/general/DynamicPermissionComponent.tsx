import React from 'react';
import { useSelector } from 'store';

interface DynamicPermissionComponentProps {
    children: React.ReactNode;
    requiredPermissions?: string[];
}

const DynamicPermissionComponent = ({ children, requiredPermissions }: DynamicPermissionComponentProps) => {
    const { adminInfo } = useSelector((state) => state.user);
    const userPermissions = adminInfo['permissions'];

    let hasRequiredPermissions = false;
    if (userPermissions) {
        if (requiredPermissions) {
            hasRequiredPermissions = requiredPermissions.every((permission) => userPermissions.includes(permission));
        } else {
            hasRequiredPermissions = true;
        }
    }

    return hasRequiredPermissions ? <React.Fragment children={children} /> : null;
};

export default DynamicPermissionComponent;
