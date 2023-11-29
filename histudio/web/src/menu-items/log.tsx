// third-party
import { FormattedMessage } from 'react-intl';

// assets
import { IconList } from '@tabler/icons';

// type
import { NavItemType } from 'types';

const icons = {
    IconList: IconList
};

// ==============================|| MENU ITEMS - DASHBOARD ||============================== //

const myAccount: NavItemType = {
    id: 'log',
    title: <FormattedMessage id="menu-items.log" />,
    type: 'group',
    children: [
        {
            id: 'general-log',
            title: <FormattedMessage id="menu-items.general-log" />,
            type: 'item',
            url: '/log/general-log',
            breadcrumbs: false,
            icon: icons.IconList
        }
    ]
};

export default myAccount;
