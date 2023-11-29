import { Fragment, useEffect, useRef, useState } from 'react';
import { useLocation } from 'react-router-dom';

// material-ui
import { styled, useTheme } from '@mui/material/styles';
import {
    Box,
    // Button,
    ClickAwayListener,
    Divider,
    List,
    ListItemButton,
    ListItemIcon,
    ListItemText,
    Paper,
    Popper,
    Typography,
    useMediaQuery
} from '@mui/material';

// third-party
import { FormattedMessage, useIntl } from 'react-intl';

// project imports
import LAYOUT_CONST from 'constant';
import NavCollapse from '../NavCollapse';
import NavItem from '../NavItem';
import useConfig from 'hooks/useConfig';
import Transitions from 'ui-component/extended/Transitions';
import { dispatch, useSelector } from 'store';

// assets
import { IconChevronDown, IconChevronRight, IconMinusVertical } from '@tabler/icons';
import { NavItemType } from 'types';
import { activeID } from 'store/slices/menu';
// import { ArrowDropDown, ArrowRight } from '@mui/icons-material';

// mini-menu - wrapper
const PopperStyled = styled(Popper)(({ theme }) => ({
    overflow: 'visible',
    zIndex: 1202,
    minWidth: 180,
    '&:before': {
        content: '""',
        display: 'block',
        position: 'absolute',
        top: 5,
        left: 32,
        width: 12,
        height: 12,
        transform: 'translateY(-50%) rotate(45deg)',
        zIndex: 120,
        borderWidth: '6px',
        borderStyle: 'solid',
        borderColor: `${theme.palette.background.paper}  transparent transparent ${theme.palette.background.paper}`
    }
}));

// ==============================|| SIDEBAR MENU LIST GROUP ||============================== //

type VirtualElement = {
    getBoundingClientRect: () => ClientRect | DOMRect;
    contextElement?: Element;
};

interface NavGroupProps {
    item: NavItemType;
    lastItem: number;
    remItems: NavItemType[];
    lastItemId: string;
}

