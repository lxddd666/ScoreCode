// third-party
import { FormattedMessage } from 'react-intl';

// assets
import { IconChartTreemap, IconUser, IconSitemap } from '@tabler/icons';

// type
import { NavItemType } from 'types';

const icons = {
    IconUser: IconUser,
    IconChartTreemap: IconChartTreemap,
    IconSitemap: IconSitemap
};

// ==============================|| MENU ITEMS - DASHBOARD ||============================== //

const organization: NavItemType = {
    id: 'organization',
    title: <FormattedMessage id="menu-items.organization" />,
    type: 'group',
    children: [
        {
            id: 'user',
            title: <FormattedMessage id="user.user" />,
            type: 'item',
            url: '/org/user',
            icon: icons.IconUser,
            breadcrumbs: false
        },
        {
            id: 'department',
            title: <FormattedMessage id="profile.department" />,
            type: 'item',
            url: '/org/department',
            icon: icons.IconChartTreemap,
            breadcrumbs: false
        },
        {
            id: 'position',
            title: <FormattedMessage id="user.position" />,
            type: 'item',
            url: '/org/position',
            icon: icons.IconSitemap,
            breadcrumbs: false
        }
    ]
};

export default organization;
