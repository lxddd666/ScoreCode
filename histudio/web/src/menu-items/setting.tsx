// third-party
import { FormattedMessage } from 'react-intl';

// assets
import { IconSettings2 } from '@tabler/icons';

// type
import { NavItemType } from 'types';

const icons = {
    IconSettings: IconSettings2
};

// ==============================|| MENU ITEMS - DASHBOARD ||============================== //

const setting: NavItemType = {
    id: 'setting',
    title: <FormattedMessage id="menu-items.setting" />,
    type: 'group',
    children: [
        {
            id: 'general',
            title: <FormattedMessage id="setting.configuration" />,
            type: 'item',
            url: '/setting/general',
            breadcrumbs: false,
            icon: icons.IconSettings
        },
        {
            id: 'blacklist',
            title: <FormattedMessage id="setting.blacklist.blacklist" />,
            type: 'item',
            url: '/setting/blacklist',
            breadcrumbs: false,
            icon: icons.IconSettings
        }
    ]
};

export default setting;
