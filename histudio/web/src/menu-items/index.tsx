import dashboard from './dashboard';
import organization from './organization'
import setting from './setting';
import myAccount from './myAccount';
import system from './system';
import log from './log';
import { NavItemType } from 'types';

// ==============================|| MENU ITEMS ||============================== //

const menuItems: { items: NavItemType[] } = {
    items: [dashboard, organization, setting, system, myAccount, log]
};

export default menuItems;
