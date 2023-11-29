// third-party
import { FormattedMessage } from 'react-intl';

// assets
import { IconUser, IconInbox } from '@tabler/icons';

// type
import { NavItemType } from 'types';

const icons = {
    IconUser: IconUser,
    IconInbox: IconInbox
};

// ==============================|| MENU ITEMS - DASHBOARD ||============================== //

const myAccount: NavItemType = {
    id: 'account',
    title: <FormattedMessage id="menu-items.my-account" />,
    type: 'group',
    children: [
        {
            id: 'profile',
            title: <FormattedMessage id="menu-items.account-setting" />,
            type: 'item',
            url: '/home/account',
            breadcrumbs: false,
            icon: icons.IconUser
        },
        {
            id: 'inbox',
            title: <FormattedMessage id="menu-items.my-inbox" />,
            type: 'item',
            url: '/home/message',
            breadcrumbs: false,
            icon: icons.IconInbox
        }
    ]
};

export default myAccount;
