import { FlatOption } from 'types/option';

export const gender = [
    { value: 1, label: 'Male' },
    { value: 2, label: 'Female' },
    { value: 3, label: 'Others' }
];

export const defaultErrorMessage = 'auth-register.default-error';

export const statusFlatOptions: FlatOption[] = [
    {
        id: 1,
        name: 'general.normal',
        value: 1
    },
    {
        id: 2,
        name: 'general.disabled',
        value: 0
    }
];

export const statusYesNoFlatOptions: FlatOption[] = [
    {
        id: 1,
        name: 'general.yes',
        value: 1
    },
    {
        id: 2,
        name: 'general.no',
        value: 0
    }
];

export const statusActiveFlatOptions: FlatOption[] = [
    {
        id: 1,
        name: 'general.active',
        value: 1
    },
    {
        id: 2,
        name: 'general.disabled',
        value: 0
    }
];

export const MenuFlatOptions: FlatOption[] = [
    {
        id: 1,
        name: 'general.directory'
    },
    {
        id: 2,
        name: 'general.menu'
    },
    {
        id: 3,
        name: 'general.button'
    }
];

export const WhatsAccountStatusOptions: FlatOption[] = [
    {
        id: 0,
        name: 'general.normal'
    },
    {
        id: 1,
        name: 'general.login-failed'
    },
    {
        id: 2,
        name: 'general.unknown'
    },
    {
        id: 3,
        name: 'general.not-exist'
    },
    {
        id: 403,
        name: 'general.banned'
    },
    {
        id: 401,
        name: 'general.verification-failed'
    }
];

export const WhatsAccountStatusString: Record<number, string> = {
    0: 'general.normal',
    1: 'general.login-failed',
    2: 'general.unknown',
    3: 'general.not-exist',
    403: 'general.banned',
    401: 'general.verification-failed'
};

