// third-party
import { FormattedMessage } from 'react-intl';

// assets
import { IconDeviceDesktopAnalytics } from '@tabler/icons';
import { IconUsers } from '@tabler/icons';

// type
import { NavItemType } from 'types';

const icons = {
    IconDeviceDesktopAnalytics: IconDeviceDesktopAnalytics,
    IconUsers: IconUsers
};

// ==============================|| MENU ITEMS - DASHBOARD ||============================== //

const system: NavItemType = {
    id: 'system',
    title: <FormattedMessage id="system.system-monitoring" />,
    type: 'group',
    children: [
        {
            id: 'monitoring',
            title: <FormattedMessage id="system.service-monitoring" />,
            type: 'item',
            url: '/system/monitoring',
            breadcrumbs: false,
            icon: icons.IconDeviceDesktopAnalytics
        },
        {
            id: 'online',
            title: <FormattedMessage id="system.online-user" />,
            type: 'item',
            url: '/system/user/online',
            breadcrumbs: false,
            icon: icons.IconUsers
        }
    ]
};

export default system;
