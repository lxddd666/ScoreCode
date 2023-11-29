// material-ui
import {
    Chip,
    Divider,
    Grid,
    List,
    ListItemButton,
    ListItemIcon,
    ListItemSecondaryAction,
    ListItemText,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableRow,
    Typography
  } from '@mui/material';
  import { useEffect } from "react";
  
  // project imports
  import Avatar from "ui-component/extended/Avatar";
  import SubCard from "ui-component/cards/SubCard";
  import { gridSpacing } from "store/constant";
  import { useIntl } from "react-intl";
  
  // assets
  import LocationOnIcon from "@mui/icons-material/LocationOn";
  import LoginIcon from "@mui/icons-material/Login";
  import AccessTimeIcon from "@mui/icons-material/AccessTime";
  
  import { getAdminInfo } from "store/slices/user";
  import { useState } from "react";
  import { useSelector, useDispatch } from "store";
  
  import ProfileCard from "ui-component/cards/Skeleton/ProfileCard";
  
  // personal details table
  /** names Don&apos;t look right */
  function createData(
    id: string,
    name: string,
    calories?: string,
    fat?: string,
    carbs?: string,
    protein?: string
  ) {
    return { id, name, calories, fat, carbs, protein };
  }
  
  // ==============================|| PROFILE 1 - PROFILE ||============================== //
  const Profile = () => {
    const [user, setUser]: any = useState({});
    const [rowsData, setRowsData]: any = useState([]);
    const userState = useSelector((state) => state.user);
    const dispatch = useDispatch();
    const intl = useIntl();
  
    const userIdLbl = intl.formatMessage({ id: "profile.userId" });
  
    useEffect(() => {
      dispatch(getAdminInfo(intl));
    }, []);
  
    useEffect(() => {
      const { adminInfo } = userState;
      setUser(adminInfo);
      setRowsData([
        createData(
          "userId",
          intl.formatMessage({ id: "profile.userId" }),
          ":",
          adminInfo?.id
        ),
        createData(
          "userName",
          intl.formatMessage({ id: "profile.userName" }),
          ":",
          adminInfo?.username
        ),
        createData(
          "email",
          intl.formatMessage({ id: "profile.email" }),
          ":",
          adminInfo?.email
        ),
        createData(
          "phone",
          intl.formatMessage({ id: "profile.phone" }),
          ":",
          adminInfo?.mobile
        ),
        createData(
          "balance",
          intl.formatMessage({ id: "profile.balance" }),
          ":",
          adminInfo?.balance ? adminInfo?.balance?.toFixed(2)?.toString() : "0.00"
        ),
        createData(
          "point",
          intl.formatMessage({ id: "profile.point" }),
          ":",
          adminInfo?.integral ? adminInfo?.integral?.toFixed(2)?.toString() : "0.00"
        ),
        createData(
          "department",
          intl.formatMessage({ id: "profile.department" }),
          ":",
          adminInfo?.deptName
        ),
        createData(
          "role",
          intl.formatMessage({ id: "profile.role" }),
          ":",
          adminInfo?.roleName
        ),
      ]);
    }, [userState]);
  
    useEffect(() => {
      const { adminInfo } = userState;
      setRowsData([
        createData(
          "userId",
          intl.formatMessage({ id: "profile.userId" }),
          ":",
          adminInfo?.id
        ),
        createData(
          "userName",
          intl.formatMessage({ id: "profile.userName" }),
          ":",
          adminInfo?.username
        ),
        createData(
          "email",
          intl.formatMessage({ id: "profile.email" }),
          ":",
          adminInfo?.email
        ),
        createData(
          "phone",
          intl.formatMessage({ id: "profile.phone" }),
          ":",
          adminInfo?.mobile
        ),
        createData(
          "balance",
          intl.formatMessage({ id: "profile.balance" }),
          ":",
          adminInfo?.balance ? adminInfo?.balance?.toFixed(2)?.toString() : "0.00"
        ),
        createData(
          "point",
          intl.formatMessage({ id: "profile.point" }),
          ":",
          adminInfo?.integral ? adminInfo?.integral?.toFixed(2)?.toString() : "0.00"
        ),
        createData(
          "department",
          intl.formatMessage({ id: "profile.department" }),
          ":",
          adminInfo?.deptName
        ),
        createData(
          "role",
          intl.formatMessage({ id: "profile.role" }),
          ":",
          adminInfo?.roleName
        ),
      ]);
    }, [userIdLbl]);
  
    return (
      <>
        <Grid container spacing={gridSpacing}>
          <Grid item lg={4} xs={12}>
            {user?.id || user?.realName || user?.roleName ? (
              <SubCard
                title={
                  <Grid container spacing={2} alignItems="center">
                    <Grid item>
                      <Avatar
                        alt="User 1"
                        src={user?.avatar}
                        sx={{ width: 70, height: 70, margin: "0 auto" }}
                      />
                    </Grid>
                    <Grid item xs zeroMinWidth>
                      <Typography align="left" variant="subtitle1">
                        {user?.realName}
                      </Typography>
                      <Typography align="left" variant="subtitle2">
                        {user?.roleName}
                      </Typography>
                    </Grid>
                    {
                      user?.deptName
                      &&
                      <Grid item>
                        <Chip size="small" label={user?.deptName} color="primary" />
                      </Grid>
                    }
                  </Grid>
                }
              >
                <List component="nav" aria-label="main mailbox folders">
                  <ListItemButton>
                    <ListItemIcon>
                      <AccessTimeIcon sx={{ fontSize: "1.3rem" }} />
                    </ListItemIcon>
                    <ListItemText
                      primary={
                        <Typography variant="subtitle1">
                          {intl.formatMessage({ id: "profile.lastLoginAt" })}
                        </Typography>
                      }
                    />
                    <ListItemSecondaryAction>
                      <Typography variant="subtitle2" align="right">
                        {user?.lastLoginAt}
                      </Typography>
                    </ListItemSecondaryAction>
                  </ListItemButton>
                  <Divider />
                  <ListItemButton>
                    <ListItemIcon>
                      <LocationOnIcon sx={{ fontSize: "1.3rem" }} />
                    </ListItemIcon>
                    <ListItemText
                      primary={
                        <Typography variant="subtitle1">
                          {intl.formatMessage({ id: "profile.lastLoginIp" })}
                        </Typography>
                      }
                    />
                    <ListItemSecondaryAction>
                      <Typography variant="subtitle2" align="right">
                        {user?.lastLoginIp}
                      </Typography>
                    </ListItemSecondaryAction>
                  </ListItemButton>
                  <Divider />
                  <ListItemButton>
                    <ListItemIcon>
                      <AccessTimeIcon sx={{ fontSize: "1.3rem" }} />
                    </ListItemIcon>
                    <ListItemText
                      primary={
                        <Typography variant="subtitle1">
                          {intl.formatMessage({ id: "profile.accountCreatedAt" })}
                        </Typography>
                      }
                    />
                    <ListItemSecondaryAction>
                      <Typography variant="subtitle2" align="right">
                        {user?.createdAt}
                      </Typography>
                    </ListItemSecondaryAction>
                  </ListItemButton>
                  <Divider />
                  <ListItemButton>
                    <ListItemIcon>
                      <LoginIcon sx={{ fontSize: "1.3rem" }} />
                    </ListItemIcon>
                    <ListItemText
                      primary={
                        <Typography variant="subtitle1">
                          {intl.formatMessage({ id: "profile.totalLogin" })}
                        </Typography>
                      }
                    />
                    <ListItemSecondaryAction>
                      <Typography variant="subtitle2" align="right">
                        {user?.loginCount}
                      </Typography>
                    </ListItemSecondaryAction>
                  </ListItemButton>
                </List>
              </SubCard>
            ) : (
              <ProfileCard />
            )}
          </Grid>
          <Grid item lg={8} xs={12}>
            {user?.id || user?.realName || user?.roleName ? (
              <Grid container spacing={gridSpacing}>
                <Grid item xs={12}>
                  <SubCard
                    title={intl.formatMessage({ id: "profile.personalInfomation" })}
                  >
                    <Grid container spacing={2}>
                      <Divider sx={{ pt: 1 }} />
                      <Grid item xs={12}>
                        <TableContainer>
                          <Table
                            sx={{
                              "& td": {
                                borderBottom: "none",
                              },
                            }}
                            size="small"
                          >
                            <TableBody>
                              {rowsData.map((rowItem: any) => {
                                return (
                                  <TableRow key={rowItem?.id}>
                                    <TableCell variant="head">
                                      {rowItem.name}
                                    </TableCell>
                                    <TableCell>{rowItem.calories}</TableCell>
                                    <TableCell>{rowItem.fat}</TableCell>
                                  </TableRow>
                                );
                              })}
                            </TableBody>
                          </Table>
                        </TableContainer>
                      </Grid>
                    </Grid>
                  </SubCard>
                </Grid>
              </Grid>
            ) : (
              <ProfileCard />
            )}
          </Grid>
        </Grid>
      </>
    );
  };
  
  export default Profile;
  