export enum AccessPermissions {
    'CONSOLE_STAT' = '/console/stat',
    'DASHBOARD_WORKPLACE' = 'dashboard_workplace',
    'DEPT_LIST' = '/dept/list',
    'POST_LIST' = '/post/list',
    'ROLE_LIST' = '/role/list',
    'MEMBER_LIST' = '/member/list',
    'DEPT_OPTION' = '/dept/option',
    'DEPT_ORGOPTION' = '/dept/orgOption',
    'MENU_LIST' = '/menu/list',
    'ROLE_DATASCOPE_SELECT' = '/role/dataScope/select',
    'ROLE_GETPERMISSIONS' = '/role/getPermissions',
    'LOG_LIST' = '/log/list',
    'LOG_VIEW' = '/log/view',
    'LOGINLOG_VIEW' = '/loginLog/view',
    'LOGINLOG_LIST' = '/loginLog/list',
    'SMSLOG_LIST' = '/smsLog/list',
    'SERVELOG_VIEW' = '/serveLog/view',
    'SERVELOG_LIST' = '/serveLog/list',
    'MONITOR_USERONLINELIST' = '/monitor/userOnlineList',
    'NOTICE_LIST' = '/notice/list',
    'ATTACHMENT_LIST' = '/attachment/list',
    'PROVINCES_LIST' = '/provinces/list',
    'PROVINCES_TREE' = '/provinces/tree',
    'PROVINCES_CHILDRENLIST' = '/provinces/childrenList',
    'PROVINCES_UNIQUEID' = '/provinces/uniqueId',
    'GENCODES_LIST' = '/genCodes/list',
    'GENCODES_SELECTS' = '/genCodes/selects',
    'GENCODES_TABLESELECT' = '/genCodes/tableSelect',
    'MEMBER_EDIT' = '/member/edit',
    'MEMBER_VIEW' = '/member/view',
    'MEMBER_DELETE' = '/member/delete',
    'MEMBER_STATUS' = '/member/status',
    'DEPT_EDIT' = '/dept/edit',
    'DEPT_DELETE' = '/dept/delete',
    'DEPT_STATUS' = '/dept/status',
    'POST_EDIT' = '/post/edit',
    'POST_DELETE' = '/post/delete',
    'POST_STATUS' = '/post/status',
    'UPLOAD_FILE' = '/upload/file',
    'HGEXAMPLE_TABLE_LIST' = '/hgexample/table/list',
    'HGEXAMPLE_TABLE_VIEW' = '/hgexample/table/view',
    'LOGINLOG_DELETE' = '/loginLog/delete',
    'LOGINLOG_EXPORT' = '/loginLog/export',
    'SERVELOG_DELETE' = '/serveLog/delete',
    'SERVELOG_EXPORT' = '/serveLog/export',
    'NOTICE_MAXSORT' = '/notice/maxSort',
    'NOTICE_DELETE' = '/notice/delete',
    'NOTICE_STATUS' = '/notice/status',
    'NOTICE_SWITCH' = '/notice/switch',
    'NOTICE_EDITNOTIFY' = '/notice/editNotify',
    'NOTICE_EDITNOTICE' = '/notice/editNotice',
    'NOTICE_EDITLETTER' = '/notice/editLetter',
    'NOTICE_EDIT' = '/notice/edit',
    'NOTICE_MESSAGELIST' = '/notice/messageList',
    'TEST_MAXSORT' = '/test/maxSort',
    'TEST_EXPORT' = '/test/export',
    'TEST_DELETE' = '/test/delete',
    'TEST_STATUS' = '/test/status',
    'TEST_SWITCH' = '/test/switch',
    'TEST_EDIT' = '/test/edit',
    'ADDONS_SELECTS' = '/addons/selects',
    'ADDONS_LIST' = '/addons/list',
    'ORDER_CREATE' = '/order/create',
    'ORDER_OPTION' = '/order/option',
    'CREDITSLOG_EXPORT' = '/creditsLog/export',
    'CREDITSLOG_LIST' = '/creditsLog/list',
    'CREDITSLOG_OPTION' = '/creditsLog/option',
    'ORDER_LIST' = '/order/list',
    'ORDER_ACCEPTREFUND' = '/order/acceptRefund',
    'ORDER_APPLYREFUND' = '/order/applyRefund',
    'ORDER_DELETE' = '/order/delete',
    'CASH_LIST' = '/cash/list',
    'CASH_APPLY' = '/cash/apply',
    'CONFIG_GETCASH' = '/config/getCash',
    'CASH_PAYMENT' = '/cash/payment',
    'CASH_VIEW' = '/cash/view',
    'MEMBER_ADDBALANCE' = '/member/addBalance',
    'MEMBER_ADDINTEGRAL' = '/member/addIntegral',
    'MEMBER_RESETPWD' = '/member/resetPwd',
    'CONFIG_GET' = '/config/get',
    'CONFIG_UPDATE' = '/config/update',
    'DICTTYPE_TREE' = '/dictType/tree',
    'DICTDATA_LIST' = '/dictData/list',
    'CONFIG_TYPESELECT' = '/config/typeSelect',
    'DICTDATA_EDIT' = '/dictData/edit',
    'DICTDATA_DELETE' = '/dictData/delete',
    'DICTTYPE_EDIT' = '/dictType/edit',
    'DICTTYPE_DELETE' = '/dictType/delete',
    'CRON_LIST' = '/cron/list',
    'CRONGROUP_SELECT' = '/cronGroup/select',
    'CRONGROUP_LIST' = '/cronGroup/list',
    'CRON_EDIT' = '/cron/edit',
    'CRON_DELETE' = '/cron/delete',
    'CRON_STATUS' = '/cron/status',
    'CRON_ONLINEEXEC' = '/cron/onlineExec',
    'CRONGROUP_EDIT' = '/cronGroup/edit',
    'CRONGROUP_DELETE' = '/cronGroup/delete',
    'BLACKLIST_LIST' = '/blacklist/list',
    'BLACKLIST_EDIT' = '/blacklist/edit',
    'BLACKLIST_STATUS' = '/blacklist/status',
    'BLACKLIST_DELETE' = '/blacklist/delete',
    'MEMBER_UPDATEPWD' = '/member/updatePwd',
    'MEMBER_UPDATEMOBILE' = '/member/updateMobile',
    'MEMBER_UPDATEEMAIL' = '/member/updateEmail',
    'SERVELICENSE_LIST' = '/serveLicense/list',
    'SERVELICENSE_EDIT' = '/serveLicense/edit',
    'SERVELICENSE_DELETE' = '/serveLicense/delete',
    'SERVELICENSE_STATUS' = '/serveLicense/status',
    'SERVELICENSE_EXPORT' = '/serveLicense/export',
    'SERVELICENSE_ASSIGNROUTER' = '/serveLicense/assignRouter',
    'MEMBER_UPDATEPROFILE' = '/member/updateProfile',
    'MEMBER_UPDATECASH' = '/member/updateCash',
    'ROLE_UPDATEPERMISSIONS' = '/role/updatePermissions',
    'ROLE_DATASCOPE_EDIT' = '/role/dataScope/edit',
    'ROLE_EDIT' = '/role/edit',
    'ROLE_DELETE' = '/role/delete',
    'MENU_EDIT' = '/menu/edit',
    'MENU_DELETE' = '/menu/delete',
    'LOG_DELETE' = '/log/delete',
    'SMSLOG_DELETE' = '/smsLog/delete',
    'MONITOR_USEROFFLINE' = '/monitor/userOffline',
    'PROVINCES_EDIT' = '/provinces/edit',
    'PROVINCES_DELETE' = '/provinces/delete',
    'GENCODES_VIEW' = '/genCodes/view',
    'GENCODES_PREVIEW' = '/genCodes/preview',
    'GENCODES_BUILD' = '/genCodes/build',
    'GENCODES_EDIT' = '/genCodes/edit',
    'ADDONS_UPGRADE' = '/addons/upgrade',
    'ADDONS_UNINSTALL' = '/addons/uninstall',
    'WHATSACCOUNT_LIST' = '/whatsAccount/list',
    'WHATSACCOUNT_VIEW' = '/whatsAccount/view',
    'WHATSACCOUNT_EDIT' = '/whatsAccount/edit',
    'WHATSACCOUNT_DELETE' = '/whatsAccount/delete',
    'WHATSACCOUNT_EXPORT' = '/whatsAccount/export',
    'WHATSPROXY_LIST' = '/whatsProxy/list',
    'WHATSPROXY_VIEW' = '/whatsProxy/view',
    'WHATSPROXY_EDIT' = '/whatsProxy/edit',
    'WHATSPROXY_DELETE' = '/whatsProxy/delete',
    'WHATSPROXY_STATUS' = '/whatsProxy/status',
    'WHATSPROXY_EXPORT' = '/whatsProxy/export',
    'WHATSCUSTOMERS_LIST' = '/whatsCustomers/list',
    'WHATSCUSTOMERS_EDIT' = '/whatsCustomers/edit',
    'WHATSCUSTOMERS_DELETE' = '/whatsCustomers/delete',
    'WHATSCUSTOMERS_EXPORT' = '/whatsCustomers/export',
    'WHATSPOSITION_EDIT' = '/whatsPosition/edit',
    'WHATSPOSITION_MAXSORT' = '/whatsPosition/maxSort',
    'WHATSPOSITION_DELETE' = '/whatsPosition/delete',
    'WHATSPOSITION_STATUS' = '/whatsPosition/status',
    'WHATSPOSITION_EXPORT' = '/whatsPosition/export',
    'WHATSMSG_LIST' = '/whatsMsg/list',
    'WHATSMSG_VIEW' = '/whatsMsg/view',
    'WHATSMSG_EDIT' = '/whatsMsg/edit',
    'WHATSMSG_DELETE' = '/whatsMsg/delete',
    'WHATSMSG_EXPORT' = '/whatsMsg/export',
    'WHATSPROXY_BIND' = '/whatsProxy/bind',
    'WHATSPROXY_UNBIND' = '/whatsProxy/unBind',
    'WHATS_LOGIN' = '/whats/login',
    'WHATS_SENDMSG' = '/whats/sendMsg',
    'WHATSCONTACTS_LIST' = '/whatsContacts/list',
    'WHATSCONTACTS_VIEW' = '/whatsContacts/view',
    'WHATSCONTACTS_EDIT' = '/whatsContacts/edit',
    'WHATSCONTACTS_DELETE' = '/whatsContacts/delete',
    'WHATSCONTACTS_EXPORT' = '/whatsContacts/export',
    'WHATSACCOUNT_BINDMEMBER' = '/whatsAccount/bindMember',
    'WHATSCONTACTS_UPLOAD' = '/whatsContacts/upload',
    'WHATSPROXY_UPLOAD' = '/whatsProxy/upload',
    'WHATSACCOUNT_UPLOAD' = '/whatsAccount/upload',
    'TWO_SCRIPTGROUP_LIST' = '/2/scriptGroup/list',
    'TWO_SCRIPTGROUP_EDIT' = '/2/scriptGroup/edit',
    'TWO_SCRIPTGROUP_DELETE' = '/2/scriptGroup/delete',
    'TWO_SCRIPTGROUP_EXPORT' = '/2/scriptGroup/export',
    'ONE_SCRIPTGROUP_LIST' = '/1/scriptGroup/list',
    'ONE_SCRIPTGROUP_EDIT' = '/1/scriptGroup/edit',
    'ONE_SCRIPTGROUP_DELETE' = '/1/scriptGroup/delete',
    'ONE_SCRIPTGROUP_EXPORT' = '/1/scriptGroup/export',
    'ONE_SYSSCRIPT_LIST' = '/1/sysScript/list',
    'ONE_SYSSCRIPT_VIEW' = '/1/sysScript/view',
    'ONE_SYSSCRIPT_EDIT' = '/1/sysScript/edit',
    'ONE_SYSSCRIPT_DELETE' = '/1/sysScript/delete',
    'ONE_SYSSCRIPT_EXPORT' = '/1/sysScript/export',
    'TWO_SYSSCRIPT_VIEW' = '/2/sysScript/view',
    'TWO_SYSSCRIPT_LIST' = '/2/sysScript/list',
    'TWO_SYSSCRIPT_EDIT' = '/2/sysScript/edit',
    'TWO_SYSSCRIPT_DELETE' = '/2/sysScript/delete',
    'TWO_SYSSCRIPT_EXPORT' = '/2/sysScript/export',
    'TGUSER_LIST' = '/tgUser/list',
    'TGUSER_VIEW' = '/tgUser/view',
    'TGUSER_EDIT' = '/tgUser/edit',
    'TGUSER_DELETE' = '/tgUser/delete',
    'TGUSER_EXPORT' = '/tgUser/export'
}