const NavGroup = ({ item, lastItem, remItems, lastItemId }: NavGroupProps) => {
    const theme = useTheme();
    const intl = useIntl();
    const buttonRef = useRef<HTMLButtonElement | null>(null);

    const { pathname } = useLocation();
    const { drawerOpen, selectedID } = useSelector((state) => state.menu);
    const { layout, borderRadius } = useConfig();
    const matchDownMd = useMediaQuery(theme.breakpoints.down('md'));
    const [anchorEl, setAnchorEl] = useState<VirtualElement | (() => VirtualElement) | null | undefined>(null);
    const [currentItem, setCurrentItem] = useState(item);
    const [expanded, setExpanded] = useState(false);

    const openMini = Boolean(anchorEl);

    useEffect(() => {
        if (lastItem) {
            if (item.id === lastItemId) {
                const localItem: any = { ...item };
                const elements = remItems.map((ele: NavItemType) => ele.elements);
                localItem.children = elements.flat(1);
                setCurrentItem(localItem);
            } else {
                setCurrentItem(item);
            }
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [item, lastItem, layout, matchDownMd]);

    useEffect(() => {
        if (expanded) {
            if (buttonRef.current) {
                buttonRef.current.scrollIntoView({ inline: 'center', behavior: 'smooth' });
            }
        }
    }, [expanded]);

    const checkOpenForParent = (child: NavItemType[], id: string) => {
        child.forEach((ele: NavItemType) => {
            if (ele.children?.length) {
                checkOpenForParent(ele.children, currentItem.id!);
            }
            if (ele.url === pathname) {
                dispatch(activeID(id));
            }
        });
    };

    const checkSelectedOnload = (data: NavItemType) => {
        const childrens = data.children ? data.children : [];
        childrens.forEach((itemCheck: NavItemType) => {
            if (itemCheck?.children?.length) {
                checkOpenForParent(itemCheck.children, currentItem.id!);
            }
            if (itemCheck?.url === pathname.replace('/1', '')) {
                dispatch(activeID(currentItem.id!));
                setExpanded(true);
            }
        });
    };

    // keep selected-menu on page load and use for horizontal menu close on change routes
    useEffect(() => {
        setExpanded(false);
        checkSelectedOnload(currentItem);
        if (openMini) setAnchorEl(null);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [pathname, currentItem]);

    const handleClick = (event: React.MouseEvent<HTMLAnchorElement> | React.MouseEvent<HTMLDivElement, MouseEvent> | undefined) => {
        if (!openMini) {
            setAnchorEl(event?.currentTarget);
        }
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    const Icon = currentItem?.icon!;
    const itemIcon = currentItem?.icon ? <Icon stroke={1.5} size="20px" /> : null;

    // menu list collapse & items
    const items = currentItem.children?.map((menu) => {
        switch (menu?.type) {
            case 'collapse':
                return <NavCollapse key={menu.id} menu={menu} level={1} parentId={currentItem.id!} />;
            case 'item':
                return <NavItem key={menu.id} item={menu} level={1} parentId={pathname} />;
            default:
                return (
                    <Typography key={menu?.id} variant="h6" color="error" align="center">
                        {intl.formatMessage({ id: 'menu-items.menu-default-error' })}
                    </Typography>
                );
        }
    });

    const moreItems = remItems.map((itemRem: NavItemType, i) => (
        <Fragment key={i}>
            {itemRem.url ? (
                <NavItem item={itemRem} level={1} />
            ) : (
                itemRem.title && (
                    <Typography variant="caption" sx={{ pl: 2 }}>
                        {itemRem.title} {itemRem.url}
                    </Typography>
                )
            )}
            {itemRem?.elements?.map((menu) => {
                switch (menu?.type) {
                    case 'collapse':
                        return <NavCollapse key={menu.id} menu={menu} level={1} parentId={currentItem.id!} />;
                    case 'item':
                        return <NavItem key={menu.id} item={menu} level={1} parentId={currentItem.id!} />;
                    default:
                        return (
                            <Typography key={menu.id} variant="h6" color="error" align="center">
                                {intl.formatMessage({ id: 'menu-items.menu-default-error' })}
                            </Typography>
                        );
                }
            })}
        </Fragment>
    ));

    const popperId = openMini ? `group-pop-${item.id}` : undefined;

    return (
        <>
            {layout === LAYOUT_CONST.VERTICAL_LAYOUT || (layout === LAYOUT_CONST.HORIZONTAL_LAYOUT && matchDownMd) ? (
                <>
                    <List
                        disablePadding={!drawerOpen}
                        subheader={
                            currentItem.title &&
                            drawerOpen && (
                                <Typography variant="caption" sx={{ ...theme.typography.menuCaption }} display="block" gutterBottom>
                                    {itemIcon && (
                                        <ListItemIcon sx={{ minWidth: 28, verticalAlign: 'text-top' }}>
                                            {currentItem.id === lastItemId ? <IconMinusVertical stroke={1.5} size="20px" /> : itemIcon}
                                        </ListItemIcon>
                                    )}
                                    {currentItem.title}
                                    {currentItem.caption && (
                                        <Typography
                                            variant="caption"
                                            sx={{ ...theme.typography.subMenuCaption }}
                                            display="block"
                                            gutterBottom
                                        >
                                            {currentItem.caption}
                                        </Typography>
                                    )}
                                </Typography>
                                //     <Button
                                //         ref={buttonRef}
                                //         sx={{ width: '100%', justifyContent: 'start' }}
                                //         variant="text"
                                //         onClick={(e) => {
                                //             e.preventDefault();
                                //             setExpanded(!expanded);
                                //         }}
                                //     >
                                //         <Typography variant="caption" sx={{ ...theme.typography.menuCaption }} display="block" gutterBottom>
                                //             {itemIcon && (
                                //                 <ListItemIcon sx={{ minWidth: 36, verticalAlign: 'text-top' }}>
                                //                     {currentItem.id === lastItemId ? <IconMinusVertical stroke={1.5} size="20px" /> : itemIcon}
                                //                 </ListItemIcon>
                                //             )}
                                //             {currentItem.title}
                                //             {currentItem.caption && (
                                //                 <Typography
                                //                     variant="caption"
                                //                     sx={{ ...theme.typography.subMenuCaption }}
                                //                     display="block"
                                //                     gutterBottom
                                //                 >
                                //                     {currentItem.caption}
                                //                 </Typography>
                                //             )}

                                //             {/* {expanded ? (
                                //                 <ArrowDropDown sx={{ verticalAlign: 'bottom' }} />
                                //             ) : (
                                //                 <ArrowRight sx={{ verticalAlign: 'bottom' }} />
                                //             )} */}
                                //         </Typography>
                                //     </Button>
                            )
                        }
                    >
                        {/* {expanded && items} */}
                        {items}
                    </List>

                    {/* group divider */}
                    {drawerOpen && <Divider sx={{ mt: 0.25, mb: 1.25 }} />}
                </>
            ) : (
                <List>
                    <ListItemButton
                        selected={selectedID === currentItem.id}
                        sx={{
                            borderRadius: `${borderRadius}px`,
                            p: 1,
                            my: 0.5,
                            mr: 1,
                            display: 'flex',
                            alignItems: 'center',
                            backgroundColor: 'inherit'
                        }}
                        onMouseEnter={handleClick}
                        onClick={handleClick}
                        onMouseLeave={handleClose}
                        aria-describedby={popperId}
                    >
                        {itemIcon && (
                            <ListItemIcon sx={{ minWidth: 28 }}>
                                {currentItem.id === lastItemId ? <IconMinusVertical stroke={1.5} size="20px" /> : itemIcon}
                            </ListItemIcon>
                        )}
                        <ListItemText
                            sx={{ mr: 1 }}
                            primary={
                                <Typography variant={selectedID === currentItem.id ? 'h5' : 'body1'} color="inherit">
                                    {currentItem.id === lastItemId ? <FormattedMessage id="more-items" /> : currentItem.title}
                                </Typography>
                            }
                        />
                        {openMini ? <IconChevronDown stroke={1.5} size="16px" /> : <IconChevronRight stroke={1.5} size="16px" />}

                        {anchorEl && (
                            <PopperStyled
                                id={popperId}
                                open={openMini}
                                anchorEl={anchorEl}
                                placement="bottom-start"
                                style={{
                                    zIndex: 2001
                                }}
                            >
                                {({ TransitionProps }) => (
                                    <Transitions in={openMini} {...TransitionProps}>
                                        <Paper
                                            sx={{
                                                mt: 0.5,
                                                py: 1.25,
                                                boxShadow: theme.shadows[8],
                                                backgroundImage: 'none'
                                            }}
                                        >
                                            <ClickAwayListener onClickAway={handleClose}>
                                                <Box
                                                    sx={{
                                                        minWidth: 200,
                                                        maxHeight: 'calc(100vh - 170px)',
                                                        overflowY: 'auto',
                                                        '&::-webkit-scrollbar': {
                                                            opacity: 0,
                                                            width: 4,
                                                            '&:hover': {
                                                                opacity: 0.7
                                                            }
                                                        },
                                                        '&::-webkit-scrollbar-track': {
                                                            background: 'transparent'
                                                        },
                                                        '&::-webkit-scrollbar-thumb': {
                                                            background: theme.palette.divider,
                                                            borderRadius: 4
                                                        }
                                                    }}
                                                >
                                                    {currentItem.id !== lastItemId ? items : moreItems}
                                                </Box>
                                            </ClickAwayListener>
                                        </Paper>
                                    </Transitions>
                                )}
                            </PopperStyled>
                        )}
                    </ListItemButton>
                </List>
            )}
        </>
    );
};

export default NavGroup;
