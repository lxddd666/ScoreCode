import { memo } from 'react';

// material-ui
import { useTheme } from '@mui/material/styles';
import { Box, Divider, List, Stack, Typography, useMediaQuery } from '@mui/material';

// project imports
import NavItem from './NavItem';
// import NavGroup from './NavGroup';

import { useDynamicRouteMenu } from 'contexts/DynamicRouteMenuContext';
import useConfig from 'hooks/useConfig';
// import { Menu } from 'menu-items/widget';

import { useSelector } from 'store';
import LAYOUT_CONST from 'constant';
import { HORIZONTAL_MAX_ITEM } from 'config';

// types
// import { NavItemType } from 'types';
import { Spin } from 'antd';
import Chip from 'ui-component/extended/Chip';
import NavCollapse from './NavCollapse';

// ==============================|| SIDEBAR MENU LIST ||============================== //

const MenuList = () => {
    const theme = useTheme();
    const { layout } = useConfig();
    const { drawerOpen } = useSelector((state) => state.menu);
    const matchDownMd = useMediaQuery(theme.breakpoints.down('md'));

    const { dynamicMenu } = useDynamicRouteMenu();

    // last menu-item to show in horizontal menu bar
    const lastItem = layout === LAYOUT_CONST.HORIZONTAL_LAYOUT && !matchDownMd ? HORIZONTAL_MAX_ITEM : null;

    let lastItemIndex = (dynamicMenu?.items!.length || 1) - 1;
    // let remItems: NavItemType[] = [];
    let lastItemId: string;

    if (dynamicMenu && lastItem && lastItem < dynamicMenu.items!.length) {
        lastItemId = dynamicMenu.items![lastItem - 1].id!;
        lastItemIndex = lastItem - 1;
        // remItems = dynamicMenu.items!.slice(lastItem - 1, dynamicMenu.items!.length).map((item: any) => ({
        //     title: item.title,
        //     elements: item.children,
        //     icon: item.icon,
        //     ...(item.url && {
        //         url: item.url
        //     })
        // }));
    }

    const navItems = dynamicMenu
        ? dynamicMenu.items!.slice(0, lastItemIndex + 1).map((item: any) => {
              switch (item.type) {
                  case 'group':
                      if (item.url && item.id !== lastItemId) {
                          return (
                              <List key={item.id}>
                                  <NavItem item={item} level={1} isParents />
                                  {layout !== LAYOUT_CONST.HORIZONTAL_LAYOUT && <Divider sx={{ py: 0.5 }} />}
                              </List>
                          );
                      }

                      //   return <NavGroup key={item.id} item={item} lastItem={lastItem!} remItems={remItems} lastItemId={lastItemId} />;
                      return <NavCollapse key={item.id} menu={item} level={1} parentId={item.id!} />;
                  default:
                      return (
                          <Typography key={item.id} variant="h6" color="error" align="center">
                              Menu Items Error
                          </Typography>
                      );
              }
          })
        : null;

    if (dynamicMenu) {
        return layout === LAYOUT_CONST.VERTICAL_LAYOUT || (layout === LAYOUT_CONST.HORIZONTAL_LAYOUT && matchDownMd) ? (
            <>
                <Box {...(drawerOpen && { sx: { mt: 1.5 } })}>{navItems}</Box>
                {layout === LAYOUT_CONST.VERTICAL_LAYOUT && drawerOpen && (
                    <Stack direction="row" justifyContent="center" sx={{ mb: 2 }}>
                        <Chip
                            label={process.env.REACT_APP_VERSION}
                            disabled
                            chipcolor="secondary"
                            size="small"
                            sx={{ cursor: 'pointer' }}
                        />
                    </Stack>
                )}
            </>
        ) : (
            <>
                {navItems}
                <Stack direction="row" justifyContent="center" sx={{ mb: 2 }}>
                    <Chip label={process.env.REACT_APP_VERSION} disabled chipcolor="secondary" size="small" sx={{ cursor: 'pointer' }} />
                </Stack>
            </>
        );
    } else return <Spin style={{ position: 'absolute', top: '50%', left: '50%', zIndex: 1, marginTop: '1rem' }} />;
};

export default memo(MenuList);
